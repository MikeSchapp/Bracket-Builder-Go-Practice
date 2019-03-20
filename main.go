package main

import (
	"encoding/json"
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

func main() {

	MyMap := map[string]string{
		"test_url": "https://www.ncaa.com/scoreboard/basketball-men/d1/2019/03/01/all-conf",
		"rootUrl":  "https://www.ncaa.com/scoreboard/basketball-men/d1/",
		"ending":   "/all-conf",
		"myRegex":  `<ul class="gamePod-game-teams">\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*<span class="gamePod-game-team-name">(.*)</span>\n.*<span class="gamePod-game-team-score">(.*)</span>\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*<span class="gamePod-game-team-name">(.*)</span>\n.*<span class="gamePod-game-team-score">(.*)</span>`,
	}
	type Game struct {
		Date string `json:"Date"`
		Team string `json:"Team"`
	}
	type JSONStruct struct {
		Wins   []Game `json:"Wins"`
		Losses []Game `json:"Losses"`
	}
	m := make(map[string]JSONStruct)
	dates := []string{"2018/11/06", "2018/11/07", "2018/11/08", "2018/11/09", "2018/11/10", "2018/11/11", "2018/11/12", "2018/11/13", "2018/11/14", "2018/11/15", "2018/11/16", "2018/11/17", "2018/11/18", "2018/11/19", "2018/11/20", "2018/11/21", "2018/11/22", "2018/11/23", "2018/11/24", "2018/11/25", "2018/11/26", "2018/11/27", "2018/11/28", "2018/11/29", "2018/11/30", "2018/12/01", "2018/12/02", "2018/12/03", "2018/12/04", "2018/12/05", "2018/12/06", "2018/12/07", "2018/12/08", "2018/12/09", "2018/12/10", "2018/12/11", "2018/12/12", "2018/12/13", "2018/12/14", "2018/12/15", "2018/12/16", "2018/12/17", "2018/12/18", "2018/12/19", "2018/12/20", "2018/12/21", "2018/12/22", "2018/12/23", "2018/12/24", "2018/12/25", "2018/12/26", "2018/12/27", "2018/12/28", "2018/12/29", "2018/12/30", "2018/12/31", "2019/01/01", "2019/01/02", "2019/01/03", "2019/01/04", "2019/01/05", "2019/01/06", "2019/01/07", "2019/01/08", "2019/01/09", "2019/01/10", "2019/01/11", "2019/01/12", "2019/01/13", "2019/01/14", "2019/01/15", "2019/01/16", "2019/01/17", "2019/01/18", "2019/01/19", "2019/01/20", "2019/01/21", "2019/01/22", "2019/01/23", "2019/01/24", "2019/01/25", "2019/01/26", "2019/01/27", "2019/01/28", "2019/01/29", "2019/01/30", "2019/01/31", "2019/02/01", "2019/02/02", "2019/02/03", "2019/02/04", "2019/02/05", "2019/02/06", "2019/02/07", "2019/02/08", "2019/02/09", "2019/02/10", "2019/02/11", "2019/02/12", "2019/02/13", "2019/02/14", "2019/02/15", "2019/02/16", "2019/02/17", "2019/02/18", "2019/02/19", "2019/02/20", "2019/02/21", "2019/02/22", "2019/02/23", "2019/02/24", "2019/02/25", "2019/02/26", "2019/02/27", "2019/02/28", "2019/03/01", "2019/03/02", "2019/03/03", "2019/03/04", "2019/03/05", "2019/03/06", "2019/03/07", "2019/3/08", "2019/03/09", "2019/03/10"}
	for dateIndex := range dates {
		dateURL := MyMap["rootUrl"] + dates[dateIndex] + MyMap["ending"]
		fmt.Println(dateURL)
		rawHTML := GetHTML(dateURL)
		iterator := RegexParser(rawHTML, MyMap["myRegex"])
		for i := range iterator {
			//iterates through all matches by group i being the group and 1-4 being the actual data points
			func(i int, iterator [][]string) {
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
			}(i, iterator)
		}
	}
	trial, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	newJSON := string(trial)
	fmt.Println(newJSON)

	TextWriter(newJSON, "test", "json")
}
