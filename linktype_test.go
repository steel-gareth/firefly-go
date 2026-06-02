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

func TestLinkTypeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LinkTypes.New(context.TODO(), firefly.LinkTypeNewParams{
		LinkType: firefly.LinkTypeParam{
			Inward:  "is (partially) paid for by",
			Name:    "Paid",
			Outward: "(partially) pays for",
		},
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

func TestLinkTypeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.LinkTypes.Get(
		context.TODO(),
		"123",
		firefly.LinkTypeGetParams{
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

func TestLinkTypeUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.LinkTypes.Update(
		context.TODO(),
		"123",
		firefly.LinkTypeUpdateParams{
			Inward:   firefly.String("is (partially) paid for by"),
			Name:     firefly.String("Paid"),
			Outward:  firefly.String("(partially) pays for"),
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

func TestLinkTypeListWithOptionalParams(t *testing.T) {
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
	_, err := client.LinkTypes.List(context.TODO(), firefly.LinkTypeListParams{
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

func TestLinkTypeDeleteWithOptionalParams(t *testing.T) {
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
	err := client.LinkTypes.Delete(
		context.TODO(),
		"123",
		firefly.LinkTypeDeleteParams{
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

func TestLinkTypeListTransactionsWithOptionalParams(t *testing.T) {
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
	_, err := client.LinkTypes.ListTransactions(
		context.TODO(),
		"123",
		firefly.LinkTypeListTransactionsParams{
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
