package handlers_test

import (
	"main/cmd/api/handlers"

	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mercadolibre/fury_go-platform/pkg/fury"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockCalculator struct {
	mock.Mock
}

func (m *MockCalculator) GetResult(operator string, operands []float64) (float64, error) {
	args := m.Called(operator, operands)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockCalculator) GetCalculatorMemory(name string) float64 {
	args := m.Called(name)
	return args.Get(0).(float64)
}

func (m *MockCalculator) ModifyMemory(name string, add bool, value float64) string {
	args := m.Called(name, add, value)
	return args.Get(0).(string)
}

func TestCalculate(t *testing.T) {
	//Arrange
	var mockCalculator MockCalculator
	mockCalculator.On("GetResult", mock.Anything, mock.Anything).Return(float64(6), nil)

	app, err := fury.NewWebApplication()
	require.NoError(t, err)
	handlers.Api(app, &mockCalculator)
	//Act
	bodyReader := strings.NewReader(`{"operator": "add", "operands": [1, 2, 3]}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/do", bodyReader)
	app.Router.ServeHTTP(w, r)
	res := handlers.ResponseCalculate{}
	json.Unmarshal(w.Body.Bytes(), &res)
	//Assert
	assert.EqualValues(t, 200, w.Code)
	assert.Nil(t, err)
	assert.EqualValues(t, 6, res.Result)
}

func TestFailCalculate(t *testing.T) {
	//Arrange
	var mockCalculator MockCalculator
	errorString := "test error"
	errMock  := errors.New(errorString)
	mockCalculator.On("GetResult", mock.Anything, mock.Anything).Return(float64(0), errMock)

	app, err := fury.NewWebApplication()
	require.NoError(t, err)
	handlers.Api(app, &mockCalculator)
	//Act
	bodyReader := strings.NewReader(`{"operator": "adda", "operands": [1, 2, 3]}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/do", bodyReader)
	app.Router.ServeHTTP(w, r)
	data, err := ioutil.ReadAll(w.Body)
	//Assert
	assert.EqualValues(t, 500, w.Code)
	assert.Nil(t, err)
	assert.EqualValues(t, errorString, data)
}

func TestCalculateMemory(t *testing.T) {
	//Arrange
	var mockCalculator MockCalculator
	mockCalculator.On("ModifyMemory", mock.Anything, mock.Anything, mock.Anything).Return("Memory updated")

	app, err := fury.NewWebApplication()
	require.NoError(t, err)
	handlers.Api(app, &mockCalculator)
	//Act
	bodyReader := strings.NewReader(`{"add": true, "value": 1}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/memory/blanco", bodyReader)
	app.Router.ServeHTTP(w, r)
	//Assert
	assert.EqualValues(t, 200, w.Code)
	assert.Nil(t, err)
}

func TestGetCalculateMemory(t *testing.T) {
	//Arrange
	var mockCalculator MockCalculator
	mockCalculator.On("GetCalculatorMemory", mock.Anything).Return(float64(10))

	app, err := fury.NewWebApplication()
	require.NoError(t, err)
	handlers.Api(app, &mockCalculator)
	//Act
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/memory/blanco",nil)
	app.Router.ServeHTTP(w, r)
	//Assert
	assert.EqualValues(t, 200, w.Code)
	assert.Nil(t, err)
}