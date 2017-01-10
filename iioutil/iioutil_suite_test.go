package iioutil_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIioutil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "iioutil suite")
}
