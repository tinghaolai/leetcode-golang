## Difficulty `Easy`

Runtime: 3 ms, faster than 27.58% of Go online submissions for Add Binary.
Memory Usage: 3 MB, less than 14.85% of Go online submissions for Add

### Self solution

* 遞迴邏輯
    * 如果數字都不需要判斷了，根據進位新增輸出
    * 計算目前位元輸出並計算進位
        * 由最小位元開始判斷，所以每次輸出都塞到前面
    * 往前推一個子元繼續遞迴
