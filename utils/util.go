package utils

import (
	"math"
	"demo/controllers"
)

//Cordinate struct
type Cordinate struct {
	x    float64
	y    float64
	d    float64
}

//"fmt"

//GetCombinations Heap's algorithm combination see https://stackoverflow.com/a/30226442/3420649
func GetCombinations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

//GetMessage return message
func GetMessage(messages ...[]string) (msg string) {
	return ""
}

//GetLocation return x, y cordenate of emisor
//see https://en.wikipedia.org/wiki/True-range_multilateration
func GetLocation(distances ...float32) (x, y float32) {
	return 0,0
}


//get location given  three-point
func trilateration(c1, c2, c3 Cordinate) (x, y float64) {
	// https://confluence.slac.stanford.edu/display/IEPM/TULIP+Algorithm+Alternative+Trilateration+Method
	// Possiblility to divide by zero. Rearranging c1,c2,c3 may result in correct answer when that happens

	d1, d2, d3, i1, i2, i3, j1, j2, j3 := &c1.d, &c2.d, &c3.d, &c1.x, &c2.x, &c3.x, &c1.y, &c2.y, &c3.y

	x = ((((math.Pow(*d1, 2)-math.Pow(*d2, 2))+(math.Pow(*i2, 2)-math.Pow(*i1, 2))+(math.Pow(*j2, 2)-math.Pow(*j1, 2)))*(2**j3-2**j2) - ((math.Pow(*d2, 2)-math.Pow(*d3, 2))+(math.Pow(*i3, 2)-math.Pow(*i2, 2))+(math.Pow(*j3, 2)-math.Pow(*j2, 2)))*(2**j2-2**j1)) / ((2**i2-2**i3)*(2**j2-2**j1) - (2**i1-2**i2)*(2**j3-2**j2)))

	y = ((math.Pow(*d1, 2) - math.Pow(*d2, 2)) + (math.Pow(*i2, 2) - math.Pow(*i1, 2)) + (math.Pow(*j2, 2) - math.Pow(*j1, 2)) + x*(2**i1-2**i2)) / (2**j2 - 2**j1)

	return x, y
}

//Solved return solution
func Solved(Satellites ){

}
