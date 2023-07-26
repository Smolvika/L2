package repository

import (
	"develop/dev11"
	"develop/dev11/pkg/handler"
	"sync"
	"time"
)

type Cache struct {
	date map[int][]dev11.Event
	mu   *sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		date: make(map[int][]dev11.Event),
		mu:   &sync.RWMutex{},
	}
}

func (c *Cache) CreateEvent(ev dev11.EventQuery) (dev11.EventAnswer, *handler.CustomError) {
	//Получение user по ID, если user не существует - добавляем его
	c.mu.RLock()
	_, ok := c.date[ev.UserId]
	c.mu.RUnlock()
	if !ok {
		c.mu.Lock()
		c.date[ev.UserId] = []dev11.Event{}
		c.mu.Unlock()
	}

	//Добавление события
	c.mu.Lock()
	upUser := c.date[ev.UserId]
	upUser = append(upUser, ev.Event)
	c.mu.Unlock()
	return dev11.EventAnswer{
		UserId:  ev.UserId,
		EventId: len(upUser),
		Event:   ev.Event,
	}, nil
}

func (c *Cache) UpdateEvent(ev dev11.EventAnswer) (dev11.EventAnswer, *handler.CustomError) {
	//Получение user по ID, если user не существует - ошибка
	c.mu.RLock()
	user, ok := c.date[ev.UserId]
	c.mu.RUnlock()
	if !ok {
		return dev11.EventAnswer{}, handler.NewCustomError(400, "Such user is not exist", "Such user is not exist")
	}

	//Добавление события в массив
	if len(user) < ev.EventId {
		return dev11.EventAnswer{}, handler.NewCustomError(400, "Such event is not exist", "Such event is not exist")
	}
	c.mu.Lock()
	user[ev.EventId-1] = ev.Event
	c.mu.Unlock()

	return dev11.EventAnswer{
		UserId:  ev.UserId,
		EventId: ev.EventId,
		Event:   ev.Event,
	}, nil
}

func (c *Cache) DeleteEvent(userId, eventId int) (dev11.EventAnswer, *handler.CustomError) {
	//Получение user по ID, если user не существует - ошибка
	c.mu.RLock()
	user, ok := c.date[userId]
	c.mu.RUnlock()
	if !ok {
		return dev11.EventAnswer{}, handler.NewCustomError(400, "Such user is not exist", "Such user is not exist")
	}

	//Удаление события из массива
	if len(user) < eventId {
		return dev11.EventAnswer{}, handler.NewCustomError(400, "Such event is not exist", "Such event is not exist")
	}
	c.mu.Lock()
	deletedEv := user[eventId-1]
	user = append(user[:eventId-1], user[eventId:]...)
	user = user[:len(user)-1]
	c.mu.Unlock()

	return dev11.EventAnswer{
		UserId:  userId,
		EventId: eventId,
		Event:   deletedEv,
	}, nil
}

func (c *Cache) GetEventsForDuration(date time.Time, userId int, dur time.Duration) ([]dev11.Event, *handler.CustomError) {
	ans := make([]dev11.Event, 0)
	c.mu.RLock()
	user, ok := c.date[userId]
	c.mu.RUnlock()
	if !ok {
		return nil, handler.NewCustomError(400, "Such user is not exist", "Such user is not exist")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, ev := range user {
		if ev.Date.Sub(date) >= 0 && ev.Date.Sub(date.Add(dur)) < 0 {
			ans = append(ans, ev)
		}
	}
	return ans, nil
}
