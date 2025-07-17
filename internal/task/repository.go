package task

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateTask(t *Task) (*Task, error) {
	if err := r.db.Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (r *Repository) GetTaskByID(id uint) (*Task, error) {
	var t Task
	if err := r.db.First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *Repository) UpdateTask(t *Task) (*Task, error) {
	if err := r.db.Save(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (r *Repository) DeleteTask(id uint) error {
	return r.db.Delete(&Task{}, id).Error
}

func (r *Repository) ListTasks(limit, offset int) ([]Task, error) {
	var tasks []Task
	if err := r.db.Limit(limit).Offset(offset).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) ListUserTasks(userID uint, limit, offset int) ([]Task, error) {
	var tasks []Task
	if err := r.db.Where("user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}
