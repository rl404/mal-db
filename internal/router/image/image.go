package image

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/image"
	"github.com/rl404/mal-db/internal/pkg/utils"
)

// Image is image api routes.
type Image struct {
	imager image.API
}

// New to create new image api routes.
func New(img image.API) *Image {
	return &Image{
		imager: img,
	}
}

// Register to register all iamge api route endpoints.
func (i *Image) Register(r chi.Router) {
	r.Get("/image/anime/{id}", i.getAnime)
	r.Get("/image/manga/{id}", i.getManga)
}

// @summary Get anime card image.
// @tags image
// @produce jpeg
// @param id path integer true "Anime ID"
// @success 200
// @router /image/anime/{id} [get]
func (i *Image) getAnime(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	img, code, err := i.imager.GetAnimeCard(id)
	utils.ResponseWithJPEG(w, img, code, err)
}

// @summary Get manga card image.
// @tags image
// @produce jpeg
// @param id path integer true "Manga ID"
// @success 200
// @router /image/manga/{id} [get]
func (i *Image) getManga(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	img, code, err := i.imager.GetMangaCard(id)
	utils.ResponseWithJPEG(w, img, code, err)
}
