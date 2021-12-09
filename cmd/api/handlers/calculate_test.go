package handlers_test

import (
	"main/cmd/api/handlers"

	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"encoding/json"

	"github.com/mercadolibre/fury_go-platform/pkg/fury"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var MemoryMap map[string]float64
var errorString string = "test error"

func mockCalculator(operator string, operands []float64) (float64, error) {
	return 6, nil
}

func mockFailCalculator(operator string, operands []float64) (float64, error) {
	var errMock error = errors.New("test error")
	return 0, errMock
}

func mockCalculatorMemory(memoryMap map[string]float64, name string, add bool, value float64) string{
	result := "Memory updated"
	return result
}

func mockGetCalculatorMemory(name string, memoryMap map[string]float64) float64{
	result := float64(10)
	return result
}

func TestCalculate(t *testing.T) {
	//Arrange
	app, err := fury.NewWebApplication()
	require.NoError(t, err)
	MemoryMap = make(map[string]float64)
	handlers.Api(app, MemoryMap)
	handlers.FuncCalculator = mockCalculator
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
	app, err := fury.NewWebApplication()
	require.NoError(t, err)
	MemoryMap = make(map[string]float64)
	handlers.Api(app, MemoryMap)
	handlers.FuncCalculator = mockFailCalculator
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
	app, err := fury.NewWebApplication()
	require.NoError(t, err)
	MemoryMap = make(map[string]float64)
	handlers.Api(app, MemoryMap)
	handlers.FuncCalculatorMemory = mockCalculatorMemory
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
	app, err := fury.NewWebApplication()
	require.NoError(t, err)
	MemoryMap = make(map[string]float64)
	handlers.Api(app, MemoryMap)
	handlers.FuncGetCalculatorMemory = mockGetCalculatorMemory
	//Act
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/memory/blanco",nil)
	app.Router.ServeHTTP(w, r)
	//Assert
	assert.EqualValues(t, 200, w.Code)
	assert.Nil(t, err)
}