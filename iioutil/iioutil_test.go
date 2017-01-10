package iioutil_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/BooleanCat/igo/iioutil"
)

var _ = Describe("iioutil", func() {
	It("is implemented by Real", func() {
		var _ iioutil.Ioutil = iioutil.New()
	})

	It("is implemented by Fake", func() {
		var _ iioutil.Ioutil = iioutil.NewFake()
	})
})
