// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly_test

import (
	"context"
	"os"
	"testing"

	"github.com/steel-gareth/firefly-go"
	"github.com/steel-gareth/firefly-go/internal/testutil"
	"github.com/steel-gareth/firefly-go/option"
)

func TestUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := firefly.NewClient(
		option.WithBaseURL(baseURL),
	)
	response, err := client.Autocomplete.ListAccounts(context.TODO(), firefly.AutocompleteListAccountsParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response)
}
