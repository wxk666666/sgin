# 概述
本版本主要实现cors中间件
## 内容
首先，创意一个middleware文件夹，以后我们的中间件将会放在这里面\
创建我们的跨域中间件Cors\

我们了解到，当使用XMLHttpRequest发送请求时，如果浏览器发现违反了同源策略就会自动加上一个请求头 origin；
后端在接受到请求后确定响应后会在 Response Headers 中加入一个属性 Access-Control-Allow-Origin；
浏览器判断响应中的 Access-Control-Allow-Origin 值是否和当前的地址相同，匹配成功后才继续响应处理，否则报错\

缺点：忽略 cookie，浏览器版本有一定要求

同时，结合项目实际，我们可以使用一个config结构体来存放我们的配置，这里可以使用建造者模式进行灵活的管理
所以cors中，我们也添加与config相同的字段\
Config:
``` GO
type Config struct {
AllowOrigins []string
AllowMethods []string
AllowHeaders []string
//如果想要让客户端可以访问到其他的标头，服务器必须将它们在 Access-Control-Expose-Headers 里面列出来.eg:"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires
ExposeHeaders []string
//这个响应头表示 preflight request （预检请求）的返回结果（即 Access-Control-Allow-Methods 和Access-Control-Allow-Headers 提供的信息）可以被缓存多久。
AccessControlMaxAge string
//告知浏览器是否可以将对请求的响应暴露给前端 JavaScript 代码。
AccessControlAllowCredentials bool
}
  ```
Cors:
 ```Go
package sgin8
type Cors struct {
AllowOrigins []string
AllowMethods []string
AllowHeaders []string
//如果想要让客户端可以访问到其他的标头，服务器必须将它们在 Access-Control-Expose-Headers 里面列出来.eg:"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires
ExposeHeaders []string
//这个响应头表示 preflight request （预检请求）的返回结果（即 Access-Control-Allow-Methods 和Access-Control-Allow-Headers 提供的信息）可以被缓存多久。
AccessControlMaxAge string
//告知浏览器是否可以将对请求的响应暴露给前端 JavaScript 代码。
AccessControlAllowCredentials bool
}
 ```
当然，我们的Config需要一些方法来控制字段的值
 ``` GO
func (c *Config) AddOrigins(origins ...string) *Config {
c.AllowOrigins = append(c.AllowOrigins, origins...)
return c
}
func (c *Config) AddMethods(methods ...string) *Config {
c.AllowMethods = append(c.AllowMethods, methods...)
return c
}
func (c *Config) AddHeaders(headers ...string) *Config {
c.AllowHeaders = append(c.AllowHeaders, headers...)
return c
}
func (c *Config) AddExposeHeaders(exposeHeaders ...string) *Config {
c.ExposeHeaders = append(c.ExposeHeaders, exposeHeaders...)
return c
}
func (c *Config) SetAccessControlMaxAge(ms string) *Config {
c.AccessControlMaxAge = ms
return c
}
func (c *Config) SetAccessControlAllowCredentials(isAllow bool) *Config {
c.AccessControlAllowCredentials = isAllow
return c
}
 ```
同时，为了方便，我们可以给出一个方法来设置默认的情况
 ``` GO
func DefaultConfig() *Config {
return &Config{
AllowOrigins:                  []string{"*"},
AllowMethods:                  []string{"POST", "POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE"},
AllowHeaders:                  []string{"Authorization", "Content-Length", "X-CSRF-Token", "Token", "session", "X_Requested_With", "Accept", "Origin", "Host", "Connection", "Accept-Encoding", "Accept-Language", "DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Pragma"},
ExposeHeaders:                 []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "Content-Languge", "Caontent-Type", "Expires", "Last-Modified", "Pragma", "FooBar"},
AccessControlMaxAge:           "200000",
AccessControlAllowCredentials: true,
}
}
 ```

最后，实现将config应用到Cors中的方法和应用Cors的设置的方法 \
将config应用到Cors中的方法：
 ``` GO
func (cors *Cors) SetConfig(c *Config) {
	cors.AccessControlMaxAge = c.AccessControlMaxAge
	cors.ExposeHeaders = c.ExposeHeaders
	cors.AllowOrigins = c.AllowOrigins
	cors.AllowHeaders = c.AllowHeaders
	cors.AllowMethods = c.AllowMethods
	cors.AccessControlAllowCredentials = c.AccessControlAllowCredentials
}
 ```
应用Cors的设置的方法 ：
 ``` GO
func (cors *Cors) Apply() sgin8.HandlerFunc {
	return func(context *sgin8.Context) {
		method := context.Req.Method
		origin := context.Req.Header.Get("Origin")

		if origin != "" {
			context.SetHeader("Access-Control-Allow-Origin", strings.Join(cors.AllowOrigins, ",")) // 设置允许访问所有域
			context.SetHeader("Access-Control-Allow-Methods", strings.Join(cors.AllowMethods, ","))
			context.SetHeader("Access-Control-Allow-Headers", strings.Join(cors.AllowHeaders, ","))
			context.SetHeader("Access-Control-Expose-Headers", strings.Join(cors.ExposeHeaders, ","))
			context.SetHeader("Access-Control-Max-Age", cors.AccessControlMaxAge)
			context.SetHeader("Access-Control-Allow-Credentials", strconv.FormatBool(cors.AccessControlAllowCredentials))
		} else {
			log.Printf("This request haven not set the 'Origin' in Header")
		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		//处理请求
		context.Next()
	}
}

 ```
你也许会问为什么实现的这么繁琐，那么下文将会解释下原因！\
外链：https://zhuanlan.zhihu.com/p/58093669