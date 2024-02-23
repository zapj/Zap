<template>
  <div>
    Terminal
    <div id="terminal"></div>
  </div>
</template>

<style></style>
<script setup>
import { onMounted } from 'vue'

import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import { AttachAddon } from 'xterm-addon-attach'
import { WebLinksAddon } from 'xterm-addon-web-links'
import 'xterm/css/xterm.css'

function ab2str(buf) {
  return String.fromCharCode.apply(null, new Uint8Array(buf))
}
onMounted(() => {
  const term = new Terminal({
    screenKeys: true,
    useStyle: true,
    cursorBlink: true
  })
  term.open(document.getElementById('terminal'))
  let wsProto = 'ws'
  if (location.protocol === 'https:') {
    wsProto = 'wss'
  }

  const socket = new WebSocket(wsProto + '://' + location.host + '/api/v1/local/ws', [
    sessionStorage.getItem('access_token')
  ])
  // socket.binaryType = 'arraybuffer'
  const attachAddon = new AttachAddon(socket)
  const fitAddon = new FitAddon()
  term.loadAddon(fitAddon)
  term.loadAddon(new WebLinksAddon())
  socket.onclose = function (event) {
    term.write('连接已关闭，请刷新重试\n')
  }

  socket.onopen = function () {
    term.loadAddon(attachAddon)
    term._initialized = true
    term.focus()
    setTimeout(function () {
      fitAddon.fit()
    })
    term.onResize(function (evt) {
      socket.send(
        new TextEncoder().encode('\x01' + JSON.stringify({ cols: evt.cols, rows: evt.rows }))
      )
    })
    term.onTitleChange(function (event) {
      console.log(event)
    })
    term.onData(function (data) {
      socket.send(new TextEncoder().encode('\x00' + data))
    })
    window.onresize = function () {
      fitAddon.fit()
    }
  }
})
</script>
