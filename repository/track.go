package repository

import "fmt"

type TrackDetails struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Artists     []string `json:"artists"`
	Likes       uint     `json:"likes"`
	GlobalPlays uint     `json:"global_plays"`
	UserPlays   uint     `json:"user_plays"`
}

func GetTrackDetails(trackID uint, userId uint) (TrackDetails, error) {
	fmt.Println("TrackID: ", trackID, "userID: ", userId)
	return TrackDetails{
		"Softly",
		"",
		[]string{"Karan Aujla"},
		12,
		7,
		2,
	}, nil
}
