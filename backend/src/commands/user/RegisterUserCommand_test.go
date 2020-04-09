package user

import (
	"github.com/gin-gonic/gin"
	"net/url"

	"strconv"
	"strings"

	"net/http"
	"net/http/httptest"


	"testing"


)

//func newTestContext(method, path string) (w *httptest.ResponseRecorder, r *http.Request) {
//	w = httptest.NewRecorder()
//	r, _ = http.NewRequest(method, path, nil)
//	r.PostForm = url.Values{}
//	return
//}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func getRegistrationPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "sangnv2222@ithink.vn")
	return params.Encode()
}

// Simple Unittest
func Test_RegisterUser_BadRequestCase1(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.POST("/api/user/register-user", RegisterUserCommand)

	input := getRegistrationPOSTPayload()

	req, err := http.NewRequest(http.MethodPost, "/api/user/register-user", strings.NewReader(input))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(input)))

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
