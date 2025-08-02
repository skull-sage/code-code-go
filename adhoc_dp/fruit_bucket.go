package adhoc_dp

import "fmt"

func totalFruit(fruits []int) int {

	type Bucket struct {
		f     int
		count int
	}

	bucketList := make([]*Bucket, 0)
	bucket := &Bucket{f: fruits[0], count: 1}

	for idx := 1; idx < len(fruits); idx++ {
		if bucket.f == fruits[idx] {
			bucket.count++
		} else {
			bucketList = append(bucketList, bucket)
			bucket = &Bucket{f: fruits[idx], count: 1}
		}
	}
	n := len(bucketList)

	if n == 1 {
		return bucketList[0].count
	}

	bucketA := bucketList[0]
	bucketB := bucketList[1]
	totalCountable := bucketA.count + bucketB.count
	max := totalCountable

	for idx := 2; idx < len(bucketList); idx++ {

		b := bucketList[idx]

		if b.f == bucketA.f || b.f == bucketB.f {
			totalCountable += b.count
		} else {
			bucketA = bucketList[idx-1]
			bucketB = bucketList[idx]
			totalCountable = bucketA.count + bucketB.count
		}

		fmt.Println("A:", bucketA, "B:", bucketB, "it:", b, "=>", totalCountable)

		if max < totalCountable {
			max = totalCountable
		}
	}

	return max
}
