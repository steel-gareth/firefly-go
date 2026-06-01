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

func TestRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Rules.New(context.TODO(), firefly.RuleNewParams{
		Actions: []firefly.RuleNewParamsAction{{
			Type:           firefly.RuleActionKeywordSetCategory,
			Value:          firefly.String("Daily groceries"),
			Active:         firefly.Bool(true),
			Order:          firefly.Int(5),
			StopProcessing: firefly.Bool(false),
		}},
		RuleGroupID: "81",
		Title:       "First rule title.",
		Trigger:     firefly.RuleTriggerTypeStoreJournal,
		Triggers: []firefly.RuleNewParamsTrigger{{
			Type:           firefly.RuleTriggerKeywordFromAccountStarts,
			Value:          "tag1",
			Active:         firefly.Bool(true),
			Order:          firefly.Int(5),
			Prohibited:     firefly.Bool(false),
			StopProcessing: firefly.Bool(false),
		}},
		Active:         firefly.Bool(true),
		Description:    firefly.String("First rule description"),
		Order:          firefly.Int(5),
		RuleGroupTitle: firefly.String("New rule group"),
		StopProcessing: firefly.Bool(false),
		Strict:         firefly.Bool(true),
		XTraceID:       firefly.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
	})
	if err != nil {
		var apierr *firefly.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Rules.Get(
		context.TODO(),
		"123",
		firefly.RuleGetParams{
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

func TestRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Rules.Update(
		context.TODO(),
		"123",
		firefly.RuleUpdateParams{
			Actions: []firefly.RuleUpdateParamsAction{{
				Active:         firefly.Bool(true),
				Order:          firefly.Int(5),
				StopProcessing: firefly.Bool(false),
				Type:           firefly.RuleActionKeywordSetCategory,
				Value:          firefly.String("Daily groceries"),
			}},
			Active:         firefly.Bool(true),
			Description:    firefly.String("First rule description"),
			Order:          firefly.Int(5),
			RuleGroupID:    firefly.String("81"),
			StopProcessing: firefly.Bool(false),
			Strict:         firefly.Bool(true),
			Title:          firefly.String("First rule title."),
			Trigger:        firefly.RuleTriggerTypeStoreJournal,
			Triggers: []firefly.RuleUpdateParamsTrigger{{
				Active:         firefly.Bool(true),
				Order:          firefly.Int(5),
				StopProcessing: firefly.Bool(false),
				Type:           firefly.RuleTriggerKeywordFromAccountStarts,
				Value:          firefly.String("tag1"),
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

func TestRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Rules.List(context.TODO(), firefly.RuleListParams{
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

func TestRuleDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Rules.Delete(
		context.TODO(),
		"123",
		firefly.RuleDeleteParams{
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

func TestRuleTestWithOptionalParams(t *testing.T) {
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
	_, err := client.Rules.Test(
		context.TODO(),
		"123",
		firefly.RuleTestParams{
			Accounts: []int64{1, 2, 3},
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

func TestRuleTriggerWithOptionalParams(t *testing.T) {
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
	err := client.Rules.Trigger(
		context.TODO(),
		"123",
		firefly.RuleTriggerParams{
			Accounts: []int64{1, 2, 3},
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
