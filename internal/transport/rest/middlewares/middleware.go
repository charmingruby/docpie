package middlewares

import (
	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/sirupsen/logrus"
)

type Middleware struct {
	logger            *logrus.Logger
	membersRepository collections.CollectionMembersRepository
}

func NewMiddleware(logger *logrus.Logger, membersRepository collections.CollectionMembersRepository) *Middleware {
	return &Middleware{
		logger:            logger,
		membersRepository: membersRepository,
	}
}
