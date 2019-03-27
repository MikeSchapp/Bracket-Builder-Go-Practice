package functions

import "bracketbuilder/structs"

//GameIterator  Takes a index value for i, the iterator object constructed from RegEx Parser, a map of JSONstructs and a list of dates
func GameIterator(i int, iterator [][]string, m map[string]structs.JSONStruct, dates []string, dateIndex int) map[string]structs.JSONStruct {
	if iterator[i][2] > iterator[i][4] {
		if _, ok := m[iterator[i][1]]; ok {
			wonGame := structs.Game{Date: dates[dateIndex], Team: iterator[i][3]}
			loss := m[iterator[i][1]].Losses
			wins := append(m[iterator[i][1]].Wins, wonGame)
			m[iterator[i][1]] = structs.JSONStruct{Wins: wins, Losses: loss}
			if _, ok := m[iterator[i][3]]; ok {
				lostGame := structs.Game{Date: dates[dateIndex], Team: iterator[i][1]}
				loss := append(m[iterator[i][3]].Losses, lostGame)
				wins := m[iterator[i][3]].Wins
				m[iterator[i][3]] = structs.JSONStruct{Wins: wins, Losses: loss}
			} else {
				lostGame := structs.Game{Date: dates[dateIndex], Team: iterator[i][1]}
				m[iterator[i][3]] = structs.JSONStruct{}
				loss := append(m[iterator[i][3]].Losses, lostGame)
				wins := m[iterator[i][3]].Wins
				m[iterator[i][3]] = structs.JSONStruct{Wins: wins, Losses: loss}
			}
		} else {
			wonGame := structs.Game{Date: dates[dateIndex], Team: iterator[i][3]}
			m[iterator[i][1]] = structs.JSONStruct{}
			loss := m[iterator[i][1]].Losses
			wins := append(m[iterator[i][1]].Wins, wonGame)
			m[iterator[i][1]] = structs.JSONStruct{Wins: wins, Losses: loss}
		}

	}
	if iterator[i][4] > iterator[i][2] {
		if _, ok := m[iterator[i][3]]; ok {
			wonGame := structs.Game{Date: dates[dateIndex], Team: iterator[i][1]}
			loss := m[iterator[i][3]].Losses
			wins := append(m[iterator[i][3]].Wins, wonGame)
			m[iterator[i][3]] = structs.JSONStruct{Wins: wins, Losses: loss}
			if _, ok := m[iterator[i][1]]; ok {
				lostGame := structs.Game{Date: dates[dateIndex], Team: iterator[i][3]}
				loss := append(m[iterator[i][1]].Losses, lostGame)
				wins := m[iterator[i][1]].Wins
				m[iterator[i][1]] = structs.JSONStruct{Wins: wins, Losses: loss}
			} else {
				lostGame := structs.Game{Date: dates[dateIndex], Team: iterator[i][3]}
				m[iterator[i][1]] = structs.JSONStruct{}
				loss := append(m[iterator[i][1]].Losses, lostGame)
				wins := m[iterator[i][1]].Wins
				m[iterator[i][1]] = structs.JSONStruct{Wins: wins, Losses: loss}
			}
		} else {
			wonGame := structs.Game{Date: dates[dateIndex], Team: iterator[i][1]}
			m[iterator[i][3]] = structs.JSONStruct{}
			loss := m[iterator[i][3]].Losses
			wins := append(m[iterator[i][3]].Wins, wonGame)
			m[iterator[i][3]] = structs.JSONStruct{Wins: wins, Losses: loss}
		}

	}
	return m
}
