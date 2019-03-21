package bracketfunctions

import (
	"bracketBuilder/myStructs"
)

func GameIterator(i int, iterator [][]string, m map[string]myStructs.JSONStruct, dates []string, dateIndex int) {
	if iterator[i][2] > iterator[i][4] {
		if _, ok := m[iterator[i][1]]; ok {
			wonGame := myStructs.Game{Date: dates[dateIndex], Team: iterator[i][3]}
			loss := m[iterator[i][1]].Losses
			wins := append(m[iterator[i][1]].Wins, wonGame)
			m[iterator[i][1]] = myStructs.JSONStruct{Wins: wins, Losses: loss}
			if _, ok := m[iterator[i][3]]; ok {
				lostGame := myStructs.Game{Date: dates[dateIndex], Team: iterator[i][1]}
				loss := append(m[iterator[i][3]].Losses, lostGame)
				wins := m[iterator[i][3]].Wins
				m[iterator[i][3]] = myStructs.JSONStruct{Wins: wins, Losses: loss}
			} else {
				lostGame := myStructs.Game{Date: dates[dateIndex], Team: iterator[i][1]}
				m[iterator[i][3]] = myStructs.JSONStruct{}
				loss := append(m[iterator[i][3]].Losses, lostGame)
				wins := m[iterator[i][3]].Wins
				m[iterator[i][3]] = myStructs.JSONStruct{Wins: wins, Losses: loss}
			}
		} else {
			wonGame := myStructs.Game{Date: dates[dateIndex], Team: iterator[i][3]}
			m[iterator[i][1]] = myStructs.JSONStruct{}
			loss := m[iterator[i][1]].Losses
			wins := append(m[iterator[i][1]].Wins, wonGame)
			m[iterator[i][1]] = myStructs.JSONStruct{Wins: wins, Losses: loss}
		}

	}
	if iterator[i][4] > iterator[i][2] {
		if _, ok := m[iterator[i][3]]; ok {
			wonGame := myStructs.Game{Date: dates[dateIndex], Team: iterator[i][1]}
			loss := m[iterator[i][3]].Losses
			wins := append(m[iterator[i][3]].Wins, wonGame)
			m[iterator[i][3]] = myStructs.JSONStruct{Wins: wins, Losses: loss}
			if _, ok := m[iterator[i][1]]; ok {
				lostGame := myStructs.Game{Date: dates[dateIndex], Team: iterator[i][3]}
				loss := append(m[iterator[i][1]].Losses, lostGame)
				wins := m[iterator[i][1]].Wins
				m[iterator[i][1]] = myStructs.JSONStruct{Wins: wins, Losses: loss}
			} else {
				lostGame := myStructs.Game{Date: dates[dateIndex], Team: iterator[i][3]}
				m[iterator[i][1]] = myStructs.JSONStruct{}
				loss := append(m[iterator[i][1]].Losses, lostGame)
				wins := m[iterator[i][1]].Wins
				m[iterator[i][1]] = myStructs.JSONStruct{Wins: wins, Losses: loss}
			}
		} else {
			wonGame := myStructs.Game{Date: dates[dateIndex], Team: iterator[i][1]}
			m[iterator[i][3]] = myStructs.JSONStruct{}
			loss := m[iterator[i][3]].Losses
			wins := append(m[iterator[i][3]].Wins, wonGame)
			m[iterator[i][3]] = myStructs.JSONStruct{Wins: wins, Losses: loss}
		}

	}
}
