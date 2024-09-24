package main

import "fmt"



func Ft_non_overlap(intervals [][]int) int {
	minLen := len(intervals)
    for _, interval := range intervals {
        if len(interval) < minLen {
            minLen = len(interval)
        }
    }
    return minLen
}
func main()  {
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	
}
