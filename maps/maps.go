package main

import "fmt"

func main() {
	// using make
	mapMake := make(map[string]int)

	// literal syntax
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	// add key
	statePopulations["Georgia"] = 10310371

	// delete
	delete(statePopulations, "Texas")

	// print full map
	fmt.Println(statePopulations)
	fmt.Println(mapMake)

	// get keys
	fmt.Println(statePopulations["Ohio"])
	fmt.Println(statePopulations["Georgia"])

	// returns zero if key doesn't exist
	// ok allows for check that key exists
	key, ok := statePopulations["Texas"]
	fmt.Println(key, ok)
	_, ok = statePopulations["New York"]
	fmt.Println(ok)

	// get number of keys
	fmt.Println(len(statePopulations))

	// maps passed by reference - deleting from sp, deletes in every ref
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}
