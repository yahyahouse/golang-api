package repository

import (
	"context"
	"exam/config"
	"exam/domain"
	"exam/internal/delivery/database"
	"log"
	"time"
)

type DatabaseRepository interface {
	GetStudent(ctx context.Context, name string) (value domain.Student, err error)
	CalcGrade(ctx context.Context, student domain.Student) (value domain.Student, err error)
}

type databaseRepository struct {
	db   *database.DB
	conf *config.DBConfig
}

func NewDatabaseRepository(db *database.DB, conf *config.DBConfig) DatabaseRepository {
	return &databaseRepository{db, conf}
}

func (repo *databaseRepository) GetStudent(ctx context.Context, name string) (value domain.Student, err error) {
	var data domain.Student

	ctxTimeOut, cancle := context.WithTimeout(ctx, 2*time.Second)
	defer cancle()
	err = repo.db.GetDBCon().WithContext(ctxTimeOut).Raw("SELECT * FROM student WHERE name=?", name).First(&data).Error
	log.Println("test", data)
	return data, err
}

func (repo *databaseRepository) CalcGrade(ctx context.Context, student domain.Student) (value domain.Student, err error) {
	var data = domain.Student{}
	query := repo.db.GetDBCon().Exec("INSERT INTO student (name, age, grade, status) VALUES (?,?,?,?)", student.Name, student.Age, student.Grade, student.Status)
	return data, query.Error
}
