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

package mac

//#cgo CFLAGS: -x objective-c -Wno-incompatible-pointer-types -Wno-deprecated
//#cgo LDFLAGS: -framework Foundation
//#include "meowAttributedString.h"
//#include "meowMemory.h"
import "C"
import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"runtime"

	"maunium.net/go/maulogger/v2"

	"go.mau.fi/mautrix-imessage/imessage"
)

type AttributeKey string

const (
	AttrBaseWritingDirection AttributeKey = "__kIMBaseWritingDirectionAttributeName"
	AttrFileTransferGUID     AttributeKey = "__kIMFileTransferGUIDAttributeName"
	AttrMessagePartIndex     AttributeKey = "__kIMMessagePartAttributeName"
	AttrURLPreviewData       AttributeKey = "__kIMDataDetectedAttributeName"
	AttrURL                  AttributeKey = "__kIMLinkAttributeName"
)

type Attribute struct {
	Location int                  `json:"location"`
	Length   int                  `json:"length"`
	Values   map[AttributeKey]any `json:"values"`
}

type AttributedString struct {
	Content    string      `json:"content"`
	Attributes []Attribute `json:"attributes"`
}

func (as *AttributedString) SortAttachments(log maulogger.Logger, attachments []*imessage.Attachment) []*imessage.Attachment {
	attachmentMap := make(map[string]*imessage.Attachment, len(attachments))
	for _, attachment := range attachments {
		attachmentMap[attachment.GUID] = attachment
	}
	output := make([]*imessage.Attachment, 0, len(attachments))
	for _, attr := range as.Attributes {
		fileGUID, ok := attr.Values[AttrFileTransferGUID].(string)
		if ok {
			attachment, ok := attachmentMap[fileGUID]
			if ok {
				output = append(output, attachment)
			} else {
				log.Warnfln("Didn't find attachment %s in message", fileGUID)
			}
		}
	}
	return output
}

func meowDecodeAttributedString(data []byte) (*AttributedString, error) {
	runtime.LockOSThread()
	pool := C.meowMakePool()
	var parsed string = C.GoString(C.meowDecodeAttributedString(C.CString(base64.StdEncoding.EncodeToString(data))))
	C.meowReleasePool(pool)
	runtime.UnlockOSThread()
	if parsed[0] != '{' {
		return nil, errors.New(parsed)
	}
	var as AttributedString
	return &as, json.Unmarshal([]byte(parsed), &as)
}
