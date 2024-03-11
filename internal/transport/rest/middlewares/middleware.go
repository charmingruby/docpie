package middlewares

import (
	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/sirupsen/logrus"
)

type Middleware struct {
	logger                *logrus.Logger
	membersRepository     collections.CollectionMembersRepository
	collectionsRepository collections.CollectionsRepository
}

func NewMiddleware(
	logger *logrus.Logger,
	membersRepository collections.CollectionMembersRepository,
	collectionsRepository collections.CollectionsRepository,
) *Middleware {
	return &Middleware{
		logger:                logger,
		membersRepository:     membersRepository,
		collectionsRepository: collectionsRepository,
	}
}
