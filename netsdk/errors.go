package netsdk

import (
	"C"
	"errors"
)

var (
	ErrNET_DVR_Init = errors.New("dvr init error")
)

const (
	NetDvrNoerror                = 0   //没有错误
	NetDvrPasswordError          = 1   //用户名密码错误
	NetDvrNoenoughpri            = 2   //权限不足
	NetDvrNoinit                 = 3   //没有初始化
	NetDvrChannelError           = 4   //通道号错误
	NetDvrOverMaxlink            = 5   //连接到DVR的客户端个数超过最大
	NetDvrVersionnomatch         = 6   //版本不匹配
	NetDvrNetworkFailConnect     = 7   //连接服务器失败
	NetDvrNetworkSendError       = 8   //向服务器发送失败
	NetDvrNetworkRecvError       = 9   //从服务器接收数据失败
	NetDvrNetworkRecvTimeout     = 10  //从服务器接收数据超时
	NetDvrNetworkErrordata       = 11  //传送的数据有误
	NetDvrOrderError             = 12  //调用次序错误
	NetDvrOpernopermit           = 13  //无此权限
	NetDvrCommandtimeout         = 14  //DVR命令执行超时
	NetDvrErrorserialport        = 15  //串口号错误
	NetDvrErroralarmport         = 16  //报警端口错误
	NetDvrParameterError         = 17  //参数错误
	NetDvrChanException          = 18  //服务器通道处于错误状态
	NetDvrNodisk                 = 19  //没有硬盘
	NetDvrErrordisknum           = 20  //硬盘号错误
	NetDvrDiskFull               = 21  //服务器硬盘满
	NetDvrDiskError              = 22  //服务器硬盘出错
	NetDvrNosupport              = 23  //服务器不支持
	NetDvrBusy                   = 24  //服务器忙
	NetDvrModifyFail             = 25  //服务器修改不成功
	NetDvrPasswordFormatError    = 26  //密码输入格式不正确
	NetDvrDiskFormating          = 27  //硬盘正在格式化，不能启动操作
	NetDvrDvrnoresource          = 28  //DVR资源不足
	NetDvrDvropratefailed        = 29  //DVR操作失败
	NetDvrOpenhostsoundFail      = 30  //打开PC声音失败
	NetDvrDvrvoiceopened         = 31  //服务器语音对讲被占用
	NetDvrTimeinputerror         = 32  //时间输入不正确
	NetDvrNospecfile             = 33  //回放时服务器没有指定的文件
	NetDvrCreatefileError        = 34  //创建文件出错
	NetDvrFileopenfail           = 35  //打开文件出错
	NetDvrOpernotfinish          = 36  //上次的操作还没有完成
	NetDvrGetplaytimefail        = 37  //获取当前播放的时间出错
	NetDvrPlayfail               = 38  //播放出错
	NetDvrFileformatError        = 39  //文件格式不正确
	NetDvrDirError               = 40  //路径错误
	NetDvrAllocResourceError     = 41  //资源分配错误
	NetDvrAudioModeError         = 42  //声卡模式错误
	NetDvrNoenoughBuf            = 43  //缓冲区太小
	NetDvrCreatesocketError      = 44  //创建SOCKET出错
	NetDvrSetsocketError         = 45  //设置SOCKET出错
	NetDvrMaxNum                 = 46  //个数达到最大
	NetDvrUsernotexist           = 47  //用户不存在
	NetDvrWriteflasherror        = 48  //写FLASH出错
	NetDvrUpgradefail            = 49  //DVR升级失败
	NetDvrCardhaveinit           = 50  //解码卡已经初始化过
	NetDvrPlayerfailed           = 51  //调用播放库中某个函数失败
	NetDvrMaxUsernum             = 52  //设备端用户数达到最大
	NetDvrGetlocalipandmacfail   = 53  //获得客户端的IP地址或物理地址失败
	NetDvrNoencodeing            = 54  //该通道没有编码
	NetDvrIpmismatch             = 55  //IP地址不匹配
	NetDvrMacmismatch            = 56  //MAC地址不匹配
	NetDvrUpgradelangmismatch    = 57  //升级文件语言不匹配
	NetDvrMaxPlayerport          = 58  //播放器路数达到最大
	NetDvrNospacebackup          = 59  //备份设备中没有足够空间进行备份
	NetDvrNodevicebackup         = 60  //没有找到指定的备份设备
	NetDvrPictureBitsError       = 61  //图像素位数不符，限24色
	NetDvrPictureDimensionError  = 62  //图片高*宽超限， 限128*256
	NetDvrPictureSizError        = 63  //图片大小超限，限100K
	NetDvrLoadplayersdkfailed    = 64  //载入当前目录下Player Sdk出错
	NetDvrLoadplayersdkprocError = 65  //找不到Player Sdk中某个函数入口
	NetDvrLoaddssdkfailed        = 66  //载入当前目录下DSsdk出错
	NetDvrLoaddssdkprocError     = 67  //找不到DsSdk中某个函数入口
	NetDvrDssdkError             = 68  //调用硬解码库DsSdk中某个函数失败
	NetDvrVoicemonopolize        = 69  //声卡被独占
	NetDvrJoinmulticastfailed    = 70  //加入多播组失败
	NetDvrCreatedirError         = 71  //建立日志文件目录失败
	NetDvrBindsocketError        = 72  //绑定套接字失败
	NetDvrSocketcloseError       = 73  //socket连接中断，此错误通常是由于连接中断或目的地不可达
	NetDvrUseridIsusing          = 74  //注销时用户ID正在进行某操作
	NetDvrSocketlistenError      = 75  //监听失败
	NetDvrProgramException       = 76  //程序异常
	NetDvrWritefileFailed        = 77  //写文件失败
	NetDvrFormatReadonly         = 78  //禁止格式化只读硬盘
	NetDvrWithsameusername       = 79  //用户配置结构中存在相同的用户名
	NetDvrDevicetypeError        = 80  /*导入参数时设备型号不匹配*/
	NetDvrLanguageError          = 81  /*导入参数时语言不匹配*/
	NetDvrParaversionError       = 82  /*导入参数时软件版本不匹配*/
	NetDvrIpchanNotalive         = 83  /*预览时外接IP通道不在线*/
	NetDvrRtspSdkError           = 84  /*加载高清IPC通讯库StreamTransClient.dll失败*/
	NetDvrConvertSdkError        = 85  /*加载转码库失败*/
	NetDvrIpcCountOverflow       = 86  /*超出最大的ip接入通道数*/
	NetDvrMaxAddNum              = 87  /*添加标签(一个文件片段64)等个数达到最大*/
	NetDvrParammodeError         = 88  //图像增强仪，参数模式错误（用于硬件设置时，客户端进行软件设置时错误值）
	NetDvrCodespitterOffline     = 89  //视频综合平台，码分器不在线
	NetDvrBackupCopying          = 90  //设备正在备份
	NetDvrChanNotsupport         = 91  // 通道不支持该操作
	NetDvrCallineinvalid         = 92  // 高度线位置太集中或长度线不够倾斜
	NetDvrCalcancelconflict      = 93  // 取消标定冲突，如果设置了规则及全局的实际大小尺寸过滤
	NetDvrCalpointoutrange       = 94  // 标定点超出范围
	NetDvrFilterrectinvalid      = 95  // 尺寸过滤器不符合要求
	NetDvrDdnsDevoffline         = 96  //设备没有注册到ddns上
	NetDvrDdnsInterError         = 97  //DDNS 服务器内部错误
	NetDvrFunctionNotSupportOs   = 98  //此功能不支持该操作系统
	NetDvrDecChanRebind          = 99  //解码通道绑定显示输出次数受限
	NetDvrIntercomSdkError       = 100 //加载当前目录下的语音对讲库失败
	NetDvrNoCurrentUpdatefile    = 101 //没有正确的升级包
	NetDvrUserNotSuccLogin       = 102 //用户还没登陆成功
	NetDvrUseLogSwitchFile       = 103 //正在使用日志开关文件
	NetDvrPoolPortExhaust        = 104 //端口池中用于绑定的端口已耗尽
	NetDvrPacketTypeNotSupport   = 105 //码流封装格式错误
	NetDvrIpparaIpidError        = 106 //IP接入配置时IPID有误

	NetDvrLoadHcpreviewSdkError       = 107 //预览组件加载失败
	NetDvrLoadHcvoicetalkSdkError     = 108 //语音组件加载失败
	NetDvrLoadHcalarmSdkError         = 109 //报警组件加载失败
	NetDvrLoadHcplaybackSdkError      = 110 //回放组件加载失败
	NetDvrLoadHcdisplaySdkError       = 111 //显示组件加载失败
	NetDvrLoadHcindustrySdkError      = 112 //行业应用组件加载失败
	NetDvrLoadHcgeneralcfgmgrSdkError = 113 //通用配置管理组件加载失败
	NetDvrLoadHccoredevcfgSdkError    = 114 //设备配置核心组件加载失败
	NetDvrLoadHcnetutilsSdkError      = 115 //HCNetUtils加载失败

	NetDvrCoreVerMismatch                = 121 //单独加载组件时，组件与core版本不匹配
	NetDvrCoreVerMismatchHcpreview       = 122 //预览组件与core版本不匹配
	NetDvrCoreVerMismatchHcvoicetalk     = 123 //语音组件与core版本不匹配
	NetDvrCoreVerMismatchHcalarm         = 124 //报警组件与core版本不匹配
	NetDvrCoreVerMismatchHcplayback      = 125 //回放组件与core版本不匹配
	NetDvrCoreVerMismatchHcdisplay       = 126 //显示组件与core版本不匹配
	NetDvrCoreVerMismatchHcindustry      = 127 //行业应用组件与core版本不匹配
	NetDvrCoreVerMismatchHcgeneralcfgmgr = 128 //通用配置管理组件与core版本不匹配

	NetDvrComVerMismatchHcpreview       = 136 //预览组件与HCNetSDK版本不匹配
	NetDvrComVerMismatchHcvoicetalk     = 137 //语音组件与HCNetSDK版本不匹配
	NetDvrComVerMismatchHcalarm         = 138 //报警组件与HCNetSDK版本不匹配
	NetDvrComVerMismatchHcplayback      = 139 //回放组件与HCNetSDK版本不匹配
	NetDvrComVerMismatchHcdisplay       = 140 //显示组件与HCNetSDK版本不匹配
	NetDvrComVerMismatchHcindustry      = 141 //行业应用组件与HCNetSDK版本不匹配
	NetDvrComVerMismatchHcgeneralcfgmgr = 142 //通用配置管理组件与HCNetSDK版本不匹配

	NetErrConfigFileImportFailed = 145 //配置文件导入失败
	NetErrConfigFileExportFailed = 146 //配置文件导出失败
	NetDvrCertificateFileError   = 147 //证书错误
	NetDvrLoadSslLibError        = 148 //加载SSL库失败（可能是版本不匹配，也可能是不存在）
	NetDvrSslVersionNotMatch     = 149 //SSL库版本不匹配

	NetDvrAliasDuplicate       = 150 //别名重复  //2011-08-31 通过别名或者序列号来访问设备的新版本ddns的配置
	NetDvrInvalidCommunication = 151 //无效通信
	NetDvrUsernameNotExist     = 152 //用户名不存在（用户名不存在，IPC5.1.7中发布出去了，所以删不掉。后续的产品这个错误码用不上）
	NetDvrUserLocked           = 153 //用户被锁定
	NetDvrInvalidUserid        = 154 //无效用户ID
	NetDvrLowLoginVersion      = 155 //登录版本低
	NetDvrLoadLibeay32DllError = 156 //加载libeay32.dll库失败
	NetDvrLoadSsleay32DllError = 157 //加载ssleay32.dll库失败
	NetErrLoadLibiconv         = 158 //加载libiconv库失败
	NetErrSslConnectFailed     = 159 //SSL连接失败
	NetErrMcastAddressError    = 160 //获取多播地址错误
	NetErrLoadZlib             = 161 //加载zlib.dll库失败
	NetErrOpensslNoInit        = 162 //Openssl库未初始化

	NetDvrServerNotExist                      = 164 //对应的服务器找不到,查找时输入的国家编号或者服务器类型错误
	NetDvrTestServerFailConnect               = 165 //连接测试服务器失败
	NetDvrNasServerInvalidDir                 = 166 //NAS服务器挂载目录失败，目录无效
	NetDvrNasServerNoenoughPri                = 167 //NAS服务器挂载目录失败，没有权限
	NetDvrEmailServerNotConfigDns             = 168 //服务器使用域名，但是没有配置DNS，可能造成域名无效。
	NetDvrEmailServerNotConfigGateway         = 169 //没有配置网关，可能造成发送邮件失败。
	NetDvrTestServerPasswordError             = 170 //用户名密码不正确，测试服务器的用户名或密码错误
	NetDvrEmailServerConnectExceptionWithSmtp = 171 //设备和smtp服务器交互异常
	NetDvrFtpServerFailCreateDir              = 172 //FTP服务器创建目录失败
	NetDvrFtpServerNoWritePir                 = 173 //FTP服务器没有写入权限
	NetDvrIpConflict                          = 174 //IP冲突
	NetDvrInsufficientStoragepoolSpace        = 175 //存储池空间已满
	NetDvrStoragepoolInvalid                  = 176 //云服务器存储池无效,没有配置存储池或者存储池ID错误
	NetDvrEffectivenessReboot                 = 177 //生效需要重启
	NetErrAnrArmingExist                      = 178 //断网续传布防连接已经存在(该错误码是在HIK私有布防连接建立的情况下，重复布防的断网续传功能时，返回。)
	NetErrUploadlinkExist                     = 179 //断网续传上传连接已经存在(EHOME协议和HIK SDK协议是不能同时支持断网续传的，当一个协议存在的时候，另外一个连接建立话，报错这个错误码。)
	NetErrIncorrectFileFormat                 = 180 //导入文件格式不正确
	NetErrIncorrectFileContent                = 181 //导入文件内容不正确
	NetErrMaxHrudpLink                        = 182 //HRUDP 连接数 超过设备限制
	NetSdkErrAccesskeySecretkey               = 183 // 接入秘钥或加密秘钥不正确
	NetSdkErrCreatePortMultiplex              = 184 //创建端口复用失败
	NetDvrNonblockingCaptureNotsupport        = 185 //不支持无阻塞抓图
	NetSdkErrFunctionInvalid                  = 186 //已开启异步，该功能无效
	NetSdkErrMaxPortMultiplex                 = 187 //已达到端口复用最大数目
)

func Err(code int) error {

	switch code {
	case NetDvrNoerror:
		return nil
	case NetDvrPasswordError:
		return errors.New("用户名密码错误")
	case NetDvrNoenoughpri:
		return errors.New("权限不足")
	case NetDvrNoinit:
		return errors.New("没有初始化")
	case NetDvrChannelError:
		return errors.New("通道号错误")
	case NetDvrOverMaxlink:
		return errors.New("连接到DVR的客户端个数超过最大")
	case NetDvrVersionnomatch:
		return errors.New("版本不匹配")
	case NetDvrNetworkFailConnect:
		return errors.New("连接服务器失败")
	case NetDvrNetworkSendError:
		return errors.New("向服务器发送失败")
	case NetDvrNetworkRecvError:
		return errors.New("从服务器接收数据失败")
	case NetDvrNetworkRecvTimeout:
		return errors.New("从服务器接收数据超时")
	case NetDvrNetworkErrordata:
		return errors.New("传送的数据有误")
	case NetDvrOrderError:
		return errors.New("调用次序错误")
	case NetDvrOpernopermit:
		return errors.New("无此权限")
	case NetDvrCommandtimeout:
		return errors.New("DVR命令执行超时")
	case NetDvrErrorserialport:
		return errors.New("串口号错误")
	case NetDvrErroralarmport:
		return errors.New("报警端口错误")
	case NetDvrParameterError:
		return errors.New("参数错误")
	case NetDvrChanException:
		return errors.New("服务器通道处于错误状态")
	case NetDvrNodisk:
		return errors.New("没有硬盘")
	case NetDvrErrordisknum:
		return errors.New("硬盘号错误")
	case NetDvrDiskFull:
		return errors.New("服务器硬盘满")
	case NetDvrDiskError:
		return errors.New("服务器硬盘出错")
	case NetDvrNosupport:
		return errors.New("服务器不支持")
	case NetDvrBusy:
		return errors.New("服务器忙")
	case NetDvrModifyFail:
		return errors.New("服务器修改不成功")
	case NetDvrPasswordFormatError:
		return errors.New("密码输入格式不正确")
	case NetDvrDiskFormating:
		return errors.New("硬盘正在格式化，不能启动操作")
	case NetDvrDvrnoresource:
		return errors.New("DVR资源不足")
	case NetDvrDvropratefailed:
		return errors.New("DVR操作失败")
	case NetDvrOpenhostsoundFail:
		return errors.New("打开PC声音失败")
	case NetDvrDvrvoiceopened:
		return errors.New("服务器语音对讲被占用")
	case NetDvrTimeinputerror:
		return errors.New("时间输入不正确")
	case NetDvrNospecfile:
		return errors.New("回放时服务器没有指定的文件")
	case NetDvrCreatefileError:
		return errors.New("创建文件出错")
	case NetDvrFileopenfail:
		return errors.New("打开文件出错")
	case NetDvrOpernotfinish:
		return errors.New("上次的操作还没有完成")
	case NetDvrGetplaytimefail:
		return errors.New("获取当前播放的时间出错")
	case NetDvrPlayfail:
		return errors.New("播放出错")
	case NetDvrFileformatError:
		return errors.New("文件格式不正确")
	case NetDvrDirError:
		return errors.New("路径错误")
	case NetDvrAllocResourceError:
		return errors.New("资源分配错误")
	case NetDvrAudioModeError:
		return errors.New("声卡模式错误")
	case NetDvrNoenoughBuf:
		return errors.New("缓冲区太小")
	case NetDvrCreatesocketError:
		return errors.New("创建SOCKET出错")
	case NetDvrSetsocketError:
		return errors.New("设置SOCKET出错")
	case NetDvrMaxNum:
		return errors.New("个数达到最大")
	case NetDvrUsernotexist:
		return errors.New("用户不存在")
	case NetDvrWriteflasherror:
		return errors.New("写FLASH出错")
	case NetDvrUpgradefail:
		return errors.New("DVR升级失败")
	case NetDvrCardhaveinit:
		return errors.New("解码卡已经初始化过")
	case NetDvrPlayerfailed:
		return errors.New("调用播放库中某个函数失败")
	case NetDvrMaxUsernum:
		return errors.New("设备端用户数达到最大")
	case NetDvrGetlocalipandmacfail:
		return errors.New("获得客户端的IP地址或物理地址失败")
	case NetDvrNoencodeing:
		return errors.New("该通道没有编码")
	case NetDvrIpmismatch:
		return errors.New("IP地址不匹配")
	case NetDvrMacmismatch:
		return errors.New("MAC地址不匹配")
	case NetDvrUpgradelangmismatch:
		return errors.New("升级文件语言不匹配")
	case NetDvrMaxPlayerport:
		return errors.New("播放器路数达到最大")
	case NetDvrNospacebackup:
		return errors.New("备份设备中没有足够空间进行备份")
	case NetDvrNodevicebackup:
		return errors.New("没有找到指定的备份设备")
	case NetDvrPictureBitsError:
		return errors.New("图像素位数不符，限24色")
	case NetDvrPictureDimensionError:
		return errors.New("图片高*宽超限， 限128*256")
	case NetDvrPictureSizError:
		return errors.New("图片大小超限，限100K")
	case NetDvrLoadplayersdkfailed:
		return errors.New("载入当前目录下Player Sdk出错")
	case NetDvrLoadplayersdkprocError:
		return errors.New("找不到Player Sdk中某个函数入口")
	case NetDvrLoaddssdkfailed:
		return errors.New("载入当前目录下DSsdk出错")
	case NetDvrLoaddssdkprocError:
		return errors.New("找不到DsSdk中某个函数入口")
	case NetDvrDssdkError:
		return errors.New("调用硬解码库DsSdk中某个函数失败")
	case NetDvrVoicemonopolize:
		return errors.New("声卡被独占")
	case NetDvrJoinmulticastfailed:
		return errors.New("加入多播组失败")
	case NetDvrCreatedirError:
		return errors.New("建立日志文件目录失败")
	case NetDvrBindsocketError:
		return errors.New("绑定套接字失败")
	case NetDvrSocketcloseError:
		return errors.New("socket连接中断，此错误通常是由于连")
	case NetDvrUseridIsusing:
		return errors.New("注销时用户ID正在进行某操作")
	case NetDvrSocketlistenError:
		return errors.New("监听失败")
	case NetDvrProgramException:
		return errors.New("程序异常")
	case NetDvrWritefileFailed:
		return errors.New("写文件失败")
	case NetDvrFormatReadonly:
		return errors.New("禁止格式化只读硬盘")
	case NetDvrWithsameusername:
		return errors.New("用户配置结构中存在相同的用户名")
	case NetDvrDevicetypeError:
		return errors.New("导入参数时设备型号不匹配")
	case NetDvrLanguageError:
		return errors.New("导入参数时语言不匹配")
	case NetDvrParaversionError:
		return errors.New("导入参数时软件版本不匹配")
	case NetDvrIpchanNotalive:
		return errors.New("预览时外接IP通道不在线")
	case NetDvrRtspSdkError:
		return errors.New("加载高清IPC通讯库StreamTransClient.")
	case NetDvrConvertSdkError:
		return errors.New("加载转码库失败")
	case NetDvrIpcCountOverflow:
		return errors.New("超出最大的ip接入通道数")
	case NetDvrMaxAddNum:
		return errors.New("添加标签(一个文件片段64)等个数达到")
	case NetDvrParammodeError:
		return errors.New("图像增强仪，参数模式错误（用于硬件")
	case NetDvrCodespitterOffline:
		return errors.New("视频综合平台，码分器不在线")
	case NetDvrBackupCopying:
		return errors.New("设备正在备份")
	case NetDvrChanNotsupport:
		return errors.New(" 通道不支持该操作")
	case NetDvrCallineinvalid:
		return errors.New(" 高度线位置太集中或长度线不够倾斜")
	case NetDvrCalcancelconflict:
		return errors.New(" 取消标定冲突，如果设置了规则及全局")
	case NetDvrCalpointoutrange:
		return errors.New(" 标定点超出范围")
	case NetDvrFilterrectinvalid:
		return errors.New(" 尺寸过滤器不符合要求")
	case NetDvrDdnsDevoffline:
		return errors.New("设备没有注册到ddns上")
	case NetDvrDdnsInterError:
		return errors.New("DDNS 服务器内部错误")
	case NetDvrFunctionNotSupportOs:
		return errors.New("此功能不支持该操作系统")
	case NetDvrDecChanRebind:
		return errors.New("解码通道绑定显示输出次数受限")
	case NetDvrIntercomSdkError:
		return errors.New("加载当前目录下的语音对讲库失败")
	case NetDvrNoCurrentUpdatefile:
		return errors.New("没有正确的升级包")
	case NetDvrUserNotSuccLogin:
		return errors.New("用户还没登陆成功")
	case NetDvrUseLogSwitchFile:
		return errors.New("正在使用日志开关文件")
	case NetDvrPoolPortExhaust:
		return errors.New("端口池中用于绑定的端口已耗尽")
	case NetDvrPacketTypeNotSupport:
		return errors.New("码流封装格式错误")
	case NetDvrIpparaIpidError:
		return errors.New("IP接入配置时IPID有误")
	case NetDvrLoadHcpreviewSdkError:
		return errors.New("预览组件加载失败")
	case NetDvrLoadHcvoicetalkSdkError:
		return errors.New("语音组件加载失败")
	case NetDvrLoadHcalarmSdkError:
		return errors.New("报警组件加载失败")
	case NetDvrLoadHcplaybackSdkError:
		return errors.New("回放组件加载失败")
	case NetDvrLoadHcdisplaySdkError:
		return errors.New("显示组件加载失败")
	case NetDvrLoadHcindustrySdkError:
		return errors.New("行业应用组件加载失败")
	case NetDvrLoadHcgeneralcfgmgrSdkError:
		return errors.New("通用配置管理组件加载失败")
	case NetDvrLoadHccoredevcfgSdkError:
		return errors.New("设备配置核心组件加载失败")
	case NetDvrLoadHcnetutilsSdkError:
		return errors.New("HCNetUtils加载失败")
	case NetDvrCoreVerMismatch:
		return errors.New("单独加载组件时，组件与core")
	case NetDvrCoreVerMismatchHcpreview:
		return errors.New("预览组件与core版本不匹配")
	case NetDvrCoreVerMismatchHcvoicetalk:
		return errors.New("语音组件与core版本不匹配")
	case NetDvrCoreVerMismatchHcalarm:
		return errors.New("报警组件与core版本不匹配")
	case NetDvrCoreVerMismatchHcplayback:
		return errors.New("回放组件与core版本不匹配")
	case NetDvrCoreVerMismatchHcdisplay:
		return errors.New("显示组件与core版本不匹配")
	case NetDvrCoreVerMismatchHcindustry:
		return errors.New("行业应用组件与core版本不匹")
	case NetDvrCoreVerMismatchHcgeneralcfgmgr:
		return errors.New("通用配置管理组件与core版本")
	case NetDvrComVerMismatchHcpreview:
		return errors.New("预览组件与HCNetSDK版本不匹配")
	case NetDvrComVerMismatchHcvoicetalk:
		return errors.New("语音组件与HCNetSDK版本不匹配")
	case NetDvrComVerMismatchHcalarm:
		return errors.New("报警组件与HCNetSDK版本不匹配")
	case NetDvrComVerMismatchHcplayback:
		return errors.New("回放组件与HCNetSDK版本不匹配")
	case NetDvrComVerMismatchHcdisplay:
		return errors.New("显示组件与HCNetSDK版本不匹配")
	case NetDvrComVerMismatchHcindustry:
		return errors.New("行业应用组件与HCNetSDK版本不")
	case NetDvrComVerMismatchHcgeneralcfgmgr:
		return errors.New("通用配置管理组件与HCNetSDK版")
	case NetErrConfigFileImportFailed:
		return errors.New("配置文件导入失败")
	case NetErrConfigFileExportFailed:
		return errors.New("配置文件导出失败")
	case NetDvrCertificateFileError:
		return errors.New("证书错误")
	case NetDvrLoadSslLibError:
		return errors.New("加载SSL库失败（可能是版本不匹配，也")
	case NetDvrSslVersionNotMatch:
		return errors.New("SSL库版本不匹配")
	case NetDvrAliasDuplicate:
		return errors.New("别名重复  //2011-08-31 通过别名或者序")
	case NetDvrInvalidCommunication:
		return errors.New("无效通信")
	case NetDvrUsernameNotExist:
		return errors.New("用户名不存在（用户名不存在，IPC5.1.7")
	case NetDvrUserLocked:
		return errors.New("用户被锁定")
	case NetDvrInvalidUserid:
		return errors.New("无效用户ID")
	case NetDvrLowLoginVersion:
		return errors.New("登录版本低")
	case NetDvrLoadLibeay32DllError:
		return errors.New("加载libeay32.dll库失败")
	case NetDvrLoadSsleay32DllError:
		return errors.New("加载ssleay32.dll库失败")
	case NetErrLoadLibiconv:
		return errors.New("加载libiconv库失败")
	case NetErrSslConnectFailed:
		return errors.New("SSL连接失败")
	case NetErrMcastAddressError:
		return errors.New("获取多播地址错误")
	case NetErrLoadZlib:
		return errors.New("加载zlib.dll库失败")
	case NetErrOpensslNoInit:
		return errors.New("Openssl库未初始化")
	case NetDvrServerNotExist:
		return errors.New("对应的服务器找不到,查")
	case NetDvrTestServerFailConnect:
		return errors.New("连接测试服务器失败")
	case NetDvrNasServerInvalidDir:
		return errors.New("NAS服务器挂载目录失败")
	case NetDvrNasServerNoenoughPri:
		return errors.New("NAS服务器挂载目录失败")
	case NetDvrEmailServerNotConfigDns:
		return errors.New("服务器使用域名，但是没")
	case NetDvrEmailServerNotConfigGateway:
		return errors.New("没有配置网关，可能造成")
	case NetDvrTestServerPasswordError:
		return errors.New("用户名密码不正确，测试")
	case NetDvrEmailServerConnectExceptionWithSmtp:
		return errors.New("设备和smtp服务器交互异")
	case NetDvrFtpServerFailCreateDir:
		return errors.New("FTP服务器创建目录失败")
	case NetDvrFtpServerNoWritePir:
		return errors.New("FTP服务器没有写入权限")
	case NetDvrIpConflict:
		return errors.New("IP冲突")
	case NetDvrInsufficientStoragepoolSpace:
		return errors.New("存储池空间已满")
	case NetDvrStoragepoolInvalid:
		return errors.New("云服务器存储池无效,没")
	case NetDvrEffectivenessReboot:
		return errors.New("生效需要重启")
	case NetErrAnrArmingExist:
		return errors.New("断网续传布防连接已经存")
	case NetErrUploadlinkExist:
		return errors.New("断网续传上传连接已经存")
	case NetErrIncorrectFileFormat:
		return errors.New("导入文件格式不正确")
	case NetErrIncorrectFileContent:
		return errors.New("导入文件内容不正确")
	case NetErrMaxHrudpLink:
		return errors.New("HRUDP 连接数 超过设备")
	case NetSdkErrAccesskeySecretkey:
		return errors.New(" 接入秘钥或加密秘钥不")
	case NetSdkErrCreatePortMultiplex:
		return errors.New("创建端口复用失败")
	case NetDvrNonblockingCaptureNotsupport:
		return errors.New("不支持无阻塞抓图")
	case NetSdkErrFunctionInvalid:
		return errors.New("已开启异步，该功能无效")
	case NetSdkErrMaxPortMultiplex:
		return errors.New("已达到端口复用最大数目")
	default:
		return errors.New("未知错误")
	}
}
