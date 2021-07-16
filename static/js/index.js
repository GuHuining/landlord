let index_choices = new Vue({
    template: `
<div class="bg" v-show="show">
    <div class="choice_frame">
        <button class="choice" id="create" v-on:click="create_room">创建房间</button>
        <button class="choice" id="join">加入房间</button>
        <button class="choice" id="random_join">随机加入</button>
    </div>
</div>
`,
    el: "#v_index_choices",
    data: {
        show: false
    },
    methods: {
        create_room: function () {
            create_room_frame.show = true
        }
    }
})

let create_room_frame = new Vue({
    el: "#v_create_room_frame",
    template: `
<div class="bg" v-show="show">
    <div class="password_frame">
        <label for="create_room_password">房间密码</label>
        <input placeholder="房间密码" id="create_room_password" maxlength="20" v-model="password">
        <button v-on:click="create_room()">创建</button>
    </div>
</div>
`,
    data: {
        show: false,
        password: "",
    },
    methods: {
        create_room: function () {
            connector.create_room()
            this.show = false
        }
    }
})