package search_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])
	// first search target in which row
	if row == 1 && col == 1 {
		return matrix[0][0] == target
	}
	targetRow, is_found := binarySearch(&matrix, 0, row-1, target, -1)
	if is_found {
		return is_found
	}
	// search that row
	left, _ := binarySearch(&matrix, 0, col-1, target, targetRow)
	return matrix[targetRow][left] == target
}

func binarySearch(matrix *[][]int, left int, right int, target int, fix_row int) (int, bool) {
	lp := left
	rp := right
	fix_col := -1
	if fix_row == -1 {
		fix_col = 0
	}
	for lp < rp {
		mid := (lp + rp) / 2
		compareEle := 0
		if fix_row != -1 {
			compareEle = (*matrix)[fix_row][mid]
		} else {
			compareEle = (*matrix)[mid][fix_col]
		}
		if compareEle > target {
			rp = mid - 1
		} else if compareEle < target {
			lp = mid + 1
		} else if compareEle == target {
			return mid, true
		}
	}
	compareEle := 0
	if fix_row != -1 {
		compareEle = (*matrix)[fix_row][lp]
	} else {
		compareEle = (*matrix)[lp][fix_col]
	}
	if compareEle == target {
		return lp, true
	} else if compareEle > target && lp > 0 {
		return lp - 1, false
	}
	return lp, false
}
