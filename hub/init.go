package hub

import "landlord/config"

func init() {
	initRooms()
}

var (
	RoomWithoutPassword        *Rooms
	RoomWithPassword           *Rooms
	RoomWithPasswordPlaying    *Rooms
	RoomWithoutPasswordPlaying *Rooms
	EmptyRooms                 *Rooms
)

// 初始化房间
func initRooms() {
	EmptyRooms = NewRooms()
	RoomWithPassword = NewRooms()
	RoomWithoutPassword = NewRooms()
	RoomWithPasswordPlaying = NewRooms()
	RoomWithoutPasswordPlaying = NewRooms()

	playConfig := config.GetPlayConfig()

	// 初始化空房间
	for i := playConfig.Rooms.Number; i > 1; i-- {
		var room = Room{
			ID: i,
		}
		room.New()
		go room.Run()
		EmptyRooms.PushBack(&room)
	}
}
