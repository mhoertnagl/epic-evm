package cpu_test

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/mhoertnagl/epic-evm/internal/cpu"
)

func TestInstrctions2(t *testing.T) {
	tstFiles(t, "../../test")
}

const MSG_EXP string = "\n[%s] Test [%d:%d]: Expected [%s] to be [%x] but is [%x]."
const MSG_FMT_TST string = "\n[%s] Test [%d:%d]: Illegal format."
const MSG_FMT_REG string = "\n[%s] Test [%d:%d]: Illegal format in test expression [%s]."

func tstFiles(t *testing.T, dir string) {
	t.Helper()

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".tst" {
			fmt.Printf("Testing [%s] ...", path)
			tstFile(t, path)
			fmt.Printf(" OK.\n")
		}
		return nil
	})
}

func tstFile(t *testing.T, path string) {
	t.Helper()

	bin, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	str := string(bin)

	// Remove all comments.
	var re = regexp.MustCompile(`//[^\n]*`)
	str = re.ReplaceAllString(str, ``)

	// Split tests into blocks.
	tests := strings.Split(str, "===")

	for i, test := range tests {
		// Remove trailing spaces.
		test = strings.TrimSpace(test)

		if test != "" {
			tstFileTest(t, path, i, test)
		}
	}
}

func tstFileTest(t *testing.T, path string, i int, test string) {
	t.Helper()

	lines := strings.Split(test, "\n")

	asm := []string{}
	bin := []string{}
	tst := []string{}

	for n, line := range lines {
		parts := strings.Split(line, "|")

		// Check the format of a single instruction.
		// The format is asm | bin | reg where asm is an essembly instruction,
		// bin is the binary code for this instruction and reg is a single register
		// test in the format reg = val.
		if len(parts) != 3 {
			t.Errorf(MSG_FMT_TST, path, i, n)
		}

		asm = append(asm, strings.TrimSpace(parts[0]))
		bin = append(bin, strings.TrimSpace(parts[1]))
		tst = append(tst, strings.TrimSpace(parts[2]))
	}

	code := strToCode(t, bin)
	m := cpu.NewCpu(code)

	for m.Running() {
		ip := m.Reg("ip")
		m.Step()
		tstRegTest(t, path, i, m, tst[ip], ip)
	}
}

func tstRegTest(t *testing.T, path string, i int, m *cpu.Cpu, tst string, ip uint32) {
	if tst != "" {
		parts := strings.Split(tst, "=")

		// A test expression has the form reg = val.
		if len(parts) != 2 {
			t.Errorf(MSG_FMT_REG, path, i, ip+1, tst)
		}

		reg := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		a := m.Reg(reg)
		e := hexStrToUint32(t, val)

		if e != a {
			t.Errorf(MSG_EXP, path, i, ip+1, reg, e, a)
		}
	}
}

func strToCode(t *testing.T, bin []string) []byte {
	t.Helper()

	len := len(bin)
	code := make([]byte, 4*len)

	for i := 0; i < len; i++ {
		v := binStrToUint32(t, bin[i])
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
