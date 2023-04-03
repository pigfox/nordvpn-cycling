package main

type AvailableData struct {
	Cities    map[string][]string
	Countries []string
}

type PotentialData struct {
	Cities    map[string][]string
	Countries []string
}

var potentialData PotentialData
var availableData AvailableData

func setPotentialData() {
	potentialData.Countries = []string{"Albania", "Argentina", "Australia", "Austria", "Belgium", "Bosnia_And_Herzegovina",
		"Brazil", "Bulgaria", "Canada", "Chile", "Colombia", "Costa_Rica", "Croatia", "Cyprus", "Czech_Republic",
		"Denmark", "Estonia", "Finland", "France", "Georgia", "Germany", "Greece", "Hong_Kong", "Hungary", "Iceland",
		"Indonesia", "Ireland", "Israel", "Italy", "Japan", "Latvia", "Lithuania", "Luxembourg", "Malaysia", "Mexico",
		"Moldova", "Netherlands", "New_Zealand", "North_Macedonia", "Norway", "Poland", "Portugal", "Romania", "Serbia",
		"Singapore", "Slovakia", "Slovenia", "South_Africa", "South_Korea", "Spain", "Sweden", "Switzerland", "Taiwan",
		"Thailand", "Turkey", "Ukraine", "United_Kingdom", "United_States", "Vietnam"}

	potentialData.Cities = make(map[string][]string)
	potentialData.Cities["Albania"] = []string{"Tirana"}
	potentialData.Cities["Argentina"] = []string{"Buenos_Aires"}
	potentialData.Cities["Australia"] = []string{"Adelaide", "Brisbane", "Melbourne", "Perth", "Sydney"}
	potentialData.Cities["Austria"] = []string{"Vienna"}
	potentialData.Cities["Belgium"] = []string{"Brussels"}
	potentialData.Cities["Bosnia_And_Herzegovina"] = []string{"Sarajevo"}
	potentialData.Cities["Brazil"] = []string{"Sao_Paulo"}
	potentialData.Cities["Bulgaria"] = []string{"Sofia"}
	potentialData.Cities["Canada"] = []string{"Montreal", "Toronto", "Vancouver"}
	potentialData.Cities["Chile"] = []string{"Santiago"}
	potentialData.Cities["Colombia"] = []string{"Bogota"}
	potentialData.Cities["Costa_Rica"] = []string{"San_Jose"}
	potentialData.Cities["Croatia"] = []string{"Zagreb"}
	potentialData.Cities["Cyprus"] = []string{"Nicosia"}
	potentialData.Cities["Czech_Republic"] = []string{"Prague"}
	potentialData.Cities["Denmark"] = []string{"Copenhagen"}
	potentialData.Cities["Estonia"] = []string{"Tallinn"}
	potentialData.Cities["Finland"] = []string{"Helsinki"}
	potentialData.Cities["France"] = []string{"Marseille", "Paris"}
	potentialData.Cities["Georgia"] = []string{"Tbilisi"}
	potentialData.Cities["Germany"] = []string{"Berlin", "Frankfurt"}
	potentialData.Cities["Greece"] = []string{"Athens"}
	potentialData.Cities["Hong_Kong"] = []string{"Hong_Kong"}
	potentialData.Cities["Hungary"] = []string{"Budapest"}
	potentialData.Cities["Iceland"] = []string{"Reykjavik"}
	potentialData.Cities["Indonesia"] = []string{"Jakarta"}
	potentialData.Cities["Ireland"] = []string{"Dublin"}
	potentialData.Cities["Israel"] = []string{"Tel_Aviv"}
	potentialData.Cities["Italy"] = []string{"Milan"}
	potentialData.Cities["Japan"] = []string{"Osaka", "Tokyo"}
	potentialData.Cities["Latvia"] = []string{"Riga"}
	potentialData.Cities["Lithuania"] = []string{"Vilnius"}
	potentialData.Cities["Luxembourg"] = []string{"Luxembourg"}
	potentialData.Cities["Malaysia"] = []string{"Kuala_Lumpur"}
	potentialData.Cities["Mexico"] = []string{"Mexico"}
	potentialData.Cities["Moldova"] = []string{"Chisinau"}
	potentialData.Cities["Netherlands"] = []string{"Amsterdam"}
	potentialData.Cities["New_Zealand"] = []string{"Auckland"}
	potentialData.Cities["North_Macedonia"] = []string{"Skopje"}
	potentialData.Cities["Norway"] = []string{"Oslo"}
	potentialData.Cities["Poland"] = []string{"Warsaw"}
	potentialData.Cities["Portugal"] = []string{"Lisbon"}
	potentialData.Cities["Romania"] = []string{"Bucharest"}
	potentialData.Cities["Serbia"] = []string{"Belgrade"}
	potentialData.Cities["Singapore"] = []string{"Singapore"}
	potentialData.Cities["Slovakia"] = []string{"Bratislava"}
	potentialData.Cities["Slovenia"] = []string{"Ljubljana"}
	potentialData.Cities["South_Africa"] = []string{"Johannesburg"}
	potentialData.Cities["South_Korea"] = []string{"Seoul"}
	potentialData.Cities["Spain"] = []string{"Madrid"}
	potentialData.Cities["Sweden"] = []string{"Stockholm"}
	potentialData.Cities["Switzerland"] = []string{"Zurich"}
	potentialData.Cities["Taiwan"] = []string{"Taipei"}
	potentialData.Cities["Thailand"] = []string{"Bangkok"}
	potentialData.Cities["Turkey"] = []string{"Istanbul"}
	potentialData.Cities["Ukraine"] = []string{"Kyiv"}
	potentialData.Cities["United_Kingdom"] = []string{"Edinburgh", "Glasgow", "London", "Manchester"}
	potentialData.Cities["United_States"] = []string{"Atlanta", "Buffalo", "Charlotte", "Chicago", "Dallas", "Denver",
		"Kansas_City", "Los_Angeles", "Manassas", "Miami", "New_York", "Phoenix", "Saint_Louis", "Salt_Lake_City",
		"San_Francisco", "Seattle"}
	potentialData.Cities["Vietnam"] = []string{"Hanoi"}

	if len(potentialData.Countries) != len(potentialData.Cities) {
		panic("Data error")
	}
}
