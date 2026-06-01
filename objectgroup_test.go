// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package emceesprodtesting5_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/emcees-prod-testing-5-go"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/testutil"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
)

func TestObjectGroupGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ObjectGroups.Get(
		context.TODO(),
		"123",
		emceesprodtesting5.ObjectGroupGetParams{
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

func TestObjectGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ObjectGroups.Update(
		context.TODO(),
		"123",
		emceesprodtesting5.ObjectGroupUpdateParams{
			Title:    "My object group",
			Order:    emceesprodtesting5.Int(1),
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

func TestObjectGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.ObjectGroups.List(context.TODO(), emceesprodtesting5.ObjectGroupListParams{
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

func TestObjectGroupDeleteWithOptionalParams(t *testing.T) {
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
	err := client.ObjectGroups.Delete(
		context.TODO(),
		"123",
		emceesprodtesting5.ObjectGroupDeleteParams{
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

func TestObjectGroupListBillsWithOptionalParams(t *testing.T) {
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
	_, err := client.ObjectGroups.ListBills(
		context.TODO(),
		"123",
		emceesprodtesting5.ObjectGroupListBillsParams{
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

func TestObjectGroupListPiggyBanksWithOptionalParams(t *testing.T) {
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
	_, err := client.ObjectGroups.ListPiggyBanks(
		context.TODO(),
		"123",
		emceesprodtesting5.ObjectGroupListPiggyBanksParams{
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
