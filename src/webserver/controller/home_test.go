package controller

import (
	"testing"
	"html/template"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
)

func TestLoginExecutesCorrectTemplate(t *testing.T) {
	h := new(home)
	expected := `login template`
	h.loginTemplate, _ = template.New("").Parse(expected)
	r := httptest.NewRequest(http.MethodGet, "/login", nil)
	w := httptest.NewRecorder()

	h.handleLogin(w, r)

	actual, _ := ioutil.ReadAll(w.Result().Body)

	if string(actual) != expected {
		t.Errorf("Failed to serve the correct template")
	}
}