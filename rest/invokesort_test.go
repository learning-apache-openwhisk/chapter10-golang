package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	homedir "github.com/mitchellh/go-homedir"
)

func init() {
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true
	spew.Config.SortKeys = true
	log.SetOutput(ioutil.Discard)
}

func loadProps(props string) {
	file, _ := homedir.Expand(props)
	prp, _ := os.Open(file)
	defer prp.Close()
	scanner := bufio.NewScanner(prp)
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), "=")
		switch a[0] {
		case "APIHOST":
			os.Setenv("__OW_API_HOST", "https://"+a[1])
		case "AUTH":
			os.Setenv("__OW_API_KEY", a[1])
		case "NAMESPACE":
			os.Setenv("__OW_NAMESPACE", a[1])
		default:
		}
	}
}

func printMap(res map[string]interface{}) {
	for k, v := range res {
		fmt.Printf("%s=%v\n", k, v)
	}
}

func Example_url_auth() {
	loadProps("wskprops")
	fmt.Println(url("action"))
	fmt.Println(auth())
	// Output:
	// https://apihost.example.com/api/v1/namespaces/guest/action
	// XXXXX yyyyyy
}

func Example_mkMap_mkErr() {
	spew.Dump(mkMap("hello", "world"))
	spew.Dump(mkErr("wrong!"))
	spew.Dump(mkErr(fmt.Errorf("very wrong")))
	spew.Dump(mkErr(nil))
	// Output:
	// (map[string]interface {}) (len=1) {
	//  (string) (len=5) "hello": (string) (len=5) "world"
	// }
	// (map[string]interface {}) (len=1) {
	//  (string) (len=5) "error": (string) (len=6) "wrong!"
	// }
	// (map[string]interface {}) (len=1) {
	//  (string) (len=5) "error": (string) (len=10) "very wrong"
	// }
	// (map[string]interface {}) (len=1) {
	//  (string) (len=5) "error": (string) (len=5) "<nil>"
	// }
}

func Example_mkPost() {
	loadProps("wskprops")
	req, _ := mkPost("utils/date", mkMap("hello", "world"))
	req.GetBody = nil
	spew.Dump(req)
	// Output:
	// (*http.Request)({
	//  Method: (string) (len=4) "POST",
	//  URL: (*url.URL)(https://apihost.example.com/api/v1/namespaces/guest/utils/date),
	//  Proto: (string) (len=8) "HTTP/1.1",
	//  ProtoMajor: (int) 1,
	//  ProtoMinor: (int) 1,
	//  Header: (http.Header) (len=2) {
	//   (string) (len=13) "Authorization": ([]string) (len=1) {
	//    (string) (len=22) "Basic WFhYWFg6eXl5eXl5"
	//   },
	//   (string) (len=12) "Content-Type": ([]string) (len=1) {
	//    (string) (len=16) "application/json"
	//   }
	//  },
	//  Body: (ioutil.nopCloser) {
	//   Reader: (*bytes.Buffer)({"hello":"world"})
	//  },
	//  GetBody: (func() (io.ReadCloser, error)) <nil>,
	//  ContentLength: (int64) 17,
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

func Example_doCall() {
	loadProps("~/.wskprops")
	args := map[string]interface{}{
		"payload":   "a,b,c",
		"separator": ",",
	}
	req, _ := mkPost("/whisk.system/actions/utils/split?blocking=true", args)
	res := doCall(req)
	spew.Dump(res["response"])
	// Output:
	// (map[string]interface {}) (len=3) {
	//  (string) (len=6) "result": (map[string]interface {}) (len=2) {
	//   (string) (len=5) "lines": ([]interface {}) (len=3) {
	//    (string) (len=1) "a",
	//    (string) (len=1) "b",
	//    (string) (len=1) "c"
	//   },
	//   (string) (len=7) "payload": (string) (len=5) "a,b,c"
	//  },
	//  (string) (len=6) "status": (string) (len=7) "success",
	//  (string) (len=7) "success": (bool) true
	// }
}

func Example_invoke() {
	loadProps("~/.wskprops")
	args := map[string]interface{}{
		"payload":   "1:2:3",
		"separator": ":",
	}
	res := whiskInvoke("utils2/split", args, true, true)
	spew.Dump(res)
	// Output:
	// (map[string]interface {}) (len=2) {
	//  (string) (len=5) "lines": ([]interface {}) (len=3) {
	//   (string) (len=1) "1",
	//   (string) (len=1) "2",
	//   (string) (len=1) "3"
	//  },
	//  (string) (len=7) "payload": (string) (len=5) "1:2:3"
	// }
}

func ExampleInvoke() {
	loadProps("~/.wskprops")
	args := map[string]interface{}{
		"text":    "b,a,d,c",
		"action":  "utils2/sort",
		"message": "Sorted:",
	}
	spew.Dump(Invoke(args))
	// Output:
	// (map[string]interface {}) (len=1) {
	//  (string) (len=6) "result": (string) (len=15) "Sorted: a b c d"
	// }
}

func ExampleInvoke_err() {
	loadProps("~/.wskprops")
	args := map[string]interface{}{}
	spew.Dump(Invoke(args))
	args = map[string]interface{}{
		"action": "missing",
	}
	spew.Dump(Invoke(args))
	args = map[string]interface{}{
		"text":   "3,1,2",
		"action": "missing",
	}
	spew.Dump(Invoke(args))
	args = map[string]interface{}{
		"text":   "1,7,2,5",
		"action": "utils2/sort",
	}
	spew.Dump(Invoke(args))
	// Output:
	// (map[string]interface {}) (len=1) {
	//  (string) (len=5) "error": (string) (len=9) "no action"
	// }
	// (map[string]interface {}) (len=1) {
	//  (string) (len=5) "error": (string) (len=7) "no text"
	// }
	// (map[string]interface {}) (len=1) {
	//  (string) (len=5) "error": (string) (len=22) "cannot retrieve result"
	// }
	// (map[string]interface {}) (len=1) {
	//  (string) (len=6) "result": (string) (len=11) ">>> 1 2 5 7"
	// }
}
