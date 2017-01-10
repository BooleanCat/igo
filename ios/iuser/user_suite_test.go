package iuser_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIuser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "iuser suite")
}
