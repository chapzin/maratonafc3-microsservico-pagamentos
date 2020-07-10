package domain

type PaymentWirecard struct {
	ID                  string
	Status              string
	Amount              AmountPayment
	InstallmentCount    int32
	DelayCapture        bool
	StatementDescriptor string
	FundingInstrument   FundingInstrumentPayment
	Fees                []Fee
	Event               []Event
	CancellationDetails CancellationDetails
	UpdatedAt           string
	CreatedAt           string
	Device              Device
	Links               LinksPayment
}

type AmountPayment struct {
	Refunds  int32
	Fees     int32
	Liquid   int32
	Currency string
	Total    int32
}

type FundingInstrumentPayment struct {
	Method          string
	CreditCard      CreditCardPayment
	Boleto          BoletoPayment
	OnlineBankDebit OnlineBankDebit
}

type CreditCardPayment struct {
	Hash             string
	Number           string
	ExpirationMounth int32
	ExpirationYear   int32
	Cvc              int32
	ID               string
	Brand            string
	First6           string
	Last4            string
	Holder           Holder
}

type BoletoPayment struct {
	LineCode         string
	ExpirationDate   string
	InstructionLines InstructionLines
	LogoURI          string
}

type InstructionLines struct {
	First  string
	Second string
	Third  string
}

type OnlineBankDebit struct {
	BankNumber     string
	BankName       string
	ExpirationDate string
	ReturnURI      string
}

type Fee struct {
	Type   string
	Amount int
}

type CancellationDetails struct {
	CancelledBy string
	Code        int32
	Description string
}

type Device struct {
	IP          string
	UserAgent   string
	FingerPrint string
	GeoLocation GeoLocation
	Self        Self
	Order       Order
	Checkout    Checkout
}

type GeoLocation struct {
	Latitude  float32
	Longitude float32
}

type Order struct {
	Title string
	Href  string
}

type LinksPayment struct {
	Self                   Self
	Order                  Order
	PayCheckout            PayCheckout
	PayCreditCard          PayCreditCard
	PayBoleto              PayBoleto
	PayOnlineBankDebitItau PayOnlineBankDebitItau
}
