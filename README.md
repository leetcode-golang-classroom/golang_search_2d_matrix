# golang_search_2d_matrix

Write an efficient algorithm that searches for a value `target` in an `m x n` integer matrix `matrix`. This matrix has the following properties:

- Integers in each row are sorted from left to right.
- The first integer of each row is greater than the last integer of the previous row.

## Examples

E**xample 1:**

![https://assets.leetcode.com/uploads/2020/10/05/mat.jpg](https://assets.leetcode.com/uploads/2020/10/05/mat.jpg)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/10/05/mat2.jpg](https://assets.leetcode.com/uploads/2020/10/05/mat2.jpg)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false

```

**Constraints:**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 100`
- `104 <= matrix[i][j], target <= $10^4$`

## 解析

題目給定一個 2 維的矩陣，然後這矩陣是一個排序過的矩陣

照著 row-base 方式排列，所以要找到有效率的方式去判斷值有沒有在矩陣

可以先把每個 row 的第一個值 去做二元搜尋找到 搜尋要搜尋數值在哪一個 row

然後在該 row 做搜二元搜尋

這樣時間複雜度會是 O(logN+logM)

## 程式碼

```go
func searchMatrix(matrix [][]int, target int) bool {
  row := len(matrix)
  col := len(matrix[0])
  // first search target in which row
  if row == 1 && col == 1 {
     return matrix[0][0] == target
  }
  left := 0
  right := row - 1
  targetRow := 0
  for left < right {
     mid := (left + right)/2
     if matrix[mid][0] > target {
        right = mid - 1
     } else if matrix[mid][0] < target {
        left = mid + 1
     } else if matrix[mid][0] == target {
         return true
     }
  }
  if matrix[left][0] == target {
      return true  
  } else if matrix[left][0] < target {
    targetRow = left
  } else if left > 0{
    targetRow = left - 1
  }
  // search that row
  left = 0
  right = col - 1
  for left < right {
     mid := (left + right)/2
     if matrix[targetRow][mid] > target {
        right = mid - 1
     } else if matrix[targetRow][mid] < target {
        left = mid + 1
     } else if matrix[targetRow][mid] == target {
         return true
     }
  }
  return matrix[targetRow][left] == target 
}
```

## 困難點

1. 邊界值收斂 需要注意當搜尋數值剛好等於搜尋收斂的那個值的判斷

## Solve Point

- [x]  Understand what the problem want to solve
- [x]  Analysis Complexity