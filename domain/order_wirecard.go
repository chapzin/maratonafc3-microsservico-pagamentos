package domain

import "github.com/asaskevich/govalidator"

type OrderInterface interface{}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Order struct {
	ID                  string              `valid:"-" json:"id"`
	OwnID               string              `valid:"required" json:"ownId"`
	Amount              Amount              `valid:"-" json:"amount"`
	Items               []Item              `valid:"-" json:"items"`
	CheckOutPreferences CheckOutPreferences `valid:"-" json:"checkoutPreferences"`
	ShippingAddress     ShippingAddress     `valid:"required" json:"shippingAddress"`
	Customer            CustomerWirecard    `valid:"-" json:"customer"`
	Receivers           []Receiver          `valid:"-" json:"receivers"`
}

func NewOrder() *Order {
	return &Order{}
}

type Amount struct {
	Currency  string    `valid:"type(string),required" json:"currency"`
	SubTotals SubTotals `valid:"-" json:"subtotals"`
}

type SubTotals struct {
	Shipping int32 `valid:"type(int),required" json:"shipping"`
	Addition int32 `valid:"type(int)" json:"addition"`
	Discount int32 `valid:"type(int)" json:"discount"`
}

type Item struct {
	Product  string `valid:"required" json:"product"`
	Category string `valid:"-" json:"category"`
	Quantity int32  `valid:"type(int)" json:"quantity"`
	Detail   string `valid:"-" json:"detail"`
	Price    int32  `valid:"type(int),required" json:"price"`
}

type CheckOutPreferences struct {
	RedirectUrls string        `valid:"-" json:"redirectUrls"`
	URLSuccess   string        `valid:"-" json:"urlSuccess"`
	URLFailure   string        `valid:"-" json:"urlFailure"`
	Installments []Installment `valid:"-" json:"installments"`
}

type Installment struct {
	Quantity []int32 `valid:"-" json:"quantity"`
	Discount int32   `valid:"-" json:"discount"`
	Addition int32   `valid:"-" json:"addition"`
}

type Receiver struct {
	Type        string         `valid:"-" json:"type"`
	FeePayor    bool           `valid:"-" json:"feePayor"`
	MoipAccount MoipAccount    `valid:"-" json:"moipAccount"`
	Amount      AmountRecivers `valid:"-" json:"amount"`
}

type MoipAccount struct {
	Login    string `valid:"-" json:"login"`
	FullName string `valid:"-" json:"fullname"`
	ID       string `valid:"-" json:"id"`
}

type AmountRecivers struct {
	Refunds int32 `valid:"type(int)" json:"refunds"`
	Fees    int32 `valid:"type(int)" json:"fees"`
	Total   int32 `valid:"type(int)" json:"total"`
}

func (o *Order) Validate() error {
	_, err := govalidator.ValidateStruct(o)
	if err != nil {
		return err
	}

	return nil
}
