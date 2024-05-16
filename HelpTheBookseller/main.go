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
		countMap[cat] = 0
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

	for ind, val := range listCat {
		if ind != len(listCat)-1 {
			total += fmt.Sprintf("(%s : %d) - ", val, countMap[val])
		} else {
			total += fmt.Sprintf("(%s : %d)", val, countMap[val])
		}
	}

	fmt.Println(totalCount)
	return total
}

func main() {
	fmt.Println(StockList([]string{"BBAR 0", "CDXE 0", "WVBSD 0", "BKWR 0", "BTSQ 0", "DRTY 0"}, []string{"A", "B", "W", "C", "D"}))
}
