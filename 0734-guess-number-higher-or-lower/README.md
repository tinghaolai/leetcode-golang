## Difficulty `Easy`

### Self solution

> The worst solution

### Online solution

```golang
func guessNumber(n int) int {
	return sort.Search(n, func(x int) bool { return guess(x) <= 0 })
}
```

> 看懂源碼就能懂解法


```golang
func Search(n int, f func(int) bool) int {
    // Define f(-1) == false and f(n) == true.
    // Invariant: f(i-1) == false, f(j) == true.
    i, j := 0, n
    for i < j {
        h := int(uint(i+j) >> 1) // avoid overflow when computing h
        // i ≤ h < j
        if !f(h) {
            i = h + 1 // preserves f(i-1) == false
        } else {
            j = h // preserves f(j) == true
        }
    }
    // i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
    return i
}
```
