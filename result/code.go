package result

// Codes:定义状态
type Codes struct {
	Message          map[uint]string
	Success          uint
	Failure          uint
	PasswordError    uint
	SystemError      uint
	ImgKrUploadError uint
	LoadFileError    uint
	FileWirteError   uint
	OpenFileError    uint
	CopyFileError    uint
	CmdStartError    uint
	DialError        uint
	ReadError        uint
	UserExist        uint
	EncryptionFailed uint
	ShouldBindError  uint
	UserNotExist     uint
	AtoiError        uint
	WriteError       uint
	SelectError      uint
	PostFormError    uint
	UploadFileError  uint
	RequsetError     uint
}

// 状态码
var APIcode = &Codes{
	Success:          200,
	Failure:          501,
	PasswordError:    600,
	SystemError:      601,
	ImgKrUploadError: 602,
	LoadFileError:    603,
	FileWirteError:   604,
	OpenFileError:    605,
	CopyFileError:    606,
	CmdStartError:    607,
	DialError:        608,
	ReadError:        609,
	UserExist:        610,
	EncryptionFailed: 612,
	ShouldBindError:  613,
	UserNotExist:     614,
	AtoiError:        615,
	WriteError:       616,
	SelectError:      617,
	PostFormError:    618,
	UploadFileError:  619,
	RequsetError:     620,
}

// 状态信息初始化
func init() {
	APIcode.Message = map[uint]string{
		APIcode.Success:          "成功",
		APIcode.Failure:          "失败",
		APIcode.PasswordError:    "密码错误，请重新输入",
		APIcode.SystemError:      "系统错误",
		APIcode.ImgKrUploadError: "图片上传错误",
		APIcode.LoadFileError:    "上传文件错误",
		APIcode.FileWirteError:   "写文件错误",
		APIcode.OpenFileError:    "打开文件错误",
		APIcode.CopyFileError:    "复制文件失败",
		APIcode.CmdStartError:    "开启cmd失败",
		APIcode.DialError:        "建立websocket连接失败",
		APIcode.ReadError:        "读取失败",
		APIcode.UserExist:        "用户已存在",
		APIcode.EncryptionFailed: "加密失败",
		APIcode.ShouldBindError:  "绑定失败",
		APIcode.UserNotExist:     "用户不存在",
		APIcode.AtoiError:        "字符串转化为整形错误",
		APIcode.WriteError:       "写错误",
		APIcode.SelectError:      "查询失败",
		APIcode.PostFormError:    "获取表单信息失败",
		APIcode.UploadFileError:  "上传文件错误",
		APIcode.RequsetError:     "请求失败",
	}
}

// GetMessage 供外部调用
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
