# go-sparse — brand assets

Official logos for the **go-sparse** organization.

<img src="avatar/go-sparse.png" width="160" alt="go-sparse">

The mark is a sparse matrix: a pair of brackets around a grid of dots, one dot
at every nonzero of a banded pattern (the main diagonal and the one on each
side of it). The empty cells are the point. `mark/main.go` writes it; nothing
is placed by hand. Its gradient (`#E88A7A` to `#A8372C`) is a brick red that
sets go-sparse apart from the violet of go-numerics, the emerald of go-ndarray
and go-fft, and go-plot's blue; its dark end reads on white at 6.47:1 and its
light end on a dark page at 7.68:1.

## Formats

- **SVG** — `svg/color`, `svg/white` (white, transparent), `svg/black` (black, transparent)
- **PNG** 16→1024 px, 3 variants — `png/<variant>/<size>/`
- **JPG** (color, white background) — `jpg/<size>/`
- **ICO** (Windows) — `ico/`
- **ICNS** (macOS) — `icns/`
- **GitHub avatar** 512 px — `avatar/`
- **Social preview** 1280×640 (repo banner) — `social/`

`check/kit.py` counts the whole matrix on every pull request, and the `mark`
job writes `svg/color/go-sparse.svg` again from `mark/main.go` and fails if a
byte differs.

---
*Generated from `svg/color/go-sparse.svg` by the fleet's Go brand pipeline (brandgen) — rounded square badge, line glyph, color / white / black variants.*
