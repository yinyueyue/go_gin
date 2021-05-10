package filter

import "net/http"

type WebFilterHandler func(w http.ResponseWriter, r *http.Request) error

type  WebFilter struct {
	filterMap map[string]WebFilterHandler
}

func NewFilter() *WebFilter  {
	return &WebFilter{filterMap:make(map[string]WebFilterHandler)}
}

// RegisterFilterUri 注册拦截器
func (filter *WebFilter) RegisterFilterUri(uri string, handler WebFilterHandler) {
	filter.filterMap[uri] = handler
}

// GetFilterHandler 根据Uri获取相应的handle
func (filter *WebFilter) GetFilterHandler(uri string)WebFilterHandler {
	return filter.filterMap[uri]
}

// WebHandler 声明新的函数类型
type WebHandler func(w http.ResponseWriter, r *http.Request)

// Handle 执行拦截器，返回函数类型
func (filter *WebFilter) Handle(webHandler WebHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for path, handler := range filter.filterMap {
			if path == r.RequestURI {
				// 执行拦截业务逻辑
				err := handler(w, r)
				if err != nil {
					_, _ = w.Write([]byte(err.Error()))
					return
				}
				break
			}
		}
		webHandler(w, r)
	}
}
