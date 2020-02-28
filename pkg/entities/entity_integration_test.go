// +build integration

package entities

import (
	"os"
	"testing"

	"github.com/newrelic/newrelic-client-go/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestIntegrationSearchEntities(t *testing.T) {
	t.Parallel()

	client := newIntegrationTestClient(t)

	params := SearchEntitiesParams{
		Name: "Dummy App",
	}

	actual, err := client.SearchEntities(params)

	require.NoError(t, err)
	require.Greater(t, len(actual), 0)
}

func TestIntegrationSearchEntitiesByTags(t *testing.T) {
	t.Parallel()

	client := newIntegrationTestClient(t)

	params := SearchEntitiesParams{
		Tags: &TagValue{
			Key:   "language",
			Value: "nodejs",
		},
	}

	actual, err := client.SearchEntities(params)

	require.NoError(t, err)
	require.NotNil(t, actual)
}

func TestIntegrationGetEntities(t *testing.T) {
	t.Parallel()

	client := newIntegrationTestClient(t)

	guids := []string{"MjUyMDUyOHxBUE18QVBQTElDQVRJT058MjE1MDM3Nzk1"}
	actual, err := client.GetEntities(guids)

	require.NoError(t, err)
	require.Greater(t, len(actual), 0)
}

func TestIntegrationGetEntity(t *testing.T) {
	t.Parallel()

	client := newIntegrationTestClient(t)

	actual, err := client.GetEntity("MjUyMDUyOHxBUE18QVBQTElDQVRJT058MjE1MDM3Nzk1")

	require.NoError(t, err)
	require.NotNil(t, actual)
}

// nolint
func newIntegrationTestClient(t *testing.T) Entities {
	apiKey := os.Getenv("NEW_RELIC_API_KEY")

	if apiKey == "" {
		t.Skipf("acceptance testing for graphql requires your personal API key")
	}

	return New(config.Config{
		PersonalAPIKey: apiKey,
		UserAgent:      "newrelic/newrelic-client-go",
		LogLevel:       "debug",
	})
}
