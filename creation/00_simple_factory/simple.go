package simplefactory

import "fmt"

type schoolmember int

const (
	student schoolmember = iota
	teacher
)

//Mouth is interface that people can Say words
type Mouth interface {
	Say(name string) string
}

//New return instance object that have `Mouth` and can speak
func New(t schoolmember) Mouth {
	switch t {
	case student:
		return &studentType{}
	case teacher:
		return &teacherType{}
	}
	return nil
}

//teacherType is one of Mouth implement
type teacherType struct{}

//Say teacher's name
func (*teacherType) Say(name string) string {
	return fmt.Sprintf("I am Teacher: %s", name)
}

//studentType is another Mouth implement
type studentType struct{}

//Say student's name
func (*studentType) Say(name string) string {
	return fmt.Sprintf("I am Student： %s", name)
}
