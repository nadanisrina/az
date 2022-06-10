package models

type Truck struct {
	Amount          int64 `json:"amount"`
	MaintenaceTime  int64 `json:"maintenaceTime"`
	OperationalTime int64 `json:"operationalTime"`
}
