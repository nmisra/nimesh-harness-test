package main

import("fmt")

func main() {
	//var [] int a
	a:=[] int { 1, 4, 2, 53, 23, 23, 42, -32, 33} 
	fmt.Println(a)
	fmt.Println("\n....sorted\n", merge_sort(a))
}

func merge_sort(arr [] int) [] int {
    // Write your code here.
    if len(arr)==1 { return arr}

	left:= merge_sort(arr[:len(arr)/2])
	right:= merge_sort(arr[len(arr)/2:]) 
	    
return merge(left, right)
}

func merge(a [] int, b [] int) (aux [] int){
    
    i:=0
    j:=0

    for (i<len(a) && j <len(b)) {
      if (a[i]<b[j]) {
          aux=append(aux, a[i])
          i++
        } else {
         aux=append(aux, b[j])
        j++
        }
	}
      for (i<len(a)) {
            aux=append(aux,a[i])
            i++
        }
     for (j<len(b)) {
         aux=append(aux,b[j])
         j++
       }
	return aux


