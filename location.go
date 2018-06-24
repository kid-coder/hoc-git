package main

func location(city string) (string, string) {
	var (
		country   string
		continent string
	)

	switch city {
	case "Ha Noi", "Da Nang", "Sai Gon":
		country, continent = "Viet Nam", "Asia"
	case "Paris":
		country, continent = "France", "European"
	default:
		country, continent = "Unknow", "Unknow"
	}
	return country, continent
}
