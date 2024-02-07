// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: s2/cs_gameevents.proto

package msgs2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ECsgoGameEvents int32

const (
	ECsgoGameEvents_GE_PlayerAnimEventId ECsgoGameEvents = 450
	ECsgoGameEvents_GE_RadioIconEventId  ECsgoGameEvents = 451
	ECsgoGameEvents_GE_FireBulletsId     ECsgoGameEvents = 452
)

// Enum value maps for ECsgoGameEvents.
var (
	ECsgoGameEvents_name = map[int32]string{
		450: "GE_PlayerAnimEventId",
		451: "GE_RadioIconEventId",
		452: "GE_FireBulletsId",
	}
	ECsgoGameEvents_value = map[string]int32{
		"GE_PlayerAnimEventId": 450,
		"GE_RadioIconEventId":  451,
		"GE_FireBulletsId":     452,
	}
)

func (x ECsgoGameEvents) Enum() *ECsgoGameEvents {
	p := new(ECsgoGameEvents)
	*p = x
	return p
}

func (x ECsgoGameEvents) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ECsgoGameEvents) Descriptor() protoreflect.EnumDescriptor {
	return file_s2_cs_gameevents_proto_enumTypes[0].Descriptor()
}

func (ECsgoGameEvents) Type() protoreflect.EnumType {
	return &file_s2_cs_gameevents_proto_enumTypes[0]
}

func (x ECsgoGameEvents) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ECsgoGameEvents) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ECsgoGameEvents(num)
	return nil
}

// Deprecated: Use ECsgoGameEvents.Descriptor instead.
func (ECsgoGameEvents) EnumDescriptor() ([]byte, []int) {
	return file_s2_cs_gameevents_proto_rawDescGZIP(), []int{0}
}

type CMsgTEPlayerAnimEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *uint32 `protobuf:"fixed32,1,opt,name=player,def=16777215" json:"player,omitempty"`
	Event  *uint32 `protobuf:"varint,2,opt,name=event" json:"event,omitempty"`
	Data   *int32  `protobuf:"varint,3,opt,name=data" json:"data,omitempty"`
}

// Default values for CMsgTEPlayerAnimEvent fields.
const (
	Default_CMsgTEPlayerAnimEvent_Player = uint32(16777215)
)

func (x *CMsgTEPlayerAnimEvent) Reset() {
	*x = CMsgTEPlayerAnimEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2_cs_gameevents_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgTEPlayerAnimEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgTEPlayerAnimEvent) ProtoMessage() {}

func (x *CMsgTEPlayerAnimEvent) ProtoReflect() protoreflect.Message {
	mi := &file_s2_cs_gameevents_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgTEPlayerAnimEvent.ProtoReflect.Descriptor instead.
func (*CMsgTEPlayerAnimEvent) Descriptor() ([]byte, []int) {
	return file_s2_cs_gameevents_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgTEPlayerAnimEvent) GetPlayer() uint32 {
	if x != nil && x.Player != nil {
		return *x.Player
	}
	return Default_CMsgTEPlayerAnimEvent_Player
}

func (x *CMsgTEPlayerAnimEvent) GetEvent() uint32 {
	if x != nil && x.Event != nil {
		return *x.Event
	}
	return 0
}

func (x *CMsgTEPlayerAnimEvent) GetData() int32 {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return 0
}

type CMsgTERadioIcon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *uint32 `protobuf:"fixed32,1,opt,name=player,def=16777215" json:"player,omitempty"`
}

// Default values for CMsgTERadioIcon fields.
const (
	Default_CMsgTERadioIcon_Player = uint32(16777215)
)

func (x *CMsgTERadioIcon) Reset() {
	*x = CMsgTERadioIcon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2_cs_gameevents_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgTERadioIcon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgTERadioIcon) ProtoMessage() {}

func (x *CMsgTERadioIcon) ProtoReflect() protoreflect.Message {
	mi := &file_s2_cs_gameevents_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgTERadioIcon.ProtoReflect.Descriptor instead.
func (*CMsgTERadioIcon) Descriptor() ([]byte, []int) {
	return file_s2_cs_gameevents_proto_rawDescGZIP(), []int{1}
}

func (x *CMsgTERadioIcon) GetPlayer() uint32 {
	if x != nil && x.Player != nil {
		return *x.Player
	}
	return Default_CMsgTERadioIcon_Player
}

type CMsgTEFireBullets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin              *CMsgVector `protobuf:"bytes,1,opt,name=origin" json:"origin,omitempty"`
	Angles              *CMsgQAngle `protobuf:"bytes,2,opt,name=angles" json:"angles,omitempty"`
	WeaponId            *uint32     `protobuf:"varint,3,opt,name=weapon_id,json=weaponId,def=16777215" json:"weapon_id,omitempty"`
	Mode                *uint32     `protobuf:"varint,4,opt,name=mode" json:"mode,omitempty"`
	Seed                *uint32     `protobuf:"varint,5,opt,name=seed" json:"seed,omitempty"`
	Player              *uint32     `protobuf:"fixed32,6,opt,name=player,def=16777215" json:"player,omitempty"`
	Inaccuracy          *float32    `protobuf:"fixed32,7,opt,name=inaccuracy" json:"inaccuracy,omitempty"`
	RecoilIndex         *float32    `protobuf:"fixed32,8,opt,name=recoil_index,json=recoilIndex" json:"recoil_index,omitempty"`
	Spread              *float32    `protobuf:"fixed32,9,opt,name=spread" json:"spread,omitempty"`
	SoundType           *int32      `protobuf:"varint,10,opt,name=sound_type,json=soundType" json:"sound_type,omitempty"`
	ItemDefIndex        *uint32     `protobuf:"varint,11,opt,name=item_def_index,json=itemDefIndex" json:"item_def_index,omitempty"`
	SoundDspEffect      *uint32     `protobuf:"fixed32,12,opt,name=sound_dsp_effect,json=soundDspEffect" json:"sound_dsp_effect,omitempty"`
	EntOrigin           *CMsgVector `protobuf:"bytes,13,opt,name=ent_origin,json=entOrigin" json:"ent_origin,omitempty"`
	NumBulletsRemaining *uint32     `protobuf:"varint,14,opt,name=num_bullets_remaining,json=numBulletsRemaining" json:"num_bullets_remaining,omitempty"`
	AttackType          *uint32     `protobuf:"varint,15,opt,name=attack_type,json=attackType" json:"attack_type,omitempty"`
}

// Default values for CMsgTEFireBullets fields.
const (
	Default_CMsgTEFireBullets_WeaponId = uint32(16777215)
	Default_CMsgTEFireBullets_Player   = uint32(16777215)
)

func (x *CMsgTEFireBullets) Reset() {
	*x = CMsgTEFireBullets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2_cs_gameevents_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgTEFireBullets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgTEFireBullets) ProtoMessage() {}

func (x *CMsgTEFireBullets) ProtoReflect() protoreflect.Message {
	mi := &file_s2_cs_gameevents_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgTEFireBullets.ProtoReflect.Descriptor instead.
func (*CMsgTEFireBullets) Descriptor() ([]byte, []int) {
	return file_s2_cs_gameevents_proto_rawDescGZIP(), []int{2}
}

func (x *CMsgTEFireBullets) GetOrigin() *CMsgVector {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *CMsgTEFireBullets) GetAngles() *CMsgQAngle {
	if x != nil {
		return x.Angles
	}
	return nil
}

func (x *CMsgTEFireBullets) GetWeaponId() uint32 {
	if x != nil && x.WeaponId != nil {
		return *x.WeaponId
	}
	return Default_CMsgTEFireBullets_WeaponId
}

func (x *CMsgTEFireBullets) GetMode() uint32 {
	if x != nil && x.Mode != nil {
		return *x.Mode
	}
	return 0
}

func (x *CMsgTEFireBullets) GetSeed() uint32 {
	if x != nil && x.Seed != nil {
		return *x.Seed
	}
	return 0
}

func (x *CMsgTEFireBullets) GetPlayer() uint32 {
	if x != nil && x.Player != nil {
		return *x.Player
	}
	return Default_CMsgTEFireBullets_Player
}

func (x *CMsgTEFireBullets) GetInaccuracy() float32 {
	if x != nil && x.Inaccuracy != nil {
		return *x.Inaccuracy
	}
	return 0
}

func (x *CMsgTEFireBullets) GetRecoilIndex() float32 {
	if x != nil && x.RecoilIndex != nil {
		return *x.RecoilIndex
	}
	return 0
}

func (x *CMsgTEFireBullets) GetSpread() float32 {
	if x != nil && x.Spread != nil {
		return *x.Spread
	}
	return 0
}

func (x *CMsgTEFireBullets) GetSoundType() int32 {
	if x != nil && x.SoundType != nil {
		return *x.SoundType
	}
	return 0
}

func (x *CMsgTEFireBullets) GetItemDefIndex() uint32 {
	if x != nil && x.ItemDefIndex != nil {
		return *x.ItemDefIndex
	}
	return 0
}

func (x *CMsgTEFireBullets) GetSoundDspEffect() uint32 {
	if x != nil && x.SoundDspEffect != nil {
		return *x.SoundDspEffect
	}
	return 0
}

func (x *CMsgTEFireBullets) GetEntOrigin() *CMsgVector {
	if x != nil {
		return x.EntOrigin
	}
	return nil
}

func (x *CMsgTEFireBullets) GetNumBulletsRemaining() uint32 {
	if x != nil && x.NumBulletsRemaining != nil {
		return *x.NumBulletsRemaining
	}
	return 0
}

func (x *CMsgTEFireBullets) GetAttackType() uint32 {
	if x != nil && x.AttackType != nil {
		return *x.AttackType
	}
	return 0
}

var File_s2_cs_gameevents_proto protoreflect.FileDescriptor

var file_s2_cs_gameevents_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x32, 0x2f, 0x63, 0x73, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x75, 0x73, 0x5f, 0x77, 0x61, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x66, 0x6f, 0x63, 0x73, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x73, 0x32, 0x1a, 0x19, 0x73, 0x32, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63,
	0x0a, 0x15, 0x43, 0x4d, 0x73, 0x67, 0x54, 0x45, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x6e,
	0x69, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x3a, 0x08, 0x31, 0x36, 0x37, 0x37, 0x37, 0x32, 0x31,
	0x35, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x0f, 0x43, 0x4d, 0x73, 0x67, 0x54, 0x45, 0x52, 0x61, 0x64,
	0x69, 0x6f, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x3a, 0x08, 0x31, 0x36, 0x37, 0x37, 0x37, 0x32, 0x31, 0x35,
	0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x97, 0x05, 0x0a, 0x11, 0x43, 0x4d, 0x73,
	0x67, 0x54, 0x45, 0x46, 0x69, 0x72, 0x65, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x12, 0x4d,
	0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x6d, 0x61, 0x72, 0x6b,
	0x75, 0x73, 0x5f, 0x77, 0x61, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x66, 0x6f, 0x63, 0x73,
	0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x73, 0x32, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x4d, 0x0a,
	0x06, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x75,
	0x73, 0x5f, 0x77, 0x61, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x66, 0x6f, 0x63, 0x73, 0x5f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x73, 0x32, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x51, 0x41,
	0x6e, 0x67, 0x6c, 0x65, 0x52, 0x06, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x09,
	0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x3a,
	0x08, 0x31, 0x36, 0x37, 0x37, 0x37, 0x32, 0x31, 0x35, 0x52, 0x08, 0x77, 0x65, 0x61, 0x70, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x06, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x07, 0x3a, 0x08, 0x31, 0x36, 0x37,
	0x37, 0x37, 0x32, 0x31, 0x35, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1e, 0x0a,
	0x0a, 0x69, 0x6e, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x69, 0x6e, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x63, 0x6f, 0x69, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x69, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x6e,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x6f,
	0x75, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x64, 0x65, 0x66, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x69, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x66, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x28, 0x0a,
	0x10, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x64, 0x73, 0x70, 0x5f, 0x65, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x73,
	0x70, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x54, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x5f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x75, 0x73, 0x5f,
	0x77, 0x61, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x66, 0x6f, 0x63, 0x73, 0x5f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x73, 0x32, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x32, 0x0a,
	0x15, 0x6e, 0x75, 0x6d, 0x5f, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x5f, 0x72, 0x65, 0x6d,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6e, 0x75,
	0x6d, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x2a, 0x5d, 0x0a, 0x0f, 0x45, 0x43, 0x73, 0x67, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x14, 0x47, 0x45, 0x5f, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x41, 0x6e, 0x69, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x10, 0xc2, 0x03,
	0x12, 0x18, 0x0a, 0x13, 0x47, 0x45, 0x5f, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x49, 0x63, 0x6f, 0x6e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x10, 0xc3, 0x03, 0x12, 0x15, 0x0a, 0x10, 0x47, 0x45,
	0x5f, 0x46, 0x69, 0x72, 0x65, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x49, 0x64, 0x10, 0xc4,
	0x03,
}

var (
	file_s2_cs_gameevents_proto_rawDescOnce sync.Once
	file_s2_cs_gameevents_proto_rawDescData = file_s2_cs_gameevents_proto_rawDesc
)

func file_s2_cs_gameevents_proto_rawDescGZIP() []byte {
	file_s2_cs_gameevents_proto_rawDescOnce.Do(func() {
		file_s2_cs_gameevents_proto_rawDescData = protoimpl.X.CompressGZIP(file_s2_cs_gameevents_proto_rawDescData)
	})
	return file_s2_cs_gameevents_proto_rawDescData
}

var file_s2_cs_gameevents_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_s2_cs_gameevents_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_s2_cs_gameevents_proto_goTypes = []interface{}{
	(ECsgoGameEvents)(0),          // 0: com.github.markus_wa.demoinfocs_golang.s2.ECsgoGameEvents
	(*CMsgTEPlayerAnimEvent)(nil), // 1: com.github.markus_wa.demoinfocs_golang.s2.CMsgTEPlayerAnimEvent
	(*CMsgTERadioIcon)(nil),       // 2: com.github.markus_wa.demoinfocs_golang.s2.CMsgTERadioIcon
	(*CMsgTEFireBullets)(nil),     // 3: com.github.markus_wa.demoinfocs_golang.s2.CMsgTEFireBullets
	(*CMsgVector)(nil),            // 4: com.github.markus_wa.demoinfocs_golang.s2.CMsgVector
	(*CMsgQAngle)(nil),            // 5: com.github.markus_wa.demoinfocs_golang.s2.CMsgQAngle
}
var file_s2_cs_gameevents_proto_depIdxs = []int32{
	4, // 0: com.github.markus_wa.demoinfocs_golang.s2.CMsgTEFireBullets.origin:type_name -> com.github.markus_wa.demoinfocs_golang.s2.CMsgVector
	5, // 1: com.github.markus_wa.demoinfocs_golang.s2.CMsgTEFireBullets.angles:type_name -> com.github.markus_wa.demoinfocs_golang.s2.CMsgQAngle
	4, // 2: com.github.markus_wa.demoinfocs_golang.s2.CMsgTEFireBullets.ent_origin:type_name -> com.github.markus_wa.demoinfocs_golang.s2.CMsgVector
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_s2_cs_gameevents_proto_init() }
func file_s2_cs_gameevents_proto_init() {
	if File_s2_cs_gameevents_proto != nil {
		return
	}
	file_s2_networkbasetypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_s2_cs_gameevents_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgTEPlayerAnimEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_s2_cs_gameevents_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgTERadioIcon); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_s2_cs_gameevents_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgTEFireBullets); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_s2_cs_gameevents_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_s2_cs_gameevents_proto_goTypes,
		DependencyIndexes: file_s2_cs_gameevents_proto_depIdxs,
		EnumInfos:         file_s2_cs_gameevents_proto_enumTypes,
		MessageInfos:      file_s2_cs_gameevents_proto_msgTypes,
	}.Build()
	File_s2_cs_gameevents_proto = out.File
	file_s2_cs_gameevents_proto_rawDesc = nil
	file_s2_cs_gameevents_proto_goTypes = nil
	file_s2_cs_gameevents_proto_depIdxs = nil
}
