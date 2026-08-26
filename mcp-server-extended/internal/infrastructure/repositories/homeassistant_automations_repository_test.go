//go:build unit

package repositories_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/infrastructure/repositories"
	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/test/domain/builders"
)

// recordedRequest captures what the repository sent, so tests can assert on the
// wire format and not only on the decoded result.
type recordedRequest struct {
	method string
	path   string
	auth   string
	body   string
}

func newServer(t *testing.T, handler http.HandlerFunc) (*httptest.Server, *recordedRequest) {
	t.Helper()

	recorded := &recordedRequest{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)

		recorded.method = r.Method
		recorded.path = r.URL.Path
		recorded.auth = r.Header.Get("Authorization")
		recorded.body = string(body)

		handler(w, r)
	}))
	t.Cleanup(server.Close)

	return server, recorded
}

func respondJSON(payload any) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(payload)
	}
}

func respondEmpty(status int) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(status)
	}
}

func TestHomeAssistantAutomationsRepositoryRead(t *testing.T) {
	t.Parallel()

	t.Run("should list automations from a bare array", func(t *testing.T) {
		// given
		server, recorded := newServer(t, respondJSON([]any{
			map[string]any{"id": "1", "alias": "First"},
			map[string]any{"id": "2", "alias": "Second"},
		}))
		repository := repositories.NewHomeAssistantAutomationsRepository(server.URL, "secret-token", nil)

		// when
		automations, err := repository.List(context.TODO())

		// then
		require.NoError(t, err)
		assert.Len(t, automations, 2)
		assert.Equal(t, http.MethodGet, recorded.method)
		assert.Equal(t, "/api/automation", recorded.path)
		assert.Equal(t, "Bearer secret-token", recorded.auth, "the bearer token should be sent")
	})

	t.Run("should list automations from a wrapped object", func(t *testing.T) {
		// given
		// Older Home Assistant builds wrap the list in an object; both shapes
		// have to keep working.
		server, _ := newServer(t, respondJSON(map[string]any{
			"automations": []any{map[string]any{"id": "1"}},
		}))
		repository := repositories.NewHomeAssistantAutomationsRepository(server.URL, "token", nil)

		// when
		automations, err := repository.List(context.TODO())

		// then
		require.NoError(t, err)
		assert.Len(t, automations, 1)
	})

	t.Run("should return an empty list for an unexpected shape", func(t *testing.T) {
		// given
		server, _ := newServer(t, respondJSON("not-a-list"))
		repository := repositories.NewHomeAssistantAutomationsRepository(server.URL, "token", nil)

		// when
		automations, err := repository.List(context.TODO())

		// then
		require.NoError(t, err)
		assert.Empty(t, automations)
	})

	t.Run("should return the automation when getting one", func(t *testing.T) {
		// given
		server, recorded := newServer(t, respondJSON(map[string]any{"id": "42", "alias": "Answer"}))
		repository := repositories.NewHomeAssistantAutomationsRepository(server.URL, "token", nil)

		// when
		automation, err := repository.Get(context.TODO(), "42")

		// then
		require.NoError(t, err)
		assert.Equal(t, "Answer", automation.Alias())
		assert.Equal(t, "/api/automation/42", recorded.path)
	})

	t.Run("should trim a trailing slash from the configured URL", func(t *testing.T) {
		// given
		// A configured URL ending in "/" would otherwise produce "//api/...",
		// which Home Assistant answers with a redirect.
		server, recorded := newServer(t, respondJSON(map[string]any{"id": "1"}))
		repository := repositories.NewHomeAssistantAutomationsRepository(server.URL+"/", "token", nil)

		// when
		_, err := repository.Get(context.TODO(), "1")

		// then
		require.NoError(t, err)
		assert.Equal(t, "/api/automation/1", recorded.path)
	})
}

func TestHomeAssistantAutomationsRepositoryWrite(t *testing.T) {
	t.Parallel()

	t.Run("should send the automation as JSON when inserting", func(t *testing.T) {
		// given
		server, recorded := newServer(t, respondJSON(map[string]any{"id": "new"}))
		repository := repositories.NewHomeAssistantAutomationsRepository(server.URL, "token", nil)
		automation := builders.NewAutomationBuilder().WithAlias("Created").Build()

		// when
		created, err := repository.Insert(context.TODO(), automation)

		// then
		require.NoError(t, err)
		assert.Equal(t, "new", created.ID())
		assert.Equal(t, http.MethodPost, recorded.method)
		assert.Contains(t, recorded.body, `"alias":"Created"`)
	})

	t.Run("should target the automation path when updating", func(t *testing.T) {
		// given
		server, recorded := newServer(t, respondJSON(map[string]any{"id": "7"}))
		repository := repositories.NewHomeAssistantAutomationsRepository(server.URL, "token", nil)
		automation := builders.NewAutomationBuilder().WithEnabled(true).Build()

		// when
		_, err := repository.Update(context.TODO(), "7", automation)

		// then
		require.NoError(t, err)
		assert.Equal(t, http.MethodPut, recorded.method)
		assert.Equal(t, "/api/automation/7", recorded.path)
	})

	t.Run("should succeed on an empty response when deleting", func(t *testing.T) {
		// given
		// Home Assistant answers several mutations with an empty 200 and no
		// JSON content type; that must not be read as a failure.
		server, recorded := newServer(t, respondEmpty(http.StatusOK))
		repository := repositories.NewHomeAssistantAutomationsRepository(server.URL, "token", nil)

		// when
		err := repository.Delete(context.TODO(), "9")

		// then
		require.NoError(t, err)
		assert.Equal(t, http.MethodDelete, recorded.method)
		assert.Equal(t, "/api/automation/9", recorded.path)
	})

	t.Run("should post to the trigger endpoint when triggering", func(t *testing.T) {
		// given
		server, recorded := newServer(t, respondEmpty(http.StatusOK))
		repository := repositories.NewHomeAssistantAutomationsRepository(server.URL, "token", nil)

		// when
		err := repository.Trigger(context.TODO(), "3")

		// then
		require.NoError(t, err)
		assert.Equal(t, http.MethodPost, recorded.method)
		assert.Equal(t, "/api/automation/3/trigger", recorded.path)
	})
}

func TestHomeAssistantAutomationsRepositoryError(t *testing.T) {
	t.Parallel()

	t.Run("should report a non-success status with its body", func(t *testing.T) {
		// given
		server, _ := newServer(t, func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("bad token"))
		})
		repository := repositories.NewHomeAssistantAutomationsRepository(server.URL, "token", nil)

		// when
		_, err := repository.Get(context.TODO(), "1")

		// then
		require.Error(t, err)
		assert.Contains(t, err.Error(), "401")
		assert.Contains(t, err.Error(), "bad token")
	})

	t.Run("should report an undecodable body", func(t *testing.T) {
		// given
		server, _ := newServer(t, func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte("{not json"))
		})
		repository := repositories.NewHomeAssistantAutomationsRepository(server.URL, "token", nil)

		// when
		_, err := repository.Get(context.TODO(), "1")

		// then
		require.Error(t, err)
		assert.Contains(t, err.Error(), "decoding response")
	})
}
