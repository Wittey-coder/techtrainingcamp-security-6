package session

type Session interface {
	Set(key string, value interface{}) error // 在Session里修改/添加一个键值对
	Get(key string) (interface{}, error)     // 根据键获取值
	Delete(key string) error                 // 删除此键
	Save() error                             // 保存Session中的键值对
}
