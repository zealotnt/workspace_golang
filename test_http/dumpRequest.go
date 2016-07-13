/******************************************************/
// Handler debug
/******************************************************/

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

if true {
	fmt.Println("**************************************")
	fmt.Println("httputil.DumpRequest")
	fmt.Println("**************************************")
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		JSON(w, err, 399)
		return
	}
	fmt.Printf("%s", dump)
}

if true {
	fmt.Println("**************************************")
	fmt.Println("ioutil.ReadAll")
	fmt.Println("**************************************")
	body, _ := ioutil.ReadAll(r.Body)
	n := bytes.IndexByte(body, 0)
	if n != -1 {
		fmt.Printf(string(body[:n]))
	} else {
		fmt.Printf(string(body))
	}
}

if true {
	fmt.Println("**************************************")
	fmt.Println("Testing form.Image <--> multipart.Fileheader")
	fmt.Println("**************************************")
	fmt.Println(form.Image.Filename)
	fmt.Println(form.Image.Header)
	fmt.Println(form.Image.Header.Get("Content-Type"))
	fmt.Println(form.Image.Header.Get("Content-Disposition"))
	form.Image.Header.Del("Content-Disposition")
	form.Image.Header.Del("Content-Type")
	form.Image.Filename = "123"
	fmt.Println(form.Image.Header.Get("Content-Type"))
	fmt.Println(form.Image.Header.Get("Content-Disposition"))
}

if true {
	fmt.Println("**************************************")
	fmt.Println("Testing form.Image <--> multipart.Fileheader Open() function")
	fmt.Println("**************************************")
	fmt.Println("Len of file = ", len(form.ImageData()))
	fmt.Println(form.Image.Filename)
	fmt.Println(form.Image.Header)
}



/******************************************************/
// Binding debug
/******************************************************/
import (
	"encoding/json"
	"fmt"
)

if true {
	errOutput, _ := json.Marshal(errs)
	fmt.Println(string(errOutput))
}