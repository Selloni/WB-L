package storage

import (
	"fmt"
	"time"
)

// RequestData exported
type RequestData struct {
	User     string `json:"user_id"`
	DateJSON string `json:"date"`
	Info     string `json:"info"`
	DataTime time.Time
}

// UserEvent exported
type UserEvent struct {
	Text string
	Date time.Time
}

// Cash exported
type Cash struct {
	Date map[string]UserEvent
}

// NewCash constructor
func NewCash() *Cash {
	return &Cash{
		Date: make(map[string]UserEvent),
	}
}

// Add called in the package handler in file event
func (c *Cash) Add(data *RequestData) error {
	if v, ok := c.Date[data.User]; ok {
		if v.Date == data.DataTime {
			return fmt.Errorf("На дату %v у пользователя %s мероприятие создано", data.DataTime, data.User)
		}
	}
	tmp := UserEvent{}
	tmp.Date = data.DataTime
	tmp.Text = data.Info
	c.Date[data.User] = tmp
	fmt.Println(c)
	return nil
}

// Update called in the package handler in file event
func (c *Cash) Update(data *RequestData) error {
	if v, ok := c.Date[data.User]; ok {
		if data.DataTime == v.Date {
			tmp := UserEvent{}
			tmp.Date = data.DataTime
			tmp.Text = data.Info
			c.Date[data.User] = tmp
			return nil
		}
	}
	return fmt.Errorf("не найдены данные для обновления")
}

// Delete called in the package handler in file event
func (c *Cash) Delete(data *RequestData) error {
	if v, ok := c.Date[data.User]; ok {
		if v.Date == data.DataTime {
			delete(c.Date, data.User)
			return nil
		}
		return fmt.Errorf("не найдена дата %v для удаления", data.DataTime)
	}
	return fmt.Errorf("не найден пользовател %s для удаления", data.User)
}

// FindDayEvent called in the package handler in file event
func (c *Cash) FindDayEvent(user string, date time.Time) ([]UserEvent, error) {
	result := make([]UserEvent, 0, 10)
	if v, ok := c.Date[user]; ok {
		if v.Date.Day() == date.Day() &&
			v.Date.Year() == date.Year() && v.Date.Month() == date.Month() {
			result = append(result, v)
		}
	} else {
		return nil, fmt.Errorf("не найден пользователь %v", user)
	}
	return result, nil
}

// FindWeekEvent called in the package handler in file event
func (c *Cash) FindWeekEvent(user string, date time.Time) ([]UserEvent, error) {
	result := make([]UserEvent, 0, 10)
	if v, ok := c.Date[user]; ok {
		weekStart := date.AddDate(0, 0, -int(date.Weekday()))
		weekEnd := date.AddDate(0, 0, 6)
		if v.Date.After(weekStart) && v.Date.Before(weekEnd) {
			result = append(result, v)
		}

	} else {
		return nil, fmt.Errorf("не найден пользователь %v", user)
	}
	return result, nil
}

// FindMonthEvent called in the package handler in file event
func (c *Cash) FindMonthEvent(user string, date time.Time) ([]UserEvent, error) {
	result := make([]UserEvent, 0, 10)
	if v, ok := c.Date[user]; ok {
		if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() {
			result = append(result, v)
		}
	} else {
		return nil, fmt.Errorf("не найден пользователь %v", user)
	}
	return result, nil
}
