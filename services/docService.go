package services

import (
	"notepad/models"
)

type DocService struct {
	*BaseService[models.Doc]
}

func NewDocService() *DocService {
	return &DocService{}
}

func (s *DocService) GetDocTree() []models.Doc {
	var list []models.Doc = []models.Doc{}
	var docMap = make(map[string]models.Doc)
	for _, doc := range s.baseList {
		docMap[doc.Id] = doc
	}
	for _, doc := range s.baseList {
		if doc.ParentId != "" {
			parentDoc, ok := docMap[doc.ParentId]
			if ok {
				parentDoc.Children = append(parentDoc.Children, doc)
			}
		} else {
			list = append(list, doc)
		}
	}
	return list
}
