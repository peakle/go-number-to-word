package go_number_to_word

import (
	"fmt"
	"testing"
)

func TestNumberToWords(t *testing.T) {
	tc := []struct {
		N      int
		Result string
	}{
		{
			N:      18975000,
			Result: "восемнадцать миллионов девятьсот семьдесят пять тысяч",
		},
		{
			N:      3225750,
			Result: "три миллиона двести двадцать пять тысяч семьсот пятьдесят",
		},
		{
			N:      15749250,
			Result: "пятнадцать миллионов семьсот сорок девять тысяч двести пятьдесят",
		},
	}

	for i, c := range tc {
		c := c
		t.Run(fmt.Sprintf("number_to_words_%d", i), func(t *testing.T) {
			r, err := NumberToWords(c.N)
			if err != nil {
				t.Fatalf("on %d: err: %s", i, err)
			}

			if c.Result != r {
				t.Fatalf("invlid result: %s: actual: %s", c.Result, r)
			}
		})
	}
}
