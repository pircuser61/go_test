package logic

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pircuser61/go_test/internal/core/mocks"
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
