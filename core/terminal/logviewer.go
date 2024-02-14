package terminal

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zapj/goutils/fileutils"
	"github.com/zapj/zap/core/utils/pathutil"
)

func LogViewerHandler(c *gin.Context) {
	logFile := c.Query("log_file")

	w := c.Writer
	r := c.Request
	h := http.Header{}
	for _, access_token := range websocket.Subprotocols(r) {
		h.Set("Sec-Websocket-Protocol", access_token)
		break
	}

	conn, err := upgrader.Upgrade(w, r, h)
	if err != nil {
		slog.Error("Unable to upgrade connection", "err", err)
		return
	}
	if logFile == "" {
		slog.Error("log file is empty")
		conn.WriteMessage(websocket.TextMessage, []byte("log file is empty"))
		return
	}
	logFile = pathutil.GetPath("data/logs", logFile)
	if !fileutils.IsFile(logFile) {
		slog.Error("log file is not exist", "log file", logFile)
		conn.WriteMessage(websocket.TextMessage, []byte("log file is not exist"))
		return
	}
	logFd, err := os.Open(logFile)
	if err != nil {
		slog.Error("Unable to open log file", "err", err, "log file", logFile)
		return
	}
	logScan := bufio.NewReader(logFd)
	defer func() {
		conn.Close()
		logFd.Close()
	}()

	go func() {

		for {
			line, err := logScan.ReadBytes('\n')
			if err != nil {
				slog.Error("Unable to read log file", "err", err, "line", line)
				if errors.Is(err, io.EOF) {
					if string(line) != "" {
						conn.WriteMessage(websocket.TextMessage, bytes.TrimRight(line, "\n"))
					}
					time.Sleep(2 * time.Second)
					continue
				} else if err != nil {
					slog.Error("Unable to read log file", "err", err)
					return
				}
			}

			err = conn.WriteMessage(websocket.TextMessage, bytes.TrimRight(line, "\n"))
			if err != nil {
				slog.Error("Unable to write message", slog.Any("err", err))
				return
			}
			// conn.WriteMessage(websocket.TextMessage, bytes.TrimRight(line, "\n"))
		}
	}()

	for {
		messageType, reader, err := conn.NextReader()
		if err != nil {
			slog.Error("Unable to grab next reader", slog.Any("err", err))
			return
		}

		if messageType == websocket.TextMessage {
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
		case 2:
			decoder := json.NewDecoder(reader)
			var msg map[string]string
			decoder.Decode(&msg)
			if logfile, ok := msg["logfile"]; ok {
				logFd.Close()
				logFd, _ = os.Open(pathutil.GetPath("data/logs", logfile))
				logScan = bufio.NewReader(logFd)
			}

			// colsBuf := make([]byte, 4)
		}

	}

}
