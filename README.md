## FuncPProf

Golang 函数级别的监控

### 使用说明

#### 定时监控采样

##### 使用说明

定时采集

```go
//s, 每隔n秒进行一次采集
TickerInterval = 100
//s, 每次采集多久
TickerDelay    = 5   

// 过滤器
// 比如一些标准库的数据可以在这里过滤掉
// true // 忽略
// false  // 不忽略
IgnoreFilter = func(name string) bool {
}

ticker := StartTickerProfiler(false)
// ticker.Stop() // 暂停ticker

// 获取监控的数据
// 参数 传入时间戳(Unix)
// 
profiles := GetProfiles(0)

```

常规采集 

```go
_ = funcpprof.StartProfile()
defer func() {
    _ = funcpprof.StopProfile(true)
    // 获取采集的数据, 并输出到 pprof.log文件
    err := funcpprof.GetProfiles(0).Flush("pprof.log")
    if err != nil {
    }
}()

```



### 广告

[ONES](https://ones.ai) 企业级研发管理解决方案