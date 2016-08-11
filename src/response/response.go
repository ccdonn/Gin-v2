package response


type Response struct {
  Status string `json:"status"`
  Timestamp int64 `json:"timstamp"`
}

type FeedResponse struct {
  Status string `json:"status"`
  Timestamp int64 `json:"timstamp"`
  Result Objes `json:"result"`
}

type FeedRedisResponse struct {
  Response
  Result RedisRes `json:"result"`
}