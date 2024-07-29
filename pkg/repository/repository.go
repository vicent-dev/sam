package repository

type Entity interface{}

type Field struct {
	Column string
	Value  interface{}
}

var (
	rs map[string]interface{}
)

type Repository[T Entity] interface {
	Find(id int) (*T, error)
	FindWithRelations(id int) (*T, error)
	FindBy(fs ...Field) ([]T, error)
	FindByWithRelations(fs ...Field) ([]T, error)
	FindFirstBy(fs ...Field) (*T, error)
	CreateBulk(ts []T) error
	Create(t *T) error
	Update(t *T, fs ...Field) error
	Delete(t *T) error
}
