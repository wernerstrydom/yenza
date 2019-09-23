package collections

// ItemList represents a list of ItemValue
type ItemList interface {
  ItemCollection
  // Adds value to this list
  Add(value ItemValue)

  // Adds values to this list
  AddRange(values ...ItemValue)

  // Returns the value at index. If the index is out of bounds, it returns an error
  At(index int) (ItemValue, error)

  // Removes all the items from this list
  Clear()

  // Returns true if this item exists in the list, otherwise false
  Contains(value ItemValue) bool

  // Inserts value at index. If index is out of range, an error is returned
  Insert(index int, value ItemValue) error

  // Removes the item at index. If the index is out of range, an error is returned
  RemoteAt(index int) error

  // Removes value from the list if it exists. If the value does not exists, this becomes a noop
  Remove(value ItemValue)
}
