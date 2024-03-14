package utils

import (
	"bytes"
	"regexp"
	"strconv"
	"unicode"
)

var (
	CEPRegexp = regexp.MustCompile(`^\d{5}-?\d{3}$`)
)

func IsCEP(doc string, ufs ...FederativeUnit) bool {
	if !CEPRegexp.MatchString(doc) {
		return false
	}

	h, _ := strconv.Atoi(doc[0:3])

	if len(ufs) == 0 {
		return h >= 10
	}

	for _, uf := range ufs {
		if (uf == SP && h >= 10 && h <= 199) ||
			(uf == RJ && h >= 200 && h <= 289) ||
			(uf == ES && h >= 290 && h <= 299) ||
			(uf == MG && h >= 300 && h <= 399) ||
			(uf == BA && h >= 400 && h <= 489) ||
			(uf == SE && h >= 490 && h <= 499) ||
			(uf == PE && h >= 500 && h <= 569) ||
			(uf == AL && h >= 570 && h <= 579) ||
			(uf == PB && h >= 580 && h <= 589) ||
			(uf == RN && h >= 590 && h <= 599) ||
			(uf == CE && h >= 600 && h <= 639) ||
			(uf == PI && h >= 640 && h <= 649) ||
			(uf == MA && h >= 650 && h <= 659) ||
			(uf == PA && h >= 660 && h <= 688) ||
			(uf == AP && h == 689) ||
			(uf == AM && h >= 690 && h <= 692) ||
			(uf == RR && h == 693) ||
			(uf == AM && h >= 694 && h <= 698) ||
			(uf == AC && h == 699) ||
			(uf == DF && h >= 700 && h <= 727) ||
			(uf == GO && h >= 728 && h <= 729) ||
			(uf == DF && h >= 730 && h <= 736) ||
			(uf == GO && h >= 737 && h <= 767) ||
			(uf == RO && h >= 768 && h <= 769) ||
			(uf == TO && h >= 770 && h <= 779) ||
			(uf == MT && h >= 780 && h <= 788) ||
			(uf == MS && h >= 790 && h <= 799) ||
			(uf == PR && h >= 800 && h <= 879) ||
			(uf == SC && h >= 880 && h <= 899) ||
			(uf == RS && h >= 900 && h <= 999) {

			return true
		}
	}

	return false
}

type FederativeUnit uint8

const (
	AC FederativeUnit = iota
	AL
	AP
	AM
	BA
	CE
	DF
	ES
	GO
	MA
	MT
	MS
	MG
	PA
	PB
	PR
	PE
	PI
	RJ
	RN
	RS
	RO
	RR
	SC
	SP
	SE
	TO
)

func toInt(r rune) int {
	return int(r - '0')
}

func allDigit(doc string) bool {
	for _, r := range doc {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func IsCNH(doc string) bool {
	if len(doc) != 11 {
		return false
	}
	if !allDigit(doc) {
		return false
	}

	sum := 0
	acc := 9
	for _, r := range doc[:len(doc)-2] {
		sum += toInt(r) * acc
		acc--
	}

	base := 0
	digit1 := sum % 11
	if digit1 == 10 {
		base = -2
	}
	if digit1 > 9 {
		digit1 = 0
	}

	sum = 0
	acc = 1
	for _, r := range doc[:len(doc)-2] {
		sum += toInt(r) * acc
		acc++
	}

	var digit2 int
	if (sum%11)+base < 0 {
		digit2 = 11 + (sum % 11) + base
	}
	if (sum%11)+base >= 0 {
		digit2 = (sum % 11) + base
	}
	if digit2 > 9 {
		digit2 = 0
	}

	return toInt(rune(doc[len(doc)-2])) == digit1 &&
		toInt(rune(doc[len(doc)-1])) == digit2
}

var (
	CPFRegexp  = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	CNPJRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
)

func IsCPF(doc string) bool {
	const (
		size = 9
		pos  = 10
	)

	return isCPFOrCNPJ(doc, CPFRegexp, size, pos)
}

func isCPFOrCNPJ(doc string, pattern *regexp.Regexp, size int, position int) bool {
	if !pattern.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	if allEq(doc) {
		return false
	}

	d := doc[:size]
	digit := calculateDigit(d, position)

	d = d + digit
	digit = calculateDigit(d, position+1)

	return doc == d+digit
}

func cleanNonDigits(doc *string) {
	buf := bytes.NewBufferString("")
	for _, r := range *doc {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}

func allEq(doc string) bool {
	base := doc[0]
	for i := 1; i < len(doc); i++ {
		if base != doc[i] {
			return false
		}
	}

	return true
}

func calculateDigit(doc string, position int) string {
	var sum int
	for _, r := range doc {

		sum += toInt(r) * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}
