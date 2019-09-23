package collections

import (
	"errors"
)

type ItemStack struct {
	items []ItemValue
}

func NewItemStack() *ItemStack {
	return &ItemStack{}
}

func (s *ItemStack) Push(value ItemValue) {
	s.items = append(s.items, value)
}

func (s *ItemStack) Pop() (ItemValue, error) {
	if len(s.items) < 1 {
		return nil, errors.New("s is empty")
	}
	result := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return result, nil
}

func (s *ItemStack) Peek() (ItemValue, error) {
	if len(s.items) < 1 {
		return nil, errors.New("s is empty")
	}
	result := s.items[len(s.items)-1]
	return result, nil
}

func (s *ItemStack) Count() int {
	return len(s.items)
}

func (s *ItemStack) IsEmpty() bool {
	return len(s.items) > 0
}