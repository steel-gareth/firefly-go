// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/emcees-prod-testing-5-go"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/testutil"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
)

func TestBudgetLimitNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Budgets.Limits.New(
		context.TODO(),
		"123",
		firefly.BudgetLimitNewParams{
			Amount:       "123.45",
			End:          time.Now(),
			Start:        time.Now(),
			CurrencyCode: firefly.String("EUR"),
			CurrencyID:   firefly.String("5"),
			FireWebhooks: firefly.Bool(true),
			Notes:        firefly.String("Some example notes"),
			XTraceID:     firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
		},
	)
	if err != nil {
		var apierr *firefly.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetLimitGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Budgets.Limits.Get(
		context.TODO(),
		1,
		firefly.BudgetLimitGetParams{
			ID:       "123",
			XTraceID: firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
		},
	)
	if err != nil {
		var apierr *firefly.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetLimitUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Budgets.Limits.Update(
		context.TODO(),
		"123",
		firefly.BudgetLimitUpdateParams{
			ID:           "123",
			Amount:       firefly.String("123.45"),
			CurrencyCode: firefly.String("EUR"),
			CurrencyID:   firefly.String("5"),
			CurrencyName: firefly.String("Euro"),
			End:          firefly.Time(time.Now()),
			FireWebhooks: firefly.Bool(true),
			Notes:        firefly.String("Some example notes"),
			Start:        firefly.Time(time.Now()),
			XTraceID:     firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
		},
	)
	if err != nil {
		var apierr *firefly.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetLimitDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Budgets.Limits.Delete(
		context.TODO(),
		"123",
		firefly.BudgetLimitDeleteParams{
			ID:       "123",
			XTraceID: firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
		},
	)
	if err != nil {
		var apierr *firefly.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetLimitList0WithOptionalParams(t *testing.T) {
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
	_, err := client.Budgets.Limits.List0(
		context.TODO(),
		"123",
		firefly.BudgetLimitList0Params{
			End:      firefly.Time(time.Now()),
			Start:    firefly.Time(time.Now()),
			XTraceID: firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
		},
	)
	if err != nil {
		var apierr *firefly.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetLimitList1WithOptionalParams(t *testing.T) {
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
	_, err := client.Budgets.Limits.List1(context.TODO(), firefly.BudgetLimitList1Params{
		End:      time.Now(),
		Start:    time.Now(),
		XTraceID: firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
	})
	if err != nil {
		var apierr *firefly.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetLimitListTransactionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Budgets.Limits.ListTransactions(
		context.TODO(),
		"123",
		firefly.BudgetLimitListTransactionsParams{
			ID:       "123",
			Limit:    firefly.Int(10),
			Page:     firefly.Int(1),
			Type:     firefly.TransactionTypeFilterAll,
			XTraceID: firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
		},
	)
	if err != nil {
		var apierr *firefly.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
