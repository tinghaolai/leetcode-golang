## Difficulty `Easy`

Runtime: 2 ms, faster than 29.29% of Go online submissions for Excel Sheet Column Title.
Memory Usage: 2 MB, less than 22.22% of Go online submissions for Excel Sheet Column Title.


### Self solution

* The trick part of this problem is that unlike other number system (ex: binary, octal, hexadecimal), the minimum number in 1, not 0
* Like regular number system, div the length of possible number which is 26
    * convert the mod number, and continually extract the div number you got
* So, what will the trick part cause regular solution error
    * let's say input is 52, which div number and mod number is 2, 0
    * But in correct calculate, only the number is greater than 26 can be the carry, so it's actually 1, 26
    * Solution is to check if mod number is 0 (which means "Z") when division
