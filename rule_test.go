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

func TestRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Rules.New(context.TODO(), emceesprodtesting5.RuleNewParams{
		Actions: []emceesprodtesting5.RuleNewParamsAction{{
			Type:           emceesprodtesting5.RuleActionKeywordSetCategory,
			Value:          emceesprodtesting5.String("Daily groceries"),
			Active:         emceesprodtesting5.Bool(true),
			Order:          emceesprodtesting5.Int(5),
			StopProcessing: emceesprodtesting5.Bool(false),
		}},
		RuleGroupID: "81",
		Title:       "First rule title.",
		Trigger:     emceesprodtesting5.RuleTriggerTypeStoreJournal,
		Triggers: []emceesprodtesting5.RuleNewParamsTrigger{{
			Type:           emceesprodtesting5.RuleTriggerKeywordFromAccountStarts,
			Value:          "tag1",
			Active:         emceesprodtesting5.Bool(true),
			Order:          emceesprodtesting5.Int(5),
			Prohibited:     emceesprodtesting5.Bool(false),
			StopProcessing: emceesprodtesting5.Bool(false),
		}},
		Active:         emceesprodtesting5.Bool(true),
		Description:    emceesprodtesting5.String("First rule description"),
		Order:          emceesprodtesting5.Int(5),
		RuleGroupTitle: emceesprodtesting5.String("New rule group"),
		StopProcessing: emceesprodtesting5.Bool(false),
		Strict:         emceesprodtesting5.Bool(true),
		XTraceID:       emceesprodtesting5.String("40c71bbb-c676-4f24-83cf-cc725d7d7a00"),
	})
	if err != nil {
		var apierr *emceesprodtesting5.Error
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
	client := emceesprodtesting5.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Rules.Get(
		context.TODO(),
		"123",
		emceesprodtesting5.RuleGetParams{
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

func TestRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Rules.Update(
		context.TODO(),
		"123",
		emceesprodtesting5.RuleUpdateParams{
			Actions: []emceesprodtesting5.RuleUpdateParamsAction{{
				Active:         emceesprodtesting5.Bool(true),
				Order:          emceesprodtesting5.Int(5),
				StopProcessing: emceesprodtesting5.Bool(false),
				Type:           emceesprodtesting5.RuleActionKeywordSetCategory,
				Value:          emceesprodtesting5.String("Daily groceries"),
			}},
			Active:         emceesprodtesting5.Bool(true),
			Description:    emceesprodtesting5.String("First rule description"),
			Order:          emceesprodtesting5.Int(5),
			RuleGroupID:    emceesprodtesting5.String("81"),
			StopProcessing: emceesprodtesting5.Bool(false),
			Strict:         emceesprodtesting5.Bool(true),
			Title:          emceesprodtesting5.String("First rule title."),
			Trigger:        emceesprodtesting5.RuleTriggerTypeStoreJournal,
			Triggers: []emceesprodtesting5.RuleUpdateParamsTrigger{{
				Active:         emceesprodtesting5.Bool(true),
				Order:          emceesprodtesting5.Int(5),
				StopProcessing: emceesprodtesting5.Bool(false),
				Type:           emceesprodtesting5.RuleTriggerKeywordFromAccountStarts,
				Value:          emceesprodtesting5.String("tag1"),
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

func TestRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Rules.List(context.TODO(), emceesprodtesting5.RuleListParams{
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

func TestRuleDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Rules.Delete(
		context.TODO(),
		"123",
		emceesprodtesting5.RuleDeleteParams{
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

func TestRuleTestWithOptionalParams(t *testing.T) {
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
	_, err := client.Rules.Test(
		context.TODO(),
		"123",
		emceesprodtesting5.RuleTestParams{
			Accounts: []int64{1, 2, 3},
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

func TestRuleTriggerWithOptionalParams(t *testing.T) {
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
	err := client.Rules.Trigger(
		context.TODO(),
		"123",
		emceesprodtesting5.RuleTriggerParams{
			Accounts: []int64{1, 2, 3},
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
