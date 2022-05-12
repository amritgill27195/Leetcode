func candy(ratings []int) int {
    
    candies := make([]int, len(ratings))
    candies[0] = 1
    
    // left pass - if element is greater than its
    // immediate left element, then leftElementVal + 1 is the value for this position
    for i := 1; i < len(ratings); i++ {
        if ratings[i] > ratings[i-1] {
            candies[i] = candies[i-1]+1
        } else {
            candies[i] = 1
        }
    }
    
    // right pass 
    // if ith rating is greater than right element
    // and the ith value for this rating is NOT greater than
    // i+1 val, then rightVal + 1
    total := candies[len(candies)-1]
    for i := len(ratings)-2; i >= 0; i-- {
        if ratings[i] > ratings[i+1] && candies[i] < candies[i+1]+1 {
            candies[i] = candies[i+1]+1
        }
        total += candies[i]
    }
    return total
    
}