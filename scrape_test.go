package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetHoshitori(t *testing.T) {
	RegisterTestingT(t)

	ret,err := GetHoshitori()
	Expect(err).To(BeNil())
	Expect(len(ret)).To(Equal(21))
}
