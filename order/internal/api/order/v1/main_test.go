package v1

import (
	"os"
	"testing"

	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
)

func TestMain(m *testing.M) {
	err := logger.Init("", true)
	if err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}
