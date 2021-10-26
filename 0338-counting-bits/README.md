## Difficulty `Easy`


### Self solution
> 沒想出好解法，最好復習一下


### Online solution 1

```golang
func countBits(num int) []int {
    answer := []int{0}
    for i := 1; i <= num; i++ {
        answer = append(answer, answer[i & (i-1)] + 1)
    }

    return answer
}

```

* 如果是奇數，位元 and 完，即是上個數，1 的數量必然是上個數 + 1
* 若是偶數，與上次數相比，有可能進位 1 - n 次，而這些進位數必爲 1, 0(任意次數)
    * 例如  1001000 / 1000111
    * 因此進位數前 1 的數量保持一致，進位數後綜合必爲 1
    * 與前數 and 後 + 1 即爲正解
    
### Online solution 2 (java)

```java
public int[] countBits(int num) {
    int[] f = new int[num + 1];
    for (int i=1; i<=num; i++) f[i] = f[i >> 1] + (i & 1);
    return f;
}

```

* 數字往前 1 位元的數量 + 最後 1 個位元 = 1 的數量
