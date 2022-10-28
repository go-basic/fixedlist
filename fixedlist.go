package fixedlist

import (
	"container/list"
	"sync"
)

type FixedList interface {
	Add(interface{})
	Len() int
	Data() []interface{}
}

type fixedList struct {
	sync.RWMutex
	length int
	data   *list.List
}

//创建定长列表
func NewFixedList(len int) FixedList {
	f := &fixedList{}
	f.length = len
	f.data = list.New()
	return f
}

//添加一条记录
func (f *fixedList) Add(val interface{}) {
	f.Lock()
	defer f.Unlock()
	f.data.PushBack(val)
	if f.data.Len() > f.length {
		for i := 0; i <= f.data.Len()-f.length; i++ {
			f.data.Remove(f.data.Front())
		}
	}
}

//获取数据长度
func (f *fixedList) Len() int {
	f.RLock()
	defer f.RUnlock()
	return f.data.Len()
}

//获取数据
func (f *fixedList) Data() []interface{} {
	f.RLock()
	defer f.RUnlock()
	var data []interface{}
	for i := f.data.Front(); i != nil; i = i.Next() {
		data = append(data, i.Value)
	}
	return data
}
