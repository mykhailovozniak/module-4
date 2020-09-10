package module_4

import "fmt"

type engine interface {
	start()
	stop()
}

type process struct {
	command command
}

type command interface {
	execute()
}

type startCommand struct {
	engine
}

type stopCommand struct {
	engine
}

func (c *startCommand) execute() {
	c.engine.start()
}

func (c * stopCommand) execute() {
	c.engine.stop()
}

type specificEngine struct {
	isLaunched bool
}

func (c *specificEngine) start() {
	c.isLaunched = true
	fmt.Println("Turning engine on")
}

func (c *specificEngine) stop() {
	c.isLaunched = false
	fmt.Println("Turning engine off")
}

func (b *process) init() {
	b.command.execute()
}

func CommandClentCode() {
	specificEngine := &specificEngine{}
	onCommand := &startCommand{
		engine: specificEngine,
	}

	offCommand := &stopCommand{
		engine: specificEngine,
	}

	startProcess := &process{
		command: onCommand,
	}

	startProcess.init()

	endProcess := &process{
		command: offCommand,
	}
	endProcess.init()
}
