// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"github.com/cridenour/go-postgis"
)

type Restaurant struct {
	ID       int64
	Name     string
	Location postgis.PointS
}
