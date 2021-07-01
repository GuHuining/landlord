package hub

import (
	"errors"
	"sync"
)

type Room struct {
	ID         int
	Password   string
	Players    []Player
	NewPlayer  chan Player
	PlayerExit chan int
	Before     *Room
	Next       *Room
	Mu         sync.Mutex
}

func (room *Room) New() {
	room.NewPlayer = make(chan Player)
}

type Rooms struct {
	Head     *Room
	Tail     *Room
	Number   int
	RoomsMap map[int]*Room
	Mu       sync.Mutex
}

// New 初始化Rooms结构体
func (rooms *Rooms) New() {
	rooms.RoomsMap = make(map[int]*Room)
}

// 插入新房间
func (rooms *Rooms) PushBack(room *Room) {
	rooms.Mu.Lock()
	defer rooms.Mu.Unlock()
	if rooms.Head == nil { // 房间列表为空时，将该房间设为头部
		rooms.Head = room
		rooms.Tail = room
		room.Before = nil
		room.Next = nil
	} else {
		rooms.Tail.Next = room
		room.Before = rooms.Tail
		room.Next = nil
		rooms.Tail = room
	}
	rooms.RoomsMap[room.ID] = room
	rooms.Number++
}

// 从头部提取一个房间
func (rooms *Rooms) PopFront() (room *Room, err error) {
	rooms.Mu.Lock()
	defer rooms.Mu.Unlock()
	if rooms.Head == nil { // 房间列表为空，报错
		err = errors.New("已无空余房间")
		return
	} else {
		room = rooms.Head
		rooms.Head = room.Next
		rooms.Head.Before = nil

		room.Before = nil
		room.Next = nil
		delete(rooms.RoomsMap, room.ID)
		rooms.Number--
		return
	}
}

// 按照ID提取一个房间
func (rooms *Rooms) GetByID(id int) (room *Room, err error) {
	room, ok := rooms.RoomsMap[id]
	if !ok {
		err = errors.New("房间不存在")
		return
	}
	// 将该房间从链表中提取出来
	if rooms.Head == room {
		rooms.Head = room.Next
	}
	if rooms.Tail == room {
		rooms.Tail = room.Before
	}
	if room.Before != nil {
		room.Before.Next = room.Next
	}
	if room.Next != nil {
		room.Next.Before = room.Before
	}
	room.Before = nil
	room.Next = nil
	rooms.Number--
	return
}
