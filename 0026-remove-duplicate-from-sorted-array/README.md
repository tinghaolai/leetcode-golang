## Difficulty `Easy`

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


**Comment**

* last 的每次操作都要再額外加一，其實就是 current 的意思，既然這樣，不如直接設置成 1，避免可讀性降低
* 只要出現任意一次的「重複元素」，在跑完第二個迴圈後，下次進入迴圈時，仍需要判斷數值是否相同，因此之後每次必然會進入第二個迴圈
    * 找到相同數值後，進行陣列賦值動作，由於賦值對象是下一個元素，因此下次迴圈必定進入第二個迴圈，而初始值相同，因此必然於每次迴圈進入第二個迴圈
        * 變得複雜、不直覺、不容易懂
    * 應該可以試著用遞迴的方式改寫
* 本質上這個原理跟我的解法相同，但既然都需要遍歷一次，為何要寫 2 個迴圈增加複雜度?
    * 我的解法：單線程的一次跑完陣列
    * 此解法：若元素皆為 unique，由於最外層的只跑到倒數第二個元素，在進行第二層回圈判斷時，finder index 不會超過陣列最大 index
        * 只有此情況才會跑最後一行 return
        * finder index (遍歷 index ) 與 last index (賦值 index) 為相同值，同為單線程
    * 若有重複元素，finder 值會有偏移，因此有可能超過最大 index
        * 當有重複元素時，finder / last index 不同步，此時為兩條不同的路徑在遍歷陣列
        * 原本不會跑最後一個元素，因此 last index 偏移，因此會跑最後一個 element，於下次迴圈開始時，進入二次迴圈，finder + 1 即會超過陣列最大 index
    
### Online answer solution 2

* 需記錄最後一個元素，因為順序會變
* 第一個迴圈只跑至倒數第二個元素，迴圈中會與下個元素判斷，因此沒有必要跑最後一個元素
    * 若陣列長度為 4，並皆為 unique，index 只會跑到 2，return 卻只需要加 1，因為這邊的迴圈應該跟其他語言一樣，i 會跑到 i < 3 為止，所以 i 最後會是 3
* 迴圈中持續判斷當前元素是否與最後一個元素相等，是的話沒有必要繼續往下計算，答案也就是當前位置 + 1
* 若下個數值與當前數值相等，進行處理
    * 假設陣列為 01222345，就是對重複的部分 (22345) 開始進行處理
    * 掃到相異的數值時，進行交換 -> 32245
        * j + 1 代表下次交換的位置
        * 最終成為 34522

**Comment**

* func `removeElement` 
    * 為何要判斷 len = 0，如果是的話，前一個判斷式不是會炸開嗎 ?
    * 變數 j 其實只是變數 start 的重命名，本身是沒有意義的，但卻讓閱讀上變得容易，了解參數的涵義，而在計算中不會造成混淆
    * 這邊在做的事情其實就是 resorting，由於是 pre-sorted，所以不需要 bubble 之類的算法，只要簡單交換位置即是正確 ordering
* 位置交換後，只要當前值與最後一個值相等，就不用繼續往下跑迴圈，但其實在處理 resorting 的時候，還是變相地遍歷了整個陣列
    * 若陣列很長，是有可能要進行非常多次數值交換的
* 相較於我的解法，是直接捨棄重複值，進行覆蓋，將無法得知原本的陣列，而此解法是進行位置交換，所以可以回推原陣列
    * 假如 list 為 122345，我的解法會是 123455，此解法會變成 123452
