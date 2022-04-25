// approach: 
// reverse full
// reverse 0 : k-1 indicies
// reverse k : end
// MAKE SURE K is relative to array size: i.e k = k % len(arr)

func rotate(nums []int, k int)  {
    k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
    reverse(nums, k, len(nums) - 1);

}


func reverse(nums []int, low, high int) {
	for low < high {
		nums[low], nums[high] = nums[high], nums[low]
		low++
		high--
	}
}