axios.post("/api/user/login_check", {})
    .then(function (response) {
        if (response.data.nickname === "") {
            nickname_frame.show = true
        } else {
            index_choices.show = true
        }
    })
    .catch(function (error) {
        login_frame.show = true
    })
