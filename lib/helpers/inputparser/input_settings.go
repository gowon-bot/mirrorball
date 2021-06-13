package inputparser

type InputSettings struct {
	ArtistPath, AlbumPath, TrackPath string
	UserRelation                     string
	DefaultSort                      string
}

func (s InputSettings) getArtistPath() string {
	if s.ArtistPath == "" {
		return "artist"
	}
	return s.ArtistPath
}

func (s InputSettings) getAlbumPath() string {
	if s.AlbumPath == "" {
		return "album"
	}
	return s.AlbumPath
}

func (s InputSettings) getTrackPath() string {
	if s.TrackPath == "" {
		return "track"
	}
	return s.TrackPath
}

func (s InputSettings) getDefaultSort() string {
	if s.DefaultSort == "" {
		return "1 desc"
	}
	return s.DefaultSort
}

func (s InputSettings) getUserRelation() string {
	if s.UserRelation == "" {
		return "User"
	}
	return s.UserRelation
}
