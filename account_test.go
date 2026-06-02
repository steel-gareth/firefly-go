// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/steel-gareth/firefly-go"
	"github.com/steel-gareth/firefly-go/internal/testutil"
	"github.com/steel-gareth/firefly-go/option"
)

func TestAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.New(context.TODO(), firefly.AccountNewParams{
		Name:               "My checking account",
		Type:               firefly.ShortAccountTypePropertyAsset,
		AccountNumber:      firefly.String("7009312345678"),
		AccountRole:        firefly.AccountRolePropertyDefaultAsset,
		Active:             firefly.Bool(false),
		Bic:                firefly.String("BOFAUS3N"),
		CreditCardType:     firefly.CreditCardTypePropertyMonthlyFull,
		CurrencyCode:       firefly.String("EUR"),
		CurrencyID:         firefly.String("12"),
		Iban:               firefly.String("GB98MIDL07009312345678"),
		IncludeNetWorth:    firefly.Bool(true),
		Interest:           firefly.String("5.3"),
		InterestPeriod:     firefly.InterestPeriodPropertyMonthly,
		Latitude:           firefly.Float(51.983333),
		LiabilityDirection: firefly.LiabilityDirectionPropertyCredit,
		LiabilityType:      firefly.LiabilityTypePropertyLoan,
		Longitude:          firefly.Float(5.916667),
		MonthlyPaymentDate: firefly.Time(time.Now()),
		Notes:              firefly.String("Some example notes"),
		OpeningBalance:     firefly.String("-1012.12"),
		OpeningBalanceDate: firefly.Time(time.Now()),
		Order:              firefly.Int(1),
		VirtualBalance:     firefly.String("123.45"),
		ZoomLevel:          firefly.Int(6),
		XTraceID:           firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
	})
	if err != nil {
		var apierr *firefly.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Get(
		context.TODO(),
		"123",
		firefly.AccountGetParams{
			Date:     firefly.Time(time.Now()),
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

func TestAccountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Update(
		context.TODO(),
		"123",
		firefly.AccountUpdateParams{
			Name:               "My checking account",
			Type:               map[string]any{},
			AccountNumber:      firefly.String("7009312345678"),
			AccountRole:        firefly.AccountRolePropertyDefaultAsset,
			Active:             firefly.Bool(false),
			Bic:                firefly.String("BOFAUS3N"),
			CreditCardType:     firefly.CreditCardTypePropertyMonthlyFull,
			CurrencyCode:       firefly.String("EUR"),
			CurrencyID:         firefly.String("12"),
			Iban:               firefly.String("GB98MIDL07009312345678"),
			IncludeNetWorth:    firefly.Bool(true),
			Interest:           firefly.String("5.3"),
			InterestPeriod:     firefly.InterestPeriodPropertyMonthly,
			Latitude:           firefly.Float(51.983333),
			LiabilityType:      firefly.LiabilityTypePropertyLoan,
			Longitude:          firefly.Float(5.916667),
			MonthlyPaymentDate: firefly.Time(time.Now()),
			Notes:              firefly.String("Some example notes"),
			OpeningBalance:     firefly.String("-1012.12"),
			OpeningBalanceDate: firefly.Time(time.Now()),
			Order:              firefly.Int(1),
			VirtualBalance:     firefly.String("123.45"),
			ZoomLevel:          firefly.Int(6),
			XTraceID:           firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
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

func TestAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.List(context.TODO(), firefly.AccountListParams{
		Date:     firefly.Time(time.Now()),
		End:      firefly.Time(time.Now()),
		Limit:    firefly.Int(10),
		Page:     firefly.Int(1),
		Start:    firefly.Time(time.Now()),
		Type:     firefly.AccountTypeFilterAll,
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

func TestAccountDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Accounts.Delete(
		context.TODO(),
		"123",
		firefly.AccountDeleteParams{
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

func TestAccountListAttachmentsWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.ListAttachments(
		context.TODO(),
		"123",
		firefly.AccountListAttachmentsParams{
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

func TestAccountListPiggyBanksWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.ListPiggyBanks(
		context.TODO(),
		"123",
		firefly.AccountListPiggyBanksParams{
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

func TestAccountListTransactionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.ListTransactions(
		context.TODO(),
		"123",
		firefly.AccountListTransactionsParams{
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
