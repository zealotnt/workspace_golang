package size

import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{0, "zero"},
	{5, "small"},
	{15, "big"},
	{105, "huge"},
	{1005, "enormous"},
}

func TestSize(t *testing.T) {
	for i, test := range tests {
		size := Size(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}

func TestSomeThing(t *testing.T) {
	SomeThing()
}
