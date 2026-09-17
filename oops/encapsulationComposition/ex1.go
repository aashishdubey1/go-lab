package encapsulationcomposition

import "fmt"

// Vehicle Embedding
// Define Engine{HorsePower int} with a Start() string method, then Car and Motorcycle that both embed Engine.
// Override Start() on Car only, and demonstrate that Motorcycle still uses the promoted version.

type Engine struct {
	HorsePower int
}

func (e *Engine) Start() string {
	return fmt.Sprintf("%d HorsePower engine is Starting", e.HorsePower)
}

type Car struct {
	Engine
	model string
	color string
}

func (c *Car) Start() string {
	return fmt.Sprintf("%s %s Car is Starting", c.color, c.model)
}

type Motorcycle struct {
	Engine
	name  string
	color string
}

func RunEx1() {
	engine := Engine{HorsePower: 800}

	car := Car{Engine: engine, color: "Black", model: "Suv 54"}

	bike := Motorcycle{Engine: engine, color: "Black", name: "Gt 650"}

	fmt.Println(engine.Start())
	fmt.Println(car.Start())
	fmt.Println(bike.Start())
}

func main() {
	RunEx1()
}
