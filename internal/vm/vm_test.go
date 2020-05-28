package vm_test

import (
	"encoding/binary"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/mhoertnagl/epic-evm/internal/vm"
)

func TestInstrctions(t *testing.T) {
	// tstFile(t, "../../test/add.tst")
	// tstFile(t, "../../test/mov.tst")
}

func tstFile(t *testing.T, path string) {
	t.Helper()

	bin, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	str := string(bin)

	var re = regexp.MustCompile(`//[^\n]*`)
	str = re.ReplaceAllString(str, ``)

	tests := strings.Split(str, "===")

	for i, test := range tests {
		test = strings.TrimSpace(test)
		if test != "" {
			tstFileTest(t, path, i, test)
		}
	}
}

func tstFileTest(t *testing.T, path string, i int, test string) {
	t.Helper()

	parts := strings.Split(test, "---")
	insParts := strings.TrimSpace(parts[0])
	insLines := strings.Split(insParts, "\n")
	tstParts := strings.TrimSpace(parts[1])
	tstLines := strings.Split(tstParts, "\n")

	code := strToCode(t, insLines)
	m := vm.NewVM(code)
	m.Run()

	for _, tstLine := range tstLines {
		parts = strings.Split(tstLine, "=")
		reg := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		a := m.Reg(reg)
		e := hexStrToUint32(t, val)
		if e != a {
			t.Errorf("\nFile: %s\nTest: %d\n\tExpected [%s] to be [%x] but is [%x].", path, i, reg, e, a)
		}
	}
}

func strToCode(t *testing.T, insLines []string) []byte {
	t.Helper()
	len := len(insLines)
	code := make([]byte, 4*len)
	for i := 0; i < len; i++ {
		v := binStrToUint32(t, insLines[i])
		binary.BigEndian.PutUint32(code[4*i:4*(i+1)], v)
	}
	return code
}

func binStrToUint32(t *testing.T, v string) uint32 {
	t.Helper()
	v = strings.TrimSpace(v)
	v = strings.Replace(v, " ", "", -1)
	n, err := strconv.ParseUint(v, 2, 32)
	if err != nil {
		t.Error(err)
	}
	return uint32(n)
}

func hexStrToUint32(t *testing.T, v string) uint32 {
	t.Helper()
	n, err := strconv.ParseUint(v, 16, 32)
	if err != nil {
		t.Error(err)
	}
	return uint32(n)
}
