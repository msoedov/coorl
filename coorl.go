package coorl

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func AsCurl(r *http.Request, body io.ReadSeeker) (cmd string) {
	if r == nil {
		return
	}
	cmd = "curl -v"
	cmd += "  --request " + r.Method
	for name, value := range r.Header {
		cmd += fmt.Sprintf(" -H '%s: %s'", name, value[0])
	}

	if body != nil {
		body.Seek(0, 0)
		payload, err := ioutil.ReadAll(body)
		if err == nil {
			cmd += fmt.Sprintf(" --data '%s'", payload)
		}
	}
	cmd += " " + r.URL.String()
	return
}
