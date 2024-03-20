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

func GetProfile(uid string) (User, error) {
	query, err := connectedDb.Query("SELECT (about, name, picture) FROM users WHERE userid=$1", uid)
	var user User
	if err != nil {
		return user, err

	}
	if query.Next() {
		err = query.Scan(&user.About, &user.Name, &user.Picture)
	}

	return user, err
}
