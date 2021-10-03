package meta

import (
	"strings"

	"github.com/jivison/gowon-indexer/lib/db"
)

type ConversionMapItem struct {
	Key   string
	Value interface{}
}

type ConversionMap struct {
	_map map[string]ConversionMapItem
}

func (cm ConversionMap) PrivateSet(key string, value interface{}) {
	cm._map[strings.ToLower(key)] = ConversionMapItem{Key: key, Value: value}
}

func (cm ConversionMap) PrivateGet(key string) (ConversionMapItem, bool) {
	value, ok := cm._map[strings.ToLower(key)]

	return value, ok
}

func (cm ConversionMap) GetMap() map[string]ConversionMapItem {
	return cm._map
}

func CreateConversionMap() *ConversionMap {
	return &ConversionMap{_map: make(map[string]ConversionMapItem)}
}

type ArtistConversionMap struct{ *ConversionMap }

func (am ArtistConversionMap) Get(artistName string) (db.Artist, string, bool) {
	item, ok := am.PrivateGet(artistName)

	if ok {
		return item.Value.(db.Artist), item.Key, ok
	}

	return db.Artist{}, item.Key, ok
}

func (am ArtistConversionMap) Set(artistName string, artist db.Artist) {
	am.PrivateSet(artistName, artist)
}

func CreateArtistConversionMap() ArtistConversionMap {
	return ArtistConversionMap{ConversionMap: CreateConversionMap()}
}

type AlbumConversionMap struct{ *ConversionMap }

func (lm AlbumConversionMap) Get(artistName, albumName string) (db.Album, string, bool) {
	if artist, ok := lm.PrivateGet(artistName); ok {
		artistMap := artist.Value.(ConversionMap)

		album, ok := artistMap.PrivateGet(albumName)

		if ok {
			return album.Value.(db.Album), album.Key, ok
		}
	}

	return db.Album{}, albumName, false
}

func (lm AlbumConversionMap) Set(artistName, albumName string, album db.Album) {
	if _, ok := lm.PrivateGet(artistName); !ok {
		lm.PrivateSet(artistName, *CreateConversionMap())
	}

	artist, _ := lm.PrivateGet(artistName)

	artistMap := artist.Value.(ConversionMap)

	artistMap.PrivateSet(albumName, album)
}

func CreateAlbumConversionMap() AlbumConversionMap {
	return AlbumConversionMap{ConversionMap: CreateConversionMap()}
}

type TrackConversionMap struct{ *ConversionMap }

func (tm TrackConversionMap) Get(artistName, albumName, trackName string) (db.Track, string, bool) {
	if artist, ok := tm.PrivateGet(artistName); ok {
		artistMap := artist.Value.(ConversionMap)

		if album, ok := artistMap.PrivateGet(albumName); ok {
			albumMap := album.Value.(ConversionMap)

			track, ok := albumMap.PrivateGet(trackName)

			if ok {
				return track.Value.(db.Track), track.Key, ok
			}
		}
	}

	return db.Track{}, trackName, false
}

func (tm TrackConversionMap) Set(artistName, albumName, trackName string, track db.Track) {
	if _, ok := tm.PrivateGet(artistName); !ok {
		tm.PrivateSet(artistName, *CreateConversionMap())
	}

	artist, _ := tm.PrivateGet(artistName)

	artistMap := artist.Value.(ConversionMap)

	if _, ok := artistMap.PrivateGet(albumName); !ok {
		artistMap.PrivateSet(albumName, *CreateConversionMap())
	}

	album, _ := artistMap.PrivateGet(albumName)

	albumMap := album.Value.(ConversionMap)

	albumMap.PrivateSet(trackName, track)
}

func CreateTrackConversionMap() TrackConversionMap {
	return TrackConversionMap{ConversionMap: CreateConversionMap()}
}

type ArtistConversionCounter struct{ *ConversionMap }

func (acm ArtistConversionCounter) Get(artistName string) int32 {
	item, ok := acm.PrivateGet(artistName)

	if !ok {
		return 0
	}

	return item.Value.(int32)
}

func (acm ArtistConversionCounter) Increment(artistName string) {
	existingValue := acm.Get(artistName)

	acm.PrivateSet(artistName, existingValue+1)
}

func (acm ArtistConversionCounter) Set(artistName string, value int32) {
	acm.PrivateSet(artistName, value)
}

func CreateArtistConversionCounter() ArtistConversionCounter {
	return ArtistConversionCounter{ConversionMap: CreateConversionMap()}
}

type AlbumConversionCounter struct{ *ConversionMap }

func (lcm AlbumConversionCounter) Get(artistName, albumName string) int32 {
	if artist, ok := lcm.PrivateGet(artistName); ok {
		artistMap := artist.Value.(ConversionMap)

		album, ok := artistMap.PrivateGet(albumName)

		if ok {
			return album.Value.(int32)
		}
	}

	return 0
}

func (lcm AlbumConversionCounter) Increment(artistName, albumName string) {
	existingValue := lcm.Get(artistName, albumName)

	if _, ok := lcm.PrivateGet(artistName); !ok {
		lcm.PrivateSet(artistName, *CreateConversionMap())
	}

	artist, _ := lcm.PrivateGet(artistName)

	artistMap := artist.Value.(ConversionMap)

	artistMap.PrivateSet(albumName, existingValue+1)
}

func CreateAlbumConversionCounter() AlbumConversionCounter {
	return AlbumConversionCounter{ConversionMap: CreateConversionMap()}
}

type TrackConversionCounter struct{ *ConversionMap }

func (tcm TrackConversionCounter) Get(artistName, albumName, trackName string) int32 {
	if artist, ok := tcm.PrivateGet(artistName); ok {
		artistMap := artist.Value.(ConversionMap)

		if album, ok := artistMap.PrivateGet(albumName); ok {
			albumMap := album.Value.(ConversionMap)

			track, ok := albumMap.PrivateGet(trackName)

			if ok {
				return track.Value.(int32)
			}
		}
	}

	return 0
}

func (tcm TrackConversionCounter) Increment(artistName, albumName, trackName string) {
	existingValue := tcm.Get(artistName, albumName, trackName)

	if _, ok := tcm.PrivateGet(artistName); !ok {
		tcm.PrivateSet(artistName, *CreateConversionMap())
	}

	artist, _ := tcm.PrivateGet(artistName)

	artistMap := artist.Value.(ConversionMap)

	if _, ok := artistMap.PrivateGet(albumName); !ok {
		artistMap.PrivateSet(albumName, *CreateConversionMap())
	}

	album, _ := artistMap.PrivateGet(albumName)

	albumMap := album.Value.(ConversionMap)

	albumMap.PrivateSet(trackName, existingValue+1)
}

func CreateTrackConversionCounter() TrackConversionCounter {
	return TrackConversionCounter{ConversionMap: CreateConversionMap()}
}

type RateYourMusicAlbumConversionMap struct{ *ConversionMap }

func (rlm RateYourMusicAlbumConversionMap) Get(rymsID string) (db.RateYourMusicAlbum, string, bool) {
	item, ok := rlm.PrivateGet(rymsID)

	if ok {
		return item.Value.(db.RateYourMusicAlbum), item.Key, ok
	}

	return db.RateYourMusicAlbum{}, item.Key, ok
}

func (rlm RateYourMusicAlbumConversionMap) Set(rymsID string, album db.RateYourMusicAlbum) {
	rlm.PrivateSet(rymsID, album)
}

func CreateRateYourMusicAlbumConversionMap() RateYourMusicAlbumConversionMap {
	return RateYourMusicAlbumConversionMap{ConversionMap: CreateConversionMap()}
}

type TagConversionMap struct{ *ConversionMap }

func (tm TagConversionMap) Get(tagName string) (db.Tag, string, bool) {
	item, ok := tm.PrivateGet(tagName)

	return item.Value.(db.Tag), item.Key, ok
}

func (tm TagConversionMap) Set(tagName string, artist db.Tag) {
	tm.PrivateSet(tagName, artist)
}

func CreateTagConversionMap() TagConversionMap {
	return TagConversionMap{ConversionMap: CreateConversionMap()}
}
