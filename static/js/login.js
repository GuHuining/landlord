let login_button = new Vue({
    el: "#login",
    methods: {
        login: function login() {
            let username = document.getElementById("username")
            let password = document.getElementById("password")

            if (username.value.length < 6 || username.value.length > 20) {
                alert("用户名长度应在6-20之间")
                return
            }
            if (password.value.length < 6 || password.value.length > 20) {
                alert("用户名长度应在6-20之间")
                return
            }
            let data = {
                username: username.value,
                password: password.value,
            }
            axios.post("/api/user/login", data)
                .then(function (response) {
                    alert("登陆成功")
                })
                .catch(function (error) {
                    alert(error.response.data.err)
                })
        }
    }
})

