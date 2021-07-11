package hub

const (
	PLAY  = iota // 出牌
	DEAL         // 发牌
	START        // 开始
	JOIN         // 加入
	DATA         // 传输信息
	QUIT         // 退出
	OK           // 成功
	ERROR        // 错误
)

type Request struct {
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}

type Response struct {
	Type    int         `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Seat 座位结构
type Seat struct {
	SeatID   int    `json:"seat_id"`
	NickName string `json:"nick_name"`
}

// SeatsData 座位信息
type SeatsData struct {
	Seats []Seat `json:"seats"`
}
