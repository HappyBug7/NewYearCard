package controller

type Controller struct {
	Card
}

func New() *Controller {
	Controller := &Controller{}
	return Controller
}
