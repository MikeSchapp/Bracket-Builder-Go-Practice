package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

// GetHTML : returns the html content of a specific url
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

// RegexParser : Used to find all regex matches within text
func RegexParser(text string, regex string) {
	r, _ := regexp.Compile(regex)
	fmt.Println(r.FindAllString(text, -1))
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

func main() {

	MyMap := map[string]string{
		"test_url": "https://www.ncaa.com/scoreboard/basketball-men/d1/2019/03/01/all-conf",
		"rootUrl":  "https://www.ncaa.com/scoreboard/basketball-men/d1/",
		"ending":   "/all-conf",
	}
	test := MyMap["test_url"]
	rawHTML := GetHTML(test)
	TextWriter(rawHTML, "my", "html")
	r, _ := regexp.Compile(`<ul class="gamePod-game-teams">\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*<span class="gamePod-game-team-name">(.*)</span>\n.*<span class="gamePod-game-team-score">(.*)</span>\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*<span class="gamePod-game-team-name">(.*)</span>\n.*<span class="gamePod-game-team-score">(.*)</span>`)
	//fmt.Println(r.FindAllString(rawHTML, -1))
	type Game struct {
		Date string `json:"Date"`
		Team string `json:"Team"`
	}
	type JSONStruct struct {
		Wins   []Game `json:"Wins"`
		Losses []Game `json:"Losses"`
	}
	m := make(map[string]JSONStruct)
	dates := []string{"2019/02/19", "2019/02/20", "2019/02/21", "2019/02/22", "2019/02/23", "2019/02/24", "2019/02/25", "2019/02/26", "2019/02/27", "2019/02/28", "2019/03/01", "2019/03/02", "2019/03/03", "2019/03/04", "2019/03/05", "2019/03/06", "2019/03/07", "2019/3/08", "2019/03/09", "2019/03/10"}
	for dateIndex := range dates {
		dateURL := MyMap["rootUrl"] + dates[dateIndex] + MyMap["ending"]
		fmt.Println(dateURL)
		rawHTML := GetHTML(dateURL)
		iterator := r.FindAllStringSubmatch(rawHTML, -1)
		for i := range iterator {
			//iterates through all matches by group i being the group and 1-4 being the actual data points
			if iterator[i][2] > iterator[i][4] {
				if _, ok := m[iterator[i][1]]; ok {
					wonGame := Game{Date: dates[dateIndex], Team: iterator[i][3]}
					loss := m[iterator[i][1]].Losses
					wins := append(m[iterator[i][1]].Wins, wonGame)
					m[iterator[i][1]] = JSONStruct{Wins: wins, Losses: loss}
					if _, ok := m[iterator[i][3]]; ok {
						lostGame := Game{Date: dates[dateIndex], Team: iterator[i][1]}
						loss := append(m[iterator[i][3]].Losses, lostGame)
						wins := m[iterator[i][3]].Wins
						m[iterator[i][3]] = JSONStruct{Wins: wins, Losses: loss}
					} else {
						lostGame := Game{Date: dates[dateIndex], Team: iterator[i][1]}
						m[iterator[i][3]] = JSONStruct{}
						loss := append(m[iterator[i][3]].Losses, lostGame)
						wins := m[iterator[i][3]].Wins
						m[iterator[i][3]] = JSONStruct{Wins: wins, Losses: loss}
					}
				} else {
					wonGame := Game{Date: dates[dateIndex], Team: iterator[i][3]}
					m[iterator[i][1]] = JSONStruct{}
					loss := m[iterator[i][1]].Losses
					wins := append(m[iterator[i][1]].Wins, wonGame)
					m[iterator[i][1]] = JSONStruct{Wins: wins, Losses: loss}
				}

			}
			if iterator[i][4] > iterator[i][2] {
				if _, ok := m[iterator[i][3]]; ok {
					wonGame := Game{Date: dates[dateIndex], Team: iterator[i][1]}
					loss := m[iterator[i][3]].Losses
					wins := append(m[iterator[i][3]].Wins, wonGame)
					m[iterator[i][3]] = JSONStruct{Wins: wins, Losses: loss}
					if _, ok := m[iterator[i][1]]; ok {
						lostGame := Game{Date: dates[dateIndex], Team: iterator[i][3]}
						loss := append(m[iterator[i][1]].Losses, lostGame)
						wins := m[iterator[i][1]].Wins
						m[iterator[i][1]] = JSONStruct{Wins: wins, Losses: loss}
					} else {
						lostGame := Game{Date: dates[dateIndex], Team: iterator[i][3]}
						m[iterator[i][1]] = JSONStruct{}
						loss := append(m[iterator[i][1]].Losses, lostGame)
						wins := m[iterator[i][1]].Wins
						m[iterator[i][1]] = JSONStruct{Wins: wins, Losses: loss}
					}
				} else {
					wonGame := Game{Date: dates[dateIndex], Team: iterator[i][1]}
					m[iterator[i][3]] = JSONStruct{}
					loss := m[iterator[i][3]].Losses
					wins := append(m[iterator[i][3]].Wins, wonGame)
					m[iterator[i][3]] = JSONStruct{Wins: wins, Losses: loss}
				}

			}
		}
	}
	fmt.Println(m)
}
