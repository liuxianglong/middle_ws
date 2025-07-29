package consts

const (
	WSHeaderLengthSize = 4
	WSMaxHeaderSize    = 1024 //头部最大1KB
	// WSMaxMessageSize 内容最大消息长度
	WSMaxMessageSize = 10 * 1024 * 1024
	// WSWriteWait 允许等待的最长写入时间
	WSWriteWait = 3
	// WSPongWait pong消息等待时长
	WSPongWait = 60
	// WSReadBufferSize 读消息缓冲
	WSReadBufferSize = 1024
	// WSWriteBufferSize 写消息缓冲
	WSWriteBufferSize = 1024
	// WSNeedStopGoroutineNum 缓冲值
	WSNeedStopGoroutineNum = 2
)
