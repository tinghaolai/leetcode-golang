## Difficulty `Medium`

Runtime: 86 ms, faster than 30.05% of Go online submissions for Perfect Squares.
Memory Usage: 2.2 MB, less than 98.39% of Go online submissions for Perfect Squares.

### Self solution

* 看完題目，這種無法確認到底需要跑幾次的題目，直覺上會用遞迴來解
    * 第一步就是設置一個起始值，直接回傳呼叫自己的答案，先當中間有個 Black Box，會自動解決
    * 順向或反向推理這個演算法需要什麼，就慢慢想出來了
* 整體邏輯：自己的數字為 Perfect square 加上某個數字，而某個數字的計算，為呼叫自己即可取得，呼叫時又會判斷為幾個 Perfect square 的總合，並且回傳每種每個字數搭配的最小相加次數 
* 先從自己的數字往回找，如果自己就是 Perfect square，即是正確答案，進行回傳
* 如果不是，代表需要的數字量必然 + 1，此時判斷是否已達四平方定理的上限 (如果沒寫這個判斷，會導致 Time Limit Exceed)
* 接著判斷此次數字搭配是否為最小值
* 最後回傳不停呼叫自己所累加的相加次數
