package jobscheduler

import (
	"bytes"
	"code.google.com/p/go-charset/charset" //for convert xml charset ISO-8859-1 to UTF-8
	_ "code.google.com/p/go-charset/data"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetSpoolerFromResponseBody(resp *http.Response) *Spooler {
	spooler := Spooler{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[ERROR]: Cannot read response body: %s \n", err)
	}
	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.CharsetReader = charset.NewReader
	err = decoder.Decode(&spooler)
	if err != nil {
		fmt.Printf("[ERROR]: Cannot decord response: %s \n", err)
	}
	return &spooler
}

func (c *Client) CallApi(params interface{}) (*http.Response, *Error) {
	buf, _ := xml.MarshalIndent(params, "", " ")
	req, _ := http.NewRequest("POST", c.Url, strings.NewReader(string(buf)))
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("[ERROR]: Cannot access JobScheduler API: %s \n", err)
		return nil, &Error{Code: "error", Text: "Cannot access JobScheduler API "}
	}
	return resp, nil
}

func contains(elements []string, input string) bool {
	for _, element := range elements {
		if input == element {
			return true
		}
	}
	return false
}

func getFolderName(fullpath string) string {
	splited := strings.Split(fullpath, "/")
	return strings.Join(splited[:len(splited)-1], "/")
}
