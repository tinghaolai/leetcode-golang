func guessNumber(n int) int {
    for (answer := guess(n); answer != 0; answer = guess(n)) {
        if answer == 1 {
                n ++
            } else {
                n --
            }
    }
    
    return n
}
