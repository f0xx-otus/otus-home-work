package hw10_program_optimization //nolint:golint,stylecheck

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strings"

	"github.com/mailru/easyjson"
)

type User struct {
	ID       int    `json:"-"`
	Name     string `json:"-"`
	Username string `json:"-"`
	Email    string
	Phone    string `json:"-"`
	Password string `json:"-"`
	Address  string `json:"-"`
}

type DomainStat map[string]int

var (
	emails         []string
	ErrEmptyDomain = errors.New("domain name can't be empty")
)

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	if domain == "" {
		return nil, ErrEmptyDomain
	}
	getEmails(r)
	return countDomains(emails, domain)
}

func getEmails(r io.Reader) {
	emails = []string{""}
	var user User
	reader := bufio.NewReader(r)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("can't read the line", err)
		}

		if err := easyjson.Unmarshal(line, &user); err != nil {
			log.Fatal("can't unmarshal the line", err)
		}
		emails = append(emails, user.Email)
	}
}

func countDomains(e []string, domain string) (DomainStat, error) {
	domain = strings.ToLower(domain)
	result := make(DomainStat)
	for _, email := range e {
		if strings.Contains(email, domain) {
			result[strings.ToLower(strings.SplitN(email, "@", 2)[1])]++
		}
	}
	return result, nil
}
