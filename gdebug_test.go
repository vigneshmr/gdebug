package gdebug

import (
	"testing"
)

type testStruct struct {
	sval string
	ival int
}

// TODO: Actually test :)
func TestDump(t *testing.T) {
	Dump(testStruct{}, "msg")
}

func TestDumpComparison(t *testing.T) {
	DumpComparison(testStruct{}, testStruct{})
}

func TestDumpComparisonIfDifferent(t *testing.T) {
	DumpComparison(testStruct{}, testStruct{})
}
