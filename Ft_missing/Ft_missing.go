package main


func Ft_missing(nums []int) int {
	
	missing := make([]int, len(nums)+1)

    for _, num := range nums {
        missing[num] = 1
    }

    for i := 0; i < len(missing); i++ {
        if missing[i] == 0 {
            return i
        }
    }

    return 0
}

