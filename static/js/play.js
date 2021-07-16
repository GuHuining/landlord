const 	PLAY  = 0      // 出牌
const 	DEAL  = 1      // 发牌
const   START = 2      // 开始
const   JOIN  = 3      // 加入
const   DATA  = 4      // 传输信息
const   QUIT  = 5      // 退出
const   OK    = 6      // 成功
const   ERROR = 7      // 错误

// 连接房间
class Connector {
    constructor() {
        this.conn = null
        if (!window["WebSocket"]) {
            alert("你的浏览器不支持WebSocket，无法进行游戏")
        }
    }
    create_room() {
        if (this.conn !== null) {
            alert("你已在游戏中")
            return
        }
        this.conn = new WebSocket("ws://" + document.location.host + "/api/play/create_room")
        this.conn.onclose = function (evt) {
            this.conn = null
            alert("退出房间")
        }
        let t = this
        this.conn.onmessage = function (evt) {
            let data = JSON.parse(evt.data)
            switch (data.type) {
                case OK: break;
                default: console.log(data)
            }
        }
        this.conn.onopen = function (evt) {
            connector.conn.send(JSON.stringify({
                password: create_room_frame.password
            }))
        }
        // 传输密码


    }
    join(data) {
        console.log(data)
    }
}

let connector = new Connector()
