package remote_package

import (
	"testing"
	// 导入github的包, cm是起别名
	cm "github.com/easierway/concurrent_map"
)

// 测试导入github远程线程安全的map包
func TestCurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
