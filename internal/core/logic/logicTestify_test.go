package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/mock"
)

type DoerMock struct {
	mock.Mock
}

func (i *DoerMock) DoIt(x string) (string, error) {
	args := i.Mock.Called(x)
	return args.Get(0).(string), args.Error(1)
}

func (*DoerMock) Finish() {}

func TestSum(t *testing.T) {
	x := Sum(2, 2)
	assert.Equal(t, x, 4)
	require.Equal(t, x, 4)
}

func TestOk(t *testing.T) {
	mockDoer := DoerMock{}
	defer mockDoer.AssertExpectations(t)

	l := New(&mockDoer)
	mockDoer.On("Do", "xxx").Return("yyy", nil).Times(1)
	_, err := l.Do("xxx")
	assert.NoError(t, err)
}

func TestErrType(t *testing.T) {

	mockDoer := DoerMock{}
	l := New(&mockDoer)

	t.Run("Validation error", func(t *testing.T) {
		input := ""
		mockDoer.On("Do", input).Return("", nil)
		_, err := l.Do(input)
		assert.ErrorIs(t, err, ValidationError{})
		mockDoer.AssertNumberOfCalls(t, "Do", 0)
	})

	t.Run("Error ErrEmptyString", func(t *testing.T) {
		input := "xxx"
		mockDoer.On("Do", input).Return("", nil)
		_, err := l.Do(input)
		assert.ErrorIs(t, err, ErrEmptyString)
		mockDoer.AssertNumberOfCalls(t, "Do", 1)
	})
}
