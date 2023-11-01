//go:build integration_handler
// +build integration_handler

package handler_test

import db2 "homework-6/tests/db"

var (
	db *db2.TDB
)

func init() {
	//запращиваем тестовые креды для бд из енв
	//cfg, err := config.FromEnv()
	db = db2.NewFromEnv()
}