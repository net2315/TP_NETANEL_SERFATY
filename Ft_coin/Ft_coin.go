package main

import (
	"math"
)

func Ft_coin(coins []int, amount int) int {
	tab := make([]int, amount+1)
    for i := range tab {
        tab[i] = math.MaxInt32
    }
    tab[0] = 0

    for _, coin := range coins {
        for i := coin; i <= amount; i++ {
            if tab[i-coin] != math.MaxInt32 {
                tab[i] = min(tab[i], tab[i-coin]+1)
            }
        }
    }
	if tab[amount] == math.MaxInt32 {
        return -1
    }
    return tab[amount]
}
