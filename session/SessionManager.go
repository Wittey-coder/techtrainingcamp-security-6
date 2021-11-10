package session

type SessionManager interface {
	Init(addr string, options ...string) error                     // 初始化
	CreateSession() (session Session, sessionId string, err error) // 添加一个新Session
	Get(sessionId string) (session *Session, err error)            // 通过ID获取Session
	Save() error                                                   // 持久化
	Delete(sessionId string) error                                 // 删除Session
}
