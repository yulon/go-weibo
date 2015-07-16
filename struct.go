package weibo

type Status struct {
	CreatedAt string `json: created_at` //微博创建时间
	Id int64 `json: id` //微博ID
	Mid int64 `json: mid` //微博MID
	Idstr string `json: idstr` //字符串型的微博ID
	Text string `json: text` //微博信息内容
	Source string `json: source` //微博来源
	Favorited bool `json: favorited` //是否已收藏，true：是，false：否
	Trucated bool `json: truncated` //是否被截断，true：是，false：否
	InReplyToStatusId string `json: in_reply_to_status_id` //（暂未支持）回复ID
	InReplyToUserId string `json: in_reply_to_user_id` //（暂未支持）回复人UID
	InReplyToScreenName string `json: in_reply_to_screen_name` //（暂未支持）回复人昵称
	ThumbnailPic string `json: thumbnail_pic` //缩略图片地址，没有时不返回此字段
	BmiddlePic string `json: bmiddle_pic` //中等尺寸图片地址，没有时不返回此字段
	OriginalPic string `json: original_pic` //原始图片地址，没有时不返回此字段
	Geo *Geo `json: geo` //地理信息字段
	User *User `json: user` //微博作者的用户信息字
	RetweetedStatus *Status `json: retweeted_status` //被转发的原微博信息字段，当该微博为转发微博时返回
	RepostsCount int `json: reposts_count` //转发数
	CommentsCount int `json: comments_count` //评论数
	AttitudesCount int `json: attitudes_count` //表态数
	Mlevel int `json: mlevel` //暂未支持
	Visible *Visible `json: visible` //微博的可见性及指定可见分组信息。
	PicIds []string `json: pic_ids` //微博配图ID。多图时返回多图ID，用来拼接图片url。用返回字段thumbnail_pic的地址配上该返回字段的图片ID，即可得到多个图片url。
	Ad []int64 `json: ad` //微博流内的推广微博ID
}

type Visible struct {
	Type int `json: type` //0：普通微博，1：私密微博，3：指定分组微博，4：密友微博
	ListId int `json: list_id` //分组的组号
}

type User struct {
	Id int64 `json: id` //用户UID
	Idstr string `json: idstr` //字符串型的用户UID
	ScreenName string `json: screen_name` //用户昵称
	Name string `json: name` //友好显示名称
	Province int `json: province` //用户所在省级ID
	City int `json: city` //用户所在城市ID
	Location string `json: location` //用户所在地
	Description string `json: description` //用户个人描述
	Url string `json: url` //用户博客地址
	ProfileImageUrl string `json: profile_image_url` //用户头像地址（中图），50×50像素
	ProfileUrl string `json: profile_url` //用户的微博统一URL地址
	Domain string `json: domain` //用户的个性化域名
	Weihao string `json: weihao` //用户的微号
	Gender string `json: gender` //性别，m：男、f：女、n：未知
	FollowersCount int `json: followers_count` //粉丝数
	FriendsCount int `json: friends_count` //关注数
	StatusesCount int `json: statuses_count` //微博数
	FavouritesCount int `json: favourites_count` //收藏数
	CreatedAt string `json: created_at` //用户创建（注册）时间
	Following bool `json: following` //暂未支持
	AllowAllActMsg bool `json: allow_all_act_msg` //是否允许所有人给我发私信，true：是，false：否
	GeoEnabled bool `json: geo_enabled` //是否允许标识用户的地理位置，true：是，false：否
	Verified bool `json: verified` //是否是微博认证用户，即加V用户，true：是，false：否
	VerifiedType int `json: verified_type` //暂未支持
	Remark string `json: remark` //用户备注信息，只有在查询用户关系时才返回此字段
	Status *Status `json: status` //用户的最近一条微博信息字段
	AllowAllComment bool `json: allow_all_comment` //是否允许所有人对我的微博进行评论，true：是，false：否
	AvatarLarge string `json: avatar_large` //用户头像地址（大图），180×180像素
	AvatarHd string `json: avatar_hd` //用户头像地址（高清），高清头像原图
	VerifiedReason string `json: verified_reason` //认证原因
	FollowMe bool `json: follow_me` //该用户是否关注当前登录用户，true：是，false：否
	OnlineStatus int `json: online_status` //用户的在线状态，0：不在线、1：在线
	BiFollowersCount int `json: bi_followers_count` //用户的互粉数
	Lang string `json: lang` //用户当前的语言版本，zh-cn：简体中文，zh-tw：繁体中文，en：英语
}

type UserCounts struct {
	Id int64 `json: id` //用户UID
	FollowersCount int `json: followers_count` //粉丝数
	FriendsCount int `json: friends_count` //关注数
	StatusesCount int `json: statuses_count` //微博数
}

type Geo struct{
	Longitude string `json: longitude` //经度坐标
	Latitude string `json: latitude` //维度坐标
	City string `json: city` //所在城市的城市代码
	Province string `json: province` //所在省份的省份代码
	CityName string `json: city_name` //所在城市的城市名称
	ProvinceName string `json: province_name` //所在省份的省份名称
	Address string `json: address` //所在的实际地址，可以为空
	Pinyin string `json: pinyin` //地址的汉语拼音，不是所有情况都会返回该字段
	More string `json: more` //更多信息，不是所有情况都会返回该字段

	//上面是官方给的数据结构但实测返回的都是下面两个属性
	Type string `json: type`
	Coordinates []float64 `json: coordinates`
}

type StatusCount struct {
	Id int64 `json: id` //微博ID
	Reposts int `json: reposts` //转发数
	Comments int `json: comments` //评论数
	Attitudes int `json: attitudes` //表态数
}

type Emotion struct {
	Category string `json: category`
	Common bool `json: common`
	Hot bool `json: hot`
	Icon string `json: icon`
	Phrase string `json: phrase`
	Picid string `json: picid`
	Type string `json: type`
	Url string `json: url`
	Value string `json: value`
}

type Comment struct {
	CreatedAt string `json: created_at` //评论创建时间
	Id int64 `json: id` //评论的ID
	Text string `json: text` //评论的内容
	Source string `json: source` //评论的来源
	User *User `json: user` //评论作者的用户信息字段
	Mid int64 `json: mid` //评论的MID
	Idstr string `json: idstr` //字符串型的评论ID
	Status *Status  `json: status` //评论的微博信息字段
	ReplyComment *Comment `json: reply_comment` //评论来源评论，当本评论属于对另一评论的回复时返回此字段
}

type Remind struct {
	Status int `json: status` //新微博未读数
	Follower int `json: follower` //新粉丝数
	Cmt int `json: cmt` //新评论数
	Dm int `json: dm` //新私信数
	MentionStatus int `json: mention_status` //新提及我的微博数
	MentionCmt int `json: mention_cmt` //新提及我的评论数
	Group int `json: group` //微群消息未读数
	PrivateGroup int `json: private_group` //私有微群消息未读数
	Notice int `json: notice` //新通知未读数
	Invite int `json: invite` //新邀请未读数
	Badge int `json: badge` //新勋章数
	Photo int `json: photo` //相册消息未读数
	Msgbox int `json: msgbox` //{{{3}}}
}
