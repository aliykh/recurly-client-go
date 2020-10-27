// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type SubscriptionChangeShippingCreate struct {
	Params `json:"-"`

	// The id of the shipping method used to deliver the subscription. To remove shipping set this to `null` and the `amount=0`. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodId *string `json:"method_id,omitempty"`

	// The code of the shipping method used to deliver the subscription. To remove shipping set this to `null` and the `amount=0`. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodCode *string `json:"method_code,omitempty"`

	// Assigns the subscription's shipping cost. If this is greater than zero then a `method_id` or `method_code` is required.
	Amount *float64 `json:"amount,omitempty"`
}

func (attr *SubscriptionChangeShippingCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
