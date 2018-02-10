package coorl

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// http -a USERNAME POST https://api.github.com/repos/jkbrzt/httpie/issues/83/comments body='HTTPie is awesome! :heart:'
func AsHttpie(r *http.Request, body io.ReadSeeker) (cmd string) {
	if r == nil {
		return
	}
	cmd = "http -v"
	cmd += " " + r.Method
	cmd += " " + r.URL.String()
	for name, value := range r.Header {
		cmd += fmt.Sprintf(" %s:%s", name, value[0])
	}
	if body != nil {
		body.Seek(0, 0)
		payload, err := ioutil.ReadAll(body)
		if err == nil {
			echo := fmt.Sprintf("echo '%s' | ", payload)
			cmd = echo + cmd
		}
	}
	return
}
