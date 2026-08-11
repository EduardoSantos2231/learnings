package main

import (
	"bytes"
	"diagnostico/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddItem(t *testing.T) {
	defer FakeDb.ResetDB()
	toSend := struct {
		Name string `json:"name"`
	}{
		Name: "arroz",
	}
	toSendMarshalled, _ := json.Marshal(toSend)
	body := bytes.NewBuffer(toSendMarshalled)
	req := httptest.NewRequest("POST", "/api/v1/items", body)
	recorder := httptest.NewRecorder()
	handleItemAddition(recorder, req)
	if recorder.Result().StatusCode != http.StatusCreated {
		t.Errorf("Expected: 201\nGot: %d", recorder.Result().StatusCode)
	}
}
func TestAddItem_InvalidBody(t *testing.T) {
	toSend := struct {
		Name string `json:"name"`
	}{
		Name: "",
	}
	bytesJson, _ := json.Marshal(toSend)
	body := bytes.NewBuffer(bytesJson)
	req := httptest.NewRequest("POST", "/api/v1/items", body)
	recorder := httptest.NewRecorder()
	handleItemAddition(recorder, req)
	if recorder.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("Expected: 400\nReceived: %d", recorder.Result().StatusCode)
	}
}

func TestListItems(t *testing.T) {
	FakeDb.Seed()
	defer FakeDb.ResetDB()
	response := struct {
		Items []storage.Item `json:"items"`
	}{}
	req := httptest.NewRequest(http.MethodGet, "/api/v1/items", nil)
	recorder := httptest.NewRecorder()
	handleItemListing(recorder, req)
	json.NewDecoder(recorder.Result().Body).Decode(&response)
	if len(response.Items) != 10 {
		t.Errorf("Expected lenght: 10\nGot: %d", len(response.Items))

	}
	if response.Items[0].Name != "Test 0" {
		t.Errorf("Expected name: Test 0\nGot %s", response.Items[0].Name)
	}
}

func TestMarkItemAsDone(t *testing.T) {
	FakeDb.Seed()
	defer FakeDb.ResetDB()
	toSend := struct {
		Ids []int `json:"ids"`
	}{
		Ids: []int{1, 2, 3, 4},
	}
	toSendMarshalled, _ := json.Marshal(toSend)
	body := bytes.NewBuffer(toSendMarshalled)
	req := httptest.NewRequest("PUT", "/api/v1/items", body)
	recorder := httptest.NewRecorder()
	handleSetItemDone(recorder, req)
	if recorder.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected: 200\nGot: %d", recorder.Result().StatusCode)
	}
}

func TestMarkItemAsDone_InvalidIds(t *testing.T) {
	FakeDb.Seed()
	defer FakeDb.ResetDB()
	toSend := struct {
		Ids []int `json:"ids"`
	}{
		Ids: []int{77, 92, 44},
	}
	toSendMarshalled, _ := json.Marshal(toSend)
	body := bytes.NewBuffer(toSendMarshalled)
	req := httptest.NewRequest("PUT", "/api/v1/items", body)
	recorder := httptest.NewRecorder()
	handleSetItemDone(recorder, req)
	if recorder.Result().StatusCode == http.StatusOK {
		t.Errorf("Expected: %d\nGot: %d", http.StatusBadRequest, http.StatusOK)
	}
}
