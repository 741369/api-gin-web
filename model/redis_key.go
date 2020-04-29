/**********************************************
** @Des: redis key
** @Author: lg1024
** @Last Modified time: 2019/1/9 下午8:44
***********************************************/

package model

var (
	// 用户是否注册
	RedisDbUserInfo  = 1
	UserInfoRedisKey = "T_USER_INFO_ID_"

	// 游戏首页数据存储
	HomeGameInfoDb                    = 2
	HomeGameSliderDataRedisKey        = "HOME_GAME_SLIDER_DATA"
	HomeGameSearchDataRedisKey        = "HOME_GAME_SEARCH_DATA"
	HomeGameRecommendGameDataRedisKey = "HOME_GAME_RECOMMEND_GAME_DATA"
	HomeGamePerfectGameDataRedisKey   = "HOME_GAME_PERFECT_GAME_DATA"
	HomeGameDailyTaskDataRedisKey     = "HOME_GAME_DAILY_TASK_DATA"
	HomeGameH5GameDataRedisKey        = "HOME_GAME_H5_GAME_DATA"

	NewHomeConfigGameListKey = "NEW_CONFIG_GAME_DATA_"

	NewHomeGameSliderDataRedisKey        = "NEW_HOME_GAME_SLIDER_DATA_"
	NewHomeGameRecommendGameDataRedisKey = "NEW_HOME_GAME_RECOMMEND_GAME_DATA_"
	NewHomeGamePerfectGameDataRedisKey   = "NEW_HOME_GAME_PERFECT_GAME_DATA_"
	NewHomeGameH5GameDataRedisKey        = "NEW_HOME_GAME_H5_GAME_DATA_"
	NewH5SDKRecommendGameListKey         = "NEW_H5_SDK_RECOMMEND_GAME_DATA_"

	// 签到获取
	ActivitySigninUserSignDays  = "ACTIVITY_SIGNIN_USER_SIGN_DAYS_%d"
	ActivitySigninGameRecommend = "ACTIVITY_SIGNIN_GAME_RECOMMEND_%d"

	// 春节活动
	SpringLotteryCountUserFirstSend = "SPRING_LOTTERY_COUNT_USER_%d"

	// 用户信息
	UserInfoNicknameKey = "USER_INFO_NICKNAME_KEY_%d"

	// 游戏信息
	GameInfoDb         = 3
	RewardInfoRedisKey = "REWARD_INFO_ID_%d"

	// 游戏资讯
	HomeGameNewsInfoDb    = 6
	GameNewsInfoKey       = "NEWS_INFO_%s_%s_%d_%d_%d"  // NEWS_INFO_coopID_gameID_type_page_pageSize
	GameNewsDetailInfoKey = "NEWS_DETAIL_INFO_%s_%s_%s" // NEWS_DETAIL_INFO_coopID_gameID_newsID

	// 游戏安装防拦截指引页
	AntiInterceptionConfig = "ANTI_INTERCEPTION_CONFIG_DATA_%s"
)
