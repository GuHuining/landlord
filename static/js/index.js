let index_choices = new Vue({
    template: `
<div class="bg" v-show="show">
    <div class="choice_frame">
        <button class="choice" id="create">创建房间</button>
        <button class="choice" id="join">加入房间</button>
        <button class="choice" id="random_join">随机加入</button>
    </div>
</div>
`,
    el: "#v_index_choices",
    data: {
        show: false
    },
})