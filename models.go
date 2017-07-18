package main

import (
	"fmt"
	"proto"
)

type Artist struct {
	Id          int32
	Name        string
	Image       string
	Thumbnail   string
	PlayerImage string
	Order       int32
	Twitter     string
	Enabled     bool
}

type Category struct {
	Id      int32
	Title   string
	Order   int32
	Enabled bool
}

type Mix struct {
	Id       int32
	Title    string
	Url      string
	Duration int32
	Featured bool
	Enabled  bool
	IsNew    bool
}

type Promo struct {
	Image string
	AdUrl string
	IsAd  bool
}

type Protobuf interface {
}

func (artist Artist) Proto() proto.Artist {
	return proto.Artist{
		Id:          artist.Id,
		Name:        artist.Name,
		Image:       artist.Image,
		Thumbnail:   artist.Thumbnail,
		PlayerImage: artist.PlayerImage,
		Order:       artist.Order,
		Twitter:     artist.Twitter,
		Enabled:     artist.Enabled,
	}
}

func (category Category) Proto() proto.Category {
	return proto.Category{
		Id:      category.Id,
		Title:   category.Title,
		Order:   category.Order,
		Enabled: category.Enabled,
	}
}

func (mix Mix) Proto() proto.Song {
	return proto.Song{
		Id:       mix.Id,
		Title:    mix.Title,
		Url:      mix.Url,
		Duration: mix.Duration,
		Featured: mix.Featured,
		Enabled:  mix.Enabled,
	}
}

func (promo Promo) Proto() proto.Promo {
	return proto.Promo{
		Image: promo.Image,
		AdUrl: promo.AdUrl,
		IsAd:  promo.IsAd,
	}
}

func (artist Artist) String() string {
	return fmt.Sprintf("%s", artist.Proto())
}

func (category Category) String() string {
	return fmt.Sprintf("%s", category.Proto())
}

func (mix Mix) String() string {
	return fmt.Sprintf("%s", mix.Proto())
}

func (promo Promo) String() string {
	return fmt.Sprintf("%s", promo.Proto())
}

// func catalog(categories []Category, songs []Mix, artists []Artist, promos []Promo) proto.Catalog {
// 	return proto.Catalog{
// 		Categories:  [],
// 		Songs:       songs,
// 		Artists:     artists,
// 		Promos:      promos,
// 		LastUpdated: 1,
// 	}
// }
