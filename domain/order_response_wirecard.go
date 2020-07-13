package domain

type OrderResponse struct {
	ID                     string                 `json:"id"`
	Status                 string                 `json:"status"`
	CreatedAt              string                 `json:"createdAt"`
	Amount                 AmountResponse         `json:"amount"`
	Payments               []PaymentWirecard      `json:"payments"`
	Refunds                []Refund               `json:"refunds"`
	Entries                []Entrie               `json:"entries"`
	Events                 []Event                `json:"events"`
	UpdatedAt              string                 `json:"updatedAt"`
	Links                  Links                  `json:"_links"`
	PayCheckout            PayCheckout            `json:"payCheckout"`
	PayCreditCard          PayCreditCard          `json:"payCreditCard"`
	PayBoleto              PayBoleto              `json:"payBoleto"`
	PayOnlineBankDebitItau PayOnlineBankDebitItau `json:"payOnlineBankDebitItau"`
}

type AmountResponse struct {
	Total          int32             `json:"total"`
	Fees           int32             `json:"fees"`
	Refunds        int32             `json:"refunds"`
	Liquid         int32             `json:"liquid"`
	OtherReceivers int32             `json:"otherReceivers"`
	SubTotals      SubTotalsResponse `json:"subtotals"`
}

type SubTotalsResponse struct {
	Items int32 `json:"items"`
}

type Event struct {
	CreatedAt   string `json:"createdAt"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Links struct {
	Self     Self     `json:"self"`
	Checkout Checkout `json:"checkout"`
}

type Self struct {
	Href string `json:"href"`
}

type Checkout struct{}

type PayCheckout struct {
	RedirectHref string `json:"redirectHref"`
}

type PayCreditCard struct {
	RedirectHref string `json:"redirectHref"`
}

type PayBoleto struct {
	RedirectHref string `json:"redirectHref"`
}

type PayOnlineBankDebitItau struct {
	RedirectHref string `json:"redirectHref"`
}
