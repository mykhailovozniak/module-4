package module_4

import "fmt"

type Flyweight struct {
	Name string
}

func NewFlyweight(name string) *Flyweight {
	return &Flyweight{name}
}

type FlyweightFactory struct {
	pool map[string] *Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{pool: make(map[string]*Flyweight)}
}

func (f *FlyweightFactory) GetFlyweight(name string) *Flyweight {
	flyweight, okay := f.pool[name]
	if !okay {
		flyweight = NewFlyweight(name)
		f.pool[name] = flyweight
	}
	return flyweight
}

func FlyweightClientCode() {
	flyweight := NewFlyweight("Chevrolet, Camaro2018, pink")
	fmt.Println(flyweight.Name)

	factory := NewFlyweightFactory()

	factory.GetFlyweight("Chevrolet, Camaro2018, pink")
	fmt.Println("Case when we already have object in storage ")
	fmt.Println(factory.GetFlyweight("Chevrolet, Camaro2018, pink"))

	fmt.Println("Case when does not have object in storage")
	fmt.Println(factory.GetFlyweight("Mercedes Benz, C300, black"))
}
