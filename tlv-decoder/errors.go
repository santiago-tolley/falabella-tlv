package decoder

type ErrEmptyInput struct{}

func (e ErrEmptyInput) Error() string {
	return "the input is empty"
}

type ErrInvalidFormat struct{}

func (e ErrInvalidFormat) Error() string {
	return "the field does not have a valid format"
}
