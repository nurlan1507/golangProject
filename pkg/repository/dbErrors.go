package repository

import (
	"errors"
	"fmt"
)

var ErrNoRecord = errors.New("model: no matching record found")

func ErrDuplicate(field string) error {
	return errors.New(fmt.Sprintf("%v already exists"))
}
