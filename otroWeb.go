package main
 
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
 
func main() {
 
	// Basic HTTP GET request
	resp, err := http.Get("http://www.google.com")
	iferr != nil {
		log.Fatal("Error getting response. ", err)
    }
	defer resp.Body.Close()
 
	// Read body from response
	body, err := ioutil.ReadAll(resp.Body)
	iferr != nil {
		log.Fatal("Error reading response ", err) 
	}
 
	fmt.Printf("%s\n", body)
}