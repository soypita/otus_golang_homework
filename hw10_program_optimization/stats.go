package hw10_program_optimization //nolint:golint,stylecheck

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	return countDomains(r, domain)
}

func countDomains(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)
	reg := regexp.MustCompile("\\." + domain)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var user User
		err := user.UnmarshalJSON(scanner.Bytes())
		if err != nil {
			return nil, err
		}
		matched := reg.Match([]byte(user.Email))
		if matched {
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]++
		}
	}
	return result, nil
}
