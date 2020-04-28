package facade

import "testing"

func TestFacadeSuperMarket(t *testing.T) {

	supermarket := NewSuperMarket()

	supermarket.Sell(1, 3, 5)
	supermarket.Sell(2, 11, 30)

	supermarket.Sell(8, 8, 30)
}
