package models

type User struct {
	ScreenName  string `json:"screen_name"`
	Location    string `json:"location"`
	ID          int64  `json:"id"`
	Description string `json:"description"`
}

type UserTweets struct {
	ID        int64  `json:"id"`
	IDStr     string `json:"id_str"`
	Retweeted bool   `json:"retweeted"`
	Text      string `json:"text"`
	FullText  string `json:"full_text"`
	User      *User
}
