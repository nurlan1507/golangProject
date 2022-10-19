package repository

type IRepository[T interface{}] interface {
	Save(x T)
}
