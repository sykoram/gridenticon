package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	svgo "github.com/ajstarks/svgo/float"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type grid [][]byte

const tileSize float64 = 10
const gridSize int = 8
const border float64 = 5
const bgColor string = "white"

// flags
var help bool
var str string
var defsFile string
var outFile string

var hasher = sha256.New()
var out *os.File

func init() {
	flag.BoolVar(&help, "h", false, "")
	flag.StringVar(&outFile, "out", "./out.svg", "Output file")
	flag.StringVar(&defsFile, "defs", "./default.defs", "Defs file")
}

func main() {
	flag.Parse()
	handleHelp()
	setup()
	h := getHash(str)
	g := bytesToGrid(h)
	generateIdenticon(g)
	out.Close()
	exit()
}

/*
Handles help flag -h. If the help is requested, prints program description and usage, and exits.
*/
func handleHelp() {
	if help {
		fmt.Println(`A SVG identicon generator!
This program generates an identicon for given string. For more advanced usage, see https://github.com/sykoram/identicon
Usage: ./identicon STRING
Additional flags:`)
		flag.PrintDefaults()
		os.Exit(0)
	}
}

/*
Checks and processes flags.
 */
func setup() {
	str = flag.Arg(0)

	var err error
	out, err = createFile(outFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
Creates and returns a file. Parent directories are created if required. File has to be closed manually!
*/
func createFile(path string) (f *os.File, err error) {
	err = os.MkdirAll(filepath.Dir(path), 666)
	if err != nil {
		return
	}
	f, err = os.Create(path)
	return
}

/*
Generates and returns hash of the given string.
os.Exit(1) on error
 */
func getHash(s string) []byte {
	reader := strings.NewReader(s)
	_, err := io.Copy(hasher, reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return hasher.Sum(nil)
}

/*
Takes []byte data, splits every byte into two parts, generates and returns a grid.
 */
func bytesToGrid(bytes []byte) [][]byte {
	// split bytes
	halfbytes := make([]byte, gridSize*gridSize)
	for i, b := range bytes {
		halfbytes[2*i] = (b & 0b11110000) >> 4
		halfbytes[2*i+1] = b & 0b00001111
	}

	// generate grid
	g := make(grid, gridSize)
	for i := 0; i < gridSize; i++ {
		g[i] = make([]byte, gridSize)
		for j := 0; j < gridSize; j++ {
			g[i][j] = halfbytes[i*gridSize+j]
		}
	}

	return g
}

/*
Generates SVG identicon from the given grid.
 */
func generateIdenticon(g grid) {
	if isGridEmpty(g) {
		fmt.Println("the grid is empty")
		os.Exit(1)
	}

	// create svg document
	w := float64(len(g))*tileSize + 2*border
	h := float64(len(g[0]))*tileSize + 2*border
	svg := svgo.New(out)
	svg.Start(w, h)

	addDefs(svg)

	// background
	svg.Rect(0, 0, w, h, "stroke:none;fill:"+bgColor)

	// tile grid
	for i := range g {
		for j := range g[i] {
			x := float64(j)*tileSize + border
			y := float64(i)*tileSize + border
			svg.Use(x, y, fmt.Sprintf("#%x", g[i][j]))
		}
	}

	svg.End()
}

/*
Returns true if the grid g has zero columns or the first column is empty; false otherwise.
 */
func isGridEmpty(g grid) bool {
	return len(g) == 0 || len(g[0]) == 0
}

/*
Creates a definition block and copies defs from an external file.
 */
func addDefs(svg *svgo.SVG) {
	svg.Def()
	defsReader, _ := os.Open(defsFile)
	_, err := io.Copy(svg.Writer, defsReader)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("error reading file %s\n", defsFile)
		os.Exit(1)
	}
	svg.DefEnd()
}

/*
Prints information about performed actions.
 */
func exit() {
	fmt.Printf("Generated an identicon for string \"%s\": %s\n", str, outFile)
	fmt.Println("See -h for help")
}