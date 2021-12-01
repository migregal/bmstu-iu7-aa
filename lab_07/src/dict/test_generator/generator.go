package main

import (
	"os"
	"strings"
	"text/template"
)

const test_cases = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const global_tpl = `package {{ .PackageName }}

import (
	"sort"
	"testing"
)

var darr = CreateArray({{ .ArrSize }})
var farr = darr.FreqAnalysis()

`

const brute_tpl = `func BenchmarkBrute{{ .TestCase | ToUpper }}(b *testing.B) {
	w := darr.Pick("{{ .TestCase }}")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

`

const binary_tpl = `func BenchmarkBinary{{ .TestCase | ToUpper }}(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["{{ .KeyName }}"].(string) < darrc[j]["{{ .KeyName }}"].(string)
	})
	w := darrc.Pick("{{ .TestCase }}")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

`

const combined_tpl = `func BenchmarkCombined{{ .TestCase | ToUpper }}(b *testing.B) {
	w := darr.Pick("{{ .TestCase }}")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

`

func generateTmpl(tmpl string, funcs, args map[string]interface{}) {
	t := template.Must(template.New("").Funcs(funcs).Parse(tmpl))
	t.Execute(os.Stdout, args)
}

// Usage: test_generator <array size>
func main() {
	if len(os.Args) < 2 {
		panic("too few args")
	}

	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
	}

	generateTmpl(global_tpl, funcMap, map[string]interface{}{
		"PackageName": "dict",
		"ArrSize":     os.Args[1],
	})

	for _, c := range test_cases {
		generateTmpl(brute_tpl, funcMap, map[string]interface{}{
			"TestCase": string(c),
		})
	}

	for _, c := range test_cases {
		generateTmpl(binary_tpl, funcMap, map[string]interface{}{
			"TestCase": string(c),
			"KeyName":  "username",
		})
	}

	for _, c := range test_cases {
		generateTmpl(combined_tpl, funcMap, map[string]interface{}{
			"TestCase": string(c),
		})
	}
}
