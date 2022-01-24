package solution

func dfsMultiBad(board, placed []uint8, n, depth int, ret *chan uint64) {
	resultReceiver := make(chan uint64)
	resultCount := 0
	var ans uint64
	for i := 0; i < n; i++ {
		if !canPlace(&board, n, depth, i) {
			continue
		}
		// it is ok to place a queen here.
		if depth == n-1 {
			ans++
		} else {
			board[depth*n+i] = 1
			placed[i] = 1
			go dfsMultiBad(
				append(make([]uint8, 0, n*n), board...),
				append(make([]uint8, 0, n), placed...),
				n,
				depth+1,
				&resultReceiver,
			)
			resultCount++ // Recording how many values we should receive.
			board[depth*n+i] = 0
			placed[i] = 0
		}
	}
	for ; resultCount > 0; resultCount-- {
		ans += <-resultReceiver
	}
	*ret <- ans
}

func SolveQueenMultiBad(n int) uint64 {
	board := make([]uint8, n*n)
	placed := make([]uint8, n)
	result := make(chan uint64)
	go dfsMultiBad(board, placed, n, 0, &result)
	return <-result
}
