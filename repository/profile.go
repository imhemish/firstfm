package repository

func UpdateProfile(user User) error {
	about := ""
	name := ""
	language := ""
	picture := ""

	if user.About.Valid {
		about = user.About.String
	}
	if user.Name.Valid {
		name = user.Name.String
	}
	if user.Language.Valid {
		language = user.Language.String
	}
	if user.Picture.Valid {
		picture = user.Picture.String
	}
	_, error := connectedDb.Exec("UPDATE users SET about = $1, name = $2, language = $3, picture = $4 WHERE userid = $5", about, name, language, picture, user.UserID)

	return error

}
