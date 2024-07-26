package web

type InputsCreateRequest struct {
	Nums   []int `validate:"required,dive,min=1,max=100"`
	Target int   `validate:"required"`
}
