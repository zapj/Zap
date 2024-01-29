<template>
    <div :id="id" style="height: 400px;width: 100%;overflow: auto;"></div>
</template>
<script setup>
import { onMounted } from 'vue'
import {EditorView, basicSetup} from "codemirror"
import {javascript} from "@codemirror/lang-javascript"
import apiRequest from '../../httpclient/client';

let editInstance = null
onMounted(()=>{
    console.log(props.filename);
    apiRequest({
        url:"/v1/filemanager/fetch",
        method:"post",
        data:{filename:props.filename},
        dataType:"form",
    }).then((resp)=>{
        editInstance = new EditorView({
            doc:resp.data,
            extensions: [basicSetup, javascript()],
            parent: document.getElementById(props.id)
            })
    }).catch((resp)=>{
        console.log(resp);
    })
    
})
const openFile = (file) => {
    alert(file)
}

const props = defineProps({
    filename:String,
    id:String,
})

</script>
<style>
</style>