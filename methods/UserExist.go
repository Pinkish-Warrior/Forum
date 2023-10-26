package forum

import "database/sql"

func UserExist(email, username string, db *sql.DB) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email=? OR username=?", email, username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
