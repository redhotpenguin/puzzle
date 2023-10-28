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

func flip() float32 {
	if rand.Intn(2) == 0 {
		return 1.0
	} else {
		return -1.0
	}
}

func main() {

	var curves = [4][3][3]point{}

	var size float32 = 100
	//	var offset float32 = 0
	var halfway float32 = 0.5
	var curveOne float32 = 0.2            // 0 to 0.5
	var curveTwo float32 = 0.1            // 0 to 0.1
	var curveThree float32 = curveOne * 2 // 0.4
	var curveFour float32 = curveTwo * 3  // 0.3

	// outie or innie??
	var outie float32 = flip() // 1.0 is outie, -1.0 is innie

	fmt.Printf("<!-- Outie %1.0f -->\n", outie)

	// top side. Every 3 lines is one C SVG declaration
	curves[0][0][0] = point{size * curveOne, 0}
	curves[0][0][1] = point{size * halfway, size * outie * curveTwo}
	curves[0][0][2] = point{size * curveThree, size * outie * curveTwo * -1}

	curves[0][1][0] = point{size * curveFour, size * outie * curveFour * -1}
	curves[0][1][1] = point{size * (1 - curveFour), size * outie * curveFour * -1}
	curves[0][1][2] = point{size * (1 - curveThree), size * outie * curveTwo * -1}

	curves[0][2][0] = point{size * halfway, size * outie * curveTwo}
	curves[0][2][1] = point{size * (1 - curveOne), 0}
	curves[0][2][2] = point{size, 0}

	// right side
	curves[1][0][0] = point{size, size * curveOne}
	curves[1][0][1] = point{size * (1 - (outie * curveTwo)), size * halfway}
	curves[1][0][2] = point{size * (1 + (outie * curveTwo)), size * 2 * curveOne}

	curves[1][1][0] = point{size * (1 + (outie * curveFour)), size * curveFour}
	curves[1][1][1] = point{size * (1 + (outie * curveFour)), size * (1 - curveFour)}
	curves[1][1][2] = point{size * (1 + (outie * curveTwo)), size * (1 - curveThree)}

	curves[1][2][0] = point{size * (1 - (outie * curveTwo)), size * halfway}
	curves[1][2][1] = point{size, size * (1 - curveOne)}
	curves[1][2][2] = point{size, size}

	// bottom side
	curves[2][0][0] = point{size * (1 - curveOne), size}
	curves[2][0][1] = point{size * halfway, size * (1 - (outie * curveTwo))}
	curves[2][0][2] = point{size * (1 - curveThree), size * (1 + (outie * curveTwo))}

	curves[2][1][0] = point{size * (1 - curveFour), size * (1 + (outie * curveFour))}
	curves[2][1][1] = point{size * curveFour, size * (1 + (outie * curveFour))}
	curves[2][1][2] = point{size * curveThree, size * (1 + (outie * curveTwo))}

	curves[2][2][0] = point{size * halfway, size * (1 - (outie * curveTwo))}
	curves[2][2][1] = point{size * (1 * curveOne), size}
	curves[2][2][2] = point{0, size}

	// left side
	curves[3][0][0] = point{0, size * (1 - curveOne)}
	curves[3][0][1] = point{size * outie * curveTwo, size * halfway}
	curves[3][0][2] = point{size * (-1 * outie * curveTwo), size * (1 - curveThree)}

	curves[3][1][0] = point{size * (-1 * outie * curveFour), size * (1 - curveFour)}
	curves[3][1][1] = point{size * (-1 * outie * curveFour), size * curveFour}
	curves[3][1][2] = point{size * (-1 * outie * curveTwo), size * curveThree}

	curves[3][2][0] = point{size * outie * curveTwo, size * halfway}
	curves[3][2][1] = point{0, size * curveOne}
	curves[3][2][2] = point{0, 0}

	strCurves := formatSvg(curves)

	fmt.Printf("%s", strCurves)
}

func formatSvg(curves [4][3][3]point) string {

	var svgHeader string = "<!-- generated with jigsaw.go -->\n<svg xmlns=\"http://www.w3.org/2000/svg\" version=\"1.0\" width=\"200mm\" height=\"200mm\" viewBox=\"-30 -30 200 200\">"

	var pathElemStart string = "<path fill=\"Blue\" stroke=\"Red\" stroke-width=\"0\" d=\""
	var strCurve string = "\tM 0,0 "

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
	var svgFooter string = "</svg>"

	// put it all together
	strCurve = strings.Join([]string{svgHeader, pathElemStart, strCurve, pathElemEnd, svgFooter}, "\n")

	return strCurve
}
