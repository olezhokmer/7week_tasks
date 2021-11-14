package cipher

import (
	"strings"
	"unicode"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Shift struct {
	shift int
}

type Vigenere struct {
	key string
}

func NewShift(n int) Cipher {
	if n == 0 || n < -25 || n > 25 {
		return nil
	} else {
		return &Shift{shift: n}
	}
}

func NewCaesar() Cipher {
	return &Shift{shift: 3}
}

func NewVigenere(s string) Cipher {
	f := false
	for _, ch := range s {
		if ch < 'a' || ch > 'z' {
			return nil
		}
		if ch > 'a' {
			f = true
		}
	}
	if !f {
		return nil
	}
	return &Vigenere{key: s}
}

func (shift *Shift) Encode(s string) string {
	a := ""
	for _, ch := range strings.ToLower(s) {
		if unicode.IsLetter(ch) {
			ch += rune(shift.shift)
			if ch > 'z' {
				ch -= 26
			} else if ch < 'a' {
				ch += 26
			}
			a += string(ch)
		}
	}
	return a
}

func (shift *Shift) Decode(s string) string {
	a := ""
	for _, ch := range s {
		ch -= rune(shift.shift)
		if ch > 'z' {
			ch -= 26
		} else if ch < 'a' {
			ch += 26
		}
		a += string(ch)
	}
	return a
}

func (vigenere *Vigenere) Encode(s string) string {
	a := ""
	i := 0
	for _, ch := range strings.ToLower(s) {
		if unicode.IsLetter(ch) {
			a += string('a' + (ch-'a'+rune(vigenere.key[i%len(vigenere.key)])-'a')%26)
			i++
		}
	}
	return a
}

func (vigenere *Vigenere) Decode(s string) string {
	a := ""
	i := 0
	for _, ch := range strings.ToLower(s) {
		if unicode.IsLetter(ch) {
			a += string('a' + (ch+26-rune(vigenere.key[i%len(vigenere.key)]))%26)
			i++
		}
	}
	return a
}
