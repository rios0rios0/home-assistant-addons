//go:build unit

package controllers_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/infrastructure/controllers"
	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/infrastructure/controllers/mappers"
	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/test/domain/builders"
	doubles "github.com/rios0rios0/home-assistant-addons/mcp-server-extended/test/domain/doubles/repositories"
)

// callTool invokes a registered tool the way a client would, so the tests cover
// registration and argument decoding as well as the handler itself.
func callTool(
	t *testing.T,
	repository *doubles.InMemoryAutomationsRepository,
	name string,
	args map[string]any,
) *mcp.CallToolResult {
	t.Helper()

	clientSession := connect(t, repository)

	result, err := clientSession.CallTool(context.TODO(), &mcp.CallToolParams{Name: name, Arguments: args})
	require.NoError(t, err, "calling %s", name)

	return result
}

func connect(t *testing.T, repository *doubles.InMemoryAutomationsRepository) *mcp.ClientSession {
	t.Helper()

	server := mcp.NewServer(&mcp.Implementation{Name: "test", Version: "test"}, nil)
	controllers.NewAutomationsController(repository, mappers.NewListAutomationsResponseMapper()).Register(server)

	client := mcp.NewClient(&mcp.Implementation{Name: "test-client", Version: "test"}, nil)
	serverTransport, clientTransport := mcp.NewInMemoryTransports()

	serverSession, err := server.Connect(context.TODO(), serverTransport, nil)
	require.NoError(t, err, "connecting server")
	t.Cleanup(func() { _ = serverSession.Close() })

	clientSession, err := client.Connect(context.TODO(), clientTransport, nil)
	require.NoError(t, err, "connecting client")
	t.Cleanup(func() { _ = clientSession.Close() })

	return clientSession
}

func textOf(t *testing.T, result *mcp.CallToolResult) string {
	t.Helper()

	require.NotEmpty(t, result.Content, "expected content in the tool result")
	text, ok := result.Content[0].(*mcp.TextContent)
	require.True(t, ok, "expected text content, got %T", result.Content[0])

	return text.Text
}

func decodeOf(t *testing.T, result *mcp.CallToolResult) map[string]any {
	t.Helper()

	var decoded map[string]any
	require.NoError(t, json.Unmarshal([]byte(textOf(t, result)), &decoded), "decoding tool output")

	return decoded
}

func TestAutomationsControllerRead(t *testing.T) {
	t.Parallel()

	t.Run("should return the count and summaries when listing", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryAutomationsRepository(
			builders.NewAutomationBuilder().WithID("1").WithAlias("First").WithTriggers("noise").Build(),
			builders.NewAutomationBuilder().WithID("2").WithAlias("Second").WithEnabled(false).Build(),
		)

		// when
		result := callTool(t, repository, "list_automations", map[string]any{})

		// then
		decoded := decodeOf(t, result)
		assert.Equal(t, float64(2), decoded["count"])

		automations, ok := decoded["automations"].([]any)
		require.True(t, ok, "expected a summary list")
		require.Len(t, automations, 2)

		first, _ := automations[0].(map[string]any)
		assert.Equal(t, true, first["enabled"], "a missing enabled should default to true")
		assert.NotContains(t, first, "triggers", "listing should summarise, not carry the full body")
	})

	t.Run("should return the full body when getting one automation", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryAutomationsRepository(
			builders.NewAutomationBuilder().WithID("1").WithAlias("First").WithTriggers("noise").Build(),
		)

		// when
		result := callTool(t, repository, "get_automation", map[string]any{"automation_id": "1"})

		// then
		decoded := decodeOf(t, result)
		assert.Equal(t, "First", decoded["alias"])
		assert.Contains(t, decoded, "triggers", "get should return the full automation")
	})

	t.Run("should register every tool", func(t *testing.T) {
		// given
		clientSession := connect(t, doubles.NewInMemoryAutomationsRepository())

		// when
		listed, err := clientSession.ListTools(context.TODO(), nil)

		// then
		require.NoError(t, err)

		found := make(map[string]bool, len(listed.Tools))
		for _, tool := range listed.Tools {
			found[tool.Name] = true
		}

		for _, name := range []string{
			"list_automations", "get_automation", "create_automation", "update_automation",
			"delete_automation", "trigger_automation", "enable_automation", "disable_automation",
		} {
			assert.True(t, found[name], "expected tool %q to be registered", name)
		}
	})
}

func TestAutomationsControllerWrite(t *testing.T) {
	t.Parallel()

	t.Run("should parse the YAML when creating", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryAutomationsRepository()
		document := "alias: From YAML\ntriggers:\n  - trigger: time\n    at: '07:00:00'\n"

		// when
		result := callTool(t, repository, "create_automation", map[string]any{"automation_yaml": document})

		// then
		assert.Equal(t, "created", decodeOf(t, result)["status"])

		stored, err := repository.Get(context.TODO(), "generated")
		require.NoError(t, err, "expected the automation to be stored")
		assert.Equal(t, "From YAML", stored.Alias())
	})

	t.Run("should apply the alias argument when the YAML omits it", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryAutomationsRepository()

		// when
		result := callTool(t, repository, "create_automation", map[string]any{
			"automation_yaml": "triggers: []\n",
			"alias":           "Named by argument",
		})

		// then
		require.False(t, result.IsError, textOf(t, result))
		stored, _ := repository.Get(context.TODO(), "generated")
		assert.Equal(t, "Named by argument", stored.Alias())
	})

	t.Run("should keep the YAML alias over the argument", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryAutomationsRepository()

		// when
		callTool(t, repository, "create_automation", map[string]any{
			"automation_yaml": "alias: From YAML\ntriggers: []\n",
			"alias":           "From argument",
		})

		// then
		stored, _ := repository.Get(context.TODO(), "generated")
		assert.Equal(t, "From YAML", stored.Alias())
	})

	t.Run("should send the parsed document when updating", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryAutomationsRepository(
			builders.NewAutomationBuilder().WithID("5").WithAlias("Before").Build(),
		)

		// when
		result := callTool(t, repository, "update_automation", map[string]any{
			"automation_id":   "5",
			"automation_yaml": "alias: After\n",
		})

		// then
		assert.Equal(t, "updated", decodeOf(t, result)["status"])
		stored, _ := repository.Get(context.TODO(), "5")
		assert.Equal(t, "After", stored.Alias())
	})

	t.Run("should remove the automation when deleting", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryAutomationsRepository(
			builders.NewAutomationBuilder().WithID("5").Build(),
		)

		// when
		result := callTool(t, repository, "delete_automation", map[string]any{"automation_id": "5"})

		// then
		decoded := decodeOf(t, result)
		assert.Equal(t, "deleted", decoded["status"])
		assert.Equal(t, "5", decoded["automation_id"])

		_, err := repository.Get(context.TODO(), "5")
		assert.Error(t, err, "expected the automation to be gone")
	})

	t.Run("should call the repository when triggering", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryAutomationsRepository(
			builders.NewAutomationBuilder().WithID("5").Build(),
		)

		// when
		result := callTool(t, repository, "trigger_automation", map[string]any{"automation_id": "5"})

		// then
		assert.Equal(t, "triggered", decodeOf(t, result)["status"])
		assert.Equal(t, []string{"5"}, repository.Triggered)
	})

	t.Run("should write enabled true when enabling", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryAutomationsRepository(
			builders.NewAutomationBuilder().WithID("5").WithEnabled(false).Build(),
		)

		// when
		result := callTool(t, repository, "enable_automation", map[string]any{"automation_id": "5"})

		// then
		assert.Equal(t, "enabled", decodeOf(t, result)["status"])
		stored, _ := repository.Get(context.TODO(), "5")
		assert.Equal(t, true, stored.IsEnabled())
	})

	t.Run("should preserve the rest of the body when disabling", func(t *testing.T) {
		// given
		// Disabling is a read-modify-write, so the rest of the automation has
		// to survive the round trip rather than being replaced by a stub.
		repository := doubles.NewInMemoryAutomationsRepository(
			builders.NewAutomationBuilder().
				WithID("5").WithAlias("Keep me").WithEnabled(true).WithTriggers("noise").Build(),
		)

		// when
		callTool(t, repository, "disable_automation", map[string]any{"automation_id": "5"})

		// then
		stored, _ := repository.Get(context.TODO(), "5")
		assert.Equal(t, false, stored.IsEnabled())
		assert.Equal(t, "Keep me", stored.Alias())
		assert.Contains(t, stored, "triggers", "triggers should survive the round trip")
	})
}

func TestAutomationsControllerError(t *testing.T) {
	t.Parallel()

	t.Run("should report malformed YAML as an error result", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryAutomationsRepository()

		// when
		result := callTool(t, repository, "create_automation", map[string]any{
			"automation_yaml": "alias: [unclosed\n",
		})

		// then
		require.True(t, result.IsError, "expected an error result for malformed YAML")
		assert.Contains(t, textOf(t, result), "parsing automation YAML")
	})

	t.Run("should report an empty YAML document as an error result", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryAutomationsRepository()

		// when
		result := callTool(t, repository, "create_automation", map[string]any{"automation_yaml": ""})

		// then
		require.True(t, result.IsError, "expected an error result for an empty document")
		assert.Contains(t, textOf(t, result), "document is empty")
	})

	t.Run("should report a repository failure as an error result", func(t *testing.T) {
		// given
		// A failure has to reach the caller as readable tool output, not as a
		// protocol fault that surfaces as a generic error in the client.
		repository := doubles.NewInMemoryAutomationsRepository().
			WithOnError(errors.New("home assistant unreachable"))

		// when
		result := callTool(t, repository, "list_automations", map[string]any{})

		// then
		require.True(t, result.IsError, "expected an error result")
		assert.Contains(t, textOf(t, result), "home assistant unreachable")
	})
}
