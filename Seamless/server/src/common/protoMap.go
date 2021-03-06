package common

import "protoMsg"
import "zeus/msgdef"

var ProtoMap = map[uint16]msgdef.IMsg{
	402: new(protoMsg.UserMainDataNotify),
	403: new(protoMsg.Vector3),
	404: new(protoMsg.PlayerRegister),
	405: new(protoMsg.PlayerLogin),
	406: new(protoMsg.PlayerLogout),
	407: new(protoMsg.MailObject),
	408: new(protoMsg.MailInfo),
	409: new(protoMsg.ReqGetMailList),
	410: new(protoMsg.RetMailList),
	411: new(protoMsg.ReqMailInfo),
	412: new(protoMsg.RetMailInfo),
	413: new(protoMsg.DelMail),
	414: new(protoMsg.GetMailObj),
	415: new(protoMsg.AnnuonceInfo),
	416: new(protoMsg.InitAnnuonceInfoRet),
	417: new(protoMsg.FriendInfo),
	418: new(protoMsg.PlatFriendStateReq),
	419: new(protoMsg.PlatFriendState),
	420: new(protoMsg.PlatFriendStateRet),
	421: new(protoMsg.FriendRankInfo),
	422: new(protoMsg.SyncFriendList),
	423: new(protoMsg.FriendApplyInfo),
	424: new(protoMsg.SyncFriendApplyList),
	425: new(protoMsg.DoSendItemReq),
	426: new(protoMsg.GameSvrState),
	427: new(protoMsg.OwnGoodsItem),
	428: new(protoMsg.OwnGoodsInfo),
	429: new(protoMsg.RectInfo),
	430: new(protoMsg.CreateSpaceReq),
	431: new(protoMsg.CreateSpaceRet),
	432: new(protoMsg.ReportCellLoad),
	433: new(protoMsg.CellInfoReq),
	434: new(protoMsg.CellInfoRet),
	435: new(protoMsg.CellBorderChangeNotify),
	436: new(protoMsg.CellInfoNotify),
	437: new(protoMsg.CreateCellNotify),
	438: new(protoMsg.DeleteCellNotify),
	439: new(protoMsg.SkillEffect),
	440: new(protoMsg.AttackReq),
	441: new(protoMsg.MoveReq),
	442: new(protoMsg.MoveUpdate),
	443: new(protoMsg.DetectCell),
	444: new(protoMsg.EnterCellOk),
	445: new(protoMsg.EffectNotify),
}
