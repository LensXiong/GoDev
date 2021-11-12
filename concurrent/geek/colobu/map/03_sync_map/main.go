package main

// 问题：如何实现线程安全的 map 类型?

// 三种方案：
// 方案①：加读写锁，扩展 map，支持并发读写。
// 方案②：分片加锁（更高效的并发`map`）。
// 方案③：应对特殊场景的 sync.Map。
