package simplefactory

import "testing"

//TestTeacher test get New with factory
func TestTeacher(t *testing.T) {
	te := New(teacher)
	name := te.Say("Tom")
	if name != "I am Teacher: Tom" {
		t.Fatal("Teacher test fail")
	}
}

func TestStudent(t *testing.T) {
	st := New(student)
	name := st.Say("Tom")
	if name != "I am student: Tom" {
		t.Fatal("student test fail")
	}
}
