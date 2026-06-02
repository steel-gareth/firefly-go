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

func TestTransactionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.New(context.TODO(), firefly.TransactionNewParams{
		Transactions: []firefly.TransactionNewParamsTransaction{{
			Amount:              "123.45",
			Date:                time.Now(),
			Description:         "Vegetables",
			Type:                firefly.TransactionTypePropertyWithdrawal,
			BillID:              firefly.String("112"),
			BillName:            firefly.String("Monthly rent"),
			BookDate:            firefly.Time(time.Now()),
			BudgetID:            firefly.String("4"),
			BudgetName:          firefly.String("Groceries"),
			CategoryID:          firefly.String("43"),
			CategoryName:        firefly.String("Groceries"),
			CurrencyCode:        firefly.String("EUR"),
			CurrencyID:          firefly.String("12"),
			DestinationID:       firefly.String("2"),
			DestinationName:     firefly.String("Buy and Large"),
			DueDate:             firefly.Time(time.Now()),
			ExternalID:          firefly.String("external_id"),
			ExternalURL:         firefly.String("external_url"),
			ForeignAmount:       firefly.String("123.45"),
			ForeignCurrencyCode: firefly.String("USD"),
			ForeignCurrencyID:   firefly.String("17"),
			InterestDate:        firefly.Time(time.Now()),
			InternalReference:   firefly.String("internal_reference"),
			InvoiceDate:         firefly.Time(time.Now()),
			Notes:               firefly.String("Some example notes"),
			Order:               firefly.Int(0),
			PaymentDate:         firefly.Time(time.Now()),
			PiggyBankID:         firefly.Int(0),
			PiggyBankName:       firefly.String("piggy_bank_name"),
			ProcessDate:         firefly.Time(time.Now()),
			Reconciled:          firefly.Bool(false),
			SepaBatchID:         firefly.String("sepa_batch_id"),
			SepaCc:              firefly.String("sepa_cc"),
			SepaCi:              firefly.String("sepa_ci"),
			SepaCountry:         firefly.String("sepa_country"),
			SepaCtID:            firefly.String("sepa_ct_id"),
			SepaCtOp:            firefly.String("sepa_ct_op"),
			SepaDB:              firefly.String("sepa_db"),
			SepaEp:              firefly.String("sepa_ep"),
			SourceID:            firefly.String("2"),
			SourceName:          firefly.String("Checking account"),
			Tags:                []string{"Barbecue preparation"},
		}},
		ApplyRules:           firefly.Bool(false),
		ErrorIfDuplicateHash: firefly.Bool(false),
		FireWebhooks:         firefly.Bool(true),
		GroupTitle:           firefly.String("Split transaction title."),
		XTraceID:             firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
	})
	if err != nil {
		var apierr *firefly.Error
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
	client := firefly.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Transactions.Get(
		context.TODO(),
		"123",
		firefly.TransactionGetParams{
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

func TestTransactionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.Update(
		context.TODO(),
		"123",
		firefly.TransactionUpdateParams{
			ApplyRules:   firefly.Bool(false),
			FireWebhooks: firefly.Bool(true),
			GroupTitle:   firefly.String("Split transaction title."),
			Transactions: []firefly.TransactionUpdateParamsTransaction{{
				Amount:               firefly.String("123.45"),
				BillID:               firefly.String("111"),
				BillName:             firefly.String("Monthly rent"),
				BookDate:             firefly.Time(time.Now()),
				BudgetID:             firefly.String("4"),
				CategoryID:           firefly.String("43"),
				CategoryName:         firefly.String("Groceries"),
				CurrencyCode:         firefly.String("EUR"),
				CurrencyID:           firefly.String("12"),
				Date:                 firefly.Time(time.Now()),
				Description:          firefly.String("Vegetables"),
				DestinationIban:      firefly.String("NL02ABNA0123456789"),
				DestinationID:        firefly.String("2"),
				DestinationName:      firefly.String("Buy and Large"),
				DueDate:              firefly.Time(time.Now()),
				ExternalID:           firefly.String("external_id"),
				ExternalURL:          firefly.String("external_url"),
				ForeignAmount:        firefly.String("123.45"),
				ForeignCurrencyCode:  firefly.String("USD"),
				ForeignCurrencyID:    firefly.String("17"),
				InterestDate:         firefly.Time(time.Now()),
				InternalReference:    firefly.String("internal_reference"),
				InvoiceDate:          firefly.Time(time.Now()),
				Notes:                firefly.String("Some example notes"),
				Order:                firefly.Int(0),
				PaymentDate:          firefly.Time(time.Now()),
				ProcessDate:          firefly.Time(time.Now()),
				Reconciled:           firefly.Bool(false),
				SepaBatchID:          firefly.String("sepa_batch_id"),
				SepaCc:               firefly.String("sepa_cc"),
				SepaCi:               firefly.String("sepa_ci"),
				SepaCountry:          firefly.String("sepa_country"),
				SepaCtID:             firefly.String("sepa_ct_id"),
				SepaCtOp:             firefly.String("sepa_ct_op"),
				SepaDB:               firefly.String("sepa_db"),
				SepaEp:               firefly.String("sepa_ep"),
				SourceIban:           firefly.String("NL02ABNA0123456789"),
				SourceID:             firefly.String("2"),
				SourceName:           firefly.String("Checking account"),
				Tags:                 []string{"Barbecue preparation"},
				TransactionJournalID: firefly.String("123"),
				Type:                 firefly.TransactionTypePropertyWithdrawal,
			}},
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

func TestTransactionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.List(context.TODO(), firefly.TransactionListParams{
		End:      firefly.Time(time.Now()),
		Limit:    firefly.Int(10),
		Page:     firefly.Int(1),
		Start:    firefly.Time(time.Now()),
		Type:     firefly.TransactionTypeFilterAll,
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

func TestTransactionDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Transactions.Delete(
		context.TODO(),
		"123",
		firefly.TransactionDeleteParams{
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

func TestTransactionListAttachmentsWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.ListAttachments(
		context.TODO(),
		"123",
		firefly.TransactionListAttachmentsParams{
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

func TestTransactionListPiggyBankEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.ListPiggyBankEvents(
		context.TODO(),
		"123",
		firefly.TransactionListPiggyBankEventsParams{
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
