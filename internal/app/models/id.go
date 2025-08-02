package models

type ID int64

func NewID(id int64) ID {
	return ID(id)
}

func (id ID) ToInt64() int64 {
	return int64(id)
}

func (id ID) ToUint64() uint64 {
	return uint64(id)
}
