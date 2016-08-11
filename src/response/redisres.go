package response

import (

)

type RedisRes struct {
  RvwId string `json:"rvwId"`
  PosterId string `json:"posterId"`
  Content string `json:"content"`
  CreateTime int64 `json:"createTime"`
  DisplayTime int64 `json:"displayTime"`
  ImageUrl1 string `json:"imageUrl1"`
  ImageUrl2 string `json:"imageUrl2"`
  ImageUrl3 string `json:"imageUrl3"`
  Privacy int `json:"privacy"`
  Score float64 `json:"score"`
  Alert int `json:"alert"`
  StatusId int `json:"statusId"`
  
  Spot Spot `json:"spot"`
}

type Spot struct {
  Name string `json:"name"`
  Address string `json:"address"`
  Provider string `json:"provider"`
  OutsideId string `json:"outsideId"`
  Lat string `json:"lat"`
  Lng string `json:"lng"`
  Id string `json:"id"`
}