package matrices

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	firstRowZero, firstColZero := false, false

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
		}
	}

	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			setZeroColumn(matrix, j)
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			setZeroRaw(matrix, i)
		}
	}

	if firstRowZero {
		setZeroRaw(matrix, 0)
	}

	if firstColZero {
		setZeroColumn(matrix, 0)
	}
}

func setZeroColumn(matrix [][]int, column int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][column] = 0
	}
}

func setZeroRaw(matrix [][]int, raw int) {
	for j := 0; j < len(matrix[0]); j++ {
		matrix[raw][j] = 0
	}
}
