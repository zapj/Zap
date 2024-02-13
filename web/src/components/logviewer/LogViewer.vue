<template>
    <div class="log-viewer">
        <div id="terminal"></div>
    </div>
</template>
<script setup>

import { createTerminal,createWebSocket } from '../../zap/terminal/terminal';



import { onMounted } from 'vue'

onMounted(() => {
    const socket = createWebSocket("/api/v1/wss/logviewer?log_file=install_32.log")
    const terminal = createTerminal("terminal",socket)
    socket.onmessage = (data) => {
        terminal.writeln(data)
    }
    socket.onclose = function (event) {
        terminal.writeln('连接已关闭，请刷新重试')
    }
    
    socket.onopen = () => {
        terminal._initialized = true
        terminal.focus()
        // terminal.writeln("Connected to server\n")
    }
})

</script>
<style scoped>

</style>