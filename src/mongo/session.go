package mongo

import (
	"time"

	"github.com/globalsign/mgo"
	logger "github.com/isayme/go-logger"
	"github.com/isayme/go-user/src/conf"
)

// Session mongo session
type Session struct {
	cfg     *conf.Mongo
	session *mgo.Session
}

// NewSession new mongo session
func NewSession(config *conf.Config) (*Session, error) {
	cfg := &config.Mongo

	dailInfo := &mgo.DialInfo{
		Addrs:          cfg.Addrs,
		Timeout:        time.Duration(cfg.Timeout) * time.Second,
		Database:       cfg.Database,
		ReplicaSetName: cfg.ReplicaSetName,
		Source:         cfg.Source,
		PoolLimit:      cfg.PoolLimit,
		Username:       cfg.Username,
		Password:       cfg.Password,
	}

	session, err := mgo.DialWithInfo(dailInfo)
	if err != nil {
		return nil, err
	}
	logger.Debugf("mongodb %v connected", cfg.Addrs)

	return &Session{
		cfg:     cfg,
		session: session,
	}, nil
}

func (s *Session) GetSesion() *mgo.Session {
	return s.session.Copy()
}

func (s *Session) GetCollection(name string) (*mgo.Session, *mgo.Collection) {
	session := s.GetSesion()
	return session, session.DB("").C(name)
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
