package test

import (
	"testing"

	"github.com/betorvs/bigestresponsetimeicmp/usecase"
)

func TestRun(t *testing.T) {

	usecase.CalculateBigestResponseTime("www.google.com")
	usecase.CalculateBigestResponseTime("172.217.168.228")
	usecase.CalculateBigestResponseTime("isbrobous.net")
}
