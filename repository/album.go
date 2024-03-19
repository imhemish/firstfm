package repository

type AlbumDetails struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Artists     []string `json:"artists"`
	Likes       int      `json:"likes"`
}

func GetAlbumDetails(album uint) (AlbumDetails, error) {
	// Database implementation later
	return AlbumDetails{"Random Album", "Sung by Random", []string{"Karan Aujla"}, 13}, nil

}
