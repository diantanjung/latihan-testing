package dependencyinjection

type Engine interface {
	MaxSpeed() int
}

func NewCar(engine Engine) *Car {
	return &Car{
		Engine: engine,
	}
}

type Car struct {
	Engine Engine
}

func (c Car) Speed() int {
	if c.Engine.MaxSpeed() < 10 {
		return 10
	}
	defaultSpeed := 60
	if defaultSpeed < c.Engine.MaxSpeed() {
		return c.Engine.MaxSpeed()
	}

	return defaultSpeed
}

type DefaultEngine struct{}

func (e DefaultEngine) MaxSpeed() int {
	return 70
}

type TurboEngine struct{}

func (e TurboEngine) MaxSpeed() int {
	return 100
}
