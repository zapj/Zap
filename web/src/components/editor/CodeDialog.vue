<template>
    <el-dialog modal-class="zeditor"  title="文件管理器" draggable :modal="false" :append-to-body="true" :width="dialogWidth" class="te-window" >
    <!-- <el-radio-group v-model="curEditFile" @change="handleSelectFile">
      <el-radio-button v-for="(value,key,index) in editorFiles"
    :key="index" :label="value.filename" :title="key" />
    </el-radio-group> -->
    <el-tag
    v-for="(value,key,index) in editorFiles"
    :key="index"
    class="mx-1"
    closable
    :title="key"
    :disable-transitions="false"
    @close="handleCloseFile(key)"
    type="success" size="large"
  >
  <el-link :underline="false" @click="handleSelectFile(key)">{{ value.filename }}</el-link> 
  </el-tag>
  <div class="edit-toolbars"></div>
  <Suspense>
    <CodeEditor v-for="(value,key,index) in editorFiles" :ref="el=>{editorRefs[key]=el}" :key="key" :filename="key" :id="value.id" v-show="showEditor(value.id)" 
        @closefile="closeFile" v-on="$listeners"></CodeEditor>
    <template #fallback>
        <el-skeleton :rows="5" />
    </template>
  </Suspense>
    
    <template #footer>
      <span class="dialog-footer">
        <!-- <el-button @click="dialogVisible = false">Cancel</el-button> -->
        <el-button type="primary" @click="saveFile"> 保存 </el-button>
      </span>
    </template>
  </el-dialog>
</template>
<script setup>
import { ElMessage } from 'element-plus';
import { onMounted,defineAsyncComponent, computed, reactive,inject } from 'vue'
const CodeEditor = defineAsyncComponent(()=>import('./CodeEditor.vue'))
var editorFiles = reactive({})
var idFileMaps = reactive({})
var editorRefs = reactive({})
const dialogWidth = ref("70%")
const curEditFile = ref()
const activeFile = ref()
const dialogVisible = inject("dialogVisible")
onMounted(()=>{
    window.addEventListener('resize',()=>{
        if (document.body.clientWidth < 768){
            dialogWidth.value="90%"
        }
    })
    // console.log(props.dialogVisible);
})
const showEditor = computed((id)=>{
    return function(id){
        return activeFile.value === id
    }
})

const handleCloseFile = (file) => {
    if (editorFiles.hasOwnProperty(file)){
        delete idFileMaps[editorFiles[file].id] 
        delete editorFiles[file]

       const keys = Object.keys(editorFiles);
       if (keys.length >= 1){
        const lastKey = keys[keys.length - 1];
        activeFile.value = editorFiles[lastKey].id;
       }
       if(keys.length === 0){
        dialogVisible.value = false
       }
    }

}
const handleSelectFile = (file) => {
    activeFile.value = editorFiles[file].id
    curEditFile.value = file
}

const openFile = (file,filename) => {
    // $emits("update:modelValue", true);
    if (!editorFiles.hasOwnProperty(file)){
        const eID = genEditID()
        activeFile.value = eID
        editorFiles[file] = {
            name : file,
            filename:filename,
            change:false,
            id: eID
        }
        idFileMaps[eID] = file
        curEditFile.value = file
    }else{
        activeFile.value = editorFiles[file].id
        curEditFile.value = file
    }
    // console.log(editorFiles);
    // console.log(activeFile.value);
}
const closeFile = (key) => {
    console.log(key);
    handleCloseFile(key)
}

// const props = defineProps({
//     dialogVisible : true
// })

defineExpose({
    openFile
})
// let $emits = defineEmits(['update:modelValue'])
const saveFile = ()=>{
    if (!editorFiles.hasOwnProperty(curEditFile.value)){
       ElMessage({message:"无法保存",type:"error"})
       return
    }
    console.log(curEditFile.value);
    editorRefs[curEditFile.value].saveFile()
    // editFile = editorFiles[curEditFile.value]
    // editFile.id
}
const genEditID = () => {
    return 'Zap' + Math.random().toString(36).substring(2,4) + Date.now() 
}
</script>
<style>
.te-window > .el-dialog__body{
    padding: 0 !important;
}
.te-window .el-dialog__header{
    background-color: #333;
    padding:10px !important;
    /* height: 30px; */
    margin-right: 0;
    line-height: 30px;
    
}
.te-window .el-dialog__header .el-dialog__title {
    color:white !important;
}

@media (max-width: 768px){
    .te-window > .el-dialog__body{
        padding: 0 !important;
    }   
}
</style>