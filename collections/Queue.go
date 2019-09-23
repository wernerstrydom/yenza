package collections

import (
  "errors"
)

type ItemQueue struct {
  items []ItemValue
}

func NewItemQueue() *ItemQueue {
  return &ItemQueue{}
}

func (q *ItemQueue) Push(value ItemValue) {
  q.items = append(q.items, value)
}

func (q *ItemQueue) Pop() (ItemValue, error) {
  s := q
  if len(s.items) < 1 {
    return nil, errors.New("q is empty")
  }
  result := s.items[len(s.items)-1]
  s.items = s.items[:len(s.items)-1]
  return result, nil
}

func (q *ItemQueue) Peek() (ItemValue, error) {
  s := q
  if len(s.items) < 1 {
    return nil, errors.New("q is empty")
  }
  result := s.items[len(s.items)-1]
  return result, nil
}

func (q *ItemQueue) Count() int {
  return len(q.items)
}

func (q *ItemQueue) IsEmpty() bool {
  return len(q.items) > 0
}


