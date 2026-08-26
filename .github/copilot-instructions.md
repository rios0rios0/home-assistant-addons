# Repository Instructions for GitHub Copilot

This document provides guidance for GitHub Copilot when working on this repository.

## Repository Overview

This is a **Home Assistant add-ons repository** that provides modular, containerized services for Home Assistant. The repository contains 19 add-ons across six categories:

- **AI & Machine Learning**: `mcp-server-extended`, `ollama-container`, `open-webui`, `localai`, `opencode`, `openclaw`
- **Voice & Speech**: `piper-tts`, `whisper-cpp`
- **Automation & Workflows**: `n8n`, `huginn`, `changedetection`
- **Monitoring & Notifications**: `uptime-kuma`, `ntfy`
- **Productivity & Utilities**: `stirling-pdf`, `it-tools`, `linkwarden`
- **Household Management**: `mealie`, `grocy`, `homebox`

Two add-ons are marked **stable** (`mcp-server-extended`, `ollama-container`); all others are **experimental**.

## Repository Structure

```
home-assistant-addons/
├── .github/
│   ├── copilot-instructions.md     # This file
│   └── workflows/
│       ├── build.yaml              # CI/CD pipeline (change detection + manifest creation)
│       └── _build-addon.yaml       # Reusable per-arch build workflow
├── <addon-name>/                   # One directory per add-on (19 total)
│   ├── config.yaml                 # Add-on metadata (name, version, arch, options…)
│   ├── build.yaml                  # Base images per architecture + build args
│   ├── Dockerfile                  # Multi-arch container with S6 overlay, Bashio, Tempio
│   ├── rootfs/                     # Files overlaid onto the container filesystem
│   │   └── etc/services.d/<slug>/
│   │       ├── run                 # S6 service startup script
│   │       └── finish              # S6 service finish script
│   └── README.md                   # Add-on documentation
├── mcp-server-extended/            # Go add-on (additional files)
│   ├── go.mod                      # Module path and dependencies
│   ├── go.sum                      # Dependency checksums
│   ├── .golangci.yaml               # Linter configuration
│   ├── cmd/mcp-ha-extended/        # Entry point and Dig wiring
│   ├── internal/domain/            # Entity and repository contract
│   ├── internal/infrastructure/    # HA REST client, MCP controller
│   ├── test/domain/                # Builders and in-memory doubles
│   └── .docs/                      # Extended documentation
├── ADDON_IDEAS.md                  # Backlog of potential future add-ons
├── CONTRIBUTING.md                 # Contribution guidelines
├── repository.json                 # Repository metadata (Home Assistant format)
├── README.md                       # Main repository documentation
├── LICENSE                         # MIT License
└── .gitignore                      # Git ignore rules
```

## Home Assistant Add-on Conventions

### Configuration Files

Each add-on directory must contain a `config.yaml` file following the Home Assistant add-on specification:

- **Required fields**: `name`, `version`, `slug`, `description`, `arch`, `url`, `startup`
- **Common optional fields**: `hassio_api`, `homeassistant`, `host_ipc`, `host_network`, `host_dbus`, `image`, `ingress`, `ingress_port`, `ingress_stream`, `init`, `map`, `options`, `schema`, `ports`, `ports_description`, `stage`
- **Version format**: Semantic versioning (e.g., `1.0.0`)
- **Security invariants**: downloads pass `--proto '=https' --proto-redir '=https'`; workflow context values reach `run:` scripts through `env:` rather than `${{ }}` interpolation; every `uses:` is pinned to a full commit SHA (Dependabot bumps them); reusable workflows receive only the secrets they declare.
- **Supported architectures**: `amd64`, `aarch64`, `armv7` (all add-ons support amd64 and aarch64; several also support armv7). An upstream image without a manifest for a target architecture does not force dropping it: arch-neutral payloads are copied from a `FROM --platform=${BUILDPLATFORM}` stage, and native binaries are compiled from source against `${BUILD_FROM}`.
- **Startup types**: `application`, `services`, `system`, `once`
- **Stage values**: `stable`, `experimental` (default to `experimental` for new add-ons)
- **Minimum HA version**: All add-ons require `homeassistant: 2024.6.0`

Example `config.yaml`:
```yaml
version: "1.0.0"
slug: my-addon
name: My Add-on
description: Short description of the add-on
url: https://github.com/rios0rios0/home-assistant-addons/tree/main/my-addon
arch:
  - amd64
  - aarch64
homeassistant: 2024.6.0
hassio_api: false
host_network: false
host_ipc: false
host_dbus: false
image: ghcr.io/rios0rios0/home-assistant-addons/my-addon
ingress: true
ingress_port: 8080
init: false
startup: application
stage: experimental
options:
  my_option: default_value
schema:
  my_option: str
```

### Build Configuration (`build.yaml`)

Each add-on must have a `build.yaml` file specifying architecture-specific base images and standard build arguments:

```yaml
build_from:
  aarch64: ghcr.io/home-assistant/aarch64-base:latest
  amd64: ghcr.io/home-assistant/amd64-base:latest
  armv7: ghcr.io/home-assistant/armv7-base:latest
args:
  BASHIO_VERSION: 0.17.1
  TEMPIO_VERSION: 2024.11.2
  S6_OVERLAY_VERSION: 3.1.6.2
```

### Add-on Documentation (`README.md`)

Each add-on should have comprehensive documentation including:

- Version and architecture badges (shields.io, `style=for-the-badge`)
- About section describing the add-on
- Features list
- Installation instructions
- Configuration examples
- Support information
- License reference

### Repository Metadata (`repository.json`)

The root `repository.json` file contains:

- Repository name
- Repository URL
- Maintainer information

## MCP Server Extended — Key Technologies & Dependencies

- **Go 1.27+**: Minimum required version
- **MCP SDK** (`github.com/modelcontextprotocol/go-sdk`): the official Go SDK for the Model Context Protocol
- **`gopkg.in/yaml.v3`**: YAML parsing for automation configurations
- **`go.uber.org/dig`**: dependency injection wiring
- **`github.com/sirupsen/logrus`**: structured logging (imported as `logger`)
- **`github.com/stretchr/testify`**: test assertions and requirements
- **Go modules**: package and dependency manager

## Development Tools

### Package Management (MCP Server Extended)
- **ALWAYS use Go modules** for dependency management
- Add dependencies: `go get <module>`
- Download dependencies: `go mod download`
- Prune and sync: `go mod tidy`
- Commit both `go.mod` and `go.sum`

### Testing
- Framework: **testify** (`assert` and `require`)
- Every test file carries a `//go:build unit` tag and lives in an external `_test` package
- Unit tests call `t.Parallel()` and group scenarios with `t.Run()`
- Run tests: `go test -tags unit ./...` — **the tag is required**, or packages report "no test files" and pass without testing anything
- Prefer in-memory doubles and builders from `test/domain/` over mocks

### Code Quality
- **gofmt**: Code formatter — `gofmt -l .` lists unformatted files
- **go vet**: `go vet -tags unit ./...`
- **golangci-lint**: Linter, configured in `.golangci.yaml`

## Coding Standards and Best Practices

### General Guidelines

1. **Minimal changes**: Make the smallest possible changes to achieve the goal
2. **Consistency**: Follow existing patterns and conventions in the repository
3. **Documentation**: Update documentation when making changes to add-ons or features
4. **JSON formatting**: Use 2-space indentation for all JSON files
5. **Markdown formatting**: Follow existing markdown style with proper headers and lists

### Go Style (MCP Server Extended)
- File names use `snake_case`
- Method receivers are a one or two letter abbreviation of the type, consistent across the type
- Entities carry **no struct tags** — JSON tags belong on the response DTOs in `internal/infrastructure/controllers/responses/`
- Accept interfaces, return structs
- Use descriptive variable names; formatting is settled by `gofmt`

### I/O Patterns
- Every call that crosses a boundary takes a `context.Context` as its first parameter
- Use the standard library `net/http` client; pass a `*http.Client` in so tests can supply their own
- Wrap errors with `fmt.Errorf("...: %w", err)` so the cause survives
- Log with `logger` (Logrus) to **stderr only** — stdout carries the MCP protocol stream

### MCP Tool Definitions
Each tool follows this pattern:
```go
// The input schema is inferred from the argument struct; `jsonschema` tags
// become the property descriptions an MCP client displays.
type toolArgs struct {
	ParamName string `json:"param_name" jsonschema:"Parameter description"`
}

mcp.AddTool(server, &mcp.Tool{
	Name:        "tool_name",
	Description: "Clear description of what the tool does",
}, handler)
```

### Home Assistant API Patterns
- Base URL: `{HA_URL}/api`
- All requests require: `Authorization: Bearer {HA_TOKEN}`
- Common endpoints:
  - `GET /api/automation` - List automations
  - `GET /api/automation/{id}` - Get automation
  - `POST /api/automation` - Create automation
  - `PUT /api/automation/{id}` - Update automation
  - `DELETE /api/automation/{id}` - Delete automation
  - `POST /api/automation/{id}/trigger` - Trigger automation

## Environment Variables

- `HA_URL`: Home Assistant URL (e.g., `http://homeassistant.local:8123`)
- `HA_TOKEN`: Long-lived access token from Home Assistant
- Both are **required** for the MCP server to function

## Version Management

- Use semantic versioning (MAJOR.MINOR.PATCH)
- Increment PATCH for bug fixes
- Increment MINOR for new features (backward compatible)
- Increment MAJOR for breaking changes
- Update version in both `config.yaml` and `README.md` badges when releasing

## Testing Guidelines

### Unit Tests
- Replace external dependencies with the in-memory double in `test/domain/doubles/`
- Build fixtures with `test/domain/builders/`, not literals
- Test both success and error paths
- Every test file carries `//go:build unit` and lives in an external `_test` package
- Example pattern:
  ```go
  //go:build unit

  func TestFeature(t *testing.T) {
  	t.Parallel()

  	t.Run("should do the thing when the condition holds", func(t *testing.T) {
  		// given
  		repository := doubles.NewInMemoryAutomationsRepository(
  			builders.NewAutomationBuilder().WithID("1").Build(),
  		)

  		// when
  		result, err := functionUnderTest(repository)

  		// then
  		require.NoError(t, err)
  		assert.Equal(t, expected, result)
  	})
  }
  ```

### Manual Testing
- Start the binary with `HA_URL` and `HA_TOKEN` set and drive it from an MCP client
- Check connectivity first: `curl -H "Authorization: Bearer $HA_TOKEN" "$HA_URL/api/"`
- Requires an actual Home Assistant instance

## Docker & Deployment

### Multi-Architecture Support
- Builds for: `aarch64`, `armv7`, `amd64`
- `aarch64` builds run natively on `ubuntu-24.04-arm`; `armv7` cross-builds on `ubuntu-latest` under QEMU. QEMU is registered only when the target platform differs from the runner's own.
- Base images defined in `build.yaml`

### Home Assistant Addon
- Configuration: `config.yaml`
- Startup scripts: `rootfs/etc/services.d/mcp-server-extended/`
- Published to GHCR: `ghcr.io/rios0rios0/home-assistant-addons/mcp-server-extended`

## Common Tasks

### Adding a New Add-on

1. Create a new directory with a descriptive slug name (lowercase, hyphens)
2. Add the standard files: `Dockerfile`, `config.yaml`, `build.yaml`, `rootfs/`, `README.md`
3. Use the standard build args in `build.yaml` (see `build.yaml` for current `BASHIO_VERSION`, `TEMPIO_VERSION`, and `S6_OVERLAY_VERSION` values)
4. Build locally to verify (replace version values with those from the add-on's `build.yaml`):
   ```bash
   docker build \
     --build-arg BUILD_FROM=ghcr.io/home-assistant/amd64-base:latest \
     --build-arg BASHIO_VERSION=<BASHIO_VERSION> \
     --build-arg TEMPIO_VERSION=<TEMPIO_VERSION> \
     --build-arg S6_OVERLAY_VERSION=<S6_OVERLAY_VERSION> \
     --build-arg BUILD_ARCH=amd64 \
     -t local/<addon-name>:test \
     ./<addon-name>
   ```
5. Validate the configuration:
   ```bash
   yq eval '.version' <addon-name>/config.yaml
   yq eval '.arch[]' <addon-name>/config.yaml
   ```
6. Add an entry to the main repository `README.md`
7. Ensure consistency with existing add-ons

### Adding a New MCP Tool

1. Add `Tool` definition in `list_tools()` function
2. Register it in `AutomationsController.Register` with `mcp.AddTool`
3. Add the handler method and, if it needs new data access, a method on the repository contract
4. Add unit tests in `internal/infrastructure/controllers/automations_controller_test.go`
5. Update documentation in `.docs/USAGE_EXAMPLES.md`

### Updating Dependencies

1. `go get <module>@<version>`
2. Prune and sync with `go mod tidy`
3. Run tests: `go test -tags unit ./...`
4. Commit both `go.mod` and `go.sum`

### Debugging

1. Check environment variables are set
2. Test Home Assistant API manually: `curl -H "Authorization: Bearer $HA_TOKEN" $HA_URL/api/`
3. Run server with debug logging
4. Check MCP client logs for protocol errors

## Security Considerations

- Never hardcode `HA_TOKEN` or `HA_URL`
- Always validate user input before sending to Home Assistant
- Use environment variables for sensitive configuration
- Token should have minimal necessary permissions in Home Assistant
- Sanitize YAML input to prevent injection attacks
- Do not commit files containing tokens or secrets (.env, credentials.json, etc.)

## CI/CD

GitHub Actions builds Docker images on every push to `main`, on tag pushes (`v*`), on pull requests targeting `main`, and on manual dispatch.

### Workflow Architecture

- **`build.yaml`** — Orchestration workflow with three jobs:
  1. **`detect-changes`**: Discovers which add-ons changed (via `git diff`), reads their supported architectures from `config.yaml`, and produces a build matrix. On pushes to `main` or tags, all add-ons are rebuilt. If any file under `.github/workflows/` changes, all add-ons are rebuilt.
  2. **`build`**: Fans out to the reusable `_build-addon.yaml` workflow for each `{addon, arch}` pair in the matrix.
  3. **`manifest`**: Creates multi-arch manifests (skipped for PRs). Tags images with the version from `config.yaml` and `latest`. Runs even when the build matrix reports failure and gates per add-on — an add-on publishes only when every architecture in its `arch[]` produced a digest — so one broken add-on blocks only itself and a partial architecture set is never published.

- **`_build-addon.yaml`** — Reusable per-arch build workflow. Accepts a `runner` input so `aarch64` builds land on a native ARM runner. Steps: checkout → arch→platform mapping → conditional QEMU setup → parse `build.yaml` for base images/args → Docker Buildx build → push digest artifact. PRs validate only (no push).
- **`claude.yaml`** — Claude Code agent triggered by issue comments, PR review comments, opened/assigned issues, and submitted PR reviews. Delegates to a reusable workflow in `rios0rios0/.github`.
- **`claude-code-review.yaml`** — Automated PR review via Claude Code on PR open/sync/reopen. Delegates to a reusable workflow in `rios0rios0/.github`.
- **`release.yaml`** — Triggers on push to `main`. Delegates to `rios0rios0/pipelines` to create Git tags when version-bump PRs merge.

### Registry

Images are published to GHCR: `ghcr.io/rios0rios0/home-assistant-addons/<addon-name>`

### Image Tags

- `<version>` (from `config.yaml`)
- `latest`

## Best Practices

1. **Minimal changes**: Make surgical, focused changes
2. **Test-driven**: Write tests before or alongside code changes
3. **Type safety**: Use type hints consistently
4. **Error handling**: Provide clear error messages
5. **Context-first**: Pass `context.Context` through every I/O boundary
6. **Documentation**: Update docs when adding features
7. **Security**: Never commit tokens or secrets
8. **Go modules only**: Do not introduce alternative dependency workflows for MCP Server Extended

## Preferred Solutions

- **HTTP Client**: standard library `net/http`
- **Package Manager**: Go modules
- **Testing**: testify with in-memory doubles (avoid mocks)
- **Formatting**: gofmt
- **Linting**: golangci-lint (see `.golangci.yaml`)
- **Dependency Injection**: Dig, with a `container.go` per layer

## Support and Contribution

- **Issues**: Report bugs and feature requests via GitHub Issues
- **Pull Requests**: Welcome for bug fixes, features, and documentation improvements
- **Maintainer**: @rios0rios0
- **License**: See LICENSE file in repository root

<!-- chlog:start -->
## Changelog (chlog) — MANDATORY

If the repository you are working in uses chlog (a `.chlog.yaml` or `.chlog.yml`
config file, or a `.changes/` directory, exists at the project root), the
following is binding and ALWAYS applies: whenever you make ANY change, you MUST
create a changelog fragment as part of the same change — automatically, without
being asked, before committing.

- Do NOT edit CHANGELOG.md directly; it is generated from fragments.
- Create the fragment with:
  `chlog new --kind <Kind> --body "<imperative description>"`
- Valid kinds: Added, Changed, Deprecated, Removed, Fixed, Security
- Choose the kind that best matches the change (e.g., new feature → Added,
  bug fix → Fixed, behavior change → Changed, removal → Removed, security fix → Security).
- If the change is backward-INCOMPATIBLE with the public API (a breaking
  change), you MUST add the `--breaking` flag:
  `chlog new --kind <Kind> --breaking --body "<description>"`.
  This is the ONLY thing that triggers a major version bump — the kind alone
  never does (per SemVer, major = incompatible change). When unsure whether a
  change breaks compatibility, ask the user instead of guessing.
- Fragments are YAML files in `.changes/unreleased/`; stage them with your commit.
- `chlog check` fails the build when a fragment is missing — never skip it.
<!-- chlog:end -->
