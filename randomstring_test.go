package randomstring

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	println("Generate:", Generate(20))
}

func TestGenerateByCharset(t *testing.T) {
	println("GenerateByCharset:", GenerateByCharset(20, Ascii))
}
