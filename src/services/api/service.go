package services

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type Service interface {
	Health() bool
	InjectServices(*logrus.Entry, []Service)
	Init() error
	Execute(*sync.WaitGroup) error
}
