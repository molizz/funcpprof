package funcpprof

import (
	"runtime"
)

/*
// 记录内存分配器的信息
type MemStats struct {
    // 堆空间分配的字节数
    Alloc uint64

    // 从服务开始运行至今分配器为分配的堆空间总和
    TotalAlloc uint64

    // 服务现在使用的内存
    Sys uint64

    // 被runtime监视的指针数
    Lookups uint64

    // 服务malloc的次数
    Mallocs uint64

    // 服务回收的heap objects
    Frees uint64

    //服务分配的堆内存
    HeapAlloc uint64

    //系统分配的堆内存
    HeapSys uint64

    //申请但是为分配的堆内存，（或者回收了的堆内存）
    HeapIdle uint64

    //正在使用的堆内存字节
    HeapInuse uint64

    //返回给OS的堆内存，类似C/C++中的free。
    HeapReleased uint64

    //堆内存块申请的量
    HeapObjects uint64

    //正在使用的栈
    StackInuse uint64

    //系统分配的作为运行栈的内存
    StackSys uint64

    //用于测试用的结构体使用的字节数
    MSpanInuse uint64

    //系统为测试用的结构体分配的字节数
    MSpanSys uint64

    //mcache结构体申请的字节数
    MCacheInuse uint64

    // 操作系统申请的堆空间用于mcache的量
    MCacheSys uint64

    //用于剖析桶散列表的堆空间
    BuckHashSys uint64

    //垃圾回收标记元信息使用的内存
    GCSys uint64

    //golang系统架构占用的额外空间
    OtherSys uint64

    // 下次GC的时间
    NextGC uint64

    // 垃圾回收器最后一次执行时间。
    LastGC uint64

    // 垃圾回收或者其他信息收集导致服务暂停的次数。
    PauseTotalNs uint64

    // 一个循环队列，记录最近垃圾回收系统中断的时间
    PauseNs [256]uint64

    //一个循环队列，记录最近垃圾回收系统中断的时间开始点。
    PauseEnd [256]uint64

    //垃圾回收的内存大小
    NumGC uint32

    //服务调用runtime.GC()强制使用垃圾回收的次数。
    NumForcedGC uint32

    //垃圾回收占用服务CPU工作的时间总和。如果有100个goroutine，垃圾回收的时间为1S,那么就占用了100S
    GCCPUFraction float64

    //是否启用GC
    EnableGC bool

    // DebugGC is currently unused.
    DebugGC bool

    //内存分配器使用情况
    BySize [61]struct {
        // Size is the maximum byte size of an object in this
        // size class.
        Size uint32

        // Mallocs is the cumulative count of heap objects
        // allocated in this size class. The cumulative bytes
        // of allocation is Size*Mallocs. The number of live
        // objects in this size class is Mallocs - Frees.
        Mallocs uint64

        // Frees is the cumulative count of heap objects freed
        // in this size class.
        Frees uint64
    }
}
*/
func MemStats() *runtime.MemStats {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return &mem
}
