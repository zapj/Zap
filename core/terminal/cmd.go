package terminal

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"unsafe"

	"github.com/creack/pty"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type windowSize struct {
	Rows uint16 `json:"rows"`
	Cols uint16 `json:"cols"`
	X    uint16
	Y    uint16
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Subprotocols:    []string{"token"},
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandlerLocalWS(c *gin.Context) {
	w := c.Writer
	r := c.Request
	h := http.Header{}
	for _, access_token := range websocket.Subprotocols(r) {
		h.Set("Sec-Websocket-Protocol", access_token)
		break
	}

	conn, err := upgrader.Upgrade(w, r, h)
	if err != nil {
		slog.Error("Unable to upgrade connection", slog.Any("err", err))
		return
	}
	// conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	// conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	cmd := exec.Command("/bin/bash", "-l")
	// cmd.SysProcAttr = &syscall.SysProcAttr{GidMappingsEnableSetgroups: true}
	// cmd.SysProcAttr.Credential = &syscall.Credential{Uid: 1000, Gid: 1000, NoSetGroups: true}
	//xterm-256color
	//xterm
	cmd.Env = append(os.Environ(), "TERM=xterm-256color")

	tty, err := pty.Start(cmd)
	if err != nil {
		slog.Error("Unable to start pty/cmd", slog.Any("err", err))
		conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}

	defer func() {
		cmd.Process.Kill()
		cmd.Process.Wait()
		tty.Close()
		conn.Close()
	}()

	go func() {
		for {
			buf := make([]byte, 1024)
			read, err := tty.Read(buf)
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
				slog.Error("Unable to read from pty/cmd", slog.Any("err", err))
				return
			}
			conn.WriteMessage(websocket.BinaryMessage, buf[:read])
		}
	}()

	for {
		messageType, reader, err := conn.NextReader()
		if err != nil {
			slog.Error("Unable to grab next reader", slog.Any("err", err))
			return
		}

		if messageType == websocket.TextMessage {
			// l.Warn("Unexpected text message")
			// conn.WriteMessage(websocket.TextMessage, []byte("Unexpected text message"))
			continue
		}

		dataTypeBuf := make([]byte, 1)
		read, err := reader.Read(dataTypeBuf)
		if err != nil {
			slog.Error("Unable to read message type from reader", slog.Any("err", err))
			conn.WriteMessage(websocket.TextMessage, []byte("Unable to read message type from reader"))
			return
		}

		if read != 1 {
			slog.Error("Unexpected number of bytes read", slog.Any("err", err), slog.Any("bytes", read))
			return
		}

		switch dataTypeBuf[0] {
		case 0:
			// buf:= bufio.NewReader(reader)

			copied, err := io.Copy(tty, reader)
			if err != nil {
				slog.Error("Error after copying ", "bytes", copied, "err", err)
			}
		case 1:
			decoder := json.NewDecoder(reader)
			resizeMessage := windowSize{}
			err := decoder.Decode(&resizeMessage)
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Error decoding resize message: "+err.Error()))
				continue
			}
			slog.Info("Resizing terminal", "resize", resizeMessage)
			_, _, errno := syscall.Syscall(
				syscall.SYS_IOCTL,
				tty.Fd(),
				syscall.TIOCSWINSZ,
				uintptr(unsafe.Pointer(&resizeMessage)),
			)
			if errno != 0 {
				slog.Info("Unable to resize terminal", "errno", syscall.Errno(errno))
			}
		default:
			slog.Info("Unknown data type", "dataType", dataTypeBuf[0])
		}
	}

}
