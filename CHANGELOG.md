# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

When a new release is proposed:

1. Create a new branch `bump/x.x.x` (this isn't a long-lived branch!!!);
2. The Unreleased section on `CHANGELOG.md` gets a version number and date;
3. Open a Pull Request with the bump version changes targeting the `main` branch;
4. When the Pull Request is merged, a new Git tag must be created using <LINK TO THE PLATFORM TO OPEN THE PULL REQUEST>.

Releases to productive environments should run from a tagged version.
Exceptions are acceptable depending on the circumstances (critical bug fixes that can be cherry-picked, etc.).

## [Unreleased]

### Fixed

- repaired the `Build and Push Docker Images` workflow, which had been failing on every commit since 2026-04-17:
  - dropped `aarch64` from `whisper-cpp/config.yaml` (`ghcr.io/ggml-org/whisper.cpp:main` only publishes an `amd64` manifest)
  - dropped `armv7` from `ollama-container/config.yaml` (`ollama/ollama:latest` does not publish an `armv7` manifest — Ollama supports `amd64` and `aarch64` only)
  - dropped `armv7` from `it-tools/config.yaml` (`corentinth/it-tools:latest` does not publish an `armv7` manifest)
  - dropped `armv7` from `grocy/config.yaml` (`lscr.io/linuxserver/grocy:latest` does not publish an `armv7` manifest)
  - bumped `opencode/build.yaml` `OPENCODE_VERSION` from `0.1.0` (does not exist) to `0.0.55` (latest published release)
  - corrected `opencode/Dockerfile` release asset URL pattern from `opencode_Linux_${OC_ARCH}.tar.gz` to `opencode-linux-${OC_ARCH}.tar.gz` (upstream renamed assets) and remapped `aarch64` → `arm64` to match the new asset naming

## [0.2.1] - 2026-04-28

### Changed

- refreshed `CLAUDE.md` and `.github/copilot-instructions.md` to document the `claude.yaml` and `claude-code-review.yaml` workflows

## [0.2.0] - 2026-04-17

### Added

- added `n8n/workflows/gg-incident.json` and companion `README.md` -- importable n8n workflow that receives GitGuardian webhook incidents and forwards them to Telegram via a bot

## [0.1.0] - 2026-03-24

The changes weren't tracked until this version.
