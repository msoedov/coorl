# Coorl

Translates `net/http` to a coooool curl command for debugging and troubleshooting

## Getting Started

Example

```go
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

fmt.Printf("%s\n", coorl.AsCurl(req, b))

```
```bash
curl -v  --request POST -H 'X-Custom-Header: myvalue' -H 'Content-Type: application/json' --data '{\"title\":\"Buy cheese and bread for breakfast.\"}' http://restapi3.apiary.io/notes
```

Translate to `httpie`

```go
fmt.Printf("%s\n", coorl.AsHttpie(req, b))
```

```shell
echo '{\"title\":\"Buy cheese and bread for breakfast.\"}' | http -v POST http://restapi3.apiary.io/notes X-Custom-Header:myvalue Content-Type:application/json
```
### Prerequisites
```go
package coorl // import "github.com/msoedov/netcurl"

func AsCurl(r *http.Request, body io.ReadSeeker) (cmd string)
func AsHttpie(r *http.Request, body io.ReadSeeker) (cmd string)
```

### Installing

A step by step series of examples that tell you have to get a development env running

```shell
go get -u github.com/msoedov/coorl
```

## Running the tests

```
go test
```


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments
