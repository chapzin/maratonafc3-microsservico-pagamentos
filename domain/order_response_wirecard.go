package domain

type OrderResponse struct {
	ID                     string
	Status                 string
	CreatedAt              string
	Amount                 AmountResponse
	Payments               []PaymentWirecard
	Refunds                []Refund
	Entries                []Entrie
	Events                 []Event
	UpdatedAt              string
	Links                  Links
	PayCheckout            PayCheckout
	PayCreditCard          PayCreditCard
	PayBoleto              PayBoleto
	PayOnlineBankDebitItau PayOnlineBankDebitItau
}

type AmountResponse struct {
	Total          int32
	Fees           int32
	Refunds        int32
	Liquid         int32
	OtherReceivers int32
	SubTotals      SubTotalsResponse
}

type SubTotalsResponse struct {
	Items int32
}

type Event struct {
	CreatedAt   string
	Type        string
	Description string
}

type Links struct {
	Self     Self
	Checkout Checkout
}

type Self struct {
	Href string
}

type Checkout struct{}

type PayCheckout struct {
	RedirectHref string
}

type PayCreditCard struct {
	RedirectHref string
}

type PayBoleto struct {
	RedirectHref string
}

type PayOnlineBankDebitItau struct {
	RedirectHref string
}
