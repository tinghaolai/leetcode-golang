Runtime: 7 ms, faster than 88.21% of Go online submissions for Remove Duplicates from Sorted Array.

Memory Usage: 4.7 MB, less than 15.80% of Go online submissions for Remove Duplicates from Sorted Array.


### Self solution

* 如果空，則必為 0
* 若不為空，長度至少為一 (currentIndex = 1)，第一個元素確定 (currentValue = nums[0])
* currentIndex 代表下一個取代值的 Index，同時代表 unique 長度
* 若迴圈當前值與 currentValue 不相等時
    * 將值寫入目前位置
       * >若數值相等，代表重複，則不加入，達成 unique 的目的
    * currentIndex + 1，代表長度 +1，也計算出下次填入 index 位置
    * 賦予新值至 currentValue 進行下次比較
    
時間複雜度應為 n，空間複雜度為 1 
