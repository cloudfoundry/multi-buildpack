package integration_test

import (
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("running supply buildpacks before the binary buildpack", func() {
	var app *cutlass.App
	AfterEach(func() {
		if app != nil {
			app.Destroy()
		}
		app = nil
	})

	Context("the app is pushed once", func() {
		BeforeEach(func() {
			app = cutlass.New(filepath.Join(bpDir, "fixtures", "fake_supply_binary_app"))
			app.Buildpack = "multi_buildpack"
		})

		It("finds the supplied dependency in the runtime container", func() {
			PushAppAndConfirm(app)

			Expect(app.Stdout.String()).ToNot(ContainSubstring("SUPPLYING DOTNET"))

			Expect(app.GetBody("/")).To(MatchRegexp("dotnet: 1.0.1"))
		})
	})
})