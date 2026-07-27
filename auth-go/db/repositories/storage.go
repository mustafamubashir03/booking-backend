package db

type Storage struct {
	UserRepository UserRepository
}

func NewStorage() *Storage {
	storage := Storage{
		UserRepository: &UserRepositoryImp{},
	}
	return &storage
}
