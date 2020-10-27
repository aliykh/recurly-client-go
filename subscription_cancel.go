// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type SubscriptionCancel struct {
	Params `json:"-"`

	// The timeframe parameter controls when the expiration takes place. The `bill_date` timeframe causes the subscription to expire when the subscription is scheduled to bill next. The `term_end` timeframe causes the subscription to continue to bill until the end of the subscription term, then expire.
	Timeframe *string `json:"timeframe,omitempty"`
}

func (attr *SubscriptionCancel) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
