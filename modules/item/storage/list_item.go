package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (s *sqlStore) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TodoItem, error) {
	var result []model.TodoItem

	db := s.db.Where("status <> ?", "Deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.
		Table(model.TodoItem{}.TableName()).
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Count(&paging.Total).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
