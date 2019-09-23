package collections

import "errors"

type ArrayList struct {
  items []ItemValue
}

func NewArrayList() *ArrayList {
  return &ArrayList{}
}

func (t *ArrayList) Count() int {
  return len(t.items)
}

func (t *ArrayList) Add(value ItemValue) {
  t.items = append(t.items, value)
}

func (t *ArrayList) AddRange(values ...ItemValue) {
  t.items = append(t.items, values...)
}

func (t *ArrayList) At(index int) (ItemValue, error) {
  if index < 0 || index >= len(t.items) {
    return nil, errors.New("index out of range")
  }
  return t.items[index], nil
}

func (t *ArrayList) Clear() {
  t.items = []ItemValue{}
}

func (t *ArrayList) Contains(value ItemValue) bool {
  return t.IndexOf(value) >= 0
}

func (t *ArrayList) IndexOf(value ItemValue) int {
  for index, item := range t.items {
    if item == value {
      return index
    }
  }
  return -1
}

func (t *ArrayList) Insert(index int, value ItemValue) error {
  if index < 0 || index > len(t.items) {
    return errors.New("index out of range")
  }
  t.items = append(t.items, 0)
  copy(t.items[index+1:], t.items[index:])
  t.items[index] = value
  return nil
}

func (t *ArrayList) RemoteAt(index int) error {
  if index < 0 || index > len(t.items) {
    return errors.New("index out of range")
  }

  copy(t.items[index:], t.items[index+1:])
  t.items[len(t.items)-1] = nil
  t.items = t.items[:len(t.items)-1]
  return nil
}

func (t *ArrayList) Remove(value ItemValue) {
  index := t.IndexOf(value)
  if index >= 0 {
    _ = t.RemoteAt(index)
  }
}
