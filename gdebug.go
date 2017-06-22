package gdebug

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
)

// Dump an object
func Dump(obj interface{}, msg string) {
	expObj, _ := json.MarshalIndent(obj, "", "    ")
	fmt.Printf("[%s]\n%s", msg, string(expObj))
}

// DumpComparison dumps the given objects for comparison
func DumpComparison(expected interface{}, actual interface{}) {
	s1 := dumpObj(expected, "Expected")
	s2 := dumpObj(actual, "Actual")

	w := tabwriter.NewWriter(os.Stdout, 3, 10, 0, ' ', 0)
	red := color.New(color.BgRed)
	fgWhiteBgRed := red.Add(color.FgHiWhite)
	diffColorSprintf := fgWhiteBgRed.SprintFunc()
	boldWhiteSprintf := color.New(color.Bold).SprintFunc()

	var ls1, ls2, eqSym string
	for i := 0; i < max(len(s1), len(s2)); i++ {
		ls1 = ""
		ls2 = ""
		if i < len(s1) {
			ls1 = s1[i]
		}
		if i < len(s2) {
			ls2 = s2[i]
		}

		isEq := true
		eqSym = ""
		if i > 0 {
			isEq = strings.EqualFold(ls1, ls2)
			if !isEq {
				eqSym = " != "
			}
		} else if i == 0 {
			ls1 = boldWhiteSprintf(ls1)
			ls2 = boldWhiteSprintf(ls2)
		}

		if isEq {
			fmt.Fprintf(w, "%s\t\t%s\n", ls1, ls2)
		} else {
			fmt.Fprintf(w, "%s\t%s\t%s\n", diffColorSprintf(ls1), eqSym, diffColorSprintf(ls2))
		}
	}
	w.Flush()
}

// DumpComparisonIfDifferent dumps the given objects for comparison if their values differ
func DumpComparisonIfDifferent(expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		DumpComparison(expected, actual)
	}
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func dumpObj(obj interface{}, msg string) []string {
	expObj, _ := json.MarshalIndent(obj, "", "    ")
	s := fmt.Sprintf("[%s]\n%s", msg, string(expObj))
	return strings.Split(s, "\n")
}
