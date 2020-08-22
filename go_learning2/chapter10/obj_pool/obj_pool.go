package obj_pool

// 对象池的一些定义
import (
	"errors"
	"time"
)

// 统一定义错误
var TimeOutError = errors.New("time out")
var OverFlowError = errors.New("overflow")

// 定义对象池中存放的对象
type ReusableObj struct {

}

// 对象池(建议使用不同的池来缓存不同的对象，并不建议池中缓存空接口来存放任何对象，这样在取值时还需要去判断对象的类型)
type ObjPool struct {
	bufChan chan *ReusableObj // 用于缓冲可重用对象
}

// 创建新的对象池
func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		// 往池中预置ReusableObj对象
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

// 定义ObjPool对象的获取对象方法
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
    select {
	case ret := <-p.bufChan:
		return ret,nil
	case <- time.After(timeout): // 超时控制
	    return nil, TimeOutError
	}
}

// 定义往ObjPool放置对象的方法
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default: // 当放不进去channel时，上面的语句不会阻塞，此时直接执行default，返回异常
		return OverFlowError
	}
}
