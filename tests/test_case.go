package tests

import (
	"github.com/goravel/framework/testing"

	"myblog/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
