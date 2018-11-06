package main

import "github.com/davecgh/go-spew/spew"

func Example_mkGet() {
	loadProps("wskprops")
	req, _ := mkGet("packages")
	spew.Dump(req)
	// Output:
	// (*http.Request)({
	//  Method: (string) (len=3) "GET",
	//  URL: (*url.URL)(https://apihost.example.com/api/v1/namespaces/guest/packages),
	//  Proto: (string) (len=8) "HTTP/1.1",
	//  ProtoMajor: (int) 1,
	//  ProtoMinor: (int) 1,
	//  Header: (http.Header) (len=1) {
	//   (string) (len=13) "Authorization": ([]string) (len=1) {
	//    (string) (len=22) "Basic WFhYWFg6eXl5eXl5"
	//   }
	//  },
	//  Body: (io.ReadCloser) <nil>,
	//  GetBody: (func() (io.ReadCloser, error)) <nil>,
	//  ContentLength: (int64) 0,
	//  TransferEncoding: ([]string) <nil>,
	//  Close: (bool) false,
	//  Host: (string) (len=19) "apihost.example.com",
	//  Form: (url.Values) <nil>,
	//  PostForm: (url.Values) <nil>,
	//  MultipartForm: (*multipart.Form)(<nil>),
	//  Trailer: (http.Header) <nil>,
	//  RemoteAddr: (string) "",
	//  RequestURI: (string) "",
	//  TLS: (*tls.ConnectionState)(<nil>),
	//  Cancel: (<-chan struct {}) <nil>,
	//  Response: (*http.Response)(<nil>),
	//  ctx: (context.Context) <nil>
	// })
}

func Example_whiskRetrieve() {
	loadProps("~/.wskprops")
	data := addMap(mkMap("payload", "1,2,3"), "separator", ",")
	id := whiskInvoke("utils2/split", data, false, false)["activationId"].(string)
	res := whiskRetrieve(id)
	spew.Dump(res["response"])
	// Output:
	// (map[string]interface {}) (len=3) {
	//  (string) (len=6) "result": (map[string]interface {}) (len=2) {
	//   (string) (len=5) "lines": ([]interface {}) (len=3) {
	//    (string) (len=1) "1",
	//    (string) (len=1) "2",
	//    (string) (len=1) "3"
	//   },
	//   (string) (len=7) "payload": (string) (len=5) "1,2,3"
	//  },
	//  (string) (len=6) "status": (string) (len=7) "success",
	//  (string) (len=7) "success": (bool) true
	// }
}

func ExampleFire() {
	loadProps("~/.wskprops")
	args := addMap(mkMap("trigger", "golang-trigger"),
		"text", "b,a,d,c")
	spew.Dump(Fire(args))
	// Output:
}
