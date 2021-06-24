package handlers

import (
	"bytes"
	"encoding/json"
	"microservice_4/data"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockStore *data.MockStore

// 메서드이름을 지을 때만큼 신중하게 테스트 이름을 생각해야함
func TestSearchHandlerReturnsBadRequestWhenNoSearchCriteriaIsSent(t *testing.T) {
	r, rw, handler := setupTest(nil)

	handler.ServeHTTP(rw, r)

	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected Bad Request got %v", rw.Code)
	}
}

func TestSearchHandlerReturnsBadRequestWhenBlankSearchCriteriaIsSent(t *testing.T) {
	r, rw, handler := setupTest(&searchRequest{})

	handler.ServeHTTP(rw, r)

	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", rw.Code)
	}
}

func TestSearchHandlerCallsDataStoreWithValidQuery(t *testing.T) {
	r, rw, handler := setupTest(&searchRequest{Query: "Fat Freddy's Cat"})
	mockStore.On("Search", "Fat Freddy's Cat").Return(make([]data.Kitten, 0))

	handler.ServeHTTP(rw, r)

	mockStore.AssertExpectations(t)
}

func setupTest(d interface{}) (*http.Request, *httptest.ResponseRecorder, Search) {
	mockStore = &data.MockStore{}

	h := Search{
		DataStore: mockStore,
	}

	rw := httptest.NewRecorder()

	if d == nil {
		return httptest.NewRequest("POST", "/search", nil), rw, h
	}

	body, _ := json.Marshal(d)
	return httptest.NewRequest("POST", "/search", bytes.NewReader(body)), rw, h

}
