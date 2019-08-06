package cards

type Cards struct {
	2	string	`json:"2"`
	Black	string	`json:"black"`
	Cards	[]Cards	`json:"cards"`
	Clubs	string	`json:"Clubs"`
}

