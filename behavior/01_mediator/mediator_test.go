package mediator

import (
	"fmt"
	"testing"
)

func TestMediator(t *testing.T) {

	med := &Mediator{}
	landlord := &Landlord{}
	//登记房源信息
	med.RegisterRoom(landlord)
	tenant := &Tenant{}
	//向中介租房
	med.RentOutRoom(tenant)

	//房东收租
	landlord.CollectRent(med)
	//租客要求修理
	tenant.AskRepair(med)

}

func TestClassCompose(t *testing.T) {

	med := &Mediator{Person: Person{Name: "mediator", WalletAssets: 1001}}

	landlord := &Landlord{Person: Person{Name: "landlord", WalletAssets: 2000}, RentAccout: 500}

	tenant := &Tenant{Person: Person{Name: "tenant", WalletAssets: 500}, furniture: "desk"}

	fmt.Println("mediator", med)
	fmt.Println("landlord", landlord.Name)
	fmt.Println("tenant", tenant.furniture)
}
