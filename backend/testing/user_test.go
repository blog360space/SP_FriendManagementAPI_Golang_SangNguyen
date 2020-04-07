package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	app "spapp/src"
	routers "spapp/src/routers"
)

func RegisterUser_Ok(t *testing.T) {
	app.Bootstrap()
	router := routers.GetRoutes()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	//assert.Equal(t, "pong", w.Body.String())
}

//func RegisterUser_BadRequestCase1(t *testing.T) {
//	total := Sum(5, 5)
//	if total != 10 {
//		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
//	}
//}