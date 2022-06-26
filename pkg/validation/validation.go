package validation

type InternalValidator interface {
	ValidateParam(param interface{}) error
}
