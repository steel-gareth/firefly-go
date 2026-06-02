// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/steel-gareth/firefly-go/internal/apijson"
	"github.com/steel-gareth/firefly-go/internal/apiquery"
	"github.com/steel-gareth/firefly-go/internal/requestconfig"
	"github.com/steel-gareth/firefly-go/option"
	"github.com/steel-gareth/firefly-go/packages/param"
	"github.com/steel-gareth/firefly-go/packages/respjson"
)

// These endpoints can be used to manage the user&#039;s webhooks and triggers them
// if necessary.
//
// WebhookMessageAttemptService contains methods and other services that help with
// interacting with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookMessageAttemptService] method instead.
type WebhookMessageAttemptService struct {
	options []option.RequestOption
}

// NewWebhookMessageAttemptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWebhookMessageAttemptService(opts ...option.RequestOption) (r WebhookMessageAttemptService) {
	r = WebhookMessageAttemptService{}
	r.options = opts
	return
}

// When a webhook message fails to send it will store the failure in an "attempt".
// You can view and analyse these. Webhooks messages that receive too many attempts
// (failures) will not be fired. You must first clear out old attempts and try
// again. This endpoint shows you the details of a single attempt. The ID of the
// attempt must match the corresponding webhook and webhook message.
func (r *WebhookMessageAttemptService) Get(ctx context.Context, attemptID int64, params WebhookMessageAttemptGetParams, opts ...option.RequestOption) (res *WebhookMessageAttemptGetResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/webhooks/%s/messages/%v/attempts/%v", url.PathEscape(params.ID), params.MessageID, attemptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// When a webhook message fails to send it will store the failure in an "attempt".
// You can view and analyse these. Webhook messages that receive too many attempts
// (failures) will not be sent again. You must first clear out old attempts before
// the message can go out again.
func (r *WebhookMessageAttemptService) List(ctx context.Context, messageID int64, params WebhookMessageAttemptListParams, opts ...option.RequestOption) (res *WebhookMessageAttemptListResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/webhooks/%s/messages/%v/attempts", url.PathEscape(params.ID), messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a webhook message attempt. If you delete all attempts for a webhook
// message, Firefly III will (once again) assume all is well with the webhook
// message and will try to send it again.
func (r *WebhookMessageAttemptService) Delete(ctx context.Context, attemptID int64, params WebhookMessageAttemptDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/webhooks/%s/messages/%v/attempts/%v", url.PathEscape(params.ID), params.MessageID, attemptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type WebhookAttempt struct {
	ID         string                   `json:"id" api:"required"`
	Attributes WebhookAttemptAttributes `json:"attributes" api:"required"`
	// Immutable value
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attributes  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookAttempt) RawJSON() string { return r.JSON.raw }
func (r *WebhookAttempt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookAttemptAttributes struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Internal log for this attempt. May contain sensitive user data.
	Logs string `json:"logs" api:"nullable"`
	// Webhook receiver response for this attempt, if any. May contain sensitive user
	// data.
	Response string `json:"response" api:"nullable"`
	// The HTTP status code of the error, if any.
	StatusCode int64     `json:"status_code" api:"nullable"`
	UpdatedAt  time.Time `json:"updated_at" format:"date-time"`
	// The ID of the webhook message this attempt belongs to.
	WebhookMessageID string `json:"webhook_message_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		Logs             respjson.Field
		Response         respjson.Field
		StatusCode       respjson.Field
		UpdatedAt        respjson.Field
		WebhookMessageID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookAttemptAttributes) RawJSON() string { return r.JSON.raw }
func (r *WebhookAttemptAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookMessageAttemptGetResponse struct {
	Data WebhookAttempt `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookMessageAttemptGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookMessageAttemptGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookMessageAttemptListResponse struct {
	Data []WebhookAttempt `json:"data" api:"required"`
	Meta Meta             `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookMessageAttemptListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookMessageAttemptListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookMessageAttemptGetParams struct {
	ID        string            `path:"id" api:"required" json:"-"`
	MessageID int64             `path:"messageId" api:"required" json:"-"`
	XTraceID  param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type WebhookMessageAttemptListParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookMessageAttemptListParams]'s query parameters as
// `url.Values`.
func (r WebhookMessageAttemptListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookMessageAttemptDeleteParams struct {
	ID        string            `path:"id" api:"required" json:"-"`
	MessageID int64             `path:"messageId" api:"required" json:"-"`
	XTraceID  param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
