package models

type ZapAccessLog struct {
	ID         int `gorm:"primarykey;autoIncrement"`
	Uid        int `gorm:"index;"`
	RemoteAddr string
	Proto      string
	TLS        string
	Host       string
	Method     string
	RequestUri string
	UserAgent  string
	Referer    string
	RequestID  string
	CreatedAt  int64
}

// msg := AccessLogRecord{
// 	When:       h.opts.Now(),
// 	RemoteAddr: r.RemoteAddr,
// 	Proto:      r.Proto,
// 	TLS:        r.TLS != nil,
// 	Host:       r.Host,
// 	Method:     r.Method,
// 	RequestURI: r.URL.RequestURI(),
// 	UserAgent:  r.UserAgent(),
// 	Referer:    r.Referer(),
// 	RequestID:  RequestIDFromContext(r.Context()),
// }
