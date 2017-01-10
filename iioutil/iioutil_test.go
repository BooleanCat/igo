package iioutil_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/BooleanCat/igo/iioutil"
)

var _ = Describe("iioutil", func() {
	Describe("Ioutil", func() {
		It("is implemented by IoutilWrap", func() {
			ioutil := new(iioutil.IoutilWrap)
			var _ = iioutil.Ioutil(ioutil)
		})

		It("is implemented by IoutilFake", func() {
			ioutil := new(iioutil.IoutilFake)
			var _ = iioutil.Ioutil(ioutil)
		})
	})
})
