package state

import (
	"context"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"

	"git.friday.cafe/fndevs/state/pb"
)

func (s *Server) fmtMessageKey(channel, message string) fdb.Key {
	return s.Subs.Messages.Pack(tuple.Tuple{channel, message})
}

func (s *Server) SetMessage(ctx context.Context, req *pb.SetMessageRequest) (*pb.SetMessageResponse, error) {
	raw, err := req.Message.Marshal()
	if err != nil {
		return nil, err
	}

	_, err = s.DB.Transact(func(tx fdb.Transaction) (interface{}, error) {
		tx.Set(s.fmtMessageKey(req.Message.ChannelId, req.Message.Id), raw)
		return nil, nil
	})

	return nil, err
}

func (s *Server) GetMessage(ctx context.Context, req *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	msg := new(pb.Message)

	_, err := s.DB.ReadTransact(func(tx fdb.ReadTransaction) (interface{}, error) {
		raw := tx.Get(s.fmtMessageKey(req.ChannelId, req.Id)).MustGet()

		err := msg.Unmarshal(raw)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetMessageResponse{
		Message: msg,
	}, nil
}
