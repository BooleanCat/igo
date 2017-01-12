package ios_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ios suite")
}
