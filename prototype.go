package module_4

import "fmt"

type Prototype interface {
	Clone() Prototype
	GetTitle() string
}

type Book struct {
	title string
}

func NewBook(title string) Prototype {
	return &Book{
		title: title,
	}
}

func (p *Book) Clone() Prototype {
	return &Book{p.title}
}

func (p *Book) GetTitle() string {
	return p.title
}

func PrototypeClientCode()  {
	product := NewBook("Some Great Book  Title")
	clonedProduct := product.Clone()


	fmt.Println(product.GetTitle() == clonedProduct.GetTitle())
}
