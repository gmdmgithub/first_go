package project

type Project struct {
	ID   int64
	Name string
}
type Repository interface {
	// FindAll() ([]Project, error)
	Store(Project) (Project, error)
	// Delete(Project) error
}
