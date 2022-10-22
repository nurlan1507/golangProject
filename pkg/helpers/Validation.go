package helpers

type Validation struct {
	Errors map[string]string
}

func NewValidation() *Validation {
	return &Validation{
		Errors: map[string]string{},
	}
}

// Check в эту функцию кидаем функцию (напимер isValidPassword(asdsad), а там уже на основе ИсВалид оно решает доавлять в мапу или нет)
func (v *Validation) Check(isValid bool, field string, message string) {
	if isValid == true {
		return
	} else {
		v.Errors[field] = message
	}
}

// IsValid if map is empty then user passed the validation
func (v *Validation) IsValid() bool {
	if len(v.Errors) == 0 {
		return true
	}
	return false
}
