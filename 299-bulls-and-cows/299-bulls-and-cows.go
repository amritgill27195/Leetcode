func getHint(secret string, guess string) string {
    bulls := 0
    cows := 0
    sFreqMap := map[byte]int{}
    for i := 0; i < len(secret); i++ {
        sFreqMap[secret[i]]++
    }
    
    for i := 0; i < len(guess); i++ {
        gChar := guess[i]
        if _, ok := sFreqMap[gChar]; ok {
            if guess[i] == secret[i] {
                bulls++
                if val := sFreqMap[gChar]; val <= 0 {cows--} 
            } else {
                if val := sFreqMap[gChar]; val > 0 {cows++}
            }
            sFreqMap[gChar]--

        }
    }
    
    return fmt.Sprintf("%vA%vB", bulls,cows)
}







