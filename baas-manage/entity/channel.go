package entity

type Channel struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	ChainId     int    `json:"chainId" xorm:"not null INT(11)"`
	Orgs        string `json:"orgs" xorm:"not null VARCHAR(255)"`
	ChannelName string `json:"channelName" xorm:"not null VARCHAR(64)"`
	OpenId      string `json:"open_id" xorm:"not null" VARCHAR(50)`
	Created     int64  `json:"created" xorm:"not null BIGINT(20)"`
}
