package gdebug

import (
	"math/rand"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randInt() int {
	return rand.Intn(5)
}

func randIntPtr() *int {
	val := randInt()
	return &val
}

func randArrIntPtr() []*int {
	val := []*int{}
	for i := 0; i <= randInt(); i++ {
		val = append(val, randIntPtr())
	}
	return val
}

func randStr() string {
	b := make([]rune, randInt())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randStrPtr() *string {
	val := randStr()
	return &val
}

func randArrStrPtr() []*string {
	val := []*string{}
	for i := 0; i <= randInt(); i++ {
		val = append(val, randStrPtr())
	}
	return val
}

type testStruct struct {
	Sval    string
	Ival    int
	Arrip   []*int
	Arrstrp []*string
}

func buildTestStruct() *testStruct {
	return &testStruct{
		Sval:    randStr(),
		Ival:    randInt(),
		Arrip:   randArrIntPtr(),
		Arrstrp: randArrStrPtr(),
	}
}

// TODO: Actually test :)
func TestDump(t *testing.T) {
	Dump(buildTestStruct(), "msg")
}

func TestDumpComparison(t *testing.T) {
	obj := buildTestStruct()
	DumpComparison(obj, obj)
}

func TestDumpComparisonIfDifferent_SameObj(t *testing.T) {
	obj := buildTestStruct()
	DumpComparisonIfDifferent(obj, obj)
}

func TestDumpComparisonIfDifferent_DiffObj(t *testing.T) {
	exp := buildTestStruct()
	act := *exp
	act.Ival++
	DumpComparisonIfDifferent(exp, act)
}
