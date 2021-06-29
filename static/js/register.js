let register_frame = new Vue({
    template: `
<div class="bg" v-show="show">
    <div class="register_frame">
        <div class="register_block">
            <label for="register_username">用户名</label>
            <input id="register_username" placeholder="长度在6-20位之间" v-model="username">
        </div>
        <div class="register_block">
            <label for="register_password">密码</label>
            <input id="register_password" placeholder="长度在6-20位之间" type="password" v-model="password">
        </div>
        <div class="register_block">
            <label for="register_password_again">重复密码</label>
            <input id="register_password_again" placeholder="长度在6-20位之间" type="password" v-model="password_again">
        </div>
        <div class="register_block">
            <label for="register_email">邮箱</label>
            <input id="register_email" v-model="email">
        </div>
        <div class="register_block">
            <label for="register_validate_code">验证码</label>
            <input id="register_validate_code" v-model="validate_code">
        </div>
        <div id="get_validate_code" @click="get_code">获取验证码</div>
        <div class="register_button">
            <button id="register" @click="register">注册</button>
            <span @click="to_login">登录</span>
        </div>
    </div>
</div>
`,
    el: "#v_register_frame",
    data: {
        show: false,
        username: "",
        password: "",
        password_again: "",
        email: "",
        validate_code: "",
    },
    methods: {
        get_code: function () {
            let email = document.getElementById("register_email")
            if (!email.value.match(/.+@.+\..+/)) {
                alert("请输入正确格式的邮箱")
                return
            }
            let data = {
                email: email.value
            }
            axios.post("/api/user/validate_code", data)
                .catch(function (error) {
                    alert(error.response.data.err)
                })
        },
        to_login: function () {
            this.show = false
            login_frame.show = true
        },
        register: function () {
            if (this.username.length < 6 || this.username.length > 20) {
                alert("用户名应在6-20位之间")
                return
            }
            if (this.password.length < 6 || this.password.length > 20) {
                alert("密码应在6-20位之间")
                return
            }
            if (!this.email.match(/.+@.+\..+/)) {
                alert("请输入正确格式的邮箱")
                return
            }
            if (this.password !== this.password_again) {
                alert("两次输入的密码不匹配")
                return
            }
            if (this.validate_code.length !== 6) {
                alert("请输入正确的验证码")
                return
            }
            let vm = this
            let data = {
                username: this.username,
                password: this.password,
                email: this.email,
                validate_code: this.validate_code
            }
            axios.post("/api/user/register", data)
                .then(function (response) {
                    register_frame.show = false
                    login_frame.show = true
                })
                .catch(function (error) {
                    alert(error.response.data.err)
                })
        }
    }
})