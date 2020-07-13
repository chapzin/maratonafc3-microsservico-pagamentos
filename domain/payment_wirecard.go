package domain

type PaymentWirecard struct {
	ID                  string                   `json:"id"`
	Status              string                   `json:"status"`
	Amount              AmountPayment            `json:"amount"`
	InstallmentCount    int32                    `json:"installmentCount"`
	DelayCapture        bool                     `json:"delayCapture"`
	StatementDescriptor string                   `json:"statementDescriptor"`
	FundingInstrument   FundingInstrumentPayment `json:"fundingInstrument"`
	Fees                []Fee                    `json:"fees"`
	Event               []Event                  `json:"event"`
	CancellationDetails CancellationDetails      `json:"cancellationDetails"`
	UpdatedAt           string                   `json:"updatedAt"`
	CreatedAt           string                   `json:"createdAt"`
	Device              Device                   `json:"device"`
	Links               LinksPayment             `json:"_links"`
}

type AmountPayment struct {
	Refunds  int32  `json:"refunds"`
	Fees     int32  `json:"fees"`
	Liquid   int32  `json:"liquid"`
	Currency string `json:"currency"`
	Total    int32  `json:"total"`
}

type FundingInstrumentPayment struct {
	Method          string            `json:"method"`
	CreditCard      CreditCardPayment `json:"creditCard"`
	Boleto          BoletoPayment     `json:"boleto"`
	OnlineBankDebit OnlineBankDebit   `json:"onlineBankDebit"`
}

type CreditCardPayment struct {
	Hash             string `json:"hash"`
	Number           string `json:"number"`
	ExpirationMounth int32  `json:"expirationMounth"`
	ExpirationYear   int32  `json:"expirationYear"`
	Cvc              int32  `json:"cvc"`
	ID               string `json:"id"`
	Brand            string `json:"brand"`
	First6           string `json:"first6"`
	Last4            string `json:"last4"`
	Holder           Holder `json:"holder"`
}

type BoletoPayment struct {
	LineCode         string           `json:"lineCode"`
	ExpirationDate   string           `json:"expirationDate"`
	InstructionLines InstructionLines `json:"instructionLines"`
	LogoURI          string           `json:"logoUri"`
}

type InstructionLines struct {
	First  string `json:"first"`
	Second string `json:"second"`
	Third  string `json:"third"`
}

type OnlineBankDebit struct {
	BankNumber     string `json:"bankNumber"`
	BankName       string `json:"bankName"`
	ExpirationDate string `json:"expirationDate"`
	ReturnURI      string `json:"returnUri"`
}

type Fee struct {
	Type   string `json:"type"`
	Amount int    `json:"amount"`
}

type CancellationDetails struct {
	CancelledBy string `json:"cancelledBy"`
	Code        int32  `json:"code"`
	Description string `json:"description"`
}

type Device struct {
	IP          string      `json:"ip"`
	UserAgent   string      `json:"userAgent"`
	FingerPrint string      `json:"fingerPrint"`
	GeoLocation GeoLocation `json:"geoLocation"`
	Self        Self        `json:"self"`
	Order       Order       `json:"order"`
	Checkout    Checkout    `json:"checkout"`
}

type GeoLocation struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Order struct {
	Title string `json:"title"`
	Href  string `json:"href"`
}

type LinksPayment struct {
	Self                   Self                   `json:"self"`
	Order                  Order                  `json:"order"`
	PayCheckout            PayCheckout            `json:"payCheckout"`
	PayCreditCard          PayCreditCard          `json:"payCreditCard"`
	PayBoleto              PayBoleto              `json:"payBoleto"`
	PayOnlineBankDebitItau PayOnlineBankDebitItau `json:"payOnlineBankDebitItau"`
}
