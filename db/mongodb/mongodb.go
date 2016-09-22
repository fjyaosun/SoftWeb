package mongodb

import (
	"pca/lib/jlog"
	"pca/lib/pool"
	"time"

	"gopkg.in/mgo.v2"
)

type MgoSrc struct {
	*mgo.Session
}

var (
	session *mgo.Session
	MgoPool pool.Pool
)

func init() {
	configInit()
	var err error
	session, err = mgo.Dial(mongoConf.URL)
	session.Clone()
	if err != nil {
		jlog.LogErrorSendMail(err)
		panic(err)
	}

	MgoPool = pool.ClassicPool(
		mongoConf.MGO_CONN_CAP,
		mongoConf.MGO_CONN_CAP/4,
		func() (pool.Src, error) {
			return &MgoSrc{session.Clone()}, nil
		},
		60*time.Second,
	)
}

func (self *MgoSrc) Usable() bool {
	if self.Session == nil || self.Session.Ping() != nil {
		return false
	}
	return true
}

func (*MgoSrc) Reset() {}

func (self *MgoSrc) Close() {
	if self.Session == nil {
		return
	}
	self.Session.Close()
}
