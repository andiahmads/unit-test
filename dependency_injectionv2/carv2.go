package dependency_injectionv2

type Speeder interface {
	MaxSpeed() int
}

//ini dependency_injection
func NewCar(speeder Speeder) *Car {
	return &Car{
		Speeder: speeder,
	}
}

type Car struct {
	Speeder Speeder
}

func (c Car) Speed() int {
	defaultSpeed := 100
	//todo: fake Test
	if c.Speeder.MaxSpeed() < 10 {
		return 20
	}
	if defaultSpeed > c.Speeder.MaxSpeed() {
		return c.Speeder.MaxSpeed()
	}
	return defaultSpeed
}

type DefaultEngine struct{}

func (e DefaultEngine) MaxSpeed() int {
	return 50
}

type TurboEngine struct{}

func (e TurboEngine) MaxSpeed() int {
	return 100
}
