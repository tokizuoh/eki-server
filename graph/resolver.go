package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tokizuoh/eki-server/graph/model"
)

type Resolver struct{}

type Station struct {
	id   string
	name string
}

var db *sql.DB

func init() {
	conn, err := sql.Open("mysql", "root:root@/station_db")
	if err != nil {
		panic(err)
	}
	db = conn
}

func (r *Resolver) station(name string) ([]*model.Station, error) {
	q := fmt.Sprintf("SELECT * FROM station WHERE name='%s'", name)
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mss []*model.Station
	for rows.Next() {
		var s Station
		err := rows.Scan(&s.id, &s.name)
		if err != nil {
			panic(err.Error())
		}

		ms := model.Station{
			DetabaseID: s.id,
			Name:       s.name,
		}
		mss = append(mss, &ms)
	}

	return mss, nil
}

func (r *Resolver) searchStation(forArg string) ([]*model.Station, error) {
	q := fmt.Sprintf("SELECT * FROM station WHERE name LIKE '%%%s%%'", forArg)
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mss []*model.Station
	for rows.Next() {
		var s Station
		err := rows.Scan(&s.id, &s.name)
		if err != nil {
			panic(err.Error())
		}

		ms := model.Station{
			DetabaseID: s.id,
			Name:       s.name,
		}
		mss = append(mss, &ms)
	}

	return mss, nil
}
