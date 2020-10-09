package DB

/******
卡片存数据库信息
说明：ImgUrl，ViewoUrl直接包含在Content中，Content就是个富文本
******/
type HomeCardInfo struct {
	UpLoadUserId	string //上传该卡片用户ID
	CardId			string //该卡片的ID
	CardType int 	//卡片类型，广告、咨询等
	AdType 	 int 	//广告类型
	InfoType int 	//资讯类型
	Title    string //卡片标题
	Content	 string //卡片内容
	HomeHtmlUrl string //卡片内容的HTML文件Url
	TypeDesc string //卡片类型说明
	CreateTime int64 //卡片创建时间
	ShowIndex	int	//卡片显示排名，热度
	IsMainCard	bool //是否为主卡片
	HomeImgUrl	string //卡片在主页显示的图片
	CardMediaType int //卡片媒体类型
	AudioFileUrl string //音频文件地址
	AudioLength string //音频长度
}