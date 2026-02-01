package hot

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	left, right := -1, m

	for right-1 > left {
		i := left + (right - left)
		j := (m+n+1)/2 - i - 2
		if nums1[i] <= nums2[j+1] {
			left = i
		} else {
			right = i
		}
	}

	i := left
	j := (m+n+1)/2 - i - 2
	left1, left2, right1, right2 := math.MinInt, math.MinInt, math.MaxInt, math.MaxInt
	if i >= 0 {
		left1 = nums1[i]
	}
	if i < m-1 {
		right1 = nums1[i+1]
	}
	if j >= 0 {
		left2 = nums2[j]
	}
	if j < n-1 {
		right2 = nums2[j+1]
	}
	leftmax := math.Max(float64(left1), float64(left2))
	rightmin := math.Min(float64(right1), float64(right2))
	if (m+n)%2 == 1 {
		return leftmax
	} else {
		return (leftmax + rightmin) / 2
	}
}
