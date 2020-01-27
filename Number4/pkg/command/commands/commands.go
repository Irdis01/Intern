package commands

type car interface {
	StartEngine()
	StopEngine() error
	TurnOnHandbrake() error
	TurnOffHandbrake()
	StartMove() error
	Accelerate() error
	Slowdown() error
}

// Commands интерйфес комманд
type Commands interface {
	Execute() error
	Cancel() error
}
