import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import { AttachAddon } from 'xterm-addon-attach'
import { WebLinksAddon } from 'xterm-addon-web-links'
import 'xterm/css/xterm.css'

export const ab2str = function (buf) {
  return String.fromCharCode.apply(null, new Uint8Array(buf))
}

export const createTerminal = function (container, socket) {
  const terminal = new Terminal({
    screenKeys: true,
    useStyle: true,
    cursorBlink: true
  })
  const fitAddon = new FitAddon()
  terminal.loadAddon(fitAddon)
  terminal.loadAddon(new WebLinksAddon())
  terminal.loadAddon(new AttachAddon(socket))
  terminal.open(document.getElementById(container))
  terminal.onResize(function (evt) {
    socket.send(
      new TextEncoder().encode('\x01' + JSON.stringify({ cols: evt.cols, rows: evt.rows }))
    )
    terminal.fit()
  })
  return terminal
}

export const resizeTerminal = function (terminal, socket) {
  terminal.fit()
  socket.send(
    JSON.stringify({
      type: 'resize',
      rows: terminal.rows,
      cols: terminal.cols
    })
  )
}

export const resizeTerminalOnResize = function (terminal, socket) {
  window.addEventListener('resize', () => {
    resizeTerminal(terminal, socket)
  })
}

export const resizeTerminalOnLoad = function (terminal, socket) {
  resizeTerminal(terminal, socket)
  resizeTerminalOnResize(terminal, socket)
}

export const resizeTerminalOnLoadAndResize = function (terminal, socket) {
  resizeTerminalOnLoad(terminal, socket)
  resizeTerminalOnResize(terminal, socket)
}

export const resizeTerminalOnLoadAndResizeAndKeydown = function (terminal, socket) {
  resizeTerminalOnLoadAndResize(terminal, socket)
  window.addEventListener('keydown', (event) => {
    if (event.key === 'Escape') {
      socket.send(
        JSON.stringify({
          type: 'resize',
          rows: terminal.rows,
          cols: terminal.cols
        })
      )
    }
  })
}

export const createWebSocket = function (url) {
  let wsProto = 'ws'
  if (location.protocol === 'https:') {
    wsProto = 'wss'
  }
  return new WebSocket(wsProto + '://' + location.host + url, [
    sessionStorage.getItem('access_token')
  ])
}

export const createTerminalAndSocket = function (container, url) {
  const socket = createWebSocket(url)
  createTerminal(container, socket)
  resizeTerminalOnLoadAndResizeAndKeydown(terminal, socket)
  return {
    terminal,
    socket
  }
}
