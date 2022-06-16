package models

import (
	"github.com/ManuelFresnedaLlamas/GymTonic/common"
	"github.com/google/uuid"
)

const (
	TranslationTable = "translations"

	TranslationLang    = "lang"
	TranslationText    = "text"
	TranslationContext = "context"
)

const (
	Spanish = "es"
	English = "en"
)

func (Translation) TableName() string {
	return TranslationTable
}

type Translation struct {
	ID          uuid.UUID `json:"id" db:"pk,id"`
	Lang        string    `json:"lang" db:"lang"`
	Text        string    `json:"text" db:"text"`
	Translation string    `json:"translation" db:"translation"`
	Context     string    `json:"context" db:"context"`
}

type TranslationQuery struct {
	ID      uuid.UUID `json:"id"`
	Lang    string    `json:"lang"`
	Text    string    `json:"text"`
	Context []string  `json:"context"`
	common.BaseQuery
}

type TranslationList struct {
	Items []Translation `json:"items"`
	Count int64         `json:"count"`
}
