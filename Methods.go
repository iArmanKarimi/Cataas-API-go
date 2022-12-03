package CataasAPI

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	pathCat     = "cat"
	pathTags    = "api/cats"
	pathAllTags = "api/tags"
)

/* Returns url to be used  */
func (c *Cataas) Encode() *url.URL {
	uri.Path = pathCat
	q := uri.Query()
	if c.Gif {
		c.Tag = "gif"
	}
	if c.Tag != "" {
		uri.Path = path.Join(uri.Path, c.Tag)
	}
	if c.Text != "" {
		uri.Path = path.Join(uri.Path, "says/"+c.Text)
		q.Set("size", strconv.Itoa(int(c.TextSize)))
		if c.TextColor != "" {
			q.Set("color", c.TextColor)
		}
	}
	if c.Size != "" {
		q.Set("type", string(c.Size))
	}
	if c.Filter != "" {
		q.Set("filter", string(c.Filter))
	}
	q.Set("width", strconv.Itoa(int(c.Width)))
	q.Set("height", strconv.Itoa(int(c.Height)))
	uri.RawQuery = q.Encode()
	return uri
}

/* sends get request then reads the response body and returns the data.
can be used after encoding using Encode or EncodeById.
*/
func (c *Cataas) Get() (img_data []byte, err error) {
	var resp *http.Response
	resp, err = http.Get(uri.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	img_data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return img_data, nil
}

/* Download image and save it to the given path */
func (c *Cataas) Download(path string) (err error) {
	// Todo: add file_name arg and join it with path
	var resp *http.Response
	resp, err = http.Get(uri.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var out *os.File
	out, err = os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

/* returns the encoded url by using the id of the image */
func (c *Cataas) EncodeById(id string) *url.URL {
	uri.Path = pathCat + "/" + id
	return uri
}

/* get all tags */
func (c *Cataas) GetAllTags() (tags []string, err error) {
	uri.Path = pathAllTags
	var data []byte
	data, err = c.Get()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tags)
	return tags, err
}

/* get cats which include the given tag(s) */
func (c *Cataas) GetCats(tags []string, options *GetCatsOptions) (cats []*CatByTag, err error) {
	var (
		skip  = "0"
		limit = "0"
	)
	if options != nil {
		skip = strconv.Itoa(int(options.Skip))
		limit = strconv.Itoa(int(options.Limit))
	}
	q := uri.Query()
	q.Set("tags", strings.Join(tags, ","))
	q.Set("skip", skip)
	q.Set("limit", limit)
	uri.Path = pathTags
	uri.RawQuery = q.Encode()
	var data []byte
	data, err = c.Get()
	if err != nil {
		return nil, err
	}
	container := make([]*CatByTag, 0, 200)
	err = json.Unmarshal(data, &container)
	if err != nil {
		return nil, err
	}
	return container, nil
}
