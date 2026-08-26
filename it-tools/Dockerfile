ARG BUILD_FROM
# The payload copied out of this stage is the compiled Vue SPA (HTML/CSS/JS) —
# no native code — so the stage is pinned to the builder's platform. Upstream
# publishes amd64/arm64 only; without the pin, buildx resolves this stage for
# the target platform and armv7 fails with "no match for platform in manifest".
FROM --platform=${BUILDPLATFORM} corentinth/it-tools:latest AS app-source

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
RUN set -x \
    && apk add --no-cache \
         nginx \
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

# Copy static files from IT Tools image
COPY --from=app-source /usr/share/nginx/html /var/www/html

# S6-Overlay
WORKDIR /root
ENTRYPOINT ["/init"]

COPY rootfs /
