package models

type BaseModel interface {
	GetCreatedAt() string
	GetUpdatedAt() string
	SetCreatedAt(t string)
	SetUpdatedAt(t string)
	GetId() string
	SetId(id string)
}

type BaseModelImpl struct {
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (b *BaseModelImpl) GetCreatedAt() string {
	return b.CreatedAt
}

func (b *BaseModelImpl) GetUpdatedAt() string {
	return b.UpdatedAt
}

func (b *BaseModelImpl) SetCreatedAt(t string) {
	b.CreatedAt = t
}

func (b *BaseModelImpl) SetUpdatedAt(t string) {
	b.UpdatedAt = t
}

func (b *BaseModelImpl) GetId() string {
	return b.Id
}

func (b *BaseModelImpl) SetId(id string) {
	b.Id = id
}
