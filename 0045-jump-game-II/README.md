## Difficulty `Medium`

Runtime: 15 ms, faster than 62.24% of Go online submissions for Jump Game II.
Memory Usage: 6.2 MB, less than 41.49% of Go online submissions for Jump Game II.

### Self solution

* 若長度小於等於 1，代表已經到終點，return 0
* 預設步數至少為 1
* 迴圈判斷目前位置加此位置步數是否能夠到達終點
    * 若可以，回傳已走步數
    * 進入迴圈，步數 + 1
    * 遍歷所有可到達步數
        * 計算當前位置 + 可前進步數
        * 找到最大值即爲前往之位置（下次的 Index）
    * 迴圈結束，回傳步數


### Online solution 1

* 迴圈跑所有元素
* 判斷當前位置加可走步數是否為最大值
    * 若已抵達最後一個元素，回傳步數
    * 更新變數
* 若目前位置已是上次最大距離，找出當前最大值
    * 增加步數
    * 更新最大距離
    
> 邏輯更爲簡單，也不會有無限迴圈、超出長度的問題
