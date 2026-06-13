# Build environment for mumbledj - golang alpine container
FROM    golang:1.26.4-alpine@sha256:7a3e50096189ad57c9f9f865e7e4aa8585ed1585248513dc5cda498e2f41812c AS builder

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
FROM    alpine:3.24.0@sha256:a2d49ea686c2adfe3c992e47dc3b5e7fa6e6b5055609400dc2acaeb241c829f4

# renovate: datasource=repology depName=alpine_3_24/ffmpeg versioning=loose
ARG     FFMPEG_VERSION="8.1.1-r0"
# renovate: datasource=repology depName=alpine_3_24/openssl versioning=loose
ARG     OPENSSL_VERSION="3.5.7-r0"
# renovate: datasource=repology depName=alpine_3_24/aria2 versioning=loose
ARG     ARIA2_VERSION="1.37.0-r2"
# renovate: datasource=repology depName=alpine_3_24/yt-dlp versioning=loose
ARG     YT_DLP_VERSION="2026.06.09-r0"
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
