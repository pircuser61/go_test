package logic

import (
	"errors"
	"reflect"

	"github.com/pircuser61/go_test/internal/core/doer"
)

type Logic struct {
	d doer.Doer
}

var ErrEmptyString = errors.New("empty string")

type ValidationError struct {
	message string
}

func (e ValidationError) Error() string { return e.message }

func (e ValidationError) Is(err error) bool {
	//fmt.Println(reflect.TypeOf(err), "==", reflect.TypeOf(err).Kind())
	return reflect.TypeOf(err) == reflect.TypeOf(e)
}

func New(d doer.Doer) Logic {
	return Logic{d}
}

func (i Logic) Do(in string) (string, error) {
	if in == "" {
		return "", ValidationError{"empty input"}
	}
	if len(in) < 3 {
		return "", errors.New("input too short") //ValidationError{"input too short"}
	}

	result, err := i.d.Do(in)
	if err != nil {
		return "", err
	}
	if result == "" {
		return "", ErrEmptyString
	}
	return "ok", nil
}
