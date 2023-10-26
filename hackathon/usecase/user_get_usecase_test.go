package usecase_test

import (
	"github.com/go-playground/assert/v2"
	"hackathon/controller"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	router := controller.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/00000000000000000000000001", nil)
	router.ServeHTTP(w, req)
	log.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)
	expected := `[{"user_id":"00000000000000000000000001","user_name":"hanako","email":"hana@gmail.com","term":0}]`
	assert.Equal(t, expected, w.Body.String())
}
