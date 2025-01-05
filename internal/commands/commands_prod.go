// +build prod

package commands

import (
	"WeManageIt/internal/migrations"
)

func init() {
	migration = migrations.Migrations
}
