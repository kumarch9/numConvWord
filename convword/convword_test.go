package convword_test

import (
	"strings"
	"testing"

	"github.com/pkchandra/numconvword/convword"
)

func TestConvNumToWords(t *testing.T) {
	dumpOutput := []string{
		"Ninety Five Crore Forty Eight Lacs Seventy Five Thousands Six Hundreds Twenty One",
		"Forty Eight Lacs Seventy Five Thousands Six Hundreds",
	}
	num := []int{954875621, 4875600}
	for i, val := range num {
		word, _ := convword.ConvNumToWords(val)
		want := strings.Compare(word, dumpOutput[i])
		if want != 0 {
			t.Errorf("want 0 but got %d", want)
		}

	}
}
