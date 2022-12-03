package CataasAPI

import "net/url"

type Size string
type Filter string

// small or sm, medium or md, square or sq, original or or
const (
	SIZE_SMALL    Size = "sm"
	SIZE_MEDIUM   Size = "md"
	SIZE_SQUARE   Size = "sq"
	SIZE_ORIGINAL Size = "or"
)

// blur, mono, sepia, negative, paint, pixel
const (
	FILTER_BLUR     Filter = "blur"
	FILTER_MONO     Filter = "mono"
	FILTER_PAINT    Filter = "paint"
	FILTER_PIXEL    Filter = "pixel"
	FILTER_SEPIA    Filter = "sepia"
	FILTER_NEGATIVE Filter = "negative"
)

var uri = &url.URL{
	Host:   "cataas.com",
	Scheme: "https",
}

type Cataas struct {
	Gif       bool
	Tag       string
	Text      string
	Width     uint
	Height    uint
	TextSize  uint
	TextColor string
	Size      Size
	Filter    Filter
}

type CatByTag struct {
	Id        string   `json:"id"`
	CreatedAt string   `json:"created_at"`
	Tags      []string `json:"tags"`
}

type GetCatsOptions struct {
	Skip  uint
	Limit uint
}
