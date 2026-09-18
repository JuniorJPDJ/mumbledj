# Build environment for mumbledj - golang alpine container
FROM    golang:1.27.1-alpine@sha256:4cb7ac979db5fcc41cae44b2227ba5ab8a51e8807f40d9ba4dee20a0ad960b5b AS builder

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
FROM    alpine:3.24.2@sha256:294b683cb724975bec92580e1e685676bd4b50bda910ddb8c51d4cabeaec77e6

# renovate: datasource=repology depName=alpine_3_24/ffmpeg versioning=loose
ARG     FFMPEG_VERSION="8.1.2-r0"
# renovate: datasource=repology depName=alpine_3_24/openssl versioning=loose
ARG     OPENSSL_VERSION="3.5.8-r0"
# renovate: datasource=repology depName=alpine_3_24/aria2 versioning=loose
ARG     ARIA2_VERSION="1.37.0-r2"
# renovate: datasource=repology depName=alpine_3_24/yt-dlp versioning=loose
ARG     YT_DLP_VERSION="2026.08.19-r0"
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
