package test

import "time"

type testModelRepo struct {
	mysql          interface{}
	log            interface{}
	mongo          interface{}
	redis          interface{}
	cache          interface{}
	cacheKeyPrefix string
	cacheDuration  time.Duration
	mq             interface{}
}

// todo NewTestModelRepoRepo .
func NewTestModelRepo() biz.TestModelRepoRepo {
	return &testModelRepo{
		mysql:          nil,
		log:            nil,
		mongo:          nil,
		redis:          nil,
		cache:          nil,
		cacheKeyPrefix: "TestModelRepo:TestModelRepo",
		cacheDuration:  time.Millisecond,
		mq:             nil,
	}
}
