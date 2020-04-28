package facade

import (
	"fmt"
)

//ISupertMarketVendor is facade interface of facade package
type ISupertMarketVendor interface {
	Sell(count int)
}

//SaltVendor 盐供应商
type SaltVendor struct{}

//MilkVendor 牛奶供应商
type MilkVendor struct{}

//RiceVendor 大米供应商
type RiceVendor struct{}

//三家供应商都能直接卖东西

//Sell 卖完了
func (SaltVendor) Sell(count int) {

	if count > 5 {
		fmt.Println("Salt out")

	}
	fmt.Println("Milk got")

}

//Sell 卖完了
func (MilkVendor) Sell(count int) {

	if count > 20 {
		fmt.Println("Milk out")

	}
	fmt.Println("Milk got")

}

//Sell 卖完了
func (RiceVendor) Sell(count int) {

	if count > 10 {
		fmt.Println("Rice out")

	}
	fmt.Println("Rice got")

}

//SuperMarket is facade implement
//SuperMarket is Facade object
//SuperMarket 具有集中进货能力
type SuperMarket struct {
	saltsVendor ISupertMarketVendor
	milksVendor ISupertMarketVendor
	ricesVendor ISupertMarketVendor
}

//ISupertMarket market can do
type ISupertMarket interface {
	Sell(salt, milk, rice int)
}

//Sell 集中购买
func (s *SuperMarket) Sell(salt, milk, rice int) {
	s.saltsVendor.Sell(salt)
	s.milksVendor.Sell(milk)
	s.ricesVendor.Sell(rice)
}

//NewSuperMarket get a market
func NewSuperMarket() ISupertMarket {

	return &SuperMarket{
		saltsVendor: MilkVendor{},
		milksVendor: MilkVendor{},
		ricesVendor: RiceVendor{},
	}
}
