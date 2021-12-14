package calculator_test

import (
	"main/internal/calculator"
	"main/internal/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenCallGetResultWithMPlusOperandThenSuccess(t *testing.T) {
	calculator := calculator.New()
	result, error := calculator.GetResult("m+", nil, storage.NewFile())

	assert.Nil(t, error)
	assert.EqualValues(t, 0, result)
}

func TestWhenCallGetResultWithMMinusOperandThenSuccess(t *testing.T) {
	calculator := calculator.New()
	result, error := calculator.GetResult("m-", nil, storage.NewFile())

	assert.Nil(t, error)
	assert.EqualValues(t, 0, result)
}

func TestWhenCallGetResultWithMPlusOperandThenShouldGetIncrementedMemory(t *testing.T) {
	calculator := calculator.New()
	calculator.GetResult("m+", nil, storage.NewFile())
	result, error := calculator.GetResult("get", nil, storage.NewFile())

	assert.Nil(t, error)
	assert.EqualValues(t, 1, result)
}

func TestWhenCallGetResultWithMMinusOperandThenShouldGetDecrementedMemory(t *testing.T) {
	calculator := calculator.New()
	calculator.GetResult("m-", []float64{}, storage.NewFile())
	result, error := calculator.GetResult("get", []float64{}, storage.NewFile())

	assert.Nil(t, error)
	assert.EqualValues(t, -1, result)
}

func TestWhenCallGetResultWithResetOperandThenGetResetMemory(t *testing.T) {
	calculator := calculator.New()
	calculator.GetResult("m-", []float64{}, storage.NewFile())
	calculator.GetResult("reset", []float64{}, storage.NewFile())
	result, error := calculator.GetResult("get", []float64{}, storage.NewFile())

	assert.Nil(t, error)
	assert.EqualValues(t, 0, result)
}

func TestWhenCallGetResultWithAddOperandThenShouldGetSuccessResultOfOperands(t *testing.T) {
	calculator := calculator.New()
	result, error := calculator.GetResult("add", []float64{1, 2, 3}, storage.NewFile())

	assert.Nil(t, error)
	assert.EqualValues(t, 6, result)
}

func TestWhenCallGetResultWithSubstractOperandThenShouldGetSuccessResultOfOperands(t *testing.T) {
	calculator := calculator.New()
	result, error := calculator.GetResult("substract", []float64{4, 1, 1}, storage.NewFile())

	assert.Nil(t, error)
	assert.EqualValues(t, 2, result)
}

func TestWhenCallGetResultWithMultiplyOperandThenShouldGetSuccessResultOfOperands(t *testing.T) {
	calculator := calculator.New()
	result, error := calculator.GetResult("multiply", []float64{1, 2, 3, 4}, storage.NewFile())

	assert.Nil(t, error)
	assert.EqualValues(t, 24, result)
}

func TestWhenCallGetResultWithDivideOperandThenShouldGetSuccessResultOfOperands(t *testing.T) {
	calculator := calculator.New()
	result, error := calculator.GetResult("divide", []float64{16, 2, 2}, storage.NewFile())

	assert.Nil(t, error)
	assert.EqualValues(t, 4, result)
}

func TestWhenCallGetResultWithDivideOperandWithZeroOperatorThenError(t *testing.T) {
	calculator := calculator.New()
	result, error := calculator.GetResult("divide", []float64{16, 2, 0}, storage.NewFile())

	assert.EqualValues(t, 0, result)
	assert.Error(t, error)
	assert.Equal(t, "impossible division by zero", error.Error())
}

func TestWhenCallGetResultWithInvalidOperandWithThenError(t *testing.T) {
	calculator := calculator.New()
	result, error := calculator.GetResult("test", []float64{16, 2, 0}, storage.NewFile())

	assert.EqualValues(t, 0, result)
	assert.Error(t, error)
	assert.Equal(t, "undefined operator", error.Error())
}