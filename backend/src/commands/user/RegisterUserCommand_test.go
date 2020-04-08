package user

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

)

// Simple Unittest
func TestRegisterUser_BadRequestCase1(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	RegisterUserCommand(c)
	//assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}