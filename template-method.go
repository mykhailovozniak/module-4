package module_4

import "fmt"

type IEmployee interface {
	GetUp()
	GoToWork()
	Relax()
	Sleep()
}

type Employee struct {
	IEmployee
}

func NewEmployee(i IEmployee) *Employee {
	return &Employee{i}
}

type Person struct {
}

func (e *Employee) LiveOneDay() {
	e.GetUp()
	e.GoToWork()
	e.Relax()
	e.Sleep()
}

func (p *Person) GetUp() {
	fmt.Println("You are getting up")
}

func (p *Person) GoToWork() {
	fmt.Println("You are going to work")
}

func (p *Person) Relax() {
	fmt.Println("You are relaxing")
}

func (p *Person) Sleep()  {
	fmt.Println("You are sleeping")
}

func TemplateMethodClientCode() {
	person := &Person{}
	employee := NewEmployee(person)

	employee.LiveOneDay()
}
