package main

import (
	"crypto/sha256"
	"fmt"
	svgo "github.com/ajstarks/svgo/float"
	"io"
	"os"
	"strings"
)

type grid [][]byte

const tileSize float64 = 10
const gridSize int = 8

var svgOut io.Writer = os.Stdout
var hasher = sha256.New()

func main() {
	reader := strings.NewReader("XXX")
	_, err := io.Copy(hasher, reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	h := hasher.Sum(nil)
	h2 := make([]byte, gridSize*gridSize)
	for i, b := range h {
		h2[2*i] = (b & 0b11110000) >> 4
		h2[2*i+1] = b & 0b00001111
	}

	g := make(grid, gridSize)
	for i := 0; i < gridSize; i++ {
		g[i] = make([]byte, gridSize)
		for j := 0; j < gridSize; j++ {
			g[i][j] = h2[i*gridSize+j]
		}
	}

	gridToSvg(g)
}

/*

 */
func gridToSvg(g grid) {
	if isGridEmpty(g) {
		fmt.Println("grid is empty")
		return
	}

	// create svg document
	w := float64(len(g))*tileSize
	h := float64(len(g[0]))*tileSize
	svg := svgo.New(svgOut)
	svg.Start(w, h)

	addDefs(svg)

	// tile grid
	for y := range g {
		for x := range g[y] {
			svg.Use(float64(x)*tileSize, float64(y)*tileSize, fmt.Sprintf("#%x", g[y][x]))
		}
	}

	//svg.Grid(0, 0, w, h, tileSize, "stroke:black;opacity:0.1") //XXX

	svg.End()
}

/*
Returns true if the grid g has zero columns or the first column is empty; false otherwise.
 */
func isGridEmpty(g grid) bool {
	return len(g) == 0 || len(g[0]) == 0
}

/*
Creates a definition block, and adds definitions using the defFun map.
 */
func addDefs(svg *svgo.SVG) {
	svg.Def()
	for i, f := range defFun {
		f(svg, i)
	}
	svg.DefEnd()
}