package asyncwriter

import "time"

const (
	// 缓冲写，默认缓冲大小。超过此大小，会触发写磁盘
	defaultBufferSize = 256 * 1024

	// 定时刷磁盘的时间间隔
	defaultFlushInterval = 30 * time.Second

	// 异步写日志，异步的 buffer 大小，即异步队列中最多缓存几条数据
	defaultAsyncBufferSize = 16 * 1024
)
