package main

import (
	"fmt"
	"math/rand"
)

var RotationCoords = map[byte][]Coordinate{
	'I': {
		{X: -1, Y: -1},
		{X: 2, Y: 1},
		{X: -2, Y: -2},
		{X: 1, Y: 2},
	},
	'O': {
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
	},
	'6': { // All tetriminos with 6 cells (T, S, Z, J, L)
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: -1, Y: -1},
		{X: 0, Y: 1},
	},
}

var tetriminos = []Tetrimino{
	{
		Value: 'I',
		Cells: [][]bool{
			{true, true, true, true},
		},
		Pos:            Coordinate{X: 3, Y: -1},
		RotationCoords: RotationCoords['I'],
	},
	{
		Value: 'O',
		Cells: [][]bool{
			{true, true},
			{true, true},
		},
		Pos:            Coordinate{X: 4, Y: -2},
		RotationCoords: RotationCoords['O'],
	},
	{
		Value: 'T',
		Cells: [][]bool{
			{false, true, false},
			{true, true, true},
		},
		Pos:            Coordinate{X: 3, Y: -2},
		RotationCoords: RotationCoords['6'],
	},
	{
		Value: 'S',
		Cells: [][]bool{
			{false, true, true},
			{true, true, false},
		},
		Pos:            Coordinate{X: 3, Y: -2},
		RotationCoords: RotationCoords['6'],
	},
	{
		Value: 'Z',
		Cells: [][]bool{
			{true, true, false},
			{false, true, true},
		},
		Pos:            Coordinate{X: 3, Y: -2},
		RotationCoords: RotationCoords['6'],
	},
	{
		Value: 'J',
		Cells: [][]bool{
			{true, false, false},
			{true, true, true},
		},
		Pos:            Coordinate{X: 3, Y: -2},
		RotationCoords: RotationCoords['6'],
	},
	{
		Value: 'L',
		Cells: [][]bool{
			{false, false, true},
			{true, true, true},
		},
		Pos:            Coordinate{X: 3, Y: -2},
		RotationCoords: RotationCoords['6'],
	},
}

type Coordinate struct {
	X, Y int
}

type Tetrimino struct {
	Value           byte
	Cells           [][]bool
	Pos             Coordinate // the top left cell of the tetrimino
	CurrentRotation int
	RotationCoords  []Coordinate
}

// TODO: update to be private. Change tests to be on NewTetrimino instead.
// randomTetrimino returns a random tetrimino.
// The tetrimino will be positioned just above the playfield.
func RandomTetrimino(playfieldHeight int) *Tetrimino {
	t := tetriminos[rand.Intn(len(tetriminos))]
	t.Pos.Y += playfieldHeight - 20
	return &t
}

func (p *Playfield) NewTetrimino() *Tetrimino {
	tet := RandomTetrimino(len(p))
	p.addCells(tet)
	return tet
}

// MoveDown moves the tetrimino down one row.
// If the tetrimino cannot move down, it will be added to the playfield and a new tetrimino will be returned.
func (t *Tetrimino) MoveDown(playfield *Playfield) (*Tetrimino, error) {
	if !t.canMoveDown(*playfield) {
		return playfield.NewTetrimino(), nil
	}
	err := playfield.removeCells(t)
	if err != nil {
		return nil, fmt.Errorf("failed to remove cells: %w", err)
	}
	t.Pos.Y++
	err = playfield.addCells(t)
	if err != nil {
		return nil, fmt.Errorf("failed to add cells: %w", err)
	}
	return nil, nil
}

// MoveLeft moves the tetrimino left one column.
// If the tetrimino cannot move left, it will not move.
func (t *Tetrimino) MoveLeft(playfield *Playfield) error {
	if !t.canMoveLeft(*playfield) {
		return nil
	}
	err := playfield.removeCells(t)
	if err != nil {
		return fmt.Errorf("failed to remove cells: %w", err)
	}
	t.Pos.X--
	err = playfield.addCells(t)
	if err != nil {
		return fmt.Errorf("failed to add cells: %w", err)
	}
	return nil
}

// MoveRight moves the tetrimino right one column.
// If the tetrimino cannot move right, it will not move.
func (t *Tetrimino) MoveRight(playfield *Playfield) error {
	if !t.canMoveRight(*playfield) {
		return nil
	}
	err := playfield.removeCells(t)
	if err != nil {
		return fmt.Errorf("failed to remove cells: %w", err)
	}
	t.Pos.X++
	err = playfield.addCells(t)
	if err != nil {
		return fmt.Errorf("failed to add cells: %w", err)
	}
	return nil
}

func (t *Tetrimino) canMoveDown(playfield Playfield) bool {
	bottomRow := len(t.Cells) - 1
	for col := range t.Cells[bottomRow] {
		if t.Cells[bottomRow][col] {
			if bottomRow+t.Pos.Y+1 >= len(playfield) {
				return false
			}
			if cell := playfield[bottomRow+t.Pos.Y+1][col+t.Pos.X]; cell != 0 && cell != 'G' {
				return false
			}
		}
	}
	return true
}

func (t *Tetrimino) canMoveLeft(playfield Playfield) bool {
	leftCol := 0
	for row := range t.Cells {
		if t.Cells[row][leftCol] {
			if leftCol+t.Pos.X-1 < 0 {
				return false
			}
			if cell := playfield[row+t.Pos.Y][leftCol+t.Pos.X-1]; cell != 0 && cell != 'G' {
				return false
			}
		}
	}
	return true
}

func (t *Tetrimino) canMoveRight(playfield Playfield) bool {
	rightCol := len(t.Cells[0]) - 1
	for row := range t.Cells {
		if t.Cells[row][rightCol] {
			if rightCol+t.Pos.X+1 >= len(playfield[0]) {
				return false
			}
			if cell := playfield[row+t.Pos.Y][rightCol+t.Pos.X+1]; cell != 0 && cell != 'G' {
				return false
			}
		}
	}
	return true
}

func (p *Playfield) removeCells(tetrimino *Tetrimino) error {
	for row := range tetrimino.Cells {
		for col := range tetrimino.Cells[row] {
			if tetrimino.Cells[row][col] { // if the cell is true
				if v := p[row+tetrimino.Pos.Y][col+tetrimino.Pos.X]; v != tetrimino.Value { // if the cell is not the expected value
					return fmt.Errorf("cell at row %d, col %d is not the expected value", row+tetrimino.Pos.Y, col+tetrimino.Pos.X)
				}
				p[row+tetrimino.Pos.Y][col+tetrimino.Pos.X] = 0
			}
		}
	}
	return nil
}

func (p *Playfield) addCells(tetrimino *Tetrimino) error {
	for row := range tetrimino.Cells {
		for col := range tetrimino.Cells[row] {
			if tetrimino.Cells[row][col] {
				if cell := p[row+tetrimino.Pos.Y][col+tetrimino.Pos.X]; cell != 0 && cell != 'G' {
					return fmt.Errorf("cell at row %d, col %d is not empty or a ghost", row+tetrimino.Pos.Y, col+tetrimino.Pos.X)
				}
				p[row+tetrimino.Pos.Y][col+tetrimino.Pos.X] = tetrimino.Value
			}
		}
	}
	return nil
}

func (t *Tetrimino) Rotate(playfield *Playfield, clockwise bool) error {
	if t.Value == 'O' {
		return nil
	}

	var rotated *Tetrimino
	if clockwise {
		rotated = t.rotateClockwise()
	} else {
		rotated = t.rotateCounterClockwise()
	}
	err := playfield.removeCells(t)
	if err != nil {
		return fmt.Errorf("failed to remove cells: %w", err)
	}
	if rotated.canRotate(playfield, t) {
		t.Cells = rotated.Cells
	}
	err = playfield.addCells(t)
	if err != nil {
		return fmt.Errorf("failed to add cells: %w", err)
	}
	return nil
}

func (t Tetrimino) rotateClockwise() *Tetrimino {
	cellsCopy := make([][]bool, len(t.Cells))
	copy(cellsCopy, t.Cells)
	t.Cells = cellsCopy

	// reverse the order of the matrix rows
	for i, j := 0, len(t.Cells)-1; i < j; i, j = i+1, j-1 {
		t.Cells[i], t.Cells[j] = t.Cells[j], t.Cells[i]
	}
	t.transpose()

	t.CurrentRotation = (t.CurrentRotation + 1) % len(t.RotationCoords)
	t.Pos.X += t.RotationCoords[t.CurrentRotation].X
	t.Pos.Y += t.RotationCoords[t.CurrentRotation].Y

	return &t
}

func (t Tetrimino) rotateCounterClockwise() *Tetrimino {
	t.transpose()
	for row := range t.Cells {
		for i, j := 0, len(t.Cells[row])-1; i < j; i, j = i+1, j-1 {
			t.Cells[row][i], t.Cells[row][j] = t.Cells[row][j], t.Cells[row][i]
		}
	}
	return &t
}

func (t *Tetrimino) transpose() {
	xl := len(t.Cells[0])
	yl := len(t.Cells)
	result := make([][]bool, xl)
	for i := range result {
		result[i] = make([]bool, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = t.Cells[j][i]
		}
	}
	t.Cells = result
}

func (t *Tetrimino) canRotate(playfield *Playfield, original *Tetrimino) bool {
	for cellRow := range t.Cells {
		for cellCol := range t.Cells[cellRow] {
			if t.Cells[cellRow][cellCol] {
				if isOutOfBoundsHorizontally(cellRow, t.Pos.Y, playfield) {
					return false
				}
				if isOutOfBoundsVertically(cellCol, t.Pos.X, playfield) {
					return false
				}
				// if isCellOccupied(Coordinate{X: cellCol, Y: cellRow}, t.Pos, playfield, original) {
				// 	return false
				// }
			}
		}
	}
	return true
}

func isOutOfBoundsHorizontally(cellRow, tetRow int, playfield *Playfield) bool {
	cellRow += tetRow
	return cellRow < 0 || cellRow >= len(playfield)
}
func isOutOfBoundsVertically(cellCol, tetCol int, playfield *Playfield) bool {
	cellCol += tetCol
	return cellCol < 0 || cellCol >= len(playfield[0])
}

func isCellOccupied(cell, tetPos Coordinate, playfield *Playfield, originalTet *Tetrimino) bool {
	cellValue := playfield[cell.Y+tetPos.Y][cell.Y+tetPos.X]
	if cellValue == 0 {
		return false
	}
	if cellValue == 'G' {
		return false
	}

	// Is the cell within the bounds of the original tet?
	// Is the cell occupied by the original tet?

	cell.Y = tetPos.Y + cell.Y
	cell.X = tetPos.X + cell.X

	isInOriginalY := cell.Y >= originalTet.Pos.Y && cell.Y < originalTet.Pos.Y+len(originalTet.Cells)
	isInOriginalX := cell.X >= originalTet.Pos.X && cell.X < originalTet.Pos.X+len(originalTet.Cells[0])
	if !isInOriginalY || !isInOriginalX {
		return false
	}

	cell.Y -= originalTet.Pos.Y
	cell.X -= originalTet.Pos.X

	return originalTet.Cells[cell.Y][cell.X]
}

// # # # # #
// N # # # #
// N # # O #
// N B O O #
// # # # # #
// # # # # #
// # # # # #
