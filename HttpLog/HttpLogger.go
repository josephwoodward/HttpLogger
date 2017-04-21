package HttpLog

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

var logDir = "./"

const LogFilename = "http_logger.log"

func handler(w http.ResponseWriter, request *http.Request) {

	if request.URL.Path[0:] == "/favicon.ico" {
		return
	}

	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(requestDump))

	f, err := os.Create(logDir + LogFilename)
	check(err)
	defer f.Close()

	x, err := f.Write([]byte("============== Request Begin ( " + time.Now().String() + ") ==============\n\n"))
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

func Log(options ConfigurationOptions) {
	//if options.RunAsDetached {
	//	fmt.Println("Detaching and listening on port " + options.Port)
	//}
	fmt.Println("Listening on port " + options.Port)

	logDir = options.LogDir

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+options.Port, nil)
}
