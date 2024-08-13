package main

func FindOddNumber(text []int) int {
	// TODO : start your code here
	mCount := make(map[int]int)

	// นับจำนวนครั้งที่ซ้ำ
	for _, num := range text {
		mCount[num]++
	}
	// หาเลขที่ซ้ำครั้งเดียว
	for num, count := range mCount {
		if count%2 != 0 {
			return num
		}
	}

	// ถ้าไม่มีเลขที่ซ้ำครั้งเดียว
	return -1
}
