package handler

import (
	"bytes"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Hamedblue1381/fibonacci_handler/models"
	"github.com/vmihailenco/msgpack/v5"
)

func TestNextNumberHandler(t *testing.T) {
	reqBody := models.FibonacciRequest{Number: big.NewInt(5)}
	reqBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/next", bytes.NewReader(reqBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	w := httptest.NewRecorder()
	NextNumberHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d; Got %d", http.StatusOK, resp.StatusCode)
	}
	expected := "application/json"

	if resp.Header.Get("Content-Type") != expected {
		t.Errorf("Expected Content-Type: %s; Got %s", expected, resp.Header.Get("Content-Type"))
	}

	var respBody models.FibonacciResponse
	json.NewDecoder(resp.Body).Decode(&respBody)
	if respBody.Result.Cmp(big.NewInt(8)) != 0 {
		t.Errorf("Expected result 8; Got %d", respBody.Result)
	}
}

func TestPrevNumberHandler(t *testing.T) {
	reqBody := models.FibonacciRequest{Number: big.NewInt(5)}
	reqBytes, _ := msgpack.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/prev", bytes.NewReader(reqBytes))
	req.Header.Set("Content-Type", "application/msgpack")
	req.Header.Set("Accept", "application/msgpack")

	w := httptest.NewRecorder()
	PrevNumberHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d; Got %d", http.StatusOK, resp.StatusCode)
	}
	expected := "application/msgpack"

	if resp.Header.Get("Content-Type") != expected {
		t.Errorf("Expected Content-Type: %s; Got %s", expected, resp.Header.Get("Content-Type"))
	}

	var respBody models.FibonacciResponse
	msgpack.NewDecoder(resp.Body).Decode(&respBody)
	if respBody.Result.Cmp(big.NewInt(3)) != 0 {
		t.Errorf("Expected result 3; Got %d", respBody.Result)
	}
}
