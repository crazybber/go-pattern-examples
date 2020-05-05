package interpreter

import "testing"

func TestMeetingActionSignInterpreter(t *testing.T) {

	p := &SignParser{}

	p.Parse("rose -> tom")
	p.Result()

	p.Parse("rose <-> tom")
	p.Result()

	p.Parse("rose <- tom")

	p.Result()

	//should error
	p.Parse("rose + tom")

	p.Result()

}

func TestCalculatorInterpreter(t *testing.T) {
	p := &Parser{}
	p.Parse("1 + 2 + 3 - 4 + 5 - 6")
	res := p.Result().Interpret()
	expect := 1
	if res != expect {
		t.Fatalf("expect %d got %d", expect, res)
	}
}
