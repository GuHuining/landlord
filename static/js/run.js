axios.post("/api/user/login_check")
    .then(function (response) {
        index_choices.show = true
    })
    .catch(function (error) {
        login_frame.show = true
    })
