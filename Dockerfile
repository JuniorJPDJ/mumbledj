# Build environment for mumbledj - golang alpine container
FROM    golang:1.26.5-alpine@sha256:0178a641fbb4858c5f1b48e34bdaabe0350a330a1b1149aabd498d0699ff5fb2 AS builder

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
FROM    alpine:3.24.1@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b

# renovate: datasource=repology depName=alpine_3_24/ffmpeg versioning=loose
ARG     FFMPEG_VERSION="8.1.2-r0"
# renovate: datasource=repology depName=alpine_3_24/openssl versioning=loose
ARG     OPENSSL_VERSION="3.5.7-r0"
# renovate: datasource=repology depName=alpine_3_24/aria2 versioning=loose
ARG     ARIA2_VERSION="1.37.0-r2"
# renovate: datasource=repology depName=alpine_3_24/yt-dlp versioning=loose
ARG     YT_DLP_VERSION="2026.07.04-r0"
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
