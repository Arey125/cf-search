package main

import (
	"encoding/json"
	"io"
	"net/http"

	_ "github.com/joho/godotenv/autoload"

	"cf-search/internal/problems"
)

type Result struct {
	Problems []problems.Problem `json:"problems"`
}

type Data struct {
	Result Result `json:"result"`
}

func parse(raw []byte) ([]problems.Problem, error) {
	data := Data{}
	err := json.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}
	return data.Result.Problems, nil
}

func fetchProblems() ([]problems.Problem, error) {
	resp, err := http.Get("https://codeforces.com/api/problemset.problems")
	if err != nil {
		return nil, err
	}

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return parse(rawData)
}
