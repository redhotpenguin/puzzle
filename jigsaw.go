package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type point struct {
	x float32
	y float32
}

var size float32 = 140

// offset which would be the jigsaw cutting width so pieces fit together
var offset float32 = 0.5

// offset multiplier for the tip of the curve
var curveOffset float32 = 2.0

// magic numbers which control the shapes of the innies/outies
var halfway float32 = 0.5
var curveOne float32 = 0.2            // 0 to 0.5
var curveTwo float32 = 0.1            // 0 to 0.1
var curveThree float32 = curveOne * 2 // 0.4
var curveFour float32 = curveTwo * 3  // 0.3

var alignmentDebug bool = true

func main() {

	var width float32 = size
	var height float32 = size

	var curves = [4][3][3]point{}

	offsetSize := size
	var offsetCurves = [4][3][3]point{}
	var offsetWidth float32 = offsetSize
	var offsetHeight float32 = offsetSize

	// no offset with debug alignment
	var alignmentOffset float32 = 0
	if alignmentDebug {
		offset = 0
		alignmentOffset = 1 // adjust this, but curve joining gets choppy above 1
	}

	// outie or innie?? random orientation for each side
	// top
	var outie float32 = flip() // 1.0 is outie, -1.0 is innie
	height = bumpDimension(height, outie, size)
	curves = setTopSide(curves, outie, size, offset)

	// create a second piece offset from the first
	if alignmentDebug {
		offsetHeight = bumpDimension(offsetHeight, outie, offsetSize)
		offsetCurves = setTopSide(offsetCurves, outie, offsetSize, alignmentOffset)
	}

	// right
	outie = flip()
	width = bumpDimension(width, outie, size)
	curves = setRightSide(curves, outie, size, offset)

	if alignmentDebug {
		offsetWidth = bumpDimension(offsetWidth, outie, offsetSize)
		offsetCurves = setRightSide(offsetCurves, outie, offsetSize, alignmentOffset)
	}

	// bottom
	outie = flip()
	height = bumpDimension(height, outie, size)
	curves = setBottomSide(curves, outie, size, offset)

	if alignmentDebug {
		offsetHeight = bumpDimension(offsetHeight, outie, offsetSize)
		offsetCurves = setBottomSide(offsetCurves, outie, offsetSize, alignmentOffset)
	}

	// left
	outie = flip()
	width = bumpDimension(width, outie, size)
	curves = setLeftSide(curves, outie, size, offset)

	if alignmentDebug {
		offsetWidth = bumpDimension(offsetWidth, outie, offsetSize)
		offsetCurves = setLeftSide(offsetCurves, outie, offsetSize, alignmentOffset)
	}

	fmt.Printf("<!-- generated with jigsaw.go -->\n<svg xmlns=\"http://www.w3.org/2000/svg\" version=\"1.0\" width=\"%1.1fmm\" height=\"%1.1fmm\" viewBox=\"-%1.1f -%1.1f %1.1f %1.1f\">\n", size+100, size+100, size/3, size/3, size+100, size+100)
	fmt.Printf("<!-- width %1.1f, height %1.1f -->\n", width, height)
	fmt.Printf("<!-- offset %1.1f, curveOffset %1.1f -->\n", offset, curveOffset)

	// format
	strCurves := formatCurves(curves, width, height, point{0, 0})
	fmt.Printf("%s", strCurves)

	if alignmentDebug {
		// debug curve
		offsetStrCurves := formatCurves(offsetCurves, offsetWidth, offsetHeight, point{0, 0})
		fmt.Printf("\n\n%s", offsetStrCurves)

		// alignment lines

		fmt.Printf("<line fill=\"none\" stroke=\"orange\" stroke-width=\"0.25\" x1=\"%1.1f\" y1=\"0\" x2=\"%1.1f\" y2=\"%1.1f\"></line>\n", size/2-10, size/2-10, size)
		fmt.Printf("<line fill=\"none\" stroke=\"orange\" stroke-width=\"0.25\" x1=\"%1.1f\" y1=\"0\" x2=\"%1.1f\" y2=\"%1.1f\"></line>\n", size/2-10-2, size/2-10-2, size)

		fmt.Printf("<line fill=\"none\" stroke=\"purple\" stroke-width=\"0.25\" x1=\"%1.1f\" y1=\"0\" x2=\"%1.1f\" y2=\"%1.1f\"></line>\n", size/2+10, size/2+10, size)
		fmt.Printf("<line fill=\"none\" stroke=\"purple\" stroke-width=\"0.25\" x1=\"%1.1f\" y1=\"0\" x2=\"%1.1f\" y2=\"%1.1f\"></line>\n", size/2+10+2, size/2+10+2, size)

		fmt.Printf("<line fill=\"none\" stroke=\"blue\" stroke-width=\"0.25\" x1=\"0\" y1=\"%1.1f\" x2=\"%1.1f\" y2=\"%1.1f\"></line>\n", size/2+10+2, size, size/2+10+2)
		fmt.Printf("<line fill=\"none\" stroke=\"blue\" stroke-width=\"0.25\" x1=\"0\" y1=\"%1.1f\" x2=\"%1.1f\" y2=\"%1.1f\"></line>\n", size/2+10, size, size/2+10)

		fmt.Printf("<line fill=\"none\" stroke=\"green\" stroke-width=\"0.25\" x1=\"0\" y1=\"%1.1f\" x2=\"%1.1f\" y2=\"%1.1f\"></line>\n", size/2-10, size, size/2-10)
		fmt.Printf("<line fill=\"none\" stroke=\"green\" stroke-width=\"0.25\" x1=\"0\" y1=\"%1.1f\" x2=\"%1.1f\" y2=\"%1.1f\"></line>\n", size/2-10-2, size, size/2-10-2)

		// diagonal cross
		fmt.Printf("<line fill=\"none\" stroke=\"Red\" stroke-width=\"0.4\" x1=\"0\" y1=\"0\" x2=\"%1.1f\" y2=\"%1.1f\"></line>\n", size, size)
		fmt.Printf("<line fill=\"none\" stroke=\"Red\" stroke-width=\"0.4\" x1=\"%1.1f\" y1=\"0\" x2=\"0\" y2=\"%1.1f\"></line>\n", size, size)

		// vertical/horizontal cross
		fmt.Printf("<line fill=\"none\" stroke=\"Red\" stroke-width=\"0.4\" x1=\"%1.1f\" y1=\"-\" x2=\"%1.1f\" y2=\"%1.1f\"></line>\n", size/2, size/2, size)
		fmt.Printf("<line fill=\"none\" stroke=\"Red\" stroke-width=\"0.4\" x1=\"0\" y1=\"%1.1f\" x2=\"%1.1f\" y2=\"%1.1f\"></line>\n", size/2, size, size/2)

		fmt.Printf("<line fill=\"none\" stroke=\"Red\" stroke-width=\"0.4\" x1=\"0\" y1=\"0\" x2=\"%1.1f\" y2=\"0\"></line>\n", size)
		fmt.Printf("<line fill=\"none\" stroke=\"Red\" stroke-width=\"0.4\" x1=\"0\" y1=\"2\" x2=\"%1.1f\" y2=\"2\"></line>\n", size)
	}
	var svgFooter string = "</svg>"
	fmt.Println(svgFooter)
}

func bumpDimension(dimension float32, outie float32, size float32) float32 {

	if outie == 1.0 {
		dimension += size * 0.25
	} else {
		dimension += size * 0.025
	}
	return dimension
}

func flip() float32 {

	return -1.0
	if rand.Intn(2) == 0 {
		return 1.0
	} else {
		return -1.0
	}
}

func setTopSide(curves [4][3][3]point, outie float32, size float32, offset float32) [4][3][3]point {

	//outie = 1.0
	// top side. Every 3 lines is one C SVG declaration
	curves[0][0][0] = point{size * curveOne, 0}
	curves[0][0][1] = point{size * halfway, size * outie * curveTwo}
	curves[0][0][2] = point{size*curveThree + offset*outie, size * outie * curveTwo * -1}

	curves[0][1][0] = point{size*curveFour + offset*outie, size * outie * curveFour * -1}
	curves[0][1][1] = point{size * (1 - curveFour), size*outie*curveFour*-1 + offset*curveOffset}
	curves[0][1][2] = point{size*(1-curveThree) - offset*outie, size * outie * curveTwo * -1}

	curves[0][2][0] = point{size*halfway - offset*outie, size * outie * curveTwo}
	curves[0][2][1] = point{size * (1 - curveOne), 0}
	curves[0][2][2] = point{size, 0}

	return curves
}

func setRightSide(curves [4][3][3]point, outie float32, size float32, offset float32) [4][3][3]point {

	//outie = 1.0
	// right side
	curves[1][0][0] = point{size, size * curveOne}
	curves[1][0][1] = point{size * (1 - (outie * curveTwo)), size * halfway}
	curves[1][0][2] = point{size * (1 + (outie * curveTwo)), size*2*curveOne + offset*outie}

	curves[1][1][0] = point{size * (1 + (outie * curveFour)), size*curveFour + offset*outie}
	curves[1][1][1] = point{size*(1+(outie*curveFour)) - offset*curveOffset, size * (1 - curveFour)}
	curves[1][1][2] = point{size * (1 + (outie * curveTwo)), size*(1-curveThree) - offset*outie}

	curves[1][2][0] = point{size * (1 - (outie * curveTwo)), size*halfway - offset*outie}
	curves[1][2][1] = point{size, size * (1 - curveOne)}
	curves[1][2][2] = point{size, size}

	return curves
}

func setBottomSide(curves [4][3][3]point, outie float32, size float32, offset float32) [4][3][3]point {
	//	outie = -1.0
	// bottom side
	curves[2][0][0] = point{size * (1 - curveOne), size}
	curves[2][0][1] = point{size * halfway, size * (1 - (outie * curveTwo))}
	curves[2][0][2] = point{size*(1-curveThree) - offset*outie, size * (1 + (outie * curveTwo))}

	curves[2][1][0] = point{size*(1-curveFour) - offset*outie, size * (1 + (outie * curveFour))}
	curves[2][1][1] = point{size * curveFour, size*(1+(outie*curveFour)) - offset*curveOffset}
	curves[2][1][2] = point{size*curveThree + offset*outie, size * (1 + (outie * curveTwo))}

	curves[2][2][0] = point{size*halfway + offset*outie, size * (1 - (outie * curveTwo))}
	curves[2][2][1] = point{size * (1 * curveOne), size}
	curves[2][2][2] = point{0, size}

	return curves
}

func setLeftSide(curves [4][3][3]point, outie float32, size float32, offset float32) [4][3][3]point {

	//	outie = -1.0
	// left side
	curves[3][0][0] = point{0, size * (1 - curveOne)}
	curves[3][0][1] = point{size * outie * curveTwo, size * halfway}
	curves[3][0][2] = point{size * (-1 * outie * curveTwo), size*(1-curveThree) - offset*outie}

	curves[3][1][0] = point{size * (-1 * outie * curveFour), size*(1-curveFour) - offset*outie}
	curves[3][1][1] = point{size*(-1*outie*curveFour) + offset*curveOffset, size * curveFour}
	curves[3][1][2] = point{size * (-1 * outie * curveTwo), size*curveThree + offset*outie}

	curves[3][2][0] = point{size * outie * curveTwo, size*halfway + offset*outie}
	curves[3][2][1] = point{0, size * curveOne}
	curves[3][2][2] = point{0, 0}

	return curves
}

func formatCurves(curves [4][3][3]point, width float32, height float32, start point) string {

	var pathElemStart string
	if alignmentDebug {
		pathElemStart = "<path fill=\"Blue\" style=\"fill-opacity: .05;\" stroke=\"Red\" stroke-width=\"0.1\" d=\""
	} else {
		pathElemStart = "<path fill=\"Blue\" style=\"fill-opacity: 1;\" stroke=\"Red\" stroke-width=\"0\" d=\""

	}
	var strCurve string = fmt.Sprintf("\tM %1.1f,%1.1f", start.x, start.y)

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			curveLine := fmt.Sprintf("\tC %0.1f,%0.1f %0.1f,%0.1f %0.1f,%0.1f",
				curves[i][j][0].x,
				curves[i][j][0].y,
				curves[i][j][1].x,
				curves[i][j][1].y,
				curves[i][j][2].x,
				curves[i][j][2].y)
			//		fmt.Println("curveline", curveLine)
			strCurve = strings.Join([]string{strCurve, curveLine}, "\n")
		}
		strCurve = strings.Join([]string{strCurve, "\n"}, "")
	}
	var pathElemEnd string = "\"></path>"

	// put it all together
	strCurve = strings.Join([]string{pathElemStart, strCurve, pathElemEnd}, "\n")

	return strCurve
}
