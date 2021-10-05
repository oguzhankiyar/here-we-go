package queries

type FindProductByIdQuery struct {
	Id string `json:"-" validate:"required"`
}
