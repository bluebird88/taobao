package domain

/* 商品视频关联记录 */
type Video struct {
	/* 视频关联记录创建时间（格式：yyyy-MM-dd HH:mm:ss） */
	Created string `json:"created"`

	/* 视频关联记录的id，和商品相对应 */
	Id int64 `json:"id"`

	/* 视频记录关联的商品的数字id(注意：iid近期即将废弃，请用num_iid参数) */
	Iid string `json:"iid"`

	/* 视频关联记录修改时间（格式：yyyy-MM-dd HH:mm:ss） */
	Modified string `json:"modified"`

	/* 视频记录所关联的商品的数字id */
	NumIid int64 `json:"num_iid"`

	/* video的url连接地址。淘秀里视频记录里面存储的url地址 */
	Url string `json:"url"`

	/* video的id，对应于视频在淘秀的存储记录 */
	VideoId int64 `json:"video_id"`
}
