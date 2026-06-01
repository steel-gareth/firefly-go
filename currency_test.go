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

func TestCurrencyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Currencies.New(context.TODO(), firefly.CurrencyNewParams{
		Code:          "AMS",
		Name:          "Ankh-Morpork dollar",
		Symbol:        "AM$",
		DecimalPlaces: firefly.Int(2),
		Enabled:       firefly.Bool(true),
		Primary:       firefly.Bool(true),
		XTraceID:      firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
	})
	if err != nil {
		var apierr *firefly.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCurrencyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Currencies.Get(
		context.TODO(),
		"USD",
		firefly.CurrencyGetParams{
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

func TestCurrencyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Currencies.Update(
		context.TODO(),
		"EUR",
		firefly.CurrencyUpdateParams{
			Code:          firefly.String("AMS"),
			DecimalPlaces: firefly.Int(2),
			Enabled:       firefly.Bool(true),
			Name:          firefly.String("Ankh-Morpork dollar"),
			Primary:       true,
			Symbol:        firefly.String("AM$"),
			XTraceID:      firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
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

func TestCurrencyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Currencies.List(context.TODO(), firefly.CurrencyListParams{
		Limit:    firefly.Int(10),
		Page:     firefly.Int(1),
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

func TestCurrencyDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Currencies.Delete(
		context.TODO(),
		"GBP",
		firefly.CurrencyDeleteParams{
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

func TestCurrencyDisableWithOptionalParams(t *testing.T) {
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
	_, err := client.Currencies.Disable(
		context.TODO(),
		"GBP",
		firefly.CurrencyDisableParams{
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

func TestCurrencyEnableWithOptionalParams(t *testing.T) {
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
	_, err := client.Currencies.Enable(
		context.TODO(),
		"USD",
		firefly.CurrencyEnableParams{
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

func TestCurrencyListAccountsWithOptionalParams(t *testing.T) {
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
	_, err := client.Currencies.ListAccounts(
		context.TODO(),
		"USD",
		firefly.CurrencyListAccountsParams{
			Date:     firefly.Time(time.Now()),
			Limit:    firefly.Int(10),
			Page:     firefly.Int(1),
			Type:     firefly.AccountTypeFilterAll,
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

func TestCurrencyListAvailableBudgetsWithOptionalParams(t *testing.T) {
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
	_, err := client.Currencies.ListAvailableBudgets(
		context.TODO(),
		"EUR",
		firefly.CurrencyListAvailableBudgetsParams{
			Limit:    firefly.Int(10),
			Page:     firefly.Int(1),
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

func TestCurrencyListBillsWithOptionalParams(t *testing.T) {
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
	_, err := client.Currencies.ListBills(
		context.TODO(),
		"USD",
		firefly.CurrencyListBillsParams{
			Limit:    firefly.Int(10),
			Page:     firefly.Int(1),
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

func TestCurrencyListBudgetLimitsWithOptionalParams(t *testing.T) {
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
	_, err := client.Currencies.ListBudgetLimits(
		context.TODO(),
		"USD",
		firefly.CurrencyListBudgetLimitsParams{
			End:      firefly.Time(time.Now()),
			Limit:    firefly.Int(10),
			Page:     firefly.Int(1),
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

func TestCurrencyListRecurrencesWithOptionalParams(t *testing.T) {
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
	_, err := client.Currencies.ListRecurrences(
		context.TODO(),
		"EUR",
		firefly.CurrencyListRecurrencesParams{
			Limit:    firefly.Int(10),
			Page:     firefly.Int(1),
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

func TestCurrencyListRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Currencies.ListRules(
		context.TODO(),
		"USD",
		firefly.CurrencyListRulesParams{
			Limit:    firefly.Int(10),
			Page:     firefly.Int(1),
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

func TestCurrencyListTransactionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Currencies.ListTransactions(
		context.TODO(),
		"USD",
		firefly.CurrencyListTransactionsParams{
			End:      firefly.Time(time.Now()),
			Limit:    firefly.Int(10),
			Page:     firefly.Int(1),
			Start:    firefly.Time(time.Now()),
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
