package main

import (
	"errors"
	"fmt"
	"os"
)

type grid [10][10][10]int

func main() {

	g := setGrid()
	//g := initgrid()
	//getGrid(&g)
	fmt.Println("The Given grid is...")
	printGrid(g)

	fmt.Println("Solving this....")
	solution := solve(g)
	fmt.Println("Solution is....")
	printGrid(solution)

}

func solve(g grid) grid {

	if !isValid(&g) {
		fmt.Println("The Given Grid is not valid.. Please check")
		os.Exit(1)
	}

	isSolved := simpleSolver(&g)
	if isSolved {
		return g
	}

	//we have not solved this yet, so need to use backtracking to solve this
	backTrack(&g)
	if !isValid(&g) {
		fmt.Println("The Given Grid is not valid.. Please check")
		os.Exit(1)
	}

	return g
}

func simpleSolver(g *grid) bool {
	var newCellCnt, currCellCnt int

	for c := 0; c < 50 && isValid(g); c++ {

		currCellCnt = countFilled(g)
		placeSingletons(g)
		placeSingleOptions(g)
		newCellCnt = countFilled(g)
		if newCellCnt == 81 {
			//we have solved the grid
			return true
		}

		if currCellCnt == newCellCnt {
			fmt.Println("We are not making progress... ")
			break
		}
	}

	return false
}

func backTrack(g *grid) error {

	if !isValid(g) {
		return errors.New("notvalid")
	}

	ig,jg := getCellToGuess(*g)
	for k := 1 ; k < 10 ; k ++ {
		if g[ig][jg][k] != 0 {
			fmt.Println("Guessing at cell ", ig, ", ", jg, " Val: ", k)
			gCopy := makeCopy(g)
			gCopy[ig][jg][0] = k
			removeFilledOptions(&gCopy,ig,jg)
			isSolved := simpleSolver(&gCopy)

			if isSolved {
				*g = gCopy
				return nil
			}

			if !isValid(&gCopy) {
				return errors.New("notvalid")
			}

			//the grid is still valif & we do not have a solution.. so back track again
			err := backTrack(&gCopy)
			if err == nil {
				*g = gCopy
				return nil
			}
		}
	}

	//fmt.Println("Solved by backtracking")
	//return gret,nil

	return nil
}

func makeCopy(g *grid) grid {
	var r grid
	for i := 0; i < 10; i ++ {
		for j := 0; j < 10 ; j ++{
			for k:=0;k<10;k++ {
				r[i][j][k] = g[i][j][k]
			}
		}
	}

	return r
}


func getCellToGuess(g grid) (int,int) {

	minCell := 10
	i1,j1 := 0 ,0

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if g[i][j][0] == 0 {
				var c int
				for k := 1; k < 10 ;k ++{
					if g[i][j][k] != 0 {
						c++
					}
				}

				if c < minCell {
					minCell = c
					i1,j1 = i,j
				}
			}
		}
	}

	return i1,j1
}

func printGrid(inp grid) {
	fmt.Println(`╔═══╤═══╤═══╦═══╤═══╤═══╦═══╤═══╤═══╗`)

	var endStr string

	for i := 1; i < 10; i++ {
		fmt.Print("║ ")
		for j := 1; j < 10; j++ {
			if j == 3 || j == 6 || j == 9 {
				endStr = ` ║ `
			} else {
				endStr = ` │ `
			}
			if inp[i][j][0] == 0 {
				fmt.Print(" ", endStr)
			} else {
				fmt.Print(inp[i][j][0], endStr)
			}
		}
		fmt.Print("\n")

		if i == 3 || i == 6 {
			fmt.Println(`╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣`)
		} else if i == 9 {
			fmt.Println(`╚═══╧═══╧═══╩═══╧═══╧═══╩═══╧═══╧═══╝`)
		} else {
			fmt.Println("╟───┼───┼───╫───┼───┼───╫───┼───┼───╢")
		}
	}
}

func getGrid(inp *grid) {

	var row, col, val int

	for {

		fmt.Print("Please enter the grid postions and it value. Set 0,0,0 to finish input: ")
		_, err := fmt.Scanf("%d %d %d", &row, &col, &val)
		if err != nil {
			fmt.Println("Improper input.. Please input again")
			continue
		}

		if row == 0 && col == 0 && val == 0 {
			break
		}
		if row > 9 || row < 1 {
			fmt.Println("Improper input.. Please input again")
			continue
		}

		if col > 9 || col < 1 {
			fmt.Println("Improper input.. Please input again")
			continue
		}

		if val > 9 || val < 1 {
			fmt.Println("Improper input.. Please input again")
			continue
		}

		inp[row][col][0] = val
		for k := 1; k < 10; k++ {
			inp[row][col][k] = 0
		}
		printGrid(*inp)
	}

}

func initgrid() grid {
	var g grid
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			for k := 1; k < 10; k++ {
				g[i][j][k] = k
			}
		}
	}

	return g
}

func setGrid() grid {

	//╔═══╤═══╤═══╦═══╤═══╤═══╦═══╤═══╤═══╗
	//║   │ 9 │   ║   │   │   ║ 8 │   │ 4 ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║ 7 │   │   ║   │ 5 │ 8 ║   │   │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │ 6 │ 8 ║ 9 │   │   ║   │   │   ║
	//╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣
	//║ 6 │ 4 │ 1 ║   │ 7 │   ║   │   │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │   │ 2 ║   │   │   ║ 4 │   │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │   │   ║   │ 1 │   ║ 6 │ 7 │ 3 ║
	//╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣
	//║   │   │   ║   │   │ 9 ║ 1 │ 5 │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │   │   ║ 2 │ 8 │   ║   │   │ 7 ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║ 4 │   │ 7 ║   │   │   ║   │ 6 │   ║
	//╚═══╧═══╧═══╩═══╧═══╧═══╩═══╧═══╧═══╝

	// SOLUTION
	//╔═══╤═══╤═══╦═══╤═══╤═══╦═══╤═══╤═══╗
	//║ 5 │ 9 │ 3 ║ 1 │ 6 │ 7 ║ 8 │ 2 │ 4 ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║ 7 │ 2 │ 4 ║ 3 │ 5 │ 8 ║ 9 │ 1 │ 6 ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║ 1 │ 6 │ 8 ║ 9 │ 2 │ 4 ║ 7 │ 3 │ 5 ║
	//╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣
	//║ 6 │ 4 │ 1 ║ 8 │ 7 │ 3 ║ 5 │ 9 │ 2 ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║ 3 │ 7 │ 2 ║ 6 │ 9 │ 5 ║ 4 │ 8 │ 1 ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║ 8 │ 5 │ 9 ║ 4 │ 1 │ 2 ║ 6 │ 7 │ 3 ║
	//╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣
	//║ 2 │ 3 │ 6 ║ 7 │ 4 │ 9 ║ 1 │ 5 │ 8 ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║ 9 │ 1 │ 5 ║ 2 │ 8 │ 6 ║ 3 │ 4 │ 7 ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║ 4 │ 8 │ 7 ║ 5 │ 3 │ 1 ║ 2 │ 6 │ 9 ║
	//╚═══╧═══╧═══╩═══╧═══╧═══╩═══╧═══╧═══╝

	g := initgrid()

	//g[1][2][0] = 9
	//g[1][7][0] = 8
	//g[1][9][0] = 4
	//
	//g[2][1][0] = 7
	//g[2][5][0] = 5
	//g[2][6][0] = 8
	//
	//g[3][2][0] = 6
	//g[3][3][0] = 8
	//g[3][4][0] = 9
	//
	//g[4][1][0] = 6
	//g[4][2][0] = 4
	//g[4][3][0] = 1
	//g[4][5][0] = 7
	//
	//g[5][3][0] = 2
	//g[5][7][0] = 4
	//
	//g[6][5][0] = 1
	//g[6][7][0] = 6
	//g[6][8][0] = 7
	//g[6][9][0] = 3
	//
	//g[7][6][0] = 9
	//g[7][7][0] = 1
	//g[7][8][0] = 5
	//
	//g[8][4][0] = 2
	//g[8][5][0] = 8
	//g[8][9][0] = 7
	//
	//g[9][1][0] = 4
	//g[9][3][0] = 7
	//g[9][8][0] = 6

	//╔═══╤═══╤═══╦═══╤═══╤═══╦═══╤═══╤═══╗
	//║   │ 3 │   ║ 6 │   │ 2 ║   │   │ 5 ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │   │   ║   │ 4 │ 8 ║   │   │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │   │ 9 ║   │   │ 7 ║ 8 │   │   ║
	//╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣
	//║   │   │ 8 ║   │   │ 4 ║   │ 3 │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │ 5 │ 3 ║   │   │   ║ 4 │ 6 │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │ 7 │   ║ 8 │   │   ║ 9 │   │   ║
	//╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣
	//║   │   │ 1 ║ 4 │   │   ║ 2 │   │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │   │   ║ 1 │ 8 │   ║   │   │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║ 3 │   │   ║ 7 │   │ 6 ║   │ 1 │   ║
	//╚═══╧═══╧═══╩═══╧═══╧═══╩═══╧═══╧═══╝

	//g[9][1][0] = 3
	//g[9][4][0] = 7
	//g[9][6][0] = 6
	//g[9][8][0] = 1
	//
	//g[8][4][0] = 1
	//g[8][5][0] = 8
	//
	//g[7][3][0] = 1
	//g[7][4][0] = 4
	//g[7][7][0] = 2
	//
	//g[6][2][0] = 7
	//g[6][4][0] = 8
	//g[6][7][0] = 9
	//
	//g[5][2][0] = 5
	//g[5][3][0] = 3
	//g[5][7][0] = 4
	//g[5][8][0] = 6
	//
	//g[4][3][0] = 8
	//g[4][6][0] = 4
	//g[4][8][0] = 3
	//
	//g[3][3][0] = 9
	//g[3][6][0] = 7
	//g[3][7][0] = 8
	//
	//g[2][5][0] = 4
	//g[2][6][0] = 8
	//
	//g[1][2][0] = 3
	//g[1][4][0] = 6
	//g[1][6][0] = 2
	//g[1][9][0] = 5

	//
	//╔═══╤═══╤═══╦═══╤═══╤═══╦═══╤═══╤═══╗
	//║   │   │ 5 ║ 2 │ 8 │   ║   │   │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │   │   ║   │   │ 4 ║ 1 │   │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │   │ 9 ║   │   │   ║ 4 │   │ 3 ║
	//╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣
	//║ 9 │   │   ║ 7 │   │   ║   │ 6 │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │ 8 │   ║   │ 1 │   ║   │ 4 │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │ 5 │   ║   │   │ 9 ║   │   │ 1 ║
	//╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣
	//║ 4 │   │ 6 ║   │   │   ║ 2 │   │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │   │ 7 ║ 4 │   │   ║   │   │   ║
	//╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
	//║   │   │   ║   │ 2 │ 5 ║ 6 │   │   ║
	//╚═══╧═══╧═══╩═══╧═══╧═══╩═══╧═══╧═══╝

	g[1][3][0] = 5
	g[1][4][0] = 2
	g[1][5][0] = 8

	g[2][6][0] = 4
	g[2][7][0] = 1

	g[3][3][0] = 9
	g[3][7][0] = 4
	g[3][9][0] = 3

	g[4][1][0] = 9
	g[4][4][0] = 7
	g[4][8][0] = 6

	g[5][2][0] = 8
	g[5][5][0] = 1
	g[5][8][0] = 4

	g[6][8][0] = 5
	g[6][6][0] = 9
	g[6][9][0] = 1

	g[7][1][0] = 4
	g[7][3][0] = 6
	g[7][7][0] = 2

	g[8][3][0] = 7
	g[8][4][0] = 4

	g[9][5][0] = 2
	g[9][6][0] = 5
	g[9][7][0] = 6

	removeFilledValues(&g)
	return g

}

//returns the 2 int to be used in a loop
// for 3x3 cell evaluation
func getcelldim(i int) (int, int) {
	switch i {
	case 1, 2, 3:
		{
			return 1, 4
		}
	case 4, 5, 6:
		{
			return 4, 7
		}
	case 7, 8, 9:
		{
			return 7, 10
		}
	}

	return 0, 0
}

func removeFilledOptions(g *grid, i, j int) {
	//first clear all the possible values in same cell
	for k := 1; k < 10; k++ {
		g[i][j][k] = 0
	}

	//now remove the possible values from cells in Row, col & mini-grid
	val := g[i][j][0]

	//remove from row first
	for k := 1; k < 10; k++ {
		g[k][j][val] = 0
	}

	//remove from col
	for k := 1; k < 10; k++ {
		g[i][k][val] = 0
	}

	//now from mini-gird
	i1, i2 := getcelldim(i)
	j1, j2 := getcelldim(j)
	for i3 := i1; i3 < i2; i3++ {
		for j3 := j1; j3 < j2; j3++ {
			g[i3][j3][val] = 0
		}
	}

}

func removeFilledValues(g *grid) {

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {

			if g[i][j][0] != 0 {
				removeFilledOptions(g, i, j)

			}
		}
	}
}

func placeSingletons(g *grid) {

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {

			if g[i][j][0] == 0 {

				var iszero bool
				var val int

				for k := 1; k < 10; k++ {
					if g[i][j][k] != 0 {
						if val != 0 {
							iszero = false
							break
						}
						iszero = true
						val = k
					}
				}

				if iszero {
					g[i][j][0] = val
					removeFilledOptions(g, i, j)
				}

			}

		}
	}
}

func countFilled(g *grid) int {
	var c int

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {

			if g[i][j][0] != 0 {
				c++
			}
		}
	}

	return c
}

func isValid(g *grid) bool {

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if g[i][j][0] != 0 {
				val := g[i][j][0]
				for l := 1; l < 10; l++ {
					if l == i {
						continue
					}
					if g[l][j][0] == val {
						fmt.Printf("Invalid at (row check) %d, %d\n", i, j)
						return false
					}
				}

				for l := 1; l < 10; l++ {
					if l == j {
						continue
					}
					if g[i][l][0] == val {
						fmt.Printf("Invalid at (col check) %d, %d\n", i, j)
						return false
					}
				}

				i1, i2 := getcelldim(i)
				j1, j2 := getcelldim(j)
				for l := i1; l < i2; l++ {
					for m := j1; m < j2; m++ {
						if l == i && m == j {
							continue
						}
						if g[l][m][0] == val {
							fmt.Printf("Invalid at (cell check) %d, %d\n", l, m)
							return false
						}
					}
				}
			}

		}
	}

	//check if [0] is != 0 but all the elemnents are 0
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if g[i][j][0] == 0 {
				var isNonZero bool

				for k := 1; k < 10; k++ {
					if g[i][j][k] != 0 {
						isNonZero = true
					}
				}

				if !isNonZero {
					fmt.Printf("all options are zero at %d,%d\n", i, j)
					return false
				}
			}
		}

	}

	return true
}

func placeSingleOptions(g *grid) {

	//check rows first
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {

			if g[i][j][0] == 0 {
				for k := 1; k < 10; k++ {
					if g[i][j][k] != 0 {

						//check rows first
						canFill := true
						for i1 := 1; i1 < 10; i1++ {
							if i1 == i {
								continue
							}

							if g[i1][j][k] != 0 {
								canFill = false
								break
							}
						}

						if canFill {
							g[i][j][0] = k
							removeFilledOptions(g, i, j)
							break
						}

						//check cols
						canFill = true
						for j1 := 1; j1 < 10; j1++ {
							if j1 == j {
								continue
							}

							if g[i][j1][k] != 0 {
								canFill = false
								break
							}
						}

						if canFill {
							g[i][j][0] = k
							removeFilledOptions(g, i, j)
							break
						}

						//check mini-grid
						canFill = true
						i1, i2 := getcelldim(i)
						j1, j2 := getcelldim(j)
						for i3 := i1; i3 < i2; i3++ {
							for j3 := j1; j3 < j2; j3++ {
								if i3 == i && j3 == j {
									continue
								}

								if g[i3][j3][k] != 0 {
									canFill = false
									break
								}
							}
						}

						if canFill {
							g[i][j][0] = k
							removeFilledOptions(g, i, j)
							break
						}
					}
				}
			}
		}
	}
}
