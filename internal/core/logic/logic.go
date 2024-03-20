package logic

import (
	"errors"
	"reflect"
	"sync"
	"time"

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

func (i Logic) DoParallel(in string, threadCount int) error {

	wg := sync.WaitGroup{}
	wg.Add(4)
	for ; threadCount > 0; threadCount-- {
		go func() {
			time.Sleep(time.Second)
			i.d.Do(in)
			wg.Done()
		}()
	}
	wg.Wait()
	i.d.Finish()
	return nil
}

func (i Logic) DoMultiple(args ...string) error {
	wg := sync.WaitGroup{}
	wg.Add(len(args))
	for ix, x := range args {
		go func(arg string, tm int) {
			time.Sleep(time.Duration(tm) * time.Millisecond)
			i.d.Do(arg)
			wg.Done()
		}(x, ix)
	}
	wg.Wait()
	return nil
}

func (i Logic) SayHello() string {
	hh := time.Now().Hour()
	if hh < 12 {
		return "good morning"
	}
	if hh > 18 {
		return "good evening"
	}
	return "Hi"

}
