package util

import (
	"regexp"
	"strings"
)

type DomainParser struct {
}

type DomainTp string

type DomainTpList []DomainTp

func (p *DomainTp) GetSeqTp() DataSeqTp {
	isIn := func(x DomainTp, vs DomainTpList) bool {
		for _, item := range vs {
			if x == item {
				return true
			}
		}
		return false
	}
	v := *p
	if isIn(v, p.intTps()) {
		return DataSeqTpInt
	} else if isIn(v, p.floatTps()) {
		return DataSeqTpFloat
	} else if isIn(v, p.stringTps()) {
		return DataSeqTpString
	} else if isIn(v, p.booleanTps()) {
		return DataSeqTpBool
	}
	return DataSeqTpInt // 不识别的类型都当做Int枚举
}

func (p *DomainTp) intTps() DomainTpList {
	return []DomainTp{
		"uint8", "uint16", "uint32", "uint64",
		"int8", "int16", "int32", "int64",
		"byte", "rune", "uint", "int", "uintptr",
		"utee.Tick"}
}

func (p *DomainTp) floatTps() DomainTpList {
	return []DomainTp{"float32", "float64", "complex64", "complex128"}
}

func (p *DomainTp) stringTps() DomainTpList {
	return []DomainTp{"string"}
}

func (p *DomainTp) booleanTps() DomainTpList {
	return []DomainTp{"bool"}
}

type DomainField struct {
	Name    string
	Type    DomainTp
	MockVal string
	GormTag string
}

type DomainStruct struct {
	Name     string
	FileName string
	Fields   []*DomainField
}

func (p *DomainParser) Parse(content string) *DomainStruct {
	pr := Parser{
		PrepareReg: []string{".*type.*struct.*\\{(.*\\n)+?.*\\}"},
	}
	structContents := pr.Exe(content)
	for _, structContent := range structContents {
		fields := p.parseFields(structContent)
		v := &DomainStruct{
			Name:   p.parseName(structContent),
			Fields: fields,
		}
		return v
	}
	return nil
}

func (p *DomainParser) parseName(structContent string) string {
	pr := Parser{
		PrepareReg: []string{"type.*struct"},
		ProcessReg: []string{"type", "struct", " "},
	}
	vs := pr.Exe(structContent)
	if len(vs) <= 0 {
		return ""
	}
	return vs[0]
}

func (p *DomainParser) parseFields(structContent string) []*DomainField {
	var vs []*DomainField
	pr := Parser{
		PrepareReg: []string{`.*\n`},
		ProcessReg: []string{"type.*struct.*", "//.*"},
	}
	contents := pr.Exe(structContent)
	seq := NewDataSeq()
	for _, content := range contents {
		var subs []string
		for _, item := range strings.Split(content, " ") {
			if item != "" {
				subs = append(subs, item)
			}
		}
		field := &DomainField{
			Name: subs[0],
			Type: DomainTp(subs[1]),
		}
		field.MockVal = seq.Next(field.Type.GetSeqTp())
		if field.Name == "Id" {
			field.MockVal = "0"
		}
		vs = append(vs, field)
	}
	return vs
}

type Parser struct {
	PrepareReg []string
	ProcessReg []string
}

func (p *Parser) Exe(content string) []string {
	prep := func(reg string, contents ...string) []string {
		var result []string
		for _, content := range contents {
			rs := regexp.MustCompile("(?i)"+reg).FindAllString(content, -1)
			result = append(result, rs...)
		}
		return result
	}

	proc := func(reg string, contents ...string) []string {
		var result []string
		for _, content := range contents {
			rs := regexp.MustCompile("(?i)"+reg).ReplaceAllString(content, "")
			result = append(result, rs)
		}
		return result
	}

	trimAndClear := func(strs ...string) []string {
		result := []string{}
		for _, v := range strs {
			v = strings.TrimSpace(v)
			if v != "" {
				result = append(result, v)
			}
		}
		return result
	}

	result := []string{content}
	for _, reg := range p.PrepareReg {
		result = prep(reg, result...)
	}
	for _, reg := range p.ProcessReg {
		result = proc(reg, result...)
	}
	return trimAndClear(result...)
}

func ParseData(file string) interface{} {
	content, err := (&FileUt{}).ReadAll(file)
	Chk(err)
	data := (&DomainParser{}).Parse(content)
	res := map[string]interface{}{
		"content":  data,
		"fileName": FilePath(file).FileName(),
	}
	return res
}
