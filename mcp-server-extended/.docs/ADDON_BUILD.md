# GitHub Actions Setup

This repository includes automated builds via GitHub Actions that build and push Docker images to GitHub Container Registry (GHCR).

## Configuration

- **Registry**: `ghcr.io`
- **Image Name**: `rios0rios0/home-assistant-addons/mcp-server-extended`
- **Image Pattern**: `ghcr.io/rios0rios0/home-assistant-addons/mcp-server-extended/{arch}-addon`

Where `{arch}` is one of: `aarch64`, `armv7`, `amd64`

## Workflow

The workflow (`.github/workflows/build.yaml`) automatically:

1. **Triggers on**:
   - Push to `main` or `master` branch
   - Creation of version tags (e.g., `v0.1.0`)
   - Pull requests (builds only, no push)
   - Manual dispatch via GitHub UI

2. **Builds for all architectures**:
   - `aarch64` (ARM 64-bit)
   - `armv7` (ARM 32-bit)
   - `amd64` (x86_64)

3. **Tags images**:
   - `latest` (only on default branch)
   - Branch name (e.g., `main`)
   - Semantic version tags (e.g., `0.1.0`, `0.1`)

## Setup Instructions

### 1. No Additional Configuration Needed

The workflow automatically uses GitHub's built-in `GITHUB_TOKEN` to authenticate with GitHub Container Registry. No additional secrets are required as GitHub Actions has permissions to publish container images to the repository's container registry by default.

### 2. Push to GitHub

Push your code to trigger the build:

```bash
git add .
git commit -m "Add GitHub Actions workflow"
git push origin main
```

**Note**: The GitHub repository is `rios0rios0/home-assistant-addons`
```

### 3. Monitor Builds

- View build progress: **Actions** tab in your GitHub repository
- Check build logs for any issues
- Images will be pushed to GitHub Container Registry automatically on successful builds

## Image URLs

After a successful build, images will be available at:

- `ghcr.io/rios0rios0/home-assistant-addons/mcp-server-extended/aarch64-addon:latest`
- `ghcr.io/rios0rios0/home-assistant-addons/mcp-server-extended/armv7-addon:latest`
- `ghcr.io/rios0rios0/home-assistant-addons/mcp-server-extended/amd64-addon:latest`

## Version Tags

To create a versioned release:

```bash
git tag -a v0.1.0 -m "Release version 0.1.0"
git push origin v0.1.0
```

This will build and tag images with:
- `ghcr.io/rios0rios0/home-assistant-addons/mcp-server-extended/{arch}-addon:0.1.0`
- `ghcr.io/rios0rios0/home-assistant-addons/mcp-server-extended/{arch}-addon:0.1`

## Troubleshooting

### Build fails with authentication error

- The workflow uses GitHub's built-in `GITHUB_TOKEN` which should work automatically
- Verify the repository has the `packages: write` permission in the workflow file
- Check that the workflow is running on the correct repository

### Build fails with "platform not supported"

- The workflow uses QEMU for cross-platform builds
- This is handled automatically by the `docker/setup-qemu-action`

### Images not appearing in GitHub Container Registry

- Check the Actions tab for build errors
- Verify the workflow completed successfully
- Check the "Packages" section in your GitHub repository

## Local Testing

To test builds locally before pushing:

```bash
./build.sh
```

This will build images locally (without pushing) for testing.

## Related Documentation

- [Addon Installation](ADDON_INSTALLATION.md) - Installing the addon
- [Addon Structure](ADDON_STRUCTURE.md) - Understanding the addon architecture
- [Addon Setup Summary](ADDON_SETUP_SUMMARY.md) - Overview of addon components
- [Documentation Index](SUMMARY.md) - Complete documentation navigation
