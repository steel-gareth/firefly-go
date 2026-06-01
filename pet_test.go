// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package emceesprodtesting5_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/emcees-prod-testing-5-go"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/testutil"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
)

func TestPetNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pets.New(context.TODO(), emceesprodtesting5.PetNewParams{
		Pet: emceesprodtesting5.PetParam{
			Name:      "doggie",
			PhotoURLs: []string{"string"},
			ID:        emceesprodtesting5.Int(10),
			Category: emceesprodtesting5.CategoryParam{
				ID:   emceesprodtesting5.Int(1),
				Name: emceesprodtesting5.String("Dogs"),
			},
			Status: emceesprodtesting5.PetStatusAvailable,
			Tags: []emceesprodtesting5.TagParam{{
				ID:   emceesprodtesting5.Int(0),
				Name: emceesprodtesting5.String("name"),
			}},
		},
	})
	if err != nil {
		var apierr *emceesprodtesting5.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetGet(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pets.Get(context.TODO(), 0)
	if err != nil {
		var apierr *emceesprodtesting5.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pets.Update(context.TODO(), emceesprodtesting5.PetUpdateParams{
		Pet: emceesprodtesting5.PetParam{
			Name:      "doggie",
			PhotoURLs: []string{"string"},
			ID:        emceesprodtesting5.Int(10),
			Category: emceesprodtesting5.CategoryParam{
				ID:   emceesprodtesting5.Int(1),
				Name: emceesprodtesting5.String("Dogs"),
			},
			Status: emceesprodtesting5.PetStatusAvailable,
			Tags: []emceesprodtesting5.TagParam{{
				ID:   emceesprodtesting5.Int(0),
				Name: emceesprodtesting5.String("name"),
			}},
		},
	})
	if err != nil {
		var apierr *emceesprodtesting5.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetDelete(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	err := client.Pets.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *emceesprodtesting5.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByStatusWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pets.FindByStatus(context.TODO(), emceesprodtesting5.PetFindByStatusParams{
		Status: emceesprodtesting5.PetFindByStatusParamsStatusAvailable,
	})
	if err != nil {
		var apierr *emceesprodtesting5.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByTagsWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pets.FindByTags(context.TODO(), emceesprodtesting5.PetFindByTagsParams{
		Tags: []string{"string"},
	})
	if err != nil {
		var apierr *emceesprodtesting5.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateByIDWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	err := client.Pets.UpdateByID(
		context.TODO(),
		0,
		emceesprodtesting5.PetUpdateByIDParams{
			Name:   emceesprodtesting5.String("name"),
			Status: emceesprodtesting5.String("status"),
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

func TestPetUploadImageWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pets.UploadImage(
		context.TODO(),
		0,
		io.Reader(bytes.NewBuffer([]byte("Example data"))),
		emceesprodtesting5.PetUploadImageParams{
			AdditionalMetadata: emceesprodtesting5.String("additionalMetadata"),
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
