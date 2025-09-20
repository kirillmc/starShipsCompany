package v1

import (
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := logger.Init("", true)
	if err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}
