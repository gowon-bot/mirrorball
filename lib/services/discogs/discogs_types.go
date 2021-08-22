package discogs

type Pagination struct {
	PerPage int `json:"per_page"`
	Pages   int `json:"pages"`
	Page    int `json:"page"`
	Items   int `json:"items"`
	URLs    struct {
		Next string `json:"next"`
		Last string `json:"last"`
	} `json:"urls"`
}

type SimpleRelease struct {
	Id               int `json:"id"`
	InstanceID       int `json:"instance_id"`
	FolderID         int `json:"folder_id"`
	Rating           int `json:"rating"`
	BasicInformation struct {
		ID          int    `json:"id"`
		Year        int    `json:"year"`
		Title       string `json:"title"`
		ResourceURL string `json:"resource_url"`
		Thumb       string `json:"thumb"`
		CoverImage  string `json:"cover_image"`
		Formats     []struct {
			Quantity     string   `json:"qty"`
			Name         string   `json:"name"`
			Descriptions []string `json:"descriptions"`
		} `json:"formats"`
		Labels []struct {
			ResourceURL   string `json:"resource_url"`
			EntityType    string `json:"entity_type"`
			CatalogNumber string `json:"catno"`
			Name          string `json:"name"`
			ID            int    `json:"id"`
		} `json:"labels"`

		Artists []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Join        string `json:"join"`
			ResourceURL string `json:"resource_url"`
			ANV         string `json:"anv"`
			Tracks      string `json:"tracks"`
			Role        string `json:"role"`
		} `json:"artists"`

		Genres []string `json:"genres"`
		Styles []string `json:"styles"`
	} `json:"basic_information"`
	Notes []struct {
		Field_id int    `json:"field_id"`
		Value    string `json:"value"`
	} `json:"notes"`
}

type Collection struct {
	Pagination Pagination      `json:"pagination"`
	Releases   []SimpleRelease `json:"releases"`
}
