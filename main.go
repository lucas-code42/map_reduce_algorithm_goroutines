package main

import (
	"fmt"
	"runtime"
	"sort"
)

func main() {
	runtime.GOMAXPROCS(8)
	list := []int{
		987, -76, 807, -878, -42, 210, 722, -260, -365, 614,
		405, -762, 500, -471, -560, -31, -555, -604, 114, -247,
		-61, 674, 311, 343, -925, 783, 451, 787, 739, 144,
		879, -948, 728, -398, 625, -496, 742, -784, -565, -588,
		-214, 320, -506, 391, 998, -287, -334, 41, 36, -21,
		740, 683, -483, -482, 75, -272, 126, -437, 173, -31,
		-485, 43, 572, -73, -620, 914, 843, -544, 30, 836,
		-721, 877, 901, 976, 433, -152, -421, -242, 164, -602,
		394, -402, 270, 138, 448, -197, -405, -641, -962, 670,
		127, 220, 520, -239, 83, -414, -117, 858, 806, -243,
		825, -420, 172, -816, 976, -989, 479, -699, -750, -792,
		-356, -476, 418, -427, -918, 748, 338, -269, 120, -763,
		-827, -878, -834, -882, 900, -554, 193, -369, 372, -489,
		-594, 865, -744, -485, -446, 235, -502, 917, -550, -329,
		764, 498, -285, 75, 441, 762, 226, 379, 899, 701,
		118, -642, 679, 883, -676, 953, 62, 418, 84, 707,
		785, 153, 20, 107, -572, 724, -367, 76, 520, -256,
		-331, -919, 642, 953, 247, -142, -745, 74, 629, -229,
		997, -15, -508, -36, -588, -399, -360, -112, 537, -461,
		109, -290, 406, -132, -329, -610, -765, -776, 138, -723,
		730, 191, 935, -911, 125, 735, -444, -315, -574, 4,
	}

	fmt.Println(list)

	cut := len(list) / 10
	list, s1 := divideList(list, cut)
	fmt.Println(list, s1)

	cut = len(list) / 8
	list, s2 := divideList(list, cut)
	fmt.Println(list, s2)

	cut = len(list) / 6
	list, s3 := divideList(list, cut)
	fmt.Println(list, s3)

	cut = len(list) / 4
	list, s4 := divideList(list, cut)
	fmt.Println(list, s4)

	c := make(chan []int)
	go sortNUmbersLowerToBigger(s1, c)
	go sortNUmbersLowerToBigger(s2, c)
	go sortNUmbersLowerToBigger(s3, c)
	go sortNUmbersLowerToBigger(s4, c)

	sorted1 := <-c
	sorted2 := <-c
	sorted3 := <-c
	sorted4 := <-c

	fmt.Println(sorted1)
	fmt.Println(sorted2)
	fmt.Println(sorted3)
	fmt.Println(sorted4)

	var finalList []int
	finalList = append(finalList, sorted1...)
	finalList = append(finalList, sorted2...)
	finalList = append(finalList, sorted3...)
	finalList = append(finalList, sorted4...)

	fmt.Println(finalList)
	tmp := finalList
	sort.Ints(tmp)

	finalList = sortNUmbers(finalList)

	fmt.Println(finalList)

	test := func() bool {
		if len(tmp) != len(finalList) {
			return false
		}
		for i, v := range tmp {
			if v != finalList[i] {
				return false
			}
		}
		return true
	}()

	fmt.Println(test)

}

func divideList(list []int, cut int) (tmp, s []int) {
	tmp = append(list[cut:])
	s = append(list[:cut])

	return tmp, s
}

// sortNUmbersBiggerToLower ordena os numeros menor para o maior
func sortNUmbersBiggerToLower(numbersList []int) []int {
	for j := 0; j < len(numbersList); j++ {
		for i := 0; i < len(numbersList); i++ {
			var temp int
			currentPosition := i
			nextPostion := i + 1

			if nextPostion == len(numbersList) {
				break
			}

			if numbersList[currentPosition] > numbersList[nextPostion] {
				temp = numbersList[currentPosition]
				numbersList[currentPosition] = numbersList[nextPostion]
				numbersList[nextPostion] = temp
			}
		}

	}
	return numbersList
}

// sortNUmbersBiggerToLower ordena os numeros do maior para o menor
func sortNUmbersLowerToBigger(numbersList []int, c chan []int) {
	for j := 0; j < len(numbersList); j++ {
		for i := 0; i < len(numbersList); i++ {

			var temp int
			currentPosition := i
			nextPostion := i + 1

			if nextPostion == len(numbersList) {
				break
			}

			if numbersList[currentPosition] < numbersList[nextPostion] {
				temp = numbersList[currentPosition]
				numbersList[currentPosition] = numbersList[nextPostion]
				numbersList[nextPostion] = temp

			}
		}

	}
	c <- numbersList
}

// sortNUmbers ordena os numeros do maior para o menor
func sortNUmbers(numbersList []int) []int {
	for j := 0; j < len(numbersList); j++ {
		for i := 0; i < len(numbersList); i++ {

			var temp int
			currentPosition := i
			nextPostion := i + 1

			if nextPostion == len(numbersList) {
				break
			}

			if numbersList[currentPosition] < numbersList[nextPostion] {
				temp = numbersList[currentPosition]
				numbersList[currentPosition] = numbersList[nextPostion]
				numbersList[nextPostion] = temp

			}
		}

	}
	return numbersList
}
