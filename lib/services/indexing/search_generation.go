package indexing

func (i Indexing) GenerateAlbumsToSearch(albumNames []AlbumToConvert) []interface{} {
	var albumsToSearch []interface{}

	for _, album := range albumNames {
		albumsToSearch = append(albumsToSearch, []interface{}{album.ArtistName, album.AlbumName})
	}

	return albumsToSearch
}

func (i Indexing) GenerateTracksToSearch(trackNames []TrackToConvert) []interface{} {
	var tracksToSearch []interface{}

	for _, track := range trackNames {
		tracksToSearch = append(tracksToSearch, []interface{}{track.ArtistName, track.TrackName, track.AlbumName})
	}

	return tracksToSearch
}
