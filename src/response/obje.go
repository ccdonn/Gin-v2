package response


type Obje struct {
    Id   string  `json:"id"`
    Name string  `json:"name"`
}

type Objes []Obje


//type FeedResponse struct {
//  status string 
//  time int64 
//  data Objes 
//}