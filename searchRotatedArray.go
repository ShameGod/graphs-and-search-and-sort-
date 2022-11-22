func search(nums []int, target int) int {

    left:=0
    right:=len(nums)-1
    
    for left<=right {
        mid := int((left + right)/2)
        fmt.Print(right , " ", mid , " ", left, " " )
        if nums[mid] == target {return mid}
        //if nums[right] == target {return right}
        //if nums[left] == target {return left}
        if nums[mid] >= nums [left] {
            if target >= nums[left] && target<nums[mid] {
                right = mid-1
            }else {
                left = mid+1
            }
        }else {
            if target > mid && target<=nums[right] {
                left = mid +1
            }else {
                right = mid -1 
            }
        }
        
    }
    return -1 
    
    
}
