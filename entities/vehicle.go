package entities

type Vehicle struct {
	ID                   string        `json:"id" bson:"_id"`
	Name                 string        `json:"name" bson:"name"`
	Model                string        `json:"model"`
	Manufacturer         string        `json:"manufacturer"`
	CostInCredits        string        `json:"cost_in_credits"`
	Length               string        `json:"length"`
	MaxAtmospheringSpeed string        `json:"max_atmosphering_speed"`
	Crew                 string        `json:"crew"`
	Passengers           string        `json:"passengers"`
	CargoCapacity        string        `json:"cargo_capacity"`
	Consumables          string        `json:"consumables"`
	VehicleClass         string        `json:"vehicle_class"`
	Pilots               []interface{} `json:"pilots"`
	Films                []string      `json:"films"`
	Created              string        `json:"created"`
	Edited               string        `json:"edited"`
	URL                  string        `json:"url"`
}
