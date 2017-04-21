package HttpLog

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func handler(w http.ResponseWriter, request *http.Request) {

	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(requestDump))

	f, err := os.Create("requestlog.log")
	check(err)
	defer f.Close()

	x, err := f.Write([]byte("============== Request Begin ( " + time.Now().String() + ") ==============\n"))
	_ = x
	check(err)

	p, err := f.Write(requestDump)
	_ = p
	check(err)

	y, err := f.Write([]byte("============== Request End ==============\n\n"))
	check(err)
	_ = y

	f.Sync()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Log(port string) {
	//fmt.Println("Listening on port " + port)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
