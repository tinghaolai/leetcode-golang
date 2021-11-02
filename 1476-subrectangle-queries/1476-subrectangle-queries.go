package leetcode

type SubrectangleQueries struct {
    Rectangle [][]int
}


func Constructor(rectangle [][]int) SubrectangleQueries {
    return SubrectangleQueries{
        Rectangle: rectangle}
}


func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
    for row1 <= row2 {
        currentCol := col1
        for currentCol <= col2 {
            this.Rectangle[row1][currentCol] = newValue

            currentCol++
        }

        row1++
    }
}


func (this *SubrectangleQueries) GetValue(row int, col int) int {
    return this.Rectangle[row][col]
}


/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
