package hw10_program_optimization //nolint:golint,stylecheck

import (
	"github.com/valyala/fastjson"
	"io"
	"io/ioutil"
	"log"
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

var emails [100_000]string
var parser fastjson.Parser

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	err := getEmails(r)
	if err != nil {
		return nil, err
	}
	return countDomains(emails, domain)
}

func getEmails(r io.Reader) (err error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		v, err := parser.Parse(line)
		if err != nil {
			log.Fatal(err)
		}
		emails[i] = string(v.GetStringBytes("Email"))
	}
	return
}

func countDomains(e [100_000]string, domain string) (DomainStat, error) {
	result := make(DomainStat)
	for _, email := range e {
		if strings.Contains(email, domain) {
			result[strings.ToLower(strings.SplitN(email, "@", 2)[1])]++
		}
	}
	return result, nil
}
