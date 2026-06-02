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

func TestRecurrenceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Recurrences.New(context.TODO(), firefly.RecurrenceNewParams{
		FirstDate:   time.Now(),
		RepeatUntil: firefly.Time(time.Now()),
		Repetitions: []firefly.RecurrenceNewParamsRepetition{{
			Moment:  "3",
			Type:    firefly.RecurrenceRepetitionTypeWeekly,
			Skip:    firefly.Int(0),
			Weekend: firefly.Int(1),
		}},
		Title: "Rent",
		Transactions: []firefly.RecurrenceNewParamsTransaction{{
			Amount:              "123.45",
			Description:         "Rent for the current month",
			DestinationID:       "258",
			SourceID:            "913",
			BillID:              firefly.String("123"),
			BudgetID:            firefly.String("4"),
			CategoryID:          firefly.String("211"),
			CurrencyCode:        firefly.String("EUR"),
			CurrencyID:          firefly.String("3"),
			ForeignAmount:       firefly.String("123.45"),
			ForeignCurrencyCode: firefly.String("GBP"),
			ForeignCurrencyID:   firefly.String("17"),
			PiggyBankID:         firefly.String("123"),
			Tags:                []string{"Barbecue preparation"},
		}},
		Type:            firefly.RecurrenceTransactionTypeWithdrawal,
		Active:          firefly.Bool(true),
		ApplyRules:      firefly.Bool(true),
		Description:     firefly.String("Recurring transaction for the monthly rent"),
		Notes:           firefly.String("Some notes"),
		NrOfRepetitions: firefly.Int(5),
		XTraceID:        firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
	})
	if err != nil {
		var apierr *firefly.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecurrenceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Recurrences.Get(
		context.TODO(),
		"123",
		firefly.RecurrenceGetParams{
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

func TestRecurrenceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Recurrences.Update(
		context.TODO(),
		"123",
		firefly.RecurrenceUpdateParams{
			Active:          firefly.Bool(true),
			ApplyRules:      firefly.Bool(true),
			Description:     firefly.String("Recurring transaction for the monthly rent"),
			FirstDate:       firefly.Time(time.Now()),
			Notes:           firefly.String("Some notes"),
			NrOfRepetitions: firefly.Int(5),
			RepeatUntil:     firefly.Time(time.Now()),
			Repetitions: []firefly.RecurrenceUpdateParamsRepetition{{
				Moment:  firefly.String("3"),
				Skip:    firefly.Int(0),
				Type:    firefly.RecurrenceRepetitionTypeWeekly,
				Weekend: firefly.Int(1),
			}},
			Title: firefly.String("Rent"),
			Transactions: []firefly.RecurrenceUpdateParamsTransaction{{
				ID:                "ID of the recurring transaction. Not to be confused with the ID of the recurrence itself. Is marked as REQUIRED but can be skipped when there is only ONE transaction.",
				Amount:            firefly.String("123.45"),
				BillID:            firefly.String("123"),
				BudgetID:          firefly.String("4"),
				CategoryID:        firefly.String("211"),
				CurrencyCode:      firefly.String("EUR"),
				CurrencyID:        firefly.String("3"),
				Description:       firefly.String("Rent for the current month"),
				DestinationID:     firefly.String("258"),
				ForeignAmount:     firefly.String("123.45"),
				ForeignCurrencyID: firefly.String("17"),
				PiggyBankID:       firefly.String("123"),
				SourceID:          firefly.String("913"),
				Tags:              []string{"Barbecue preparation"},
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

func TestRecurrenceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Recurrences.List(context.TODO(), firefly.RecurrenceListParams{
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

func TestRecurrenceDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Recurrences.Delete(
		context.TODO(),
		"123",
		firefly.RecurrenceDeleteParams{
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

func TestRecurrenceListTransactionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Recurrences.ListTransactions(
		context.TODO(),
		"123",
		firefly.RecurrenceListTransactionsParams{
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

func TestRecurrenceTriggerTransactionWithOptionalParams(t *testing.T) {
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
	_, err := client.Recurrences.TriggerTransaction(
		context.TODO(),
		"123",
		firefly.RecurrenceTriggerTransactionParams{
			Date:     time.Now(),
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
