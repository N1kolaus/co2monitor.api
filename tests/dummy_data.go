package tests

import "github.com/fminister/co2monitor.api/models"

var Locations = []models.Location{
	{
		Name: "test location 1",
	},
	{
		Name: "test location 2",
	},
}

var CO2 = []models.Co2Data{
	{
		LocationID: 1,
		CO2:        1001,
		Temp:       20.1,
	},
	{
		LocationID: 1,
		CO2:        2001,
		Temp:       22.1,
	},
	{
		LocationID: 2,
		CO2:        1002,
		Temp:       20.2,
	},
	{
		LocationID: 2,
		CO2:        2002,
		Temp:       22.2,
	},
}
