package main

// CrawlersMan manage primary crawler and maintain the list of remote crawlers used
type CrawlersMan struct {
	mongoStore MongoStore
}

func (cm *CrawlersMan) init(config *Config) (err error) {
	cm.mongoStore.init(config.mongoAddress, crawler)
	return
}
func (cm *CrawlersMan) query(request *Request) (response *QueryResponse) { return }
func (cm *CrawlersMan) get(request *Request) (response *QueryResponse)   { return }
func (cm *CrawlersMan) create(request *Request) (response Response)      { return }
func (cm *CrawlersMan) update(request *Request) (response Response)      { return }
func (cm *CrawlersMan) delete(request *Request) (response Response)      { return }