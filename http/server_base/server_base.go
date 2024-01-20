package server_base

import (
	"net/http"
	"regexp"
)

// パスパラメータを表現
type RouterParameter map[string]string

// パスパラメータを渡す関数type
type routerFunction func(RouterParameter, http.ResponseWriter, *http.Request)

/*
 * what: ルーティングのためのtype
 * why: ルーティングには
 *  method: メソッド名(GET, POST, etc...)
 *  macher: リクエストのパターンマッチ
 *  function: パスパラメータと関数の呼び出し
 * が必要
 */
type routerItem struct {
	method   string
	matcher  *regexp.Regexp
	function routerFunction
}

/*
 * サーバ起動の際にここに対してルーティングを追加する
 */
type Router struct {
	routingItems []routerItem
}

/*
 * regist GET
 */
func (rt *Router) GET(prefix string, function routerFunction) {
	rt.routingItems = append(rt.routingItems, routerItem{
		method:   http.MethodGet,
		matcher:  regexp.MustCompile(prefix),
		function: function,
	})
}

/*
 * regist POST
 */
func (rt *Router) POST(prefix string, function routerFunction) {
	rt.routingItems = append(rt.routingItems, routerItem{
		method:   http.MethodPost,
		matcher:  regexp.MustCompile(prefix),
		function: function,
	})
}

/*
 * match request to registered route
 * このメソッドを実装していないとhttp.ListenAndServerにhandlerとして渡せません
 */
func (rt *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, registedItem := range rt.routingItems {
		// skip unmatched items
		if registedItem.method != request.Method ||
			!registedItem.matcher.MatchString(request.RequestURI) {
			continue
		}
		// extract parameter
		match := registedItem.matcher.FindStringSubmatch(request.RequestURI)
		parameter := make(RouterParameter)
		for i, name := range registedItem.matcher.SubexpNames() { // Ex. SubexpNames retruns name when the case ^/(?P<name>\w+)$
			parameter[name] = match[i]
		}
		registedItem.function(parameter, writer, request)
		return
	}
	http.NotFound(writer, request)
}
