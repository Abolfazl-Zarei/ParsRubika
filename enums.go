package ParsRubika

// سازنده ابوالفضل زارعی
// آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

type ChatTypeEnum string

const (
	UserChat    ChatTypeEnum = "User"
	BotChat     ChatTypeEnum = "Bot"
	GroupChat   ChatTypeEnum = "Group"
	ChannelChat ChatTypeEnum = "Channel"
)

type FileTypeEnum string

const (
	FileType    FileTypeEnum = "File"
	ImageType   FileTypeEnum = "Image"
	VoiceType   FileTypeEnum = "Voice"
	VideoType   FileTypeEnum = "Video"
	MusicType   FileTypeEnum = "Music"
	GifType     FileTypeEnum = "Gif"
	PhotoType   FileTypeEnum = "Photo"
	StickerType FileTypeEnum = "Sticker"
)

type UpdateTypeEnum string

const (
	UpdatedMessage UpdateTypeEnum = "UpdatedMessage"
	NewMessage     UpdateTypeEnum = "NewMessage"
	RemovedMessage UpdateTypeEnum = "RemovedMessage"
	StartedBot     UpdateTypeEnum = "StartedBot"
	StoppedBot     UpdateTypeEnum = "StoppedBot"
	UpdatedPayment UpdateTypeEnum = "UpdatedPayment"
)

type ButtonTypeEnum string

const (
	ButtonTypeSimple           ButtonTypeEnum = "Simple"
	ButtonTypeSelection        ButtonTypeEnum = "Selection"
	ButtonTypeCalendar         ButtonTypeEnum = "Calendar"
	ButtonTypeNumberPicker     ButtonTypeEnum = "NumberPicker"
	ButtonTypeStringPicker     ButtonTypeEnum = "StringPicker"
	ButtonTypeLocation         ButtonTypeEnum = "Location"
	ButtonTypePayment          ButtonTypeEnum = "Payment"
	ButtonTypeCameraImage      ButtonTypeEnum = "CameraImage"
	ButtonTypeCameraVideo      ButtonTypeEnum = "CameraVideo"
	ButtonTypeGalleryImage     ButtonTypeEnum = "GalleryImage"
	ButtonTypeGalleryVideo     ButtonTypeEnum = "GalleryVideo"
	ButtonTypeFile             ButtonTypeEnum = "File"
	ButtonTypeAudio            ButtonTypeEnum = "Audio"
	ButtonTypeRecordAudio      ButtonTypeEnum = "RecordAudio"
	ButtonTypeMyPhoneNumber    ButtonTypeEnum = "MyPhoneNumber"
	ButtonTypeMyLocation       ButtonTypeEnum = "MyLocation"
	ButtonTypeTextbox          ButtonTypeEnum = "Textbox"
	ButtonTypeLink             ButtonTypeEnum = "Link"
	ButtonTypeAskMyPhoneNumber ButtonTypeEnum = "AskMyPhoneNumber"
	ButtonTypeAskMyLocation    ButtonTypeEnum = "AskMyLocation"
	ButtonTypeBarcode          ButtonTypeEnum = "Barcode"
)

type ChatKeypadTypeEnum string

const (
	NoneKeypad   ChatKeypadTypeEnum = "None"
	NewKeypad    ChatKeypadTypeEnum = "New"
	RemoveKeypad ChatKeypadTypeEnum = "Remove"
)

type UpdateEndpointTypeEnum string

const (
	ReceiveUpdate        UpdateEndpointTypeEnum = "ReceiveUpdate"
	ReceiveInlineMessage UpdateEndpointTypeEnum = "ReceiveInlineMessage"
	ReceiveQuery         UpdateEndpointTypeEnum = "ReceiveQuery"
	GetSelectionItem     UpdateEndpointTypeEnum = "GetSelectionItem"
	SearchSelectionItems UpdateEndpointTypeEnum = "SearchSelectionItems"
)

type MessageSenderEnum string

const (
	UserSender MessageSenderEnum = "User"
	BotSender  MessageSenderEnum = "Bot"
)

type ChatMemberStatusEnum string

const (
	MemberCreator       ChatMemberStatusEnum = "Creator"
	MemberAdministrator ChatMemberStatusEnum = "Administrator"
	MemberMember        ChatMemberStatusEnum = "Member"
	MemberRestricted    ChatMemberStatusEnum = "Restricted"
	MemberLeft          ChatMemberStatusEnum = "Left"
	MemberKicked        ChatMemberStatusEnum = "Kicked"
)

type PollStatusEnum string

const (
	OpenPoll   PollStatusEnum = "Open"
	ClosedPoll PollStatusEnum = "Closed"
)

type ForwardedFromEnum string

const (
	ForwardedFromUser    ForwardedFromEnum = "User"
	ForwardedFromChannel ForwardedFromEnum = "Channel"
	ForwardedFromBot     ForwardedFromEnum = "Bot"
)

type PaymentStatusEnum string

const (
	PaymentPaid    PaymentStatusEnum = "Paid"
	PaymentNotPaid PaymentStatusEnum = "NotPaid"
)

type ButtonSelectionTypeEnum string

const (
	ButtonSelectionTextOnly   ButtonSelectionTypeEnum = "TextOnly"
	ButtonSelectionTextImgThu ButtonSelectionTypeEnum = "TextImgThu"
	ButtonSelectionTextImgBig ButtonSelectionTypeEnum = "TextImgBig"
)

type ButtonSelectionSearchEnum string

const (
	ButtonSelectionSearchNone  ButtonSelectionSearchEnum = "None"
	ButtonSelectionSearchLocal ButtonSelectionSearchEnum = "Local"
	ButtonSelectionSearchApi   ButtonSelectionSearchEnum = "Api"
)

type ButtonSelectionGetEnum string

const (
	ButtonSelectionGetLocal ButtonSelectionGetEnum = "Local"
	ButtonSelectionGetApi   ButtonSelectionGetEnum = "Api"
)

type ButtonCalendarTypeEnum string

const (
	ButtonCalendarDatePersian   ButtonCalendarTypeEnum = "DatePersian"
	ButtonCalendarDateGregorian ButtonCalendarTypeEnum = "DateGregorian"
)

type ButtonTextboxTypeLineEnum string

const (
	ButtonTextboxLineSingle ButtonTextboxTypeLineEnum = "SingleLine"
	ButtonTextboxLineMulti  ButtonTextboxTypeLineEnum = "MultiLine"
)

type ButtonTextboxTypeKeypadEnum string

const (
	ButtonTextboxKeypadString ButtonTextboxTypeKeypadEnum = "String"
	ButtonTextboxKeypadNumber ButtonTextboxTypeKeypadEnum = "Number"
)

type ButtonLocationTypeEnum string

const (
	ButtonLocationPicker ButtonLocationTypeEnum = "Picker"
	ButtonLocationView   ButtonLocationTypeEnum = "View"
)
