package loan

import (
	"gorm.io/gorm"

	loanData "redbank/data/loan"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(e loanData.Loan) (loanData.Loan, error) {
	model := toModel(e)
	if err := r.db.Create(&model).Error; err != nil {
		return loanData.Loan{}, err
	}
	return toEntity(model), nil
}

func (r *Repo) Find(id string) (loanData.Loan, error) {
	var model Model
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return loanData.Loan{}, err
	}
	return toEntity(model), nil
}

func (r *Repo) FindByAccount(accountID string) ([]loanData.Loan, error) {
	var models []Model
	if err := r.db.Find(&models, "account_id = ?", accountID).Error; err != nil {
		return nil, err
	}
	entities := make([]loanData.Loan, 0, len(models))
	for _, m := range models {
		entities = append(entities, toEntity(m))
	}
	return entities, nil
}

func (r *Repo) Update(e loanData.Loan) (loanData.Loan, error) {
	model := toModel(e)
	if err := r.db.Save(&model).Error; err != nil {
		return loanData.Loan{}, err
	}
	return toEntity(model), nil
}

func (r *Repo) UpdateStatus(id string, status loanData.Status) (loanData.Loan, error) {
	if err := r.db.Model(&Model{}).Where("id = ?", id).Update("status", string(status)).Error; err != nil {
		return loanData.Loan{}, err
	}
	return r.Find(id)
}

func (r *Repo) Delete(id string) error {
	return r.db.Delete(&Model{}, "id = ?", id).Error
}
