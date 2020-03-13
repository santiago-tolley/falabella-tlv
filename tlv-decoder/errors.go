package decoder

type ErrEmptyInput struct{}

func (e ErrEmptyInput) Error() string {
	return "the input is empty"
}

type ErrInvalidFormat struct{}

func (e ErrInvalidFormat) Error() string {
	return "the field does not have a valid format"
}

type ErrInvalidType struct{}

func (e ErrInvalidType) Error() string {
	return "the field does not have a valid type"
}

type ErrInvalidLength struct{}

func (e ErrInvalidLength) Error() string {
	return "the field does not have a valid length"
}
