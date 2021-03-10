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
	NetDvrDevicetypeError        = 80  // 导入参数时设备型号不匹配
	NetDvrLanguageError          = 81  // 导入参数时语言不匹配
	NetDvrParaversionError       = 82  // 导入参数时软件版本不匹配
	NetDvrIpchanNotalive         = 83  // 预览时外接IP通道不在线
	NetDvrRtspSdkError           = 84  // 加载高清IPC通讯库StreamTransClient.dll失败
	NetDvrConvertSdkError        = 85  // 加载转码库失败
	NetDvrIpcCountOverflow       = 86  // 超出最大的ip接入通道数
	NetDvrMaxAddNum              = 87  // 添加标签(一个文件片段64)等个数达到最大
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
	// 阵列错误码
	RAID_ERROR_INDEX                = 200
	NET_DVR_NAME_NOT_ONLY           = (RAID_ERROR_INDEX + 0)  // 名称已存在
	NET_DVR_OVER_MAX_ARRAY          = (RAID_ERROR_INDEX + 1)  // 阵列达到上限
	NET_DVR_OVER_MAX_VD             = (RAID_ERROR_INDEX + 2)  // 虚拟磁盘达到上限
	NET_DVR_VD_SLOT_EXCEED          = (RAID_ERROR_INDEX + 3)  // 虚拟磁盘槽位已满
	NET_DVR_PD_STATUS_INVALID       = (RAID_ERROR_INDEX + 4)  // 重建阵列所需物理磁盘状态错误
	NET_DVR_PD_BE_DEDICATE_SPARE    = (RAID_ERROR_INDEX + 5)  // 重建阵列所需物理磁盘为指定热备
	NET_DVR_PD_NOT_FREE             = (RAID_ERROR_INDEX + 6)  // 重建阵列所需物理磁盘非空闲
	NET_DVR_CANNOT_MIG2NEWMODE      = (RAID_ERROR_INDEX + 7)  // 不能从当前的阵列类型迁移到新的阵列类型
	NET_DVR_MIG_PAUSE               = (RAID_ERROR_INDEX + 8)  // 迁移操作已暂停
	NET_DVR_MIG_CANCEL              = (RAID_ERROR_INDEX + 9)  // 正在执行的迁移操作已取消
	NET_DVR_EXIST_VD                = (RAID_ERROR_INDEX + 10) // 阵列上阵列上存在虚拟磁盘，无法删除阵列
	NET_DVR_TARGET_IN_LD_FUNCTIONAL = (RAID_ERROR_INDEX + 11) // 对象物理磁盘为虚拟磁盘组成部分且工作正常
	NET_DVR_HD_IS_ASSIGNED_ALREADY  = (RAID_ERROR_INDEX + 12) // 指定的物理磁盘被分配为虚拟磁盘
	NET_DVR_INVALID_HD_COUNT        = (RAID_ERROR_INDEX + 13) // 物理磁盘数量与指定的RAID等级不匹配
	NET_DVR_LD_IS_FUNCTIONAL        = (RAID_ERROR_INDEX + 14) // 阵列正常，无法重建
	NET_DVR_BGA_RUNNING             = (RAID_ERROR_INDEX + 15) // 存在正在执行的后台任务
	NET_DVR_LD_NO_ATAPI             = (RAID_ERROR_INDEX + 16) // 无法用ATAPI盘创建虚拟磁盘
	NET_DVR_MIGRATION_NOT_NEED      = (RAID_ERROR_INDEX + 17) // 阵列无需迁移
	NET_DVR_HD_TYPE_MISMATCH        = (RAID_ERROR_INDEX + 18) // 物理磁盘不属于同意类型
	NET_DVR_NO_LD_IN_DG             = (RAID_ERROR_INDEX + 19) // 无虚拟磁盘，无法进行此项操作
	NET_DVR_NO_ROOM_FOR_SPARE       = (RAID_ERROR_INDEX + 20) // 磁盘空间过小，无法被指定为热备盘
	NET_DVR_SPARE_IS_IN_MULTI_DG    = (RAID_ERROR_INDEX + 21) // 磁盘已被分配为某阵列热备盘
	NET_DVR_DG_HAS_MISSING_PD       = (RAID_ERROR_INDEX + 22) // 阵列缺少盘

	// x86 64bit nvr新增 2012-02-04
	NET_DVR_NAME_EMPTY              = (RAID_ERROR_INDEX + 23) // 名称为空
	NET_DVR_INPUT_PARAM             = (RAID_ERROR_INDEX + 24) // 输入参数有误
	NET_DVR_PD_NOT_AVAILABLE        = (RAID_ERROR_INDEX + 25) // 物理磁盘不可用
	NET_DVR_ARRAY_NOT_AVAILABLE     = (RAID_ERROR_INDEX + 26) // 阵列不可用
	NET_DVR_PD_COUNT                = (RAID_ERROR_INDEX + 27) // 物理磁盘数不正确
	NET_DVR_VD_SMALL                = (RAID_ERROR_INDEX + 28) // 虚拟磁盘太小
	NET_DVR_NO_EXIST                = (RAID_ERROR_INDEX + 29) // 不存在
	NET_DVR_NOT_SUPPORT             = (RAID_ERROR_INDEX + 30) // 不支持该操作
	NET_DVR_NOT_FUNCTIONAL          = (RAID_ERROR_INDEX + 31) // 阵列状态不是正常状态
	NET_DVR_DEV_NODE_NOT_FOUND      = (RAID_ERROR_INDEX + 32) // 虚拟磁盘设备节点不存在
	NET_DVR_SLOT_EXCEED             = (RAID_ERROR_INDEX + 33) // 槽位达到上限
	NET_DVR_NO_VD_IN_ARRAY          = (RAID_ERROR_INDEX + 34) // 阵列上不存在虚拟磁盘
	NET_DVR_VD_SLOT_INVALID         = (RAID_ERROR_INDEX + 35) // 虚拟磁盘槽位无效
	NET_DVR_PD_NO_ENOUGH_SPACE      = (RAID_ERROR_INDEX + 36) // 所需物理磁盘空间不足
	NET_DVR_ARRAY_NONFUNCTION       = (RAID_ERROR_INDEX + 37) // 只有处于正常状态的阵列才能进行迁移
	NET_DVR_ARRAY_NO_ENOUGH_SPACE   = (RAID_ERROR_INDEX + 38) // 阵列空间不足
	NET_DVR_STOPPING_SCANNING_ARRAY = (RAID_ERROR_INDEX + 39) // 正在执行安全拔盘或重新扫描
	NET_DVR_NOT_SUPPORT_16T         = (RAID_ERROR_INDEX + 40) // 不支持创建大于16T的阵列
	NET_DVR_ARRAY_FORMATING         = (RAID_ERROR_INDEX + 41) // 正在执行格式化的阵列无法删除
	NET_DVR_QUICK_SETUP_PD_COUNT    = (RAID_ERROR_INDEX + 42) // 一键配置至少需要三块空闲盘

	//设备未激活时，登录失败，返回错误码
	NET_DVR_ERROR_DEVICE_NOT_ACTIVATED = 250 //设备未激活
	//老SDK接新设备，设置用户密码或者激活的时候为风险密码时，错误码
	NET_DVR_ERROR_RISK_PASSWORD = 251 //有风险的密码
	//已激活的设备，再次激活时返回错误码
	NET_DVR_ERROR_DEVICE_HAS_ACTIVATED = 252 //设备已激活

	// 智能错误码
	VCA_ERROR_INDEX                   = 300                    // 智能错误码索引
	NET_DVR_ID_ERROR                  = (VCA_ERROR_INDEX + 0)  // 配置ID不合理
	NET_DVR_POLYGON_ERROR             = (VCA_ERROR_INDEX + 1)  // 多边形不符合要求
	NET_DVR_RULE_PARAM_ERROR          = (VCA_ERROR_INDEX + 2)  // 规则参数不合理
	NET_DVR_RULE_CFG_CONFLICT         = (VCA_ERROR_INDEX + 3)  // 配置信息冲突
	NET_DVR_CALIBRATE_NOT_READY       = (VCA_ERROR_INDEX + 4)  // 当前没有标定信息
	NET_DVR_CAMERA_DATA_ERROR         = (VCA_ERROR_INDEX + 5)  // 摄像机参数不合理
	NET_DVR_CALIBRATE_DATA_UNFIT      = (VCA_ERROR_INDEX + 6)  // 长度不够倾斜，不利于标定
	NET_DVR_CALIBRATE_DATA_CONFLICT   = (VCA_ERROR_INDEX + 7)  // 标定出错，以为所有点共线或者位置太集中
	NET_DVR_CALIBRATE_CALC_FAIL       = (VCA_ERROR_INDEX + 8)  // 摄像机标定参数值计算失败
	NET_DVR_CALIBRATE_LINE_OUT_RECT   = (VCA_ERROR_INDEX + 9)  // 输入的样本标定线超出了样本外接矩形框
	NET_DVR_ENTER_RULE_NOT_READY      = (VCA_ERROR_INDEX + 10) // 没有设置进入区域
	NET_DVR_AID_RULE_NO_INCLUDE_LANE  = (VCA_ERROR_INDEX + 11) // 交通事件规则中没有包括车道（特值拥堵和逆行）
	NET_DVR_LANE_NOT_READY            = (VCA_ERROR_INDEX + 12) // 当前没有设置车道
	NET_DVR_RULE_INCLUDE_TWO_WAY      = (VCA_ERROR_INDEX + 13) // 事件规则中包含2种不同方向
	NET_DVR_LANE_TPS_RULE_CONFLICT    = (VCA_ERROR_INDEX + 14) // 车道和数据规则冲突
	NET_DVR_NOT_SUPPORT_EVENT_TYPE    = (VCA_ERROR_INDEX + 15) // 不支持的事件类型
	NET_DVR_LANE_NO_WAY               = (VCA_ERROR_INDEX + 16) // 车道没有方向
	NET_DVR_SIZE_FILTER_ERROR         = (VCA_ERROR_INDEX + 17) // 尺寸过滤框不合理
	NET_DVR_LIB_FFL_NO_FACE           = (VCA_ERROR_INDEX + 18) // 特征点定位时输入的图像没有人脸
	NET_DVR_LIB_FFL_IMG_TOO_SMALL     = (VCA_ERROR_INDEX + 19) // 特征点定位时输入的图像太小
	NET_DVR_LIB_FD_IMG_NO_FACE        = (VCA_ERROR_INDEX + 20) // 单张图像人脸检测时输入的图像没有人脸
	NET_DVR_LIB_FACE_TOO_SMALL        = (VCA_ERROR_INDEX + 21) // 建模时人脸太小
	NET_DVR_LIB_FACE_QUALITY_TOO_BAD  = (VCA_ERROR_INDEX + 22) // 建模时人脸图像质量太差
	NET_DVR_KEY_PARAM_ERR             = (VCA_ERROR_INDEX + 23) //高级参数设置错误
	NET_DVR_CALIBRATE_DATA_ERR        = (VCA_ERROR_INDEX + 24) //标定样本数目错误，或数据值错误，或样本点超出地平线
	NET_DVR_CALIBRATE_DISABLE_FAIL    = (VCA_ERROR_INDEX + 25) //所配置规则不允许取消标定
	NET_DVR_VCA_LIB_FD_SCALE_OUTRANGE = (VCA_ERROR_INDEX + 26) //最大过滤框的宽高最小值超过最小过滤框的宽高最大值两倍以上
	NET_DVR_LIB_FD_REGION_TOO_LARGE   = (VCA_ERROR_INDEX + 27) //当前检测区域范围过大。检测区最大为图像的2/3
	NET_DVR_TRIAL_OVERDUE             = (VCA_ERROR_INDEX + 28) //试用版评估期已结束
	NET_DVR_CONFIG_FILE_CONFLICT      = (VCA_ERROR_INDEX + 29) //设备类型与配置文件冲突（加密狗类型与现有分析仪配置不符错误码提示）
	//算法库相关错误码
	NET_DVR_FR_FPL_FAIL          = (VCA_ERROR_INDEX + 30) // 人脸特征点定位失败
	NET_DVR_FR_IQA_FAIL          = (VCA_ERROR_INDEX + 31) // 人脸评分失败
	NET_DVR_FR_FEM_FAIL          = (VCA_ERROR_INDEX + 32) // 人脸特征提取失败
	NET_DVR_FPL_DT_CONF_TOO_LOW  = (VCA_ERROR_INDEX + 33) // 特征点定位时人脸检测置信度过低
	NET_DVR_FPL_CONF_TOO_LOW     = (VCA_ERROR_INDEX + 34) // 特征点定位置信度过低
	NET_DVR_E_DATA_SIZE          = (VCA_ERROR_INDEX + 35) // 数据长度不匹配
	NET_DVR_FR_MODEL_VERSION_ERR = (VCA_ERROR_INDEX + 36) // 人脸模型数据中的模型版本错误
	NET_DVR_FR_FD_FAIL           = (VCA_ERROR_INDEX + 37) // 识别库中人脸检测失败
	NET_DVR_FA_NORMALIZE_ERR     = (VCA_ERROR_INDEX + 38) // 人脸归一化出错
	//其他错误码
	NET_DVR_DOG_PUSTREAM_NOT_MATCH     = (VCA_ERROR_INDEX + 39) // 加密狗与前端取流设备类型不匹配
	NET_DVR_DEV_PUSTREAM_NOT_MATCH     = (VCA_ERROR_INDEX + 40) // 前端取流设备版本不匹配
	NET_DVR_PUSTREAM_ALREADY_EXISTS    = (VCA_ERROR_INDEX + 41) // 设备的其他通道已经添加过该前端设备
	NET_DVR_SEARCH_CONNECT_FAILED      = (VCA_ERROR_INDEX + 42) // 连接检索服务器失败
	NET_DVR_INSUFFICIENT_DISK_SPACE    = (VCA_ERROR_INDEX + 43) // 可存储的硬盘空间不足
	NET_DVR_DATABASE_CONNECTION_FAILED = (VCA_ERROR_INDEX + 44) // 数据库连接失败
	NET_DVR_DATABASE_ADM_PW_ERROR      = (VCA_ERROR_INDEX + 45) // 数据库用户名、密码错误
	NET_DVR_DECODE_YUV                 = (VCA_ERROR_INDEX + 46) // 解码失败
	NET_DVR_IMAGE_RESOLUTION_ERROR     = (VCA_ERROR_INDEX + 47) //
	NET_DVR_CHAN_WORKMODE_ERROR        = (VCA_ERROR_INDEX + 48) //

	NET_DVR_RTSP_ERROR_NOENOUGHPRI    = 401 //无权限：服务器返回401时，转成这个错误码
	NET_DVR_RTSP_ERROR_ALLOC_RESOURCE = 402 //分配资源失败
	NET_DVR_RTSP_ERROR_PARAMETER      = 403 //参数错误
	NET_DVR_RTSP_ERROR_NO_URL         = 404 //指定的URL地址不存在：服务器返回404时，转成这个错误码
	NET_DVR_RTSP_ERROR_FORCE_STOP     = 406 //用户中途强行退出

	NET_DVR_RTSP_GETPORTFAILED        = 407 //rtsp 得到端口错误
	NET_DVR_RTSP_DESCRIBERROR         = 410 //rtsp decribe 交互错误
	NET_DVR_RTSP_DESCRIBESENDTIMEOUT  = 411 //rtsp decribe 发送超时
	NET_DVR_RTSP_DESCRIBESENDERROR    = 412 //rtsp decribe 发送失败
	NET_DVR_RTSP_DESCRIBERECVTIMEOUT  = 413 //rtsp decribe 接收超时
	NET_DVR_RTSP_DESCRIBERECVDATALOST = 414 //rtsp decribe 接收数据错误
	NET_DVR_RTSP_DESCRIBERECVERROR    = 415 //rtsp decribe 接收失败
	NET_DVR_RTSP_DESCRIBESERVERERR    = 416 //rtsp decribe 服务器返回错误状态

	NET_DVR_RTSP_SETUPERROR        = 420 //rtsp setup 交互错误
	NET_DVR_RTSP_SETUPSENDTIMEOUT  = 421 //rtsp setup 发送超时
	NET_DVR_RTSP_SETUPSENDERROR    = 422 //rtsp setup 发送错误
	NET_DVR_RTSP_SETUPRECVTIMEOUT  = 423 //rtsp setup 接收超时
	NET_DVR_RTSP_SETUPRECVDATALOST = 424 //rtsp setup 接收数据错误
	NET_DVR_RTSP_SETUPRECVERROR    = 425 //rtsp setup 接收失败
	NET_DVR_RTSP_OVER_MAX_CHAN     = 426 //超过服务器最大连接数，或者服务器资源不足，服务器返回453时，转成这个错误码。
	NET_DVR_RTSP_SETUPSERVERERR    = 427 //rtsp setup 服务器返回错误状态

	NET_DVR_RTSP_PLAYERROR        = 430 //rtsp play 交互错误
	NET_DVR_RTSP_PLAYSENDTIMEOUT  = 431 //rtsp play 发送超时
	NET_DVR_RTSP_PLAYSENDERROR    = 432 //rtsp play 发送错误
	NET_DVR_RTSP_PLAYRECVTIMEOUT  = 433 //rtsp play 接收超时
	NET_DVR_RTSP_PLAYRECVDATALOST = 434 //rtsp play 接收数据错误
	NET_DVR_RTSP_PLAYRECVERROR    = 435 //rtsp play 接收失败
	NET_DVR_RTSP_PLAYSERVERERR    = 436 //rtsp play 服务器返回错误状态

	NET_DVR_RTSP_TEARDOWNERROR        = 440 //rtsp teardown 交互错误
	NET_DVR_RTSP_TEARDOWNSENDTIMEOUT  = 441 //rtsp teardown 发送超时
	NET_DVR_RTSP_TEARDOWNSENDERROR    = 442 //rtsp teardown 发送错误
	NET_DVR_RTSP_TEARDOWNRECVTIMEOUT  = 443 //rtsp teardown 接收超时
	NET_DVR_RTSP_TEARDOWNRECVDATALOST = 444 //rtsp teardown 接收数据错误
	NET_DVR_RTSP_TEARDOWNRECVERROR    = 445 //rtsp teardown 接收失败
	NET_DVR_RTSP_TEARDOWNSERVERERR    = 446 //rtsp teardown 服务器返回错误状态

	NET_PLAYM4_NOERROR                = 500 //no error
	NET_PLAYM4_PARA_OVER              = 501 //input parameter is invalid;
	NET_PLAYM4_ORDER_ERROR            = 502 //The order of the function to be called is error.
	NET_PLAYM4_TIMER_ERROR            = 503 //Create multimedia clock failed;
	NET_PLAYM4_DEC_VIDEO_ERROR        = 504 //Decode video data failed.
	NET_PLAYM4_DEC_AUDIO_ERROR        = 505 //Decode audio data failed.
	NET_PLAYM4_ALLOC_MEMORY_ERROR     = 506 //Allocate memory failed.
	NET_PLAYM4_OPEN_FILE_ERROR        = 507 //Open the file failed.
	NET_PLAYM4_CREATE_OBJ_ERROR       = 508 //Create thread or event failed
	NET_PLAYM4_CREATE_DDRAW_ERROR     = 509 //Create DirectDraw object failed.
	NET_PLAYM4_CREATE_OFFSCREEN_ERROR = 510 //failed when creating off-screen surface.
	NET_PLAYM4_BUF_OVER               = 511 //buffer is overflow
	NET_PLAYM4_CREATE_SOUND_ERROR     = 512 //failed when creating audio device.
	NET_PLAYM4_SET_VOLUME_ERROR       = 513 //Set volume failed
	NET_PLAYM4_SUPPORT_FILE_ONLY      = 514 //The function only support play file.
	NET_PLAYM4_SUPPORT_STREAM_ONLY    = 515 //The function only support play stream.
	NET_PLAYM4_SYS_NOT_SUPPORT        = 516 //System not support.
	NET_PLAYM4_FILEHEADER_UNKNOWN     = 517 //No file header.
	NET_PLAYM4_VERSION_INCORRECT      = 518 //The version of decoder and encoder is not adapted.
	NET_PALYM4_INIT_DECODER_ERROR     = 519 //Initialize decoder failed.
	NET_PLAYM4_CHECK_FILE_ERROR       = 520 //The file data is unknown.
	NET_PLAYM4_INIT_TIMER_ERROR       = 521 //Initialize multimedia clock failed.
	NET_PLAYM4_BLT_ERROR              = 522 //Blt failed.
	NET_PLAYM4_UPDATE_ERROR           = 523 //Update failed.
	NET_PLAYM4_OPEN_FILE_ERROR_MULTI  = 524 //openfile error, streamtype is multi
	NET_PLAYM4_OPEN_FILE_ERROR_VIDEO  = 525 //openfile error, streamtype is video
	NET_PLAYM4_JPEG_COMPRESS_ERROR    = 526 //JPEG compress error
	NET_PLAYM4_EXTRACT_NOT_SUPPORT    = 527 //Don't support the version of this file.
	NET_PLAYM4_EXTRACT_DATA_ERROR     = 528 //extract video data failed.

	//转封装库错误码
	NET_CONVERT_ERROR_NOT_SUPPORT = 581 //convert not support

	//语音对讲库错误码
	NET_AUDIOINTERCOM_OK              = 600 //无错误
	NET_AUDIOINTECOM_ERR_NOTSUPORT    = 601 //不支持
	NET_AUDIOINTECOM_ERR_ALLOC_MEMERY = 602 //内存申请错误
	NET_AUDIOINTECOM_ERR_PARAMETER    = 603 //参数错误
	NET_AUDIOINTECOM_ERR_CALL_ORDER   = 604 //调用次序错误
	NET_AUDIOINTECOM_ERR_FIND_DEVICE  = 605 //未发现设备
	NET_AUDIOINTECOM_ERR_OPEN_DEVICE  = 606 //不能打开设备诶
	NET_AUDIOINTECOM_ERR_NO_CONTEXT   = 607 //设备上下文出错
	NET_AUDIOINTECOM_ERR_NO_WAVFILE   = 608 //WAV文件出错
	NET_AUDIOINTECOM_ERR_INVALID_TYPE = 609 //无效的WAV参数类型
	NET_AUDIOINTECOM_ERR_ENCODE_FAIL  = 610 //编码失败
	NET_AUDIOINTECOM_ERR_DECODE_FAIL  = 611 //解码失败
	NET_AUDIOINTECOM_ERR_NO_PLAYBACK  = 612 //播放失败
	NET_AUDIOINTECOM_ERR_DENOISE_FAIL = 613 //降噪失败
	NET_AUDIOINTECOM_ERR_UNKOWN       = 619 //未知错误

	NET_QOS_OK                                   = 700               //no error
	NET_QOS_ERROR                                = (NET_QOS_OK - 1)  //qos error
	NET_QOS_ERR_INVALID_ARGUMENTS                = (NET_QOS_OK - 2)  //invalid arguments
	NET_QOS_ERR_SESSION_NOT_FOUND                = (NET_QOS_OK - 3)  //session net found
	NET_QOS_ERR_LIB_NOT_INITIALIZED              = (NET_QOS_OK - 4)  //lib not initialized
	NET_QOS_ERR_OUTOFMEM                         = (NET_QOS_OK - 5)  //outtofmem
	NET_QOS_ERR_PACKET_UNKNOW                    = (NET_QOS_OK - 10) //packet unknow
	NET_QOS_ERR_PACKET_VERSION                   = (NET_QOS_OK - 11) //packet version error
	NET_QOS_ERR_PACKET_LENGTH                    = (NET_QOS_OK - 12) //packet length error
	NET_QOS_ERR_PACKET_TOO_BIG                   = (NET_QOS_OK - 13) //packet too big
	NET_QOS_ERR_SCHEDPARAMS_INVALID_BANDWIDTH    = (NET_QOS_OK - 20) //schedparams invalid bandwidth
	NET_QOS_ERR_SCHEDPARAMS_BAD_FRACTION         = (NET_QOS_OK - 21) //schedparams bad fraction
	NET_QOS_ERR_SCHEDPARAMS_BAD_MINIMUM_INTERVAL = (NET_QOS_OK - 22) //schedparams bad minimum interval

	NET_ERROR_TRUNK_LINE                      = 711 //子系统已被配成干线
	NET_ERROR_MIXED_JOINT                     = 712 //不能进行混合拼接
	NET_ERROR_DISPLAY_SWITCH                  = 713 //不能进行显示通道切换
	NET_ERROR_USED_BY_BIG_SCREEN              = 714 //解码资源被大屏占用
	NET_ERROR_USE_OTHER_DEC_RESOURCE          = 715 //不能使用其他解码子系统资源
	NET_ERROR_DISP_MODE_SWITCH                = 716 //显示通道显示状态切换中
	NET_ERROR_SCENE_USING                     = 717 //场景正在使用
	NET_ERR_NO_ENOUGH_DEC_RESOURCE            = 718 //解码资源不足
	NET_ERR_NO_ENOUGH_FREE_SHOW_RESOURCE      = 719 //畅显资源不足
	NET_ERR_NO_ENOUGH_VIDEO_MEMORY            = 720 //显存资源不足
	NET_ERR_MAX_VIDEO_NUM                     = 721 //一拖多资源不足
	NET_ERR_WIN_COVER_FREE_SHOW_AND_NORMAL    = 722 //窗口跨越了畅显输出口和非畅显输出口
	NET_ERR_FREE_SHOW_WIN_SPLIT               = 723 //畅显窗口不支持分屏
	NET_ERR_INAPPROPRIATE_WIN_FREE_SHOW       = 724 //不是输出口整数倍的窗口不支持开启畅显
	NET_DVR_TRANSPARENT_WIN_NOT_SUPPORT_SPLIT = 725 //开启透明度的窗口不支持分屏
	NET_DVR_SPLIT_WIN_NOT_SUPPORT_TRANSPARENT = 726 //开启多分屏的窗口不支持透明度设置
	NET_ERR_MAX_LOGO_NUM                      = 727 //logo数达到上限
	NET_ERR_MAX_WIN_LOOP_NUM                  = 728 //轮巡窗口数达到上限
	NET_ERR_VIRTUAL_LED_VERTICAL_CROSS        = 729 //虚拟LED不能纵向跨屏
	NET_ERR_MAX_VIRTUAL_LED_HEIGHT            = 730 //虚拟LED高度超限
	NET_ERR_VIRTUAL_LED_ILLEGAL_CHARACTER     = 731 //虚拟LED内容包含非法字符
	NET_ERR_BASEMAP_NOT_EXIST                 = 732 //底图图片不存在
	NET_ERR_LED_NOT_SUPPORT_VIRTUAL_LED       = 733 //LED屏幕不支持虚拟LED
	NET_ERR_LED_RESOLUTION_NOT_SUPPORT        = 734 //LED分辨率不支持
	NET_ERR_PLAN_OVERDUE                      = 735 //预案超期，不能再调用
	NET_ERR_PROCESSER_MAX_SCREEN_BLK          = 736 //单个处理器接入的信号跨越的屏幕个数超限
	NET_ERR_WND_SIZE_TOO_SMALL                = 737 //开窗窗口宽高太小
	NET_ERR_WND_SPLIT_NOT_SUPPORT_ROAM        = 738 //分屏窗口不支持漫游
	NET_ERR_OUTPUT_ONE_BOARD_ONE_WALL         = 739 //同一个子板的输出口只能绑定到同一面墙上
	NET_ERR_WND_CANNOT_LCD_AND_LED_OUTPUT     = 740 //窗口不能同时跨LCD和LED输出口
	NET_ERR_MAX_OSD_NUM                       = 741 //OSD数量达到最大

	NET_SDK_CANCEL_WND_TOPKEEP_ATTR_FIRST     = 751 //先取消置顶保持窗口的置顶保持属性才能进行置底操作
	NET_SDK_ERR_LED_SCREEN_CHECKING           = 752 //正在校正LED屏幕
	NET_SDK_ERR_NOT_SUPPORT_SINGLE_RESOLUTION = 753 //LCD/LED输出口绑定之后不支持单个输出口的分辨率配置
	NET_SDK_ERR_LED_RESOLUTION_MISMATCHED     = 754 //该输出口的LED分辨率和其他输出口的LED分辨率不匹配��需要满足同行等高、同列等宽

	NET_SDK_ERR_MAX_VIRTUAL_LED_WIDTH        = 755 //虚拟LED宽度超限，包括最大值和最小值
	NET_SDK_ERR_MAX_VIRTUAL_LED_IN_SCREEN    = 756 //单屏虚拟LED数量超限
	NET_SDK_ERR_MAX_VIRTUAL_LED_IN_WALL      = 757 //单墙虚拟LED数量超限
	NET_SDK_ERR_VIRTUAL_LED_OVERLAP          = 758 //虚拟LED重叠错误
	NET_SDK_ERR_VIRTUAL_LED_TYPE             = 759 //类型错误
	NET_SDK_ERR_VIRTUAL_LED_COLOUR           = 760 //颜色错误
	NET_SDK_ERR_VIRTUAL_LED_MOVE_DIRECTION   = 761 //移动方向错误
	NET_SDK_ERR_VIRTUAL_LED_MOVE_MODE        = 762 //移动模式错误
	NET_SDK_ERR_VIRTUAL_LED_MOVE_SPEED       = 763 //移动速度错误
	NET_SDK_ERR_VIRTUAL_LED_DISP_MODE        = 764 //显示模式有误
	NET_SDK_ERR_VIRTUAL_LED_NO               = 765 //虚拟LED序号错误
	NET_SDK_ERR_VIRTUAL_LED_PARA             = 766 //虚拟LED参数配置错误，包括结构体内其他参数
	NET_SDK_ERR_BASEMAP_POSITION             = 767 //底图窗口宽高参数错误
	NET_SDK_ERR_BASEMAP_PICTURE_LEN          = 768 //底图图片长度超限
	NET_SDK_ERR_BASEMAP_PICTURE_RESOLUTION   = 769 //底图图片分辨率错误
	NET_SDK_ERR_BASEMAP_PICTURE_FORMAT       = 770 //底图图片格式错误
	NET_SDK_ERR_MAX_VIRTUAL_LED_NUM          = 771 //设备支持的虚拟LED数量超限
	NET_SDK_ERR_MAX_TIME_VIRTUAL_LED_IN_WALL = 772 //单面电视墙支持的时间虚拟LED的数量超限

	NET_ERR_TERMINAL_BUSY = 780 //终端忙，终端处于会议中

	NET_ERR_DATA_RETURNED_ILLEGAL         = 790 //设备返回的数据不合法
	NET_DVR_FUNCTION_RESOURCE_USAGE_ERROR = 791 //设备其它功能占用资源，导致该功能无法开启

	NET_DVR_ERR_IMPORT_EMPTY_FILE        = 792 //导入文件为空
	NET_DVR_ERR_IMPORT_TOO_LARGE_FILE    = 793 //导入文件过大
	NET_DVR_ERR_BAD_IPV4_ADDRESS         = 794 //IPV4地址无效
	NET_DVR_ERR_BAD_NET_MASK             = 795 //子网掩码地址无效
	NET_DVR_ERR_INVALID_NET_GATE_ADDRESS = 796 //网关地址无效
	NET_DVR_ERR_BAD_DNS                  = 797 //DNS地址无效
	NET_DVR_ERR_ILLEGAL_PASSWORD         = 798 //密码不能包含用户名

	NET_DVR_DEV_NET_OVERFLOW                   = 800 //网络流量超过设备能力上限
	NET_DVR_STATUS_RECORDFILE_WRITING_NOT_LOCK = 801 //录像文件在录像，无法被锁定
	NET_DVR_STATUS_CANT_FORMAT_LITTLE_DISK     = 802 //由于硬盘太小无法格式化

	//N+1错误码
	NET_SDK_ERR_REMOTE_DISCONNECT  = 803 //远端无法连接
	NET_SDK_ERR_RD_ADD_RD          = 804 //备机不能添加备机
	NET_SDK_ERR_BACKUP_DISK_EXCEPT = 805 //备份盘异常
	NET_SDK_ERR_RD_LIMIT           = 806 //备机数已达上限
	NET_SDK_ERR_ADDED_RD_IS_WD     = 807 //添加的备机是工作机
	NET_SDK_ERR_ADD_ORDER_WRONG    = 808 //添加顺序出错，比如没有被工作机添加为备机，就添加工作机
	NET_SDK_ERR_WD_ADD_WD          = 809 //工作机不能添加工作机
	NET_SDK_ERR_WD_SERVICE_EXCETP  = 810 //工作机CVR服务异常
	NET_SDK_ERR_RD_SERVICE_EXCETP  = 811 //备机CVR服务异常
	NET_SDK_ERR_ADDED_WD_IS_RD     = 812 //添加的工作机是备机
	NET_SDK_ERR_PERFORMANCE_LIMIT  = 813 //性能达到上限
	NET_SDK_ERR_ADDED_DEVICE_EXIST = 814 //添加的设备已经存在

	//审讯机错误码
	NET_SDK_ERR_INQUEST_RESUMING  = 815 //审讯恢复中
	NET_SDK_ERR_RECORD_BACKUPING  = 816 //审讯备份中
	NET_SDK_ERR_DISK_PLAYING      = 817 //光盘回放中
	NET_SDK_ERR_INQUEST_STARTED   = 818 //审讯已开启
	NET_SDK_ERR_LOCAL_OPERATING   = 819 //本地操作进行中
	NET_SDK_ERR_INQUEST_NOT_START = 820 //审讯未开启
	//Netra3.1.0错误码
	NET_SDK_ERR_CHAN_AUDIO_BIND = 821 //通道未绑定或绑定语音对讲失败
	//云存储错误码
	NET_DVR_N_PLUS_ONE_MODE      = 822 //设备当前处于N+1模式
	NET_DVR_CLOUD_STORAGE_OPENED = 823 //云存储模式已开启

	NET_DVR_ERR_OPER_NOT_ALLOWED = 824 //设备处于N+0被接管状态，不允许该操作
	NET_DVR_ERR_NEED_RELOCATE    = 825 //设备处于N+0被接管状态，需要获取重定向信息，再重新操作

	//庭审主机错误码
	NET_SDK_ERR_IR_PORT_ERROR                        = 830 //红外输出口错误
	NET_SDK_ERR_IR_CMD_ERROR                         = 831 //红外输出口的命令号错误
	NET_SDK_ERR_NOT_INQUESTING                       = 832 //设备处于非审讯状态
	NET_SDK_ERR_INQUEST_NOT_PAUSED                   = 833 //设备处于非暂停状态
	NET_DVR_CHECK_PASSWORD_MISTAKE_ERROR             = 834 //校验密码错误
	NET_DVR_CHECK_PASSWORD_NULL_ERROR                = 835 //校验密码不能为空
	NET_DVR_UNABLE_CALIB_ERROR                       = 836 // 当前无法标定
	NET_DVR_PLEASE_CALIB_ERROR                       = 837 //请先完成标定
	NET_DVR_ERR_PANORAMIC_CAL_EMPTY                  = 838 //Flash中全景标定为空
	NET_DVR_ERR_CALIB_FAIL_PLEASEAGAIN               = 839 //标定失败，请重新标定(Calibration failed. Please calibrate again.)
	NET_DVR_ERR_DETECTION_LINE                       = 840 //规则线配置错误，请重新配置规则线，确保规则线位于红色区域内(Please set detection line again. The detection line should be within the red count area.)
	NET_DVR_ERR_TURN_OFF_IMAGE_PARA                  = 841 //请先关闭图像参数切换功能(Please turn off the image parameters switch first.)
	NET_DVR_EXCEED_FACE_IMAGES_ERROR                 = 843 //超过人脸图片最大张数
	NET_DVR_ANALYSIS_FACE_IMAGES_ERROR               = 844 //图片数据识别失败
	NET_ERR_ALARM_INPUT_OCCUPIED                     = 845 //A<-1报警号已用于触发车辆抓拍Alarm Input No. A<-1 is used to trigger vehicle capture.
	NET_DVR_FACELIB_DATABASE_ERROR                   = 846 //人脸库中数据库版本不匹配
	NET_DVR_FACELIB_DATA_ERROR                       = 847 //人脸库数据错误
	NET_DVR_FACE_DATA_ID_ERROR                       = 848 //人脸数据PID无效
	NET_DVR_FACELIB_ID_ERROR                         = 849 //人脸库ID无效
	NET_DVR_EXCEED_FACE_LIBARY_ERROR                 = 850 //超过人脸库最大个数
	NET_DVR_PIC_ANALYSIS_NO_TARGET_ERROR             = 851 //图片未识别到目标
	NET_DVR_SUBPIC_ANALYSIS_MODELING_ERROR           = 852 //子图建模失败
	NET_DVR_PIC_ANALYSIS_NO_RESOURCE_ERROR           = 853 //无对应智能分析引擎支持图片二次识别
	NET_DVR_ANALYSIS_ENGINES_NO_RESOURCE_ERROR       = 854 //无分析引擎资源
	NET_DVR_ANALYSIS_ENGINES_USAGE_EXCEED_ERROR      = 855 //引擎使用率超负荷，已达100%
	NET_DVR_EXCEED_HUMANMISINFO_FILTER_ENABLED_ERROR = 856 //超过开启人体去误报最大通道个数
	NET_DVR_NAME_ERROR                               = 857 //名称错误
	NET_DVR_NAME_EXIST_ERROR                         = 858 //名称已存在
	NET_DVR_FACELIB_PIC_IMPORTING_ERROR              = 859 //人脸库导入图片中
	NET_DVR_ERR_CALIB_POSITION                       = 860 //标定位置超出摄像机运动范围
	NET_DVR_ERR_DELETE                               = 861 //无法删除
	NET_DVR_ERR_SCENE_ID                             = 862 //场景ID无效
	NET_DVR_ERR_CALIBING                             = 863 //标定中
	NET_DVR_PIC_FORMAT_ERROR                         = 864 //图片格式错误
	NET_DVR_PIC_RESOLUTION_INVALID_ERROR             = 865 //图片分辨率无效错误
	NET_DVR_PIC_SIZE_EXCEED_ERROR                    = 866 //图片过大
	NET_DVR_PIC_ANALYSIS_TARGRT_NUM_EXCEED_ERROR     = 867 //图片目标个数超过上限
	NET_DVR_ANALYSIS_ENGINES_LOADING_ERROR           = 868 //分析引擎初始化中
	NET_DVR_ANALYSIS_ENGINES_ABNORMA_ERROR           = 869 //分析引擎异常
	NET_DVR_ANALYSIS_ENGINES_FACELIB_IMPORTING       = 870 //分析引擎正在导入人脸库
	NET_DVR_NO_DATA_FOR_MODELING_ERROR               = 871 //无待建模数据
	NET_DVR_FACE_DATA_MODELING_ERROR                 = 872 //设备正在进行图片建模操作，不支持并发处理
	NET_ERR_FACELIBDATA_OVERLIMIT                    = 873 //超过设备中支持导入人脸数最大个数限制（导入的人脸库中数据）
	NET_DVR_ANALYSIS_ENGINES_ASSOCIATED_CHANNEL      = 874 //分析引擎已关联通道
	NET_DVR_ERR_CUSTOMID_LEN                         = 875 //上层自定义ID的长度最小32字符长度
	NET_DVR_ERR_CUSTOMFACELIBID_REPEAT               = 876 //上层下发重复的自定义人脸库ID
	NET_DVR_ERR_CUSTOMHUMANID_REPEAT                 = 877 //上层下发重复的自定义人员ID
	NET_DVR_ERR_URL_DOWNLOAD_FAIL                    = 878 //url下载失败
	NET_DVR_ERR_URL_DOWNLOAD_NOTSTART                = 879 //url未开始下载

	NET_DVR_CFG_FILE_SECRETKEY_ERROR = 880 //配置文件安全校验密钥错误
	NET_DVR_WDR_NOTDISABLE_ERROR     = 881 //请先关闭所有通道当前日夜参数转换模式下的宽动态
	NET_DVR_HLC_NOTDISABLE_ERROR     = 882 //请先关闭所有通道当前日夜参数转换模式下的强光抑制

	NET_DVR_THERMOMETRY_REGION_OVERSTEP_ERROR = 883 //测温区域越界

	NET_DVR_ERR_MODELING_DEVICEINTERNAL = 884 //建模失败，设备内部错误
	NET_DVR_ERR_MODELING_FACE           = 885 //建模失败，人脸建模错误
	NET_DVR_ERR_MODELING_FACEGRADING    = 886 //建模失败，人脸质量评分错误
	NET_DVR_ERR_MODELING_FACEGFEATURE   = 887 //建模失败，特征点提取错误
	NET_DVR_ERR_MODELING_FACEGANALYZING = 888 //建模失败，属性提取错误

	NET_DVR_ERR_STREAM_LIMIT            = 889 //码流性能超过上限，请减少取流路数
	NET_DVR_ERR_STREAM_DESCRIPTION      = 890 //请输入码流描述
	NET_DVR_ERR_STREAM_DELETE           = 891 //码流正在使用无法删除
	NET_DVR_ERR_CUSTOMSTREAM_NAME       = 892 //自定义码流名称为空或不合法
	NET_DVR_ERR_CUSTOMSTREAM_NOTEXISTED = 893 //该自定义码流不存在

	NET_DVR_ERR_TOO_SHORT_CALIBRATING_TIME = 894 //标定时间太短
	NET_DVR_ERR_AUTO_CALIBRATE_FAILED      = 895 //自动标定失败
	NET_DVR_ERR_VERIFICATION_FAILED        = 896 //校验失败

	NET_DVR_NO_TEMP_SENSOR_ERROR          = 897 //无温度传感器
	NET_DVR_PUPIL_DISTANCE_OVERSIZE_ERROR = 898 //瞳距过大
	NET_DVR_ERR_UNOPENED_FACE_SNAP        = 899 //操作无效，请先开启人脸抓拍
	//2011-10-25多屏控制器错误码（900-950）
	NET_ERR_CUT_INPUTSTREAM_OVERLIMIT   = 900 //信号源裁剪数值超限
	NET_ERR_WINCHAN_IDX                 = 901 // 开窗通道号错误
	NET_ERR_WIN_LAYER                   = 902 // 窗口层数错误，单个屏幕上最多覆盖的窗口层数
	NET_ERR_WIN_BLK_NUM                 = 903 // 窗口的块数错误，单个窗口可覆盖的屏幕个数
	NET_ERR_OUTPUT_RESOLUTION           = 904 // 输出分辨率错误
	NET_ERR_LAYOUT                      = 905 // 布局号错误
	NET_ERR_INPUT_RESOLUTION            = 906 // 输入分辨率不支持
	NET_ERR_SUBDEVICE_OFFLINE           = 907 // 子设备不在线
	NET_ERR_NO_DECODE_CHAN              = 908 // 没有空闲解码通道
	NET_ERR_MAX_WINDOW_ABILITY          = 909 // 开窗能力上限, 分布式多屏控制器中解码子设备能力上限或者显示处理器能力上限导致
	NET_ERR_ORDER_ERROR                 = 910 // 调用顺序有误
	NET_ERR_PLAYING_PLAN                = 911 // 正在执行预案
	NET_ERR_DECODER_USED                = 912 // 解码板正在使用
	NET_ERR_OUTPUT_BOARD_DATA_OVERFLOW  = 913 // 输出板数据量超限
	NET_ERR_SAME_USER_NAME              = 914 // 用户名相同
	NET_ERR_INVALID_USER_NAME           = 915 // 无效用户名
	NET_ERR_MATRIX_USING                = 916 // 输入矩阵正在使用
	NET_ERR_DIFFERENT_CHAN_TYPE         = 917 // 通道类型不同（矩阵输出通道和控制器的输入为不同的类型）
	NET_ERR_INPUT_CHAN_BINDED           = 918 // 输入通道已经被其他矩阵绑定
	NET_ERR_BINDED_OUTPUT_CHAN_OVERFLOW = 919 // 正在使用的矩阵输出通道个数超过矩阵与控制器绑定的通道个数
	NET_ERR_MAX_SIGNAL_NUM              = 920 // 输入信号源个数达到上限
	NET_ERR_INPUT_CHAN_USING            = 921 // 输入通道正在使用
	NET_ERR_MANAGER_LOGON               = 922 // 管理员已经登陆，操作失败
	NET_ERR_USERALREADY_LOGON           = 923 // 该用户已经登陆，操作失败
	NET_ERR_LAYOUT_INIT                 = 924 // 布局正在初始化，操作失败
	NET_ERR_BASEMAP_SIZE_NOT_MATCH      = 925 // 底图大小不符
	NET_ERR_WINDOW_OPERATING            = 926 // 窗口正在执行其他操作，本次操作失败
	NET_ERR_SIGNAL_UPLIMIT              = 927 // 信号源开窗个数达到上限
	NET_ERR_SIGNAL_MAX_ENLARGE_TIMES    = 928 // 信号源放大倍数超限
	NET_ERR_ONE_SIGNAL_MULTI_CROSS      = 929 // 单个信号源不能多次跨屏
	NET_ERR_ULTRA_HD_SIGNAL_MULTI_WIN   = 930 // 超高清信号源不能重复开窗
	NET_ERR_MAX_VIRTUAL_LED_WIDTH       = 931 //虚拟LED宽度大于限制值
	NET_ERR_MAX_VIRTUAL_LED_WORD_LEN    = 932 //虚拟LED字符数大于限制值
	NET_ERR_SINGLE_OUTPUTPARAM_CONFIG   = 933 //不支持单个显示输出参数设置
	NET_ERR_MULTI_WIN_BE_COVER          = 934 //多分屏窗口被覆盖
	NET_ERR_WIN_NOT_EXIST               = 935 //窗口不存在
	NET_ERR_WIN_MAX_SIGNALSOURCE        = 936 //窗口信号源数超过限制值
	NET_ERR_MULTI_WIN_MOVE              = 937 //对多分屏窗口移动
	NET_ERR_MULTI_WIN_YPBPR_SDI         = 938 // YPBPR 和SDI信号源不支持9/16分屏
	NET_ERR_DIFF_TYPE_OUTPUT_MIXUSE     = 939 //不同类型输出板混插
	NET_ERR_SPLIT_WIN_CROSS             = 940 //对跨屏窗口分屏
	NET_ERR_SPLIT_WIN_NOT_FULL_SCREEN   = 941 //对未满屏窗口分屏
	NET_ERR_SPLIT_WIN_MANY_WIN          = 942 //对单个输出口上有多个窗口的窗口分屏
	NET_ERR_WINDOW_SIZE_OVERLIMIT       = 943 //窗口大小超限
	NET_ERR_INPUTSTREAM_ALREADY_JOINT   = 944 //信号源已加入拼接
	NET_ERR_JOINT_INPUTSTREAM_OVERLIMIT = 945 //拼接信号源个数超限

	NET_ERR_LED_RESOLUTION                 = 946 //LED 分辨率大于输出分辨率
	NET_ERR_JOINT_SCALE_OVERLIMIT          = 947 //拼接信号源的规模超限
	NET_ERR_INPUTSTREAM_ALREADY_DECODE     = 948 //信号源已上墙
	NET_ERR_INPUTSTREAM_NOTSUPPORT_CAPTURE = 949 //信号源不支持抓图
	NET_ERR_JOINT_NOTSUPPORT_SPLITWIN      = 950 //拼接信号源不支持分屏

	//解码器错误码（951-999）
	NET_ERR_MAX_WIN_OVERLAP          = 951 //达到最大窗口重叠数
	NET_ERR_STREAMID_CHAN_BOTH_VALID = 952 //stream ID和通道号同时有效
	NET_ERR_NO_ZERO_CHAN             = 953 //设备无零通道
	NEED_RECONNECT                   = 955 //需要重定向（转码子系统使用）
	NET_ERR_NO_STREAM_ID             = 956 //流ID不存在
	NET_DVR_TRANS_NOT_START          = 957 //转码未启动
	NET_ERR_MAXNUM_STREAM_ID         = 958 //流ID数达到上限
	NET_ERR_WORKMODE_MISMATCH        = 959 //工作模式不匹配
	NET_ERR_MODE_IS_USING            = 960 //已工作在当前模式
	NET_ERR_DEV_PROGRESSING          = 961 //设备正在处理中
	NET_ERR_PASSIVE_TRANSCODING      = 962 //正在被动转码

	NET_ERR_RING_NOT_CONFIGURE = 964 //环网未配置

	NET_ERR_CLOSE_WINDOW_FIRST                  = 971 //切换全帧率畅显时必须先关闭对应的已上墙的窗口
	NET_ERR_SPLIT_WINDOW_NUM_NOT_SUPPORT        = 972 //VGA/DVI/DP/HDMI/HDBase_T输入源在全帧率畅显下不支持9/16画面
	NET_ERR_REACH_ONE_SIGNAL_PREVIEW_MAX_LINK   = 973 //单信号源回显连接数量超限
	NET_ERR_ONLY_SPLITWND_SUPPORT_AMPLIFICATION = 974 //只有分屏窗口支持子窗口放大
	NET_DVR_ERR_WINDOW_SIZE_PLACE               = 975 //窗口位置错误
	NET_DVR_ERR_RGIONAL_RESTRICTIONS            = 976 //屏幕距离超限
	NET_ERR_WNDZOOM_NOT_SUPPORT                 = 977 //单窗口不支持子窗口全屏功能
	NET_ERR_LED_SCREEN_SIZE                     = 978 //LED屏宽高不正确
	NET_ERR_OPEN_WIN_IN_ERROR_AREA              = 979 //在非法区域开窗(包括跨LCD/LED屏)
	NET_ERR_TITLE_WIN_NOT_SUPPORT_MOVE          = 980 //平铺模式不支持漫游
	NET_ERR_TITLE_WIN_NOT_SUPPORT_COVER         = 981 //平铺模式不支持图层覆盖
	NET_ERR_TITLE_WIN_NOT_SUPPORT_SPLIT         = 982 //平铺模式不支持分屏
	NET_DVR_LED_WINDOWS_ALREADY_CLOSED          = 983 //LED区域内输出口的分辨率发生变化，设备已关闭该区域内的所有LED窗口
	NET_DVR_ERR_CLOSE_WINDOWS                   = 984 //操作失败，请先关闭窗口
	NET_DVR_ERR_MATRIX_LOOP_ABILITY             = 985 //超出轮巡解码能力限制
	NET_DVR_ERR_MATRIX_LOOP_TIME                = 986 //轮巡解码时间不支持
	NET_DVR_ERR_LINKED_OUT_ABILITY              = 987 //联动通道数超过上限
	NET_ERR_REACH_SCENE_MAX_NUM                 = 988 //场景数量达到上限
	NET_ERR_SCENE_MEM_NOT_ENOUGH                = 989 //内存不足，无法新建场景
	NET_ERR_RESOLUTION_NOT_SUPPORT_ODD_VOUT     = 990 //奇口不支持该分辨率
	NET_ERR_RESOLUTION_NOT_SUPPORT_EVEN_VOUT    = 991 //偶口不支持该分辨率

	NET_DVR_CANCEL_WND_OPENKEEP_ATTR_FIRST        = 992 //常开窗口需要先取消常开属性才能被关闭
	NET_SDK_LED_MODE_NOT_SUPPORT_SPLIT            = 993 //LED模式下不支持窗口分屏
	NET_ERR_VOICETALK_ONLY_SUPPORT_ONE_TALK       = 994 //同时只支持一路语音对讲
	NET_ERR_WND_POSITION_ADJUSTED                 = 995 //窗口位置被设备调整，上层需要重新获取下窗口位置
	NET_SDK_ERR_STARTTIME_CANNOT_LESSTHAN_CURTIME = 996 //开始时间不能小于当前时间
	NET_SDK_ERR_NEED_ADJUST_PLAN                  = 997 //场景已被预案关联，请先将该场景从预案中删除
	NET_ERR_UnitConfig_Failed                     = 998 //当“启用单位统一”勾选时，测温下配置的单位与系统设置下的单位不同返回单位配置错误

	//能力集解析库错误码
	XML_ABILITY_NOTSUPPORT            = 1000 //不支持能力节点获取
	XML_ANALYZE_NOENOUGH_BUF          = 1001 //输出内存不足
	XML_ANALYZE_FIND_LOCALXML_ERROR   = 1002 //无法找到对应的本地xml
	XML_ANALYZE_LOAD_LOCALXML_ERROR   = 1003 //加载本地xml出错
	XML_NANLYZE_DVR_DATA_FORMAT_ERROR = 1004 //设备能力数据格式错误
	XML_ANALYZE_TYPE_ERROR            = 1005 //能力集类型错误
	XML_ANALYZE_XML_NODE_ERROR        = 1006 //XML能力节点格式错误
	XML_INPUT_PARAM_ERROR             = 1007 //输入的能力XML节点值错误

	NET_DVR_ERR_RETURNED_XML_DATA = 1008 //设备返回的XML数据有误

	//传显错误码
	NET_ERR_LEDAREA_EXIST_WINDOW   = 1051 //LED区域有窗口存在(如果LED区域上已经有窗口存在，不允许修改LED区域）
	NET_ERR_AUDIO_EXIST            = 1052 //输出口上存在音频输出，不允许解除绑定
	NET_ERR_MATERIAL_NAME_EXIST    = 1053 //素材名称已存在
	NET_ERR_MATERIAL_APPROVE_STATE = 1054 //素材审核状态错误
	NET_ERR_DATAHD_SIGNAL_FORMAT   = 1055 //已使用的硬盘不允许单个格式化

	NET_ERR_SCENE_SWITCHING                                 = 1056 //场景正在切换
	NER_ERR_DATA_TRANSFER                                   = 1057 //设备正在数据转移中
	NET_ERR_DATA_RESTORE                                    = 1058 //设备正在数据还原中
	NET_ERR_CHECK_NOT_ENABLE                                = 1059 //校正使能未开启
	NET_ERR_AREA_OFFLINE                                    = 1060 //区域不在线
	NET_ERR_SCREEN_TYPE                                     = 1061 //屏幕类型不匹配
	NET_ERR_MIN_OPERATE_UNIT                                = 1062 //最小操作单元不匹配
	NET_ERR_MAINHD_NOT_BACKUP                               = 1063 //第一槽位上的普通盘（主盘）禁止设置成备份盘
	NET_ERR_ONE_BACKUP_HD                                   = 1064 //设置备份盘时，设备至少有一块普通盘
	NET_ERR_CONNECT_SUB_SYSTEM_ABNORMAL                     = 1065 //连接子系统异常
	NET_ERR_SERIAL_PORT_VEST                                = 1066 //错误的串口归属
	NET_ERR_WHITE_LIST_FULL                                 = 1067 //白名单列表数量已满
	NET_ERR_NOT_MATCH_SOURCE                                = 1068 //不匹配的信号源类型
	NET_ERR_CLOCK_VIRTUAL_LED_FULL                          = 1069 //开启时钟功能的虚拟LED达上限
	NET_ERR_MAX_WIN_SIGNAL_LOOP_NUM                         = 1070 //窗口轮巡信号源个数已达上限
	NET_ERR_RESOLUTION_NO_MATCH_FRAME                       = 1071 //分辨率与当前帧数不匹配
	NET_ERR_NOT_UPDATE_LOW_VERSION                          = 1072 //不支持向低版本升级
	NET_ERR_NO_CUSTOM_TO_UPDATE                             = 1073 //非定制程序无法升级
	NET_ERR_CHAN_RESOLUTION_NOT_SUPPORT_SPLIT               = 1074 //该输出口分辨率不支持分屏
	NET_ERR_HIGH_DEFINITION_NOT_SUPPORT_SPLIT               = 1075 //超高清不支持9/16画面分割
	NET_ERR_MIRROR_IMAGE_BY_VIDEO_WALL                      = 1076 //电视墙镜像出错
	NET_ERR_MAX_OSD_FONT_SIZE                               = 1077 //超过OSD最大支持字符数
	NET_ERR_HIGH_DEFINITION_NOT_SUPPORT_VIDEO_SET           = 1078 //超清不支持视频参数设置
	NET_ERR_TILE_MODE_NOT_SUPPORT_JOINT                     = 1079 //平铺模式不支持拼接窗口
	NET_ERR_ADD_AUDIO_MATRIX_FAILED                         = 1080 //创建音频矩阵失败
	NET_ERR_ONE_VIRTUAL_LED_AREA_BIND_ONE_AUDIO_AREA        = 1081 //一个虚拟LED区域只能绑定一个音频区域
	NET_ERR_NAT_NOT_MODIFY_SERVER_NETWORK_PARAM             = 1082 //NAT下无法修改服务器网络参数
	NET_ERR_ORIGINAL_CHECH_DATA_ERROR                       = 1083 //原始校正数据错误
	NET_ERR_INPUT_BOARD_SPLICED_IN_DIFFERENT_NETWORKAREAS   = 1084 //不同网络区域的输入板不能拼接
	NET_ERR_SPLICINGSOURCE_ONWALL_IN_DIFFERENT_NETWORKAREAS = 1085 //不同网络区域的拼接源不能上墙
	NET_ERR_ONWALL_OUTPUTBOARD_MODIFY_NETWORKAREAS          = 1086 //已绑定在电视墙上的输出板不能修改网络区域
	NET_ERR_LAN_AND_WAN_CANNOT_SAME_NET_SEGMENT             = 1087 //LAN口IP和WAN口IP不能处于同一网段
	NET_ERR_USERNAME_REPETITIVE                             = 1088 //用户名重复
	NET_ERR_ASSOCIATED_SAMEWALL_IN_DIFFERENT_NETWORKAREAS   = 1089 //不同数据网络区域的输出口不能关联到同一电视墙
	NET_ERR_BASEMAP_ROAM_IN_LED_AREA                        = 1090 //LED区域不允许底图漫游
	NET_ERR_VIRTUAL_LED_NOT_SUPPORT_4K_OUTPUT               = 1091 //虚拟LED不支持4K输出口显示
	NET_ERR_BASEMAP_NOT_SUPPORT_4K_OUTPUT                   = 1092 //底图不支持4K输出口显示
	NET_ERR_MIN_BLOCK_IN_VIRTUAL_LED_AND_OUTPUT             = 1093 //虚拟LED与输出口相交出现最小块
	NET_ERR_485FIlE_VERSION_INVALID                         = 1094 //485文件版本无效
	NET_ERR_485FIlE_CHECK_ERROR                             = 1095 //485文件校验错误
	NET_ERR_485FIlE_ABNORMAL_SIZE                           = 1096 //485文件大小异常效
	NET_ERR_MODIFY_SUBBOARD_NETCFG_IN_NAT                   = 1097 //NAT下无法修改子板网络参
	NET_ERR_OSD_CONTENT_WITH_ILLEGAL_CHARACTERS             = 1098 //OSD内容包含非法字符
	NET_ERR_NON_SLAVE_DEVICE_INSERT_SYNC_LINE               = 1099 //非从设备禁止插入同步线
	//民用错误码（1100～1200）
	NET_ERR_PLT_USERID                    = 1100 //验证平台userid错误
	NET_ERR_TRANS_CHAN_START              = 1101 //透明通道已打开，当前操作无法完成
	NET_ERR_DEV_UPGRADING                 = 1102 //设备正在升级
	NET_ERR_MISMATCH_UPGRADE_PACK_TYPE    = 1103 //升级包类型不匹配
	NET_ERR_DEV_FORMATTING                = 1104 //设备正在格式化
	NET_ERR_MISMATCH_UPGRADE_PACK_VERSION = 1105 //升级包版本不匹配
	NET_ERR_PT_LOCKED                     = 1106 //PT锁定功能

	NET_DVR_LOGO_OVERLAY_WITHOUT_UPLOAD_PIC                          = 1110 //logo叠加失败，没有上传logo图片
	NET_DVR_ERR_ILLEGAL_VERIFICATION_CODE                            = 1111 //不合法验证码
	NET_DVR_ERR_LACK_VERIFICATION_CODE                               = 1112 //缺少验证码
	NET_DVR_ERR_FORBIDDEN_IP                                         = 1113 //该IP地址已被禁止，不允许配置(设备支持的IP地址过滤功能)
	NET_DVR_ERR_UNLOCKPTZ                                            = 1114 //操作无效，请先锁定云台
	NET_DVR_ERR_COUNTAREA_LARGE                                      = 1116 //计数区域绘制错误，区域面积大于设备允许值
	NET_DVR_ERR_LABEL_ID_EXCEED                                      = 1117 //标签ID超限
	NET_DVR_ERR_LABEL_TYPE                                           = 1118 //标签类型错误
	NET_DVR_ERR_LABEL_FULL                                           = 1119 //标签满
	NET_DVR_ERR_LABEL_DISABLED                                       = 1120 //标签未使能
	NET_DVR_ERR_DOME_PT_TRANS_TO_DOME_XY                             = 1121 //球机PT转球机XY失败
	NET_DVR_ERR_DOME_PT_TRANS_TO_PANORAMA_XY                         = 1122 //球机PT转全景XY失败
	NET_DVR_ERR_PANORAMA_XY_TRANS_TO_DOME_PT                         = 1123 //全景XY坐标转球机PT错误
	NET_DVR_ERR_SCENE_DUR_TIME_LESS_THAN_INTERV_TIME                 = 1124 //场景停留时间要大于检测周期
	NET_DVR_ERR_HTTP_BKN_EXCEED_ONE                                  = 1125 //断网续传布防只支持一路
	NET_DVR_ERR_DELETING_FAILED_TURN_OFF_HTTPS_ESDK_WEBSOCKETS_FIRST = 1126 //删除失败，请先关闭HTTPS和网络服务中的增强型SDK服务及WebSockets服务。
	NET_DVR_ERR_DELETING_FAILED_TURN_OFF_HTTPS_ESDK_FIRST            = 1127 //删除失败，请先关闭HTTPS和网络服务中的增强型SDK服务
	NET_DVR_ERR_PTZ_OCCUPIED_PRIORITY                                = 1128 // 有高优先级云台控制权限用户操作
	NET_DVR_ERR_INCORRECT_VIDEOAUDIO_ID                              = 1129 // 视频通道编码ID或语音输出通道编码ID错误
	NET_DVR_ERR_REPETITIONTIME_OVER_MAXIMUM                          = 1130 // 去重时间最大不超过最大值
	NET_DVR_ERR_FORMATTING_FAILED                                    = 1131 // 格式化错误，请重新
	NET_DVR_ERR_ENCRYPTED_FORMATTING_FAILED                          = 1132 // 加密格式化失败，请重试
	NET_DVR_ERR_WRONG_PASSWORD                                       = 1133 // 密码错误,请输入正确的密码（SD卡 密码校验失败）
	NET_DVR_ERR_EXPOSURE_SYNC                                        = 1134 // 镜头间曝光同步已开启，不允许配置手动RGB

	//2012-10-16 报警设备错误码（1200~1300）
	NET_ERR_SEARCHING_MODULE                  = 1201 // 正在搜索外接模块
	NET_ERR_REGISTERING_MODULE                = 1202 // 正在注册外接模块
	NET_ERR_GETTING_ZONES                     = 1203 // 正在获取防区参数
	NET_ERR_GETTING_TRIGGERS                  = 1204 // 正在获取触发器
	NET_ERR_ARMED_STATUS                      = 1205 // 系统处于布防状态
	NET_ERR_PROGRAM_MODE_STATUS               = 1206 // 系统处于编程模式
	NET_ERR_WALK_TEST_MODE_STATUS             = 1207 // 系统处于步测模式
	NET_ERR_BYPASS_STATUS                     = 1208 // 旁路状态
	NET_ERR_DISABLED_MODULE_STATUS            = 1209 // 功能未使能
	NET_ERR_NOT_SUPPORT_OPERATE_ZONE          = 1210 // 防区不支持该操作
	NET_ERR_NOT_SUPPORT_MOD_MODULE_ADDR       = 1211 // 模块地址不能被修改
	NET_ERR_UNREGISTERED_MODULE               = 1212 // 模块未注册
	NET_ERR_PUBLIC_SUBSYSTEM_ASSOCIATE_SELF   = 1213 // 公共子系统关联自身
	NET_ERR_EXCEEDS_ASSOCIATE_SUBSYSTEM_NUM   = 1214 // 超过公共子系统最大关联个数
	NET_ERR_BE_ASSOCIATED_BY_PUBLIC_SUBSYSTEM = 1215 // 子系统被其他公共子系统关联
	NET_ERR_ZONE_FAULT_STATUS                 = 1216 // 防区处于故障状态
	NET_ERR_SAME_EVENT_TYPE                   = 1217 // 事件触发报警输出开启和事件触发报警输出关闭中有相同事件类型
	NET_ERR_ZONE_ALARM_STATUS                 = 1218 // 防区处于报警状态
	NET_ERR_EXPANSION_BUS_SHORT_CIRCUIT       = 1219 //扩展总线短路
	NET_ERR_PWD_CONFLICT                      = 1220 //密码冲突
	NET_ERR_DETECTOR_GISTERED_BY_OTHER_ZONE   = 1221 //探测器已被其他防区注册
	NET_ERR_DETECTOR_GISTERED_BY_OTHER_PU     = 1222 //探测器已被其他主机注册
	NET_ERR_DETECTOR_DISCONNECT               = 1223 //探测器不在线
	NET_ERR_CALL_BUSY                         = 1224 //设备正在通话中
	NET_DVR_ERR_ZONE_TAMPER_STAUS             = 1225 //防区处于防拆状态
	NET_DVR_ERR_WIRELESS_DEV_REGISTER         = 1226 //无线外设已被其他主机注册
	NET_DVR_ERR_WIRELESS_DEV_ADDED            = 1227 //无线外设已被添加
	NET_DVR_ERR_WIRELESS_DEV_OFFLINE          = 1228 //无线外设不在线
	NET_DVR_ERR_WIRELESS_DEV_TAMPER_STATUS    = 1229 //无线外设处于防拆状态
	NET_DVR_ERR_GPRS_PHONE_CONFLICT           = 1230 //电话报警与无线报警中心冲突
	//信息发布主机
	NET_ERR_GET_ALL_RETURN_OVER = 1300 //获取所有返回数目超限
	NET_ERR_RESOURCE_USING      = 1301 //信息发布资源正在使用，不能修改
	NET_ERR_FILE_SIZE_OVERLIMIT = 1302 //文件大小超限

	//信息发布服务器错误码
	NET_ERR_MATERIAL_NAME                                                   = 1303 //素材名称非法
	NET_ERR_MATERIAL_NAME_LEN                                               = 1304 //素材名称长度非法
	NET_ERR_MATERIAL_REMARK                                                 = 1305 //素材描述非法
	NET_ERR_MATERIAL_REMARK_LEN                                             = 1306 //素材描述长度非法
	NET_ERR_MATERIAL_SHARE_PROPERTY                                         = 1307 //素材共享属性非法
	NET_ERR_UNSUPPORT_MATERIAL_TYPE                                         = 1308 //素材类型不支持
	NET_ERR_MATERIAL_NOT_EXIST                                              = 1309 //素材不存在
	NET_ERR_READ_FROM_DISK                                                  = 1310 //从硬盘读取素材文件失败
	NET_ERR_WRITE_TO_DISK                                                   = 1311 //向硬盘写素材文件失败
	NET_ERR_WRITE_DATA_BASE                                                 = 1312 //素材写数据库失败
	NET_ERR_NO_APPROVED_NOT_EXPORT                                          = 1313 //未审核内容禁止发布
	NET_ERR_MATERIAL_EXCEPTION                                              = 1314 //素材异常
	NET_ERR_NO_MISINFO                                                      = 1315 //无误报信息
	NET_ERR_LAN_NOT_SUP_DHCP_CLIENT_CONFIGURATION                           = 1316 //网桥在路由模式下，配置DHCP客户端返回错误
	NET_ERR_VIDEOWALL_OPTPORT_RESOLUTION_INCONSISTENT                       = 1317 //电视墙上各输出口分辨率不一致(主要用于设置输出分辨率为4K出现异常时报错)
	NET_ERR_VIDEOWALL_OPTPORT_RESOLUTION_INCONSISTENT_UNBIND_OPTPORT_FIRST  = 1318 //电视墙上各输出口分辨率不一致，请先解绑定输出口(主要用于绑定输出口出现异常时报错)
	NET_ERR_FOUR_K_OUTPUT_RESOLUTION_UNSUPPORT_NINE_TO_SIXTEEN_SPLIT_SCREEN = 1319 //4K输出分辨率不支持9/16分屏
	NET_ERR_SIGNAL_SOURCE_UNSUPPORT_CUSTOM_RESOLUTION                       = 1320 //信号源不支持该自定义分辨率
	NET_ERR_DVI_UNSUPPORT_FOURK_OUTPUT_RESOLUTION                           = 1321 //DVI不支持4K输出分辨率
	NET_ERR_BNC_UNSUPPORT_SOURCE_CROPPING                                   = 1322 //BNC不支持信号源裁剪

	//多屏互动错误码
	NET_ERR_MAX_SCREEN_CTRL_NUM   = 1351 //屏幕控制连接数达到上限
	NET_ERR_FILE_NOT_EXIST        = 1352 //文件不存在
	NET_ERR_THUMBNAIL_NOT_EXIST   = 1353 //缩略图不存在
	NET_ERR_DEV_OPEN_FILE_FAIL    = 1354 //设备端打开文件失败
	NET_ERR_SERVER_READ_FILE_FAIL = 1355 //设备端读取文件失败
	NET_ERR_FILE_SIZE             = 1356 //文件大小错误
	NET_ERR_FILE_NAME             = 1357 //文件名称错误，为空或不合法

	//分段错误码（1351-1400）
	NET_ERR_BROADCAST_BUSY = 1358 //设备正在广播中

	//2012-12-20抓拍机错误码（1400-1499）
	NET_DVR_ERR_LANENUM_EXCEED           = 1400 //车道数超出能力
	NET_DVR_ERR_PRAREA_EXCEED            = 1401 //牌识区域过大
	NET_DVR_ERR_LIGHT_PARAM              = 1402 //信号灯接入参数错误
	NET_DVR_ERR_LANE_LINE_INVALID        = 1403 //车道线配置错误
	NET_DVR_ERR_STOP_LINE_INVALID        = 1404 //停止线配置错误
	NET_DVR_ERR_LEFTORRIGHT_LINE_INVALID = 1405 //左/右转分界线配置错误
	NET_DVR_ERR_LANE_NO_REPEAT           = 1406 //叠加车道号重复
	NET_DVR_ERR_PRAREA_INVALID           = 1407 //牌识多边形不符合要求
	NET_DVR_ERR_LIGHT_NUM_EXCEED         = 1408 //视频检测交通灯信号灯数目超出最大值
	NET_DVR_ERR_SUBLIGHT_NUM_INVALID     = 1409 //视频检测交通灯信号灯子灯数目不合法
	NET_DVR_ERR_LIGHT_AREASIZE_INVALID   = 1410 //视频检测交通灯输入信号灯框大小不合法
	NET_DVR_ERR_LIGHT_COLOR_INVALID      = 1411 //视频检测交通灯输入信号灯颜色不合法
	NET_DVR_ERR_LIGHT_DIRECTION_INVALID  = 1412 //视频检测交通灯输入灯方向属性不合法
	NET_DVR_ERR_LACK_IOABLITY            = 1413 //IO口实际支持的能力不足

	NET_DVR_ERR_FTP_PORT                          = 1414 //FTP端口号非法（端口号重复或者异常）
	NET_DVR_ERR_FTP_CATALOGUE                     = 1415 //FTP目录名非法（启用多级目录，多级目录传值为空）
	NET_DVR_ERR_FTP_UPLOAD_TYPE                   = 1416 //FTP上传类型非法（单ftp只支持全部/双ftp只支持卡口和违章）
	NET_DVR_ERR_FLASH_PARAM_WRITE                 = 1417 //配置参数时写FLASH失败
	NET_DVR_ERR_FLASH_PARAM_READ                  = 1418 //配置参数时读FLASH失败
	NET_DVR_ERR_PICNAME_DELIMITER                 = 1419 //FTP图片命名分隔符非法
	NET_DVR_ERR_PICNAME_ITEM                      = 1420 //FTP图片命名项非法（例如 分隔符）
	NET_DVR_ERR_PLATE_RECOGNIZE_TYPE              = 1421 //牌识区域类型非法 （矩形和多边形有效性校验）
	NET_DVR_ERR_CAPTURE_TIMES                     = 1422 //抓拍次数非法 （有效值是0～5）
	NET_DVR_ERR_LOOP_DISTANCE                     = 1423 //线圈距离非法 （有效值是0～2000ms）
	NET_DVR_ERR_LOOP_INPUT_STATUS                 = 1424 //线圈输入状态非法 （有效值）
	NET_DVR_ERR_RELATE_IO_CONFLICT                = 1425 //测速组IO关联冲突
	NET_DVR_ERR_INTERVAL_TIME                     = 1426 //连拍间隔时间非法 （0～6000ms）
	NET_DVR_ERR_SIGN_SPEED                        = 1427 //标志限速值非法（大车标志限速不能大于小车标志限速 ）
	NET_DVR_ERR_PIC_FLIP                          = 1428 //图像配置翻转 （配置交互影响）
	NET_DVR_ERR_RELATE_LANE_NUMBER                = 1429 //关联车道数错误 (重复 有效值校验1～99)
	NET_DVR_ERR_TRIGGER_MODE                      = 1430 //配置抓拍机触发模式非法
	NET_DVR_ERR_DELAY_TIME                        = 1431 //触发延时时间错误(2000ms)
	NET_DVR_ERR_EXCEED_RS485_COUNT                = 1432 //超过最大485个数限制
	NET_DVR_ERR_RADAR_TYPE                        = 1433 //雷达类型错误
	NET_DVR_ERR_RADAR_ANGLE                       = 1434 //雷达角度错误
	NET_DVR_ERR_RADAR_SPEED_VALID_TIME            = 1435 //雷达有效时间错误
	NET_DVR_ERR_RADAR_LINE_CORRECT                = 1436 //雷达线性矫正参数错误
	NET_DVR_ERR_RADAR_CONST_CORRECT               = 1437 //雷达常量矫正参数错误
	NET_DVR_ERR_RECORD_PARAM                      = 1438 //录像参数无效（预录时间不超过10s）
	NET_DVR_ERR_LIGHT_WITHOUT_COLOR_AND_DIRECTION = 1439 //视频检测信号灯配置信号灯个数，但是没有勾选信号灯方向和颜色的
	NET_DVR_ERR_LIGHT_WITHOUT_DETECTION_REGION    = 1440 //视频检测信号灯配置信号灯个数，但是没有画检测区域
	NET_DVR_ERR_RECOGNIZE_PROVINCE_PARAM          = 1441 //牌识参数省份参数的合法性

	NET_DVR_ERR_SPEED_TIMEOUT        = 1442 //IO测速超时时间非法（有效值大于0）
	NET_DVR_ERR_NTP_TIMEZONE         = 1443 //ntp时区参数错误
	NET_DVR_ERR_NTP_INTERVAL_TIME    = 1444 //ntp校时间隔错误
	NET_DVR_ERR_NETWORK_CARD_NUM     = 1445 //可配置网卡数目错误
	NET_DVR_ERR_DEFAULT_ROUTE        = 1446 //默认路由错误
	NET_DVR_ERR_BONDING_WORK_MODE    = 1447 //bonding网卡工作模式错误
	NET_DVR_ERR_SLAVE_CARD           = 1448 //slave网卡错误
	NET_DVR_ERR_PRIMARY_CARD         = 1449 //Primary网卡错误
	NET_DVR_ERR_DHCP_PPOE_WORK       = 1450 //dhcp和pppoE不能同时启动
	NET_DVR_ERR_NET_INTERFACE        = 1451 //网络接口错误
	NET_DVR_ERR_MTU                  = 1452 //MTU错误
	NET_DVR_ERR_NETMASK              = 1453 //子网掩码错误
	NET_DVR_ERR_IP_INVALID           = 1454 //IP地址不合法
	NET_DVR_ERR_MULTICAST_IP_INVALID = 1455 //多播地址不合法
	NET_DVR_ERR_GATEWAY_INVALID      = 1456 //网关不合法
	NET_DVR_ERR_DNS_INVALID          = 1457 //DNS不合法
	NET_DVR_ERR_ALARMHOST_IP_INVALID = 1458 //告警主机地址不合法
	NET_DVR_ERR_IP_CONFLICT          = 1459 //IP冲突
	NET_DVR_ERR_NETWORK_SEGMENT      = 1460 //IP不支持同网段
	NET_DVR_ERR_NETPORT              = 1461 //端口错误

	NET_DVR_ERR_PPPOE_NOSUPPORT        = 1462 //PPPOE不支持
	NET_DVR_ERR_DOMAINNAME_NOSUPPORT   = 1463 //域名不支持
	NET_DVR_ERR_NO_SPEED               = 1464 //未启用测速功能
	NET_DVR_ERR_IOSTATUS_INVALID       = 1465 //IO状态错误
	NET_DVR_ERR_BURST_INTERVAL_INVALID = 1466 //连拍间隔非法
	NET_DVR_ERR_RESERVE_MODE           = 1467 //备用模式错误

	NET_DVR_ERR_LANE_NO            = 1468 //叠加车道号错误
	NET_DVR_ERR_COIL_AREA_TYPE     = 1469 //线圈区域类型错误
	NET_DVR_ERR_TRIGGER_AREA_PARAM = 1470 //触发区域参数错误
	NET_DVR_ERR_SPEED_LIMIT_PARAM  = 1471 //违章限速参数错误
	NET_DVR_ERR_LANE_PROTOCOL_TYPE = 1472 //车道关联协议类型错误

	NET_DVR_ERR_INTERVAL_TYPE               = 1473 //连拍间隔类型非法
	NET_DVR_ERR_INTERVAL_DISTANCE           = 1474 //连拍间隔距离非法
	NET_DVR_ERR_RS485_ASSOCIATE_DEVTYPE     = 1475 //RS485关联类型非法
	NET_DVR_ERR_RS485_ASSOCIATE_LANENO      = 1476 //RS485关联车道号非法
	NET_DVR_ERR_LANENO_ASSOCIATE_MULTIRS485 = 1477 //车道号关联多个RS485口
	NET_DVR_ERR_LIGHT_DETECTION_REGION      = 1478 //视频检测信号灯配置信号灯个数，但是检测区域宽或高为0

	NET_DVR_ERR_DN2D_NOSUPPORT     = 1479 //不支持抓拍帧2D降噪
	NET_DVR_ERR_IRISMODE_NOSUPPORT = 1480 //不支持的镜头类型
	NET_DVR_ERR_WB_NOSUPPORT       = 1481 //不支持的白平衡模式
	NET_DVR_ERR_IO_EFFECTIVENESS   = 1482 //IO口的有效性
	NET_DVR_ERR_LIGHTNO_MAX        = 1483 //信号灯检测器接入红/黄灯超限(16)
	NET_DVR_ERR_LIGHTNO_CONFLICT   = 1484 //信号灯检测器接入红/黄灯冲突

	NET_DVR_ERR_CANCEL_LINE        = 1485 //直行触发线
	NET_DVR_ERR_STOP_LINE          = 1486 //待行区停止线
	NET_DVR_ERR_RUSH_REDLIGHT_LINE = 1487 //闯红灯触发线
	NET_DVR_ERR_IOOUTNO_MAX        = 1488 //IO输出口编号越界

	NET_DVR_ERR_IOOUTNO_AHEADTIME_MAX             = 1489 //IO输出口提前时间超限
	NET_DVR_ERR_IOOUTNO_IOWORKTIME                = 1490 //IO输出口有效持续时间超限
	NET_DVR_ERR_IOOUTNO_FREQMULTI                 = 1491 //IO输出口脉冲模式下倍频出错
	NET_DVR_ERR_IOOUTNO_DUTYRATE                  = 1492 //IO输出口脉冲模式下占空比出错
	NET_DVR_ERR_VIDEO_WITH_EXPOSURE               = 1493 //以曝闪起效，工作方式不支持视频
	NET_DVR_ERR_PLATE_BRIGHTNESS_WITHOUT_FLASHDET = 1494 //车牌亮度自动使能闪光灯仅在车牌亮度补偿模式下起效

	NET_DVR_ERR_RECOGNIZE_TYPE_PARAM       = 1495 //识别类型非法 车牌识别参数（如大车、小车、背向、正向、车标识别等）
	NET_DVR_ERR_PALTE_RECOGNIZE_AREA_PARAM = 1496 //牌识参数非法 牌识区域配置时判断出错
	NET_DVR_ERR_PORT_CONFLICT              = 1497 //端口有冲突
	NET_DVR_ERR_LOOP_IP                    = 1498 //IP不能设置为回环地址
	NET_DVR_ERR_DRIVELINE_SENSITIVE        = 1499 //压线灵敏度出错(视频电警模式下)

	//2013-3-6VQD错误码（1500～1550）
	NET_ERR_VQD_TIME_CONFLICT = 1500 //VQD诊断时间段冲突
	NET_ERR_VQD_PLAN_NO_EXIST = 1501 //VQD诊断计划不存在
	NET_ERR_VQD_CHAN_NO_EXIST = 1502 //VQD监控点不存在
	NET_ERR_VQD_CHAN_MAX      = 1503 //VQD计划数已达上限
	NET_ERR_VQD_TASK_MAX      = 1504 //VQD任务数已达上限

	NET_SDK_GET_INPUTSTREAMCFG          = 1551 //获取信号源
	NET_SDK_AUDIO_SWITCH_CONTROL        = 1552 //子窗口音频开关控制
	NET_SDK_GET_VIDEOWALLDISPLAYNO      = 1553 //获取设备显示输出号
	NET_SDK_GET_ALLSUBSYSTEM_BASIC_INFO = 1554 //获取所有子系统基本信息
	NET_SDK_SET_ALLSUBSYSTEM_BASIC_INFO = 1555 //设置所有子系统基本信息
	NET_SDK_GET_AUDIO_INFO              = 1556 //获取所有音频信息
	NET_SDK_GET_MATRIX_STATUS_V50       = 1557 // 获取视频综合平台状态_V50
	NET_SDK_DELETE_MONITOR_INFO         = 1558 //删除Monitor信息
	NET_SDK_DELETE_CAMERA_INFO          = 1559 //删除Camaera信息

	//抓拍机错误码新增扩展(1600~1900)
	NET_DVR_ERR_EXCEED_MAX_CAPTURE_TIMES = 1600 //抓拍模式为频闪时最大抓拍张数为2张(IVT模式下)
	NET_DVR_ERR_REDAR_TYPE_CONFLICT      = 1601 //相同485口关联雷达类型冲突
	NET_DVR_ERR_LICENSE_PLATE_NULL       = 1602 //车牌号为空
	NET_DVR_ERR_WRITE_DATABASE           = 1603 //写入数据库失败
	NET_DVR_ERR_LICENSE_EFFECTIVE_TIME   = 1604 //车牌有效时间错误
	//视频电警
	NET_DVR_ERR_PRERECORDED_STARTTIME_LONG = 1605 //预录开始时间大于违法抓拍张数
	//混合卡口
	NET_DVR_ERR_TRIGGER_RULE_LINE                 = 1606 //触发规则线错误
	NET_DVR_ERR_LEFTRIGHT_TRIGGERLINE_NOTVERTICAL = 1607 //左/右触发线不垂直
	NET_DVR_ERR_FLASH_LAMP_MODE                   = 1608 //闪光灯闪烁模式错误
	NET_DVR_ERR_ILLEGAL_SNAPSHOT_NUM              = 1609 //违章抓拍张数错误
	NET_DVR_ERR_ILLEGAL_DETECTION_TYPE            = 1610 //违章检测类型错误
	NET_DVR_ERR_POSITIVEBACK_TRIGGERLINE_HIGH     = 1611 //正背向触发线高度错误
	NET_DVR_ERR_MIXEDMODE_CAPTYPE_ALLTARGETS      = 1612 //混合模式下只支持机非人抓拍类型

	NET_DVR_ERR_CARSIGNSPEED_GREATERTHAN_LIMITSPEED                                 = 1613 //小车标志限速大于限速值
	NET_DVR_ERR_BIGCARSIGNSPEED_GREATERTHAN_LIMITSPEED                              = 1614 //大车标志限速大于限速值
	NET_DVR_ERR_BIGCARSIGNSPEED_GREATERTHAN_CARSIGNSPEED                            = 1615 //大车标志限速大于小车标志限速值
	NET_DVR_ERR_BIGCARLIMITSPEED_GREATERTHAN_CARLIMITSPEED                          = 1616 //大车限速值大于小车限速值
	NET_DVR_ERR_BIGCARLOWSPEEDLIMIT_GREATERTHAN_CARLOWSPEEDLIMIT                    = 1617 //大车低速限速值大于小车低速限速值
	NET_DVR_ERR_CARLIMITSPEED_GREATERTHAN_EXCEPHIGHSPEED                            = 1618 //小车限速大于异常高速值
	NET_DVR_ERR_BIGCARLIMITSPEED_GREATERTHAN_EXCEPHIGHSPEED                         = 1619 //大车限速大于异常高速值
	NET_DVR_ERR_STOPLINE_MORETHAN_TRIGGERLINE                                       = 1620 //停止线超过直行触发线
	NET_DVR_ERR_YELLOWLIGHTTIME_INVALID                                             = 1621 // 视频检测黄灯持续时间不合法报错
	NET_DVR_ERR_TRIGGERLINE1_FOR_NOT_YIELD_TO_PEDESTRIAN_CANNOT_EXCEED_TRIGGERLINE2 = 1622 //第一条不礼让行人触发线的位置超过了第二条不礼让行人触发线
	NET_DVR_ERR_TRIGGERLINE2_FOR_NOT_YIELD_TO_PEDESTRIAN_CANNOT_EXCEED_TRIGGERLINE1 = 1623 //第二条不礼让行人触发线的位置超过了第一条不礼让行人触发线

	//门禁主机错误码
	NET_ERR_TIME_OVERLAP                            = 1900 //时间段重叠
	NET_ERR_HOLIDAY_PLAN_OVERLAP                    = 1901 //假日计划重叠
	NET_ERR_CARDNO_NOT_SORT                         = 1902 //卡号未排序
	NET_ERR_CARDNO_NOT_EXIST                        = 1903 //卡号不存在
	NET_ERR_ILLEGAL_CARDNO                          = 1904 //卡号错误
	NET_ERR_ZONE_ALARM                              = 1905 //防区处于布防状态(参数修改不允许)
	NET_ERR_ZONE_OPERATION_NOT_SUPPORT              = 1906 //防区不支持该操作
	NET_ERR_INTERLOCK_ANTI_CONFLICT                 = 1907 //多门互锁和反潜回同时配置错误
	NET_ERR_DEVICE_CARD_FULL                        = 1908 //卡已满（卡达到10W后返回）
	NET_ERR_HOLIDAY_GROUP_DOWNLOAD                  = 1909 //假日组下载失败
	NET_ERR_LOCAL_CONTROL_OFF                       = 1910 //就地控制器离线
	NET_ERR_LOCAL_CONTROL_DISADD                    = 1911 //就地控制器未添加
	NET_ERR_LOCAL_CONTROL_HASADD                    = 1912 //就地控制器已添加
	NET_ERR_LOCAL_CONTROL_DOORNO_CONFLICT           = 1913 //与已添加的就地控制器门编号冲突
	NET_ERR_LOCAL_CONTROL_COMMUNICATION_FAIL        = 1914 //就地控制器通信失败
	NET_ERR_OPERAND_INEXISTENCE                     = 1915 //操作对象不存在（对门、报警输出、报警输入相关操作，当对象未添加时返回）
	NET_ERR_LOCAL_CONTROL_OVER_LIMIT                = 1916 //就地控制器超出设备最大能力（主控对就地数量有限制）
	NET_ERR_DOOR_OVER_LIMIT                         = 1917 //门超出设备最大能力
	NET_ERR_ALARM_OVER_LIMIT                        = 1918 //报警输入输出超出设备最大能力
	NET_ERR_LOCAL_CONTROL_ADDRESS_INCONFORMITY_TYPE = 1919 //就地控制器地址与类型不符
	NET_ERR_NOT_SUPPORT_ONE_MORE_CARD               = 1920 //不支持一人多卡
	NET_ERR_DELETE_NO_EXISTENCE_FACE                = 1921 //删除的人脸不存在
	NET_ERR_DOOR_SPECIAL_PASSWORD_REPEAT            = 1922 //与设备门特殊密码重复
	NET_ERR_AUTH_CODE_REPEAT                        = 1923 //与设备认证码重复
	NET_ERR_DEPLOY_EXCEED_MAX                       = 1924 //布防超过最大连接数
	NET_ERR_NOT_SUPPORT_DEL_FP_BY_ID                = 1925 //读卡器不支持按手指ID删除指纹
	NET_ERR_TIME_RANGE                              = 1926 //有效期参数配置范围有误
	NET_ERR_CAPTURE_TIMEOUT                         = 1927 //采集超时
	NET_ERR_LOW_SCORE                               = 1928 //采集质量低
	NET_ERR_OFFLINE_CAPTURING                       = 1929 //离线采集中，无法响应

	//可视对讲错误码
	NET_DVR_ERR_OUTDOOR_COMMUNICATION       = 1950 //与门口机通信异常
	NET_DVR_ERR_ROOMNO_UNDEFINED            = 1951 //未设置房间号
	NET_DVR_ERR_NO_CALLING                  = 1952 //无呼叫
	NET_DVR_ERR_RINGING                     = 1953 //响铃
	NET_DVR_ERR_IS_CALLING_NOW              = 1954 //正在通话
	NET_DVR_ERR_LOCK_PASSWORD_WRONG         = 1955 //智能锁密码错误
	NET_DVR_ERR_CONTROL_LOCK_FAILURE        = 1956 //开关锁失败
	NET_DVR_ERR_CONTROL_LOCK_OVERTIME       = 1957 //开关锁超时
	NET_DVR_ERR_LOCK_DEVICE_BUSY            = 1958 //智能锁设备繁忙
	NET_DVR_ERR_UNOPEN_REMOTE_LOCK_FUNCTION = 1959 //远程开锁功能未打开

	//后端错误码 （2100 - 3000）
	NET_DVR_ERR_FILE_NOT_COMPLETE   = 2100 //下载的文件不完整
	NET_DVR_ERR_IPC_EXIST           = 2101 //该IPC已经存在
	NET_DVR_ERR_ADD_IPC             = 2102 //该通道已添加IPC
	NET_DVR_ERR_OUT_OF_RES          = 2103 //网络带宽能力不足
	NET_DVR_ERR_CONFLICT_TO_LOCALIP = 2104 //IPC的ip地址跟DVR的ip地址冲突
	NET_DVR_ERR_IP_SET              = 2105 //非法ip地址
	NET_DVR_ERR_PORT_SET            = 2106 //非法的端口号

	NET_ERR_WAN_NOTSUPPORT                     = 2107 //不在同一个局域网，无法设置安全问题或导出GUID文件
	NET_ERR_MUTEX_FUNCTION                     = 2108 //功能互斥
	NET_ERR_QUESTION_CONFIGNUM                 = 2109 //安全问题配置数量错误
	NET_ERR_FACECHAN_NORESOURCE                = 2110 //人脸智能通道资源已用完
	NET_ERR_DATA_CALLBACK                      = 2111 //正在数据回迁
	NET_ERR_ATM_VCA_CHAN_IS_RELATED            = 2112 //ATM智能通道已被关联
	NET_ERR_ATM_VCA_CHAN_IS_OVERLAPED          = 2113 //ATM智能通道已被叠加
	NET_ERR_FACE_CHAN_UNOVERLAP_EACH_OTHER     = 2114 //人脸通道不能互相叠加
	NET_ERR_ACHIEVE_MAX_CHANNLE_LIMIT          = 2115 //达到最大路数限制
	NET_DVR_SMD_ENCODING_NORESOURSE            = 2116 //SMD编码资源不足
	NET_DVR_SMD_DECODING_NORESOURSE            = 2117 //SMD解码资源不足
	NET_DVR_FACELIB_DATA_PROCESSING            = 2118 //人脸库数据正在处理
	NET_DVR_ERR_LARGE_TIME_DIFFRENCE           = 2119 //设备和服务器之间的时间差异太大
	NET_DVR_NO_SUPPORT_WITH_PLAYBACK           = 2120 //已开启回放，不支持本功能
	NET_DVR_CHANNEL_NO_SUPPORT_WITH_SMD        = 2121 //通道已开启SMD，不支持本功能
	NET_DVR_CHANNEL_NO_SUPPORT_WITH_FD         = 2122 //通道已开启人脸抓拍，不支持本功能
	NET_DVR_ILLEGAL_PHONE_NUMBER               = 2123 //非法的电话号码
	NET_DVR_ILLEGAL_CERITIFICATE_NUMBER        = 2124 //非法的证件号码
	NET_DVR_ERR_CHANNEL_RESOLUTION_NO_SUPPORT  = 2125 //通道分辨率不支持
	NET_DVR_ERR_CHANNEL_COMPRESSION_NO_SUPPORT = 2126 //通道编码格式不支持

	NET_DVR_ERR_CLUSTER_DEVICE_TOO_LESS              = 2127 //设备数少，不允许删除
	NET_DVR_ERR_CLUSTER_DEL_DEVICE_CM_PLAYLOAD       = 2128 //该设备是集群主机，不允许删除
	NET_DVR_ERR_CLUSTER_DEVNUM_OVER_UPPER_LIMIT      = 2129 //设备数达到上限，不允许增加
	NET_DVR_ERR_CLUSTER_DEVICE_TYPE_INCONFORMITY     = 2130 //设备类型不一致
	NET_DVR_ERR_CLUSTER_DEVICE_VERSION_INCONFORMITY  = 2131 //设备版本不一致
	NET_DVR_ERR_CLUSTER_IP_CONFLICT                  = 2132 //集群系统IP地址冲突：ipv4地址冲突、ipv6地址冲突
	NET_DVR_ERR_CLUSTER_IP_INVALID                   = 2133 //集群系统IP地址无效：ipv4非法、ipv6非法
	NET_DVR_ERR_CLUSTER_PORT_CONFLICT                = 2134 //集群系统端口冲突
	NET_DVR_ERR_CLUSTER_PORT_INVALID                 = 2135 //集群系统端口非法
	NET_DVR_ERR_CLUSTER_USERNAEM_OR_PASSWORD_INVALID = 2136 //组网用户名或密码非法
	NET_DVR_ERR_CLUSTER_DEVICE_ALREADY_EXIST         = 2137 //存在相同设备
	NET_DVR_ERR_CLUSTER_DEVICE_NOT_EXIST             = 2138 //设备不存在(组网时下发的cs列表中的设备信息在任何一台cs上都找不到，删除的时候下发的id不对)
	NET_DVR_ERR_CLUSTER_NON_CLUSTER_MODE             = 2139 //设备处于非集群模式
	NET_DVR_ERR_CLUSTER_IP_NOT_SAME_LAN              = 2140 //IP地址不在同一局域网，不同区域网不允许组网/扩容

	NET_DVR_ERR_CAPTURE_PACKAGE_FAILED                = 2141 //抓包失败
	NET_DVR_ERR_CAPTURE_PACKAGE_PROCESSING            = 2142 //正在抓包
	NET_DVR_ERR_SAFETY_HELMET_NO_RESOURCE             = 2143 //安全帽检测资源不足
	NET_DVR_NO_SUPPORT_WITH_ABSTRACT                  = 2144 //已开启视频摘要，不支持本功能
	NET_DVR_ERR_TAPE_LIB_NEED_STOP_ARCHIVE            = 2145 //磁带库归档需要停止
	NET_DVR_INSUFFICIENT_DEEP_LEARNING_RESOURCES      = 2146 //深度学习资源超限
	NET_DVR_ERR_IDENTITY_KEY                          = 2147 //交互口令错误
	NET_DVR_MISSING_IDENTITY_KEY                      = 2148 //交互口令缺失
	NET_DVR_NO_SUPPORT_WITH_PERSON_DENSITY_DETECT     = 2149 //已开启人员密度检测，不支持本功能
	NET_DVR_IPC_RESOLUTION_OVERFLOW                   = 2150 //IPC分辨率超限
	NET_DVR_IPC_BITRATE_OVERFLOW                      = 2151 //IPC码率超限
	NET_DVR_ERR_INVALID_TASKID                        = 2152 //无效的taskID
	NET_DVR_PANEL_MODE_NOT_CONFIG                     = 2153 //没有配置面板路智能
	NET_DVR_NO_HUMAN_ENGINES_RESOURCE                 = 2154 //人体引擎资源不足
	NET_DVR_ERR_TASK_NUMBER_OVERFLOW                  = 2155 //任务数据超过上限
	NET_DVR_ERR_COLLISION_TIME_OVERFLOW               = 2156 //碰撞时间超过上限
	NET_DVR_ERR_CAPTURE_PACKAGE_NO_USB                = 2157 //未识别到U盘，请插入U盘或重新插入
	NET_DVR_ERR_NO_SET_SECURITY_EMAIL                 = 2158 //未设置安全邮箱
	NET_DVR_ERR_EVENT_NOTSUPPORT                      = 2159 //订阅事件不支持
	NET_DVR_ERR_PASSWORD_FORMAT                       = 2160 //密码格式不对
	NET_DVR_ACCESS_FRONT_DEVICE_PARAM_FAILURE         = 2161 //获取前端设备参数失败
	NET_DVR_ACCESS_FRONT_DEVICE_STREAM_FAILURE        = 2162 //对前端设备取流失败
	NET_DVR_ERR_USERNAME_FORMAT                       = 2163 //用户名格式不对
	NET_DVR_ERR_UNOPENED_HIGH_RESOLUTION_MODE         = 2164 //超高分辨率模式未开启
	NET_DVR_ERR_TOO_SMALL_QUATO                       = 2165 //配额设置太小
	NET_DVR_ERR_EMAIL_FORMAT                          = 2166 //邮箱格式不对
	NET_DVR_ERR_SECURITY_CODE_FORMAT                  = 2167 //安全码格式不对
	NET_DVR_PD_SPACE_TOO_SMALL                        = 2168 //阵列硬盘容量太小
	NET_DVR_PD_NUM_TOO_BIG                            = 2169 //阵列硬盘总数超过总盘数的二分之一
	NET_DVR_ERR_USB_IS_FULL                           = 2170 //U盘已满
	NET_DVR_EXCEED_MAX_SMD_TYPE                       = 2171 //达到最大SMD事件种类上限
	NET_DVR_CHANNEL_NO_SUPPORT_WITH_BEHAVIOR          = 2172 //通道已开启行为分析，不支持本功能
	NET_DVR_NO_BEHAVIOR_ENGINES_RESOURCE              = 2173 //行为分析资源不足
	NET_DVR_NO_RETENTION_ENGINES_RESOURCE             = 2174 //人员滞留检测资源不足
	NET_DVR_NO_LEAVE_POSITION_ENGINES_RESOURCE        = 2175 //离岗检测资源不足
	NET_DVR_NO_PEOPLE_NUM_CHANGE_ENGINES_RESOURCE     = 2176 //人数异常资源不足
	NET_DVR_PANEL_MODE_NUM_OVER_LIMIT                 = 2177 //超过面板路最大路数
	NET_DVR_SURROUND_MODE_NUM_OVER_LIMIT              = 2178 //超过环境路最大路数
	NET_DVR_FACE_MODE_NUM_OVER_LIMIT                  = 2179 //超过人脸路最大路数
	NET_DVR_SAFETYCABIN_MODE_NUM_OVER_LIMIT           = 2180 //超过防护舱路最大路数
	NET_DVR_DETECT_REGION_RANGE_INVALID               = 2181 //检测区域范围非法
	NET_DVR_CHANNEL_CAPTURE_PICTURE_FAILURE           = 2182 //通道抓图失败
	NET_DVR_VCACHAN_IS_NORESOURCE                     = 2183 //智能通道资源用完
	NET_DVR_IPC_NUM_REACHES_LIMIT                     = 2184 // Ipc通道数目达到上限
	NET_DVR_IOT_NUM_REACHES_LIMIT                     = 2185 // IOT通道数目达到上限
	NET_DVR_IOT_CHANNEL_DEVICE_EXIST                  = 2186 //当前IOT通道已经添加设备
	NET_DVR_IOT_CHANNEL_DEVICE_NOT_EXIST              = 2187 //当前IOT通道不存在设备
	NET_DVR_INVALID_IOT_PROTOCOL_TYPE                 = 2188 //非法的IOT协议类型
	NET_DVR_INVALID_EZVIZ_SECRET_KEY                  = 2189 //非法的萤石注册验证码
	NET_DVR_DUPLICATE_IOT_DEVICE                      = 2190 //重复的IOT设备
	NET_DVR_SADP_MODIFY_FALIURE                       = 2191 // SADP修改失败
	NET_DVR_IPC_NETWORK_ABNORMAL                      = 2192 // IPC网络异常
	NET_DVR_IPC_PASSWORD_ERROR                        = 2193 // IPC用户名密码错误
	NET_DVR_ERROR_IPC_TYPE                            = 2194 //IPC类型不对
	NET_DVR_ERROR_IPC_LIST_NOT_EMPTY                  = 2195 //已添加IPC列表不为空，不支持一键配置
	NET_DVR_ERROR_IPC_LIST_NOT_MATCH_PAIRING          = 2196 //IPC列表和配单不匹配
	NET_DVR_ERROR_IPC_BAD_LANGUAGE                    = 2197 //IPC语言和设备不匹配
	NET_DVR_ERROR_IPC_IS_LOCKING                      = 2198 //IPC已被锁
	NET_DVR_ERROR_IPC_NOT_ACTIVATED                   = 2199 //IPC未激活
	NET_DVR_FIELD_CODING_NOT_SUPPORT                  = 2200 //场编码不支持
	NET_DVR_ERROR_H323_NOT_SUPPORT_H265               = 2201 //H323视频会议就不支持H265码流
	NET_DVR_ERROR_EXPOSURE_TIME_TOO_BIG_IN_MODE_P     = 2202 //P制式下，曝光时间过大
	NET_DVR_ERROR_EXPOSURE_TIME_TOO_BIG_IN_MODE_N     = 2203 //N制式下，曝光时间过大
	NET_DVR_ERROR_PING_PROCESSING                     = 2204 //正在PING
	NET_DVR_ERROR_PING_NOT_START                      = 2205 //Ping功能未开始
	NET_DVR_ERROR_NEED_DOUBLE_VERIFICATION            = 2206 //需要二次认证
	NET_DVR_NO_DOUBLE_VERIFICATION_USER               = 2207 //无二次认证用户
	NET_DVR_CHANNEL_OFFLINE                           = 2208 //通道离线
	NET_DVR_TIMESPAN_NUM_OVER_LIMIT                   = 2209 //时间段超出支持最大数目
	NET_DVR_CHANNEL_NUM_OVER_LIMIT                    = 2210 //通道数目超出支持最大数目
	NET_DVR_NO_SEARCH_ID_RESOURCE                     = 2211 //分页查询的searchID资源不足
	NET_DVR_ERROR_ONEKEY_EXPORT                       = 2212 //正在进行导出操作，请稍后再试
	NET_DVR_NO_CITY_MANAGEMENT_ENGINES_RESOURCE       = 2213 //城管算法引擎资源不足
	NET_DVR_NO_SITUATION_ANALYSIS_ENGINES_RESOURCE    = 2214 //态势分析引擎资源不足
	NET_DVR_INTELLIGENT_ANALYSIS_IPC_CANNT_DELETE     = 2215 //正在进行智能分析的IPC无法删除
	NET_DVR_NOSUPPORT_RESET_PASSWORD                  = 2216 //NVR不支持对IPC重置密码
	NET_DVR_ERROR_IPC_NEED_ON_LAN                     = 2217 // IPC需要在局域网内
	NET_DVR_CHANNEL_NO_SUPPORT_WITH_SAFETY_HELMET     = 2218 //通道已开启安全帽检测，不支持本功能
	NET_DVR_ERROR_GET_RESETPASSWORDTYPE_IS_ABNORMAL   = 2219 // IPC重置密码时，获取IPC的重置密码类型异常
	NET_DVR_ERROR_IPC_NOSUPPORT_RESET_PASSWORD        = 2220 //  IPC不支持重置密码
	NET_DVR_ERROR_IP_IS_NOT_ONLY_ONE                  = 2221 // IPC的IP不唯一，有重复
	NET_DVR_NO_SUPPORT_WITH_SMD_OR_SCD                = 2222 //已开启SMD/SCD，不支持本功能（SCD为场景变更）
	NET_DVR_NO_SUPPORT_WITH_FD                        = 2223 //已开启人脸抓拍，不支持本功能
	NET_DVR_NO_FD_ENGINES_RESOURCE                    = 2224 //人脸抓拍资源不足
	NET_DVR_ERROR_ONEKEY_REMOVE                       = 2225 //正在进行删除操作，请稍后再试
	NET_DVR_FACE_PIP_BACKGROUND_CHANNEL_OVERFLOW      = 2226 //人脸画中画背景通道超限
	NET_DVR_MICIN_CHANNEL_OCCUPIED                    = 2227 //micin通道被占用
	NET_DVR_IPC_CHANNEL_IS_IN_PIP                     = 2228 //操作失败，该通道已关联到审讯通道，请先取消画中画配置关联
	NET_DVR_CHANNEL_NO_SUPPORT_WITH_FACE_CONTRAST     = 2229 //通道已开启人脸比对，不支持本功能
	NET_DVR_INVALID_RECHARGE_CARD                     = 2230 //无效的充值卡
	NET_DVR_CLOUD_PLATFORM_SERVER_EXCEPTION           = 2231 //云平台服务器异常
	NET_DVR_OPERATION_FAILURE_WITHOUT_LOGIN           = 2232 //未登录操作失败
	NET_DVR_INVALID_ASSOCIATED_SERIAL_NUMBER          = 2233 //关联序列号非法
	NET_DVR_CLOUD_PLATFORM_ACCOUNT_NOT_EXIST          = 2234 //云平台帐号不存在
	NET_DVR_DEVICE_SERIAL_NUMBER_REGISTERED           = 2235 //设备序列号已注册
	NET_DVR_CONFERENCE_ROOM_NOT_EXIST                 = 2236 //会议室不存在
	NET_DVR_NEED_DISABLED_ANALOG_CHANNEL              = 2237 //需禁用模拟通道
	NET_DVR_STUDENT_ROLL_CALL_FAILURE                 = 2238 //学生点名失败
	NET_DVR_SUB_DEVICE_NOT_ENABLE_INDIVIDUAL_BEHAVIOR = 2239 //子设备未启用个体行为模式
	NET_DVR_SUB_DEVICE_CHANNEL_CONTROL_FAILED         = 2240 //子设备通道控制失败
	NET_DVR_DEVICE_NOT_IN_CONFERENCE                  = 2241 //设备不在会议中
	NET_DVR_ALREADY_EXIST_CONFERENCE                  = 2242 //当前已经存在会议
	NET_DVR_NO_SUPPORT_WITH_VIDEO_CONFERENCE          = 2243 //当前正在视频会议中，不支持本功能
	NET_DVR_START_INTERACTION_FAILURE                 = 2244 //互动开始失败
	NET_DVR_ASK_QUESTION_STARTED                      = 2245 //已开始提问
	NET_DVR_ASK_QUESTION_CLOSED                       = 2246 //已结束提问
	NET_DVR_UNABLE_OPERATE_BY_HOST                    = 2247 //已被主持人禁用，无法操作
	NET_DVR_REPEATED_ASK_QUESTION                     = 2248 //重复提问
	NET_DVR_SWITCH_TIMEDIFF_LESS_LIMIT                = 2249 // 开关机时间差小于限制值(10分钟)
	NET_DVR_CHANNEL_DEVICE_EXIST                      = 2250 //当前通道已经添加设备
	NET_DVR_CHANNEL_DEVICE_NOT_EXIST                  = 2251 //当前通道不存在设备
	NET_DVR_ERROR_ADJUSTING_RESOLUTION                = 2252 //先关闭摄像机的裁剪，再调整分辨率
	NET_DVR_SSD_FILE_SYSTEM_IS_UPGRADING              = 2253 //SSD文件系统正在升级
	NET_DVR_SSD_FILE_SYSTEM_IS_FORMAT                 = 2254 //SSD正在格式化
	NET_DVR_CHANNEL_IS_CONNECTING                     = 2255 //当前通道正在连接

	NET_DVR_CHANNEL_STREAM_TYPE_NOT_SUPPORT = 2257 //当前通道码流类型不支持
	NET_DVR_CHANNEL_USERNAME_NOT_EXIST      = 2258 //当前通道用户名不存在
	NET_DVR_CHANNEL_ACCESS_PARAM_FAILURE    = 2259 //当前通道获取参数失败
	NET_DVR_CHANNEL_GET_STREAM_FAILURE      = 2260 //当前通道取流失败
	NET_DVR_CHANNEL_RISK_PASSWORD           = 2261 //当前通道密码为风险密码
	NET_DVR_NO_SUPPORT_DELETE_STRANGER_LIB  = 2262 //不支持删除陌生人库
	NET_DVR_NO_SUPPORT_CREATE_STRANGER_LIB  = 2263 //不支持创建陌生人库
	NET_DVR_NETWORK_PORT_CONFLICT           = 2264 //网络配置端口冲突
	NET_DVR_TRANSCODE_NO_RESOURCES          = 2265 //转码资源不足
	NET_DVR_SSD_FILE_SYSTEM_ERROR           = 2266 //SSD文件系统错误
	NET_DVR_INSUFFICIENT_SSD__FOR_FPD       = 2267 //用于人员频次业务的SSD容量不够
	NET_DVR_ASSOCIATED_FACELIB_OVER_LIMIT   = 2268 //关联人脸库已达上限
	NET_DVR_NEED_DELETE_DIGITAL_CHANNEL     = 2269 //需删除数字通道
	//热成像产线相关错误码（3001 - 3500）
	NET_DVR_ERR_NOTSUPPORT_DEICING             = 3001 //设备当前状态不支持除冰功能（只在POE+、AC24V、DC12V供电下支持除冰功能）
	NET_DVR_ERR_THERMENABLE_CLOSE              = 3002 //测温功能总使能未开启。(即NET_DVR_THERMOMETRY_BASICPARAM使能未开启)
	NET_DVR_ERR_NOTMEET_DEICING                = 3003 //当前空腔温度不满足手动除冰开启条件（需空腔温度小于30度才可开启）
	NET_DVR_ERR_PANORAMIC_LIMIT_OPERATED       = 3004 //全景地图和限位不可同时操作
	NET_DVR_ERR_SMARTH264_ROI_OPERATED         = 3005 //SmartH264和ROI不可同时操作
	NET_DVR_ERR_RULENUM_LIMIT                  = 3006 //规则数上限
	NET_DVR_ERR_LASER_DEICING_OPERATED         = 3007 //激光以及除冰功能不可同时操作
	NET_DVR_ERR_OFFDIGITALZOOM_OR_MINZOOMLIMIT = 3008 //请先关闭数据变倍功能或变倍限制设置为最小值。当执行烟火检测、行为分析、船只检测、坏点矫正、测温、烟火屏蔽功能时，若没有关闭数据变倍或者变倍限制没有设置为最小值时，将会提示该错误码。
	NET_DVR_ERR_FIREWAITING                    = 3009 //设备云台正在火点等待中
	NET_DVR_SYNCHRONIZEFOV_ERROR               = 3010 //同步视场角错误
	NET_DVR_CERTIFICATE_VALIDATION_ERROR       = 3011 //证书验证不通过
	NET_DVR_CERTIFICATES_NUM_EXCEED_ERROR      = 3012 //证书个数超过上限
	NET_DVR_RULE_SHIELDMASK_CONFLICT_ERROR     = 3013 //规则区域与屏蔽区域冲突

	//前端产品线错误码（3501-4000）
	NET_DVR_ERR_NO_SAFETY_HELMET_REGION = 3501 //未配置安全帽检测区域
	NET_DVR_ERR_UNCLOSED_SAFETY_HELMET  = 3502 //未关闭安全帽检测使能
	NET_DVR_ERR_MUX_RECV_STATE          = 3503 //多路复用接收状态异常
	NET_DVR_UPLOAD_HBDLIBID_ERROR       = 3504 // 上传人体库ID（HBDID or customHBDID）错误

	NET_ERR_NPQ_BASE_INDEX   = 8000                         //NPQ库错误码
	NET_ERR_NPQ_PARAM        = (NET_ERR_NPQ_BASE_INDEX + 1) //NPQ库参数有误
	NET_ERR_NPQ_SYSTEM       = (NET_ERR_NPQ_BASE_INDEX + 2) //NPQ库操作系统调用错误(包括资源申请失败或内部错误等）
	NET_ERR_NPQ_GENRAL       = (NET_ERR_NPQ_BASE_INDEX + 3) //NPQ库内部通用错误
	NET_ERR_NPQ_PRECONDITION = (NET_ERR_NPQ_BASE_INDEX + 4) //NPQ库调用顺序错误
	NET_ERR_NPQ_NOTSUPPORT   = (NET_ERR_NPQ_BASE_INDEX + 5) //NPQ库功能不支持

	NET_ERR_NPQ_NOTCALLBACK = (NET_ERR_NPQ_BASE_INDEX + 100) //数据没有回调上来
	NET_ERR_NPQ_LOADLIB     = (NET_ERR_NPQ_BASE_INDEX + 101) //NPQ库加载失败
	NET_ERR_NPQ_STEAM_CLOSE = (NET_ERR_NPQ_BASE_INDEX + 104) //本路码流NPQ功能未开启
	NET_ERR_NPQ_MAX_LINK    = (NET_ERR_NPQ_BASE_INDEX + 110) //NPQ取流路数达到上限
	NET_ERR_NPQ_STREAM_CFG  = (NET_ERR_NPQ_BASE_INDEX + 111) //编码参数存在冲突配置

	NET_EZVIZ_P2P_BASE_INDEX                     = 8300 //萤石P2P组件错误码
	NET_DVR_EZVIZ_P2P_REGISTER_ERROR             = (NET_EZVIZ_P2P_BASE_INDEX + 1)
	NET_DVR_EZVIZ_P2P_LOGIN_2C_ERROR             = (NET_EZVIZ_P2P_BASE_INDEX + 2)
	NET_DVR_EZVIZ_P2P_LOGIN_2B_ERROR             = (NET_EZVIZ_P2P_BASE_INDEX + 3)
	NET_DVR_EZVIZ_P2P_BUILDLINK_ERROR            = (NET_EZVIZ_P2P_BASE_INDEX + 4)
	NET_DVR_EZVIZ_P2P_PORTMAPPING_ERROR          = (NET_EZVIZ_P2P_BASE_INDEX + 5)
	NET_DVR_EZVIZ_P2P_COULDNT_RESOLVE_HOST       = (NET_EZVIZ_P2P_BASE_INDEX + 6)  //P2PCLOUD_ER_COULDNT_RESOLVE_HOST    1006
	NET_DVR_EZVIZ_P2P_COULDNT_CONNECT            = (NET_EZVIZ_P2P_BASE_INDEX + 7)  //P2PCLOUD_ER_COULDNT_CONNECT 1007
	NET_DVR_EZVIZ_P2P_OPERATION_TIMEOUT          = (NET_EZVIZ_P2P_BASE_INDEX + 8)  //P2PCLOUD_ER_OPERATION_TIMEOUT   1028
	NET_DVR_EZVIZ_P2P_NOT_INITIALIZED            = (NET_EZVIZ_P2P_BASE_INDEX + 9)  //P2PCLOUD_ER_NOT_INITIALIZED 2001
	NET_DVR_EZVIZ_P2P_INVALID_ARG                = (NET_EZVIZ_P2P_BASE_INDEX + 10) //P2PCLOUD_ER_INVALID_ARG 2002
	NET_DVR_EZVIZ_P2P_EXCEED_MAX_SERVICE         = (NET_EZVIZ_P2P_BASE_INDEX + 11) //P2PCLOUD_ER_EXCEED_MAX_SERVICE  2003
	NET_DVR_EZVIZ_P2P_TOKEN_NOT_EXIST            = (NET_EZVIZ_P2P_BASE_INDEX + 12) //P2PCLOUD_ER_TOKEN_NOT_EXIST 2004
	NET_DVR_EZVIZ_P2P_DISCONNECTED               = (NET_EZVIZ_P2P_BASE_INDEX + 13) //P2PCLOUD_ER_DISCONNECTED    2005
	NET_DVR_EZVIZ_P2P_RELAY_ADDR_NOT_EXIST       = (NET_EZVIZ_P2P_BASE_INDEX + 14) //P2PCLOUD_ER_RELAY_ADDR_NOT_EXIST    2006
	NET_DVR_EZVIZ_P2P_DEV_NOT_ONLINE             = (NET_EZVIZ_P2P_BASE_INDEX + 15) //P2PCLOUD_ER_DEV_NOT_ONLINE  3121
	NET_DVR_EZVIZ_P2P_DEV_CONNECT_EXCEED         = (NET_EZVIZ_P2P_BASE_INDEX + 16) //P2PCLOUD_ER_DEV_CONNECT_EXCEED  3123
	NET_DVR_EZVIZ_P2P_DEV_CONNECT_FAILED         = (NET_EZVIZ_P2P_BASE_INDEX + 17) //P2PCLOUD_ER_DEV_CONNECT_FAILED  3209
	NET_DVR_EZVIZ_P2P_DEV_RECV_TIMEOUT           = (NET_EZVIZ_P2P_BASE_INDEX + 18) //P2PCLOUD_ER_DEV_RECV_TIMEOUT    3213
	NET_DVR_EZVIZ_P2P_USER_FORCE_STOP            = (NET_EZVIZ_P2P_BASE_INDEX + 19) //P2PCLOUD_ER_USER_FORCE_STOP 3216
	NET_DVR_EZVIZ_P2P_NO_PERMISSION              = (NET_EZVIZ_P2P_BASE_INDEX + 20) //P2PCLOUD_ER_NO_PERMISSION   3255
	NET_DVR_EZVIZ_P2P_DEV_PU_NOT_FOUND           = (NET_EZVIZ_P2P_BASE_INDEX + 21) //P2PCLOUD_ER_DEV_PU_NOT_FOUND    3297
	NET_DVR_EZVIZ_P2P_DEV_CONN_NOLONGER_AVAIL    = (NET_EZVIZ_P2P_BASE_INDEX + 22) //P2PCLOUD_ER_DEV_CONN_NOLONGER_AVAIL 3351
	NET_DVR_EZVIZ_P2P_DEV_NOT_LISTENING          = (NET_EZVIZ_P2P_BASE_INDEX + 23) //P2PCLOUD_ER_DEV_NOT_LISTENING   3610
	NET_DVR_EZVIZ_P2P_DEV_TUNNEL_SOCKET_LIMITED  = (NET_EZVIZ_P2P_BASE_INDEX + 24) //P2PCLOUD_ER_DEV_TUNNEL_SOCKET_LIMITED   3612
	NET_DVR_EZVIZ_P2P_FAIL_CREATE_THREAD         = (NET_EZVIZ_P2P_BASE_INDEX + 25) //TUNNEL_ER_FAIL_CREATE_THREAD    4001
	NET_DVR_EZVIZ_P2P_FAIL_ALLOC_BUFFERS         = (NET_EZVIZ_P2P_BASE_INDEX + 26) //P2PCLOUD_ER_FAIL_ALLOC_BUFFERS  4002
	NET_DVR_EZVIZ_P2P_FAIL_CREATE_SOCKET         = (NET_EZVIZ_P2P_BASE_INDEX + 27) //P2PCLOUD_ER_FAIL_CREATE_SOCKET  4003
	NET_DVR_EZVIZ_P2P_BIND_LOCAL_SERVICE         = (NET_EZVIZ_P2P_BASE_INDEX + 28) //P2PCLOUD_ER_BIND_LOCAL_SERVICE  4004
	NET_DVR_EZVIZ_P2P_LISTEN_LOCAL_SERVICE       = (NET_EZVIZ_P2P_BASE_INDEX + 29) //P2PCLOUD_ER_LISTEN_LOCAL_SERVICE    4005
	NET_DVR_EZVIZ_P2P_SVR_RSP_BAD                = (NET_EZVIZ_P2P_BASE_INDEX + 30) //P2PCLOUD_ER_SVR_RSP_BAD 5001
	NET_DVR_EZVIZ_P2P_SVR_RSP_INVALID            = (NET_EZVIZ_P2P_BASE_INDEX + 31) //P2PCLOUD_ER_SVR_RSP_INVALID 5002
	NET_DVR_EZVIZ_P2P_SVR_LOGIN_FAILED           = (NET_EZVIZ_P2P_BASE_INDEX + 32) //P2PCLOUD_ER_SVR_LOGIN_FAILED    5003
	NET_DVR_EZVIZ_P2P_SVR_TOKEN_EXPIRED          = (NET_EZVIZ_P2P_BASE_INDEX + 33) //P2PCLOUD_ER_SVR_TOKEN_EXPIRED   5004
	NET_DVR_EZVIZ_P2P_SVR_DEV_NOT_BELONG_TO_USER = (NET_EZVIZ_P2P_BASE_INDEX + 34) //P2PCLOUD_ER_SVR_DEV_NOT_BELONG_TO_USER  5005

	//传显错误码 8501~9500
	NET_ERR_UPGRADE_PROG_ERR        = 8501 //程序执行出错
	NET_ERR_UPGRADE_NO_DEVICE       = 8502 //没有设备(指LED控制器没有接接收卡)
	NET_ERR_UPGRADE_NO_FILE         = 8503 //没有找到升级文件
	NET_ERR_UPGRADE_DATA_ERROR      = 8504 //升级文件数据不兼容
	NET_ERR_UPGRADE_LINK_SERVER_ERR = 8505 //与服务器连接失败
	NET_ERR_UPGRADE_OEMCODE_NOMATCH = 8506 //oemCode不匹配
	NET_ERR_UPGRADE_FLASH_NOENOUGH  = 8507 //flash不足
	NET_ERR_UPGRADE_RAM_NOENOUGH    = 8508 //RAM不足
	NET_ERR_UPGRADE_DSPRAM_NOENOUGH = 8509 //DSP RAM不足
	NET_ERR_NOT_SUPPORT_CHECK       = 8510 //该屏幕型号不支持校正
	NET_ERR_LED_DEVICE_BUSY_CHECK   = 8511 //LED设备忙（正在校正）
	NET_ERR_DEVICE_MEM_NOT_ENOUGH   = 8512 //设备内存不足
	NET_ERR_CHECK_PARAM             = 8513 //校正参数错误
	NET_ERR_RESOLUTION_OVER_LIMIT   = 8514 //输入分辨率超过限制
	NET_ERR_NO_CUSTOM_BASE          = 8515 //无自定义底图
	NET_ERR_PRIORITY_LOWER          = 8516 //优先级低于当前模式
	NET_ERR_SEND_MESSAGE_EXCEPT     = 8517 //消息发送异常
	NET_ERR_SENDCARD_UPGRADING      = 8518 //发送卡升级中
	NET_ERR_NO_WIRELESS_NETCARD     = 8519 //未插入无线网卡
	NET_ERR_LOAD_FS_FAIL            = 8520 //从屏幕加载失败
	NET_ERR_FLASH_UNSTORAGE_RECCARD = 8521 //Flash中未存储接收卡参数
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
	// 阵列错误码
	case NET_DVR_NAME_NOT_ONLY:
		return errors.New("// 名称已存在")
	case NET_DVR_OVER_MAX_ARRAY:
		return errors.New("// 阵列达到上限")
	case NET_DVR_OVER_MAX_VD:
		return errors.New("// 虚拟磁盘达到上限")
	case NET_DVR_VD_SLOT_EXCEED:
		return errors.New("// 虚拟磁盘槽位已满")
	case NET_DVR_PD_STATUS_INVALID:
		return errors.New("// 重建阵列所需物理磁盘状态错误")
	case NET_DVR_PD_BE_DEDICATE_SPARE:
		return errors.New("// 重建阵列所需物理磁盘为指定热备")
	case NET_DVR_PD_NOT_FREE:
		return errors.New("// 重建阵列所需物理磁盘非空闲")
	case NET_DVR_CANNOT_MIG2NEWMODE:
		return errors.New("// 不能从当前的阵列类型迁移到新的阵列类型")
	case NET_DVR_MIG_PAUSE:
		return errors.New("// 迁移操作已暂停")
	case NET_DVR_MIG_CANCEL:
		return errors.New("// 正在执行的迁移操作已取消")
	case NET_DVR_EXIST_VD:
		return errors.New("// 阵列上阵列上存在虚拟磁盘，无法删除阵列")
	case NET_DVR_TARGET_IN_LD_FUNCTIONAL:
		return errors.New("// 对象物理磁盘为虚拟磁盘组成部分且工作正常")
	case NET_DVR_HD_IS_ASSIGNED_ALREADY:
		return errors.New("// 指定的物理磁盘被分配为虚拟磁盘")
	case NET_DVR_INVALID_HD_COUNT:
		return errors.New("// 物理磁盘数量与指定的RAID等级不匹配")
	case NET_DVR_LD_IS_FUNCTIONAL:
		return errors.New("// 阵列正常，无法重建")
	case NET_DVR_BGA_RUNNING:
		return errors.New("// 存在正在执行的后台任务")
	case NET_DVR_LD_NO_ATAPI:
		return errors.New("// 无法用ATAPI盘创建虚拟磁盘")
	case NET_DVR_MIGRATION_NOT_NEED:
		return errors.New("// 阵列无需迁移")
	case NET_DVR_HD_TYPE_MISMATCH:
		return errors.New("// 物理磁盘不属于同意类型")
	case NET_DVR_NO_LD_IN_DG:
		return errors.New("// 无虚拟磁盘，无法进行此项操作")
	case NET_DVR_NO_ROOM_FOR_SPARE:
		return errors.New("// 磁盘空间过小，无法被指定为热备盘")
	case NET_DVR_SPARE_IS_IN_MULTI_DG:
		return errors.New("// 磁盘已被分配为某阵列热备盘")
	case NET_DVR_DG_HAS_MISSING_PD:
		return errors.New("// 阵列缺少盘")

	// x86 64bit nvr新增 2012-02-04
	case NET_DVR_NAME_EMPTY:
		return errors.New("// 名称为空")
	case NET_DVR_INPUT_PARAM:
		return errors.New("// 输入参数有误")
	case NET_DVR_PD_NOT_AVAILABLE:
		return errors.New("// 物理磁盘不可用")
	case NET_DVR_ARRAY_NOT_AVAILABLE:
		return errors.New("// 阵列不可用")
	case NET_DVR_PD_COUNT:
		return errors.New("// 物理磁盘数不正确")
	case NET_DVR_VD_SMALL:
		return errors.New("// 虚拟磁盘太小")
	case NET_DVR_NO_EXIST:
		return errors.New("// 不存在")
	case NET_DVR_NOT_SUPPORT:
		return errors.New("// 不支持该操作")
	case NET_DVR_NOT_FUNCTIONAL:
		return errors.New("// 阵列状态不是正常状态")
	case NET_DVR_DEV_NODE_NOT_FOUND:
		return errors.New("// 虚拟磁盘设备节点不存在")
	case NET_DVR_SLOT_EXCEED:
		return errors.New("// 槽位达到上限")
	case NET_DVR_NO_VD_IN_ARRAY:
		return errors.New("// 阵列上不存在虚拟磁盘")
	case NET_DVR_VD_SLOT_INVALID:
		return errors.New("// 虚拟磁盘槽位无效")
	case NET_DVR_PD_NO_ENOUGH_SPACE:
		return errors.New("// 所需物理磁盘空间不足")
	case NET_DVR_ARRAY_NONFUNCTION:
		return errors.New("// 只有处于正常状态的阵列才能进行迁移")
	case NET_DVR_ARRAY_NO_ENOUGH_SPACE:
		return errors.New("// 阵列空间不足")
	case NET_DVR_STOPPING_SCANNING_ARRAY:
		return errors.New("// 正在执行安全拔盘或重新扫描")
	case NET_DVR_NOT_SUPPORT_16T:
		return errors.New("// 不支持创建大于16T的阵列")
	case NET_DVR_ARRAY_FORMATING:
		return errors.New("// 正在执行格式化的阵列无法删除")
	case NET_DVR_QUICK_SETUP_PD_COUNT:
		return errors.New("// 一键配置至少需要三块空闲盘")

	//设备未激活时，登录失败，返回错误码
	case NET_DVR_ERROR_DEVICE_NOT_ACTIVATED:
		return errors.New("//设备未激活")
	//老SDK接新设备，设置用户密码或者激活的时候为风险密码时，错误码
	case NET_DVR_ERROR_RISK_PASSWORD:
		return errors.New("//有风险的密码")
	//已激活的设备，再次激活时返回错误码
	case NET_DVR_ERROR_DEVICE_HAS_ACTIVATED:
		return errors.New("//设备已激活")

		// 智能错误码
		// case VCA_ERROR_INDEX:
		// 	return errors.New("// 智能错误码索引")
		// case NET_DVR_ID_ERROR:
		// return errors.New("// 配置ID不合理")
	case NET_DVR_POLYGON_ERROR:
		return errors.New("// 多边形不符合要求")
	case NET_DVR_RULE_PARAM_ERROR:
		return errors.New("// 规则参数不合理")
	case NET_DVR_RULE_CFG_CONFLICT:
		return errors.New("// 配置信息冲突")
	case NET_DVR_CALIBRATE_NOT_READY:
		return errors.New("// 当前没有标定信息")
	case NET_DVR_CAMERA_DATA_ERROR:
		return errors.New("// 摄像机参数不合理")
	case NET_DVR_CALIBRATE_DATA_UNFIT:
		return errors.New("// 长度不够倾斜，不利于标定")
	case NET_DVR_CALIBRATE_DATA_CONFLICT:
		return errors.New("// 标定出错，以为所有点共线或者位置太集中")
	case NET_DVR_CALIBRATE_CALC_FAIL:
		return errors.New("// 摄像机标定参数值计算失败")
	case NET_DVR_CALIBRATE_LINE_OUT_RECT:
		return errors.New("// 输入的样本标定线超出了样本外接矩形框")
	case NET_DVR_ENTER_RULE_NOT_READY:
		return errors.New("// 没有设置进入区域")
	case NET_DVR_AID_RULE_NO_INCLUDE_LANE:
		return errors.New("// 交通事件规则中没有包括车道（特值拥堵和逆行）")
	case NET_DVR_LANE_NOT_READY:
		return errors.New("// 当前没有设置车道")
	case NET_DVR_RULE_INCLUDE_TWO_WAY:
		return errors.New("// 事件规则中包含2种不同方向")
	case NET_DVR_LANE_TPS_RULE_CONFLICT:
		return errors.New("// 车道和数据规则冲突")
	case NET_DVR_NOT_SUPPORT_EVENT_TYPE:
		return errors.New("// 不支持的事件类型")
	case NET_DVR_LANE_NO_WAY:
		return errors.New("// 车道没有方向")
	case NET_DVR_SIZE_FILTER_ERROR:
		return errors.New("// 尺寸过滤框不合理")
	case NET_DVR_LIB_FFL_NO_FACE:
		return errors.New("// 特征点定位时输入的图像没有人脸")
	case NET_DVR_LIB_FFL_IMG_TOO_SMALL:
		return errors.New("// 特征点定位时输入的图像太小")
	case NET_DVR_LIB_FD_IMG_NO_FACE:
		return errors.New("// 单张图像人脸检测时输入的图像没有人脸")
	case NET_DVR_LIB_FACE_TOO_SMALL:
		return errors.New("// 建模时人脸太小")
	case NET_DVR_LIB_FACE_QUALITY_TOO_BAD:
		return errors.New("// 建模时人脸图像质量太差")
	case NET_DVR_KEY_PARAM_ERR:
		return errors.New("//高级参数设置错误")
	case NET_DVR_CALIBRATE_DATA_ERR:
		return errors.New("//标定样本数目错误，或数据值错误，或样本点超出地平线")
	case NET_DVR_CALIBRATE_DISABLE_FAIL:
		return errors.New("//所配置规则不允许取消标定")
	case NET_DVR_VCA_LIB_FD_SCALE_OUTRANGE:
		return errors.New("//最大过滤框的宽高最小值超过最小过滤框的宽高最大值两倍以上")
	case NET_DVR_LIB_FD_REGION_TOO_LARGE:
		return errors.New("//当前检测区域范围过大。检测区最大为图像的2/3")
	case NET_DVR_TRIAL_OVERDUE:
		return errors.New("//试用版评估期已结束")
	case NET_DVR_CONFIG_FILE_CONFLICT:
		return errors.New("//设备类型与配置文件冲突（加密狗类型与现有分析仪配置不符错误码提示）")
	//算法库相关错误码
	case NET_DVR_FR_FPL_FAIL:
		return errors.New("// 人脸特征点定位失败")
	case NET_DVR_FR_IQA_FAIL:
		return errors.New("// 人脸评分失败")
	case NET_DVR_FR_FEM_FAIL:
		return errors.New("// 人脸特征提取失败")
	case NET_DVR_FPL_DT_CONF_TOO_LOW:
		return errors.New("// 特征点定位时人脸检测置信度过低")
	case NET_DVR_FPL_CONF_TOO_LOW:
		return errors.New("// 特征点定位置信度过低")
	case NET_DVR_E_DATA_SIZE:
		return errors.New("// 数据长度不匹配")
	case NET_DVR_FR_MODEL_VERSION_ERR:
		return errors.New("// 人脸模型数据中的模型版本错误")
	case NET_DVR_FR_FD_FAIL:
		return errors.New("// 识别库中人脸检测失败")
	case NET_DVR_FA_NORMALIZE_ERR:
		return errors.New("// 人脸归一化出错")
	//其他错误码
	case NET_DVR_DOG_PUSTREAM_NOT_MATCH:
		return errors.New("// 加密狗与前端取流设备类型不匹配")
	case NET_DVR_DEV_PUSTREAM_NOT_MATCH:
		return errors.New("// 前端取流设备版本不匹配")
	case NET_DVR_PUSTREAM_ALREADY_EXISTS:
		return errors.New("// 设备的其他通道已经添加过该前端设备")
	case NET_DVR_SEARCH_CONNECT_FAILED:
		return errors.New("// 连接检索服务器失败")
	case NET_DVR_INSUFFICIENT_DISK_SPACE:
		return errors.New("// 可存储的硬盘空间不足")
	case NET_DVR_DATABASE_CONNECTION_FAILED:
		return errors.New("// 数据库连接失败")
	case NET_DVR_DATABASE_ADM_PW_ERROR:
		return errors.New("// 数据库用户名、密码错误")
	case NET_DVR_DECODE_YUV:
		return errors.New("// 解码失败")
	case NET_DVR_IMAGE_RESOLUTION_ERROR:
		return errors.New("//")
	case NET_DVR_CHAN_WORKMODE_ERROR:
		return errors.New("//")

	case NET_DVR_RTSP_ERROR_NOENOUGHPRI:
		return errors.New("//无权限：服务器返回401时，转成这个错误码")
	case NET_DVR_RTSP_ERROR_ALLOC_RESOURCE:
		return errors.New("//分配资源失败")
	case NET_DVR_RTSP_ERROR_PARAMETER:
		return errors.New("//参数错误")
	case NET_DVR_RTSP_ERROR_NO_URL:
		return errors.New("//指定的URL地址不存在：服务器返回404时，转成这个错误码")
	case NET_DVR_RTSP_ERROR_FORCE_STOP:
		return errors.New("//用户中途强行退出")

	case NET_DVR_RTSP_GETPORTFAILED:
		return errors.New("//rtsp 得到端口错误")
	case NET_DVR_RTSP_DESCRIBERROR:
		return errors.New("//rtsp decribe 交互错误")
	case NET_DVR_RTSP_DESCRIBESENDTIMEOUT:
		return errors.New("//rtsp decribe 发送超时")
	case NET_DVR_RTSP_DESCRIBESENDERROR:
		return errors.New("//rtsp decribe 发送失败")
	case NET_DVR_RTSP_DESCRIBERECVTIMEOUT:
		return errors.New("//rtsp decribe 接收超时")
	case NET_DVR_RTSP_DESCRIBERECVDATALOST:
		return errors.New("//rtsp decribe 接收数据错误")
	case NET_DVR_RTSP_DESCRIBERECVERROR:
		return errors.New("//rtsp decribe 接收失败")
	case NET_DVR_RTSP_DESCRIBESERVERERR:
		return errors.New("//rtsp decribe 服务器返回错误状态")

	case NET_DVR_RTSP_SETUPERROR:
		return errors.New("//rtsp setup 交互错误")
	case NET_DVR_RTSP_SETUPSENDTIMEOUT:
		return errors.New("//rtsp setup 发送超时")
	case NET_DVR_RTSP_SETUPSENDERROR:
		return errors.New("//rtsp setup 发送错误")
	case NET_DVR_RTSP_SETUPRECVTIMEOUT:
		return errors.New("//rtsp setup 接收超时")
	case NET_DVR_RTSP_SETUPRECVDATALOST:
		return errors.New("//rtsp setup 接收数据错误")
	case NET_DVR_RTSP_SETUPRECVERROR:
		return errors.New("//rtsp setup 接收失败")
	case NET_DVR_RTSP_OVER_MAX_CHAN:
		return errors.New("//超过服务器最大连接数，或者服务器资源不足，服务器返回453时，转成这个错误码。")
	case NET_DVR_RTSP_SETUPSERVERERR:
		return errors.New("//rtsp setup 服务器返回错误状态")

	case NET_DVR_RTSP_PLAYERROR:
		return errors.New("//rtsp play 交互错误")
	case NET_DVR_RTSP_PLAYSENDTIMEOUT:
		return errors.New("//rtsp play 发送超时")
	case NET_DVR_RTSP_PLAYSENDERROR:
		return errors.New("//rtsp play 发送错误")
	case NET_DVR_RTSP_PLAYRECVTIMEOUT:
		return errors.New("//rtsp play 接收超时")
	case NET_DVR_RTSP_PLAYRECVDATALOST:
		return errors.New("//rtsp play 接收数据错误")
	case NET_DVR_RTSP_PLAYRECVERROR:
		return errors.New("//rtsp play 接收失败")
	case NET_DVR_RTSP_PLAYSERVERERR:
		return errors.New("//rtsp play 服务器返回错误状态")

	case NET_DVR_RTSP_TEARDOWNERROR:
		return errors.New("//rtsp teardown 交互错误")
	case NET_DVR_RTSP_TEARDOWNSENDTIMEOUT:
		return errors.New("//rtsp teardown 发送超时")
	case NET_DVR_RTSP_TEARDOWNSENDERROR:
		return errors.New("//rtsp teardown 发送错误")
	case NET_DVR_RTSP_TEARDOWNRECVTIMEOUT:
		return errors.New("//rtsp teardown 接收超时")
	case NET_DVR_RTSP_TEARDOWNRECVDATALOST:
		return errors.New("//rtsp teardown 接收数据错误")
	case NET_DVR_RTSP_TEARDOWNRECVERROR:
		return errors.New("//rtsp teardown 接收失败")
	case NET_DVR_RTSP_TEARDOWNSERVERERR:
		return errors.New("//rtsp teardown 服务器返回错误状态")

	case NET_PLAYM4_NOERROR:
		return errors.New("//no error")
	case NET_PLAYM4_PARA_OVER:
		return errors.New("//input parameter is invalid;")
	case NET_PLAYM4_ORDER_ERROR:
		return errors.New("//The order of the function to be called is error.")
	case NET_PLAYM4_TIMER_ERROR:
		return errors.New("//Create multimedia clock failed;")
	case NET_PLAYM4_DEC_VIDEO_ERROR:
		return errors.New("//Decode video data failed.")
	case NET_PLAYM4_DEC_AUDIO_ERROR:
		return errors.New("//Decode audio data failed.")
	case NET_PLAYM4_ALLOC_MEMORY_ERROR:
		return errors.New("//Allocate memory failed.")
	case NET_PLAYM4_OPEN_FILE_ERROR:
		return errors.New("//Open the file failed.")
	case NET_PLAYM4_CREATE_OBJ_ERROR:
		return errors.New("//Create thread or event failed")
	case NET_PLAYM4_CREATE_DDRAW_ERROR:
		return errors.New("//Create DirectDraw object failed.")
	case NET_PLAYM4_CREATE_OFFSCREEN_ERROR:
		return errors.New("//failed when creating off-screen surface.")
	case NET_PLAYM4_BUF_OVER:
		return errors.New("//buffer is overflow")
	case NET_PLAYM4_CREATE_SOUND_ERROR:
		return errors.New("//failed when creating audio device.")
	case NET_PLAYM4_SET_VOLUME_ERROR:
		return errors.New("//Set volume failed")
	case NET_PLAYM4_SUPPORT_FILE_ONLY:
		return errors.New("//The function only support play file.")
	case NET_PLAYM4_SUPPORT_STREAM_ONLY:
		return errors.New("//The function only support play stream.")
	case NET_PLAYM4_SYS_NOT_SUPPORT:
		return errors.New("//System not support.")
	case NET_PLAYM4_FILEHEADER_UNKNOWN:
		return errors.New("//No file header.")
	case NET_PLAYM4_VERSION_INCORRECT:
		return errors.New("//The version of decoder and encoder is not adapted.")
	case NET_PALYM4_INIT_DECODER_ERROR:
		return errors.New("//Initialize decoder failed.")
	case NET_PLAYM4_CHECK_FILE_ERROR:
		return errors.New("//The file data is unknown.")
	case NET_PLAYM4_INIT_TIMER_ERROR:
		return errors.New("//Initialize multimedia clock failed.")
	case NET_PLAYM4_BLT_ERROR:
		return errors.New("//Blt failed.")
	case NET_PLAYM4_UPDATE_ERROR:
		return errors.New("//Update failed.")
	case NET_PLAYM4_OPEN_FILE_ERROR_MULTI:
		return errors.New("//openfile error, streamtype is multi")
	case NET_PLAYM4_OPEN_FILE_ERROR_VIDEO:
		return errors.New("//openfile error, streamtype is video")
	case NET_PLAYM4_JPEG_COMPRESS_ERROR:
		return errors.New("//JPEG compress error")
	case NET_PLAYM4_EXTRACT_NOT_SUPPORT:
		return errors.New("//Don't support the version of this file.")
	case NET_PLAYM4_EXTRACT_DATA_ERROR:
		return errors.New("//extract video data failed.")

	//转封装库错误码
	case NET_CONVERT_ERROR_NOT_SUPPORT:
		return errors.New("//convert not support")

	//语音对讲库错误码
	case NET_AUDIOINTERCOM_OK:
		return errors.New("//无错误")
	case NET_AUDIOINTECOM_ERR_NOTSUPORT:
		return errors.New("//不支持")
	case NET_AUDIOINTECOM_ERR_ALLOC_MEMERY:
		return errors.New("//内存申请错误")
	case NET_AUDIOINTECOM_ERR_PARAMETER:
		return errors.New("//参数错误")
	case NET_AUDIOINTECOM_ERR_CALL_ORDER:
		return errors.New("//调用次序错误")
	case NET_AUDIOINTECOM_ERR_FIND_DEVICE:
		return errors.New("//未发现设备")
	case NET_AUDIOINTECOM_ERR_OPEN_DEVICE:
		return errors.New("//不能打开设备诶")
	case NET_AUDIOINTECOM_ERR_NO_CONTEXT:
		return errors.New("//设备上下文出错")
	case NET_AUDIOINTECOM_ERR_NO_WAVFILE:
		return errors.New("//WAV文件出错")
	case NET_AUDIOINTECOM_ERR_INVALID_TYPE:
		return errors.New("//无效的WAV参数类型")
	case NET_AUDIOINTECOM_ERR_ENCODE_FAIL:
		return errors.New("//编码失败")
	case NET_AUDIOINTECOM_ERR_DECODE_FAIL:
		return errors.New("//解码失败")
	case NET_AUDIOINTECOM_ERR_NO_PLAYBACK:
		return errors.New("//播放失败")
	case NET_AUDIOINTECOM_ERR_DENOISE_FAIL:
		return errors.New("//降噪失败")
	case NET_AUDIOINTECOM_ERR_UNKOWN:
		return errors.New("//未知错误")

	case NET_QOS_OK:
		return errors.New("//no error")
	case NET_QOS_ERROR:
		return errors.New("//qos error")
	case NET_QOS_ERR_INVALID_ARGUMENTS:
		return errors.New("//invalid arguments")
	case NET_QOS_ERR_SESSION_NOT_FOUND:
		return errors.New("//session net found")
	case NET_QOS_ERR_LIB_NOT_INITIALIZED:
		return errors.New("//lib not initialized")
	case NET_QOS_ERR_OUTOFMEM:
		return errors.New("//outtofmem")
	case NET_QOS_ERR_PACKET_UNKNOW:
		return errors.New("//packet unknow")
	case NET_QOS_ERR_PACKET_VERSION:
		return errors.New("//packet version error")
	case NET_QOS_ERR_PACKET_LENGTH:
		return errors.New("//packet length error")
	case NET_QOS_ERR_PACKET_TOO_BIG:
		return errors.New("//packet too big")
	case NET_QOS_ERR_SCHEDPARAMS_INVALID_BANDWIDTH:
		return errors.New("//schedparams invalid bandwidth")
	case NET_QOS_ERR_SCHEDPARAMS_BAD_FRACTION:
		return errors.New("//schedparams bad fraction")
	case NET_QOS_ERR_SCHEDPARAMS_BAD_MINIMUM_INTERVAL:
		return errors.New("//schedparams bad minimum interval")

	case NET_ERROR_TRUNK_LINE:
		return errors.New("//子系统已被配成干线")
	case NET_ERROR_MIXED_JOINT:
		return errors.New("//不能进行混合拼接")
	case NET_ERROR_DISPLAY_SWITCH:
		return errors.New("//不能进行显示通道切换")
	case NET_ERROR_USED_BY_BIG_SCREEN:
		return errors.New("//解码资源被大屏占用")
	case NET_ERROR_USE_OTHER_DEC_RESOURCE:
		return errors.New("//不能使用其他解码子系统资源")
	case NET_ERROR_DISP_MODE_SWITCH:
		return errors.New("//显示通道显示状态切换中")
	case NET_ERROR_SCENE_USING:
		return errors.New("//场景正在使用")
	case NET_ERR_NO_ENOUGH_DEC_RESOURCE:
		return errors.New("//解码资源不足")
	case NET_ERR_NO_ENOUGH_FREE_SHOW_RESOURCE:
		return errors.New("//畅显资源不足")
	case NET_ERR_NO_ENOUGH_VIDEO_MEMORY:
		return errors.New("//显存资源不足")
	case NET_ERR_MAX_VIDEO_NUM:
		return errors.New("//一拖多资源不足")
	case NET_ERR_WIN_COVER_FREE_SHOW_AND_NORMAL:
		return errors.New("//窗口跨越了畅显输出口和非畅显输出口")
	case NET_ERR_FREE_SHOW_WIN_SPLIT:
		return errors.New("//畅显窗口不支持分屏")
	case NET_ERR_INAPPROPRIATE_WIN_FREE_SHOW:
		return errors.New("//不是输出口整数倍的窗口不支持开启畅显")
	case NET_DVR_TRANSPARENT_WIN_NOT_SUPPORT_SPLIT:
		return errors.New("//开启透明度的窗口不支持分屏")
	case NET_DVR_SPLIT_WIN_NOT_SUPPORT_TRANSPARENT:
		return errors.New("//开启多分屏的窗口不支持透明度设置")
	case NET_ERR_MAX_LOGO_NUM:
		return errors.New("//logo数达到上限")
	case NET_ERR_MAX_WIN_LOOP_NUM:
		return errors.New("//轮巡窗口数达到上限")
	case NET_ERR_VIRTUAL_LED_VERTICAL_CROSS:
		return errors.New("//虚拟LED不能纵向跨屏")
	case NET_ERR_MAX_VIRTUAL_LED_HEIGHT:
		return errors.New("//虚拟LED高度超限")
	case NET_ERR_VIRTUAL_LED_ILLEGAL_CHARACTER:
		return errors.New("//虚拟LED内容包含非法字符")
	case NET_ERR_BASEMAP_NOT_EXIST:
		return errors.New("//底图图片不存在")
	case NET_ERR_LED_NOT_SUPPORT_VIRTUAL_LED:
		return errors.New("//LED屏幕不支持虚拟LED")
	case NET_ERR_LED_RESOLUTION_NOT_SUPPORT:
		return errors.New("//LED分辨率不支持")
	case NET_ERR_PLAN_OVERDUE:
		return errors.New("//预案超期，不能再调用")
	case NET_ERR_PROCESSER_MAX_SCREEN_BLK:
		return errors.New("//单个处理器接入的信号跨越的屏幕个数超限")
	case NET_ERR_WND_SIZE_TOO_SMALL:
		return errors.New("//开窗窗口宽高太小")
	case NET_ERR_WND_SPLIT_NOT_SUPPORT_ROAM:
		return errors.New("//分屏窗口不支持漫游")
	case NET_ERR_OUTPUT_ONE_BOARD_ONE_WALL:
		return errors.New("//同一个子板的输出口只能绑定到同一面墙上")
	case NET_ERR_WND_CANNOT_LCD_AND_LED_OUTPUT:
		return errors.New("//窗口不能同时跨LCD和LED输出口")
	case NET_ERR_MAX_OSD_NUM:
		return errors.New("//OSD数量达到最大")

	case NET_SDK_CANCEL_WND_TOPKEEP_ATTR_FIRST:
		return errors.New("//先取消置顶保持窗口的置顶保持属性才能进行置底操作")
	case NET_SDK_ERR_LED_SCREEN_CHECKING:
		return errors.New("//正在校正LED屏幕")
	case NET_SDK_ERR_NOT_SUPPORT_SINGLE_RESOLUTION:
		return errors.New("//LCD/LED输出口绑定之后不支持单个输出口的分辨率配置")
	case NET_SDK_ERR_LED_RESOLUTION_MISMATCHED:
		return errors.New("//该输出口的LED分辨率和其他输出口的LED分辨率不匹配��需要满足同行等高、同列等宽")

	case NET_SDK_ERR_MAX_VIRTUAL_LED_WIDTH:
		return errors.New("//虚拟LED宽度超限，包括最大值和最小值")
	case NET_SDK_ERR_MAX_VIRTUAL_LED_IN_SCREEN:
		return errors.New("//单屏虚拟LED数量超限")
	case NET_SDK_ERR_MAX_VIRTUAL_LED_IN_WALL:
		return errors.New("//单墙虚拟LED数量超限")
	case NET_SDK_ERR_VIRTUAL_LED_OVERLAP:
		return errors.New("//虚拟LED重叠错误")
	case NET_SDK_ERR_VIRTUAL_LED_TYPE:
		return errors.New("//类型错误")
	case NET_SDK_ERR_VIRTUAL_LED_COLOUR:
		return errors.New("//颜色错误")
	case NET_SDK_ERR_VIRTUAL_LED_MOVE_DIRECTION:
		return errors.New("//移动方向错误")
	case NET_SDK_ERR_VIRTUAL_LED_MOVE_MODE:
		return errors.New("//移动模式错误")
	case NET_SDK_ERR_VIRTUAL_LED_MOVE_SPEED:
		return errors.New("//移动速度错误")
	case NET_SDK_ERR_VIRTUAL_LED_DISP_MODE:
		return errors.New("//显示模式有误")
	case NET_SDK_ERR_VIRTUAL_LED_NO:
		return errors.New("//虚拟LED序号错误")
	case NET_SDK_ERR_VIRTUAL_LED_PARA:
		return errors.New("//虚拟LED参数配置错误，包括结构体内其他参数")
	case NET_SDK_ERR_BASEMAP_POSITION:
		return errors.New("//底图窗口宽高参数错误")
	case NET_SDK_ERR_BASEMAP_PICTURE_LEN:
		return errors.New("//底图图片长度超限")
	case NET_SDK_ERR_BASEMAP_PICTURE_RESOLUTION:
		return errors.New("//底图图片分辨率错误")
	case NET_SDK_ERR_BASEMAP_PICTURE_FORMAT:
		return errors.New("//底图图片格式错误")
	case NET_SDK_ERR_MAX_VIRTUAL_LED_NUM:
		return errors.New("//设备支持的虚拟LED数量超限")
	case NET_SDK_ERR_MAX_TIME_VIRTUAL_LED_IN_WALL:
		return errors.New("//单面电视墙支持的时间虚拟LED的数量超限")

	case NET_ERR_TERMINAL_BUSY:
		return errors.New("//终端忙，终端处于会议中")

	case NET_ERR_DATA_RETURNED_ILLEGAL:
		return errors.New("//设备返回的数据不合法")
	case NET_DVR_FUNCTION_RESOURCE_USAGE_ERROR:
		return errors.New("//设备其它功能占用资源，导致该功能无法开启")

	case NET_DVR_ERR_IMPORT_EMPTY_FILE:
		return errors.New("//导入文件为空")
	case NET_DVR_ERR_IMPORT_TOO_LARGE_FILE:
		return errors.New("//导入文件过大")
	case NET_DVR_ERR_BAD_IPV4_ADDRESS:
		return errors.New("//IPV4地址无效")
	case NET_DVR_ERR_BAD_NET_MASK:
		return errors.New("//子网掩码地址无效")
	case NET_DVR_ERR_INVALID_NET_GATE_ADDRESS:
		return errors.New("//网关地址无效")
	case NET_DVR_ERR_BAD_DNS:
		return errors.New("//DNS地址无效")
	case NET_DVR_ERR_ILLEGAL_PASSWORD:
		return errors.New("//密码不能包含用户名")

	case NET_DVR_DEV_NET_OVERFLOW:
		return errors.New("//网络流量超过设备能力上限")
	case NET_DVR_STATUS_RECORDFILE_WRITING_NOT_LOCK:
		return errors.New("//录像文件在录像，无法被锁定")
	case NET_DVR_STATUS_CANT_FORMAT_LITTLE_DISK:
		return errors.New("//由于硬盘太小无法格式化")

	//N+1错误码
	case NET_SDK_ERR_REMOTE_DISCONNECT:
		return errors.New("//远端无法连接")
	case NET_SDK_ERR_RD_ADD_RD:
		return errors.New("//备机不能添加备机")
	case NET_SDK_ERR_BACKUP_DISK_EXCEPT:
		return errors.New("//备份盘异常")
	case NET_SDK_ERR_RD_LIMIT:
		return errors.New("//备机数已达上限")
	case NET_SDK_ERR_ADDED_RD_IS_WD:
		return errors.New("//添加的备机是工作机")
	case NET_SDK_ERR_ADD_ORDER_WRONG:
		return errors.New("//添加顺序出错，比如没有被工作机添加为备机，就添加工作机")
	case NET_SDK_ERR_WD_ADD_WD:
		return errors.New("//工作机不能添加工作机")
	case NET_SDK_ERR_WD_SERVICE_EXCETP:
		return errors.New("//工作机CVR服务异常")
	case NET_SDK_ERR_RD_SERVICE_EXCETP:
		return errors.New("//备机CVR服务异常")
	case NET_SDK_ERR_ADDED_WD_IS_RD:
		return errors.New("//添加的工作机是备机")
	case NET_SDK_ERR_PERFORMANCE_LIMIT:
		return errors.New("//性能达到上限")
	case NET_SDK_ERR_ADDED_DEVICE_EXIST:
		return errors.New("//添加的设备已经存在")

	//审讯机错误码
	case NET_SDK_ERR_INQUEST_RESUMING:
		return errors.New("//审讯恢复中")
	case NET_SDK_ERR_RECORD_BACKUPING:
		return errors.New("//审讯备份中")
	case NET_SDK_ERR_DISK_PLAYING:
		return errors.New("//光盘回放中")
	case NET_SDK_ERR_INQUEST_STARTED:
		return errors.New("//审讯已开启")
	case NET_SDK_ERR_LOCAL_OPERATING:
		return errors.New("//本地操作进行中")
	case NET_SDK_ERR_INQUEST_NOT_START:
		return errors.New("//审讯未开启")
	//Netra3.1.0错误码
	case NET_SDK_ERR_CHAN_AUDIO_BIND:
		return errors.New("//通道未绑定或绑定语音对讲失败")
	//云存储错误码
	case NET_DVR_N_PLUS_ONE_MODE:
		return errors.New("//设备当前处于N+1模式")
	case NET_DVR_CLOUD_STORAGE_OPENED:
		return errors.New("//云存储模式已开启")

	case NET_DVR_ERR_OPER_NOT_ALLOWED:
		return errors.New("//设备处于N+0被接管状态，不允许该操作")
	case NET_DVR_ERR_NEED_RELOCATE:
		return errors.New("//设备处于N+0被接管状态，需要获取重定向信息，再重新操作")

	//庭审主机错误码
	case NET_SDK_ERR_IR_PORT_ERROR:
		return errors.New("//红外输出口错误")
	case NET_SDK_ERR_IR_CMD_ERROR:
		return errors.New("//红外输出口的命令号错误")
	case NET_SDK_ERR_NOT_INQUESTING:
		return errors.New("//设备处于非审讯状态")
	case NET_SDK_ERR_INQUEST_NOT_PAUSED:
		return errors.New("//设备处于非暂停状态")
	case NET_DVR_CHECK_PASSWORD_MISTAKE_ERROR:
		return errors.New("//校验密码错误")
	case NET_DVR_CHECK_PASSWORD_NULL_ERROR:
		return errors.New("//校验密码不能为空")
	case NET_DVR_UNABLE_CALIB_ERROR:
		return errors.New("// 当前无法标定")
	case NET_DVR_PLEASE_CALIB_ERROR:
		return errors.New("//请先完成标定")
	case NET_DVR_ERR_PANORAMIC_CAL_EMPTY:
		return errors.New("//Flash中全景标定为空")
	case NET_DVR_ERR_CALIB_FAIL_PLEASEAGAIN:
		return errors.New("//标定失败，请重新标定(Calibration failed. Please calibrate again.)")
	case NET_DVR_ERR_DETECTION_LINE:
		return errors.New("//规则线配置错误，请重新配置规则线，确保规则线位于红色区域内(Please set detection line again. The detection line should be within the red count area.)")
	case NET_DVR_ERR_TURN_OFF_IMAGE_PARA:
		return errors.New("//请先关闭图像参数切换功能(Please turn off the image parameters switch first.)")
	case NET_DVR_EXCEED_FACE_IMAGES_ERROR:
		return errors.New("//超过人脸图片最大张数")
	case NET_DVR_ANALYSIS_FACE_IMAGES_ERROR:
		return errors.New("//图片数据识别失败")
	case NET_ERR_ALARM_INPUT_OCCUPIED:
		return errors.New("//A<-1报警号已用于触发车辆抓拍Alarm Input No. A<-1 is used to trigger vehicle capture.")
	case NET_DVR_FACELIB_DATABASE_ERROR:
		return errors.New("//人脸库中数据库版本不匹配")
	case NET_DVR_FACELIB_DATA_ERROR:
		return errors.New("//人脸库数据错误")
	case NET_DVR_FACE_DATA_ID_ERROR:
		return errors.New("//人脸数据PID无效")
	case NET_DVR_FACELIB_ID_ERROR:
		return errors.New("//人脸库ID无效")
	case NET_DVR_EXCEED_FACE_LIBARY_ERROR:
		return errors.New("//超过人脸库最大个数")
	case NET_DVR_PIC_ANALYSIS_NO_TARGET_ERROR:
		return errors.New("//图片未识别到目标")
	case NET_DVR_SUBPIC_ANALYSIS_MODELING_ERROR:
		return errors.New("//子图建模失败")
	case NET_DVR_PIC_ANALYSIS_NO_RESOURCE_ERROR:
		return errors.New("//无对应智能分析引擎支持图片二次识别")
	case NET_DVR_ANALYSIS_ENGINES_NO_RESOURCE_ERROR:
		return errors.New("//无分析引擎资源")
	case NET_DVR_ANALYSIS_ENGINES_USAGE_EXCEED_ERROR:
		return errors.New("//引擎使用率超负荷，已达100%")
	case NET_DVR_EXCEED_HUMANMISINFO_FILTER_ENABLED_ERROR:
		return errors.New("//超过开启人体去误报最大通道个数")
	case NET_DVR_NAME_ERROR:
		return errors.New("//名称错误")
	case NET_DVR_NAME_EXIST_ERROR:
		return errors.New("//名称已存在")
	case NET_DVR_FACELIB_PIC_IMPORTING_ERROR:
		return errors.New("//人脸库导入图片中")
	case NET_DVR_ERR_CALIB_POSITION:
		return errors.New("//标定位置超出摄像机运动范围")
	case NET_DVR_ERR_DELETE:
		return errors.New("//无法删除")
	case NET_DVR_ERR_SCENE_ID:
		return errors.New("//场景ID无效")
	case NET_DVR_ERR_CALIBING:
		return errors.New("//标定中")
	case NET_DVR_PIC_FORMAT_ERROR:
		return errors.New("//图片格式错误")
	case NET_DVR_PIC_RESOLUTION_INVALID_ERROR:
		return errors.New("//图片分辨率无效错误")
	case NET_DVR_PIC_SIZE_EXCEED_ERROR:
		return errors.New("//图片过大")
	case NET_DVR_PIC_ANALYSIS_TARGRT_NUM_EXCEED_ERROR:
		return errors.New("//图片目标个数超过上限")
	case NET_DVR_ANALYSIS_ENGINES_LOADING_ERROR:
		return errors.New("//分析引擎初始化中")
	case NET_DVR_ANALYSIS_ENGINES_ABNORMA_ERROR:
		return errors.New("//分析引擎异常")
	case NET_DVR_ANALYSIS_ENGINES_FACELIB_IMPORTING:
		return errors.New("//分析引擎正在导入人脸库")
	case NET_DVR_NO_DATA_FOR_MODELING_ERROR:
		return errors.New("//无待建模数据")
	case NET_DVR_FACE_DATA_MODELING_ERROR:
		return errors.New("//设备正在进行图片建模操作，不支持并发处理")
	case NET_ERR_FACELIBDATA_OVERLIMIT:
		return errors.New("//超过设备中支持导入人脸数最大个数限制（导入的人脸库中数据）")
	case NET_DVR_ANALYSIS_ENGINES_ASSOCIATED_CHANNEL:
		return errors.New("//分析引擎已关联通道")
	case NET_DVR_ERR_CUSTOMID_LEN:
		return errors.New("//上层自定义ID的长度最小32字符长度")
	case NET_DVR_ERR_CUSTOMFACELIBID_REPEAT:
		return errors.New("//上层下发重复的自定义人脸库ID")
	case NET_DVR_ERR_CUSTOMHUMANID_REPEAT:
		return errors.New("//上层下发重复的自定义人员ID")
	case NET_DVR_ERR_URL_DOWNLOAD_FAIL:
		return errors.New("//url下载失败")
	case NET_DVR_ERR_URL_DOWNLOAD_NOTSTART:
		return errors.New("//url未开始下载")

	case NET_DVR_CFG_FILE_SECRETKEY_ERROR:
		return errors.New("//配置文件安全校验密钥错误")
	case NET_DVR_WDR_NOTDISABLE_ERROR:
		return errors.New("//请先关闭所有通道当前日夜参数转换模式下的宽动态")
	case NET_DVR_HLC_NOTDISABLE_ERROR:
		return errors.New("//请先关闭所有通道当前日夜参数转换模式下的强光抑制")

	case NET_DVR_THERMOMETRY_REGION_OVERSTEP_ERROR:
		return errors.New("//测温区域越界")

	case NET_DVR_ERR_MODELING_DEVICEINTERNAL:
		return errors.New("//建模失败，设备内部错误")
	case NET_DVR_ERR_MODELING_FACE:
		return errors.New("//建模失败，人脸建模错误")
	case NET_DVR_ERR_MODELING_FACEGRADING:
		return errors.New("//建模失败，人脸质量评分错误")
	case NET_DVR_ERR_MODELING_FACEGFEATURE:
		return errors.New("//建模失败，特征点提取错误")
	case NET_DVR_ERR_MODELING_FACEGANALYZING:
		return errors.New("//建模失败，属性提取错误")

	case NET_DVR_ERR_STREAM_LIMIT:
		return errors.New("//码流性能超过上限，请减少取流路数")
	case NET_DVR_ERR_STREAM_DESCRIPTION:
		return errors.New("//请输入码流描述")
	case NET_DVR_ERR_STREAM_DELETE:
		return errors.New("//码流正在使用无法删除")
	case NET_DVR_ERR_CUSTOMSTREAM_NAME:
		return errors.New("//自定义码流名称为空或不合法")
	case NET_DVR_ERR_CUSTOMSTREAM_NOTEXISTED:
		return errors.New("//该自定义码流不存在")

	case NET_DVR_ERR_TOO_SHORT_CALIBRATING_TIME:
		return errors.New("//标定时间太短")
	case NET_DVR_ERR_AUTO_CALIBRATE_FAILED:
		return errors.New("//自动标定失败")
	case NET_DVR_ERR_VERIFICATION_FAILED:
		return errors.New("//校验失败")

	case NET_DVR_NO_TEMP_SENSOR_ERROR:
		return errors.New("//无温度传感器")
	case NET_DVR_PUPIL_DISTANCE_OVERSIZE_ERROR:
		return errors.New("//瞳距过大")
	case NET_DVR_ERR_UNOPENED_FACE_SNAP:
		return errors.New("//操作无效，请先开启人脸抓拍")
	//2011-10-25多屏控制器错误码（900-950）
	case NET_ERR_CUT_INPUTSTREAM_OVERLIMIT:
		return errors.New("//信号源裁剪数值超限")
	case NET_ERR_WINCHAN_IDX:
		return errors.New("// 开窗通道号错误")
	case NET_ERR_WIN_LAYER:
		return errors.New("// 窗口层数错误，单个屏幕上最多覆盖的窗口层数")
	case NET_ERR_WIN_BLK_NUM:
		return errors.New("// 窗口的块数错误，单个窗口可覆盖的屏幕个数")
	case NET_ERR_OUTPUT_RESOLUTION:
		return errors.New("// 输出分辨率错误")
	case NET_ERR_LAYOUT:
		return errors.New("// 布局号错误")
	case NET_ERR_INPUT_RESOLUTION:
		return errors.New("// 输入分辨率不支持")
	case NET_ERR_SUBDEVICE_OFFLINE:
		return errors.New("// 子设备不在线")
	case NET_ERR_NO_DECODE_CHAN:
		return errors.New("// 没有空闲解码通道")
	case NET_ERR_MAX_WINDOW_ABILITY:
		return errors.New("// 开窗能力上限, 分布式多屏控制器中解码子设备能力上限或者显示处理器能力上限导致")
	case NET_ERR_ORDER_ERROR:
		return errors.New("// 调用顺序有误")
	case NET_ERR_PLAYING_PLAN:
		return errors.New("// 正在执行预案")
	case NET_ERR_DECODER_USED:
		return errors.New("// 解码板正在使用")
	case NET_ERR_OUTPUT_BOARD_DATA_OVERFLOW:
		return errors.New("// 输出板数据量超限")
	case NET_ERR_SAME_USER_NAME:
		return errors.New("// 用户名相同")
	case NET_ERR_INVALID_USER_NAME:
		return errors.New("// 无效用户名")
	case NET_ERR_MATRIX_USING:
		return errors.New("// 输入矩阵正在使用")
	case NET_ERR_DIFFERENT_CHAN_TYPE:
		return errors.New("// 通道类型不同（矩阵输出通道和控制器的输入为不同的类型）")
	case NET_ERR_INPUT_CHAN_BINDED:
		return errors.New("// 输入通道已经被其他矩阵绑定")
	case NET_ERR_BINDED_OUTPUT_CHAN_OVERFLOW:
		return errors.New("// 正在使用的矩阵输出通道个数超过矩阵与控制器绑定的通道个数")
	case NET_ERR_MAX_SIGNAL_NUM:
		return errors.New("// 输入信号源个数达到上限")
	case NET_ERR_INPUT_CHAN_USING:
		return errors.New("// 输入通道正在使用")
	case NET_ERR_MANAGER_LOGON:
		return errors.New("// 管理员已经登陆，操作失败")
	case NET_ERR_USERALREADY_LOGON:
		return errors.New("// 该用户已经登陆，操作失败")
	case NET_ERR_LAYOUT_INIT:
		return errors.New("// 布局正在初始化，操作失败")
	case NET_ERR_BASEMAP_SIZE_NOT_MATCH:
		return errors.New("// 底图大小不符")
	case NET_ERR_WINDOW_OPERATING:
		return errors.New("// 窗口正在执行其他操作，本次操作失败")
	case NET_ERR_SIGNAL_UPLIMIT:
		return errors.New("// 信号源开窗个数达到上限")
	case NET_ERR_SIGNAL_MAX_ENLARGE_TIMES:
		return errors.New("// 信号源放大倍数超限")
	case NET_ERR_ONE_SIGNAL_MULTI_CROSS:
		return errors.New("// 单个信号源不能多次跨屏")
	case NET_ERR_ULTRA_HD_SIGNAL_MULTI_WIN:
		return errors.New("// 超高清信号源不能重复开窗")
	case NET_ERR_MAX_VIRTUAL_LED_WIDTH:
		return errors.New("//虚拟LED宽度大于限制值")
	case NET_ERR_MAX_VIRTUAL_LED_WORD_LEN:
		return errors.New("//虚拟LED字符数大于限制值")
	case NET_ERR_SINGLE_OUTPUTPARAM_CONFIG:
		return errors.New("//不支持单个显示输出参数设置")
	case NET_ERR_MULTI_WIN_BE_COVER:
		return errors.New("//多分屏窗口被覆盖")
	case NET_ERR_WIN_NOT_EXIST:
		return errors.New("//窗口不存在")
	case NET_ERR_WIN_MAX_SIGNALSOURCE:
		return errors.New("//窗口信号源数超过限制值")
	case NET_ERR_MULTI_WIN_MOVE:
		return errors.New("//对多分屏窗口移动")
	case NET_ERR_MULTI_WIN_YPBPR_SDI:
		return errors.New("// YPBPR 和SDI信号源不支持9/16分屏")
	case NET_ERR_DIFF_TYPE_OUTPUT_MIXUSE:
		return errors.New("//不同类型输出板混插")
	case NET_ERR_SPLIT_WIN_CROSS:
		return errors.New("//对跨屏窗口分屏")
	case NET_ERR_SPLIT_WIN_NOT_FULL_SCREEN:
		return errors.New("//对未满屏窗口分屏")
	case NET_ERR_SPLIT_WIN_MANY_WIN:
		return errors.New("//对单个输出口上有多个窗口的窗口分屏")
	case NET_ERR_WINDOW_SIZE_OVERLIMIT:
		return errors.New("//窗口大小超限")
	case NET_ERR_INPUTSTREAM_ALREADY_JOINT:
		return errors.New("//信号源已加入拼接")
	case NET_ERR_JOINT_INPUTSTREAM_OVERLIMIT:
		return errors.New("//拼接信号源个数超限")

	case NET_ERR_LED_RESOLUTION:
		return errors.New("//LED 分辨率大于输出分辨率")
	case NET_ERR_JOINT_SCALE_OVERLIMIT:
		return errors.New("//拼接信号源的规模超限")
	case NET_ERR_INPUTSTREAM_ALREADY_DECODE:
		return errors.New("//信号源已上墙")
	case NET_ERR_INPUTSTREAM_NOTSUPPORT_CAPTURE:
		return errors.New("//信号源不支持抓图")
	case NET_ERR_JOINT_NOTSUPPORT_SPLITWIN:
		return errors.New("//拼接信号源不支持分屏")

	//解码器错误码（951-999）
	case NET_ERR_MAX_WIN_OVERLAP:
		return errors.New("//达到最大窗口重叠数")
	case NET_ERR_STREAMID_CHAN_BOTH_VALID:
		return errors.New("//stream ID和通道号同时有效")
	case NET_ERR_NO_ZERO_CHAN:
		return errors.New("//设备无零通道")
	case NEED_RECONNECT:
		return errors.New("//需要重定向（转码子系统使用）")
	case NET_ERR_NO_STREAM_ID:
		return errors.New("//流ID不存在")
	case NET_DVR_TRANS_NOT_START:
		return errors.New("//转码未启动")
	case NET_ERR_MAXNUM_STREAM_ID:
		return errors.New("//流ID数达到上限")
	case NET_ERR_WORKMODE_MISMATCH:
		return errors.New("//工作模式不匹配")
	case NET_ERR_MODE_IS_USING:
		return errors.New("//已工作在当前模式")
	case NET_ERR_DEV_PROGRESSING:
		return errors.New("//设备正在处理中")
	case NET_ERR_PASSIVE_TRANSCODING:
		return errors.New("//正在被动转码")

	case NET_ERR_RING_NOT_CONFIGURE:
		return errors.New("//环网未配置")

	case NET_ERR_CLOSE_WINDOW_FIRST:
		return errors.New("//切换全帧率畅显时必须先关闭对应的已上墙的窗口")
	case NET_ERR_SPLIT_WINDOW_NUM_NOT_SUPPORT:
		return errors.New("//VGA/DVI/DP/HDMI/HDBase_T输入源在全帧率畅显下不支持9/16画面")
	case NET_ERR_REACH_ONE_SIGNAL_PREVIEW_MAX_LINK:
		return errors.New("//单信号源回显连接数量超限")
	case NET_ERR_ONLY_SPLITWND_SUPPORT_AMPLIFICATION:
		return errors.New("//只有分屏窗口支持子窗口放大")
	case NET_DVR_ERR_WINDOW_SIZE_PLACE:
		return errors.New("//窗口位置错误")
	case NET_DVR_ERR_RGIONAL_RESTRICTIONS:
		return errors.New("//屏幕距离超限")
	case NET_ERR_WNDZOOM_NOT_SUPPORT:
		return errors.New("//单窗口不支持子窗口全屏功能")
	case NET_ERR_LED_SCREEN_SIZE:
		return errors.New("//LED屏宽高不正确")
	case NET_ERR_OPEN_WIN_IN_ERROR_AREA:
		return errors.New("//在非法区域开窗(包括跨LCD/LED屏)")
	case NET_ERR_TITLE_WIN_NOT_SUPPORT_MOVE:
		return errors.New("//平铺模式不支持漫游")
	case NET_ERR_TITLE_WIN_NOT_SUPPORT_COVER:
		return errors.New("//平铺模式不支持图层覆盖")
	case NET_ERR_TITLE_WIN_NOT_SUPPORT_SPLIT:
		return errors.New("//平铺模式不支持分屏")
	case NET_DVR_LED_WINDOWS_ALREADY_CLOSED:
		return errors.New("//LED区域内输出口的分辨率发生变化，设备已关闭该区域内的所有LED窗口")
	case NET_DVR_ERR_CLOSE_WINDOWS:
		return errors.New("//操作失败，请先关闭窗口")
	case NET_DVR_ERR_MATRIX_LOOP_ABILITY:
		return errors.New("//超出轮巡解码能力限制")
	case NET_DVR_ERR_MATRIX_LOOP_TIME:
		return errors.New("//轮巡解码时间不支持")
	case NET_DVR_ERR_LINKED_OUT_ABILITY:
		return errors.New("//联动通道数超过上限")
	case NET_ERR_REACH_SCENE_MAX_NUM:
		return errors.New("//场景数量达到上限")
	case NET_ERR_SCENE_MEM_NOT_ENOUGH:
		return errors.New("//内存不足，无法新建场景")
	case NET_ERR_RESOLUTION_NOT_SUPPORT_ODD_VOUT:
		return errors.New("//奇口不支持该分辨率")
	case NET_ERR_RESOLUTION_NOT_SUPPORT_EVEN_VOUT:
		return errors.New("//偶口不支持该分辨率")

	case NET_DVR_CANCEL_WND_OPENKEEP_ATTR_FIRST:
		return errors.New("//常开窗口需要先取消常开属性才能被关闭")
	case NET_SDK_LED_MODE_NOT_SUPPORT_SPLIT:
		return errors.New("//LED模式下不支持窗口分屏")
	case NET_ERR_VOICETALK_ONLY_SUPPORT_ONE_TALK:
		return errors.New("//同时只支持一路语音对讲")
	case NET_ERR_WND_POSITION_ADJUSTED:
		return errors.New("//窗口位置被设备调整，上层需要重新获取下窗口位置")
	case NET_SDK_ERR_STARTTIME_CANNOT_LESSTHAN_CURTIME:
		return errors.New("//开始时间不能小于当前时间")
	case NET_SDK_ERR_NEED_ADJUST_PLAN:
		return errors.New("//场景已被预案关联，请先将该场景从预案中删除")
	case NET_ERR_UnitConfig_Failed:
		return errors.New("//当“启用单位统一”勾选时，测温下配置的单位与系统设置下的单位不同返回单位配置错误")

	//能力集解析库错误码
	case XML_ABILITY_NOTSUPPORT:
		return errors.New("//不支持能力节点获取")
	case XML_ANALYZE_NOENOUGH_BUF:
		return errors.New("//输出内存不足")
	case XML_ANALYZE_FIND_LOCALXML_ERROR:
		return errors.New("//无法找到对应的本地xml")
	case XML_ANALYZE_LOAD_LOCALXML_ERROR:
		return errors.New("//加载本地xml出错")
	case XML_NANLYZE_DVR_DATA_FORMAT_ERROR:
		return errors.New("//设备能力数据格式错误")
	case XML_ANALYZE_TYPE_ERROR:
		return errors.New("//能力集类型错误")
	case XML_ANALYZE_XML_NODE_ERROR:
		return errors.New("//XML能力节点格式错误")
	case XML_INPUT_PARAM_ERROR:
		return errors.New("//输入的能力XML节点值错误")

	case NET_DVR_ERR_RETURNED_XML_DATA:
		return errors.New("//设备返回的XML数据有误")

	//传显错误码
	case NET_ERR_LEDAREA_EXIST_WINDOW:
		return errors.New("//LED区域有窗口存在(如果LED区域上已经有窗口存在，不允许修改LED区域）")
	case NET_ERR_AUDIO_EXIST:
		return errors.New("//输出口上存在音频输出，不允许解除绑定")
	case NET_ERR_MATERIAL_NAME_EXIST:
		return errors.New("//素材名称已存在")
	case NET_ERR_MATERIAL_APPROVE_STATE:
		return errors.New("//素材审核状态错误")
	case NET_ERR_DATAHD_SIGNAL_FORMAT:
		return errors.New("//已使用的硬盘不允许单个格式化")

	case NET_ERR_SCENE_SWITCHING:
		return errors.New("//场景正在切换")
	case NER_ERR_DATA_TRANSFER:
		return errors.New("//设备正在数据转移中")
	case NET_ERR_DATA_RESTORE:
		return errors.New("//设备正在数据还原中")
	case NET_ERR_CHECK_NOT_ENABLE:
		return errors.New("//校正使能未开启")
	case NET_ERR_AREA_OFFLINE:
		return errors.New("//区域不在线")
	case NET_ERR_SCREEN_TYPE:
		return errors.New("//屏幕类型不匹配")
	case NET_ERR_MIN_OPERATE_UNIT:
		return errors.New("//最小操作单元不匹配")
	case NET_ERR_MAINHD_NOT_BACKUP:
		return errors.New("//第一槽位上的普通盘（主盘）禁止设置成备份盘")
	case NET_ERR_ONE_BACKUP_HD:
		return errors.New("//设置备份盘时，设备至少有一块普通盘")
	case NET_ERR_CONNECT_SUB_SYSTEM_ABNORMAL:
		return errors.New("//连接子系统异常")
	case NET_ERR_SERIAL_PORT_VEST:
		return errors.New("//错误的串口归属")
	case NET_ERR_WHITE_LIST_FULL:
		return errors.New("//白名单列表数量已满")
	case NET_ERR_NOT_MATCH_SOURCE:
		return errors.New("//不匹配的信号源类型")
	case NET_ERR_CLOCK_VIRTUAL_LED_FULL:
		return errors.New("//开启时钟功能的虚拟LED达上限")
	case NET_ERR_MAX_WIN_SIGNAL_LOOP_NUM:
		return errors.New("//窗口轮巡信号源个数已达上限")
	case NET_ERR_RESOLUTION_NO_MATCH_FRAME:
		return errors.New("//分辨率与当前帧数不匹配")
	case NET_ERR_NOT_UPDATE_LOW_VERSION:
		return errors.New("//不支持向低版本升级")
	case NET_ERR_NO_CUSTOM_TO_UPDATE:
		return errors.New("//非定制程序无法升级")
	case NET_ERR_CHAN_RESOLUTION_NOT_SUPPORT_SPLIT:
		return errors.New("//该输出口分辨率不支持分屏")
	case NET_ERR_HIGH_DEFINITION_NOT_SUPPORT_SPLIT:
		return errors.New("//超高清不支持9/16画面分割")
	case NET_ERR_MIRROR_IMAGE_BY_VIDEO_WALL:
		return errors.New("//电视墙镜像出错")
	case NET_ERR_MAX_OSD_FONT_SIZE:
		return errors.New("//超过OSD最大支持字符数")
	case NET_ERR_HIGH_DEFINITION_NOT_SUPPORT_VIDEO_SET:
		return errors.New("//超清不支持视频参数设置")
	case NET_ERR_TILE_MODE_NOT_SUPPORT_JOINT:
		return errors.New("//平铺模式不支持拼接窗口")
	case NET_ERR_ADD_AUDIO_MATRIX_FAILED:
		return errors.New("//创建音频矩阵失败")
	case NET_ERR_ONE_VIRTUAL_LED_AREA_BIND_ONE_AUDIO_AREA:
		return errors.New("//一个虚拟LED区域只能绑定一个音频区域")
	case NET_ERR_NAT_NOT_MODIFY_SERVER_NETWORK_PARAM:
		return errors.New("//NAT下无法修改服务器网络参数")
	case NET_ERR_ORIGINAL_CHECH_DATA_ERROR:
		return errors.New("//原始校正数据错误")
	case NET_ERR_INPUT_BOARD_SPLICED_IN_DIFFERENT_NETWORKAREAS:
		return errors.New("//不同网络区域的输入板不能拼接")
	case NET_ERR_SPLICINGSOURCE_ONWALL_IN_DIFFERENT_NETWORKAREAS:
		return errors.New("//不同网络区域的拼接源不能上墙")
	case NET_ERR_ONWALL_OUTPUTBOARD_MODIFY_NETWORKAREAS:
		return errors.New("//已绑定在电视墙上的输出板不能修改网络区域")
	case NET_ERR_LAN_AND_WAN_CANNOT_SAME_NET_SEGMENT:
		return errors.New("//LAN口IP和WAN口IP不能处于同一网段")
	case NET_ERR_USERNAME_REPETITIVE:
		return errors.New("//用户名重复")
	case NET_ERR_ASSOCIATED_SAMEWALL_IN_DIFFERENT_NETWORKAREAS:
		return errors.New("//不同数据网络区域的输出口不能关联到同一电视墙")
	case NET_ERR_BASEMAP_ROAM_IN_LED_AREA:
		return errors.New("//LED区域不允许底图漫游")
	case NET_ERR_VIRTUAL_LED_NOT_SUPPORT_4K_OUTPUT:
		return errors.New("//虚拟LED不支持4K输出口显示")
	case NET_ERR_BASEMAP_NOT_SUPPORT_4K_OUTPUT:
		return errors.New("//底图不支持4K输出口显示")
	case NET_ERR_MIN_BLOCK_IN_VIRTUAL_LED_AND_OUTPUT:
		return errors.New("//虚拟LED与输出口相交出现最小块")
	case NET_ERR_485FIlE_VERSION_INVALID:
		return errors.New("//485文件版本无效")
	case NET_ERR_485FIlE_CHECK_ERROR:
		return errors.New("//485文件校验错误")
	case NET_ERR_485FIlE_ABNORMAL_SIZE:
		return errors.New("//485文件大小异常效")
	case NET_ERR_MODIFY_SUBBOARD_NETCFG_IN_NAT:
		return errors.New("//NAT下无法修改子板网络参")
	case NET_ERR_OSD_CONTENT_WITH_ILLEGAL_CHARACTERS:
		return errors.New("//OSD内容包含非法字符")
	case NET_ERR_NON_SLAVE_DEVICE_INSERT_SYNC_LINE:
		return errors.New("//非从设备禁止插入同步线")
	//民用错误码（1100～1200）
	case NET_ERR_PLT_USERID:
		return errors.New("//验证平台userid错误")
	case NET_ERR_TRANS_CHAN_START:
		return errors.New("//透明通道已打开，当前操作无法完成")
	case NET_ERR_DEV_UPGRADING:
		return errors.New("//设备正在升级")
	case NET_ERR_MISMATCH_UPGRADE_PACK_TYPE:
		return errors.New("//升级包类型不匹配")
	case NET_ERR_DEV_FORMATTING:
		return errors.New("//设备正在格式化")
	case NET_ERR_MISMATCH_UPGRADE_PACK_VERSION:
		return errors.New("//升级包版本不匹配")
	case NET_ERR_PT_LOCKED:
		return errors.New("//PT锁定功能")

	case NET_DVR_LOGO_OVERLAY_WITHOUT_UPLOAD_PIC:
		return errors.New("//logo叠加失败，没有上传logo图片")
	case NET_DVR_ERR_ILLEGAL_VERIFICATION_CODE:
		return errors.New("//不合法验证码")
	case NET_DVR_ERR_LACK_VERIFICATION_CODE:
		return errors.New("//缺少验证码")
	case NET_DVR_ERR_FORBIDDEN_IP:
		return errors.New("//该IP地址已被禁止，不允许配置(设备支持的IP地址过滤功能)")
	case NET_DVR_ERR_UNLOCKPTZ:
		return errors.New("//操作无效，请先锁定云台")
	case NET_DVR_ERR_COUNTAREA_LARGE:
		return errors.New("//计数区域绘制错误，区域面积大于设备允许值")
	case NET_DVR_ERR_LABEL_ID_EXCEED:
		return errors.New("//标签ID超限")
	case NET_DVR_ERR_LABEL_TYPE:
		return errors.New("//标签类型错误")
	case NET_DVR_ERR_LABEL_FULL:
		return errors.New("//标签满")
	case NET_DVR_ERR_LABEL_DISABLED:
		return errors.New("//标签未使能")
	case NET_DVR_ERR_DOME_PT_TRANS_TO_DOME_XY:
		return errors.New("//球机PT转球机XY失败")
	case NET_DVR_ERR_DOME_PT_TRANS_TO_PANORAMA_XY:
		return errors.New("//球机PT转全景XY失败")
	case NET_DVR_ERR_PANORAMA_XY_TRANS_TO_DOME_PT:
		return errors.New("//全景XY坐标转球机PT错误")
	case NET_DVR_ERR_SCENE_DUR_TIME_LESS_THAN_INTERV_TIME:
		return errors.New("//场景停留时间要大于检测周期")
	case NET_DVR_ERR_HTTP_BKN_EXCEED_ONE:
		return errors.New("//断网续传布防只支持一路")
	case NET_DVR_ERR_DELETING_FAILED_TURN_OFF_HTTPS_ESDK_WEBSOCKETS_FIRST:
		return errors.New("//删除失败，请先关闭HTTPS和网络服务中的增强型SDK服务及WebSockets服务。")
	case NET_DVR_ERR_DELETING_FAILED_TURN_OFF_HTTPS_ESDK_FIRST:
		return errors.New("//删除失败，请先关闭HTTPS和网络服务中的增强型SDK服务")
	case NET_DVR_ERR_PTZ_OCCUPIED_PRIORITY:
		return errors.New("// 有高优先级云台控制权限用户操作")
	case NET_DVR_ERR_INCORRECT_VIDEOAUDIO_ID:
		return errors.New("// 视频通道编码ID或语音输出通道编码ID错误")
	case NET_DVR_ERR_REPETITIONTIME_OVER_MAXIMUM:
		return errors.New("// 去重时间最大不超过最大值")
	case NET_DVR_ERR_FORMATTING_FAILED:
		return errors.New("// 格式化错误，请重新")
	case NET_DVR_ERR_ENCRYPTED_FORMATTING_FAILED:
		return errors.New("// 加密格式化失败，请重试")
	case NET_DVR_ERR_WRONG_PASSWORD:
		return errors.New("// 密码错误,请输入正确的密码（SD卡 密码校验失败）")
	case NET_DVR_ERR_EXPOSURE_SYNC:
		return errors.New("// 镜头间曝光同步已开启，不允许配置手动RGB")

	//2012-10-16 报警设备错误码（1200~1300）
	case NET_ERR_SEARCHING_MODULE:
		return errors.New("// 正在搜索外接模块")
	case NET_ERR_REGISTERING_MODULE:
		return errors.New("// 正在注册外接模块")
	case NET_ERR_GETTING_ZONES:
		return errors.New("// 正在获取防区参数")
	case NET_ERR_GETTING_TRIGGERS:
		return errors.New("// 正在获取触发器")
	case NET_ERR_ARMED_STATUS:
		return errors.New("// 系统处于布防状态")
	case NET_ERR_PROGRAM_MODE_STATUS:
		return errors.New("// 系统处于编程模式")
	case NET_ERR_WALK_TEST_MODE_STATUS:
		return errors.New("// 系统处于步测模式")
	case NET_ERR_BYPASS_STATUS:
		return errors.New("// 旁路状态")
	case NET_ERR_DISABLED_MODULE_STATUS:
		return errors.New("// 功能未使能")
	case NET_ERR_NOT_SUPPORT_OPERATE_ZONE:
		return errors.New("// 防区不支持该操作")
	case NET_ERR_NOT_SUPPORT_MOD_MODULE_ADDR:
		return errors.New("// 模块地址不能被修改")
	case NET_ERR_UNREGISTERED_MODULE:
		return errors.New("// 模块未注册")
	case NET_ERR_PUBLIC_SUBSYSTEM_ASSOCIATE_SELF:
		return errors.New("// 公共子系统关联自身")
	case NET_ERR_EXCEEDS_ASSOCIATE_SUBSYSTEM_NUM:
		return errors.New("// 超过公共子系统最大关联个数")
	case NET_ERR_BE_ASSOCIATED_BY_PUBLIC_SUBSYSTEM:
		return errors.New("// 子系统被其他公共子系统关联")
	case NET_ERR_ZONE_FAULT_STATUS:
		return errors.New("// 防区处于故障状态")
	case NET_ERR_SAME_EVENT_TYPE:
		return errors.New("// 事件触发报警输出开启和事件触发报警输出关闭中有相同事件类型")
	case NET_ERR_ZONE_ALARM_STATUS:
		return errors.New("// 防区处于报警状态")
	case NET_ERR_EXPANSION_BUS_SHORT_CIRCUIT:
		return errors.New("//扩展总线短路")
	case NET_ERR_PWD_CONFLICT:
		return errors.New("//密码冲突")
	case NET_ERR_DETECTOR_GISTERED_BY_OTHER_ZONE:
		return errors.New("//探测器已被其他防区注册")
	case NET_ERR_DETECTOR_GISTERED_BY_OTHER_PU:
		return errors.New("//探测器已被其他主机注册")
	case NET_ERR_DETECTOR_DISCONNECT:
		return errors.New("//探测器不在线")
	case NET_ERR_CALL_BUSY:
		return errors.New("//设备正在通话中")
	case NET_DVR_ERR_ZONE_TAMPER_STAUS:
		return errors.New("//防区处于防拆状态")
	case NET_DVR_ERR_WIRELESS_DEV_REGISTER:
		return errors.New("//无线外设已被其他主机注册")
	case NET_DVR_ERR_WIRELESS_DEV_ADDED:
		return errors.New("//无线外设已被添加")
	case NET_DVR_ERR_WIRELESS_DEV_OFFLINE:
		return errors.New("//无线外设不在线")
	case NET_DVR_ERR_WIRELESS_DEV_TAMPER_STATUS:
		return errors.New("//无线外设处于防拆状态")
	case NET_DVR_ERR_GPRS_PHONE_CONFLICT:
		return errors.New("//电话报警与无线报警中心冲突")
	//信息发布主机
	case NET_ERR_GET_ALL_RETURN_OVER:
		return errors.New("//获取所有返回数目超限")
	case NET_ERR_RESOURCE_USING:
		return errors.New("//信息发布资源正在使用，不能修改")
	case NET_ERR_FILE_SIZE_OVERLIMIT:
		return errors.New("//文件大小超限")

	//信息发布服务器错误码
	case NET_ERR_MATERIAL_NAME:
		return errors.New("//素材名称非法")
	case NET_ERR_MATERIAL_NAME_LEN:
		return errors.New("//素材名称长度非法")
	case NET_ERR_MATERIAL_REMARK:
		return errors.New("//素材描述非法")
	case NET_ERR_MATERIAL_REMARK_LEN:
		return errors.New("//素材描述长度非法")
	case NET_ERR_MATERIAL_SHARE_PROPERTY:
		return errors.New("//素材共享属性非法")
	case NET_ERR_UNSUPPORT_MATERIAL_TYPE:
		return errors.New("//素材类型不支持")
	case NET_ERR_MATERIAL_NOT_EXIST:
		return errors.New("//素材不存在")
	case NET_ERR_READ_FROM_DISK:
		return errors.New("//从硬盘读取素材文件失败")
	case NET_ERR_WRITE_TO_DISK:
		return errors.New("//向硬盘写素材文件失败")
	case NET_ERR_WRITE_DATA_BASE:
		return errors.New("//素材写数据库失败")
	case NET_ERR_NO_APPROVED_NOT_EXPORT:
		return errors.New("//未审核内容禁止发布")
	case NET_ERR_MATERIAL_EXCEPTION:
		return errors.New("//素材异常")
	case NET_ERR_NO_MISINFO:
		return errors.New("//无误报信息")
	case NET_ERR_LAN_NOT_SUP_DHCP_CLIENT_CONFIGURATION:
		return errors.New("//网桥在路由模式下，配置DHCP客户端返回错误")
	case NET_ERR_VIDEOWALL_OPTPORT_RESOLUTION_INCONSISTENT:
		return errors.New("//电视墙上各输出口分辨率不一致(主要用于设置输出分辨率为4K出现异常时报错)")
	case NET_ERR_VIDEOWALL_OPTPORT_RESOLUTION_INCONSISTENT_UNBIND_OPTPORT_FIRST:
		return errors.New("//电视墙上各输出口分辨率不一致，请先解绑定输出口(主要用于绑定输出口出现异常时报错)")
	case NET_ERR_FOUR_K_OUTPUT_RESOLUTION_UNSUPPORT_NINE_TO_SIXTEEN_SPLIT_SCREEN:
		return errors.New("//4K输出分辨率不支持9/16分屏")
	case NET_ERR_SIGNAL_SOURCE_UNSUPPORT_CUSTOM_RESOLUTION:
		return errors.New("//信号源不支持该自定义分辨率")
	case NET_ERR_DVI_UNSUPPORT_FOURK_OUTPUT_RESOLUTION:
		return errors.New("//DVI不支持4K输出分辨率")
	case NET_ERR_BNC_UNSUPPORT_SOURCE_CROPPING:
		return errors.New("//BNC不支持信号源裁剪")

	//多屏互动错误码
	case NET_ERR_MAX_SCREEN_CTRL_NUM:
		return errors.New("//屏幕控制连接数达到上限")
	case NET_ERR_FILE_NOT_EXIST:
		return errors.New("//文件不存在")
	case NET_ERR_THUMBNAIL_NOT_EXIST:
		return errors.New("//缩略图不存在")
	case NET_ERR_DEV_OPEN_FILE_FAIL:
		return errors.New("//设备端打开文件失败")
	case NET_ERR_SERVER_READ_FILE_FAIL:
		return errors.New("//设备端读取文件失败")
	case NET_ERR_FILE_SIZE:
		return errors.New("//文件大小错误")
	case NET_ERR_FILE_NAME:
		return errors.New("//文件名称错误，为空或不合法")

	//分段错误码（1351-1400）
	case NET_ERR_BROADCAST_BUSY:
		return errors.New("//设备正在广播中")

	//2012-12-20抓拍机错误码（1400-1499）
	case NET_DVR_ERR_LANENUM_EXCEED:
		return errors.New("//车道数超出能力")
	case NET_DVR_ERR_PRAREA_EXCEED:
		return errors.New("//牌识区域过大")
	case NET_DVR_ERR_LIGHT_PARAM:
		return errors.New("//信号灯接入参数错误")
	case NET_DVR_ERR_LANE_LINE_INVALID:
		return errors.New("//车道线配置错误")
	case NET_DVR_ERR_STOP_LINE_INVALID:
		return errors.New("//停止线配置错误")
	case NET_DVR_ERR_LEFTORRIGHT_LINE_INVALID:
		return errors.New("//左/右转分界线配置错误")
	case NET_DVR_ERR_LANE_NO_REPEAT:
		return errors.New("//叠加车道号重复")
	case NET_DVR_ERR_PRAREA_INVALID:
		return errors.New("//牌识多边形不符合要求")
	case NET_DVR_ERR_LIGHT_NUM_EXCEED:
		return errors.New("//视频检测交通灯信号灯数目超出最大值")
	case NET_DVR_ERR_SUBLIGHT_NUM_INVALID:
		return errors.New("//视频检测交通灯信号灯子灯数目不合法")
	case NET_DVR_ERR_LIGHT_AREASIZE_INVALID:
		return errors.New("//视频检测交通灯输入信号灯框大小不合法")
	case NET_DVR_ERR_LIGHT_COLOR_INVALID:
		return errors.New("//视频检测交通灯输入信号灯颜色不合法")
	case NET_DVR_ERR_LIGHT_DIRECTION_INVALID:
		return errors.New("//视频检测交通灯输入灯方向属性不合法")
	case NET_DVR_ERR_LACK_IOABLITY:
		return errors.New("//IO口实际支持的能力不足")

	case NET_DVR_ERR_FTP_PORT:
		return errors.New("//FTP端口号非法（端口号重复或者异常）")
	case NET_DVR_ERR_FTP_CATALOGUE:
		return errors.New("//FTP目录名非法（启用多级目录，多级目录传值为空）")
	case NET_DVR_ERR_FTP_UPLOAD_TYPE:
		return errors.New("//FTP上传类型非法（单ftp只支持全部/双ftp只支持卡口和违章）")
	case NET_DVR_ERR_FLASH_PARAM_WRITE:
		return errors.New("//配置参数时写FLASH失败")
	case NET_DVR_ERR_FLASH_PARAM_READ:
		return errors.New("//配置参数时读FLASH失败")
	case NET_DVR_ERR_PICNAME_DELIMITER:
		return errors.New("//FTP图片命名分隔符非法")
	case NET_DVR_ERR_PICNAME_ITEM:
		return errors.New("//FTP图片命名项非法（例如 分隔符）")
	case NET_DVR_ERR_PLATE_RECOGNIZE_TYPE:
		return errors.New("//牌识区域类型非法 （矩形和多边形有效性校验）")
	case NET_DVR_ERR_CAPTURE_TIMES:
		return errors.New("//抓拍次数非法 （有效值是0～5）")
	case NET_DVR_ERR_LOOP_DISTANCE:
		return errors.New("//线圈距离非法 （有效值是0～2000ms）")
	case NET_DVR_ERR_LOOP_INPUT_STATUS:
		return errors.New("//线圈输入状态非法 （有效值）")
	case NET_DVR_ERR_RELATE_IO_CONFLICT:
		return errors.New("//测速组IO关联冲突")
	case NET_DVR_ERR_INTERVAL_TIME:
		return errors.New("//连拍间隔时间非法 （0～6000ms）")
	case NET_DVR_ERR_SIGN_SPEED:
		return errors.New("//标志限速值非法（大车标志限速不能大于小车标志限速 ）")
	case NET_DVR_ERR_PIC_FLIP:
		return errors.New("//图像配置翻转 （配置交互影响）")
	case NET_DVR_ERR_RELATE_LANE_NUMBER:
		return errors.New("//关联车道数错误 (重复 有效值校验1～99)")
	case NET_DVR_ERR_TRIGGER_MODE:
		return errors.New("//配置抓拍机触发模式非法")
	case NET_DVR_ERR_DELAY_TIME:
		return errors.New("//触发延时时间错误(2000ms)")
	case NET_DVR_ERR_EXCEED_RS485_COUNT:
		return errors.New("//超过最大485个数限制")
	case NET_DVR_ERR_RADAR_TYPE:
		return errors.New("//雷达类型错误")
	case NET_DVR_ERR_RADAR_ANGLE:
		return errors.New("//雷达角度错误")
	case NET_DVR_ERR_RADAR_SPEED_VALID_TIME:
		return errors.New("//雷达有效时间错误")
	case NET_DVR_ERR_RADAR_LINE_CORRECT:
		return errors.New("//雷达线性矫正参数错误")
	case NET_DVR_ERR_RADAR_CONST_CORRECT:
		return errors.New("//雷达常量矫正参数错误")
	case NET_DVR_ERR_RECORD_PARAM:
		return errors.New("//录像参数无效（预录时间不超过10s）")
	case NET_DVR_ERR_LIGHT_WITHOUT_COLOR_AND_DIRECTION:
		return errors.New("//视频检测信号灯配置信号灯个数，但是没有勾选信号灯方向和颜色的")
	case NET_DVR_ERR_LIGHT_WITHOUT_DETECTION_REGION:
		return errors.New("//视频检测信号灯配置信号灯个数，但是没有画检测区域")
	case NET_DVR_ERR_RECOGNIZE_PROVINCE_PARAM:
		return errors.New("//牌识参数省份参数的合法性")

	case NET_DVR_ERR_SPEED_TIMEOUT:
		return errors.New("//IO测速超时时间非法（有效值大于0）")
	case NET_DVR_ERR_NTP_TIMEZONE:
		return errors.New("//ntp时区参数错误")
	case NET_DVR_ERR_NTP_INTERVAL_TIME:
		return errors.New("//ntp校时间隔错误")
	case NET_DVR_ERR_NETWORK_CARD_NUM:
		return errors.New("//可配置网卡数目错误")
	case NET_DVR_ERR_DEFAULT_ROUTE:
		return errors.New("//默认路由错误")
	case NET_DVR_ERR_BONDING_WORK_MODE:
		return errors.New("//bonding网卡工作模式错误")
	case NET_DVR_ERR_SLAVE_CARD:
		return errors.New("//slave网卡错误")
	case NET_DVR_ERR_PRIMARY_CARD:
		return errors.New("//Primary网卡错误")
	case NET_DVR_ERR_DHCP_PPOE_WORK:
		return errors.New("//dhcp和pppoE不能同时启动")
	case NET_DVR_ERR_NET_INTERFACE:
		return errors.New("//网络接口错误")
	case NET_DVR_ERR_MTU:
		return errors.New("//MTU错误")
	case NET_DVR_ERR_NETMASK:
		return errors.New("//子网掩码错误")
	case NET_DVR_ERR_IP_INVALID:
		return errors.New("//IP地址不合法")
	case NET_DVR_ERR_MULTICAST_IP_INVALID:
		return errors.New("//多播地址不合法")
	case NET_DVR_ERR_GATEWAY_INVALID:
		return errors.New("//网关不合法")
	case NET_DVR_ERR_DNS_INVALID:
		return errors.New("//DNS不合法")
	case NET_DVR_ERR_ALARMHOST_IP_INVALID:
		return errors.New("//告警主机地址不合法")
	case NET_DVR_ERR_IP_CONFLICT:
		return errors.New("//IP冲突")
	case NET_DVR_ERR_NETWORK_SEGMENT:
		return errors.New("//IP不支持同网段")
	case NET_DVR_ERR_NETPORT:
		return errors.New("//端口错误")

	case NET_DVR_ERR_PPPOE_NOSUPPORT:
		return errors.New("//PPPOE不支持")
	case NET_DVR_ERR_DOMAINNAME_NOSUPPORT:
		return errors.New("//域名不支持")
	case NET_DVR_ERR_NO_SPEED:
		return errors.New("//未启用测速功能")
	case NET_DVR_ERR_IOSTATUS_INVALID:
		return errors.New("//IO状态错误")
	case NET_DVR_ERR_BURST_INTERVAL_INVALID:
		return errors.New("//连拍间隔非法")
	case NET_DVR_ERR_RESERVE_MODE:
		return errors.New("//备用模式错误")

	case NET_DVR_ERR_LANE_NO:
		return errors.New("//叠加车道号错误")
	case NET_DVR_ERR_COIL_AREA_TYPE:
		return errors.New("//线圈区域类型错误")
	case NET_DVR_ERR_TRIGGER_AREA_PARAM:
		return errors.New("//触发区域参数错误")
	case NET_DVR_ERR_SPEED_LIMIT_PARAM:
		return errors.New("//违章限速参数错误")
	case NET_DVR_ERR_LANE_PROTOCOL_TYPE:
		return errors.New("//车道关联协议类型错误")

	case NET_DVR_ERR_INTERVAL_TYPE:
		return errors.New("//连拍间隔类型非法")
	case NET_DVR_ERR_INTERVAL_DISTANCE:
		return errors.New("//连拍间隔距离非法")
	case NET_DVR_ERR_RS485_ASSOCIATE_DEVTYPE:
		return errors.New("//RS485关联类型非法")
	case NET_DVR_ERR_RS485_ASSOCIATE_LANENO:
		return errors.New("//RS485关联车道号非法")
	case NET_DVR_ERR_LANENO_ASSOCIATE_MULTIRS485:
		return errors.New("//车道号关联多个RS485口")
	case NET_DVR_ERR_LIGHT_DETECTION_REGION:
		return errors.New("//视频检测信号灯配置信号灯个数，但是检测区域宽或高为0")

	case NET_DVR_ERR_DN2D_NOSUPPORT:
		return errors.New("//不支持抓拍帧2D降噪")
	case NET_DVR_ERR_IRISMODE_NOSUPPORT:
		return errors.New("//不支持的镜头类型")
	case NET_DVR_ERR_WB_NOSUPPORT:
		return errors.New("//不支持的白平衡模式")
	case NET_DVR_ERR_IO_EFFECTIVENESS:
		return errors.New("//IO口的有效性")
	case NET_DVR_ERR_LIGHTNO_MAX:
		return errors.New("//信号灯检测器接入红/黄灯超限(16)")
	case NET_DVR_ERR_LIGHTNO_CONFLICT:
		return errors.New("//信号灯检测器接入红/黄灯冲突")

	case NET_DVR_ERR_CANCEL_LINE:
		return errors.New("//直行触发线")
	case NET_DVR_ERR_STOP_LINE:
		return errors.New("//待行区停止线")
	case NET_DVR_ERR_RUSH_REDLIGHT_LINE:
		return errors.New("//闯红灯触发线")
	case NET_DVR_ERR_IOOUTNO_MAX:
		return errors.New("//IO输出口编号越界")

	case NET_DVR_ERR_IOOUTNO_AHEADTIME_MAX:
		return errors.New("//IO输出口提前时间超限")
	case NET_DVR_ERR_IOOUTNO_IOWORKTIME:
		return errors.New("//IO输出口有效持续时间超限")
	case NET_DVR_ERR_IOOUTNO_FREQMULTI:
		return errors.New("//IO输出口脉冲模式下倍频出错")
	case NET_DVR_ERR_IOOUTNO_DUTYRATE:
		return errors.New("//IO输出口脉冲模式下占空比出错")
	case NET_DVR_ERR_VIDEO_WITH_EXPOSURE:
		return errors.New("//以曝闪起效，工作方式不支持视频")
	case NET_DVR_ERR_PLATE_BRIGHTNESS_WITHOUT_FLASHDET:
		return errors.New("//车牌亮度自动使能闪光灯仅在车牌亮度补偿模式下起效")

	case NET_DVR_ERR_RECOGNIZE_TYPE_PARAM:
		return errors.New("//识别类型非法 车牌识别参数（如大车、小车、背向、正向、车标识别等）")
	case NET_DVR_ERR_PALTE_RECOGNIZE_AREA_PARAM:
		return errors.New("//牌识参数非法 牌识区域配置时判断出错")
	case NET_DVR_ERR_PORT_CONFLICT:
		return errors.New("//端口有冲突")
	case NET_DVR_ERR_LOOP_IP:
		return errors.New("//IP不能设置为回环地址")
	case NET_DVR_ERR_DRIVELINE_SENSITIVE:
		return errors.New("//压线灵敏度出错(视频电警模式下)")

	//2013-3-6VQD错误码（1500～1550）
	case NET_ERR_VQD_TIME_CONFLICT:
		return errors.New("//VQD诊断时间段冲突")
	case NET_ERR_VQD_PLAN_NO_EXIST:
		return errors.New("//VQD诊断计划不存在")
	case NET_ERR_VQD_CHAN_NO_EXIST:
		return errors.New("//VQD监控点不存在")
	case NET_ERR_VQD_CHAN_MAX:
		return errors.New("//VQD计划数已达上限")
	case NET_ERR_VQD_TASK_MAX:
		return errors.New("//VQD任务数已达上限")

	case NET_SDK_GET_INPUTSTREAMCFG:
		return errors.New("//获取信号源")
	case NET_SDK_AUDIO_SWITCH_CONTROL:
		return errors.New("//子窗口音频开关控制")
	case NET_SDK_GET_VIDEOWALLDISPLAYNO:
		return errors.New("//获取设备显示输出号")
	case NET_SDK_GET_ALLSUBSYSTEM_BASIC_INFO:
		return errors.New("//获取所有子系统基本信息")
	case NET_SDK_SET_ALLSUBSYSTEM_BASIC_INFO:
		return errors.New("//设置所有子系统基本信息")
	case NET_SDK_GET_AUDIO_INFO:
		return errors.New("//获取所有音频信息")
	case NET_SDK_GET_MATRIX_STATUS_V50:
		return errors.New("// 获取视频综合平台状态_V50")
	case NET_SDK_DELETE_MONITOR_INFO:
		return errors.New("//删除Monitor信息")
	case NET_SDK_DELETE_CAMERA_INFO:
		return errors.New("//删除Camaera信息")

	//抓拍机错误码新增扩展(1600~1900)
	case NET_DVR_ERR_EXCEED_MAX_CAPTURE_TIMES:
		return errors.New("//抓拍模式为频闪时最大抓拍张数为2张(IVT模式下)")
	case NET_DVR_ERR_REDAR_TYPE_CONFLICT:
		return errors.New("//相同485口关联雷达类型冲突")
	case NET_DVR_ERR_LICENSE_PLATE_NULL:
		return errors.New("//车牌号为空")
	case NET_DVR_ERR_WRITE_DATABASE:
		return errors.New("//写入数据库失败")
	case NET_DVR_ERR_LICENSE_EFFECTIVE_TIME:
		return errors.New("//车牌有效时间错误")
	//视频电警
	case NET_DVR_ERR_PRERECORDED_STARTTIME_LONG:
		return errors.New("//预录开始时间大于违法抓拍张数")
	//混合卡口
	case NET_DVR_ERR_TRIGGER_RULE_LINE:
		return errors.New("//触发规则线错误")
	case NET_DVR_ERR_LEFTRIGHT_TRIGGERLINE_NOTVERTICAL:
		return errors.New("//左/右触发线不垂直")
	case NET_DVR_ERR_FLASH_LAMP_MODE:
		return errors.New("//闪光灯闪烁模式错误")
	case NET_DVR_ERR_ILLEGAL_SNAPSHOT_NUM:
		return errors.New("//违章抓拍张数错误")
	case NET_DVR_ERR_ILLEGAL_DETECTION_TYPE:
		return errors.New("//违章检测类型错误")
	case NET_DVR_ERR_POSITIVEBACK_TRIGGERLINE_HIGH:
		return errors.New("//正背向触发线高度错误")
	case NET_DVR_ERR_MIXEDMODE_CAPTYPE_ALLTARGETS:
		return errors.New("//混合模式下只支持机非人抓拍类型")

	case NET_DVR_ERR_CARSIGNSPEED_GREATERTHAN_LIMITSPEED:
		return errors.New("//小车标志限速大于限速值")
	case NET_DVR_ERR_BIGCARSIGNSPEED_GREATERTHAN_LIMITSPEED:
		return errors.New("//大车标志限速大于限速值")
	case NET_DVR_ERR_BIGCARSIGNSPEED_GREATERTHAN_CARSIGNSPEED:
		return errors.New("//大车标志限速大于小车标志限速值")
	case NET_DVR_ERR_BIGCARLIMITSPEED_GREATERTHAN_CARLIMITSPEED:
		return errors.New("//大车限速值大于小车限速值")
	case NET_DVR_ERR_BIGCARLOWSPEEDLIMIT_GREATERTHAN_CARLOWSPEEDLIMIT:
		return errors.New("//大车低速限速值大于小车低速限速值")
	case NET_DVR_ERR_CARLIMITSPEED_GREATERTHAN_EXCEPHIGHSPEED:
		return errors.New("//小车限速大于异常高速值")
	case NET_DVR_ERR_BIGCARLIMITSPEED_GREATERTHAN_EXCEPHIGHSPEED:
		return errors.New("//大车限速大于异常高速值")
	case NET_DVR_ERR_STOPLINE_MORETHAN_TRIGGERLINE:
		return errors.New("//停止线超过直行触发线")
	case NET_DVR_ERR_YELLOWLIGHTTIME_INVALID:
		return errors.New("// 视频检测黄灯持续时间不合法报错")
	case NET_DVR_ERR_TRIGGERLINE1_FOR_NOT_YIELD_TO_PEDESTRIAN_CANNOT_EXCEED_TRIGGERLINE2:
		return errors.New("//第一条不礼让行人触发线的位置超过了第二条不礼让行人触发线")
	case NET_DVR_ERR_TRIGGERLINE2_FOR_NOT_YIELD_TO_PEDESTRIAN_CANNOT_EXCEED_TRIGGERLINE1:
		return errors.New("//第二条不礼让行人触发线的位置超过了第一条不礼让行人触发线")

	//门禁主机错误码
	case NET_ERR_TIME_OVERLAP:
		return errors.New("//时间段重叠")
	case NET_ERR_HOLIDAY_PLAN_OVERLAP:
		return errors.New("//假日计划重叠")
	case NET_ERR_CARDNO_NOT_SORT:
		return errors.New("//卡号未排序")
	case NET_ERR_CARDNO_NOT_EXIST:
		return errors.New("//卡号不存在")
	case NET_ERR_ILLEGAL_CARDNO:
		return errors.New("//卡号错误")
	case NET_ERR_ZONE_ALARM:
		return errors.New("//防区处于布防状态(参数修改不允许)")
	case NET_ERR_ZONE_OPERATION_NOT_SUPPORT:
		return errors.New("//防区不支持该操作")
	case NET_ERR_INTERLOCK_ANTI_CONFLICT:
		return errors.New("//多门互锁和反潜回同时配置错误")
	case NET_ERR_DEVICE_CARD_FULL:
		return errors.New("//卡已满（卡达到10W后返回）")
	case NET_ERR_HOLIDAY_GROUP_DOWNLOAD:
		return errors.New("//假日组下载失败")
	case NET_ERR_LOCAL_CONTROL_OFF:
		return errors.New("//就地控制器离线")
	case NET_ERR_LOCAL_CONTROL_DISADD:
		return errors.New("//就地控制器未添加")
	case NET_ERR_LOCAL_CONTROL_HASADD:
		return errors.New("//就地控制器已添加")
	case NET_ERR_LOCAL_CONTROL_DOORNO_CONFLICT:
		return errors.New("//与已添加的就地控制器门编号冲突")
	case NET_ERR_LOCAL_CONTROL_COMMUNICATION_FAIL:
		return errors.New("//就地控制器通信失败")
	case NET_ERR_OPERAND_INEXISTENCE:
		return errors.New("//操作对象不存在（对门、报警输出、报警输入相关操作，当对象未添加时返回）")
	case NET_ERR_LOCAL_CONTROL_OVER_LIMIT:
		return errors.New("//就地控制器超出设备最大能力（主控对就地数量有限制）")
	case NET_ERR_DOOR_OVER_LIMIT:
		return errors.New("//门超出设备最大能力")
	case NET_ERR_ALARM_OVER_LIMIT:
		return errors.New("//报警输入输出超出设备最大能力")
	case NET_ERR_LOCAL_CONTROL_ADDRESS_INCONFORMITY_TYPE:
		return errors.New("//就地控制器地址与类型不符")
	case NET_ERR_NOT_SUPPORT_ONE_MORE_CARD:
		return errors.New("//不支持一人多卡")
	case NET_ERR_DELETE_NO_EXISTENCE_FACE:
		return errors.New("//删除的人脸不存在")
	case NET_ERR_DOOR_SPECIAL_PASSWORD_REPEAT:
		return errors.New("//与设备门特殊密码重复")
	case NET_ERR_AUTH_CODE_REPEAT:
		return errors.New("//与设备认证码重复")
	case NET_ERR_DEPLOY_EXCEED_MAX:
		return errors.New("//布防超过最大连接数")
	case NET_ERR_NOT_SUPPORT_DEL_FP_BY_ID:
		return errors.New("//读卡器不支持按手指ID删除指纹")
	case NET_ERR_TIME_RANGE:
		return errors.New("//有效期参数配置范围有误")
	case NET_ERR_CAPTURE_TIMEOUT:
		return errors.New("//采集超时")
	case NET_ERR_LOW_SCORE:
		return errors.New("//采集质量低")
	case NET_ERR_OFFLINE_CAPTURING:
		return errors.New("//离线采集中，无法响应")

	//可视对讲错误码
	case NET_DVR_ERR_OUTDOOR_COMMUNICATION:
		return errors.New("//与门口机通信异常")
	case NET_DVR_ERR_ROOMNO_UNDEFINED:
		return errors.New("//未设置房间号")
	case NET_DVR_ERR_NO_CALLING:
		return errors.New("//无呼叫")
	case NET_DVR_ERR_RINGING:
		return errors.New("//响铃")
	case NET_DVR_ERR_IS_CALLING_NOW:
		return errors.New("//正在通话")
	case NET_DVR_ERR_LOCK_PASSWORD_WRONG:
		return errors.New("//智能锁密码错误")
	case NET_DVR_ERR_CONTROL_LOCK_FAILURE:
		return errors.New("//开关锁失败")
	case NET_DVR_ERR_CONTROL_LOCK_OVERTIME:
		return errors.New("//开关锁超时")
	case NET_DVR_ERR_LOCK_DEVICE_BUSY:
		return errors.New("//智能锁设备繁忙")
	case NET_DVR_ERR_UNOPEN_REMOTE_LOCK_FUNCTION:
		return errors.New("//远程开锁功能未打开")

	//后端错误码 （2100 - 3000）
	case NET_DVR_ERR_FILE_NOT_COMPLETE:
		return errors.New("//下载的文件不完整")
	case NET_DVR_ERR_IPC_EXIST:
		return errors.New("//该IPC已经存在")
	case NET_DVR_ERR_ADD_IPC:
		return errors.New("//该通道已添加IPC")
	case NET_DVR_ERR_OUT_OF_RES:
		return errors.New("//网络带宽能力不足")
	case NET_DVR_ERR_CONFLICT_TO_LOCALIP:
		return errors.New("//IPC的ip地址跟DVR的ip地址冲突")
	case NET_DVR_ERR_IP_SET:
		return errors.New("//非法ip地址")
	case NET_DVR_ERR_PORT_SET:
		return errors.New("//非法的端口号")

	case NET_ERR_WAN_NOTSUPPORT:
		return errors.New("//不在同一个局域网，无法设置安全问题或导出GUID文件")
	case NET_ERR_MUTEX_FUNCTION:
		return errors.New("//功能互斥")
	case NET_ERR_QUESTION_CONFIGNUM:
		return errors.New("//安全问题配置数量错误")
	case NET_ERR_FACECHAN_NORESOURCE:
		return errors.New("//人脸智能通道资源已用完")
	case NET_ERR_DATA_CALLBACK:
		return errors.New("//正在数据回迁")
	case NET_ERR_ATM_VCA_CHAN_IS_RELATED:
		return errors.New("//ATM智能通道已被关联")
	case NET_ERR_ATM_VCA_CHAN_IS_OVERLAPED:
		return errors.New("//ATM智能通道已被叠加")
	case NET_ERR_FACE_CHAN_UNOVERLAP_EACH_OTHER:
		return errors.New("//人脸通道不能互相叠加")
	case NET_ERR_ACHIEVE_MAX_CHANNLE_LIMIT:
		return errors.New("//达到最大路数限制")
	case NET_DVR_SMD_ENCODING_NORESOURSE:
		return errors.New("//SMD编码资源不足")
	case NET_DVR_SMD_DECODING_NORESOURSE:
		return errors.New("//SMD解码资源不足")
	case NET_DVR_FACELIB_DATA_PROCESSING:
		return errors.New("//人脸库数据正在处理")
	case NET_DVR_ERR_LARGE_TIME_DIFFRENCE:
		return errors.New("//设备和服务器之间的时间差异太大")
	case NET_DVR_NO_SUPPORT_WITH_PLAYBACK:
		return errors.New("//已开启回放，不支持本功能")
	case NET_DVR_CHANNEL_NO_SUPPORT_WITH_SMD:
		return errors.New("//通道已开启SMD，不支持本功能")
	case NET_DVR_CHANNEL_NO_SUPPORT_WITH_FD:
		return errors.New("//通道已开启人脸抓拍，不支持本功能")
	case NET_DVR_ILLEGAL_PHONE_NUMBER:
		return errors.New("//非法的电话号码")
	case NET_DVR_ILLEGAL_CERITIFICATE_NUMBER:
		return errors.New("//非法的证件号码")
	case NET_DVR_ERR_CHANNEL_RESOLUTION_NO_SUPPORT:
		return errors.New("//通道分辨率不支持")
	case NET_DVR_ERR_CHANNEL_COMPRESSION_NO_SUPPORT:
		return errors.New("//通道编码格式不支持")

	case NET_DVR_ERR_CLUSTER_DEVICE_TOO_LESS:
		return errors.New("//设备数少，不允许删除")
	case NET_DVR_ERR_CLUSTER_DEL_DEVICE_CM_PLAYLOAD:
		return errors.New("//该设备是集群主机，不允许删除")
	case NET_DVR_ERR_CLUSTER_DEVNUM_OVER_UPPER_LIMIT:
		return errors.New("//设备数达到上限，不允许增加")
	case NET_DVR_ERR_CLUSTER_DEVICE_TYPE_INCONFORMITY:
		return errors.New("//设备类型不一致")
	case NET_DVR_ERR_CLUSTER_DEVICE_VERSION_INCONFORMITY:
		return errors.New("//设备版本不一致")
	case NET_DVR_ERR_CLUSTER_IP_CONFLICT:
		return errors.New("//集群系统IP地址冲突：ipv4地址冲突、ipv6地址冲突")
	case NET_DVR_ERR_CLUSTER_IP_INVALID:
		return errors.New("//集群系统IP地址无效：ipv4非法、ipv6非法")
	case NET_DVR_ERR_CLUSTER_PORT_CONFLICT:
		return errors.New("//集群系统端口冲突")
	case NET_DVR_ERR_CLUSTER_PORT_INVALID:
		return errors.New("//集群系统端口非法")
	case NET_DVR_ERR_CLUSTER_USERNAEM_OR_PASSWORD_INVALID:
		return errors.New("//组网用户名或密码非法")
	case NET_DVR_ERR_CLUSTER_DEVICE_ALREADY_EXIST:
		return errors.New("//存在相同设备")
	case NET_DVR_ERR_CLUSTER_DEVICE_NOT_EXIST:
		return errors.New("//设备不存在(组网时下发的cs列表中的设备信息在任何一台cs上都找不到，删除的时候下发的id不对)")
	case NET_DVR_ERR_CLUSTER_NON_CLUSTER_MODE:
		return errors.New("//设备处于非集群模式")
	case NET_DVR_ERR_CLUSTER_IP_NOT_SAME_LAN:
		return errors.New("//IP地址不在同一局域网，不同区域网不允许组网/扩容")

	case NET_DVR_ERR_CAPTURE_PACKAGE_FAILED:
		return errors.New("//抓包失败")
	case NET_DVR_ERR_CAPTURE_PACKAGE_PROCESSING:
		return errors.New("//正在抓包")
	case NET_DVR_ERR_SAFETY_HELMET_NO_RESOURCE:
		return errors.New("//安全帽检测资源不足")
	case NET_DVR_NO_SUPPORT_WITH_ABSTRACT:
		return errors.New("//已开启视频摘要，不支持本功能")
	case NET_DVR_ERR_TAPE_LIB_NEED_STOP_ARCHIVE:
		return errors.New("//磁带库归档需要停止")
	case NET_DVR_INSUFFICIENT_DEEP_LEARNING_RESOURCES:
		return errors.New("//深度学习资源超限")
	case NET_DVR_ERR_IDENTITY_KEY:
		return errors.New("//交互口令错误")
	case NET_DVR_MISSING_IDENTITY_KEY:
		return errors.New("//交互口令缺失")
	case NET_DVR_NO_SUPPORT_WITH_PERSON_DENSITY_DETECT:
		return errors.New("//已开启人员密度检测，不支持本功能")
	case NET_DVR_IPC_RESOLUTION_OVERFLOW:
		return errors.New("//IPC分辨率超限")
	case NET_DVR_IPC_BITRATE_OVERFLOW:
		return errors.New("//IPC码率超限")
	case NET_DVR_ERR_INVALID_TASKID:
		return errors.New("//无效的taskID")
	case NET_DVR_PANEL_MODE_NOT_CONFIG:
		return errors.New("//没有配置面板路智能")
	case NET_DVR_NO_HUMAN_ENGINES_RESOURCE:
		return errors.New("//人体引擎资源不足")
	case NET_DVR_ERR_TASK_NUMBER_OVERFLOW:
		return errors.New("//任务数据超过上限")
	case NET_DVR_ERR_COLLISION_TIME_OVERFLOW:
		return errors.New("//碰撞时间超过上限")
	case NET_DVR_ERR_CAPTURE_PACKAGE_NO_USB:
		return errors.New("//未识别到U盘，请插入U盘或重新插入")
	case NET_DVR_ERR_NO_SET_SECURITY_EMAIL:
		return errors.New("//未设置安全邮箱")
	case NET_DVR_ERR_EVENT_NOTSUPPORT:
		return errors.New("//订阅事件不支持")
	case NET_DVR_ERR_PASSWORD_FORMAT:
		return errors.New("//密码格式不对")
	case NET_DVR_ACCESS_FRONT_DEVICE_PARAM_FAILURE:
		return errors.New("//获取前端设备参数失败")
	case NET_DVR_ACCESS_FRONT_DEVICE_STREAM_FAILURE:
		return errors.New("//对前端设备取流失败")
	case NET_DVR_ERR_USERNAME_FORMAT:
		return errors.New("//用户名格式不对")
	case NET_DVR_ERR_UNOPENED_HIGH_RESOLUTION_MODE:
		return errors.New("//超高分辨率模式未开启")
	case NET_DVR_ERR_TOO_SMALL_QUATO:
		return errors.New("//配额设置太小")
	case NET_DVR_ERR_EMAIL_FORMAT:
		return errors.New("//邮箱格式不对")
	case NET_DVR_ERR_SECURITY_CODE_FORMAT:
		return errors.New("//安全码格式不对")
	case NET_DVR_PD_SPACE_TOO_SMALL:
		return errors.New("//阵列硬盘容量太小")
	case NET_DVR_PD_NUM_TOO_BIG:
		return errors.New("//阵列硬盘总数超过总盘数的二分之一")
	case NET_DVR_ERR_USB_IS_FULL:
		return errors.New("//U盘已满")
	case NET_DVR_EXCEED_MAX_SMD_TYPE:
		return errors.New("//达到最大SMD事件种类上限")
	case NET_DVR_CHANNEL_NO_SUPPORT_WITH_BEHAVIOR:
		return errors.New("//通道已开启行为分析，不支持本功能")
	case NET_DVR_NO_BEHAVIOR_ENGINES_RESOURCE:
		return errors.New("//行为分析资源不足")
	case NET_DVR_NO_RETENTION_ENGINES_RESOURCE:
		return errors.New("//人员滞留检测资源不足")
	case NET_DVR_NO_LEAVE_POSITION_ENGINES_RESOURCE:
		return errors.New("//离岗检测资源不足")
	case NET_DVR_NO_PEOPLE_NUM_CHANGE_ENGINES_RESOURCE:
		return errors.New("//人数异常资源不足")
	case NET_DVR_PANEL_MODE_NUM_OVER_LIMIT:
		return errors.New("//超过面板路最大路数")
	case NET_DVR_SURROUND_MODE_NUM_OVER_LIMIT:
		return errors.New("//超过环境路最大路数")
	case NET_DVR_FACE_MODE_NUM_OVER_LIMIT:
		return errors.New("//超过人脸路最大路数")
	case NET_DVR_SAFETYCABIN_MODE_NUM_OVER_LIMIT:
		return errors.New("//超过防护舱路最大路数")
	case NET_DVR_DETECT_REGION_RANGE_INVALID:
		return errors.New("//检测区域范围非法")
	case NET_DVR_CHANNEL_CAPTURE_PICTURE_FAILURE:
		return errors.New("//通道抓图失败")
	case NET_DVR_VCACHAN_IS_NORESOURCE:
		return errors.New("//智能通道资源用完")
	case NET_DVR_IPC_NUM_REACHES_LIMIT:
		return errors.New("// Ipc通道数目达到上限")
	case NET_DVR_IOT_NUM_REACHES_LIMIT:
		return errors.New("// IOT通道数目达到上限")
	case NET_DVR_IOT_CHANNEL_DEVICE_EXIST:
		return errors.New("//当前IOT通道已经添加设备")
	case NET_DVR_IOT_CHANNEL_DEVICE_NOT_EXIST:
		return errors.New("//当前IOT通道不存在设备")
	case NET_DVR_INVALID_IOT_PROTOCOL_TYPE:
		return errors.New("//非法的IOT协议类型")
	case NET_DVR_INVALID_EZVIZ_SECRET_KEY:
		return errors.New("//非法的萤石注册验证码")
	case NET_DVR_DUPLICATE_IOT_DEVICE:
		return errors.New("//重复的IOT设备")
	case NET_DVR_SADP_MODIFY_FALIURE:
		return errors.New("// SADP修改失败")
	case NET_DVR_IPC_NETWORK_ABNORMAL:
		return errors.New("// IPC网络异常")
	case NET_DVR_IPC_PASSWORD_ERROR:
		return errors.New("// IPC用户名密码错误")
	case NET_DVR_ERROR_IPC_TYPE:
		return errors.New("//IPC类型不对")
	case NET_DVR_ERROR_IPC_LIST_NOT_EMPTY:
		return errors.New("//已添加IPC列表不为空，不支持一键配置")
	case NET_DVR_ERROR_IPC_LIST_NOT_MATCH_PAIRING:
		return errors.New("//IPC列表和配单不匹配")
	case NET_DVR_ERROR_IPC_BAD_LANGUAGE:
		return errors.New("//IPC语言和设备不匹配")
	case NET_DVR_ERROR_IPC_IS_LOCKING:
		return errors.New("//IPC已被锁")
	case NET_DVR_ERROR_IPC_NOT_ACTIVATED:
		return errors.New("//IPC未激活")
	case NET_DVR_FIELD_CODING_NOT_SUPPORT:
		return errors.New("//场编码不支持")
	case NET_DVR_ERROR_H323_NOT_SUPPORT_H265:
		return errors.New("//H323视频会议就不支持H265码流")
	case NET_DVR_ERROR_EXPOSURE_TIME_TOO_BIG_IN_MODE_P:
		return errors.New("//P制式下，曝光时间过大")
	case NET_DVR_ERROR_EXPOSURE_TIME_TOO_BIG_IN_MODE_N:
		return errors.New("//N制式下，曝光时间过大")
	case NET_DVR_ERROR_PING_PROCESSING:
		return errors.New("//正在PING")
	case NET_DVR_ERROR_PING_NOT_START:
		return errors.New("//Ping功能未开始")
	case NET_DVR_ERROR_NEED_DOUBLE_VERIFICATION:
		return errors.New("//需要二次认证")
	case NET_DVR_NO_DOUBLE_VERIFICATION_USER:
		return errors.New("//无二次认证用户")
	case NET_DVR_CHANNEL_OFFLINE:
		return errors.New("//通道离线")
	case NET_DVR_TIMESPAN_NUM_OVER_LIMIT:
		return errors.New("//时间段超出支持最大数目")
	case NET_DVR_CHANNEL_NUM_OVER_LIMIT:
		return errors.New("//通道数目超出支持最大数目")
	case NET_DVR_NO_SEARCH_ID_RESOURCE:
		return errors.New("//分页查询的searchID资源不足")
	case NET_DVR_ERROR_ONEKEY_EXPORT:
		return errors.New("//正在进行导出操作，请稍后再试")
	case NET_DVR_NO_CITY_MANAGEMENT_ENGINES_RESOURCE:
		return errors.New("//城管算法引擎资源不足")
	case NET_DVR_NO_SITUATION_ANALYSIS_ENGINES_RESOURCE:
		return errors.New("//态势分析引擎资源不足")
	case NET_DVR_INTELLIGENT_ANALYSIS_IPC_CANNT_DELETE:
		return errors.New("//正在进行智能分析的IPC无法删除")
	case NET_DVR_NOSUPPORT_RESET_PASSWORD:
		return errors.New("//NVR不支持对IPC重置密码")
	case NET_DVR_ERROR_IPC_NEED_ON_LAN:
		return errors.New("// IPC需要在局域网内")
	case NET_DVR_CHANNEL_NO_SUPPORT_WITH_SAFETY_HELMET:
		return errors.New("//通道已开启安全帽检测，不支持本功能")
	case NET_DVR_ERROR_GET_RESETPASSWORDTYPE_IS_ABNORMAL:
		return errors.New("// IPC重置密码时，获取IPC的重置密码类型异常")
	case NET_DVR_ERROR_IPC_NOSUPPORT_RESET_PASSWORD:
		return errors.New("//  IPC不支持重置密码")
	case NET_DVR_ERROR_IP_IS_NOT_ONLY_ONE:
		return errors.New("// IPC的IP不唯一，有重复")
	case NET_DVR_NO_SUPPORT_WITH_SMD_OR_SCD:
		return errors.New("//已开启SMD/SCD，不支持本功能（SCD为场景变更）")
	case NET_DVR_NO_SUPPORT_WITH_FD:
		return errors.New("//已开启人脸抓拍，不支持本功能")
	case NET_DVR_NO_FD_ENGINES_RESOURCE:
		return errors.New("//人脸抓拍资源不足")
	case NET_DVR_ERROR_ONEKEY_REMOVE:
		return errors.New("//正在进行删除操作，请稍后再试")
	case NET_DVR_FACE_PIP_BACKGROUND_CHANNEL_OVERFLOW:
		return errors.New("//人脸画中画背景通道超限")
	case NET_DVR_MICIN_CHANNEL_OCCUPIED:
		return errors.New("//micin通道被占用")
	case NET_DVR_IPC_CHANNEL_IS_IN_PIP:
		return errors.New("//操作失败，该通道已关联到审讯通道，请先取消画中画配置关联")
	case NET_DVR_CHANNEL_NO_SUPPORT_WITH_FACE_CONTRAST:
		return errors.New("//通道已开启人脸比对，不支持本功能")
	case NET_DVR_INVALID_RECHARGE_CARD:
		return errors.New("//无效的充值卡")
	case NET_DVR_CLOUD_PLATFORM_SERVER_EXCEPTION:
		return errors.New("//云平台服务器异常")
	case NET_DVR_OPERATION_FAILURE_WITHOUT_LOGIN:
		return errors.New("//未登录操作失败")
	case NET_DVR_INVALID_ASSOCIATED_SERIAL_NUMBER:
		return errors.New("//关联序列号非法")
	case NET_DVR_CLOUD_PLATFORM_ACCOUNT_NOT_EXIST:
		return errors.New("//云平台帐号不存在")
	case NET_DVR_DEVICE_SERIAL_NUMBER_REGISTERED:
		return errors.New("//设备序列号已注册")
	case NET_DVR_CONFERENCE_ROOM_NOT_EXIST:
		return errors.New("//会议室不存在")
	case NET_DVR_NEED_DISABLED_ANALOG_CHANNEL:
		return errors.New("//需禁用模拟通道")
	case NET_DVR_STUDENT_ROLL_CALL_FAILURE:
		return errors.New("//学生点名失败")
	case NET_DVR_SUB_DEVICE_NOT_ENABLE_INDIVIDUAL_BEHAVIOR:
		return errors.New("//子设备未启用个体行为模式")
	case NET_DVR_SUB_DEVICE_CHANNEL_CONTROL_FAILED:
		return errors.New("//子设备通道控制失败")
	case NET_DVR_DEVICE_NOT_IN_CONFERENCE:
		return errors.New("//设备不在会议中")
	case NET_DVR_ALREADY_EXIST_CONFERENCE:
		return errors.New("//当前已经存在会议")
	case NET_DVR_NO_SUPPORT_WITH_VIDEO_CONFERENCE:
		return errors.New("//当前正在视频会议中，不支持本功能")
	case NET_DVR_START_INTERACTION_FAILURE:
		return errors.New("//互动开始失败")
	case NET_DVR_ASK_QUESTION_STARTED:
		return errors.New("//已开始提问")
	case NET_DVR_ASK_QUESTION_CLOSED:
		return errors.New("//已结束提问")
	case NET_DVR_UNABLE_OPERATE_BY_HOST:
		return errors.New("//已被主持人禁用，无法操作")
	case NET_DVR_REPEATED_ASK_QUESTION:
		return errors.New("//重复提问")
	case NET_DVR_SWITCH_TIMEDIFF_LESS_LIMIT:
		return errors.New("// 开关机时间差小于限制值(10分钟)")
	case NET_DVR_CHANNEL_DEVICE_EXIST:
		return errors.New("//当前通道已经添加设备")
	case NET_DVR_CHANNEL_DEVICE_NOT_EXIST:
		return errors.New("//当前通道不存在设备")
	case NET_DVR_ERROR_ADJUSTING_RESOLUTION:
		return errors.New("//先关闭摄像机的裁剪，再调整分辨率")
	case NET_DVR_SSD_FILE_SYSTEM_IS_UPGRADING:
		return errors.New("//SSD文件系统正在升级")
	case NET_DVR_SSD_FILE_SYSTEM_IS_FORMAT:
		return errors.New("//SSD正在格式化")
	case NET_DVR_CHANNEL_IS_CONNECTING:
		return errors.New("//当前通道正在连接")

	case NET_DVR_CHANNEL_STREAM_TYPE_NOT_SUPPORT:
		return errors.New("//当前通道码流类型不支持")
	case NET_DVR_CHANNEL_USERNAME_NOT_EXIST:
		return errors.New("//当前通道用户名不存在")
	case NET_DVR_CHANNEL_ACCESS_PARAM_FAILURE:
		return errors.New("//当前通道获取参数失败")
	case NET_DVR_CHANNEL_GET_STREAM_FAILURE:
		return errors.New("//当前通道取流失败")
	case NET_DVR_CHANNEL_RISK_PASSWORD:
		return errors.New("//当前通道密码为风险密码")
	case NET_DVR_NO_SUPPORT_DELETE_STRANGER_LIB:
		return errors.New("//不支持删除陌生人库")
	case NET_DVR_NO_SUPPORT_CREATE_STRANGER_LIB:
		return errors.New("//不支持创建陌生人库")
	case NET_DVR_NETWORK_PORT_CONFLICT:
		return errors.New("//网络配置端口冲突")
	case NET_DVR_TRANSCODE_NO_RESOURCES:
		return errors.New("//转码资源不足")
	case NET_DVR_SSD_FILE_SYSTEM_ERROR:
		return errors.New("//SSD文件系统错误")
	case NET_DVR_INSUFFICIENT_SSD__FOR_FPD:
		return errors.New("//用于人员频次业务的SSD容量不够")
	case NET_DVR_ASSOCIATED_FACELIB_OVER_LIMIT:
		return errors.New("//关联人脸库已达上限")
	case NET_DVR_NEED_DELETE_DIGITAL_CHANNEL:
		return errors.New("//需删除数字通道")
	//热成像产线相关错误码（3001 - 3500）
	case NET_DVR_ERR_NOTSUPPORT_DEICING:
		return errors.New("//设备当前状态不支持除冰功能（只在POE+、AC24V、DC12V供电下支持除冰功能）")
	case NET_DVR_ERR_THERMENABLE_CLOSE:
		return errors.New("//测温功能总使能未开启。(即NET_DVR_THERMOMETRY_BASICPARAM使能未开启)")
	case NET_DVR_ERR_NOTMEET_DEICING:
		return errors.New("//当前空腔温度不满足手动除冰开启条件（需空腔温度小于30度才可开启）")
	case NET_DVR_ERR_PANORAMIC_LIMIT_OPERATED:
		return errors.New("//全景地图和限位不可同时操作")
	case NET_DVR_ERR_SMARTH264_ROI_OPERATED:
		return errors.New("//SmartH264和ROI不可同时操作")
	case NET_DVR_ERR_RULENUM_LIMIT:
		return errors.New("//规则数上限")
	case NET_DVR_ERR_LASER_DEICING_OPERATED:
		return errors.New("//激光以及除冰功能不可同时操作")
	case NET_DVR_ERR_OFFDIGITALZOOM_OR_MINZOOMLIMIT:
		return errors.New("//请先关闭数据变倍功能或变倍限制设置为最小值。当执行烟火检测、行为分析、船只检测、坏点矫正、测温、烟火屏蔽功能时，若没有关闭数据变倍或者变倍限制没有设置为最小值时，将会提示该错误码。")
	case NET_DVR_ERR_FIREWAITING:
		return errors.New("//设备云台正在火点等待中")
	case NET_DVR_SYNCHRONIZEFOV_ERROR:
		return errors.New("//同步视场角错误")
	case NET_DVR_CERTIFICATE_VALIDATION_ERROR:
		return errors.New("//证书验证不通过")
	case NET_DVR_CERTIFICATES_NUM_EXCEED_ERROR:
		return errors.New("//证书个数超过上限")
	case NET_DVR_RULE_SHIELDMASK_CONFLICT_ERROR:
		return errors.New("//规则区域与屏蔽区域冲突")

	//前端产品线错误码（3501-4000）
	case NET_DVR_ERR_NO_SAFETY_HELMET_REGION:
		return errors.New("//未配置安全帽检测区域")
	case NET_DVR_ERR_UNCLOSED_SAFETY_HELMET:
		return errors.New("//未关闭安全帽检测使能")
	case NET_DVR_ERR_MUX_RECV_STATE:
		return errors.New("//多路复用接收状态异常")
	case NET_DVR_UPLOAD_HBDLIBID_ERROR:
		return errors.New("// 上传人体库ID（HBDID or customHBDID）错误")

	case NET_ERR_NPQ_BASE_INDEX:
		return errors.New("//NPQ库错误码")
	case NET_ERR_NPQ_PARAM:
		return errors.New("//NPQ库参数有误")
	case NET_ERR_NPQ_SYSTEM:
		return errors.New("//NPQ库操作系统调用错误(包括资源申请失败或内部错误等）")
	case NET_ERR_NPQ_GENRAL:
		return errors.New("//NPQ库内部通用错误")
	case NET_ERR_NPQ_PRECONDITION:
		return errors.New("//NPQ库调用顺序错误")
	case NET_ERR_NPQ_NOTSUPPORT:
		return errors.New("//NPQ库功能不支持")

	case NET_ERR_NPQ_NOTCALLBACK:
		return errors.New("//数据没有回调上来")
	case NET_ERR_NPQ_LOADLIB:
		return errors.New("//NPQ库加载失败")
	case NET_ERR_NPQ_STEAM_CLOSE:
		return errors.New("//本路码流NPQ功能未开启")
	case NET_ERR_NPQ_MAX_LINK:
		return errors.New("//NPQ取流路数达到上限")
	case NET_ERR_NPQ_STREAM_CFG:
		return errors.New("//编码参数存在冲突配置")

	case NET_EZVIZ_P2P_BASE_INDEX:
		return errors.New("萤石P2P组件错误码")

	case NET_DVR_EZVIZ_P2P_COULDNT_RESOLVE_HOST:
		return errors.New("P2PCLOUD_ER_COULDNT_RESOLVE_HOST    1006")
	case NET_DVR_EZVIZ_P2P_COULDNT_CONNECT:
		return errors.New("P2PCLOUD_ER_COULDNT_CONNECT 1007")
	case NET_DVR_EZVIZ_P2P_OPERATION_TIMEOUT:
		return errors.New("P2PCLOUD_ER_OPERATION_TIMEOUT   1028")
	case NET_DVR_EZVIZ_P2P_NOT_INITIALIZED:
		return errors.New("P2PCLOUD_ER_NOT_INITIALIZED 2001")
	case NET_DVR_EZVIZ_P2P_INVALID_ARG:
		return errors.New("P2PCLOUD_ER_INVALID_ARG 2002")
	case NET_DVR_EZVIZ_P2P_EXCEED_MAX_SERVICE:
		return errors.New("P2PCLOUD_ER_EXCEED_MAX_SERVICE  2003")
	case NET_DVR_EZVIZ_P2P_TOKEN_NOT_EXIST:
		return errors.New("P2PCLOUD_ER_TOKEN_NOT_EXIST 2004")
	case NET_DVR_EZVIZ_P2P_DISCONNECTED:
		return errors.New("P2PCLOUD_ER_DISCONNECTED    2005")
	case NET_DVR_EZVIZ_P2P_RELAY_ADDR_NOT_EXIST:
		return errors.New("P2PCLOUD_ER_RELAY_ADDR_NOT_EXIST    2006")
	case NET_DVR_EZVIZ_P2P_DEV_NOT_ONLINE:
		return errors.New("P2PCLOUD_ER_DEV_NOT_ONLINE  3121")
	case NET_DVR_EZVIZ_P2P_DEV_CONNECT_EXCEED:
		return errors.New("P2PCLOUD_ER_DEV_CONNECT_EXCEED  3123")
	case NET_DVR_EZVIZ_P2P_DEV_CONNECT_FAILED:
		return errors.New("P2PCLOUD_ER_DEV_CONNECT_FAILED  3209")
	case NET_DVR_EZVIZ_P2P_DEV_RECV_TIMEOUT:
		return errors.New("P2PCLOUD_ER_DEV_RECV_TIMEOUT    3213")
	case NET_DVR_EZVIZ_P2P_USER_FORCE_STOP:
		return errors.New("P2PCLOUD_ER_USER_FORCE_STOP 3216")
	case NET_DVR_EZVIZ_P2P_NO_PERMISSION:
		return errors.New("P2PCLOUD_ER_NO_PERMISSION   3255")
	case NET_DVR_EZVIZ_P2P_DEV_PU_NOT_FOUND:
		return errors.New("P2PCLOUD_ER_DEV_PU_NOT_FOUND    3297")
	case NET_DVR_EZVIZ_P2P_DEV_CONN_NOLONGER_AVAIL:
		return errors.New("P2PCLOUD_ER_DEV_CONN_NOLONGER_AVAIL 3351")
	case NET_DVR_EZVIZ_P2P_DEV_NOT_LISTENING:
		return errors.New("P2PCLOUD_ER_DEV_NOT_LISTENING   3610")
	case NET_DVR_EZVIZ_P2P_DEV_TUNNEL_SOCKET_LIMITED:
		return errors.New("P2PCLOUD_ER_DEV_TUNNEL_SOCKET_LIMITED   3612")
	case NET_DVR_EZVIZ_P2P_FAIL_CREATE_THREAD:
		return errors.New("TUNNEL_ER_FAIL_CREATE_THREAD    4001")
	case NET_DVR_EZVIZ_P2P_FAIL_ALLOC_BUFFERS:
		return errors.New("P2PCLOUD_ER_FAIL_ALLOC_BUFFERS  4002")
	case NET_DVR_EZVIZ_P2P_FAIL_CREATE_SOCKET:
		return errors.New("P2PCLOUD_ER_FAIL_CREATE_SOCKET  4003")
	case NET_DVR_EZVIZ_P2P_BIND_LOCAL_SERVICE:
		return errors.New("P2PCLOUD_ER_BIND_LOCAL_SERVICE  4004")
	case NET_DVR_EZVIZ_P2P_LISTEN_LOCAL_SERVICE:
		return errors.New("P2PCLOUD_ER_LISTEN_LOCAL_SERVICE    4005")
	case NET_DVR_EZVIZ_P2P_SVR_RSP_BAD:
		return errors.New("P2PCLOUD_ER_SVR_RSP_BAD 5001")
	case NET_DVR_EZVIZ_P2P_SVR_RSP_INVALID:
		return errors.New("P2PCLOUD_ER_SVR_RSP_INVALID 5002")
	case NET_DVR_EZVIZ_P2P_SVR_LOGIN_FAILED:
		return errors.New("P2PCLOUD_ER_SVR_LOGIN_FAILED    5003")
	case NET_DVR_EZVIZ_P2P_SVR_TOKEN_EXPIRED:
		return errors.New("P2PCLOUD_ER_SVR_TOKEN_EXPIRED   5004")
	case NET_DVR_EZVIZ_P2P_SVR_DEV_NOT_BELONG_TO_USER:
		return errors.New("P2PCLOUD_ER_SVR_DEV_NOT_BELONG_TO_USER  5005")

	//传显错误码 8501~9500
	case NET_ERR_UPGRADE_PROG_ERR:
		return errors.New("程序执行出错")
	case NET_ERR_UPGRADE_NO_DEVICE:
		return errors.New("没有设备(指LED控制器没有接接收卡)")
	case NET_ERR_UPGRADE_NO_FILE:
		return errors.New("没有找到升级文件")
	case NET_ERR_UPGRADE_DATA_ERROR:
		return errors.New("升级文件数据不兼容")
	case NET_ERR_UPGRADE_LINK_SERVER_ERR:
		return errors.New("与服务器连接失败")
	case NET_ERR_UPGRADE_OEMCODE_NOMATCH:
		return errors.New("oemCode不匹配")
	case NET_ERR_UPGRADE_FLASH_NOENOUGH:
		return errors.New("flash不足")
	case NET_ERR_UPGRADE_RAM_NOENOUGH:
		return errors.New("RAM不足")
	case NET_ERR_UPGRADE_DSPRAM_NOENOUGH:
		return errors.New("DSP RAM不足")
	case NET_ERR_NOT_SUPPORT_CHECK:
		return errors.New("该屏幕型号不支持校正")
	case NET_ERR_LED_DEVICE_BUSY_CHECK:
		return errors.New("LED设备忙（正在校正）")
	case NET_ERR_DEVICE_MEM_NOT_ENOUGH:
		return errors.New("设备内存不足")
	case NET_ERR_CHECK_PARAM:
		return errors.New("校正参数错误")
	case NET_ERR_RESOLUTION_OVER_LIMIT:
		return errors.New("输入分辨率超过限制")
	case NET_ERR_NO_CUSTOM_BASE:
		return errors.New("无自定义底图")
	case NET_ERR_PRIORITY_LOWER:
		return errors.New("优先级低于当前模式")
	case NET_ERR_SEND_MESSAGE_EXCEPT:
		return errors.New("消息发送异常")
	case NET_ERR_SENDCARD_UPGRADING:
		return errors.New("发送卡升级中")
	case NET_ERR_NO_WIRELESS_NETCARD:
		return errors.New("未插入无线网卡")
	case NET_ERR_LOAD_FS_FAIL:
		return errors.New("从屏幕加载失败")
	case NET_ERR_FLASH_UNSTORAGE_RECCARD:
		return errors.New("Flash中未存储接收卡参数")
	default:
		return errors.New("未知错误")
	}
}
