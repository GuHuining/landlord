package hub

import (
	"context"
	"errors"
	"fmt"
	"landlord/log"
	"sync"
)

type Room struct {
	ID         int
	Password   string
	Players    []*Player
	NewPlayer  chan *Player
	PlayerExit chan *Player
	State      int
	Before     *Room
	Next       *Room
	Mu         sync.Mutex
}

const (
	EMPTY   = iota // 空房间
	WAITING        // 等待开始
	RUNNING        // 正在运行
)

func (room *Room) New() {
	room.Players = make([]*Player, 3)
	room.NewPlayer = make(chan *Player, 1)
	room.PlayerExit = make(chan *Player, 1)
	room.State = EMPTY
}

// Destroy 将房间恢复原状
func (room *Room) Destroy() {
	room.State = EMPTY
	room.Password = ""
	room.Players = make([]*Player, 3)
}

// Run 进行房间内情况的处理
func (room *Room) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	for {
		select {
		case player := <-room.NewPlayer:
			room.Join(ctx, player) // 有人加入房间
		case player := <-room.PlayerExit:
			room.Exit(cancel, player)
		}
	}
}

// Join 处理有人加入的情况
func (room *Room) Join(ctx context.Context, player *Player) {
	room.Mu.Lock()
	defer room.Mu.Unlock()
	if room.State == RUNNING { // 当游戏已经开始
		player.Conn.WriteJSON(Response{ERROR, "游戏已开始", nil})
		player.Conn.Close()
	} else if room.State == EMPTY { // 房间还未分配
		player.Conn.WriteJSON(Response{ERROR, "房间不存在", nil})
		player.Conn.Close()
	} else {
		if len(room.Players) == 3 { // 房间人数已满
			player.Conn.WriteJSON(Response{ERROR, "该房间已满", nil})
			player.Conn.Close()
			return
		}
		// 选择一个空座位
		var newSeat Seat  // 新
		seatsData := room.generateSeatData()
		for i := 0; i < 3; i++ {
			if room.Players[i] == nil {
				newSeat = Seat{i, player.Nickname}
				writeErr := player.Conn.WriteJSON(Response{DATA, "", newSeat})
				writeErr = player.Conn.WriteJSON(Response{DATA, "", seatsData})
				if writeErr != nil {
					return
				}
				room.Players[i] = player
				break
			}
		}

		// 向其他人发送加入信息
		for i := 0; i < 3; i++ {
			if room.Players[i] != nil && room.Players[i].UserID != player.UserID {
				player.Conn.WriteJSON(Response{JOIN, player.Nickname+"加入了房间", newSeat})
			}
		}
		// TODO 若人满则开始游戏
	}

}

// generateSeatData 获取座位信息
func (room *Room) generateSeatData() SeatsData {
	seats := SeatsData{make([]Seat, 3)}
	for i := 0; i < 3; i++ {
		if room.Players[i] != nil {
			seats.Seats[i] = Seat{i, room.Players[i].Nickname}
		}
	}
	return seats
}

// Exit 处理有人退出的情况
func (room *Room) Exit(cancel context.CancelFunc, player *Player) {
	room.Mu.Lock()
	defer room.Mu.Unlock()
	if room.State == WAITING {
		if len(room.Players) == 1 { // 最后一人退出则将此房间放回空房间列表
			if room.Password == "" { // 处理无密码的房间
				_, err := RoomWithoutPassword.PopByID(room.ID)
				if err != nil {
					log.MyLog.Printf("exit: 释放房间失败")
					return
				}
			} else { // 处理有密码的房间
				_, err := RoomWithPassword.PopByID(room.ID)
				if err != nil {
					log.MyLog.Printf("exit: 释放房间失败")
					return
				}
			}
			room.Destroy()
			EmptyRooms.PushBack(room)
		} else { // 房间内还有其他人
			for i := 0; i < 3; i++ {
				if room.Players[i].UserID == player.UserID { // 将退出的用户剔出用户列表
					room.Players[i] = nil
					break
				}
			}
			for i := 0; i < 3; i++ { // 向其他用户发信
				if room.Players[i] != nil {
					room.Players[i].Conn.WriteJSON(Response{QUIT, fmt.Sprintf("用户%s退出", player.Nickname), nil})
				}
			}
		}
	} else if room.State == RUNNING { // 房间正在游戏时，则解散房间
		for i := 0; i < 3; i++ {
			if room.Players[i].UserID == player.UserID { // 将退出的用户剔出用户列表
				room.Players[i] = nil
				break
			}
		}
		for i := 0; i < 3; i++ { // 向其他用户发信
			if room.Players[i] != nil {
				room.Players[i].Conn.WriteJSON(Response{QUIT, fmt.Sprintf("用户%s退出", player.Nickname), nil})
			}
		}
		if room.Password == "" { // 处理无密码的房间
			_, err := RoomWithoutPasswordPlaying.PopByID(room.ID)
			if err != nil {
				log.MyLog.Printf("exit: 释放房间失败")
				return
			}
		} else { // 处理有密码的房间
			_, err := RoomWithPassword.PopByID(room.ID)
			if err != nil {
				log.MyLog.Printf("exit: 释放房间失败")
				return
			}
		}
		room.Destroy()
		EmptyRooms.PushBack(room)
	}

}

type Rooms struct {
	Head     *Room
	Tail     *Room
	Number   int
	RoomsMap map[int]*Room
	Mu       sync.Mutex
}

// NewRooms 初始化Rooms结构体
func NewRooms() *Rooms {
	var rooms Rooms
	rooms.RoomsMap = make(map[int]*Room)
	return &rooms
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
func (rooms *Rooms) PopByID(id int) (room *Room, err error) {
	rooms.Mu.Lock()
	defer rooms.Mu.Unlock()
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
	delete(rooms.RoomsMap, room.ID)
	return
}

// 按照ID获取一个房间，但不从房间列表中移除
func (rooms *Rooms) GetByID(id int) (room *Room, err error) {
	rooms.Mu.Lock()
	defer rooms.Mu.Unlock()
	room, ok := rooms.RoomsMap[id]
	if !ok {
		err = errors.New("房间不存在")
		return
	}
	return
}
