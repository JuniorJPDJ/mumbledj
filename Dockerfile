# Build environment for mumbledj - golang alpine container
FROM    golang:1.25.6-alpine@sha256:bc2596742c7a01aa8c520a075515c7fee21024b05bfaa18bd674fe82c100a05d AS builder

# renovate: datasource=repology depName=alpine_3_23/opus-dev versioning=loose
ARG     OPUS_VERSION="1.5.2-r1"

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
FROM    alpine:3.23.2@sha256:865b95f46d98cf867a156fe4a135ad3fe50d2056aa3f25ed31662dff6da4eb62

# renovate: datasource=repology depName=alpine_3_23/ffmpeg versioning=loose
ARG     FFMPEG_VERSION="8.0.1-r1"
# renovate: datasource=repology depName=alpine_3_23/openssl versioning=loose
ARG     OPENSSL_VERSION="3.5.4-r0"
# renovate: datasource=repology depName=alpine_3_23/aria2 versioning=loose
ARG     ARIA2_VERSION="1.37.0-r1"
# renovate: datasource=repology depName=alpine_3_23/yt-dlp versioning=loose
ARG     YT_DLP_VERSION="2025.12.08-r0"
# renovate: datasource=repology depName=alpine_3_23/opus versioning=loose
ARG     OPUS_VERSION="1.5.2-r1"

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
