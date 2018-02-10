package coorl

import (
	"bytes"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
)

var curlPost = "curl -v  --request POST -H 'X-Custom-Header: myvalue' -H 'Content-Type: application/json' --data '{\"title\":\"Buy cheese and bread for breakfast.\"}' http://restapi3.apiary.io/notes"
var httpiePost = "echo '{\"title\":\"Buy cheese and bread for breakfast.\"}' | http -v POST http://restapi3.apiary.io/notes X-Custom-Header:myvalue Content-Type:application/json"

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Spec", func() {
		url := "http://restapi3.apiary.io/notes"
		jsonStr := []byte(`{"title":"Buy cheese and bread for breakfast."}`)
		b := bytes.NewReader(jsonStr)
		req, err := http.NewRequest("POST", url, b)
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		_, err = client.Do(req)
		if err != nil {
			panic(err)
		}
		g.It("Should parse curl", func() {
			curlCmd := AsCurl(req, b)
			g.Assert(curlCmd).Eql(curlPost)
		})
		g.It("Should parse httpie", func() {
			curlCmd := AsHttpie(req, b)
			g.Assert(curlCmd).Eql(httpiePost)
		})
	})
	return
}
