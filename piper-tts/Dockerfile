ARG BUILD_FROM
FROM lscr.io/linuxserver/piper:latest AS app-source

FROM ${BUILD_FROM}

ENV \
    LANG="C.UTF-8" \
    DEBIAN_FRONTEND="noninteractive" \
    CURL_CA_BUNDLE="/etc/ssl/certs/ca-certificates.crt" \
    S6_BEHAVIOUR_IF_STAGE2_FAILS=2 \
    S6_CMD_WAIT_FOR_SERVICES_MAXTIME=0 \
    S6_CMD_WAIT_FOR_SERVICES=1 \
    S6_SERVICES_READYTIME=50

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

ARG \
    BASHIO_VERSION \
    TEMPIO_VERSION \
    S6_OVERLAY_VERSION

WORKDIR /usr/src
ARG BUILD_ARCH

RUN set -x \
    && apk add --no-cache \
         python3 \
         py3-pip \
         xz \
    && rm -rf /var/cache/apk/*

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
    && curl -L -f -s "https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-${S6_ARCH}.tar.xz" \
       | tar Jxvf - -C / \
    && curl -L -f -s "https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-noarch.tar.xz" \
       | tar Jxvf - -C / \
    && curl -L -f -s "https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-symlinks-arch.tar.xz" \
       | tar Jxvf - -C / \
    && curl -L -f -s "https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-symlinks-noarch.tar.xz" \
       | tar Jxvf - -C / \
    && mkdir -p /etc/fix-attrs.d /etc/services.d \
    && curl -L -f -s -o /usr/bin/tempio \
         "https://github.com/home-assistant/tempio/releases/download/${TEMPIO_VERSION}/tempio_${BUILD_ARCH}" \
    && chmod a+x /usr/bin/tempio \
    && mkdir -p /usr/src/bashio \
    && curl -L -f -s "https://github.com/hassio-addons/bashio/archive/v${BASHIO_VERSION}.tar.gz" \
       | tar -xzf - --strip 1 -C /usr/src/bashio \
    && mv /usr/src/bashio/lib /usr/lib/bashio \
    && ln -sf /usr/lib/bashio/bashio /usr/bin/bashio \
    && rm -rf /usr/src/bashio

COPY --from=app-source /app /app

WORKDIR /root
ENTRYPOINT ["/init"]

COPY rootfs /
