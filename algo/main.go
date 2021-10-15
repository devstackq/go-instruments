package main

import (
	"errors"
	"log"
)



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

//tree
// значения всегда отсортированы,
// Поиск, вставка, Удаление o log n - binary search
// Корень - если парент = нил
// Листы - узел без потомков
// Узел - указыатель на левый и правый дочерний элемент
// Ребро связывает 2 узал

func main() {
	//countRepeat, rotate, seq
	// rotateLeft(2, []int32{1, 2, 3, 4, 5})
	// map[key], 1,5
	// pairs(2, []int32{1, 5, 3, 4, 2})
	// sherlockAndAnagrams("abba")

	//nil data tree
	t := &TreeNode{data: 10} // set root
	t.Insert(2) //left
	t.Insert(12)//rigjht
	t.Insert(11)

 t.PrintTree()
 
 t.FindMin()
t.FindMax()

log.Println(t.SearchTree(11), "search")

}

type TreeNode struct {
	data interface{}
	left *TreeNode
	right *TreeNode
	parent *TreeNode
}


func (t *TreeNode)PrintTree() {

if  t == nil {
	return
}
log.Println(t.data.(int), "print")

t.left.PrintTree()
t.right.PrintTree()


// if t.right!= nil {
// 	t.right.PrintTree()
// }else {
// 	log.Println(t.right, "right list nil")
// }
	// switch v := t.data.(type){
	// case int:
	// 	// printInt(v)
	// 	fmt.Printf("v=digit: %d", v)
	// case string:
	// 	fmt.Printf("v=str data: %s", v)
	// }
}

find item, then concat prev with next Tree, Parent

func (t *TreeNode)RemoveTree(n int) {
	
	if t.data.(int) != n {
		if t.data.(int) > n {
			 t.left.SearchTree(n)
		}
		if t.data.(int) < n {
			 t.right.SearchTree(n)	
		}
	}else {
		t.left
	}
	//replace prev, next

}

func (t *TreeNode) SearchTree(n int) *TreeNode{
	
	if t.data.(int)  != n {
		//current node data > arg ? 
		if t.data.(int) > n {
			return t.left.SearchTree(n)	
		}
		if t.data.(int) < n {
			return t.right.SearchTree(n)	
		}
		// return t.data.(int)
	}
log.Println(t.data.(int), "find  epta")
	return t
}

func (t *TreeNode) FindMin() int {
	if t.left == nil {
		log.Println(t.data.(int), "min value ")
		return t.data.(int)
	}
	return t.left.FindMin()
}

func (t *TreeNode) FindMax() int {
	if t.right == nil {
		log.Println(t.data.(int), "find max value ")
		return t.data.(int)
	}
	return t.right.FindMax()
}

func(t *TreeNode)Insert(n int)error{

	if t == nil{
		return 	errors.New("tree is nil")
	}

	//case1 : 
	//check type - reflect ?
	if t.data.(int) > n {
	 if t.left == nil {
		// t.left.parent = t.left
		t.left = &TreeNode{data : n}
		return nil
	}

	return  t.left.Insert(n)
}
//case2 : nodeValue < N, insert right side
if t.data.(int) < n {
	if t.right == nil {
		//if list right == nil
		// t.parent = t
	   t.right = &TreeNode{data : n}
	   return nil		
   }
   //else start again - with current value, current scope
   return  t.right.Insert(n)
}
return nil
}