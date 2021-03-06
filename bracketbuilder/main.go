package main

import (
	"bracketbuilder/functions"
	"bracketbuilder/structs"
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	dates := []string{"2018/11/06", "2018/11/07", "2018/11/08", "2018/11/09", "2018/11/10", "2018/11/11", "2018/11/12", "2018/11/13", "2018/11/14", "2018/11/15", "2018/11/16", "2018/11/17", "2018/11/18", "2018/11/19", "2018/11/20", "2018/11/21", "2018/11/22", "2018/11/23", "2018/11/24", "2018/11/25", "2018/11/26", "2018/11/27", "2018/11/28", "2018/11/29", "2018/11/30", "2018/12/01", "2018/12/02", "2018/12/03", "2018/12/04", "2018/12/05", "2018/12/06", "2018/12/07", "2018/12/08", "2018/12/09", "2018/12/10", "2018/12/11", "2018/12/12", "2018/12/13", "2018/12/14", "2018/12/15", "2018/12/16", "2018/12/17", "2018/12/18", "2018/12/19", "2018/12/20", "2018/12/21", "2018/12/22", "2018/12/23", "2018/12/24", "2018/12/25", "2018/12/26", "2018/12/27", "2018/12/28", "2018/12/29", "2018/12/30", "2018/12/31", "2019/01/01", "2019/01/02", "2019/01/03", "2019/01/04", "2019/01/05", "2019/01/06", "2019/01/07", "2019/01/08", "2019/01/09", "2019/01/10", "2019/01/11", "2019/01/12", "2019/01/13", "2019/01/14", "2019/01/15", "2019/01/16", "2019/01/17", "2019/01/18", "2019/01/19", "2019/01/20", "2019/01/21", "2019/01/22", "2019/01/23", "2019/01/24", "2019/01/25", "2019/01/26", "2019/01/27", "2019/01/28", "2019/01/29", "2019/01/30", "2019/01/31", "2019/02/01", "2019/02/02", "2019/02/03", "2019/02/04", "2019/02/05", "2019/02/06", "2019/02/07", "2019/02/08", "2019/02/09", "2019/02/10", "2019/02/11", "2019/02/12", "2019/02/13", "2019/02/14", "2019/02/15", "2019/02/16", "2019/02/17", "2019/02/18", "2019/02/19", "2019/02/20", "2019/02/21", "2019/02/22", "2019/02/23", "2019/02/24", "2019/02/25", "2019/02/26", "2019/02/27", "2019/02/28", "2019/03/01", "2019/03/02", "2019/03/03", "2019/03/04", "2019/03/05", "2019/03/06", "2019/03/07", "2019/3/08", "2019/03/09", "2019/03/10"}
	MyMap := map[string]string{
		"test_url": "https://www.ncaa.com/scoreboard/basketball-men/d1/2019/03/01/all-conf",
		"rootUrl":  "https://www.ncaa.com/scoreboard/basketball-men/d1/",
		"ending":   "/all-conf",
		"myRegex":  `<ul class="gamePod-game-teams">\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*<span class="gamePod-game-team-name">(.*)</span>\n.*<span class="gamePod-game-team-score">(.*)</span>\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*<span class="gamePod-game-team-name">(.*)</span>\n.*<span class="gamePod-game-team-score">(.*)</span>`,
	}
	m := make(map[string]structs.JSONStruct)
	var lock sync.RWMutex
	var wg sync.WaitGroup
	for dateIndex := range dates {
		dateURL := MyMap["rootUrl"] + dates[dateIndex] + MyMap["ending"]
		fmt.Println(dateURL)
		rawHTML, err := functions.GetHTML(dateURL)
		if err != nil {
			panic(err)
		}
		iterator := functions.RegexParser(rawHTML, MyMap["myRegex"])
		for i := range iterator {
			wg.Add(1)
			lock.Lock()
			go func(i int) {
				newMap := functions.GameIterator(i, iterator, m, dates, dateIndex)
				for k, v := range newMap {
					m[k] = v
				}
				lock.Unlock()
			}(i)
			wg.Done()
		}
	}
	wg.Wait()
	trial, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	newJSON := string(trial)
	fmt.Println(newJSON)

	functions.TextWriter(newJSON, "test", "json")
}
