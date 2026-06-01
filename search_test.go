// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package emceesprodtesting5_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/emcees-prod-testing-5-go"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/testutil"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
)

func TestSearchAccountsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := emceesprodtesting5.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Search.Accounts(context.TODO(), emceesprodtesting5.SearchAccountsParams{
		Field:    emceesprodtesting5.SearchAccountsParamsFieldAll,
		Query:    "checking",
		Limit:    emceesprodtesting5.Int(10),
		Page:     emceesprodtesting5.Int(1),
		Type:     emceesprodtesting5.AccountTypeFilterAll,
		XTraceID: emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
	})
	if err != nil {
		var apierr *emceesprodtesting5.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchTransactionsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := emceesprodtesting5.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Search.Transactions(context.TODO(), emceesprodtesting5.SearchTransactionsParams{
		Query:    "groceries",
		Limit:    emceesprodtesting5.Int(10),
		Page:     emceesprodtesting5.Int(1),
		XTraceID: emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
	})
	if err != nil {
		var apierr *emceesprodtesting5.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
