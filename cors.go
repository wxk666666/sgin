package SGin

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

// CorsConfig 使用了建造者模式
type CorsConfig struct {
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

func (c *CorsConfig) AddOrigins(origins ...string) *CorsConfig {
	c.AllowOrigins = append(c.AllowOrigins, origins...)
	return c
}
func (c *CorsConfig) AddMethods(methods ...string) *CorsConfig {
	c.AllowMethods = append(c.AllowMethods, methods...)
	return c
}
func (c *CorsConfig) AddHeaders(headers ...string) *CorsConfig {
	c.AllowHeaders = append(c.AllowHeaders, headers...)
	return c
}
func (c *CorsConfig) AddExposeHeaders(exposeHeaders ...string) *CorsConfig {
	c.ExposeHeaders = append(c.ExposeHeaders, exposeHeaders...)
	return c
}
func (c *CorsConfig) SetAccessControlMaxAge(ms string) *CorsConfig {
	c.AccessControlMaxAge = ms
	return c
}
func (c *CorsConfig) SetAccessControlAllowCredentials(isAllow bool) *CorsConfig {
	c.AccessControlAllowCredentials = isAllow
	return c
}
func (c *CorsConfig) Build() HandlerFunc {
	cors := &Cors{}
	cors.SetCorsConfig(c)
	return cors.Apply()
}
func DefaultCorsConfig() *CorsConfig {
	return &CorsConfig{
		AllowOrigins:                  []string{"*"},
		AllowMethods:                  []string{"POST", "POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE"},
		AllowHeaders:                  []string{"Authorization", "Content-Length", "X-CSRF-Token", "Token", "session", "X_Requested_With", "Accept", "Origin", "Host", "Connection", "Accept-Encoding", "Accept-Language", "DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Pragma"},
		ExposeHeaders:                 []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "Content-Languge", "Caontent-Type", "Expires", "Last-Modified", "Pragma", "FooBar"},
		AccessControlMaxAge:           "200000",
		AccessControlAllowCredentials: true,
	}
}

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

func (cors *Cors) SetCorsConfig(c *CorsConfig) {
	cors.AccessControlMaxAge = c.AccessControlMaxAge
	cors.ExposeHeaders = c.ExposeHeaders
	cors.AllowOrigins = c.AllowOrigins
	cors.AllowHeaders = c.AllowHeaders
	cors.AllowMethods = c.AllowMethods
	cors.AccessControlAllowCredentials = c.AccessControlAllowCredentials
}
func (cors *Cors) Apply() HandlerFunc {
	return func(context *Context) {
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
