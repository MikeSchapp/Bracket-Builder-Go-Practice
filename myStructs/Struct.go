package myStructs

type Game struct {
	Date string `json:"Date"`
	Team string `json:"Team"`
}
type JSONStruct struct {
	Wins   []Game `json:"Wins"`
	Losses []Game `json:"Losses"`
}
