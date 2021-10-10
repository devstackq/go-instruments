package main

import "log"

func main() {
	//countRepeat, rotate, seq
	// rotateLeft(2, []int32{1, 2, 3, 4, 5})
	// map[key], 1,5
	// pairs(2, []int32{1, 5, 3, 4, 2})

	sherlockAndAnagrams("abba")
}

func rotateLeft(d int32, arr []int32) []int32 {
	// Write your code here
	l := len(arr)
	result := make([]int32, l) //len, cap = 4
	dInt := int(d)             //cast

	for i, n := range arr {
		temp := i - dInt
		if temp < 0 {
			temp *= -1
			result[l-temp] = n // set value by index
		} else { // else new index = set value result array
			result[temp] = n //
		}
	}
	return result
}

// 1 5 3 4 2   arr = [1, 5, 3, 4, 2]
func pairs(k int32, arr []int32) int32 {
	//map
	var count int32
	//sort reverse, minus, compare?

	m := make(map[int32]bool)

	for _, n := range arr {
		m[n] = true
	}

	for _, n := range arr {
		if m[n-k] {
			count++
		}
	}

	//O(n)
	// for i := 0; i< len(arr);i++ {
	// 	m[arr[i]]=true//init key
	// 	// log.Println(arr[i])
	// 	if m[arr[i]- k]  { // map[1- 3]//3-1 == 1, true -> count+1
	// 		log.Println(arr[i]-k, "find key", arr[i], m)
	// 		count++
	// 	}
	// 	if  m[arr[i]+ k] {
	// 		log.Println(arr[i]-k, "find key2", arr[i], m)
	// 		count++
	// 	}
	// }

	// for i := len(sl)-1; i > 0; i-- {
	// 	for j := 0; j < len(sl); j++ {
	// 		// if sl[i] > sl[i] {
	// 			// log.Println(sl[j], sl[i])
	// 			if sl[i]- sl[j] == int(k) {
	// 				count++
	// 				break
	// 			}
	// 	}
	// }
	log.Println(count, "res")
	return count

}

func sherlockAndAnagrams(s string) int32 {
	// Write your code here
	var count int32

	m := make(map[string]bool)
	// m := make(map[int]bool)

	for i := range s {
		m[s[0:i]] = true // 0 :1, 0:2
	}
	log.Println(m, "bef")
	idx := 0

	for i := len(s) - 1; i >= 0; i-- {
		log.Println(i, i+idx, idx, "ww", s[i:i+idx+1])
		if m[s[i:i+idx+1]] {
			count++
		}
		idx++
	}
	log.Println(m, count, "res")
	return count
}
