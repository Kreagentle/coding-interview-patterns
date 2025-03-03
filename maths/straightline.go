package maths

func checkStraightLine(coordinates [][]int) bool {
	dx1 := coordinates[1][0] - coordinates[0][0]
	dy1 := coordinates[1][1] - coordinates[0][1]
	for i := 2; i < len(coordinates); i++ {
		dxn := coordinates[i][0] - coordinates[0][0]
		dyn := coordinates[i][1] - coordinates[0][1]
		if dxn*dy1 != dyn*dx1 {
			return false
		}
	}

	return true
}
