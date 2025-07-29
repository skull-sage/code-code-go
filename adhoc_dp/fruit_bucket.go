package adhoc_dp

import "fmt"

func totalFruit(fruits []int) int {

	n := len(fruits)

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	type Bucket struct {
		f     int
		count int
	}

	bucketA := &Bucket{f: -1, count: 0}
	bucketB := &Bucket{f: -1, count: 0}

	max := 0

	//fmt.Println(bucketA, *bucketA)

	for idx := 0; idx < n; idx++ {

		f := fruits[idx]

		if f == bucketA.f {
			bucketA, bucketB = bucketB, bucketA
			bucketB.count++
		} else if f == bucketB.f {
			bucketB.count++
		} else {
			bucketA, bucketB = bucketB, bucketA

			bucketB.f = f
			bucketB.count = 1
		}

		fmt.Println(bucketA, bucketB)
		if max < bucketA.count+bucketB.count {
			max = bucketA.count + bucketB.count
		}
	}

	return max
}
