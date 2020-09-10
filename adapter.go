package module_4

import "fmt"

type Target interface {
	Request() string
}

type Adaptee struct {
}

type Adapter struct {
	*Adaptee
}

func (a *Adaptee) SpecificRequest() string {
	return "Data from specific Request"
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

func (a *Adapter) Request() string {
	specificResponse := a.SpecificRequest()

	return specificResponse + " some data from adapter "
}

func AdapterClientCode() {
	adaptee := new(Adaptee)
	fmt.Println(adaptee.SpecificRequest())

	adapter := NewAdapter(&Adaptee{})
	req := adapter.Request()
	fmt.Println(req)
}
