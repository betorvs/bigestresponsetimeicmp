package test

import (
	"testing"

	"github.com/betorvs/biggestresponsetimeicmp/usecase"
)

func TestRun(t *testing.T) {

	fakemap := make(map[[4]byte]int64)
	time1 := int64(5057475)
	addr1 := [4]byte{10, 10, 27, 1}
	time2 := int64(459153154)
	addr2 := [4]byte{10, 10, 20, 1}

	fakemap[addr1] = time1
	fakemap[addr2] = time2

	key, big := usecase.FindBiggestResponseTime(fakemap)
	if big == time1 {
		t.Fatalf("Error occurred while trying to calculate the biggest response time")
	}
	if key != addr2 {
		t.Fatalf("Error occurred while trying to calculate the biggest response time")
	}
	// usecase.CalculateBiggestResponseTime("www.google.com")
	// usecase.CalculateBiggestResponseTime("172.217.168.228")
	// usecase.CalculateBiggestResponseTime("isbrobous.net")
}
