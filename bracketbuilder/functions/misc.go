package functions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

// GetHTML : returns the html content of a specified url
func GetHTML(url string) string {
	resp, err := http.Get(url)
	// handle the error
	if err != nil {
		panic(err)
	}
	// remember to close the response body
	defer resp.Body.Close()
	// read the html into byte array
	html, _ := ioutil.ReadAll(resp.Body)
	//determine length of byte array for string conversion
	length := len(html)
	// convert byte array into a string
	newHTML := string(html[:length])
	return newHTML
}

// RegexParser : Used to find all regex matches within text, returns an array of an array of strings
func RegexParser(text string, regex string) [][]string {
	r, _ := regexp.Compile(regex)
	data := r.FindAllStringSubmatch(text, -1)
	return data
}

/*TextWriter : creates textfile and then writes to it
text : raw text to be added to a file
fileName : Name of the file
fileExt : File extension type (html, txt, json etc.)
*/
func TextWriter(text string, fileName string, fileExt string) {
	file, err := os.Create(fileName + "." + fileExt)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprintf(file, text)
}
