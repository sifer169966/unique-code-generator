package generator

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

var rememberer = map[string]bool{}

type Generator struct {
	mixer    *mixer
	ruler    *ruler
	Quantity int
}

func NewGenerator() *Generator {
	mixer := defaultMixer()
	ruler := defaultRuler()
	return &Generator{
		mixer: mixer,
		ruler: ruler,
	}
}

func (gen *Generator) GenerateCodes(amount int) (code []string) {

	for i := 0; i < amount; i++ {
		result := gen.generateCode()
		code = append(code, result)
		fmt.Println(strings.Join(code, "\n"))
	}
	return code
}

func (gen *Generator) generateCode() string {
	codeBuilder := make([]byte, gen.ruler.length)
	var mixUpperCaseEN string
	mixUpperCaseEN = gen.mixer.upperCaseEN
	for i := range codeBuilder {
		target := rand.Intn(len(mixUpperCaseEN))
		codeBuilder[i] = mixUpperCaseEN[target]
		if gen.ruler.isUniqueCharacter {
			mixUpperCaseEN = gen.removeUsedCharacter(mixUpperCaseEN, target)
		}
	}

	code := string(codeBuilder)
	if gen.ruler.isUniqueCode && !gen.isCodeAlredayExist(code) {
		gen.remember(code)
	} else {
		panic("found duplicate code " + code)
	}
	return gen.ruler.prefix + code
}

func (gen *Generator) removeUsedCharacter(mixer string, target int) string {
	var buffer bytes.Buffer
	for i := range mixer {
		if i == target {
			continue
		}
		buffer.WriteByte(mixer[i])
	}
	return buffer.String()
}

func (gen *Generator) isCodeAlredayExist(code string) bool {
	_, exist := rememberer[code]
	return exist
}

func (gen *Generator) remember(code string) {
	rememberer[code] = true
}
