package repository

import (
	"github.com/nekizz/telegram-bot/internal/model"
)

type ModelListRepository[T any, query model.ListQuery] interface {
	List(query, ...string) ([]*T, int64, error)
}

type ModelReadRepository[T any, id uint] interface {
	Read(id, ...string) (*T, error)
}

type ModelReadByConditionRepository[T any] interface {
	ReadByCondition(interface{}, ...string) (*T, error)
}

type ModelInsertRepository[T any] interface {
	Insert(*T) (*T, error)
}

type ModelUpdateRepository[T any] interface {
	Update(*T) (*T, error)
}

type ModelDeleteRepository[T any] interface {
	Delete(*T) error
}
