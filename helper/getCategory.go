package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/BurntSushi/toml"
	"github.com/mozillazg/request"
)

const (
	filename     = "leetcode.toml"
	loginPageURL = "https://leetcode.com/accounts/login/"
	refererURL   = "https://leetcode.com/"
)

var req *request.Request
var username string

func init() {
	cfg := config{}
	if _, err := toml.DecodeFile(filename, &cfg); err != nil {
		log.Fatalf(err.Error())
	}
	username = cfg.Login
	fmt.Printf("Hi, %s. I'm working for you.\n", cfg.Login)

	req = request.NewRequest(new(http.Client))
	req.Headers = map[string]string{
		"Accept-Encoding": "",
		"Referer":         refererURL,
	}

	// login
	csrfToken, err := getCSRFToken(req)
	if err != nil {
		log.Fatal(err)
	}
	req.Data = map[string]string{
		"csrfmiddlewaretoken": csrfToken,
		"login":               cfg.Login,
		"password":            cfg.Password,
	}

	if err = login(req); err != nil {
		log.Fatal(err)
	}
}

func getCategory(name string) *Category {
	URL := url(name)

	data := getJSON(URL)

	res := new(Category)
	if err := json.Unmarshal(data, res); err != nil {
		log.Fatal("无法把json转换成Category: " + err.Error())
	}

	return res
}

func url(s string) string {
	format := "https://leetcode.com/api/problems/%s/"
	return fmt.Sprintf(format, s)
}

func getJSON(URL string) []byte {
	log.Printf("开始下载 %s 的数据", URL)
	resp, err := req.Get(URL)
	if err != nil {
		log.Fatal("getJSON: Get Error: " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("getJSON: Read Error: " + err.Error())
	}
	return body
}

func getCSRFToken(req *request.Request) (string, error) {
	resp, err := req.Get(loginPageURL)
	if err != nil {
		return "", err
	}
	s, err := resp.Text()
	if err != nil {
		return "", err
	}

	reInput := regexp.MustCompile(
		`<input\s+[^>]*?name=['"]csrfmiddlewaretoken['"'][^>]*>`,
	)
	input := reInput.FindString(s)
	reValue := regexp.MustCompile(`value=['"]([^'"]+)['"]`)
	csrfToken := reValue.FindStringSubmatch(input)
	if len(csrfToken) < 2 {
		return "", err
	}
	return csrfToken[1], err
}

func login(req *request.Request) error {
	resp, err := req.Post(loginPageURL)
	defer resp.Body.Close() // **Don't forget close the response body**
	return err
}

type config struct {
	Login, Password string
}
