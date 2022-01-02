package main

// [etcd issue 6857](https://github.com/etcd-io/etcd/pull/6857/commits/7afc490c95789c408fbc256d8e790273d331c984)

// raft/node.go

// If the node is stopped, then Status can hang forever because there is no
// event loop to answer. So, just return empty status to avoid deadlocks.

/**
func (n *node) Status() Status {
	c := make(chan Status)
	// 修改前： n.status <- c
	// 修改前： return <-c
	select {
	case n.status <- c:
		return <-c
	case <-n.done:
		return Status{}
	}
}
*/
