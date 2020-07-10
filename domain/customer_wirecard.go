package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type CustomerWirecard struct {
	ID                string            `json:"id" valid:"-"`
	OwnID             string            `json:"ownId" valid:"required~O id do cliente é obrigatorio"`
	FullName          string            `json:"fullname" valid:"required~O nome completo é obrigatorio"`
	Email             string            `json:"email" valid:"email, required~O email não é valido"`
	BirthDate         string            `json:"birthDate" valid:"-"`
	TaxDocument       TaxDocument       `json:"taxDocument" valid:"-"`
	Phone             Phone             `json:"phone" valid:"-"`
	ShippingAddress   ShippingAddress   `json:"shippingAddress" valid:"-"`
	FundingInstrument FundingInstrument `json:"fundingInstrument" valid:"-"`
}

func NewCustomer() *CustomerWirecard {
	return &CustomerWirecard{}
}

type TaxDocument struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Phone struct {
	CountryCode int32 `json:"countryCode"`
	AreaCode    int32 `json:"areaCode"`
	Number      int32 `json:"number"`
}

type ShippingAddress struct {
	Street       string `json:"street"`
	StreetNumber string `json:"streetNumber"`
	Complement   string `json:"complement"`
	District     string `json:"district"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      int32  `json:"zipCode"`
}

type FundingInstrument struct {
	Method     string     `json:"method"`
	CreditCard CreditCard `json:"creditCard"`
}

type CreditCard struct {
	Hash            string `json:"hash"`
	Number          string `json:"number"`
	ExpirationMonth int32  `json:"expirationMonth"`
	ExpirationYear  int32  `json:"expirationYear"`
	Cvc             string `json:"cvc"`
	Holder          Holder `json:"holder"`
}

type Holder struct {
	FullName       string          `json:"fullname"`
	BirthDate      *time.Time      `json:"birthdate"`
	TaxDocument    TaxDocument     `json:"taxDocument"`
	Phone          Phone           `json:"phone"`
	BillingAddress ShippingAddress `json:"billingAddress"`
}

func (c *CustomerWirecard) Validate() error {
	_, err := govalidator.ValidateStruct(c)

	if err != nil {
		return err
	}

	return nil
}
