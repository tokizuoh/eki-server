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
	id      string
	name    string
	isLiked bool
}

var db *sql.DB

func init() {
	conn, err := sql.Open("mysql", "root:root@/station_db")
	if err != nil {
		panic(err)
	}
	db = conn
}

// Query

func (r *Resolver) station(name string) ([]*model.Station, error) {
	q := fmt.Sprintf("SELECT * FROM station WHERE name='%s'", name)
	return search(q)
}

func (r *Resolver) searchStation(forArg string) ([]*model.Station, error) {
	q := fmt.Sprintf("SELECT * FROM station WHERE name LIKE '%%%s%%'", forArg)
	return search(q)
}

func search(query string) ([]*model.Station, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mss []*model.Station
	for rows.Next() {
		var s Station
		err := rows.Scan(&s.id, &s.name, &s.isLiked)
		if err != nil {
			panic(err.Error())
		}

		ms := model.Station{
			DetabaseID: s.id,
			Name:       s.name,
			IsLiked:    s.isLiked,
		}
		mss = append(mss, &ms)
	}

	return mss, nil
}

// Mutation

func (r *Resolver) likeStation(databaseID string) (*model.Station, error) {
	upd, err := db.Prepare("UPDATE station SET is_liked=? WHERE id=?")
	if err != nil {
		return nil, err
	}

	_, err = upd.Exec(true, databaseID)
	if err != nil {
		return nil, err
	}

	q := fmt.Sprintf("SELECT * FROM station WHERE id='%s'", databaseID)
	ss, err := search(q)
	if err != nil {
		return nil, err
	}

	if len(ss) != 1 {
		return nil, fmt.Errorf("unexpected")
	}

	return ss[0], nil
}
