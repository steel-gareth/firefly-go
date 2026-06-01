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

func TestTransactionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.New(context.TODO(), emceesprodtesting5.TransactionNewParams{
		Transactions: []emceesprodtesting5.TransactionNewParamsTransaction{{
			Amount:              "123.45",
			Date:                time.Now(),
			Description:         "Vegetables",
			Type:                emceesprodtesting5.TransactionTypePropertyWithdrawal,
			BillID:              emceesprodtesting5.String("112"),
			BillName:            emceesprodtesting5.String("Monthly rent"),
			BookDate:            emceesprodtesting5.Time(time.Now()),
			BudgetID:            emceesprodtesting5.String("4"),
			BudgetName:          emceesprodtesting5.String("Groceries"),
			CategoryID:          emceesprodtesting5.String("43"),
			CategoryName:        emceesprodtesting5.String("Groceries"),
			CurrencyCode:        emceesprodtesting5.String("EUR"),
			CurrencyID:          emceesprodtesting5.String("12"),
			DestinationID:       emceesprodtesting5.String("2"),
			DestinationName:     emceesprodtesting5.String("Buy and Large"),
			DueDate:             emceesprodtesting5.Time(time.Now()),
			ExternalID:          emceesprodtesting5.String("external_id"),
			ExternalURL:         emceesprodtesting5.String("external_url"),
			ForeignAmount:       emceesprodtesting5.String("123.45"),
			ForeignCurrencyCode: emceesprodtesting5.String("USD"),
			ForeignCurrencyID:   emceesprodtesting5.String("17"),
			InterestDate:        emceesprodtesting5.Time(time.Now()),
			InternalReference:   emceesprodtesting5.String("internal_reference"),
			InvoiceDate:         emceesprodtesting5.Time(time.Now()),
			Notes:               emceesprodtesting5.String("Some example notes"),
			Order:               emceesprodtesting5.Int(0),
			PaymentDate:         emceesprodtesting5.Time(time.Now()),
			PiggyBankID:         emceesprodtesting5.Int(0),
			PiggyBankName:       emceesprodtesting5.String("piggy_bank_name"),
			ProcessDate:         emceesprodtesting5.Time(time.Now()),
			Reconciled:          emceesprodtesting5.Bool(false),
			SepaBatchID:         emceesprodtesting5.String("sepa_batch_id"),
			SepaCc:              emceesprodtesting5.String("sepa_cc"),
			SepaCi:              emceesprodtesting5.String("sepa_ci"),
			SepaCountry:         emceesprodtesting5.String("sepa_country"),
			SepaCtID:            emceesprodtesting5.String("sepa_ct_id"),
			SepaCtOp:            emceesprodtesting5.String("sepa_ct_op"),
			SepaDB:              emceesprodtesting5.String("sepa_db"),
			SepaEp:              emceesprodtesting5.String("sepa_ep"),
			SourceID:            emceesprodtesting5.String("2"),
			SourceName:          emceesprodtesting5.String("Checking account"),
			Tags:                []string{"Barbecue preparation"},
		}},
		ApplyRules:           emceesprodtesting5.Bool(false),
		ErrorIfDuplicateHash: emceesprodtesting5.Bool(false),
		FireWebhooks:         emceesprodtesting5.Bool(true),
		GroupTitle:           emceesprodtesting5.String("Split transaction title."),
		XTraceID:             emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
	})
	if err != nil {
		var apierr *emceesprodtesting5.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.Get(
		context.TODO(),
		"123",
		emceesprodtesting5.TransactionGetParams{
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

func TestTransactionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.Update(
		context.TODO(),
		"123",
		emceesprodtesting5.TransactionUpdateParams{
			ApplyRules:   emceesprodtesting5.Bool(false),
			FireWebhooks: emceesprodtesting5.Bool(true),
			GroupTitle:   emceesprodtesting5.String("Split transaction title."),
			Transactions: []emceesprodtesting5.TransactionUpdateParamsTransaction{{
				Amount:               emceesprodtesting5.String("123.45"),
				BillID:               emceesprodtesting5.String("111"),
				BillName:             emceesprodtesting5.String("Monthly rent"),
				BookDate:             emceesprodtesting5.Time(time.Now()),
				BudgetID:             emceesprodtesting5.String("4"),
				CategoryID:           emceesprodtesting5.String("43"),
				CategoryName:         emceesprodtesting5.String("Groceries"),
				CurrencyCode:         emceesprodtesting5.String("EUR"),
				CurrencyID:           emceesprodtesting5.String("12"),
				Date:                 emceesprodtesting5.Time(time.Now()),
				Description:          emceesprodtesting5.String("Vegetables"),
				DestinationIban:      emceesprodtesting5.String("NL02ABNA0123456789"),
				DestinationID:        emceesprodtesting5.String("2"),
				DestinationName:      emceesprodtesting5.String("Buy and Large"),
				DueDate:              emceesprodtesting5.Time(time.Now()),
				ExternalID:           emceesprodtesting5.String("external_id"),
				ExternalURL:          emceesprodtesting5.String("external_url"),
				ForeignAmount:        emceesprodtesting5.String("123.45"),
				ForeignCurrencyCode:  emceesprodtesting5.String("USD"),
				ForeignCurrencyID:    emceesprodtesting5.String("17"),
				InterestDate:         emceesprodtesting5.Time(time.Now()),
				InternalReference:    emceesprodtesting5.String("internal_reference"),
				InvoiceDate:          emceesprodtesting5.Time(time.Now()),
				Notes:                emceesprodtesting5.String("Some example notes"),
				Order:                emceesprodtesting5.Int(0),
				PaymentDate:          emceesprodtesting5.Time(time.Now()),
				ProcessDate:          emceesprodtesting5.Time(time.Now()),
				Reconciled:           emceesprodtesting5.Bool(false),
				SepaBatchID:          emceesprodtesting5.String("sepa_batch_id"),
				SepaCc:               emceesprodtesting5.String("sepa_cc"),
				SepaCi:               emceesprodtesting5.String("sepa_ci"),
				SepaCountry:          emceesprodtesting5.String("sepa_country"),
				SepaCtID:             emceesprodtesting5.String("sepa_ct_id"),
				SepaCtOp:             emceesprodtesting5.String("sepa_ct_op"),
				SepaDB:               emceesprodtesting5.String("sepa_db"),
				SepaEp:               emceesprodtesting5.String("sepa_ep"),
				SourceIban:           emceesprodtesting5.String("NL02ABNA0123456789"),
				SourceID:             emceesprodtesting5.String("2"),
				SourceName:           emceesprodtesting5.String("Checking account"),
				Tags:                 []string{"Barbecue preparation"},
				TransactionJournalID: emceesprodtesting5.String("123"),
				Type:                 emceesprodtesting5.TransactionTypePropertyWithdrawal,
			}},
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

func TestTransactionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.List(context.TODO(), emceesprodtesting5.TransactionListParams{
		End:      emceesprodtesting5.Time(time.Now()),
		Limit:    emceesprodtesting5.Int(10),
		Page:     emceesprodtesting5.Int(1),
		Start:    emceesprodtesting5.Time(time.Now()),
		Type:     emceesprodtesting5.TransactionTypeFilterAll,
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

func TestTransactionDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Transactions.Delete(
		context.TODO(),
		"123",
		emceesprodtesting5.TransactionDeleteParams{
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

func TestTransactionListAttachmentsWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.ListAttachments(
		context.TODO(),
		"123",
		emceesprodtesting5.TransactionListAttachmentsParams{
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

func TestTransactionListPiggyBankEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.ListPiggyBankEvents(
		context.TODO(),
		"123",
		emceesprodtesting5.TransactionListPiggyBankEventsParams{
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
