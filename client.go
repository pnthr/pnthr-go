package main 

import (
	"fmt"
	"net/http"
	"os"
	"bytes"
	"net/url"
)

func main() {
	// Take url input from Command Line

	// If there are not 2 args, [0] is GO default,
	// [1] is what you type in until next space  
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "Please use http://host:port/page")
		os.Exit(1)
	}

	// Maually set payload
	payload := "PtDh+VSHLceJSRGnNOk=-53a1c59f62393700"

	// Parse the URL
	host, err := url.Parse(os.Args[1])
	checkError(err)

	// Set http Client Struct
	client := &http.Client{}

	// Call New Post Request to host
	// NewRequest(method, urlStr string, body io.Reader)

	// Make sure your URL has http:// !!!

	request, err := http.NewRequest("POST", "http://"+host.String(), bytes.NewBufferString(payload))
	checkError(err)

	// Add Header PNTHR
	// For Now we have default instead of dynamic
	request.Header.Add("pnthr", "534c33bb6637350002000000")

	response, err := client.Do(request)
	checkError(err)

	// Check Resppnse Status
	fmt.Println(response.Status)
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	var buf [512]byte
	reader := response.Body
	fmt.Println("got body")
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Println(string(buf[0:n]))
	}
	defer response.Body.Close()
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}