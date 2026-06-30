package transaction

import (
	"gorm.io/gorm"

	transactionData "redbank/data/transaction"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(e transactionData.Transaction) (transactionData.Transaction, error) {
	model := toModel(e)
	if err := r.db.Create(&model).Error; err != nil {
		return transactionData.Transaction{}, err
	}
	return toEntity(model), nil
}

func (r *Repo) Find(id string) (transactionData.Transaction, error) {
	var model Model
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return transactionData.Transaction{}, err
	}
	return toEntity(model), nil
}

func (r *Repo) FindByAccount(accountID string) ([]transactionData.Transaction, error) {
	var models []Model
	if err := r.db.Find(&models, "account_id = ?", accountID).Error; err != nil {
		return nil, err
	}
	entities := make([]transactionData.Transaction, 0, len(models))
	for _, m := range models {
		entities = append(entities, toEntity(m))
	}
	return entities, nil
}

func (r *Repo) UpdateStatus(id string, status transactionData.Status) (transactionData.Transaction, error) {
	if err := r.db.Model(&Model{}).Where("id = ?", id).Update("status", string(status)).Error; err != nil {
		return transactionData.Transaction{}, err
	}
	return r.Find(id)
}
