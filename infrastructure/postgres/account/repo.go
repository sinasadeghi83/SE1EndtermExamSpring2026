package account

import (
	"gorm.io/gorm"

	accountData "redbank/data/account"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(e accountData.Account) (accountData.Account, error) {
	model := toModel(e)
	if err := r.db.Create(&model).Error; err != nil {
		return accountData.Account{}, err
	}
	return toEntity(model), nil
}

func (r *Repo) Find(id string) (accountData.Account, error) {
	var model Model
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return accountData.Account{}, err
	}
	return toEntity(model), nil
}

func (r *Repo) FindByOwner(ownerID string) ([]accountData.Account, error) {
	var models []Model
	if err := r.db.Find(&models, "owner_id = ?", ownerID).Error; err != nil {
		return nil, err
	}
	entities := make([]accountData.Account, 0, len(models))
	for _, m := range models {
		entities = append(entities, toEntity(m))
	}
	return entities, nil
}

func (r *Repo) Update(e accountData.Account) (accountData.Account, error) {
	model := toModel(e)
	if err := r.db.Save(&model).Error; err != nil {
		return accountData.Account{}, err
	}
	return toEntity(model), nil
}

func (r *Repo) UpdateBalance(id string, delta int64) (accountData.Account, error) {
	if err := r.db.Model(&Model{}).Where("id = ?", id).
		Update("balance", gorm.Expr("balance + ?", delta)).Error; err != nil {
		return accountData.Account{}, err
	}
	return r.Find(id)
}

func (r *Repo) Delete(id string) error {
	return r.db.Delete(&Model{}, "id = ?", id).Error
}
