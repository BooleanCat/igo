package iexec_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIexec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "iexec suite")
}
