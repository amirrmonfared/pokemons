// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"time"
)

type Pokemon struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Type1      string    `json:"type1"`
	Type2      string    `json:"type2"`
	Hp         int32     `json:"hp"`
	Attack     int32     `json:"attack"`
	Defense    int32     `json:"defense"`
	SpAtk      int32     `json:"sp_atk"`
	SpDef      int32     `json:"sp_def"`
	Speed      int32     `json:"speed"`
	Generation int32     `json:"generation"`
	Legendary  bool      `json:"legendary"`
	CreatedAt  time.Time `json:"created_at"`
}
