package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

func create_user_table() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(`
                CREATE TABLE IF NOT EXISTS usermanage (
                    userid INT AUTO_INCREMENT PRIMARY KEY,
                    username VARCHAR(100) NOT NULL,
                    userphone VARCHAR(15),
                    useremail VARCHAR(100)
                )
            `)
			return err
		},
	}
}
