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

func flip() int {
	if rand.Intn(1) == 0 {
		return 1
	} else {
		return -1
	}
}

func main() {

	var curves = [4][3][3]point{}

	var size float32 = 100
	var offset float32 = 0
	var halfway float32 = 0.5
	var curveOne float32 = 0.2            // 0 to 0.5
	var curveTwo float32 = 0.1            // 0 to 0.1
	var curveThree float32 = curveOne * 2 // 0.4
	var curveFour float32 = curveTwo * 3  // 0.3

	// top side
	curves[0][0][0] = point{size*curveOne + offset, 0 + offset}
	curves[0][0][1] = point{size*halfway + offset, size*curveTwo + offset}
	curves[0][0][2] = point{size*curveThree + offset, size*curveTwo*-1 + offset}

	curves[0][1][0] = point{size*curveFour + offset, -1*size*curveFour - offset}
	curves[0][1][1] = point{size*(1-curveFour) + offset, -1*size*curveFour - offset}
	curves[0][1][2] = point{size*(1-curveThree) + offset, size*curveTwo*-1 + offset}

	curves[0][2][0] = point{size*halfway - offset, size*curveTwo + offset}
	curves[0][2][1] = point{size*(1-curveOne) - offset, 0 + offset}
	curves[0][2][2] = point{size - offset, 0 + offset}

	// right side
	curves[1][0][0] = point{size - offset, size*curveOne + offset}
	curves[1][0][1] = point{size*(1-curveTwo) - offset, size*halfway + offset}
	curves[1][0][2] = point{size*(1+curveTwo) - offset, size*2*curveOne + offset}

	curves[1][1][0] = point{size*(1+curveFour) - offset, size*curveFour + offset}
	curves[1][1][1] = point{size*(1+curveFour) - offset, size*(1-curveFour) + offset}
	curves[1][1][2] = point{size*(1+curveTwo) - offset, size*(1-curveThree) + offset}

	curves[1][2][0] = point{size*(1-curveTwo) - offset, size*halfway + offset}
	curves[1][2][1] = point{size - offset, size*(1-curveOne) + offset}
	curves[1][2][2] = point{size + offset, size + offset}

	// bottom side
	curves[2][0][0] = point{size*(1-curveOne) + offset, size + offset}
	curves[2][0][1] = point{size*halfway + offset, size*(1-curveTwo) + offset}
	curves[2][0][2] = point{size*(1-curveThree) + offset, size*(1+curveTwo) + offset}

	curves[2][1][0] = point{size*(1-curveFour) + offset, size*(1+curveFour) + offset}
	curves[2][1][1] = point{size*curveFour + offset, size*(1+curveFour) + offset}
	curves[2][1][2] = point{size*curveThree + offset, size*(1+curveTwo) + offset}

	curves[2][2][0] = point{size*halfway + offset, size*(1-curveTwo) + offset}
	curves[2][2][1] = point{size*(1*curveOne) + offset, size + offset}
	curves[2][2][2] = point{0 + offset, size + offset}

	// left side
	curves[3][0][0] = point{0 + offset, size*(1-curveOne) + offset}
	curves[3][0][1] = point{size*curveTwo + offset, size*halfway + offset}
	curves[3][0][2] = point{size*(-1*curveTwo) + offset, size*(1-curveThree) + offset}

	curves[3][1][0] = point{size*(-1*curveFour) + offset, size*(1-curveFour) + offset}
	curves[3][1][1] = point{size*(-1*curveFour) + offset, size*curveFour + offset}
	curves[3][1][2] = point{size*(-1*curveTwo) + offset, size*curveThree + offset}

	curves[3][2][0] = point{size*curveTwo + offset, size*halfway + offset}
	curves[3][2][1] = point{0 + offset, size*curveOne + offset}
	curves[3][2][2] = point{0 + offset, 0 + offset}

	strCurves := formatSvg(curves)

	fmt.Println("string", strCurves)
}

func formatSvg(curves [4][3][3]point) string {

	var strCurve string = "M 0,0 "

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			curveLine := fmt.Sprintf("C %0.1f,%0.1f %0.1f,%0.1f %0.1f,%0.1f",
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
	return strCurve
}
