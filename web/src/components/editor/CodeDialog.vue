<template>
    <el-dialog modal-class="zeditor"  title="文件管理器" draggable :modal="false" :append-to-body="true"  class="te-window" >
    <el-radio-group v-model="curEditFile" @change="handleSelectFile">
      <el-radio-button v-for="(value,key,index) in editorFiles"
    :key="index" :label="key" />
    </el-radio-group>
  
    <CodeEditor v-for="(value,key,index) in editorFiles" :key="index" :filename="key" :id="value.id" v-show="showEditor(value.id)" ></CodeEditor>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="dialogVisible = false"> Confirm </el-button>
      </span>
    </template>
  </el-dialog>
</template>
<script setup>
import { onMounted,defineExpose,defineAsyncComponent, computed, reactive } from 'vue'
const CodeEditor = defineAsyncComponent(()=>import('../../components/editor/CodeEditor.vue'))
var editorFiles = reactive({})
onMounted(()=>{
    
})
const curEditFile = ref()
const showEditor = computed((id)=>{
    return function(id){
        return activeFile.value === id
    }
})
const activeFile = ref()
const handleCloseFile = (file) => {

}
const handleSelectFile = (value) => {
    activeFile.value = editorFiles[value].id
}

const openFile = (file) => {
    if (!editorFiles.hasOwnProperty(file)){
        const eID = genEditID()
        activeFile.value = eID
        editorFiles[file] = {
            name : file,
            change:false,
            id: eID
        }
        curEditFile.value = file
    }else{
        activeFile.value = editorFiles[file].id
        curEditFile.value = file
    }
    // console.log(editorFiles);
    // console.log(activeFile.value);
}
// defineProps({
//     openDialog
// })
defineExpose({
    openFile
})

const genEditID = () => {
    return 'Zap' + Math.random().toString(36).substring(2,4) + Date.now() 
}
</script>