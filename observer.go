package module_4

import "fmt"

type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}

type observer interface {
	update(string)
	getID() string
}

func newProduct(name string) *product {
	return &product{
		name: name,
	}
}

type product struct {
	observerList []observer
	name         string
	available      bool
}

func (p *product) updateAvailability() {
	p.available = true
	p.notifyAll()
	fmt.Println("Product is available for buying")
}

func (p *product) register(o observer) {
	p.observerList = append(p.observerList, o)
}

func (p *product) deregister(o observer) {
	p.observerList = removeFromslice(p.observerList, o)
}

func (p *product) notifyAll() {
	for _, observer := range p.observerList {
		observer.update(p.name)
	}
}

func removeFromslice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

type customer struct {
	id string
}

func (c *customer) update(productName string) {
	fmt.Println("Notify customers about product " + c.id + " " + productName)
}

func (c *customer) getID() string {
	return c.id
}

func ObserverClientCode() {
	tv := newProduct("apple tv")
	observerFirst := &customer{id: "john.doe@mail.com"}
	observerSecond := &customer{id: "bob.bob@mail.com"}
	tv.register(observerFirst)
	tv.register(observerSecond)
	tv.updateAvailability()
}
