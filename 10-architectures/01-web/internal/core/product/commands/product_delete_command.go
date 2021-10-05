package commands

type DeleteProductCommand struct {
	Id string `json:"-" validate:"required"`
}
