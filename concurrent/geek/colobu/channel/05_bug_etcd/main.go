package main

// [etcd issue 6857](https://github.com/etcd-io/etcd/pull/6857/commits/7afc490c95789c408fbc256d8e790273d331c984)
// [etcd issue 5505](https://github.com/etcd-io/etcd/pull/5505/commits/09e8f5782e91a3287c5f38b190f7b74e93bb2049)
// [etcd issue 11256](https://github.com/etcd-io/etcd/issues/11256)

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
