package user

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

func CheckUserExists(username string) bool {
	users, err := ReadSystemUsers()
	if err != nil {
		slog.Error("读取passwd失败")
		return false
	}
	if _, ok := users[username]; ok {
		return true
	}
	return false
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

type ShadowEntry struct {
	Username string

	Password string `json:"-"`

	// Days since Jan 1, 1970 password was last changed.
	LastChange string

	// -1 if password aging is disabled.
	MinPasswordAge string

	// 用户必须更改密码的天数。
	//
	// -1 is password aging is disabled.
	MaxPasswordAge string

	// 密码到期前的天数（请参阅上面的密码最长使用期限），在此期间应向用户发出警告。
	//
	// -1 is password aging is disabled.
	WarnPeriod string

	// 密码过期后的天数（请参阅上面的密码最长使用期限），在此期间密码仍应被接受。
	//
	// -1 is password aging is disabled.
	InactivityPeriod string

	// 帐户的到期日期，表示为自1970年1月1日以来的天数。
	//
	// -1 is account never expires.
	AccountExpiry string

	// Unused now.
	Flags string
}

// vboxadd:!:19720::::::

func (u *ShadowEntry) ToString() string {
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%s:%s\n", u.Username, u.Password,
		fmt.Sprint(u.LastChange),
		fmt.Sprint(u.MinPasswordAge),
		fmt.Sprint(u.MaxPasswordAge),
		fmt.Sprint(u.WarnPeriod),
		fmt.Sprint(u.InactivityPeriod),
		u.GetAccountExpiry(),
		u.GetFlags())
}

func (u *ShadowEntry) GetAccountExpiry() string {
	if u.AccountExpiry == "-1" || u.AccountExpiry == "0" {
		return ""
	}
	return fmt.Sprint(u.AccountExpiry)
}

func (u *ShadowEntry) GetFlags() string {
	if u.Flags == "0" {
		return ""
	}
	return fmt.Sprint(u.Flags)
}

func (u *UserEntry) ToString() string {
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s\n", u.Username, u.Password, u.Uid, u.Gid, u.Gecos, u.Home, u.Shell)
}

// Parse opens the '/etc/passwd' file and parses it into a map from usernames
// to Entries
func ReadSystemUsers() (map[string]UserEntry, error) {
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

func parseShadowLine(line string) (string, ShadowEntry, error) {
	fs := strings.Split(line, ":")
	if len(fs) != 9 {
		return "", ShadowEntry{}, errors.New("unexpected number of fields in /etc/shadow")
	}
	return fs[0], ShadowEntry{
		Username:         fs[0],
		Password:         fs[1],
		LastChange:       fs[2],
		MinPasswordAge:   fs[3],
		MaxPasswordAge:   fs[4],
		WarnPeriod:       fs[5],
		InactivityPeriod: fs[6],
		AccountExpiry:    fs[7],
		Flags:            fs[8],
	}, nil
}

func ReadSystemShadow() (map[string]ShadowEntry, error) {
	return ReadShadow("/etc/shadow")
}

func ReadShadow(path string) (map[string]ShadowEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	lines := bufio.NewReader(file)
	entries := make(map[string]ShadowEntry)
	for {
		line, _, err := lines.ReadLine()
		if err != nil {
			break
		}
		name, entry, err := parseShadowLine(string(line))
		if err != nil {
			return nil, err
		}
		entries[name] = entry
	}
	return entries, nil
}
