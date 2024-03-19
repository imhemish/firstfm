package repository

func UpdatePassword(userid string, hashedpwd string) error {
	_, error := connectedDb.Exec("UPDATE users SET hashedpwd = $1 WHERE userid = $2", hashedpwd, userid)

	return error

}

func GetHashedPassword(userid string) (string, error) {
	var pwd string
	query := connectedDb.QueryRow("SELECT hashedpwd from users WHERE userid=$1", userid)
	err := query.Scan(&pwd)
	return pwd, err
}
