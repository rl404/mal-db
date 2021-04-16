package image

import (
	"image"
	"image/color"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

func (i *Image) getImageFromURL(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	res, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (i *Image) drawBackground(dc *gg.Context, url string) error {
	if url == "" {
		// Draw default background.
		i.drawDefaultBackground(dc)
		return nil
	}

	bg, err := i.getImageFromURL(url)
	if err != nil {
		return err
	}

	// Resize.
	bg = imaging.Fill(bg, i.width, i.height, imaging.Center, imaging.Lanczos)
	bg = imaging.Blur(bg, 10)

	dc.DrawImage(bg, 0, 0)
	return nil
}

func (i *Image) drawDefaultBackground(dc *gg.Context) {
	dc.SetColor(color.White)
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.Fill()
}

func (i *Image) drawOverlay(dc *gg.Context) {
	w := float64(dc.Width()) - (2.0 * i.margin)
	h := float64(dc.Height()) - (2.0 * i.margin)
	dc.SetColor(color.RGBA{0, 0, 0, 204})
	dc.DrawRectangle(i.margin, i.margin, w, h)
	dc.Fill()
}

func (i *Image) drawTitle(dc *gg.Context, title, enTitle string) {
	dc.SetColor(color.White)
	dc.SetFontFace(i.title.face)

	title = strings.Join(dc.WordWrap(title, i.title.maxWidth), "\n")

	// Resize font.
	fontSize := i.title.size
	w, h := dc.MeasureMultilineString(title, 1.25)
	for h > i.title.maxHeight {
		fontSize--
		f, _ := gg.LoadFontFace(i.fontPath, fontSize)
		dc.SetFontFace(f)
		w, h = dc.MeasureMultilineString(title, 1.25)
	}

	x := 2 * i.margin
	y := 2*i.margin + (i.title.maxHeight-h)/2.0
	dc.DrawStringWrapped(title, x, y, 0, 0, i.title.maxWidth, 1.25, gg.AlignCenter)

	// Line.
	dc.DrawLine(
		2*i.margin+(i.enTitle.maxWidth-w)/2,
		y+h+i.margin,
		2*i.margin+(i.enTitle.maxWidth-w)/2+w,
		y+h+i.margin,
	)
	dc.SetRGBA(1, 1, 1, 0.5)
	dc.Stroke()

	if enTitle == "" {
		return
	}

	// English title.
	dc.SetColor(color.RGBA{255, 255, 255, 127})
	dc.SetFontFace(i.enTitle.face)

	// Resize font.
	fontSize2 := i.enTitle.size
	w2, _ := dc.MeasureString(enTitle)
	for w2 > w {
		fontSize2--
		f, _ := gg.LoadFontFace(i.fontPath, fontSize2)
		dc.SetFontFace(f)
		w2, _ = dc.MeasureString(enTitle)
	}

	x2 := 2 * i.margin
	y2 := y + h + i.margin
	dc.DrawStringWrapped(enTitle, x2, y2, 0, 0, i.enTitle.maxWidth, 1, gg.AlignCenter)
}

func (i *Image) drawStats(dc *gg.Context, order int, title, value string) {
	dc.SetColor(color.RGBA{255, 255, 255, 127})

	// Title.
	dc.SetFontFace(i.subtitle.face)
	x1 := (float64(order) * i.subtitle.maxWidth) + i.margin
	y1 := float64(dc.Height()) - i.stats.maxHeight - i.subtitle.maxHeight - (i.margin * 4)
	dc.DrawStringWrapped(title, x1, y1, 0, 0, i.subtitle.maxWidth, 1, gg.AlignCenter)

	// Line.
	dc.DrawLine(
		2*i.margin+float64(order)*i.stats.maxWidth,
		float64(dc.Height())-i.stats.maxHeight-(i.margin*3),
		2*i.margin+float64(order)*i.stats.maxWidth+i.stats.maxWidth-2*i.margin,
		float64(dc.Height())-i.stats.maxHeight-(i.margin*3),
	)
	dc.SetRGBA(1, 1, 1, 0.5)
	dc.Stroke()

	// Value.
	dc.SetColor(color.White)
	dc.SetFontFace(i.stats.face)
	fontSize := i.stats.size
	w, h := dc.MeasureString(value)
	for w > i.stats.maxWidth-(2*i.margin) {
		fontSize--
		f, _ := gg.LoadFontFace(i.fontPath, fontSize)
		dc.SetFontFace(f)
		w, h = dc.MeasureString(value)
	}

	x2 := (float64(order) * i.stats.maxWidth) + (2 * i.margin)
	y2 := float64(dc.Height()) - i.stats.maxHeight - (i.margin * 3) + (i.stats.maxHeight-h)/2
	dc.DrawStringWrapped(value, x2, y2, 0, 0, i.stats.maxWidth-(2*i.margin), 1, gg.AlignCenter)
}
