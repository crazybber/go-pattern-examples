package mediator

import "testing"

func TestMediator(t *testing.T) {

	med := &mediator{}

	landlord := &Landlord{}

	med.publishRoom(landlord)

	tenant := &Tenant{}

}
