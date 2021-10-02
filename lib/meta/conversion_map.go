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

func (cm ConversionMap) set(key string, value interface{}) {
	cm._map[strings.ToLower(key)] = ConversionMapItem{Key: key, Value: value}
}

func (cm ConversionMap) get(key string) (ConversionMapItem, bool) {
	value, ok := cm._map[strings.ToLower(key)]

	return value, ok
}

func (cm ConversionMap) GetMap() map[string]ConversionMapItem {
	return cm._map
}

func createConversionMap() *ConversionMap {
	return &ConversionMap{_map: make(map[string]ConversionMapItem)}
}

type ArtistConversionMap struct{ *ConversionMap }

func (am ArtistConversionMap) Get(artistName string) (db.Artist, string, bool) {
	item, ok := am.get(artistName)

	if ok {
		return item.Value.(db.Artist), item.Key, ok
	}

	return db.Artist{}, item.Key, ok
}

func (am ArtistConversionMap) Set(artistName string, artist db.Artist) {
	am.set(artistName, artist)
}

func CreateArtistConversionMap() ArtistConversionMap {
	return ArtistConversionMap{ConversionMap: createConversionMap()}
}

type AlbumConversionMap struct{ *ConversionMap }

func (lm AlbumConversionMap) Get(artistName, albumName string) (db.Album, string, bool) {
	if artist, ok := lm.get(artistName); ok {
		artistMap := artist.Value.(ConversionMap)

		album, ok := artistMap.get(albumName)

		if ok {
			return album.Value.(db.Album), album.Key, ok
		}
	}

	return db.Album{}, albumName, false
}

func (lm AlbumConversionMap) Set(artistName, albumName string, album db.Album) {
	if _, ok := lm.get(artistName); !ok {
		lm.set(artistName, *createConversionMap())
	}

	artist, _ := lm.get(artistName)

	artistMap := artist.Value.(ConversionMap)

	artistMap.set(albumName, album)
}

func CreateAlbumConversionMap() AlbumConversionMap {
	return AlbumConversionMap{ConversionMap: createConversionMap()}
}

type TrackConversionMap struct{ *ConversionMap }

func (tm TrackConversionMap) Get(artistName, albumName, trackName string) (db.Track, string, bool) {
	if artist, ok := tm.get(artistName); ok {
		artistMap := artist.Value.(ConversionMap)

		if album, ok := artistMap.get(albumName); ok {
			albumMap := album.Value.(ConversionMap)

			track, ok := albumMap.get(trackName)

			if ok {
				return track.Value.(db.Track), track.Key, ok
			}
		}
	}

	return db.Track{}, trackName, false
}

func (tm TrackConversionMap) Set(artistName, albumName, trackName string, track db.Track) {
	if _, ok := tm.get(artistName); !ok {
		tm.set(artistName, *createConversionMap())
	}

	artist, _ := tm.get(artistName)

	artistMap := artist.Value.(ConversionMap)

	if _, ok := artistMap.get(albumName); !ok {
		artistMap.set(albumName, *createConversionMap())
	}

	album, _ := artistMap.get(albumName)

	albumMap := album.Value.(ConversionMap)

	albumMap.set(trackName, track)
}

func CreateTrackConversionMap() TrackConversionMap {
	return TrackConversionMap{ConversionMap: createConversionMap()}
}

type ArtistConversionCounter struct{ *ConversionMap }

func (acm ArtistConversionCounter) Get(artistName string) int32 {
	item, ok := acm.get(artistName)

	if !ok {
		return 0
	}

	return item.Value.(int32)
}

func (acm ArtistConversionCounter) Increment(artistName string) {
	existingValue := acm.Get(artistName)

	acm.set(artistName, existingValue+1)
}

func (acm ArtistConversionCounter) Set(artistName string, value int32) {
	acm.set(artistName, value)
}

func CreateArtistConversionCounter() ArtistConversionCounter {
	return ArtistConversionCounter{ConversionMap: createConversionMap()}
}

type AlbumConversionCounter struct{ *ConversionMap }

func (lcm AlbumConversionCounter) Get(artistName, albumName string) int32 {
	if artist, ok := lcm.get(artistName); ok {
		artistMap := artist.Value.(ConversionMap)

		album, ok := artistMap.get(albumName)

		if ok {
			return album.Value.(int32)
		}
	}

	return 0
}

func (lcm AlbumConversionCounter) Increment(artistName, albumName string) {
	existingValue := lcm.Get(artistName, albumName)

	if _, ok := lcm.get(artistName); !ok {
		lcm.set(artistName, *createConversionMap())
	}

	artist, _ := lcm.get(artistName)

	artistMap := artist.Value.(ConversionMap)

	artistMap.set(albumName, existingValue+1)
}

func CreateAlbumConversionCounter() AlbumConversionCounter {
	return AlbumConversionCounter{ConversionMap: createConversionMap()}
}

type TrackConversionCounter struct{ *ConversionMap }

func (tcm TrackConversionCounter) Get(artistName, albumName, trackName string) int32 {
	if artist, ok := tcm.get(artistName); ok {
		artistMap := artist.Value.(ConversionMap)

		if album, ok := artistMap.get(albumName); ok {
			albumMap := album.Value.(ConversionMap)

			track, ok := albumMap.get(trackName)

			if ok {
				return track.Value.(int32)
			}
		}
	}

	return 0
}

func (tcm TrackConversionCounter) Increment(artistName, albumName, trackName string) {
	existingValue := tcm.Get(artistName, albumName, trackName)

	if _, ok := tcm.get(artistName); !ok {
		tcm.set(artistName, *createConversionMap())
	}

	artist, _ := tcm.get(artistName)

	artistMap := artist.Value.(ConversionMap)

	if _, ok := artistMap.get(albumName); !ok {
		artistMap.set(albumName, *createConversionMap())
	}

	album, _ := artistMap.get(albumName)

	albumMap := album.Value.(ConversionMap)

	albumMap.set(trackName, existingValue+1)
}

func CreateTrackConversionCounter() TrackConversionCounter {
	return TrackConversionCounter{ConversionMap: createConversionMap()}
}

type RateYourMusicAlbumConversionMap struct{ *ConversionMap }

func (rlm RateYourMusicAlbumConversionMap) Get(artistName string) (db.RateYourMusicAlbum, string, bool) {
	item, ok := rlm.get(artistName)

	if ok {
		return item.Value.(db.RateYourMusicAlbum), item.Key, ok
	}

	return db.RateYourMusicAlbum{}, item.Key, ok
}

func (rlm RateYourMusicAlbumConversionMap) Set(artistName string, artist db.RateYourMusicAlbum) {
	rlm.set(artistName, artist)
}

func CreateRateYourMusicAlbumConversionMap() RateYourMusicAlbumConversionMap {
	return RateYourMusicAlbumConversionMap{ConversionMap: createConversionMap()}
}

type TagConversionMap struct{ *ConversionMap }

func (tm TagConversionMap) Get(tagName string) (db.Tag, string, bool) {
	item, ok := tm.get(tagName)

	return item.Value.(db.Tag), item.Key, ok
}

func (tm TagConversionMap) Set(tagName string, artist db.Tag) {
	tm.set(tagName, artist)
}

func CreateTagConversionMap() TagConversionMap {
	return TagConversionMap{ConversionMap: createConversionMap()}
}
