package solution

func canPlace(board *[]uint8, n, depth, i int) bool {
	ok := true
	// No need to check horizontal
	// Check vertical
	for j := 0; j < depth; j++ {
		if (*board)[j*n+i] != 0 {
			ok = false
			break
		}
	}
	if !ok {
		return false
	}
	// Check diagonals.
	// search towards top left
	for j, k := i-1, depth-1; j >= 0 && k >= 0; j, k = j-1, k-1 {
		if (*board)[k*n+j] != 0 {
			ok = false
			break
		}
	}
	if !ok {
		return false
	}
	// search towards top right
	for j, k := i+1, depth-1; j < n && k >= 0; j, k = j+1, k-1 {
		if (*board)[k*n+j] != 0 {
			ok = false
			break
		}
	}
	return ok
}

// board[x * n + y] <- board[x][y]
func dfs(board, placed *[]uint8, n, depth int) (ret uint64) {
	for i := 0; i < n; i++ {
		if (*placed)[i] != 0 || !canPlace(board, n, depth, i) {
			continue
		}
		// it is ok to place a queen here.
		if depth == n-1 {
			return 1
		} else {
			(*board)[depth*n+i] = 1
			(*placed)[i] = 1
			ret += dfs(board, placed, n, depth+1)
			(*board)[depth*n+i] = 0
			(*placed)[i] = 0
		}
	}
	return
}

func SolveQueen(n int) (ret uint64) {
	board := make([]uint8, n*n)
	placed := make([]uint8, n)
	return dfs(&board, &placed, n, 0)
}
