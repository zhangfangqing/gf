// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

package gparser_test

import (
    "bytes"
    "testing"
    "gitee.com/johng/gf/g/encoding/gparser"
    "fmt"
)

func Test_Set1(t *testing.T) {
    e := []byte(`{"k1":{"k11":[1,2,3]},"k2":"v2"}`)
    p := gparser.New(map[string]string{
        "k1" : "v1",
        "k2" : "v2",
    })
    p.Set("k1.k11", []int{1,2,3})
    if c, err := p.ToJson(); err == nil {
        fmt.Println(string(c))
        if bytes.Compare(c, []byte(`{"k1":{"k11":[1,2,3]},"k2":"v2"}`)) != 0 {
            t.Error("expect:", string(e))
        }
    } else {
        t.Error(err)
    }
}

func Test_Set2(t *testing.T) {
    e := []byte(`[[null,1]]`)
    p := gparser.New([]string{"a"})
    p.Set("0.1", 1)
    if c, err := p.ToJson(); err == nil {
        fmt.Println(string(c))
        if bytes.Compare(c, e) != 0 {
            t.Error("expect:", string(e))
        }
    } else {
        t.Error(err)
    }
}

func Test_Set3(t *testing.T) {
    e := []byte(`{"kv":{"k1":"v1"}}`)
    p := gparser.New([]string{"a"})
    p.Set("kv", map[string]string{
        "k1" : "v1",
    })
    if c, err := p.ToJson(); err == nil {
        fmt.Println(string(c))
        if bytes.Compare(c, e) != 0 {
            t.Error("expect:", string(e))
        }
    } else {
        t.Error(err)
    }
}

func Test_Set4(t *testing.T) {
    e := []byte(`["a",[{"k1":"v1"}]]`)
    p := gparser.New([]string{"a"})
    p.Set("1.0", map[string]string{
        "k1" : "v1",
    })
    if c, err := p.ToJson(); err == nil {
        fmt.Println(string(c))
        if bytes.Compare(c, e) != 0 {
            t.Error("expect:", string(e))
        }
    } else {
        t.Error(err)
    }
}

func Test_Set5(t *testing.T) {
    e := []byte(`[[[[[[[[[[[[[[[[[[[[[1,2,3]]]]]]]]]]]]]]]]]]]]]`)
    p := gparser.New([]string{"a"})
    p.Set("0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0", []int{1,2,3})
    if c, err := p.ToJson(); err == nil {
        fmt.Println(string(c))
        if bytes.Compare(c, e) != 0 {
            t.Error("expect:", string(e))
        }
    } else {
        t.Error(err)
    }
}

func Test_Set6(t *testing.T) {
    e := []byte(`["a",[1,2,3]]`)
    p := gparser.New([]string{"a"})
    p.Set("1", []int{1,2,3})
    if c, err := p.ToJson(); err == nil {
        fmt.Println(string(c))
        if bytes.Compare(c, e) != 0 {
            t.Error("expect:", string(e))
        }
    } else {
        t.Error(err)
    }
}

func Test_Set7(t *testing.T) {
    e := []byte(`{"0":[null,[1,2,3]],"k1":"v1","k2":"v2"}`)
    p := gparser.New(map[string]string{
        "k1" : "v1",
        "k2" : "v2",
    })
    p.Set("0.1", []int{1,2,3})
    if c, err := p.ToJson(); err == nil {
        fmt.Println(string(c))
        if bytes.Compare(c, e) != 0 {
            t.Error("expect:", string(e))
        }
    } else {
        t.Error(err)
    }
}

func Test_Set8(t *testing.T) {
    e := []byte(`{"0":[[[[[[null,[1,2,3]]]]]]],"k1":"v1","k2":"v2"}`)
    p := gparser.New(map[string]string{
        "k1" : "v1",
        "k2" : "v2",
    })
    p.Set("0.0.0.0.0.0.1", []int{1,2,3})
    if c, err := p.ToJson(); err == nil {
        fmt.Println(string(c))
        if bytes.Compare(c, e) != 0 {
            t.Error("expect:", string(e))
        }
    } else {
        t.Error(err)
    }
}

func Test_Set9(t *testing.T) {
    e := []byte(`{"k1":[null,[1,2,3]],"k2":"v2"}`)
    p := gparser.New(map[string]string{
        "k1" : "v1",
        "k2" : "v2",
    })
    p.Set("k1.1", []int{1,2,3})
    if c, err := p.ToJson(); err == nil {
        fmt.Println(string(c))
        if bytes.Compare(c, e) != 0 {
            t.Error("expect:", string(e))
        }
    } else {
        t.Error(err)
    }
}






