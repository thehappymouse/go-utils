package rabbitmq

type ReceiveCall func([]byte) bool

// Receiver 观察者模式需要的接口
// 观察者用于接收指定的queue到来的数据
type Receiver interface {
	GetBindInfo() BindInfo	// 获取绑定信息
	OnError(error)         // 处理遇到的错误，当RabbitMQ对象发生了错误，他需要告诉接收者处理错误
	OnReceive([]byte) bool // 处理收到的消息, 这里需要告知RabbitMQ对象消息是否处理成功
}
