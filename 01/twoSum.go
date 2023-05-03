// twoSum - returns indices of the two numbers such that they add up to target
func twoSum(nums []int, target int) []int {
    //init hashmap
    indexMap := make(map[int]int)
    
    for currIndex, currNum := range nums {
        if requiredIndex, ok := indexMap[target - currNum]; ok {
            return []int{requiredIndex, currIndex}
        }
        indexMap[currNum] = currIndex
    }

    return nil
}
