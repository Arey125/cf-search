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

var columns []string = []string{
	"contest_id",
	"problemset_name",
	"ind",
	"name",
	"type",
	"points",
	"rating",
	"tags",
}

func scanProblem(rows *sql.Rows) (*Problem, error) {
	problem := Problem{}
	var tagString string
    rating := sql.NullInt32{}
	err := rows.Scan(
		&problem.ContestId,
		&problem.ProblemsetName,
		&problem.Index,
		&problem.Name,
		&problem.Type,
		&problem.Points,
		&rating,
		&tagString,
	)
	if tagString != "" {
		problem.Tags = strings.Split(tagString, ", ")
	}
    if rating.Valid {
        r := int(rating.Int32)
        problem.Rating = &r
    }

	if err != nil {
		return nil, err
	}

	return &problem, nil
}

func (m ProblemModel) addBatch(problems []Problem) error {
	q := sq.Insert("problems").
		Columns(columns...)
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
		err := m.addBatch(problems[i*batchSize : min((i+1)*batchSize, count)])
		if err != nil {
			return err
		}
	}

	return nil
}

func (m ProblemModel) GetPage(page int) ([]Problem, error) {
	pageSize := 100
	q := sq.Select(columns...).
		From("problems").
		Limit(uint64(pageSize)).
		Offset(uint64(pageSize * page))

	rows, err := q.RunWith(m.db).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	problems := []Problem{}
	for rows.Next() {
		problem, err := scanProblem(rows)
		if err != nil {
			return nil, err
		}
		problems = append(problems, *problem)
	}

	return problems, nil
}
