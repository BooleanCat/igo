package inet_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "inet suite")
}
