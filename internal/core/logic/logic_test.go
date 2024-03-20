package logic

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pircuser61/go_test/internal/core/mocks"
	"github.com/tkuchiki/faketime"
)

func TestLogicOk(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	input := "test"
	want := "ok"
	mockDoer := mocks.NewMockDoer(mockCtrl)
	l := New(mockDoer)
	mockDoer.EXPECT().Do(input).Return("result", nil).Times(1)

	got, err := l.Do(input)
	if err != nil {
		t.Error("unexpected error", err.Error())
	}
	if want != got {
		t.Errorf("want %s got %s", want, got)
	}
}

func TestLogicErr(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	input := "test"
	want := ""
	mockDoer := mocks.NewMockDoer(mockCtrl)
	l := New(mockDoer)
	mockDoer.EXPECT().Do(input).Return("", errors.New("testError")).Times(1)

	got, err := l.Do(input)
	if err == nil {
		t.Error("must be error")
		return
	}
	if want != got {
		t.Errorf("want %s got %s", want, got)
	}
}

func TestLogicConcreteErr(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	input := "test"
	wantErr := ErrEmptyString
	mockDoer := mocks.NewMockDoer(mockCtrl)
	l := New(mockDoer)
	mockDoer.EXPECT().Do(input).Return("", nil).Times(1)

	_, err := l.Do(input)
	if err != wantErr {
		t.Error("must be error 'EmptyString', got err:", err)
		return
	}
}

func TestLogicErrType(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	l := New(mockDoer)
	t.Run("Validation error Empty", func(t *testing.T) {
		mockDoer.EXPECT().Do("").Return("", nil).Times(0)
		_, err := l.Do("")
		if err == nil {
			t.Error("must be error")
			return
		}
		if !errors.Is(err, ValidationError{}) {
			t.Error("must be error type 'ValidationError', got:", err)
			return
		}
	})

	t.Run("Validation error Short", func(t *testing.T) {
		mockDoer.EXPECT().Do("").Return("", nil).Times(0)
		_, err := l.Do("1")
		if err == nil {
			t.Error("must be error")
			return
		}

		//тест падает т.к. тир ошибки не ValidationError
		if !errors.Is(err, &ValidationError{}) {
			t.Error("must be error type 'ValidationError', got:", err)
			return
		}
	})
}

func TestLogicParallel(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	doer := mocks.NewMockDoer(mockController)

	doer.EXPECT().Finish().After(doer.EXPECT().Do("test").Times(4))
	l := New(doer)
	l.DoParallel("test", 4)
}

func TestLogicOrder(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	doer := mocks.NewMockDoer(mockController)
	gomock.InOrder(
		doer.EXPECT().Do("first"),
		doer.EXPECT().Do("second"),
		doer.EXPECT().Do("third"))

	l := New(doer)
	err := l.DoMultiple("first", "second", "third")
	if err != nil {
		t.Error("unexpected error", err)
	}
}

func TestLogicTime(t *testing.T) {
	tm := time.Date(2009, time.November, 10, 11, 0, 0, 0, time.UTC)
	f := faketime.NewFaketimeWithTime(tm)
	defer f.Undo()
	f.Do()
	l := New(nil)

	want := "good morning"
	got := l.SayHello()
	if want != got {
		t.Errorf("want %s got %s", want, got)
	}
}
