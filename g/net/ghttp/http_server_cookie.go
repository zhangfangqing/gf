// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.
//
// HTTP Cookie管理对象，
// 由于Cookie是和HTTP请求挂钩的，因此被包含到 ghttp 包中进行管理

package ghttp

import (
    "sync"
    "time"
    "strings"
    "net/http"
    "gitee.com/johng/gf/g/os/gtime"
    "gitee.com/johng/gf/g/container/gmap"
)

const (
    gDEFAULT_PATH    = "/"           // 默认path
    gDEFAULT_MAX_AGE = 86400*365     // 默认cookie有效期(一年)
    SESSION_ID_NAME  = "gfsessionid" // 默认存放Cookie中的SessionId名称
)

// cookie对象
type Cookie struct {
    mu       sync.RWMutex          // 并发安全互斥锁
    data     map[string]CookieItem // 数据项
    domain   string                // 默认的cookie域名
    request  *Request              // 所属HTTP请求对象
    response *Response             // 所属HTTP返回对象
}

// cookie项
type CookieItem struct {
    value  string
    domain string // 有效域名
    path   string // 有效路径
    expire int    // 过期时间
}

// 包含所有当前服务器正在服务的Cookie(每个请求一个Cookie对象)
var cookies = gmap.NewUintInterfaceMap()

// 创建一个cookie对象，与传入的请求对应
func NewCookie(r *Request) *Cookie {
    if r := GetCookie(r.Id); r != nil {
        return r
    }
    c := &Cookie {
        data     : make(map[string]CookieItem),
        domain   : defaultDomain(r),
        request  : r,
        response : r.Response,
    }
    c.init()
    cookies.Set(uint(r.Id), c)
    return c
}

// 获取一个已经存在的Cookie对象
func GetCookie(requestid int) *Cookie {
    if r := cookies.Get(uint(requestid)); r != nil {
        return r.(*Cookie)
    }
    return nil
}

// 获取默认的domain参数
func defaultDomain(r *Request) string {
    return strings.Split(r.Host, ":")[0]
}

// 从请求流中初始化
func (c *Cookie) init() {
    c.mu.Lock()
    defer c.mu.Unlock()
    for _, v := range c.request.Cookies() {
        c.data[v.Name] = CookieItem {
            v.Value, v.Domain, v.Path, v.Expires.Second(),
        }
    }
}

// 获取SessionId
func (c *Cookie) SessionId() string {
    v := c.Get(SESSION_ID_NAME)
    if v == "" {
        v = makeSessionId()
        c.SetSessionId(v)
    }
    return v
}

// 设置SessionId
func (c *Cookie) SetSessionId(sessionid string)  {
    c.Set(SESSION_ID_NAME, sessionid)
}

// 设置cookie，使用默认参数
func (c *Cookie) Set(key, value string) {
    c.SetCookie(key, value, c.domain, gDEFAULT_PATH, gDEFAULT_MAX_AGE)
}

// 设置cookie，带详细cookie参数
func (c *Cookie) SetCookie(key, value, domain, path string, maxage int) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = CookieItem {
        value, domain, path, int(gtime.Second()) + maxage,
    }
}

// 查询cookie
func (c *Cookie) Get(key string) string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    if r, ok := c.data[key]; ok {
        if r.expire >= 0 {
            return r.value
        } else {
            return ""
        }
    }
    return ""
}

// 标记该cookie在对应的域名和路径失效
// 删除cookie的重点是需要通知浏览器客户端cookie已过期
func (c *Cookie) Remove(key, domain, path string) {
    c.SetCookie(key, "", domain, path, -86400)
}

// 请求完毕后删除已经存在的Cookie对象
func (c *Cookie) Close() {
    cookies.Remove(uint(c.request.Id))
}

// 输出到客户端
func (c *Cookie) Output() {
    c.mu.RLock()
    defer c.mu.RUnlock()
    for k, v := range c.data {
        if v.expire == 0 {
            continue
        }
        http.SetCookie(
            c.response.ResponseWriter,
            &http.Cookie {
                Name    : k,
                Value   : v.value,
                Domain  : v.domain,
                Path    : v.path,
                Expires : time.Unix(int64(v.expire), 0),
            },
        )
    }
}
