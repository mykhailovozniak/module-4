package module_4

import (
	"fmt"
	"strings"
)

type Parser struct {
}

func (p *Parser) Parse() string {
	return "parsing data"
}

type Process struct {
}

func (p *Process) Process() string {
	return "processing data"
}

type Storage struct {
}

func (s *Storage) Store() string {
	return "storing data"
}

type Application struct {
	parser *Parser
	process *Process
	storage *Storage
}

func (a *Application) Start() string{
	stages := []string{
		a.parser.Parse(),
		a.process.Process(),
		a.storage.Store(),
	}

	return strings.Join(stages, ", ")
}

func FacadeClintCode() {
	app := new(Application)

	fmt.Println(app.Start())
}
