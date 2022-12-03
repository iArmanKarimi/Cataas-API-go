package CataasAPI

import (
	"errors"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
	"time"
)

func print_test(t *testing.T, want interface{}, got interface{}) {
	t.Errorf(`want: %v ; got: %v`, want, got)
}

func TestEncode(t *testing.T) {
	// want := "https://"
	// got := c.Encode().String()
	const (
		gif       = false
		tag       = "cute,fluffy"
		text      = "hello"
		width     = 200
		height    = 200
		textSize  = 20
		textColor = "Blue"
		size      = SIZE_SMALL
		filter    = FILTER_BLUR
	)
	c := &Cataas{
		Gif:       gif,
		Tag:       tag,
		Text:      text,
		Width:     width,
		Height:    height,
		TextSize:  textSize,
		TextColor: textColor,
		Size:      SIZE_SMALL,
		Filter:    FILTER_BLUR,
	}
	uri := c.Encode()

	if _text := path.Base(uri.Path); _text != text {
		print_test(t, text, _text)
	}
	if _tag := strings.Split(path.Dir(uri.Path), "/")[1]; _tag != tag {
		print_test(t, tag, _tag)
	}
	if _width := uri.Query().Get("width"); _width != strconv.Itoa(width) {
		print_test(t, width, _width)
	}
	if _height := uri.Query().Get("height"); _height != strconv.Itoa(height) {
		print_test(t, height, _height)
	}
	if _text_size := uri.Query().Get("size"); _text_size != strconv.Itoa(textSize) {
		print_test(t, textSize, _text_size)
	}
	if _text_color := uri.Query().Get("color"); _text_color != textColor {
		print_test(t, textColor, _text_color)
	}
	if _size := uri.Query().Get("type"); _size != string(size) {
		print_test(t, string(size), _size)
	}
	if _filter := uri.Query().Get("filter"); _filter != string(filter) {
		print_test(t, string(filter), _filter)
	}
	if _gif := uri.Query().Get("gif"); _gif != "" {
		t.Error("Gif option not be set when tag is")
	}
}

func TestGet(t *testing.T) {
	c := &Cataas{Text: "Hello", TextSize: 20}
	img, err := c.Get()
	if err != nil {
		t.Error(err)
	}
	if len(img) < 1 {
		t.Error("Empty image")
	}
}

func TestDownload(t *testing.T) {
	c := &Cataas{}
	const file_path = "__img__test__.png"
	err := c.Download(file_path)
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat(file_path); err == nil {
		os.Remove(file_path)
	} else if errors.Is(err, os.ErrNotExist) {
		t.Error("File was not created. download failed.")
	} else {
		t.Error(err)
	}
}

func TestEncodeById(t *testing.T) {
	c := &Cataas{}
	id := rand.New(rand.NewSource(time.Now().Unix())).Int()
	uri := c.EncodeById(strconv.Itoa(id))
	if _id, err := strconv.Atoi(path.Base(uri.Path)); err != nil {
		t.Error(err)
	} else if _id != id {
		print(t, id, _id)
	}
}

func TestGetAllTags(t *testing.T) {
	c := &Cataas{}
	if tags, err := c.GetAllTags(); err != nil {
		t.Error(err)
	} else if len(tags) < 1 {
		t.Error("Empty tags list")
	}
}

// Todo
// func TestGetCats(t *testing.T) { }

// Benchmarks //
func BenchmarkTestGet_5(b *testing.B) {
	c := &Cataas{}
	for i := 0; i < 5; i++ {
		_, err := c.Get()
		if err != nil {
			_ = err.Error()
		}
	}
}
