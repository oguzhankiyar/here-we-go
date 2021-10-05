package main

import (
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	Describe("the strings package", func() {
		Context("strings.Contains()", func() {
			When("the string contains the substring in the middle", func() {
				It("returns `true`", func() {
					Expect(strings.Contains("Ginkgo is awesome", "is")).To(BeTrue())
				})
			})
		})
	})

	RunSpecs(t, "samples")
}