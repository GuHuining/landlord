let login_frame = new Vue({
    template: `
<div class="bg" v-show="show">
    <div class="login_frame">
        <div class="login_block">
            <label for="login_username">用户名</label>
            <input id="login_username" v-model="username">
        </div>
        <div class="login_block">
            <label for="login_password">密码</label>
            <input id="login_password" type="password" v-model="password">
        </div>
        <div class="login_button">
            <button id="login" @click="login">登录</button>
            <span @click="to_register">注册</span>
        </div>
    </div>
</div>

`,
    el: "#v_login_frame",
    data: {
        show: false,
        username: "",
        password: ""
    },
    methods: {
        login: function () {
            let vm = this
            this.username = this.username.trim()
            this.password = this.password.trim()
            if (this.username.length < 6 || this.username.length > 20) {
                alert("用户名长度应在6-20之间")
                return
            }
            if (this.password.length < 6 || this.password.length > 20) {
                alert("密码长度应在6-20之间")
                return
            }
            let data = {
                username: this.username,
                password: this.password,
            }
            axios.post("/api/user/login", data)
                .then(function (response) {
                    vm.show = false
                    // 检查是否绑定昵称
                    axios.post("/api/user/login_check", {})
                        .then(function (response) {
                            if (response.data.nickname === "") {
                                nickname_frame.show = true
                            } else {
                                index_choices.show = true
                            }
                        })
                        .catch(function (error) {
                            alert("登录失败")
                            vm.show = true
                        })
                })
                .catch(function (error) {
                    alert(error.response.data.err)
                })
        },
        to_register: function () {
            this.show = false
            register_frame.show = true
        }
    }
})

let nickname_frame = new Vue({
    template: `
<div class="bg" v-show="show">
    <div class="nickname_frame">
        <div class="nickname_block">
            <label for="input_nickname">昵称(20字符以内)</label>
            <input id="input_nickname" maxlength="20" v-model="nickname">
            <button @click="commit">确认</button>
        </div>
        
    </div>
</div>
`,
    el: "#v_nickname_frame",
    data: {
        show: false,
        nickname: ""
    },
    methods: {
        commit: function () {
            this.nickname = this.nickname.trim()
            if (this.nickname === "") {
                alert("请输入昵称")
                return
            }
            let data = {
                nickname: this.nickname
            }
            axios.post("/api/user/bind_nickname", data)
                .then(function (response) {
                    nickname_frame.show = false
                    index_choices.show = true
                })
                .catch(function (error) {
                    alert(error.response.data.err)
                })
        }
    }
})
