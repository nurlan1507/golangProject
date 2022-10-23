package helpers

import "fmt"

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
	fmt.Println(ok)
	if !ok {
		v.Errors[field] = message
		return
	}
	return
}

// IsValid if map is empty then user passed the validation
func (v *Validation) Valid() bool {
	if len(v.Errors) == 0 {
		fmt.Println(v.Errors)
		return true
	}
	return false
}
