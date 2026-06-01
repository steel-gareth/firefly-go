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

func TestRecurrenceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Recurrences.New(context.TODO(), emceesprodtesting5.RecurrenceNewParams{
		FirstDate:   time.Now(),
		RepeatUntil: emceesprodtesting5.Time(time.Now()),
		Repetitions: []emceesprodtesting5.RecurrenceNewParamsRepetition{{
			Moment:  "3",
			Type:    emceesprodtesting5.RecurrenceRepetitionTypeWeekly,
			Skip:    emceesprodtesting5.Int(0),
			Weekend: emceesprodtesting5.Int(1),
		}},
		Title: "Rent",
		Transactions: []emceesprodtesting5.RecurrenceNewParamsTransaction{{
			Amount:              "123.45",
			Description:         "Rent for the current month",
			DestinationID:       "258",
			SourceID:            "913",
			BillID:              emceesprodtesting5.String("123"),
			BudgetID:            emceesprodtesting5.String("4"),
			CategoryID:          emceesprodtesting5.String("211"),
			CurrencyCode:        emceesprodtesting5.String("EUR"),
			CurrencyID:          emceesprodtesting5.String("3"),
			ForeignAmount:       emceesprodtesting5.String("123.45"),
			ForeignCurrencyCode: emceesprodtesting5.String("GBP"),
			ForeignCurrencyID:   emceesprodtesting5.String("17"),
			PiggyBankID:         emceesprodtesting5.String("123"),
			Tags:                []string{"Barbecue preparation"},
		}},
		Type:            emceesprodtesting5.RecurrenceTransactionTypeWithdrawal,
		Active:          emceesprodtesting5.Bool(true),
		ApplyRules:      emceesprodtesting5.Bool(true),
		Description:     emceesprodtesting5.String("Recurring transaction for the monthly rent"),
		Notes:           emceesprodtesting5.String("Some notes"),
		NrOfRepetitions: emceesprodtesting5.Int(5),
		XTraceID:        emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
	})
	if err != nil {
		var apierr *emceesprodtesting5.Error
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
	client := emceesprodtesting5.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Recurrences.Get(
		context.TODO(),
		"123",
		emceesprodtesting5.RecurrenceGetParams{
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

func TestRecurrenceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Recurrences.Update(
		context.TODO(),
		"123",
		emceesprodtesting5.RecurrenceUpdateParams{
			Active:          emceesprodtesting5.Bool(true),
			ApplyRules:      emceesprodtesting5.Bool(true),
			Description:     emceesprodtesting5.String("Recurring transaction for the monthly rent"),
			FirstDate:       emceesprodtesting5.Time(time.Now()),
			Notes:           emceesprodtesting5.String("Some notes"),
			NrOfRepetitions: emceesprodtesting5.Int(5),
			RepeatUntil:     emceesprodtesting5.Time(time.Now()),
			Repetitions: []emceesprodtesting5.RecurrenceUpdateParamsRepetition{{
				Moment:  emceesprodtesting5.String("3"),
				Skip:    emceesprodtesting5.Int(0),
				Type:    emceesprodtesting5.RecurrenceRepetitionTypeWeekly,
				Weekend: emceesprodtesting5.Int(1),
			}},
			Title: emceesprodtesting5.String("Rent"),
			Transactions: []emceesprodtesting5.RecurrenceUpdateParamsTransaction{{
				ID:                "ID of the recurring transaction. Not to be confused with the ID of the recurrence itself. Is marked as REQUIRED but can be skipped when there is only ONE transaction.",
				Amount:            emceesprodtesting5.String("123.45"),
				BillID:            emceesprodtesting5.String("123"),
				BudgetID:          emceesprodtesting5.String("4"),
				CategoryID:        emceesprodtesting5.String("211"),
				CurrencyCode:      emceesprodtesting5.String("EUR"),
				CurrencyID:        emceesprodtesting5.String("3"),
				Description:       emceesprodtesting5.String("Rent for the current month"),
				DestinationID:     emceesprodtesting5.String("258"),
				ForeignAmount:     emceesprodtesting5.String("123.45"),
				ForeignCurrencyID: emceesprodtesting5.String("17"),
				PiggyBankID:       emceesprodtesting5.String("123"),
				SourceID:          emceesprodtesting5.String("913"),
				Tags:              []string{"Barbecue preparation"},
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

func TestRecurrenceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Recurrences.List(context.TODO(), emceesprodtesting5.RecurrenceListParams{
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

func TestRecurrenceDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Recurrences.Delete(
		context.TODO(),
		"123",
		emceesprodtesting5.RecurrenceDeleteParams{
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

func TestRecurrenceListTransactionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Recurrences.ListTransactions(
		context.TODO(),
		"123",
		emceesprodtesting5.RecurrenceListTransactionsParams{
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

func TestRecurrenceTriggerTransactionWithOptionalParams(t *testing.T) {
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
	_, err := client.Recurrences.TriggerTransaction(
		context.TODO(),
		"123",
		emceesprodtesting5.RecurrenceTriggerTransactionParams{
			Date:     time.Now(),
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
