func isValidSudoku(board [][]byte) bool {
	for i:=0; i<9; i++ {
		// declaring outside cuases reuse, leads to bug
		rowSet := make(map[byte]bool);
		colSet := make(map[byte]bool);
		for j:=0; j<9; j++ {
			// [j][i] --> row
			if board[j][i] != '.' {
				if colSet[board[j][i]] {
					return false;
				}
				colSet[board[j][i]] = true;
			}
			// [i][j] --> col
			if board[i][j] != '.' {
				if rowSet[board[i][j]] {
					return false;
				}
			    rowSet[board[i][j]] = true;
			}
		}
	}

	for sq:=0; sq<9; sq++ {
		sqSet := make(map[byte]bool);
		for i:=0;i<3;i++ {
			for j:=0;j<3;j++ {
				row := (sq / 3) * 3 + i;
				col := (sq % 3) * 3 + j;
				if board[row][col] != '.' {
					if sqSet[board[row][col]] {
						return false;
					}
					sqSet[board[row][col]] = true;
				}
			}
		}
	}

	return true;
}
