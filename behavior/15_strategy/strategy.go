package strategy

import "fmt"

//money kind
const (
	RMB = "RMB"
	HK  = "HK"
)

//StoreContext for Store 要包存钱的上下文信息
type StoreContext struct {
	Kind, CardID string
	Money        int
}

//IStore 要实现的存钱接口
type IStore interface {
	Store(*StoreContext)
}

//MainLandCitizen 大陆居民
type MainLandCitizen struct{ Name string }

//Store Money  to bank
func (m *MainLandCitizen) Store(ctx *StoreContext) {
	fmt.Println("i am: ", m.Name, "i want to store: ", ctx.Money, ctx.Kind, "to: ", ctx.CardID)
}

//HongKongCitizen 香港居民
type HongKongCitizen struct{ Name string }

//Store Money  to bank
func (h *HongKongCitizen) Store(ctx *StoreContext) {
	fmt.Println("i am: ", h.Name, "i want to store: ", ctx.Money, ctx.Kind, "to: ", ctx.CardID)
}

//Bank handle moneyholder
type Bank struct {
	moneyHolder IStore
}

//Recept a user
func (b *Bank) Recept(moneyHolder IStore) {
	b.moneyHolder = moneyHolder
	fmt.Println("Bank: ", "Recept a New User")
}

//AccountUserMoney 动态替换的过程在这里,这里调用任何实现了Store的接口对象
//AccountUserMoney to handle User's Money
func (b *Bank) AccountUserMoney(ctx *StoreContext) {
	b.moneyHolder.Store(ctx)
	fmt.Println("Bank: ", "Processing Store", ctx.Money, ctx.Kind, "to: ", ctx.CardID)
}
