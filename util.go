package go_number_to_word

import (
	"fmt"
	"strings"
)

var (
	units = []string{
		"ноль", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять",
	}
	teenagers = []string{
		"десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать",
		"семнадцать", "восемнадцать", "девятнадцать",
	}
	tens = []string{
		"двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто",
	}
	hundreds = []string{
		"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот",
	}
	thousands = []string{
		"", "тысяча", "тысячи", "тысяч",
	}
	millions = []string{
		"", "миллион", "миллиона", "миллионов",
	}
	billions = []string{
		"", "миллиард", "миллиарда", "миллиардов",
	}
)

func NumberToWords(n int) (string, error) {
	if n == 0 {
		return units[0], nil
	}

	var parts []string
	i := 0
	prefix := ""

	for n > 0 {
		if nMod := n % 1000; nMod != 0 {
			j := 3
			if nMod < 4 {
				j = 2
			}

			switch i {
			case 0:
			case 1:
				prefix = " " + thousands[j]
			case 2:
				prefix = " " + millions[j]
			case 3:
				prefix = " " + billions[j]
			default:
				return "", fmt.Errorf("undefined prefix")
			}

			parts = append([]string{convertHundreds(nMod) + prefix}, parts...)
		}
		n /= 1000

		i++
	}

	result := strings.Join(parts, " ")

	return result, nil
}

func convertHundreds(n int) string {
	if n == 0 {
		return ""
	} else if n < 10 {
		return units[n]
	} else if n < 20 {
		return teenagers[n-10]
	} else if n < 100 {
		tensDigit := n / 10
		unitsDigit := n % 10

		if unitsDigit == 0 {
			return tens[tensDigit-2]
		}

		return fmt.Sprintf("%s %s", tens[tensDigit-2], units[unitsDigit])
	} else {
		hundredsDigit := n / 100
		tensDigit := n % 100 / 10
		unitsDigit := n % 10

		if tensDigit == 0 && unitsDigit == 0 {
			return hundreds[hundredsDigit]
		}

		return fmt.Sprintf("%s %s", hundreds[hundredsDigit], convertHundreds(n%100))
	}
}
