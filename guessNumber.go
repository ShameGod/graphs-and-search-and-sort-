/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    left:=0
    right:=n
    mid:= int((left+right)/2)
    for true{
        guess:= guess(mid)
        if guess==0 {return mid}
        if guess<0 {
            right=mid-1
        }else{
            left=mid+1
        }
        mid=int((left+right)/2)
    }
    return mid
}
