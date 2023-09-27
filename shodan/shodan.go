package shodan

//shodan URL被定义未一个常量值
const BaseURL = "https://api.shodan.io"

//定义结构体Client，该结构体用于维护请求中的API令牌
type Client struct {
	apiKey string
}

//定义一个辅助函数，以API令牌作为输入，创建并返回一个初始化的实例Client
func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
