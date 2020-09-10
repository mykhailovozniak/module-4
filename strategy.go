package module_4

import "fmt"

type HeadersStrategy interface {
	extractHeaders()
}

type DevelopmentHeaders struct {
}

type ProductionHeaders struct {
}

type AppContext struct {
	strategy HeadersStrategy
}

func NewDevelopmentHeaders() HeadersStrategy {
	return &DevelopmentHeaders{}
}

func NewProductionHeaders() HeadersStrategy {
	return &ProductionHeaders{}
}

func (DevelopmentHeaders) extractHeaders() {
	fmt.Println("Extraction auth header from Authorization")
}

func (ProductionHeaders) extractHeaders() {
	fmt.Println("Extraction auth header from x-api-custom")
}

func NewContext() *AppContext {
	return &AppContext{}
}

func (a *AppContext) SetStrategy(s HeadersStrategy) {
	a.strategy = s
}

func (a *AppContext) ExtractHeaders() {
	a.strategy.extractHeaders()
}

func StrategyClientCode() {
	context := NewContext()

	prodHeadersStrategy := NewProductionHeaders()
	devHeadersStrategy := NewDevelopmentHeaders()

	context.SetStrategy(prodHeadersStrategy)
	context.ExtractHeaders()

	context.SetStrategy(devHeadersStrategy)
	context.ExtractHeaders()
}
