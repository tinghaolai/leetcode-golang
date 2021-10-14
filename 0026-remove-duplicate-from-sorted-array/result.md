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


### Online answer solution 1

* finder 為迴圈 index，last 表示目前 unique 位置
* 第一個迴圈跑全部的 element，第二個迴圈找出相同值
    * 若數值相同，則繼續尋找
    * Finder 可能超過長度，判斷並返回
* 若找到迴圈結束，代表找到相異值
    * 進行 index 賦值
    * last index ++ 計算下次位置
* 迴圈結束，回傳 last + 1


** Comment **

* last 的每次操作都要再額外加一，其實就是 current 的意思，既然這樣，不如直接設置成 1，避免可讀性降低
* 只要出現任意一次的「重複元素」，在跑完第二個迴圈後，下次進入迴圈時，仍需要判斷數值是否相同，因此之後每次必然會進入第二個迴圈
    * 變得複雜、不容易懂
    * 應該可以試著用遞迴的方式改寫
