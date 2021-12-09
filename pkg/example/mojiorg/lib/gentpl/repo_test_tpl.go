package gentpl

import (
	"github.com/mojiorg/gentpl/pkg/example/mojiorg/lib/gentpl/util"
	"strings"
)

type RepoTestTpl struct {
	BaseDir  string
	Data     interface{}
	FileName string
}

func (p *RepoTestTpl) GetTemplate() string {
	return `package impl_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("{{ .content.Name }}", func() {

	var ro *impl.{{ .content.Name }}RepoImpl

	BeforeEach(func() {
		ro = &impl.{{ .content.Name }}RepoImpl{
			Ds: Ds,
		}

		f := func() {
		    testData := testdata.{{ .content.Name }}()
			req := model.{{ .content.Name }}{}.New(testData)
			err := ro.Ds.Gdb().Create(req).Error
			Ω(err).ShouldNot(HaveOccurred())
		}

		f()
	})

	It("Create", func() {
        err := ro.Create(ctx, testdata.{{ .content.Name }}(2))
        Ω(err).ShouldNot(HaveOccurred())
    })

    It("Update", func() {
        err := ro.Update(ctx, testdata.{{ .content.Name }}(1))
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

})`
}

func (p *RepoTestTpl) GetPath() string {
	return p.GetBaseDir() + `../repo/impl`
}

func (p *RepoTestTpl) GetFileName() string {
	return strings.ReplaceAll(p.FileName, ".go", "_test.go")
}

func (p *RepoTestTpl) GetData() interface{} {
	return p.Data
}

func (p *RepoTestTpl) ParseData(file string) {
	f := util.FilePath(file)
	p.Data = util.ParseData(file)
	p.FileName = f.FileName()
	if p.BaseDir == "" {
		p.BaseDir = f.FolderName()
	}
	return
}

func (p *RepoTestTpl) SetBaseDir(dir string) {
	p.BaseDir = dir
}

func (p *RepoTestTpl) GetBaseDir() string {
	return p.BaseDir
}
