package inventory

type Service interface {
	CreateItem(item *Item) error
	GetItems() ([]Item, error)
	GetItem(id uint) (*Item, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateItem(item *Item) error {
	return s.repo.Create(item)
}

func (s *service) GetItems() ([]Item, error) {
	return s.repo.GetAll()
}

func (s *service) GetItem(id uint) (*Item, error) {
	return s.repo.GetByID(id)
}
