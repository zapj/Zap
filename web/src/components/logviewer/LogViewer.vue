<template>
  <div class="log-viewer">
    <div id="terminal"></div>
  </div>
</template>
<script setup>
import { FitAddon } from 'xterm-addon-fit'
import { createTerminal, createWebSocket } from '../../zap/terminal/terminal'

import { computed, onMounted, watch, watchEffect } from 'vue'
const props = defineProps({
  logfile: String
})
const logfile = computed(() => props.logfile)
var socket
var terminal
watch(logfile, (newVal, oldVal) => {
  //clean terminal
  terminal.clear()
  socket.send(new TextEncoder().encode('\x02' + JSON.stringify({ logfile: newVal })))
})
onMounted(() => {
  socket = createWebSocket('/api/v1/wss/logviewer?log_file=' + props.logfile)
  terminal = createTerminal('terminal', socket)
  socket.onmessage = (data) => {
    terminal.writeln(data)
  }
  socket.onclose = function (event) {
    terminal.writeln('连接已关闭，请刷新重试')
  }

  socket.onopen = () => {
    terminal._initialized = true
    terminal.focus()
    // terminal.loadAddons(FitAddon)
    // terminal.writeln("Connected to server\n")
  }
})
</script>
<style scoped></style>
