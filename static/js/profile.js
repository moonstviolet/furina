import {createApp} from 'https://cdn.bootcdn.net/ajax/libs/petite-vue/0.4.1/petite-vue.es.min.js'

function fetchData(update, recalculate) {
    const xhr = new XMLHttpRequest();
    xhr.open('POST', '/user/profile', false);
    xhr.setRequestHeader('Content-Type', 'application/json;charset=utf8');
    xhr.send(JSON.stringify({Update: update, Recalculate: recalculate}));
    if (xhr.status !== 200) {
        return null
    }
    const data = JSON.parse(xhr.responseText)
    console.log(data)
    return data
}

createApp({
    user: Object((fetchData(false, false).Data || {}).User),
    updateMsg: "",
    updateList: [],
    update() {
        const data = fetchData(true, false)
        if (data.Code !== 0) {
            console.log(data.Code, data.Msg)
            return
        }
        this.updateMsg = data.Data.UpdateMsg
        if (this.updateMsg === "获取角色面板数据成功") {
            this.user = data.Data.User
            this.updateList = data.Data.UpdateList
        }
    },
    recalculate() {
        const data = fetchData(false, true)
        if (data.Code !== 0) {
            console.log(data.Code, data.Msg)
            return
        }
        this.updateMsg = data.Data.UpdateMsg
    },
}).mount()