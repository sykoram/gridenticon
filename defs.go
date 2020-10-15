package main

import (
	"fmt"
	svgo "github.com/ajstarks/svgo/float"
)

var defFun = map[byte]func(*svgo.SVG,byte) {
	0x0: circle,
	0x1: circle,
	0x2: circle,
	0x3: circle,
	0x4: circle,
	0x5: circle,
	0x6: circle,
	0x7: circle,
	0x8: circleSmall,
	0x9: circleSmall,
	0xA: circleSmall,
	0xB: circleSmall,
	0xC: circleSmall,
	0xD: circleSmall,
	0xE: circleSmall,
	0xF: circleSmall,
}

func circle(svg *svgo.SVG, id byte) {
	svg.Circle(5, 5, 4, fmt.Sprintf("id='%x'", id))
}

func circleSmall(svg *svgo.SVG, id byte) {
	svg.Circle(5, 5, 2, fmt.Sprintf("id='%x'", id))
}

//svg.Circle(50, 50, 10, "stroke:red;fill:blue")
//svg.Line(25, 5, 75, 55, "stroke:black;stroke-width:10;stroke-linecap:round")
//svg.Arc(20, 80, 20, 20, 0, false, true, 60, 80, "stroke:black;stroke-width:1;fill:none;stroke-linecap:round")
//svg.Arc(10, 50, 20, 20, 0, false, true, 30, 30, "stroke:black;stroke-width:1;fill:none;stroke-linecap:round")