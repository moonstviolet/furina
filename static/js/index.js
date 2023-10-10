import {createApp} from 'https://cdn.bootcdn.net/ajax/libs/petite-vue/0.4.1/petite-vue.es.min.js'


createApp({
    uid: "",
    msg: "",
    login() {
        if (this.uid.length !== 9) {
            this.msg = "不合法的UID"
            return
        }
        const xhr = new XMLHttpRequest();
        xhr.open('POST', '/login', false);
        xhr.setRequestHeader('Content-Type', 'application/json;charset=utf8');
        console.log(this.uid)
        xhr.send(JSON.stringify({Uid: this.uid}));
        if (xhr.status !== 200) {
            return null
        }
        const data = JSON.parse(xhr.responseText)
        if (data.Code === 0) {
            location.reload()
            return
        }
        console.log(data.Msg)
        this.msg = data.Msg
    }
}).mount()