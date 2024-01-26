package w32

// https://github.com/tpn/winsdk-10/blob/9b69fd26ac0c7d0b83d378dba01080e93349c2ed/Include/10.0.10240.0/um/WinNls.h#L898-L899
const (
	MUI_LANGUAGE_ID               = 0x4   // Use traditional language ID convention
	MUI_LANGUAGE_NAME             = 0x8   // Use ISO language (culture) name convention
	MUI_MACHINE_LANGUAGE_SETTINGS = 0x400 // 取決於系統設定
)
