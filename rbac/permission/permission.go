package permission

import "github.com/traPtitech/traQ/rbac"

var List = rbac.Permissions{}

func init() {
	l := []rbac.Permission{
		GetWebhook,
		CreateWebhook,
		EditWebhook,
		DeleteWebhook,
		AccessOthersWebhook,

		GetBot,
		CreateBot,
		EditBot,
		DeleteBot,
		InstallBot,
		UninstallBot,

		CreateChannel,
		GetChannel,
		EditChannel,
		DeleteChannel,
		ChangeParentChannel,
		EditChannelTopic,

		GetMyTokens,
		RevokeMyToken,
		GetClients,
		CreateClient,
		EditMyClient,
		DeleteMyClient,

		GetClip,
		CreateClip,
		DeleteClip,
		GetClipFolder,
		CreateClipFolder,
		PatchClipFolder,
		DeleteClipFolder,

		UploadFile,
		DownloadFile,
		DeleteFile,

		GetHeartbeat,
		PostHeartbeat,

		GetMessage,
		PostMessage,
		EditMessage,
		DeleteMessage,
		ReportMessage,
		GetMessageReports,

		GetChannelMute,
		EditChannelMute,

		GetChannelSubscription,
		EditChannelSubscription,
		ConnectNotificationStream,
		RegisterFCMDevice,

		CreateMessagePin,
		DeleteMessagePin,

		GetMySessions,
		DeleteMySessions,

		GetStamp,
		CreateStamp,
		EditStamp,
		EditStampName,
		EditStampCreatedByOthers,
		DeleteStamp,
		AddMessageStamp,
		RemoveMessageStamp,
		GetMyStampHistory,
		GetFavoriteStamp,
		EditFavoriteStamp,

		GetChannelStar,
		EditChannelStar,

		GetUnread,
		DeleteUnread,

		GetUser,
		RegisterUser,
		GetMe,
		EditMe,
		ChangeMyIcon,
		ChangeMyPassword,
		EditOtherUsers,
		GetUserQRCode,
		GetUserGroup,
		CreateUserGroup,
		CreateSpecialUserGroup,
		EditUserGroup,
		DeleteUserGroup,

		GetUserTag,
		EditUserTag,
	}
	for _, p := range l {
		List.Add(p)
	}
}