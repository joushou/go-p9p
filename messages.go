package p9pnew

import "fmt"

// Message represents the target of an fcall.
type Message interface{}

// newMessage returns a new instance of the message based on the Fcall type.
func newMessage(typ FcallType) (Message, error) {
	switch typ {
	case Tversion:
		return MessageTversion{}, nil
	case Rversion:
		return MessageRversion{}, nil
	case Tauth:
		return MessageTauth{}, nil
	case Rauth:
		return MessageRauth{}, nil
	case Tattach:
		return MessageTattach{}, nil
	case Rattach:
		return MessageRattach{}, nil
	case Rerror:
		return MessageRerror{}, nil
	case Tflush:
		return MessageTflush{}, nil
	case Rflush:
		return MessageRflush{}, nil // No message body for this response.
	case Twalk:
		return MessageTwalk{}, nil
	case Rwalk:
		return MessageRwalk{}, nil
	case Topen:
		return MessageTopen{}, nil
	case Ropen:
		return MessageRopen{}, nil
	case Tcreate:
		return MessageTcreate{}, nil
	case Rcreate:
		return MessageRcreate{}, nil
	case Tread:
		return MessageTread{}, nil
	case Rread:
		return MessageRread{}, nil
	case Twrite:
		return MessageTwrite{}, nil
	case Rwrite:
		return MessageRwrite{}, nil
	case Tclunk:
		return MessageTclunk{}, nil
	case Rclunk:
		return MessageRclunk{}, nil // no response body
	case Tremove:
		return MessageTremove{}, nil
	case Rremove:
		return MessageRremove{}, nil
	case Tstat:
		return MessageTstat{}, nil
	case Rstat:
		return MessageRstat{}, nil
	case Twstat:
		return MessageTwstat{}, nil
	case Rwstat:
		return MessageRwstat{}, nil
	}

	return nil, fmt.Errorf("unknown message type")
}

func messageType(m Message) FcallType {
	switch v := m.(type) {
	case MessageTversion:
		return Tversion
	case MessageRversion:
		return Rversion
	case MessageTauth:
		return Tauth
	case MessageRauth:
		return Rauth
	case MessageTflush:
		return Tflush
	case MessageRflush:
		return Rflush
	case MessageRerror:
		return Rerror
	case MessageTattach:
		return Tattach
	case MessageRattach:
		return Rattach
	case MessageTwalk:
		return Twalk
	case MessageRwalk:
		return Rwalk
	case MessageTopen:
		return Topen
	case MessageRopen:
		return Ropen
	case MessageTcreate:
		return Tcreate
	case MessageRcreate:
		return Rcreate
	case MessageTread:
		return Tread
	case MessageRread:
		return Rread
	case MessageTwrite:
		return Twrite
	case MessageRwrite:
		return Rwrite
	case MessageTclunk:
		return Tclunk
	case MessageRclunk:
		return Rclunk
	case MessageTremove:
		return Tremove
	case MessageRremove:
		return Rremove
	case MessageTstat:
		return Tstat
	case MessageRstat:
		return Rstat
	case MessageTwstat:
		return Twstat
	case MessageRwstat:
		return Rwstat
	case error:
		return Rerror
	default:
		// NOTE(stevvooe): This is considered a programming error.
		panic(fmt.Sprintf("unsupported message type: %T", v))
	}
}

// MessageVersion encodes the message body for Tversion and Rversion RPC
// calls. The body is identical in both directions.
type MessageTversion struct {
	MSize   uint32
	Version string
}

type MessageRversion struct {
	MSize   uint32
	Version string
}

type MessageTauth struct {
	Afid  Fid
	Uname string
	Aname string
}

type MessageRauth struct {
	Qid Qid
}

type MessageRerror struct {
	Ename string
}

func (e MessageRerror) Error() string {
	return fmt.Sprintf("9p: %v", e.Ename)
}

type MessageTflush struct {
	Oldtag Tag
}

type MessageRflush struct{}

type MessageTattach struct {
	Fid   Fid
	Afid  Fid
	Uname string
	Aname string
}

type MessageRattach struct {
	Qid Qid
}

type MessageTwalk struct {
	Fid    Fid
	Newfid Fid
	Wnames []string
}

type MessageRwalk struct {
	Qids []Qid
}

type MessageTopen struct {
	Fid  Fid
	Mode uint8
}

type MessageRopen struct {
	Qid    Qid
	IOUnit uint32
}

type MessageTcreate struct {
	Fid  Fid
	Name string
	Perm uint32
	Mode uint8
}

type MessageRcreate struct {
	Qid    Qid
	IOUnit uint32
}

type MessageTread struct {
	Fid    Fid
	Offset uint64
	Count  uint32
}

type MessageRread struct {
	Data []byte
}

type MessageTwrite struct {
	Fid    Fid
	Offset uint64
	Data   []byte
}

type MessageRwrite struct {
	Count uint32
}

type MessageTclunk struct {
	Fid Fid
}

type MessageRclunk struct{}

type MessageTremove struct {
	Fid Fid
}

type MessageRremove struct{}

type MessageTstat struct {
	Fid Fid
}

type MessageRstat struct {
	Stat Dir
}

type MessageTwstat struct {
	Fid  Fid
	Stat Dir
}

type MessageRwstat struct{}
