package algorithm

import (
	"container/list"
)

type entry struct {
	key, val int
}

type MyHashMap struct {
	data []list.List
}

func constructor() MyHashMap {
	return MyHashMap{make([]list.List, 769)}
}

func (m *MyHashMap) hash(key int) int {
	return key % 769
}

func (m *MyHashMap) Put(key int, value int) {
	h := m.hash(key)
	front := m.data[h].Front()
	for e := front; e != nil; e = e.Next() {
		if ee := e.Value.(entry); ee.key == key {
			e.Value = entry{key, value}
			return
		}
	}

	m.data[h].PushBack(entry{key, value})
}

func (m *MyHashMap) Get(key int) int {
	h := m.hash(key)
	front := m.data[h].Front()
	for e := front; e != nil; e = e.Next() {
		if ee := e.Value.(entry); ee.key == key {
			return ee.val
		}
	}

	return -1
}

func (m *MyHashMap) Remove(key int) {
	h := m.hash(key)
	front := m.data[h].Front()
	for e := front; e != nil; e = e.Next() {
		if ee := e.Value.(entry); ee.key == key {
			m.data[h].Remove(e)
		}
	}
}
