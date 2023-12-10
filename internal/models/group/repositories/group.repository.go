package groupRepository

import (
	"social-api/config"
	"social-api/internal/entity"
)

func Create(group *entity.Group) (*entity.Group, error) {
	if err := config.DB.Create(group).Error; err != nil {
		return nil, err
	}

	return group, nil
}
