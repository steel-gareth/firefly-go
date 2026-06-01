// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package emceesprodtesting5_test

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

func TestAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.New(context.TODO(), emceesprodtesting5.AccountNewParams{
		Name:               "My checking account",
		Type:               emceesprodtesting5.ShortAccountTypePropertyAsset,
		AccountNumber:      emceesprodtesting5.String("7009312345678"),
		AccountRole:        emceesprodtesting5.AccountRolePropertyDefaultAsset,
		Active:             emceesprodtesting5.Bool(false),
		Bic:                emceesprodtesting5.String("BOFAUS3N"),
		CreditCardType:     emceesprodtesting5.CreditCardTypePropertyMonthlyFull,
		CurrencyCode:       emceesprodtesting5.String("EUR"),
		CurrencyID:         emceesprodtesting5.String("12"),
		Iban:               emceesprodtesting5.String("GB98MIDL07009312345678"),
		IncludeNetWorth:    emceesprodtesting5.Bool(true),
		Interest:           emceesprodtesting5.String("5.3"),
		InterestPeriod:     emceesprodtesting5.InterestPeriodPropertyMonthly,
		Latitude:           emceesprodtesting5.Float(51.983333),
		LiabilityDirection: emceesprodtesting5.LiabilityDirectionPropertyCredit,
		LiabilityType:      emceesprodtesting5.LiabilityTypePropertyLoan,
		Longitude:          emceesprodtesting5.Float(5.916667),
		MonthlyPaymentDate: emceesprodtesting5.Time(time.Now()),
		Notes:              emceesprodtesting5.String("Some example notes"),
		OpeningBalance:     emceesprodtesting5.String("-1012.12"),
		OpeningBalanceDate: emceesprodtesting5.Time(time.Now()),
		Order:              emceesprodtesting5.Int(1),
		VirtualBalance:     emceesprodtesting5.String("123.45"),
		ZoomLevel:          emceesprodtesting5.Int(6),
		XTraceID:           emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
	})
	if err != nil {
		var apierr *emceesprodtesting5.Error
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
	client := emceesprodtesting5.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Accounts.Get(
		context.TODO(),
		"123",
		emceesprodtesting5.AccountGetParams{
			Date:     emceesprodtesting5.Time(time.Now()),
			End:      emceesprodtesting5.Time(time.Now()),
			Start:    emceesprodtesting5.Time(time.Now()),
			XTraceID: emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
		},
	)
	if err != nil {
		var apierr *emceesprodtesting5.Error
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
	client := emceesprodtesting5.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Accounts.Update(
		context.TODO(),
		"123",
		emceesprodtesting5.AccountUpdateParams{
			Name:               "My checking account",
			Type:               map[string]any{},
			AccountNumber:      emceesprodtesting5.String("7009312345678"),
			AccountRole:        emceesprodtesting5.AccountRolePropertyDefaultAsset,
			Active:             emceesprodtesting5.Bool(false),
			Bic:                emceesprodtesting5.String("BOFAUS3N"),
			CreditCardType:     emceesprodtesting5.CreditCardTypePropertyMonthlyFull,
			CurrencyCode:       emceesprodtesting5.String("EUR"),
			CurrencyID:         emceesprodtesting5.String("12"),
			Iban:               emceesprodtesting5.String("GB98MIDL07009312345678"),
			IncludeNetWorth:    emceesprodtesting5.Bool(true),
			Interest:           emceesprodtesting5.String("5.3"),
			InterestPeriod:     emceesprodtesting5.InterestPeriodPropertyMonthly,
			Latitude:           emceesprodtesting5.Float(51.983333),
			LiabilityType:      emceesprodtesting5.LiabilityTypePropertyLoan,
			Longitude:          emceesprodtesting5.Float(5.916667),
			MonthlyPaymentDate: emceesprodtesting5.Time(time.Now()),
			Notes:              emceesprodtesting5.String("Some example notes"),
			OpeningBalance:     emceesprodtesting5.String("-1012.12"),
			OpeningBalanceDate: emceesprodtesting5.Time(time.Now()),
			Order:              emceesprodtesting5.Int(1),
			VirtualBalance:     emceesprodtesting5.String("123.45"),
			ZoomLevel:          emceesprodtesting5.Int(6),
			XTraceID:           emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
		},
	)
	if err != nil {
		var apierr *emceesprodtesting5.Error
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
	client := emceesprodtesting5.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Accounts.List(context.TODO(), emceesprodtesting5.AccountListParams{
		Date:     emceesprodtesting5.Time(time.Now()),
		End:      emceesprodtesting5.Time(time.Now()),
		Limit:    emceesprodtesting5.Int(10),
		Page:     emceesprodtesting5.Int(1),
		Start:    emceesprodtesting5.Time(time.Now()),
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

func TestAccountDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Accounts.Delete(
		context.TODO(),
		"123",
		emceesprodtesting5.AccountDeleteParams{
			XTraceID: emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
		},
	)
	if err != nil {
		var apierr *emceesprodtesting5.Error
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
	client := emceesprodtesting5.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Accounts.ListAttachments(
		context.TODO(),
		"123",
		emceesprodtesting5.AccountListAttachmentsParams{
			Limit:    emceesprodtesting5.Int(10),
			Page:     emceesprodtesting5.Int(1),
			XTraceID: emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
		},
	)
	if err != nil {
		var apierr *emceesprodtesting5.Error
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
	client := emceesprodtesting5.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Accounts.ListPiggyBanks(
		context.TODO(),
		"123",
		emceesprodtesting5.AccountListPiggyBanksParams{
			Limit:    emceesprodtesting5.Int(10),
			Page:     emceesprodtesting5.Int(1),
			XTraceID: emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
		},
	)
	if err != nil {
		var apierr *emceesprodtesting5.Error
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
	client := emceesprodtesting5.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Accounts.ListTransactions(
		context.TODO(),
		"123",
		emceesprodtesting5.AccountListTransactionsParams{
			End:      emceesprodtesting5.Time(time.Now()),
			Limit:    emceesprodtesting5.Int(10),
			Page:     emceesprodtesting5.Int(1),
			Start:    emceesprodtesting5.Time(time.Now()),
			Type:     emceesprodtesting5.TransactionTypeFilterAll,
			XTraceID: emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
		},
	)
	if err != nil {
		var apierr *emceesprodtesting5.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
