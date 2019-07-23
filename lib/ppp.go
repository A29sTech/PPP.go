package a29

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
)

func M2P(mm float32) int {
	return int((mm / 25.4) * 300.0)
}

// Main  Object.
type Paper struct {
	Img    image.Image
	layers []image.Image
	names  []string
	noc    []int
}

// Constructer.
func (paper *Paper) New(width int, height int) {
	paper.Img = imaging.New(width, height, color.RGBA{255, 255, 255, 0xff})
}

//"Add Layers"
func (paper *Paper) Add(img image.Image, width float32, height float32, noc int, name string) int {
	if width < 1 && height < 1 {
		return 1
	}
	var bw = 2
	paper.names = append(paper.names, name)
	paper.noc = append(paper.noc, noc)
	//w := img.Bounds().Size().X
	//h := img.Bounds().Size().Y
	resized := imaging.Resize(img, int(M2P(width)), int(M2P(height)), imaging.Lanczos)
	bg := imaging.New(resized.Bounds().Size().X+(bw+bw), resized.Bounds().Size().Y+(bw+bw), color.RGBA{0, 0, 0, 0xff})
	border := imaging.Paste(bg, resized, image.Pt(bw, bw))
	paper.layers = append(paper.layers, border)
	return 0

}

//"Render Func"
func (paper *Paper) Render(x float32, w float32, s float32, r bool) image.Image {
	var copy int
	var pw int
	var img = paper.Img
	var cx = M2P(x)
	var cy = M2P(3.0) // Hard Coded.
	var mxh = 0
	px := M2P(x)
	ps := M2P(s)
	if w < 1 {
		pw = img.Bounds().Size().X - px - M2P(3.0) // Y Border.
	} else {
		pw = M2P(w)
	}
	for i, l := range paper.layers {
		//println(paper.names[i])
		copy = 0

		for paper.noc[i] > copy {
			//println(copy)

			lw := l.Bounds().Size().X
			lh := l.Bounds().Size().Y
			if r == true {
				if cx+lw < px+pw {
					img = imaging.Paste(img, l, image.Pt(cx, cy))
					cx += lw + ps
					if lh > mxh {
						mxh = lh
					}
				} else {
					cx = M2P(x)
					cy += mxh + ps
					img = imaging.Paste(img, l, image.Pt(cx, cy))
					cx += lw + ps
					mxh = lh
				}
			}

			copy++

		}

	}

	return img
}

// Remove ElementBy Name;
func (paper *Paper) Remove(name string) bool {
	for i, n := range paper.names {
		if n == name {
			if len(paper.names) < 2 {
				paper.layers = []image.Image{}
				paper.names = []string{}
				paper.noc = []int{}
				return true
			}
			paper.layers = append(paper.layers[:i], paper.layers[i+1:]...)
			paper.names = append(paper.names[:i], paper.names[i+1:]...)
			paper.noc = append(paper.noc[:i], paper.noc[i+1:]...)
			return true
		}
	}
	return false
}

// get Info
func (paper *Paper) PrintInfo() {
	if len(paper.names) == 0 {
		println("Empty Paper !")
	}
	for i, n := range paper.names {
		print(i+1, ": "+n+" #", paper.noc[i], "\n")
	}
}
