// mautrix-imessage - A Matrix-iMessage puppeting bridge.
// Copyright (C) 2022 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package ios

import (
	"maunium.net/go/mautrix/id"

	"go.mau.fi/mautrix-imessage/imessage"
	"go.mau.fi/mautrix-imessage/ipc"
)

const (
	ReqSendMessage         ipc.Command = "send_message"
	ReqSendMedia           ipc.Command = "send_media"
	ReqSendTapback         ipc.Command = "send_tapback"
	ReqSendReadReceipt     ipc.Command = "send_read_receipt"
	ReqSetTyping           ipc.Command = "set_typing"
	ReqGetChats            ipc.Command = "get_chats"
	ReqGetChat             ipc.Command = "get_chat"
	ReqGetChatAvatar       ipc.Command = "get_chat_avatar"
	ReqGetContact          ipc.Command = "get_contact"
	ReqGetContactList      ipc.Command = "get_contact_list"
	ReqGetMessagesAfter    ipc.Command = "get_messages_after"
	ReqGetRecentMessages   ipc.Command = "get_recent_messages"
	ReqGetMessage          ipc.Command = "get_message"
	ReqPreStartupSync      ipc.Command = "pre_startup_sync"
	ReqPostStartupSync     ipc.Command = "post_startup_sync"
	ReqResolveIdentifier   ipc.Command = "resolve_identifier"
	ReqPrepareDM           ipc.Command = "prepare_dm"
	ReqMessageBridgeResult ipc.Command = "message_bridge_result"
	ReqChatBridgeResult    ipc.Command = "chat_bridge_result"
	ReqBackfillResult      ipc.Command = "backfill_result"
	ReqUpcomingMessage     ipc.Command = "upcoming_message"
)

type SendMessageRequest struct {
	ChatGUID    string                   `json:"chat_guid"`
	Text        string                   `json:"text"`
	ReplyTo     string                   `json:"reply_to"`
	ReplyToPart int                      `json:"reply_to_part"`
	RichLink    *imessage.RichLink       `json:"rich_link,omitempty"`
	Metadata    imessage.MessageMetadata `json:"metadata,omitempty"`
}

type SendMediaRequest struct {
	ChatGUID string `json:"chat_guid"`
	Text     string `json:"text"`
	imessage.Attachment
	ReplyTo        string                   `json:"reply_to"`
	ReplyToPart    int                      `json:"reply_to_part"`
	IsAudioMessage bool                     `json:"is_audio_message"`
	Metadata       imessage.MessageMetadata `json:"metadata,omitempty"`
}

type SendTapbackRequest struct {
	ChatGUID   string               `json:"chat_guid"`
	TargetGUID string               `json:"target_guid"`
	TargetPart int                  `json:"target_part"`
	Type       imessage.TapbackType `json:"type"`
}

type SendReadReceiptRequest struct {
	ChatGUID string `json:"chat_guid"`
	ReadUpTo string `json:"read_up_to"`
}

type SetTypingRequest struct {
	ChatGUID string `json:"chat_guid"`
	Typing   bool   `json:"typing"`
}

type GetChatRequest struct {
	ChatGUID string `json:"chat_guid"`
	ThreadID string `json:"thread_id"`
}

type GetChatsRequest struct {
	MinTimestamp float64 `json:"min_timestamp"`
}

type GetContactRequest struct {
	UserGUID string `json:"user_guid"`
}

type GetContactListResponse struct {
	Contacts []*imessage.Contact `json:"contacts"`
}

type GetRecentMessagesRequest struct {
	ChatGUID   string `json:"chat_guid"`
	Limit      int    `json:"limit"`
	BackfillID string `json:"backfill_id"`
}

type GetMessageRequest struct {
	GUID string `json:"guid"`
}

type GetMessagesAfterRequest struct {
	ChatGUID   string  `json:"chat_guid"`
	Timestamp  float64 `json:"timestamp"`
	BackfillID string  `json:"backfill_id"`
}

type PingServerResponse struct {
	Start  float64 `json:"start_ts"`
	Server float64 `json:"server_ts"`
	End    float64 `json:"end_ts"`
}

type ResolveIdentifierRequest struct {
	Identifier string `json:"identifier"`
}

type ResolveIdentifierResponse struct {
	GUID string `json:"guid"`
}

type PrepareDMRequest struct {
	GUID string `json:"guid"`
}

type MessageBridgeResult struct {
	ChatGUID string     `json:"chat_guid"`
	GUID     string     `json:"message_guid"`
	EventID  id.EventID `json:"event_id,omitempty"`
	Success  bool       `json:"success"`
}

type ChatBridgeResult struct {
	ChatGUID string    `json:"chat_guid"`
	MXID     id.RoomID `json:"mxid"`
}

type BackfillResult struct {
	ChatGUID   string `json:"chat_guid"`
	BackfillID string `json:"backfill_id"`
	Success    bool   `json:"success"`

	MessageIDs map[string][]id.EventID `json:"message_ids"`
}

type UpcomingMessage struct {
	EventID id.EventID `json:"event_id"`
}
