package entities

type ImportData struct {
	Articles  []Article  `json:"articles"`
	People    []Person   `json:"people"`
	Starships []Starship `json:"starships"`
	Vehicles  []Vehicle  `json:"vehicles"`
	Users     []User     `json:"users"`
}
