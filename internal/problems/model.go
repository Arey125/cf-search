package problems

import (
	"database/sql"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

type Problem struct {
	ContestId      int      `json:"contestId"`
	ProblemsetName string   `json:"problemsetName"`
	Index          string   `json:"index"`
	Name           string   `json:"name"`
	Type           string   `json:"type"`
	Points         float32  `json:"points"`
	Rating         *int     `json:"rating"`
	Tags           []string `json:"tags"`
}

type ProblemModel struct {
	db *sql.DB
}

func NewModel(db *sql.DB) ProblemModel {
	return ProblemModel{db}
}

func (m ProblemModel) addBatch(problems []Problem) error {
	q := sq.Insert("problems").
		Columns("contest_id", "problemset_name", "ind", "name", "type", "points", "rating", "tags")
	for _, p := range problems {

		q = q.Values(
			p.ContestId,
			p.ProblemsetName,
			p.Index,
			p.Name,
			p.Type,
			p.Points,
			p.Rating,
			strings.Join(p.Tags, ", "),
		)
	}
	_, err := q.RunWith(m.db).Exec()

	if err != nil {
		return err
	}
	return nil
}

func (m ProblemModel) AddMany(problems []Problem) error {
	batchSize := 500

	count := len(problems)
	batchCount := (len(problems) + batchSize - 1) / batchSize

	for i := range batchCount {
		m.addBatch(problems[i*batchSize : min((i+1)*batchSize, count)])
	}

	return nil
}
