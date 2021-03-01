package model

// SafeGetArtistName returns the artist name or nil
func (ti TrackInput) SafeGetArtistName() *string {
	if ti.Artist != nil {
		return ti.Artist.Name
	}
	return nil
}

// SafeGetAlbumName returns the artist name or nil
func (ti TrackInput) SafeGetAlbumName() *string {
	if ti.Album != nil {
		return ti.Album.Name
	}
	return nil
}

// SafeGetArtistName returns the artist name or nil
func (li AlbumInput) SafeGetArtistName() *string {
	if li.Artist != nil {
		return li.Artist.Name
	}
	return nil
}
