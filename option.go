package jpush

const (
	DistributionJPush         = "jpush"          //值为 jpush 表示推送强制走极光通道下发；
	DistributionOSPush        = "ospush"         //值为 ospush 表示推送强制走厂商通道下发；
	DistributionFirstOSPush   = "first_ospush"   //值为 first_ospush 时表示推送优先走厂商通道下发，无效走极光通道下发；
	DistributionSecondaryPush = "secondary_push" //值为 secondary_push 表示推送优先走极光，极光不在线再走厂商，厂商作为辅助；【建议此种方式】

	DistributionFcm                 = "fcm"                //值为 fcm 表示推送强制走 fcm 通道下发；
	DistributionFcmPns              = "pns"                //值为 pns 表示推送强制走小米/华为/魅族/oppo/vivo 通道下发；
	DistributionFcmJPush            = "jpush"              //值为 jpush 表示推送强制走极光通道下发；
	DistributionFcmSecondaryFcmPush = "secondary_fcm_push" //值为 secondary_fcm_push 表示针对fcm+国内厂商组合类型用户，推送优先走极光，极光不在线再走fcm通道，fcm作为辅助；
	DistributionFcmSecondaryPnsPush = "secondary_pns_push" //值为 secondary_pns_push 表示针对fcm+国内厂商组合类型用户，推送优先走极光，极光不在线再走厂商通道，厂商作为辅助；
)
