package xdp

import (
	"github.com/cilium/ebpf"
	"golang.org/x/sys/unix"
)

// DefaultSocketOptions is the default SocketOptions used by an xdp.Socket created without specifying options.
var DefaultSocketOptions = SocketOptions{
	NumFrames:              128,
	FrameSize:              2048,
	FillRingNumDescs:       64,
	CompletionRingNumDescs: 64,
	RxRingNumDescs:         64,
	TxRingNumDescs:         64,
}

type umemRing struct {
	Producer *uint32
	Consumer *uint32
	Descs    []uint64
}

type rxTxRing struct {
	Producer *uint32
	Consumer *uint32
	Descs    []Desc
}

// A Socket is an implementation of the AF_XDP Linux socket type for reading packets from a device.
type Socket struct {
	fd                       int
	umem                     []byte
	fillRing                 umemRing
	rxRing                   rxTxRing
	txRing                   rxTxRing
	completionRing           umemRing
	qidconfMap               *ebpf.Map
	xsksMap                  *ebpf.Map
	program                  *ebpf.Program
	ifindex                  int
	numTransmitted           int
	numFilled                int
	freeRXDescs, freeTXDescs []bool
	options                  SocketOptions
	rxDescs                  []Desc
	getTXDescs, getRXDescs   []Desc
}

// SocketOptions are configuration settings used to bind an XDP socket.
type SocketOptions struct {
	NumFrames              int
	FrameSize              int
	FillRingNumDescs       int
	CompletionRingNumDescs int
	RxRingNumDescs         int
	TxRingNumDescs         int
}

// Desc represents an XDP Rx/Tx descriptor.
type Desc unix.XDPDesc

// Stats contains various counters of the XDP socket, such as numbers of
// sent/received frames.
type Stats struct {
	// Filled is the number of items consumed thus far by the Linux kernel
	// from the Fill ring queue.
	Filled uint64

	// Received is the number of items consumed thus far by the user of
	// this package from the Rx ring queue.
	Received uint64

	// Transmitted is the number of items consumed thus far by the Linux
	// kernel from the Tx ring queue.
	Transmitted uint64

	// Completed is the number of items consumed thus far by the user of
	// this package from the Completion ring queue.
	Completed uint64

	// KernelStats contains the in-kernel statistics of the corresponding
	// XDP socket, such as the number of invalid descriptors that were
	// submitted into Fill or Tx ring queues.
	KernelStats unix.XDPStatistics
}

// DefaultSocketFlags are the flags which are passed to bind(2) system call
// when the XDP socket is bound, possible values include unix.XDP_SHARED_UMEM,
// unix.XDP_COPY, unix.XDP_ZEROCOPY.
var DefaultSocketFlags uint16

// DefaultXdpFlags are the flags which are passed when the XDP program is
// attached to the network link, possible values include
// unix.XDP_FLAGS_DRV_MODE, unix.XDP_FLAGS_HW_MODE, unix.XDP_FLAGS_SKB_MODE,
// unix.XDP_FLAGS_UPDATE_IF_NOEXIST.
var DefaultXdpFlags uint32
