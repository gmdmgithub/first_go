package project

// Project -main struct
type Project struct {
	ID   int64
	Name string
	Description string
}

// Repository - abstract the repository to make it independent (ie Mongo, postgrsql, mysql)
type Repository interface {
	// FindAll() ([]Project, error)
	Store(Project) (Project, error)
	// Delete(Project) error
}
