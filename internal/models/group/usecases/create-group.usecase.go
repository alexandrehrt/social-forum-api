package groupUsecases

import (
	"social-api/internal/entity"
	groupRepository "social-api/internal/models/group/repositories"
)

func Create(group *entity.Group) (*entity.Group, error) {
	group, err := groupRepository.Create(group)
	if err != nil {
		return nil, err
	}

	return group, nil
}
