// Command mark writes the go-sparse mark, svg/color/go-sparse.svg:
//
//	cd mark && go run . > ../svg/color/go-sparse.svg
//
// The mark is a sparse matrix: a pair of brackets around a grid of dots, one
// dot at every position where the matrix holds a nonzero. The pattern is a
// band (the main diagonal and the one on each side of it), the shape a
// discretised one-dimensional operator takes and the one a banded solver
// exploits; every other cell is empty, and the empty cells are the point.
//
// Nothing is placed by hand. The pattern below says which cells are nonzero,
// and the geometry follows: the cells are spaced evenly and centred on the
// square, and the brackets stand off the grid by the same clear gap on every
// side. Change n or the bandwidth and the mark redraws itself.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient, a white line glyph
// at stroke-width 13 with round caps and joins (a zero-length line is a dot).
// logos/brandgen renders the format matrix from the file this writes; the kit
// workflow runs this again and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#E88A7A" // the gradient's light stop
	dark  = "#A8372C" // its dark stop

	stroke = 13.0 // the family's line width

	n         = 5    // the matrix is n by n
	bandwidth = 1    // a cell (i, j) is nonzero when |i-j| <= bandwidth
	pitch     = 30.0 // distance between the centres of two cells
	arm       = 20.0 // how far a bracket's arms reach toward the grid
	margin    = 22.0 // from the outermost cell centre to the bracket's spine
)

// num prints a coordinate with at most one decimal and no trailing zero.
func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	half := pitch * float64(n-1) / 2
	cell := func(k int) float64 { return 128 - half + pitch*float64(k) }

	var paths []string
	// The brackets: a spine and two arms, one on each side of the grid.
	top, bottom := 128-half-margin, 128+half+margin
	for _, side := range []float64{-1, 1} {
		spine := 128 + side*(half+margin)
		tip := spine - side*arm
		paths = append(paths, fmt.Sprintf(`<path d="M%s %sH%sV%sH%s"/>`,
			num(tip), num(top), num(spine), num(bottom), num(tip)))
	}
	// The nonzeros: one dot per cell inside the band, row by row.
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if abs(i-j) <= bandwidth {
				paths = append(paths, fmt.Sprintf(`<path d="M%s %sh0"/>`, num(cell(j)), num(cell(i))))
			}
		}
	}

	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-sparse">
  <defs><linearGradient id="spgrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#spgrad)"/>
  <g fill="none" stroke="#FFFFFF" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">
    %s
  </g>
</svg>
`, light, dark, num(stroke), strings.Join(paths, "\n    "))
}
