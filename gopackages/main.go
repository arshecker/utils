//файл ~/mygo/gopackges/main.go

package main

import (
	"fmt"
	newcolor "gopackages/color"
	"gopackages/random"
	"gopackages/wordz"

	"github.com/huandu/xstrings"

	"github.com/arshecker/utils"
	utilsV2 "github.com/arshecker/utils/v2"
	utilsV3 "github.com/arshecker/utils/v3"
)

func main() {
	isExistV3 := utilsV3.InSliceInt([]int{1, 2, 3, 4, 5}, 5)
	if isExistV3 {
		fmt.Println("Slice Int contain finding value")
		return
	}

	isExistV2 := utilsV2.InSlice(wordz.Words, "Two")
	if isExistV2 {
		fmt.Println("Slice Int contain finding value")
		return
	}

	isExistInt := utils.ContainsInt([]int{1, 2, 3, 4, 5}, 5)
	if isExistInt {
		fmt.Println("Slice Int contain finding value")
		return
	}

	isExist := utils.Contains(wordz.Words, "Two")
	if isExist {
		fmt.Println("Slice Words contain finding value")
		return
	}
	newcolor.Greet()
	city := random.City([]string{"Asghabat", "London", "Baku", "Ibiza", "New York", "Lebap"})
	digit := random.Digit([]string{"12121", "123213", "3311311", "555"})
	fmt.Println(city)
	fmt.Println(digit)
	shuffle_city := xstrings.Shuffle(city)
	fmt.Printf("Guess the city: %s", shuffle_city)

}
