package image

import (
	"fmt"
	"image"
	"net/http"

	"github.com/fogleman/gg"
	"github.com/rl404/mal-db/internal/loader"
	"github.com/rl404/mal-db/internal/pkg/utils"
	"golang.org/x/image/font"
)

// API to generate entry cards.
type API interface {
	GetAnimeCard(id int) (image.Image, int, error)
	GetMangaCard(id int) (image.Image, int, error)
}

// Image contains basic image generator config.
type Image struct {
	api loader.API

	// Layout.
	width  int
	height int
	margin float64

	// Font.
	fontPath string
	title    fontConfig
	enTitle  fontConfig
	subtitle fontConfig
	stats    fontConfig
}

type fontConfig struct {
	face      font.Face
	size      float64
	maxHeight float64
	maxWidth  float64
}

// New to create new image generator.
func New(api loader.API) (i *Image, err error) {
	img := Image{
		api:      api,
		width:    1200,
		height:   630,
		margin:   20.0,
		fontPath: "../../fonts/Roboto-Bold.ttf",
		title: fontConfig{
			size:      70,
			maxHeight: 630 - (20 * 2) - (20 * 2) - 60 - (20 * 2) - 40 - (20 * 2) - 40 - (20 * 2),
			maxWidth:  1200 - (20 * 4),
		},
		enTitle: fontConfig{
			size:      40,
			maxHeight: 40,
			maxWidth:  1200 - (20 * 4),
		},
		subtitle: fontConfig{
			size:      40,
			maxHeight: 40,
			maxWidth:  ((1200.0 - (20.0 * 2.0)) / 5.0),
		},
		stats: fontConfig{
			size:      60,
			maxHeight: 60,
			maxWidth:  ((1200.0 - (20.0 * 2.0)) / 5.0),
		},
	}

	img.title.face, err = gg.LoadFontFace(img.fontPath, img.title.size)
	if err != nil {
		return nil, err
	}

	img.enTitle.face, err = gg.LoadFontFace(img.fontPath, img.enTitle.size)
	if err != nil {
		return nil, err
	}

	img.subtitle.face, err = gg.LoadFontFace(img.fontPath, img.subtitle.size)
	if err != nil {
		return nil, err
	}

	img.stats.face, err = gg.LoadFontFace(img.fontPath, img.stats.size)
	if err != nil {
		return nil, err
	}

	return &img, nil
}

// GetAnimeCard to generate anime card image.
func (i *Image) GetAnimeCard(id int) (image.Image, int, error) {
	// Prepare data.
	data, _, code, err := i.api.GetAnime(id)
	if err != nil {
		return nil, code, err
	}

	// Init.
	dc := gg.NewContext(i.width, i.height)

	// Prepare background.
	if err = i.drawBackground(dc, data.Image); err != nil {
		i.drawDefaultBackground(dc)
	}

	// Prepare overlay.
	i.drawOverlay(dc)

	// Prepare texts.
	i.drawTitle(dc, data.Title, data.AlternativeTitles.English)
	i.drawStats(dc, 0, "Rank", "#"+utils.Thousands(data.Rank))
	i.drawStats(dc, 1, "Score", fmt.Sprintf("%.2f", data.Score))
	i.drawStats(dc, 2, "Popularity", "#"+utils.Thousands(data.Popularity))
	i.drawStats(dc, 3, "Member", utils.Thousands(data.Member))
	i.drawStats(dc, 4, "Favorite", utils.Thousands(data.Favorite))

	return dc.Image(), http.StatusOK, nil
}

// GetMangaCard to generate manga card image.
func (i *Image) GetMangaCard(id int) (image.Image, int, error) {
	// Prepare data.
	data, _, code, err := i.api.GetManga(id)
	if err != nil {
		return nil, code, err
	}

	// Init.
	dc := gg.NewContext(i.width, i.height)

	// Prepare background.
	if err = i.drawBackground(dc, data.Image); err != nil {
		i.drawDefaultBackground(dc)
	}

	// Prepare overlay.
	i.drawOverlay(dc)

	// Prepare texts.
	i.drawTitle(dc, data.Title, data.AlternativeTitles.English)
	i.drawStats(dc, 0, "Rank", "#"+utils.Thousands(data.Rank))
	i.drawStats(dc, 1, "Score", fmt.Sprintf("%.2f", data.Score))
	i.drawStats(dc, 2, "Popularity", "#"+utils.Thousands(data.Popularity))
	i.drawStats(dc, 3, "Member", utils.Thousands(data.Member))
	i.drawStats(dc, 4, "Favorite", utils.Thousands(data.Favorite))

	return dc.Image(), http.StatusOK, nil
}
