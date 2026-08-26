ARG BUILD_FROM

# The binary is cross-compiled on the builder's own architecture rather than
# emulated. This add-on used to be a Python package, and installing it on armv7
# meant compiling `msgpack` from source under QEMU because no musl armv7 wheel
# exists — roughly twenty minutes per build. Go cross-compiles the same target
# in seconds, so `--platform=${BUILDPLATFORM}` pins this stage to the runner and
# GOARCH/GOARM select the target.
FROM --platform=${BUILDPLATFORM} golang:1.27-alpine AS app-build

ARG TARGETARCH
ARG TARGETVARIANT
ARG ADDON_VERSION=dev

WORKDIR /src

# Dependencies first: this layer is cached unless go.mod/go.sum change.
COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN set -x \
    && case "${TARGETVARIANT}" in \
         v7) export GOARM=7 ;; \
         v6) export GOARM=6 ;; \
       esac \
    && CGO_ENABLED=0 GOOS=linux GOARCH="${TARGETARCH}" \
       go build -trimpath \
         -ldflags "-s -w -X main.version=${ADDON_VERSION}" \
         -o /out/mcp-ha-extended ./cmd/mcp-ha-extended

FROM ${BUILD_FROM}

# Default ENV
ENV \
    LANG="C.UTF-8" \
    DEBIAN_FRONTEND="noninteractive" \
    CURL_CA_BUNDLE="/etc/ssl/certs/ca-certificates.crt" \
    S6_BEHAVIOUR_IF_STAGE2_FAILS=2 \
    S6_CMD_WAIT_FOR_SERVICES_MAXTIME=0 \
    S6_CMD_WAIT_FOR_SERVICES=1 \
    S6_SERVICES_READYTIME=50

# Set shell
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# Build Args
ARG \
    BASHIO_VERSION \
    TEMPIO_VERSION \
    S6_OVERLAY_VERSION

# Base system
WORKDIR /usr/src
ARG BUILD_ARCH

# Install system dependencies
# Note: bash, jq, tzdata, curl, ca-certificates are already in base image.
# The application is a statically linked binary, so no language runtime is
# installed here at all — only what S6 and bashio need.
RUN set -x \
    && apk add --no-cache \
         xz \
    && rm -rf /var/cache/apk/*

# Install S6 Overlay, bashio, tempio
RUN set -x \
    && mkdir -p /usr/share/man/man1 \
    && if [ "${BUILD_ARCH}" = "armv7" ]; then \
           export S6_ARCH="arm"; \
       elif [ "${BUILD_ARCH}" = "i386" ]; then \
           export S6_ARCH="i686"; \
       elif [ "${BUILD_ARCH}" = "amd64" ]; then \
           export S6_ARCH="x86_64"; \
       elif [ "${BUILD_ARCH}" = "aarch64" ]; then \
           export S6_ARCH="aarch64"; \
       else \
           export S6_ARCH="${BUILD_ARCH}"; \
       fi \
    && curl --proto '=https' --proto-redir '=https' -L -f -s "https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-${S6_ARCH}.tar.xz" \
       | tar Jxvf - -C / \
    && curl --proto '=https' --proto-redir '=https' -L -f -s "https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-noarch.tar.xz" \
       | tar Jxvf - -C / \
    && curl --proto '=https' --proto-redir '=https' -L -f -s "https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-symlinks-arch.tar.xz" \
       | tar Jxvf - -C / \
    && curl --proto '=https' --proto-redir '=https' -L -f -s "https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-symlinks-noarch.tar.xz" \
       | tar Jxvf - -C / \
    && mkdir -p /etc/fix-attrs.d /etc/services.d \
    && curl --proto '=https' --proto-redir '=https' -L -f -s -o /usr/bin/tempio \
         "https://github.com/home-assistant/tempio/releases/download/${TEMPIO_VERSION}/tempio_${BUILD_ARCH}" \
    && chmod a+x /usr/bin/tempio \
    && mkdir -p /usr/src/bashio \
    && curl --proto '=https' --proto-redir '=https' -L -f -s "https://github.com/hassio-addons/bashio/archive/v${BASHIO_VERSION}.tar.gz" \
       | tar -xzf - --strip 1 -C /usr/src/bashio \
    && mv /usr/src/bashio/lib /usr/lib/bashio \
    && ln -sf /usr/lib/bashio/bashio /usr/bin/bashio \
    && rm -rf /usr/src/bashio

# Install the application
COPY --from=app-build /out/mcp-ha-extended /usr/bin/mcp-ha-extended

# S6-Overlay
WORKDIR /root
ENTRYPOINT ["/init"]

COPY rootfs /
