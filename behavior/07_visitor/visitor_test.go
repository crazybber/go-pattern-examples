package visitor

import "testing"

func TestVisitor(t *testing.T) {

	//汽油提供给，制衣工厂
	g := gas{density: 100}

	//柴油，提供给军工厂
	d := diesel{energy: 897}

	//购买石油的客户
	m := &militaryFactory{}

	c := &clothFactory{}

	g.Accept(c)

	d.Accept(m)

}

func ExampleAnalysis() {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewIndividualCustomer("bob"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Accept(&AnalysisVisitor{})
	// Output:
	// analysis enterprise customer A company
	// analysis enterprise customer B company
}
