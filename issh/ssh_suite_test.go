package issh_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIssh(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "issh suite")
}
