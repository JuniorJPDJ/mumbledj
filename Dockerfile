# Build environment for mumbledj - golang alpine container
FROM    golang:1.27.2-alpine@sha256:f92b6ef800e499660581efdabdf25d9d817a9d124eaf900924f0504e7e27e12d AS builder

# renovate: datasource=repology depName=alpine_3_24/opus-dev versioning=loose
ARG     OPUS_VERSION="1.6.1-r0"

ARG     branch=master

RUN     apk add --no-cache \
          ca-certificates \
          make \
          git \
          build-base \
          opus-dev=${OPUS_VERSION}

COPY    . $GOPATH/src/github.com/leoverto/mumbledj
WORKDIR $GOPATH/src/github.com/leoverto/mumbledj

RUN     make build install


# Export binary only from builder environment
FROM    alpine:3.24.3@sha256:be2656061b536c1883c4034790e21d0a09c5aaaf51968b89368a3c0371ea2c9a

# renovate: datasource=repology depName=alpine_3_24/ffmpeg versioning=loose
ARG     FFMPEG_VERSION="8.1.3-r0"
# renovate: datasource=repology depName=alpine_3_24/openssl versioning=loose
ARG     OPENSSL_VERSION="3.5.9-r0"
# renovate: datasource=repology depName=alpine_3_24/aria2 versioning=loose
ARG     ARIA2_VERSION="1.37.0-r2"
# renovate: datasource=repology depName=alpine_3_24/yt-dlp versioning=loose
ARG     YT_DLP_VERSION="2026.10.21-r0"
# renovate: datasource=repology depName=alpine_3_24/opus versioning=loose
ARG     OPUS_VERSION="1.6.1-r0"

RUN     apk add --no-cache \
          ffmpeg=${FFMPEG_VERSION} \
          openssl=${OPENSSL_VERSION} \
          aria2=${ARIA2_VERSION} \
          yt-dlp=${YT_DLP_VERSION} \
          opus=${OPUS_VERSION}

COPY    --from=builder /usr/local/bin/mumbledj /usr/local/bin/mumbledj

# Drop to user level privileges
RUN     addgroup -S mumbledj && adduser -S mumbledj -G mumbledj && chmod 750 /home/mumbledj
WORKDIR /home/mumbledj
USER    mumbledj
ENV     HOME=/home/mumbledj

RUN     mkdir -p .config/mumbledj && \
        mkdir -p .cache/mumbledj

ENTRYPOINT ["/usr/local/bin/mumbledj"]
