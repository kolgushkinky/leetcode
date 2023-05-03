// twoSum - returns indices of the two numbers such that they add up to target
// O(n^2) solution
func twoSum(nums []int, target int) []int {
    for i, num := range nums {
        for j, num2 := range nums[i+1:] {
            if num + num2 == target {
                return []int{i, j+i+1}
            }
        }
    }
    return nil
}
