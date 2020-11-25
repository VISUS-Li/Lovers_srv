package Channel


var SportChMng *SportChManage


type SportChManage struct{
	SportChMap map[string]*Channel //管理整个MsgPush的中的所有运动模块Channel连接。键为UserId
}

func NewSportChMng() *SportChManage{
	s := new(SportChManage)
	s.SportChMap = make(map[string]*Channel)
	return s
}

func (chMng *SportChManage)Get(userId string) *Channel{
	ch, ok := chMng.SportChMap[userId]
	if !ok {
		return nil
	}
	return ch
}

func (chMng *SportChManage)Set(userId string, ch *Channel){
	chMng.SportChMap[userId] = ch
}

func (chMng *SportChManage)Del(userId string){
	delete(chMng.SportChMap,userId)
}

