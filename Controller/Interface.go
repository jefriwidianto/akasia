package Controller

type controllerInterface interface {
	ProductInterface
}

type Controller struct {
	controllerInterface
}
