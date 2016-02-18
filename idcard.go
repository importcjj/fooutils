package fooutils

import (
	"strconv"
	
)


var(
	weightFactorys = [17]int{7,9,10,5,8,4,2,1,6,3,7,9,10,5,8,4,2}
	validBits = [17]string{"1","0","X","9","8","7","6","5","4","3","2"}
	
)


func IsValid(str_id string) bool {
	intSlice, last := NumsAndlast(str_id)
	sum := 0
	for i, v := range intSlice {
		sum += weightFactorys[i] * v
		
	}
	remainder := sum % 11
	if check := validBits[remainder]; check == last {
		return true
		
	}
	return false
	
}

func NumsAndlast(str_id string) ([]int, string) {
	l := len(str_id)
	var intSlice = make([]int, l-1)
	var last string
	for i, v := range str_id {
		if i == l-1 {
			last = string(v)
			break
			
		}
		n, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
			
		}
		intSlice[i] = n
		
	}
	return intSlice, last
	
}
