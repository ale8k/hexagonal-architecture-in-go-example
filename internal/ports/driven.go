package ports

type PersistencePort interface {
	Save(any)
}
