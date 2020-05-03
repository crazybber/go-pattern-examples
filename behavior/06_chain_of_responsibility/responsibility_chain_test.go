package chain

import "testing"

func TestChainsOfResponsibilityInApprovementList(t *testing.T) {

	request := FeeRequest{
		Name:          "sophia",
		Mount:         1e+9,
		RequiredLevel: 8,
	}

	flow := &FeeRequestChainFlow{}

	gm1 := &GM{level: 7}

	cfo1 := &CFO{level: 9}

	flow.AddApprover(gm1)
	flow.AddApprover(cfo1)

	flow.RunApprovalFlow(request)

	t.Log("------second flow--------")
	request = FeeRequest{
		Name:          "peter",
		Mount:         1e+13,
		RequiredLevel: 8,
	}

	ceo := &CEO{}
	flow.AddApprover(ceo)

	flow.RunApprovalFlow(request)

}

func TestChainsOfResponsibilityInApprovementLink(t *testing.T) {

	request := FeeRequest{
		Name:          "sophia",
		Mount:         1e+11,
		RequiredLevel: 8,
	}

	gm := &GM{level: 7}

	cfo := &CFO{level: 9}

	ceo := &CEO{}

	gm.SetNext(cfo)

	cfo.SetNext(ceo)

	gm.HandleApproval(request)

}
