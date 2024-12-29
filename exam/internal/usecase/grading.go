package usecase

import (
	"context"
	"encoding/json"
	"exam/config"
	"exam/domain"
	"exam/internal/delivery/kafka"
	"exam/internal/repository"
	"log"
)

type GradingUseCase interface {
	CalcGrade(ctx context.Context, student domain.Student) (value domain.Student, err error)
	GetStudent(ctx context.Context, name string) (value domain.Student, err error)
}

func NewGradingUseCase(cfg *config.Config, db repository.DatabaseRepository, redis repository.RedisRepository) GradingUseCase {
	return gradingUseCase{cfg, db, redis}
}

type gradingUseCase struct {
	cfg   *config.Config
	db    repository.DatabaseRepository
	redis repository.RedisRepository
}

func (g gradingUseCase) CalcGrade(ctx context.Context, student domain.Student) (value domain.Student, err error) {

	if student.Grade >= 80 {
		student.Status = "A"
	} else if student.Grade >= 70 {
		student.Status = "B"
	} else {
		student.Status = "C"
	}
	studentJSON, _ := json.Marshal(student)
	err = kafka.PushCommentToQueue("test", string(studentJSON))
	_, err = g.db.CalcGrade(ctx, student)
	if err != nil {
		return domain.Student{}, err
	}
	return student, nil
}
func (g gradingUseCase) GetStudent(ctx context.Context, name string) (value domain.Student, err error) {
	var student domain.Student

	student, err = g.db.GetStudent(ctx, name)
	log.Println("tes2", student)
	err = kafka.PushCommentToQueue("test", name)
	log.Println("tes3", err)
	if err != nil {
		return domain.Student{}, err
	}
	return student, nil
}
