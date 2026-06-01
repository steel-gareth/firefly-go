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

func TestBillNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Bills.New(context.TODO(), emceesprodtesting5.BillNewParams{
		AmountMax:        "123.45",
		AmountMin:        "123.45",
		Date:             time.Now(),
		Name:             "Rent",
		RepeatFreq:       emceesprodtesting5.BillRepeatFrequencyMonthly,
		Active:           emceesprodtesting5.Bool(true),
		CurrencyCode:     emceesprodtesting5.String("EUR"),
		CurrencyID:       emceesprodtesting5.String("5"),
		EndDate:          emceesprodtesting5.Time(time.Now()),
		ExtensionDate:    emceesprodtesting5.Time(time.Now()),
		Notes:            emceesprodtesting5.String("Some example notes"),
		ObjectGroupID:    emceesprodtesting5.String("5"),
		ObjectGroupTitle: emceesprodtesting5.String("Example Group"),
		Skip:             emceesprodtesting5.Int(0),
		XTraceID:         emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
	})
	if err != nil {
		var apierr *emceesprodtesting5.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBillGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Bills.Get(
		context.TODO(),
		"123",
		emceesprodtesting5.BillGetParams{
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

func TestBillUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Bills.Update(
		context.TODO(),
		"123",
		emceesprodtesting5.BillUpdateParams{
			Name:             "Rent",
			Active:           emceesprodtesting5.Bool(true),
			AmountMax:        emceesprodtesting5.String("123.45"),
			AmountMin:        emceesprodtesting5.String("123.45"),
			CurrencyCode:     emceesprodtesting5.String("EUR"),
			CurrencyID:       emceesprodtesting5.String("5"),
			Date:             emceesprodtesting5.Time(time.Now()),
			EndDate:          emceesprodtesting5.Time(time.Now()),
			ExtensionDate:    emceesprodtesting5.Time(time.Now()),
			Notes:            emceesprodtesting5.String("Some example notes"),
			ObjectGroupID:    emceesprodtesting5.String("5"),
			ObjectGroupTitle: emceesprodtesting5.String("Example Group"),
			RepeatFreq:       emceesprodtesting5.BillRepeatFrequencyMonthly,
			Skip:             emceesprodtesting5.Int(0),
			XTraceID:         emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
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

func TestBillListWithOptionalParams(t *testing.T) {
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
	_, err := client.Bills.List(context.TODO(), emceesprodtesting5.BillListParams{
		End:      emceesprodtesting5.Time(time.Now()),
		Limit:    emceesprodtesting5.Int(10),
		Page:     emceesprodtesting5.Int(1),
		Start:    emceesprodtesting5.Time(time.Now()),
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

func TestBillDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Bills.Delete(
		context.TODO(),
		"123",
		emceesprodtesting5.BillDeleteParams{
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

func TestBillListAttachmentsWithOptionalParams(t *testing.T) {
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
	_, err := client.Bills.ListAttachments(
		context.TODO(),
		"123",
		emceesprodtesting5.BillListAttachmentsParams{
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

func TestBillListRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Bills.ListRules(
		context.TODO(),
		"123",
		emceesprodtesting5.BillListRulesParams{
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

func TestBillListTransactionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Bills.ListTransactions(
		context.TODO(),
		"123",
		emceesprodtesting5.BillListTransactionsParams{
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
