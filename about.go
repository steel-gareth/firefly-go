// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/steel-gareth/firefly-go/internal/apijson"
	"github.com/steel-gareth/firefly-go/internal/requestconfig"
	"github.com/steel-gareth/firefly-go/option"
	"github.com/steel-gareth/firefly-go/packages/param"
	"github.com/steel-gareth/firefly-go/packages/respjson"
)

// These endpoints deliver general system information, version- and meta
// information.
//
// AboutService contains methods and other services that help with interacting with
// the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAboutService] method instead.
type AboutService struct {
	options []option.RequestOption
}

// NewAboutService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAboutService(opts ...option.RequestOption) (r AboutService) {
	r = AboutService{}
	r.options = opts
	return
}

// Returns general system information and versions of the (supporting) software.
func (r *AboutService) GetInfo(ctx context.Context, query AboutGetInfoParams, opts ...option.RequestOption) (res *AboutGetInfoResponse, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/about"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the currently authenticated user.
func (r *AboutService) GetUser(ctx context.Context, query AboutGetUserParams, opts ...option.RequestOption) (res *UserSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/about/user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UserRead struct {
	ID         string     `json:"id" api:"required"`
	Attributes User       `json:"attributes" api:"required"`
	Links      ObjectLink `json:"links" api:"required"`
	// Immutable value
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attributes  respjson.Field
		Links       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserRead) RawJSON() string { return r.JSON.raw }
func (r *UserRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserSingle struct {
	Data UserRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserSingle) RawJSON() string { return r.JSON.raw }
func (r *UserSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AboutGetInfoResponse struct {
	Data AboutGetInfoResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AboutGetInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *AboutGetInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AboutGetInfoResponseData struct {
	// Same value as the version field.
	APIVersion string `json:"api_version"`
	Driver     string `json:"driver"`
	Os         string `json:"os"`
	PhpVersion string `json:"php_version"`
	Version    string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion  respjson.Field
		Driver      respjson.Field
		Os          respjson.Field
		PhpVersion  respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AboutGetInfoResponseData) RawJSON() string { return r.JSON.raw }
func (r *AboutGetInfoResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AboutGetInfoParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type AboutGetUserParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
