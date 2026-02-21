#!/bin/bash
# Build script for MCP Server Extended Home Assistant Addon

set -e

# Configuration
ADDON_NAME="mcp-server-extended"
VERSION="0.1.0"
REGISTRY="rios0rios0"

# Build arguments
BASHIO_VERSION="0.17.1"
TEMPIO_VERSION="2024.11.2"
S6_OVERLAY_VERSION="3.1.6.2"

# Base images
BASE_AARCH64="ghcr.io/home-assistant/aarch64-base:latest"
BASE_ARMV7="ghcr.io/home-assistant/armv7-base:latest"
BASE_AMD64="ghcr.io/home-assistant/amd64-base:latest"

# Function to build for a specific architecture
build_arch() {
    local arch=$1
    local base_image=$2

    echo "Building for ${arch}..."

    docker buildx build \
        --platform linux/${arch} \
        --build-arg BUILD_FROM="${base_image}" \
        --build-arg BUILD_ARCH="${arch}" \
        --build-arg S6_OVERLAY_VERSION="${S6_OVERLAY_VERSION}" \
        --build-arg TEMPIO_VERSION="${TEMPIO_VERSION}" \
        --build-arg BASHIO_VERSION="${BASHIO_VERSION}" \
        -t "${REGISTRY}/${arch}-addon-${ADDON_NAME}:${VERSION}" \
        -t "${REGISTRY}/${arch}-addon-${ADDON_NAME}:latest" \
        --load \
        .
}

# Build for all architectures
echo "Building MCP Server Extended Addon v${VERSION}"

# Uncomment the architectures you want to build
# build_arch "aarch64" "${BASE_AARCH64}"
# build_arch "armv7" "${BASE_ARMV7}"
# build_arch "amd64" "${BASE_AMD64}"

echo "Build complete!"
