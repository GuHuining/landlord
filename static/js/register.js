let register_frame = new Vue({
    template: `
<div class="bg" v-show="show">
    <div class="register_frame">
        <div class="register_block">
            <label for="register_username">用户名</label>
            <input id="register_username">
        </div>
        <div class="register_block">
            <label for="register_password">密码</label>
            <input id="register_password">
        </div>
        <div class="register_block">
            <label for="register_password_again">重复密码</label>
            <input id="register_password_again">
        </div>
        <div class="register_block">
            <label for="register_email">邮箱</label>
            <input id="register_email">
        </div>
        <div class="register_block">
            <label for="register_validate_code">验证码</label>
            <input id="register_validate_code">
        </div>
        <div id="get_validate_code" @click="get_code">获取验证码</div>
        <div class="register_button">
            <button id="register">注册</button>
            <span @click="to_login">登录</span>
        </div>
    </div>
</div>
`,
    el: "#v_register_frame",
    data: {
        show: false
    },
    methods: {
        get_code: function () {
            let email = document.getElementById("email")
            if (!email.value.match(/.*@.*\..*/)) {
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
        }
    }
})