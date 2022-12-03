# Cataas API

Cat as a service API wrapper for go.

> Cat as a service is a REST API to spread peace and love (or not) thanks to cats.

## Install

```bash
go get github.com/iArmanKarimi/cataas-api-go
```

## Examples

First, import the library:

```go 
import CataasAPI "github.com/iArmanKarimi/cataas-api-go"
```

`Get()`

```go 
c := new(CataasAPI.Cataas)
c.Gif = true
c.Text = "Hello"
c.Color = "LightBlue"
c.TextSize = 30
c.Size = CataasAPI.SIZE_MEDIUM
c.Encode()
data, err := c.Get()
```

`Download(string)`

```go 
c := new(CataasAPI.Cataas)
c.Filter = CataasAPI.FILTER_SEPIA
c.Encode()
err := c.Download("cat.png")
```

`GetAllTags()`

```go 
c := new(CataasAPI.Cataas)
tags, err := c.GetAllTags()
fmt.Printf("total tags: %d", len(tags))
```

`GetTags([]string, *GetCatsOptions)`

```go
c := new(CataasAPI.Cataas)
// get images with 'cute' tag
cats, err := c.GetCats([]string{"cute"}, nil)
// use options
cats, err := c.GetCats([]string{"cute"}, &CataasAPI.GetCatsOptions{
    Skip:  0,
    Limit: 10,
})
// print cat ids
for i, cat := range cats {
    fmt.Printf("%d) id: %s\n", i+1, cat.Id)
}
```

### Tips

+ After setting `Gif = true`, `Tag` is ignored.
+ Don't forget to call `Encode()` or `EncodeById` before `Get()` or `Download()`.

## Reference

[API website](https://cataas.com/)
