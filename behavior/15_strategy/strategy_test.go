package strategy

import "testing"

func TestStoreMoney(t *testing.T) {

	bank := Bank{&MainLandCitizen{"Miss White"}}
	ctx := &StoreContext{
		Kind:   "RMB",
		CardID: "12345678921",
		Money:  10000,
	}
	bank.AccountUserMoney(ctx)

	hkUser := &HongKongCitizen{"Miss Black"}

	bank.Recept(hkUser)

	ctx = &StoreContext{
		Kind:   "HK",
		CardID: "987345678456",
		Money:  8723,
	}
	bank.AccountUserMoney(ctx)
}
