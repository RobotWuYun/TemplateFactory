package constants

// init
var InitStructFormat = `type %s struct{
	mysql                 interface{}
	log                    interface{}
	mongo					interface{}
	redis					interface{}
	cache					interface{}
    cacheKeyPrefix 			string
	cacheDuration  			time.Duration
	mq             			interface{}
	}
	`

var InitNewFuncFormat = `// todo New%sRepo .
func New%s() biz.%sRepo {
	return &%s{
		mysql :               nil,
	log :                   nil,
	mongo :					nil,
	redis : 					nil,
	cache :					nil,
    cacheKeyPrefix : 			"%s%s",
	cacheDuration : 			time.Millisecond,
	mq :            			nil,
	}
}`
