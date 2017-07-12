package chipmunk

import "fmt"

type UsernameGenerator struct {
	l *lcg
}

func NewUsernameGenerator(seed int) *UsernameGenerator {
	return &UsernameGenerator{l: newlcg(seed)}
}

func (u *UsernameGenerator) Generate() string {
	n := u.l.Next()
	return u.convert(n)
}

func (u *UsernameGenerator) convert(n int) string {
	// last 4 digit
	tail := n % 10000
	n = n / 10000
	secondLetter := string('A' + n%26)
	n = n / 26
	firstLetter := string('A' + n)
	return fmt.Sprintf("%s%s%04d", firstLetter, secondLetter, tail)
}
