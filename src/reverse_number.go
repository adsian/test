package main

import (
	"fmt"
	"math"
	"strconv"
)

func main(){
	var numArray []int

	for i := 1000; i < 9999; i++{
    num := strconv.Itoa(i)
    reverseI := 0

    for i, v := range num {
    	v, _ := strconv.Atoi(string(v))
    	reverseI += v * int(math.Pow10(i))
    }
    if reverseI == i * 9 {
    	numArray = append(numArray, i)
    }
	}

	fmt.Println(numArray)
}



//	qianWei := i / 1000 % 10
//	baiWei := i / 100 % 10
//	shiWei := i / 10 % 10
//	geWei := i / 1 % 10
//	// fmt.Println(geWei, shiWei, baiWei ,qianWei)
//
//	if geWei == 0 {
//		continue
//	}
//
//	ReverseI := geWei * int(math.Pow10(3)) + shiWei * int(math.Pow10(2)) + baiWei * int(math.Pow10(1)) + qianWei
//	fmt.Println(ReverseI)
//
//	if ReverseI < i {
//		continue
//	} else if ReverseI == i * 9{
//		num = append(num, i)
//	} else {
//		continue
//	}
//}
//
//fmt.Println(num)