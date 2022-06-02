/*
    To make a palindrome, we need 2 SAME EXACT CHARACTERS
    
    approach: freqMap
    - We can build out a freqMap { $char: count }
    - Then since we know we can need 2 count of same exact characters - this also means if the count is 4 , 6,8 ( any even num ) - we can take them all
    - So we can loop over the freqMap
    - and if the char count is even, take them all and add to our len var. delete this key from map
    - if the count is not even, than we can take max even from this count ( count-1 on a odd count will give us the highest even count we can get )
    - Then we will add the even count from odd count and update the count of this odd count char with whatever is left after taking even out. ( count-evenCountThatWeTook )
    - Finally once we have looped over the entire freqMap;
    - if our map still has any data ( i.e len(freqMap) != 0 )
    - We can atleast take 1 character and add to our len ( remember we are not forming the string but just the longest possible len of a palindrome string )
    - Why can we take 1 character and add that to our len count?
    - Because a string is still a palindrome if we can 1 character in the middle.
        - For example: input; abccccdd
        - dccccd - but we can add either a or b in the middle and this will still be a palindrome
        - dccaccd or dccbccd
    
    
*/

// func longestPalindrome(s string) int {
//     freqMap := map[byte]int{}
//     for i := 0; i < len(s); i++ {
//         freqMap[s[i]]++
//     }
    
//     maxLen := 0
//     for char, count := range freqMap {
//         if count % 2 == 0 {
//             maxLen+=count
//             delete(freqMap, char)
//         } else {
//             evenCount := count-1
//             maxLen += evenCount
//             freqMap[char] = count-evenCount
//         }
//     }
    
//     if len(freqMap) != 0 {maxLen++}
//     return maxLen
// }


func longestPalindrome(s string) int {
    seen := map[byte]bool{}
    maxLen := 0
    for i := 0; i < len(s); i++ {
        _, ok := seen[s[i]]
        if ok {
            maxLen += 2
            delete(seen, s[i])
        } else {
            seen[s[i]]=false
        }
    }
    if len(seen) != 0 {
        maxLen++
    }
    return maxLen
}