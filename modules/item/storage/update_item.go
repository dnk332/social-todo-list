package storage

import (
	"context"
	"social-todo-list/modules/item/model"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdated *model.TodoItemUpdate) error {
	if err := s.db.Where(cond).Updates(dataUpdated).Error; err != nil {
		return err
	}

	return nil
}
