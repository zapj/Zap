package user

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/zapj/zap/assist/global"
)

func CreateUser(username, password string) error {
	return nil
}

func CheckUserExists(username string) bool {
	users, err := Parse()
	if err != nil {
		global.LOG.Error("读取passwd失败")
		return false
	}
	if _, ok := users[username]; ok {
		return true
	}
	return false
}

func WriteToPasswd(users map[string]UserEntry) {

}

type UserEntry struct {
	Username string
	Password string
	Uid      string
	Gid      string
	Gecos    string
	Home     string
	Shell    string
}

// Parse opens the '/etc/passwd' file and parses it into a map from usernames
// to Entries
func Parse() (map[string]UserEntry, error) {
	return ParseFile("/etc/passwd")
}

func ParseFile(path string) (map[string]UserEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return ParseReader(file)
}

func ParseReader(r io.Reader) (map[string]UserEntry, error) {
	lines := bufio.NewReader(r)
	entries := make(map[string]UserEntry)
	for {
		line, _, err := lines.ReadLine()
		if err != nil {
			break
		}
		name, entry, err := parseLine(string(line))
		if err != nil {
			return nil, err
		}
		entries[name] = entry
	}
	return entries, nil
}

func parseLine(line string) (string, UserEntry, error) {
	fs := strings.Split(line, ":")
	if len(fs) != 7 {
		return "", UserEntry{}, errors.New("unexpected number of fields in /etc/passwd")
	}
	return fs[0], UserEntry{
		Username: fs[0],
		Password: fs[1],
		Uid:      fs[2],
		Gid:      fs[3],
		Gecos:    fs[4],
		Home:     fs[5],
		Shell:    fs[6],
	}, nil
}

func copyBytes(x []byte) []byte {
	y := make([]byte, len(x))
	copy(y, x)
	return y
}
