package state

import (
	"context"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"

	"git.friday.cafe/fndevs/state/pb"
)

func (s *Server) fmtChannelKey(guild, channel string) fdb.Key {
	return s.Subs.Channels.Pack(tuple.Tuple{guild, channel})
}

func (s *Server) GetChannel(ctx context.Context, req *pb.GetChannelRequest) (*pb.GetChannelResponse, error) {
	ch := new(pb.Channel)

	_, err := s.DB.ReadTransact(func(tx fdb.ReadTransaction) (interface{}, error) {
		raw := tx.Get(s.fmtChannelKey(req.GuildId, req.Id)).MustGet()

		err := ch.Unmarshal(raw)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetChannelResponse{
		Channel: ch,
	}, nil
}

func (s *Server) SetChannel(ctx context.Context, req *pb.SetChannelRequest) (*pb.SetChannelResponse, error) {
	raw, err := req.Channel.Marshal()
	if err != nil {
		return nil, err
	}

	_, err = s.DB.Transact(func(tx fdb.Transaction) (interface{}, error) {
		tx.Set(s.fmtChannelKey(req.Channel.GuildId, req.Channel.Id), raw)
		return nil, nil
	})

	return nil, err
}

func (s *Server) UpdateChannel(ctx context.Context, req *pb.UpdateChannelRequest) (*pb.UpdateChannelResponse, error) {
	ch := new(pb.Channel)

	_, err := s.DB.Transact(func(tx fdb.Transaction) (interface{}, error) {
		raw := tx.Get(s.fmtChannelKey(req.GuildId, req.Id)).MustGet()

		err := ch.Unmarshal(raw)
		if err != nil {
			return nil, err
		}

		if req.Channel.Name != nil {
			ch.Name = req.Channel.Name.Value
		}
		if req.Channel.Topic != nil {
			ch.Topic = req.Channel.Topic.Value
		}
		if req.Channel.Nsfw != nil {
			ch.Nsfw = req.Channel.Nsfw.Value
		}
		if req.Channel.Position != nil {
			ch.Position = req.Channel.Position.Value
		}
		if req.Channel.Bitrate != nil {
			ch.Bitrate = req.Channel.Bitrate.Value
		}
		if req.Channel.Overwrites != nil {
			ch.Overwrites = req.Channel.Overwrites
		}
		if req.Channel.ParentId != nil {
			ch.ParentId = req.Channel.ParentId.Value
		}

		raw, err = ch.Marshal()
		if err != nil {
			return nil, err
		}

		tx.Set(s.fmtChannelKey(ch.GuildId, ch.Id), raw)
		return nil, nil
	})

	return nil, err
}

func (s *Server) DeleteChannel(ctx context.Context, req *pb.DeleteChannelRequest) (*pb.DeleteChannelResponse, error) {
	_, err := s.DB.Transact(func(tx fdb.Transaction) (interface{}, error) {
		tx.Clear(s.fmtChannelKey(req.GuildId, req.Id))
		return nil, nil
	})

	return nil, err
}
