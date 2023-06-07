package main

func shipWithinDays(weights []int, days int) int {
	capLeft, capRight := 0, 0
	for i := 0; i < len(weights); i++ {
		capRight += weights[i]
		if capLeft < weights[i] {
			capLeft = weights[i]
		}
	}
	capLeft--
	for capRight-capLeft > 1 {
		mid := (capLeft + capRight) / 2
		day := calcDays(weights, mid, days)
		if day <= days && day != -1 {
			capRight = mid
		} else {
			capLeft = mid
		}
	}
	return capRight
}

func calcDays(weights []int, cap int, targetDays int) int {
	days := 1
	acc := 0
	for i := 0; i < len(weights); i++ {
		acc += weights[i]
		if acc > cap {
			days++
			if days > targetDays {
				return -1
			}
			acc = weights[i]
		}
	}

	return days
}
