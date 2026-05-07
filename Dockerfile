# Build environment for mumbledj - golang alpine container
FROM    golang:1.26.3-alpine@sha256:91eda9776261207ea25fd06b5b7fed8d397dd2c0a283e77f2ab6e91bfa71079d AS builder

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
FROM    alpine:3.23.4@sha256:5b10f432ef3da1b8d4c7eb6c487f2f5a8f096bc91145e68878dd4a5019afde11

# renovate: datasource=repology depName=alpine_3_23/ffmpeg versioning=loose
ARG     FFMPEG_VERSION="8.0.1-r1"
# renovate: datasource=repology depName=alpine_3_23/openssl versioning=loose
ARG     OPENSSL_VERSION="3.5.6-r0"
# renovate: datasource=repology depName=alpine_3_23/aria2 versioning=loose
ARG     ARIA2_VERSION="1.37.0-r1"
# renovate: datasource=repology depName=alpine_3_23/yt-dlp versioning=loose
ARG     YT_DLP_VERSION="2026.03.17-r0"
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
