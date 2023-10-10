import {createApp} from 'https://cdn.bootcdn.net/ajax/libs/petite-vue/0.4.1/petite-vue.es.min.js'

function fetchData() {
    const xhr = new XMLHttpRequest();
    xhr.open('POST', '/character', false);
    xhr.setRequestHeader('Content-Type', 'application/json;charset=utf8');
    const path = location.pathname
    const cid = Number(path.substring(path.lastIndexOf('/') + 1))
    xhr.send(JSON.stringify({Cid: cid}));
    if (xhr.status !== 200) {
        return null
    }
    const data = JSON.parse(xhr.responseText)
    console.log(data)
    return data
}

createApp({
    char: (fetchData().Data || {}).Character,
}).mount()