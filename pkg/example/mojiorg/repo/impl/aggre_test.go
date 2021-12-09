package impl_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Aggre", func() {

	var ro *impl.AggreRepoImpl

	BeforeEach(func() {
		ro = &impl.AggreRepoImpl{
			Ds: Ds,
		}

		f := func() {
			testData := testdata.Aggre()
			req := model.Aggre{}.New(testData)
			err := ro.Ds.Gdb().Create(req).Error
			Ω(err).ShouldNot(HaveOccurred())
		}

		f()
	})

	It("Create", func() {
		err := ro.Create(ctx, testdata.Aggre(2))
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("Update", func() {
		err := ro.Update(ctx, testdata.Aggre(1))
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("MustGet", func() {
		_, err := ro.MustGet(ctx, 1)
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("MultiGet", func() {
		l, err := ro.MultiGet(ctx, 1, 2, 3)
		Ω(err).NotTo(HaveOccurred())
		Ω(l).To(HaveLen(1))
	})

})
