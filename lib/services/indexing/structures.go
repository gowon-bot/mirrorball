package indexing

type TrackToConvert struct {
	ArtistName string
	TrackName  string
	AlbumName  *string
}

type AlbumToConvert struct {
	ArtistName string
	AlbumName  string
}
