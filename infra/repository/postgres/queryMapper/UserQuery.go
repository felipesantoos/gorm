package queryMapper

import (
	"database/sql"
	"gorm/core"
)

const UserQueryBase = `
SELECT
user_id,
user_login,
user_email,
user_avatar,
user_active,
user_last_login,
user_created,
user_updated
`

const UserQueryKey = UserQueryBase + `
FROM users
WHERE user_id = :user_id
`

func ToParams(u *core.User) map[string]interface{} {
	return map[string]interface{}{
		"user_id":         u.ID,
		"user_login":      u.Login,
		"user_email":      u.Email,
		"user_avatar":     u.Avatar,
		"user_active":     u.Active,
		"user_last_login": u.LastLogin,
		"user_created":    u.Created,
		"user_updated":    u.Updated,
	}
}

func ScanRow(scanner *sql.Row, dest *core.User) error {
	return scanner.Scan(
		&dest.ID,
		&dest.Login,
		&dest.Email,
		&dest.Avatar,
		&dest.Active,
		&dest.LastLogin,
		&dest.Created,
		&dest.Updated,
	)
}
