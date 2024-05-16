package main

import (
	"fmt"
	"strconv"
)

func StockList(listArt []string, listCat []string) string {
	countBookInCatInt := 0
	total := ""
	countMap := make(map[string]int)
	for _, cat := range listCat {
		countBookInCatInt = 0
		for _, art := range listArt {
			if art[0] == cat[0] {
				for ind, val := range art {
					if val == ' ' {
						res, _ := strconv.Atoi(art[ind+1:])
						countBookInCatInt += res
					}
				}
				countMap[cat] = countBookInCatInt
			} else {
				countMap[cat] = countBookInCatInt
			}
		}
	}

	// fmt.Println(countMap)
	// fmt.Println(countBookInCatInt)
	for key, value := range countMap {
		total += fmt.Sprintf("(%s : %d)", key, value)
	}
	return total
}

func main() {
	fmt.Println(StockList([]string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}, []string{"A", "B", "C", "D"}))
}
