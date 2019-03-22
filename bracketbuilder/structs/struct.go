package structs

//Game defines a date and a team
type Game struct {
	Date string `json:"Date"`
	Team string `json:"Team"`
}

//JSONStruct defines a array of Game structs
type JSONStruct struct {
	Wins   []Game `json:"Wins"`
	Losses []Game `json:"Losses"`
}
