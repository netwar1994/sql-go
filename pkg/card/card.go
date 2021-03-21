package card

import (
	"errors"
	"log"
	"sync"
)

type Card struct {
	Id int64
	Issuer string
	Number string
	Balance int64
	Status string
	OwnerId	int64
	Virtual  bool
}

type Service struct {
	mu sync.RWMutex
	cards []*Card
}

func NewService() *Service {
	return &Service{mu: sync.RWMutex{}, cards: []*Card{}}
}

func (s *Service) All() []*Card {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cards
}

func (s *Service) checkUserId(userId int64) error {
	for _, c := range s.cards {
		log.Println(c)
		if c.Id == userId {
			return nil
		}
	}
	return errors.New("user not found")
}