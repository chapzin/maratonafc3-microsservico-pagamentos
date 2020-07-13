package domain

type Refund struct {
	ID                  string              `json:"id"`
	Status              string              `json:"status"`
	Amount              AmountRefund        `json:"amount"`
	Type                string              `json:"type"`
	RefundingInstrument RefundingInstrument `json:"refundingInstrument"`
	Events              []Event             `json:"events"`
	CreatedAt           string              `json:"createdAt"`
	Links               LinksRefund         `json:"_links"`
}

type AmountRefund struct {
	Fees     int32  `json:"fees"`
	Currency string `json:"currency"`
	Total    int32  `json:"total"`
}

type RefundingInstrument struct {
	Method     string            `json:"method"`
	CreditCard CreditCardPayment `json:"creditCard"`
}

type LinksRefund struct {
	Self    Self        `json:"self"`
	Order   Order       `json:"order"`
	Payment PaymentLink `json:"payment"`
}

type PaymentLink struct {
	Title string `json:"title"`
	Href  string `json:"href"`
}
