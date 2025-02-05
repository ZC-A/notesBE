package eventbus

import (
	eb "github.com/asaskevich/EventBus"
)

// 初始化时直接赋值
// 没有在init方法中初始化，是防止在各个模块init中调用有顺序的依赖
var EventBus = eb.New()

const (
	EventSignalConfigPreParse  = "sys:config:pre-parse"
	EventSignalConfigPostParse = "sys:config:post-parse"
)
