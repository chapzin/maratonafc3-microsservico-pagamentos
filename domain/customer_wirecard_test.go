package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustomerValidate(t *testing.T) {
	c := NewCustomer()
	err := c.Validate()

	require.Error(t, err)
}

func TestCustomer(t *testing.T) {
	c := NewCustomer()

	c.OwnID = "abcd"
	c.Email = "chapzin@gmail.com"
	c.FullName = "Ricardo Augusto Barroso Gomes"

	err := c.Validate()

	require.Nil(t, err)
	require.NotNil(t, c)
}

func TestCustomerEmail(t *testing.T) {
	c := NewCustomer()

	c.OwnID = "abcd"
	c.Email = "testeteste.com"
	c.FullName = "Ricardo Augusto Barroso Gomes"

	err := c.Validate()

	require.Error(t, err)

}
