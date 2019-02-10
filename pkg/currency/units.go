package currency

// Units maps Roman Numerals and its value
var Units map[string]int

// RegisterUnits registers Roman Numerals to its value
func RegisterUnits() {
	Units = make(map[string]int)
	Units["I"] = 1
	Units["V"] = 5
	Units["X"] = 10
	Units["L"] = 50
	Units["C"] = 100
	Units["D"] = 500
	Units["M"] = 1000
}
