package module_4

import (
	"fmt"
	"log"
)

const (
	BOOK="BOOK"
	TV="TV"
	KEYBOARD="KEYBOARD"
)

type HomeProductCreator interface {
	CreateProduct(named string) HomeProduct
}

type HomeProduct interface {
	Use() string
}

type BookProduct struct {
	action string
}

type TVProduct struct {
	action string
}

type KeyBoardProduct struct {
	action string
}

func (p *BookProduct) Use() string {
	return p.action
}

func (p *TVProduct) Use() string {
	return p.action
}

func (p *KeyBoardProduct) Use() string {
	return p.action
}

type ProductCreator struct {

}

func NewHomeProductCreator() HomeProductCreator {
	return &ProductCreator{}
}

func (p *ProductCreator) CreateProduct(named string) HomeProduct {
	var homeProduct HomeProduct

	switch named {
	case BOOK:
		homeProduct = &BookProduct{named}
	case TV:
		homeProduct = &TVProduct{named}
	case KEYBOARD:
		homeProduct = &KeyBoardProduct{named}
	default:
		log.Fatalln("Unknown product chosen")
	}

	return homeProduct
}

func factoryClientCode() {
	factory := NewHomeProductCreator()

	homeProducts := []HomeProduct{
		factory.CreateProduct(BOOK),
		factory.CreateProduct(TV),
		factory.CreateProduct(KEYBOARD),
	}

	for _, homeProduct := range homeProducts {
		fmt.Println(homeProduct.Use())
	}
}
