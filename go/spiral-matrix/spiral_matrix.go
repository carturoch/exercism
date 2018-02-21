package spiralmatrix

func initEmptyMatrix(size int) (m [][]int) {
	m = make([][]int, size)
	for i := 0; i < size; i++ {
		m[i] = make([]int, size)
	}
	return
}

func cycle(array []rune) []rune {
	cycled := array[1:]
	cycled = append(cycled, array[0])
	return cycled
}

func SpiralMatrix(n int) [][]int {
	m := initEmptyMatrix(n)
	dirs := []rune{'R', 'D', 'L', 'U'}
	offsets := map[rune]int{'R': 0, 'D': 0, 'L': 0, 'U': 0}
	step, idx := 0, 0
	size := n
	for size > 0 {
		step++
		dir := dirs[0]
		start := offsets[dir]
		switch dir {
		case 'R':
			for i := 0; i < size; i++ {
				idx++
				m[start][start+i] = idx
			}
		case 'D':
			for i := 0; i < size; i++ {
				idx++
				m[start+i+1][n-start-1] = idx
			}
		case 'L':
			for i := 0; i < size; i++ {
				idx++
				m[n-start-1][n-2-start-i] = idx
			}
		case 'U':
			for i := 0; i < size; i++ {
				idx++
				m[n-2-start-i][start] = idx
			}
		}
		if step%2 != 0 {
			size--
		}
		offsets[dir] = start + 1
		dirs = cycle(dirs)
	}
	return m
}
