package xdp

const (
	XDP_FLAGS_DRV_MODE          = 0x4
	XDP_FLAGS_HW_MODE           = 0x8
	XDP_FLAGS_SKB_MODE          = 0x2
	XDP_FLAGS_UPDATE_IF_NOEXIST = 0x1
)

// DefaultSocketFlags are the flags which are passed to bind(2) system call
// when the XDP socket is bound, possible values include unix.XDP_SHARED_UMEM,
// unix.XDP_COPY, unix.XDP_ZEROCOPY.
var DefaultSocketFlags uint16

// DefaultXdpFlags are the flags which are passed when the XDP program is
// attached to the network link, possible values include
// unix.XDP_FLAGS_DRV_MODE, unix.XDP_FLAGS_HW_MODE, unix.XDP_FLAGS_SKB_MODE,
// unix.XDP_FLAGS_UPDATE_IF_NOEXIST.
var DefaultXdpFlags uint32

// SocketOptions are configuration settings used to bind an XDP socket.
type SocketOptions struct {
	NumFrames              int
	FrameSize              int
	FillRingNumDescs       int
	CompletionRingNumDescs int
	RxRingNumDescs         int
	TxRingNumDescs         int
}

// DefaultSocketOptions is the default SocketOptions used by an xdp.Socket created without specifying options.
var DefaultSocketOptions = SocketOptions{
	NumFrames:              128,
	FrameSize:              2048,
	FillRingNumDescs:       64,
	CompletionRingNumDescs: 64,
	RxRingNumDescs:         64,
	TxRingNumDescs:         64,
}
