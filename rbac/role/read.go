package role

import (
	"github.com/traPtitech/traQ/rbac"
	"github.com/traPtitech/traQ/rbac/permission"
)

// Read 読み取り専用ユーザーロール
const Read = "read"

var readPerms = []rbac.Permission{
	permission.GetChannel,
	permission.GetMessage,
	permission.GetChannelSubscription,
	permission.ConnectNotificationStream,
	permission.GetUser,
	permission.GetMe,
	permission.GetClip,
	permission.GetClipFolder,
	permission.GetChannelStar,
	permission.GetUnread,
	permission.GetChannelMute,
	permission.GetUserTag,
	permission.GetUserGroup,
	permission.GetStamp,
	permission.GetMyStampHistory,
	permission.GetFavoriteStamp,
	permission.DownloadFile,
	permission.GetHeartbeat,
	permission.GetWebhook,
	permission.GetBot,
}