package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var checkPath string

var _ = BeforeSuite(func() {

})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func TestCheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Check Suite")
}
