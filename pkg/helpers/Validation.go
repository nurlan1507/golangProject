package helpers

type Validation struct {
	Errors map[string]string
}

func NewValidation() *Validation {
	return &Validation{
		Errors: make(map[string]string),
	}
}

// Check в эту функцию кидаем функцию (напимер isValidPassword(asdsad), а там уже на основе ИсВалид оно решает доавлять в мапу или нет)
func (v *Validation) Check(ok bool, field string, message string) {
	if !ok {
		v.Errors[field] = message
		return
	}
	return
}

func (v *Validation) CheckQuestions(ok bool, field string, message string) bool {
	if !ok {
		v.Errors[field] = message
		return false
	}
	return true
}

// IsValid if map is empty then user passed the validation
func (v *Validation) Valid() bool {
	if len(v.Errors) == 0 {
		return true
	}
	return false
}
