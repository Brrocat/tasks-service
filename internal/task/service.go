package task

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(t *Task) (*Task, error) {
	return s.repo.CreateTask(t)
}

func (s *Service) GetTaskByID(id uint) (*Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *Service) UpdateTask(t *Task) (*Task, error) {
	return s.repo.UpdateTask(t)
}

func (s *Service) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}

func (s *Service) ListTasks(limit, offset int) ([]Task, error) {
	return s.repo.ListTasks(limit, offset)
}

func (s *Service) ListUserTasks(userID uint, limit, offset int) ([]Task, error) {
	return s.repo.ListUserTasks(userID, limit, offset)
}
