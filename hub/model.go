package hub

const (
	PLAY  = iota // 出牌
	DEAL         // 发牌
	START        // 开始
	QUIT         // 退出
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

