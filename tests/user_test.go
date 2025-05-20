package tests

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreatUser(t *testing.T) {
	// Set gin to test mode
	gin.SetMode(gin.TestMode)

	// Recreate a clean database for testing
	DatabaseRefresh()
}
