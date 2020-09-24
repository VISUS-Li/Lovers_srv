package DB

type HomeCardInfo struct {
	CardType int 	//卡片类型，广告、咨询等
	AdType 	 int 	//广告类型
	InfoType int 	//资讯类型
	ImgUrl	 string //图片路径
	Title    string //卡片标题
	Content	 string //卡片内容
	TypeDesc string //卡片类型说明
	CreateTime uint64 //卡片创建时间
	ShowIndex	int	//卡片显示排名，热度
}
