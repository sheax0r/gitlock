package gh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	user  string
	token string
	repo  string
}

func NewClient(user string, token string, repo string) Client {
	return Client{
		user:  user,
		token: token,
		repo:  repo,
	}
}

func (c Client) getSha() (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/git/refs/heads/master", c.repo)
	resp, err := c.req("GET", url, "")

	var s struct {
		Object struct {
			Sha string `json:"sha"`
		} `json:"object"`
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&s)
	return s.Object.Sha, err
}

func (c Client) req(m string, url string, body string) (*http.Response, error) {
	r, err := http.NewRequest(m, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	r.SetBasicAuth(c.user, c.token)
	return http.DefaultClient.Do(r)
}

type ref struct {
	Ref string `json:"ref"`
	Sha string `json:"sha"`
}

func (c Client) Lock() error {
	sha, err := c.getSha()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("https://api.github.com/repos/%s/git/refs", c.repo)
	bytes, err := json.Marshal(&ref{Ref: "refs/heads/LOCK", Sha: sha})
	if err != nil {
		return err
	}
	resp, err := c.req("POST", url, string(bytes))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		bytes, _ = ioutil.ReadAll(resp.Body)
		return fmt.Errorf(string(bytes))
	}

	return nil
}

func (c Client) Unlock() error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/git/refs/heads/LOCK", c.repo)
	resp, err := c.req("DELETE", url, "")
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusAccepted {
		bytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf(string(bytes))
	}

	return nil
}
