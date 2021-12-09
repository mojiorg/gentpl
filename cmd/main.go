package main

import (
	"github.com/mojiorg/gentpl/pkg"
	"github.com/mojiorg/gentpl/pkg/example/mojiorg/lib/gentpl"
)

func main() {
	file := "pkg/example/mojiorg/domain/aggre.go"
	tplService := pkg.GenTpl{}
	l := []pkg.ITemplate{
		&gentpl.TestDataTpl{},
		&gentpl.RepoTpl{},
		&gentpl.RepoImplTpl{},
		&gentpl.RepoModelTpl{},
		&gentpl.RepoTestTpl{},
	}
	tplService.Gen(file, l)
}
