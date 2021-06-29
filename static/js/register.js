let get_validate_code = new Vue({
    el: "#get_validate_code",
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
        }
    }
})