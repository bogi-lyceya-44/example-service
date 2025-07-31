package models

type Id int64

func NewId(id int64) Id {
	return Id(id)
}

func (id Id) ToInt64() int64 {
	return int64(id)
}

func (id Id) ToUint64() uint64 {
	return uint64(id)
}
