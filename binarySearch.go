import "fmt"
func search(nums []int, target int) int {
    return binarySearch(nums, target,0, len(nums)-1)
}

func binarySearch(nums []int, target int,left int, right int) int{
    fmt.Print("Checking if the target is out of the scode of the array")
    if nums[left]>target || nums[right]<target{
        return -1
    }
    mid := int((left + right)/2)
    if nums[mid]==target{return mid}
    if left==right{return -1}
    if nums[mid]>target{
        return binarySearch(nums, target, left, mid-1)   
    }
    return binarySearch(nums,target,mid+1,right)
}
