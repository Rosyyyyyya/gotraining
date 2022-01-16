package myclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	myHttpGet()

}

func myHttpGet() {

	rsp, err := http.Get("http://localhost:80/")
	if err != nil {
		fmt.Println("myHttpGet error is ", err)
		return
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("myHttpGet error is ", err)
		return
	}

	fmt.Println("response statuscode is ", rsp.StatusCode,
		"\nhead[name]=", rsp.Header["Name"],
		"\nbody is ", string(body))
}
