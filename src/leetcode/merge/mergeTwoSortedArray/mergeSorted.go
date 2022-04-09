package main

import "fmt"

func main(){
	fmt.Print(merge([]int{1,2,3,0,0,0},3,[]int{2,5,6},3))
}

func merge(nums1 []int, m int, nums2 []int, n int)  []int{
    res := make([]int,m+n)
	//m1,n1 := len(nums1),len(nums2)
    for i,j,k:=0,0,0;k < (m+n);k++{
        if(i<m && j<n){
			//fmt.Print("In",i,j,m,n,"\n")
            if(nums1[i] == nums2[j]){
                res[k] = nums1[i]
                k++
                i++
                res[k] = nums2[j]
                j++
            }else if(nums1[i] < nums2[j]){
                res[k] = nums1[i]
                i++
            }else{
                res[k] = nums2[j]
                j++
            }
        }else if(j==n){
				//fmt.Print("Just i")
                res[k] = nums1[i]
                i++
        }else if(i==m){
			//fmt.Print("Just j")
                res[k] = nums2[j]
                j++
        }
    }
    return res
}
