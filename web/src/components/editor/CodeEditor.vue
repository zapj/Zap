<template>
    <div :id="id" style="height: 400px;width: 100%;overflow: auto;">
        <template v-if="showEmpty" >
            <el-empty description="文件类型不支持，无法打开"  />
        </template>
    </div>
</template>
<script setup>
import { onMounted,inject } from 'vue'
import {EditorView, basicSetup} from "codemirror"
import {javascript} from "@codemirror/lang-javascript"
import apiRequest from '../../httpclient/client';
import { ElMessage } from 'element-plus';
const dialogVisible = inject("dialogVisible")
let editInstance = null
const showEmpty = ref(true)
onMounted(()=>{
    apiRequest({
        url:"/v1/filemanager/fetch",
        method:"post",
        data:{filename:props.filename},
        dataType:"form",
    }).then((resp)=>{
        if (resp.code !== 0) {
            showEmpty.value = true
            return
        }
        showEmpty.value = false
        editInstance = new EditorView({
            doc:resp.data,
            extensions: [basicSetup, javascript()],
            parent: document.getElementById(props.id)
            })
    }).catch(()=>{
        showEmpty.value = true
    })
    
})

const props = defineProps({
    filename:String,
    id:String,
})
const saveFile = ()=>{
 
    apiRequest({
        url:"/v1/filemanager/putfile",
        method:"post",
        data:{filename:props.filename,content:editInstance.state.doc.toString()},
        dataType:"form",
    }).then((resp)=>{
        if(resp.code === 0 ){
            ElMessage({message:resp.msg,type:"success"})
            $emits('closefile',props.filename)
        }else{
            ElMessage({message:resp.msg,type:"error"})
        }
    }).catch((resp)=>{
        ElMessage({message:resp.msg,type:"error"})
    })
}
defineExpose({
    saveFile
})
let $emits = defineEmits(['closefile'])
</script>
<style>
</style>