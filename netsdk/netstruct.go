package netsdk

/*
#cgo CFLAGS: -I ../include
#cgo LDFLAGS: -L ../lib -lhcnetsdk

#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type (
	LLONG = int64
	LONG  = int32
	DWORD = uint32
	WORD  = uint16
	BYTE  = byte
	BOOL  = int32
)

type DeviceAbilityKind int

const (
	DADeviceSofthardware         DeviceAbilityKind = 0x001 //设备软硬件能力
	DADeviceNetwork              DeviceAbilityKind = 0x002 //设备网络能力
	DADeviceEncodeAll            DeviceAbilityKind = 0x003 //设备所有编码能力
	DADeviceEncodeCurrent        DeviceAbilityKind = 0x004 //设备当前编码能力
	DAIpcFrontParameter          DeviceAbilityKind = 0x005 //ipc前端参数1.0
	DAIpcUpgradeDescription      DeviceAbilityKind = 0x006 //ipc升级信息
	DADeviceRaid                 DeviceAbilityKind = 0x007 //RAID能力
	DADeviceEncodeAllV20         DeviceAbilityKind = 0x008 //设备所有编码能力2.0
	DAIpcFrontParameterV20       DeviceAbilityKind = 0x009 //ipc前端参数2.0
	DADeviceAlarm                DeviceAbilityKind = 0x00a //辅助报警能力
	DADeviceDynchan              DeviceAbilityKind = 0x00b //设备数字通道能力
	DADeviceUser                 DeviceAbilityKind = 0x00c //设备用户管理参数能力
	DADeviceNetapp               DeviceAbilityKind = 0x00d //设备网络应用参数能力
	DADeviceVideopic             DeviceAbilityKind = 0x00e //设备图像参数能力
	DADeviceJpegCap              DeviceAbilityKind = 0x00f //设备JPEG抓图能力
	DADeviceSerial               DeviceAbilityKind = 0x010 //RS232和RS485串口能力
	DADeviceInfo                 DeviceAbilityKind = 0x011 //设备通用能力类型，具体能力根据发送的能力节点来区分
	DAStream                     DeviceAbilityKind = 0x012 //流能力
	DASystemManagement           DeviceAbilityKind = 0x013 //设备系统管理能力
	DAIpViewDev                  DeviceAbilityKind = 0x014 //IP可视对讲分机能力
	DAVcaDev                     DeviceAbilityKind = 0x100 //设备智能分析的总能力
	DAVcaChan                    DeviceAbilityKind = 0x110 //行为分析能力
	DATransfer                   DeviceAbilityKind = 0x120
	DAMatrixdecoder              DeviceAbilityKind = 0x200 //多路解码器显示、解码能力
	DAVideoplatform              DeviceAbilityKind = 0x210 //视频综合平台能力集
	DAVideoplatformSbucodesystem DeviceAbilityKind = 0x211 //视频综合平台编码子系统能力集
	DAWall                       DeviceAbilityKind = 0x212 //电视墙能力集
	DAMatrix                     DeviceAbilityKind = 0x213 //SDI矩阵能力
	DADecodecard                 DeviceAbilityKind = 0x220 //解码卡服务器能力集
	DAVideoplatformV40           DeviceAbilityKind = 0x230 //视频综合平台能力集
	DAMatrixmanagedevice         DeviceAbilityKind = 0x240 //矩阵接入网关能力集
	DAMatrixdecoderV41           DeviceAbilityKind = 0x260 //解码器能力集
	DADecoder                    DeviceAbilityKind = 0x261 //解码器xml能力集
	DADecodecardV41              DeviceAbilityKind = 0x270 //解码卡服务器能力集V41
	DACodecard                   DeviceAbilityKind = 0x271 //编码卡能力集
	DASnapcamera                 DeviceAbilityKind = 0x300 //抓拍机能力集
	DAItcTriggerMode             DeviceAbilityKind = 0x301 //智能IPC设备的触发模式能力
	DACompressioncfg             DeviceAbilityKind = 0x400 //获取压缩参数能力集合
	DACompressionLimit           DeviceAbilityKind = 0x401 //获取主子码流压缩参数能力限制
	DAPicCapture                 DeviceAbilityKind = 0x402 //获图片分辨率能力集合
	DAAlarmhost                  DeviceAbilityKind = 0x500 //网络报警主机能力集
	DAItDevice                   DeviceAbilityKind = 0x501 //智能交通能力集
	DAScreencontrol              DeviceAbilityKind = 0x600 //大屏控制器能力集
	DAScreenserver               DeviceAbilityKind = 0x610 //大屏服务器能力集
	DAFisheye                    DeviceAbilityKind = 0x700 //鱼眼能力集
	DALcdScreen                  DeviceAbilityKind = 0x800 //LCD屏幕能力 2013-10-12
	DAAcs                        DeviceAbilityKind = 0x801 //门禁能力
	DAMergedev                   DeviceAbilityKind = 0x802 //合码器能力集
	DACamFusion                  DeviceAbilityKind = 0x803 //相机拼接能力
	DAOpticalDevAccess           DeviceAbilityKind = 0x805 //光端机接入能力
	DANetRing                    DeviceAbilityKind = 0x806 //环网能力集
	DALed                        DeviceAbilityKind = 0x807 //LED屏能力集
	DAPublishdev                 DeviceAbilityKind = 0x80a //信息发布能力
	DAScreenExchange             DeviceAbilityKind = 0x80b //屏幕互动能力
	DARemoteNetmgrFot            DeviceAbilityKind = 0x80e //远端网管收发器能力
)

type NET_DVR_DEVICEINFO_V30 struct {
	ST_sSerialNumber     [48]BYTE //序列号
	ST_byAlarmInPortNum  BYTE     //报警输入个数
	ST_byAlarmOutPortNum BYTE     //报警输出个数
	ST_byDiskNum         BYTE     //硬盘个数
	ST_byDVRType         BYTE     //设备类型, 1:DVR 2:ATM DVR 3:DVS ......
	ST_byChanNum         BYTE     //模拟通道个数
	ST_byStartChan       BYTE     //起始通道号,例如DVS-1,DVR - 1
	ST_byAudioChanNum    BYTE     //语音通道数
	ST_byIPChanNum       BYTE     //最大数字通道个数，低位
	ST_byZeroChanNum     BYTE     //零通道编码个数 //2010-01-16
	ST_byMainProto       BYTE     //主码流传输协议类型 0-private, 1-rtsp,2-同时支持private和rtsp
	ST_bySubProto        BYTE     //子码流传输协议类型0-private, 1-rtsp,2-同时支持private和rtsp
	ST_bySupport         BYTE     //能力，位与结果为0表示不支持，1表示支持，
	//bySupport & 0x1, 表示是否支持智能搜索
	//bySupport & 0x2, 表示是否支持备份
	//bySupport & 0x4, 表示是否支持压缩参数能力获取
	//bySupport & 0x8, 表示是否支持多网卡
	//bySupport & 0x10, 表示支持远程SADP
	//bySupport & 0x20, 表示支持Raid卡功能
	//bySupport & 0x40, 表示支持IPSAN 目录查找
	//bySupport & 0x80, 表示支持rtp over rtsp
	ST_bySupport1 BYTE // 能力集扩充，位与结果为0表示不支持，1表示支持
	//bySupport1 & 0x1, 表示是否支持snmp v30
	//bySupport1 & 0x2, 支持区分回放和下载
	//bySupport1 & 0x4, 是否支持布防优先级
	//bySupport1 & 0x8, 智能设备是否支持布防时间段扩展
	//bySupport1 & 0x10, 表示是否支持多磁盘数（超过33个）
	//bySupport1 & 0x20, 表示是否支持rtsp over http
	//bySupport1 & 0x80, 表示是否支持车牌新报警信息2012-9-28, 且还表示是否支持NET_DVR_IPPARACFG_V40结构体
	ST_bySupport2 BYTE /*能力，位与结果为0表示不支持，非0表示支持
	                 bySupport2 & 0x1, 表示解码器是否支持通过URL取流解码
	                 bySupport2 & 0x2,  表示支持FTPV40
	                 bySupport2 & 0x4,  表示支持ANR
	                 bySupport2 & 0x8,  表示支持CCD的通道参数配置
	                 bySupport2 & 0x10,  表示支持布防报警回传信息（仅支持抓拍机报警 新老报警结构）
	                 bySupport2 & 0x20,  表示是否支持单独获取设备状态子项
	bySupport2 & 0x40,  表示是否是码流加密设备*/
	ST_wDevType   WORD //设备型号
	ST_bySupport3 BYTE //能力集扩展，位与结果为0表示不支持，1表示支持
	//bySupport3 & 0x1, 表示是否支持批量配置多码流参数
	// bySupport3 & 0x4 表示支持按组配置， 具体包含 通道图像参数、报警输入参数、IP报警输入、输出接入参数、
	// 用户参数、设备工作状态、JPEG抓图、定时和时间抓图、硬盘盘组管理
	//bySupport3 & 0x8为1 表示支持使用TCP预览、UDP预览、多播预览中的"延时预览"字段来请求延时预览（后续都将使用这种方式请求延时预览）。而当bySupport3 & 0x8为0时，将使用 "私有延时预览"协议。
	//bySupport3 & 0x10 表示支持"获取报警主机主要状态（V40）"。
	//bySupport3 & 0x20 表示是否支持通过DDNS域名解析取流

	ST_byMultiStreamProto BYTE //是否支持多码流,按位表示,0-不支持,1-支持,bit1-码流3,bit2-码流4,bit7-主码流，bit-8子码流
	ST_byStartDChan       BYTE //起始数字通道号,0表示无效
	ST_byStartDTalkChan   BYTE //起始数字对讲通道号，区别于模拟对讲通道号，0表示无效
	ST_byHighDChanNum     BYTE //数字通道个数，高位
	ST_bySupport4         BYTE //能力集扩展，位与结果为0表示不支持，1表示支持
	//bySupport4 & 0x4表示是否支持拼控统一接口
	// bySupport4 & 0x80 支持设备上传中心报警使能。表示判断调用接口是 NET_DVR_PDC_RULE_CFG_V42还是 NET_DVR_PDC_RULE_CFG_V41
	ST_byLanguageType BYTE // 支持语种能力,按位表示,每一位0-不支持,1-支持
	//  byLanguageType 等于0 表示 老设备
	//  byLanguageType & 0x1表示支持中文
	//  byLanguageType & 0x2表示支持英文
	ST_byVoiceInChanNum     BYTE //音频输入通道数
	ST_byStartVoiceInChanNo BYTE //音频输入起始通道号 0表示无效
	ST_bySupport5           BYTE //按位表示,0-不支持,1-支持,bit0-支持多码流
	//bySupport5 &0x01表示支持wEventTypeEx ,兼容dwEventType 的事件类型（支持行为事件扩展）--先占住，防止冲突
	//bySupport5 &0x04表示是否支持使用扩展的场景模式接口
	/*
	   bySupport5 &0x08 设备返回该值表示是否支持计划录像类型V40接口协议(DVR_SET_RECORDCFG_V40/ DVR_GET_RECORDCFG_V40)(在该协议中设备支持类型类型13的配置)
	   之前的部分发布的设备，支持录像类型13，则配置录像类型13。如果不支持，统一转换成录像类型3兼容处理。SDK通过命令探测处理)
	   bySupport5 &0x10 设备返回改值表示支持超过255个预置点
	*/
	ST_bySupport6 BYTE //能力，按位表示，0-不支持,1-支持
	//bySupport6 0x1  表示设备是否支持压缩
	//bySupport6 0x2 表示是否支持流ID方式配置流来源扩展命令，DVR_SET_STREAM_SRC_INFO_V40
	//bySupport6 0x4 表示是否支持事件搜索V40接口
	//bySupport6 0x8 表示是否支持扩展智能侦测配置命令
	//bySupport6 0x40表示图片查询结果V40扩展
	ST_byMirrorChanNum    BYTE //镜像通道个数，<录播主机中用于表示导播通道>
	ST_wStartMirrorChanNo WORD //起始镜像通道号
	ST_bySupport7         BYTE //能力,按位表示,0-不支持,1-支持
	// bySupport7 & 0x1  表示设备是否支持 INTER_VCA_RULECFG_V42 扩展
	// bySupport7 & 0x2  表示设备是否支持 IPC HVT 模式扩展
	// bySupport7 & 0x04  表示设备是否支持 返回锁定时间
	// bySupport7 & 0x08  表示设置云台PTZ位置时，是否支持带通道号
	// bySupport7 & 0x10  表示设备是否支持双系统升级备份
	// bySupport7 & 0x20  表示设备是否支持 OSD字符叠加 V50
	// bySupport7 & 0x40  表示设备是否支持 主从跟踪（从摄像机）
	// bySupport7 & 0x80  表示设备是否支持 报文加密
	ST_byRes2 BYTE //保留
}

const MAX_RESOLUTIONNUM = 64

type NET_DVR_MATRIX_ABILITY struct {
	ST_dwSize              DWORD
	ST_byDecNums           BYTE
	ST_byStartChan         BYTE
	ST_byVGANums           BYTE
	ST_byBNCNums           BYTE
	ST_byVGAWindowMode     [8][12]BYTE /*VGA支持的窗口模式，VGA1可能由混合输出*/
	ST_byBNCWindowMode     [4]BYTE     /*BNC支持的窗口模式*/
	ST_byDspNums           BYTE
	ST_byHDMINums          BYTE //HDMI显示通道个数（从25开始）
	ST_byDVINums           BYTE //DVI显示通道个数（从29开始）
	ST_byRes1              [13]BYTE
	ST_bySupportResolution [MAX_RESOLUTIONNUM]BYTE //按照上面的枚举定义,一个字节代表一个分辨率是//否支持，1：支持，0：不支持
	ST_byHDMIWindowMode    [4][8]BYTE              //HDMI支持的窗口模式
	ST_byDVIWindowMode     [4][8]BYTE              //DVI支持的窗口模式
	ST_byRes2              [24]BYTE
}

type NET_DVR_ALARMER struct {
	ST_byUserIDValid     BYTE               /* userid是否有效 0-无效，1-有效 */
	ST_bySerialValid     BYTE               /* 序列号是否有效 0-无效，1-有效 */
	ST_byVersionValid    BYTE               /* 版本号是否有效 0-无效，1-有效 */
	ST_byDeviceNameValid BYTE               /* 设备名字是否有效 0-无效，1-有效 */
	ST_byMacAddrValid    BYTE               /* MAC地址是否有效 0-无效，1-有效 */
	ST_byLinkPortValid   BYTE               /* login端口是否有效 0-无效，1-有效 */
	ST_byDeviceIPValid   BYTE               /* 设备IP是否有效 0-无效，1-有效 */
	ST_bySocketIPValid   BYTE               /* socket ip是否有效 0-无效，1-有效 */
	ST_lUserID           LONG               /* NET_DVR_Login()返回值, 布防时有效 */
	ST_sSerialNumber     [SERIALNO_LEN]BYTE /* 序列号 */
	ST_dwDeviceVersion   DWORD              /* 版本信息 高16位表示主版本，低16位表示次版本*/
	ST_sDeviceName       [NAME_LEN]byte     /* 设备名字 */
	ST_byMacAddr         [MACADDR_LEN]BYTE  /* MAC地址 */
	ST_wLinkPort         WORD               /* link port */
	ST_sDeviceIP         [128]byte          /* IP地址 */
	ST_sSocketIP         [128]byte          /* 报警主动上传时的socket IP地址 */
	ST_byIpProtocol      BYTE               /* Ip协议 0-IPV4, 1-IPV6 */
	ST_byRes1            [2]BYTE
	ST_bJSONBroken       BYTE //JSON断网续传标志。0：不续传；1：续传
	ST_wSocketPort       WORD
	ST_byRes2            [6]BYTE
}

type NET_DVR_SETUPALARM_PARAM struct {
	ST_dwSize              DWORD
	ST_byLevel             BYTE //布防优先级，0-一等级（高），1-二等级（中），2-三等级（低）
	ST_byAlarmInfoType     BYTE //上传报警信息类型（抓拍机支持），0-老报警信息（NET_DVR_PLATE_RESULT），1-新报警信息(NET_ITS_PLATE_RESULT)2012-9-28
	ST_byRetAlarmTypeV40   BYTE //0--返回NET_DVR_ALARMINFO_V30或NET_DVR_ALARMINFO, 1--设备支持NET_DVR_ALARMINFO_V40则返回NET_DVR_ALARMINFO_V40，不支持则返回NET_DVR_ALARMINFO_V30或NET_DVR_ALARMINFO
	ST_byRetDevInfoVersion BYTE //CVR上传报警信息回调结构体版本号 0-COMM_ALARM_DEVICE， 1-COMM_ALARM_DEVICE_V40
	ST_byRetVQDAlarmType   BYTE //VQD报警上传类型，0-上传报报警NET_DVR_VQD_DIAGNOSE_INFO，1-上传报警NET_DVR_VQD_ALARM
	//1-表示人脸侦测报警扩展(INTER_FACE_DETECTION),0-表示原先支持结构(INTER_FACESNAP_RESULT)
	ST_byFaceAlarmDetection BYTE
	//Bit0- 表示二级布防是否上传图片: 0-上传，1-不上传
	//Bit1- 表示开启数据上传确认机制；0-不开启，1-开启
	//Bit6- 表示雷达检测报警(eventType:radarDetection)是否开启实时上传；0-不开启，1-开启（用于web插件实时显示雷达目标轨迹）
	ST_bySupport BYTE
	//断网续传类型

	//bit0-车牌检测（IPC） （0-不续传，1-续传）
	//bit1-客流统计（IPC）  （0-不续传，1-续传）
	//bit2-热度图统计（IPC） （0-不续传，1-续传）
	//bit3-人脸抓拍（IPC） （0-不续传，1-续传）
	//bit4-人脸对比（IPC） （0-不续传，1-续传）
	ST_byBrokenNetHttp BYTE
	ST_wTaskNo         WORD //任务处理号 和 (上传数据NET_DVR_VEHICLE_RECOG_RESULT中的字段dwTaskNo对应 同时 下发任务结构 NET_DVR_VEHICLE_RECOG_COND中的字段dwTaskNo对应)
	ST_byDeployType    BYTE //布防类型：0-客户端布防，1-实时布防
	ST_bySubScription  BYTE //订阅，按位表示，未开启订阅不上报  //占位
	//Bit7-移动侦测人车分类是否传图；0-不传图(V30上报)，1-传图(V40上报)
	ST_byRes1         [2]BYTE
	ST_byAlarmTypeURL BYTE
	//bit0-表示人脸抓拍报警上传（INTER_FACESNAP_RESULT）；0-表示二进制传输，1-表示URL传输（设备支持的情况下，设备支持能力根据具体报警能力集判断,同时设备需要支持URL的相关服务，当前是”云存储“）
	//bit1-表示EVENT_JSON中图片数据长传类型；0-表示二进制传输，1-表示URL传输（设备支持的情况下，设备支持能力根据具体报警能力集判断）
	//bit2 - 人脸比对(报警类型为COMM_SNAP_MATCH_ALARM)中图片数据上传类型：0 - 二进制传输，1 - URL传输
	//bit3 - 行为分析(报警类型为COMM_ALARM_RULE)中图片数据上传类型：0 - 二进制传输，1 - URL传输，本字段设备是否支持，对应软硬件能力集中<isSupportBehaviorUploadByCloudStorageURL>节点是否返回且为true
	ST_byCustomCtrl BYTE
	//Bit0- 表示支持副驾驶人脸子图上传: 0-不上传,1-上传
}

type NET_DVR_IPADDR struct {
	ST_sIpV4  [16]byte  /* IPv4地址 */
	ST_byIPv6 [128]BYTE /* 保留 */
}

type NET_VCA_DEV_INFO struct {
	ST_struDevIP     NET_DVR_IPADDR //前端设备地址，
	ST_wPort         WORD           //前端设备端口号，
	ST_byChannel     BYTE           //前端设备通道，
	ST_byIvmsChannel BYTE           // Ivms 通道
}

type NET_DVR_LLI_PARAM struct {
	ST_fSec     float32 //秒[0.000000,60.000000]
	ST_byDegree BYTE    //度:纬度[0,90] 经度[0,180]
	ST_byMinute BYTE    //分[0,59]
	ST_byRes    [6]BYTE
}

type NET_DVR_SENSOR_PARAM struct {
	ST_bySensorType BYTE //SensorType:0-CCD,1-CMOS
	ST_byRes        [31]BYTE
	ST_fHorWidth    float32 //水平宽度 精确到小数点后两位 *10000
	ST_fVerWidth    float32 //垂直宽度 精确到小数点后两位 *10000
	ST_fFold        float32 //zoom=1没变时的焦距 精确到小数点后两位 *100
}

//球机位置信息
type NET_DVR_PTZPOS_PARAM struct {
	ST_fPanPos  float32 //水平参数，精确到小数点后1位
	ST_fTiltPos float32 //垂直参数，精确到小数点后1位
	ST_fZoomPos float32 //变倍参数，精确到小数点后1位
	ST_byRes    [16]BYTE
}

type NET_DVR_GIS_UPLOADINFO struct {
	ST_dwSize           DWORD            //结构体大小
	ST_dwRelativeTime   DWORD            //相对时标
	ST_dwAbsTime        DWORD            //绝对时标
	ST_struDevInfo      NET_VCA_DEV_INFO //前端设备
	ST_fAzimuth         float32          //电子罗盘的方位信息；方位角[0.00°,360.00°)
	ST_byLatitudeType   BYTE             //纬度类型，0-北纬，1-南纬
	ST_byLongitudeType  BYTE             // 经度类型，0-东度，1-西度
	ST_byRes1           [2]BYTE
	ST_struLatitude     NET_DVR_LLI_PARAM    /*纬度*/
	ST_struLongitude    NET_DVR_LLI_PARAM    /*经度*/
	ST_fHorizontalValue float32              //水平视场角，精确到小数点后面两位
	ST_fVerticalValue   float32              //垂直视场角，精确到小数点后面两位
	ST_fVisibleRadius   float32              //当前可视半径，精确到小数点后面两位
	ST_fMaxViewRadius   float32              //最大可视半径，精确到小数点后面0位（预留处理）
	ST_struSensorParam  NET_DVR_SENSOR_PARAM //Sensor信息
	ST_struPtzPos       NET_DVR_PTZPOS_PARAM //ptz坐标
	ST_byRes            [256]BYTE
}

type NET_DVR_PTZPOS struct {
	ST_wAction  WORD //获取时该字段无效
	ST_wPanPos  WORD //水平参数
	ST_wTiltPos WORD //垂直参数
	ST_wZoomPos WORD //变倍参数
}
type NET_DVR_SCENE_INFO struct {
	ST_dwSceneID   DWORD          //场景ID, 0 - 表示该场景无效
	ST_bySceneName [NAME_LEN]BYTE //场景名称
	ST_byDirection BYTE           //监测方向 1-上行，2-下行，3-双向，4-由东向西，5-由南向北，6-由西向东，7-由北向南，8-其它
	ST_byRes1      [3]BYTE        //保留
	ST_struPtzPos  NET_DVR_PTZPOS //Ptz 坐标
	ST_byRes2      [64]BYTE       //保留
}

type NET_DVR_TIME_EX struct {
	ST_wYear    WORD
	ST_byMonth  BYTE
	ST_byDay    BYTE
	ST_byHour   BYTE
	ST_byMinute BYTE
	ST_bySecond BYTE
	ST_byRes    BYTE
}

type NET_VCA_POINT struct {
	ST_fX float32
	ST_fY float32
}

type NET_DVR_DIRECTION struct {
	ST_struStartPoint NET_VCA_POINT // 方
	ST_struEndPoint   NET_VCA_POINT // 方
}

type NET_DVR_AID_INFO struct {
	ST_byRuleID            BYTE // 规则序号，为规则配置结构下标，0-16
	ST_byVisibilityLevel   BYTE // 能见度等级：0-保留；1-无雾~薄雾；2-薄雾~中雾；3-大雾~浓雾；4-浓雾及以上
	ST_byRes1              [2]BYTE
	ST_byRuleName          [NAME_LEN]BYTE    //  规则名称
	ST_dwAIDType           DWORD             // 报警事件类型
	ST_struDirect          NET_DVR_DIRECTION // 报警指向区域
	ST_bySpeedLimit        BYTE              //限速值，单位km/h[0,255]
	ST_byCurrentSpeed      BYTE              //当前速度值，单位km/h[0,255]
	ST_byVehicleEnterState BYTE              //车辆出入状态 0-无效 1-驶入 2-驶出
	ST_byState             BYTE              //0-变化上传，1-轮巡上传，2-当前设备定时抓拍的数据上传，实际作用于平台形成图片序列，用于反查算法没有检测到的停车车辆（索引值2在“dwAIDType;//报警事件类型”为 “停车事件”的时候生效）
	ST_byParkingID         [16]BYTE          //停车位编号
	ST_dwAIDTypeEx         DWORD             // 报警事件类型扩展,参考TRAFFIC_AID_TYPE_EX
	ST_byRes2              [16]BYTE          // 保留字节
}

type NET_VCA_RECT struct {
	ST_fX      float32 //边界框左上角点的X轴坐标, 0.000~1
	ST_fY      float32 //边界框左上角点的Y轴坐标, 0.000~1
	ST_fWidth  float32 //边界框的宽度, 0.000~1
	ST_fHeight float32 //边界框的高度, 0.000~1
}

type NET_DVR_PLATE_INFO struct {
	ST_byPlateType     BYTE //车牌类型
	ST_byColor         BYTE //车牌颜色
	ST_byBright        BYTE //车牌亮度
	ST_byLicenseLen    BYTE //车牌字符个数
	ST_byEntireBelieve BYTE //整个车牌的置信度，-100
	ST_byRegion        BYTE // 区域索引值 0-保留，1-欧洲(EU)，2-俄语区域(ER)，3-欧洲&俄罗斯(EU&CIS) ,4-中东(ME),0xff-所有
	ST_byCountry       BYTE // 国家索引值，参照枚举COUNTRY_INDEX（不支持"COUNTRY_ALL = 0xff, //ALL  全部"）
	ST_byArea          BYTE //区域（省份），各国家内部区域枚举，阿联酋参照 EMI_AREA
	ST_byPlateSize     BYTE //车牌尺寸，0~未知，1~long, 2~short(中东车牌使用)
	/*附加信息标识（即是否有NET_DVR_VEHICLE_ADDINFO结构体）,0-无附加信息, 1-有附加信息。*/
	ST_byAddInfoFlag BYTE
	//该字段是在byCountry索引基础上，扩展了区域索引，之后使用该字段代替byCountry，优先使用wCRIndex字段
	//为了兼容老用户，如果该字段值大于256（即新增区域），则byCountry赋值为0xfd（国家字段无效）。
	ST_wCRIndex       WORD    //国家/地区索引，索引值参考_CR_ INDEX_
	ST_byRes          [4]BYTE //保留
	ST_pAddInfoBuffer *BYTE
	ST_sPlateCategory [MAX_CATEGORY_LEN]byte //车牌附加信息, 即中东车牌中车牌号码旁边的小字信息，(目前只有中东地区支持)
	ST_dwXmlLen       DWORD                  //XML报警信息长度
	ST_pXmlBuf        *byte                  // XML报警信息指针,报警类型为 COMM_ITS_PLATE_RESUL时有效，其XML对应到EventNotificationAlert XML Block
	ST_struPlateRect  NET_VCA_RECT           //车牌位置
	ST_sLicense       [MAX_LICENSE_LEN]byte  //车牌号码,注：中东车牌需求把小字也纳入车牌号码，小字和车牌号中间用空格分隔
	ST_byBelieve      [MAX_LICENSE_LEN]BYTE  //各个识别字符的置信度，如检测到车牌"浙A12345", 置信度为,20,30,40,50,60,70，则表示"浙"字正确的可能性只有%，"A"字的正确的可能性是%
}

type NET_DVR_VEHICLE_INFO struct {
	ST_dwIndex       DWORD //车辆序号
	ST_byVehicleType BYTE  //车辆类型 0 表示其它车型，1 表示小型车，2 表示大型车 ,3表示行人触发 ,4表示二轮车触发 5表示三轮车触发(3.5Ver)  6表示机动车触发
	ST_byColorDepth  BYTE  //车身颜色深浅
	ST_byColor       BYTE  //车身颜色,参考VCR_CLR_CLASS
	ST_byRadarState  BYTE
	ST_wSpeed        WORD //单位km/h
	ST_wLength       WORD //前一辆车的车身长度
	/*违规类型，0-正常，1-低速，2-超速，3-逆行，4-闯红灯,5-压车道线,6-不按导向，7-路口滞留，
	                          8-机占非，9-违法变道，10-不按车道 11-违反禁令，12-路口停车，13-绿灯停车, 14-未礼让行人(违法代码1357),
	  15-违章停车，16-违章掉头,17-占用应急车道,18-禁右,19-禁左,20-压黄线,21-未系安全带,22-行人闯红灯,23-加塞,24-违法使用远光灯，
	  25-驾驶时拨打接听手持电话，26-左转不让直行，27-右转不让左转，28-掉头不让直行，29-大弯小转, 30-闯绿灯，31-未带头盔，
	  32-非机动车载人，33-非机动车占用机动车道，34-非机动车打伞棚, 35-黑烟车, 36-鸣笛,37-压线停车,38-跨位停车,39-压线且跨位停车
	  40-不让右方道路来车先行 41-进入环形路口未让已在路口内的机动车先行 42-机动车从匝道进入主路未让行
	*/
	ST_byIllegalType             BYTE
	ST_byVehicleLogoRecog        BYTE     //参考枚举类型 VLR_VEHICLE_CLASS
	ST_byVehicleSubLogoRecog     BYTE     //车辆品牌子类型识别；参考VSB_VOLKSWAGEN_CLASS等子类型枚举。
	ST_byVehicleModel            BYTE     //车辆子品牌年款，0-未知，参考"车辆子品牌年款.xlsx"
	ST_byCustomInfo              [16]BYTE //自定义信息
	ST_wVehicleLogoRecog         WORD     //车辆主品牌，参考"车辆主品牌.xlsx" (该字段兼容byVehicleLogoRecog);
	ST_byIsParking               BYTE     //是否停车 0-无效，1-停车，2-未停车
	ST_byRes                     BYTE     //保留字节
	ST_dwParkingTime             DWORD    //停车时间，单位：s
	ST_byBelieve                 BYTE     //byIllegalType置信度，1-100
	ST_byCurrentWorkerNumber     BYTE     //当前作业人数
	ST_byCurrentGoodsLoadingRate BYTE     //当前货物装载率 0-空 1-少 2-中 3-多 4-满
	ST_byDoorsStatus             BYTE     //车门状态 0-车门关闭 1-车门开启
	ST_byRes3                    [4]BYTE
}

// 图片信息（后续会加入码流）
type NET_ITS_PICTURE_INFO struct {
	ST_dwDataLen DWORD //媒体数据长度
	// 0:车牌图;1:车辆图;2:合成图; 3:特写图;4:二直图;5:码流;6:人脸子图(主驾驶);7:人脸子图(副驾驶)成图;8-非机动车;9-行人;10-称重原始裸数据;11-目标图;12-主驾驶室图 ;13-副驾驶室图;14-人脸图抠小图
	//15 - 自定义图片(用户自己上传进行违法检测的图片)
	ST_byType BYTE
	// 0-数据直接上传; 1-云存储服务器URL(3.7Ver)原先的图片数据变成URL数据，图片长度变成URL长度
	ST_byDataType         BYTE
	ST_byCloseUpType      BYTE         //特写图类型，0-保留,1-非机动车,2-行人
	ST_byPicRecogMode     BYTE         //图片背向识别：0-正向车牌识别，1-背向识别(尾牌识别) ；
	ST_dwRedLightTime     DWORD        //经过的红灯时间  （s）
	ST_byAbsTime          [32]BYTE     //绝对时间点,yyyymmddhhmmssxxx,e.g.20090810235959999  最后三位为毫秒数
	ST_struPlateRect      NET_VCA_RECT //车牌位置,当byType为8-非机动车;9-行人时，表示人体坐标
	ST_struPlateRecgRect  NET_VCA_RECT //牌识区域坐标，当图片类型为12-主驾驶室图13-副驾驶室图是，该坐标为驾驶员坐标
	ST_pBuffer            *BYTE        //数据指针
	ST_dwUTCTime          DWORD        //UTC时间定义
	ST_byCompatibleAblity BYTE         //兼容能力字段 0表示无效，1表示有效; bit0-表示dwUTCTime字段有效
	ST_byTimeDiffFlag     BYTE         /*时差字段是否有效  0-时差无效， 1-时差有效 */
	ST_cTimeDifferenceH   byte         /*与UTC的时差（小时），-12 ... +14， +表示东区,，byTimeDiffFlag为1时有效*/
	ST_cTimeDifferenceM   byte         /*与UTC的时差（分钟），-30, 30, 45， +表示东区，byTimeDiffFlag为1时有效*/
	ST_byRes2             [4]BYTE      //保留
}

type NET_DVR_TIME_V30 struct {
	ST_wYear            WORD
	ST_byMonth          BYTE
	ST_byDay            BYTE
	ST_byHour           BYTE
	ST_byMinute         BYTE
	ST_bySecond         BYTE
	ST_byISO8601        BYTE /*是否是8601的时间格式，即时差字段是否有效0-时差无效，年月日时分秒为设备本地时间 1-时差有效 */
	ST_wMilliSec        WORD //毫秒，精度不够，默认为0
	ST_cTimeDifferenceH byte //与UTC的时差（小时），-12 ... +14，+表示东区, byISO8601为1时有效
	ST_cTimeDifferenceM byte //与UTC的时差（分钟），-30, 30, 45，+表示东区，byISO8601为1时有效
}

type NET_DVR_TFS_ALARM struct {
	ST_dwSize                  DWORD                    //结构体大小
	ST_dwRelativeTime          DWORD                    //相对时标
	ST_dwAbsTime               DWORD                    //绝对时标
	ST_dwIllegalType           DWORD                    //违章类型，采用国标定义，当dwIllegalType值为0xffffffff时使用 DWORD byIllegalCode
	ST_dwIllegalDuration       DWORD                    //违法持续时间（单位：秒） = 抓拍最后一张图片的时间 - 抓拍第一张图片的时间
	ST_byMonitoringSiteID      [MONITORSITE_ID_LEN]BYTE //监测点编号（路口编号、内部编号）
	ST_byDeviceID              [DEVICE_ID_LEN]BYTE      //设备编号
	ST_struDevInfo             NET_VCA_DEV_INFO         //前端设备信息
	ST_struSceneInfo           NET_DVR_SCENE_INFO       //场景信息
	ST_struBeginRecTime        NET_DVR_TIME_EX          //录像开始时间
	ST_struEndRecTime          NET_DVR_TIME_EX          //录像结束时间
	ST_struAIDInfo             NET_DVR_AID_INFO         //交通事件信息
	ST_struPlateInfo           NET_DVR_PLATE_INFO       //车牌信息
	ST_struVehicleInfo         NET_DVR_VEHICLE_INFO     //车辆信息
	ST_dwPicNum                DWORD                    //图片数量
	ST_struPicInfo             [8]NET_ITS_PICTURE_INFO  //图片信息，最多8张
	ST_bySpecificVehicleType   BYTE                     //具体车辆种类  参考识别结果类型 BYTE VTR_RESULT
	ST_byLaneNo                BYTE                     //关联车道号
	ST_byRes1                  [2]BYTE                  //保留
	ST_struTime                NET_DVR_TIME_V30         //手动跟踪定位，当前时间。
	ST_dwSerialNo              DWORD                    //序号；
	ST_byVehicleAttribute      BYTE                     //车辆属性，按位表示，0- 无附加属性(普通车)，bit1- 黄标车(类似年检的标志)，bit2- 危险品车辆，值：0- 否，1- 是
	ST_byPilotSafebelt         BYTE                     //0-表示未知,1-系安全带,2-不系安全带
	ST_byCopilotSafebelt       BYTE                     //0-表示未知,1-系安全带,2-不系安全带
	ST_byPilotSunVisor         BYTE                     //0-表示未知,1-不打开遮阳板,2-打开遮阳板
	ST_byCopilotSunVisor       BYTE                     //0-表示未知, 1-不打开遮阳板,2-打开遮阳板
	ST_byPilotCall             BYTE                     // 0-表示未知, 1-不打电话,2-打电话
	ST_byRes2                  [2]BYTE                  //保留
	ST_byIllegalCode           [ILLEGAL_LEN]BYTE        //违法代码扩展，当dwIllegalType值为0xffffffff；使用这个值
	ST_wCountry                WORD                     // 国家索引值,参照枚举BYTE COUNTRY_INDEX
	ST_byRegion                BYTE                     //区域索引值,0-保留，1-欧洲(Europe Region)，2-俄语区域(Russian Region)，3-欧洲&俄罗斯(EU&CIS) , 4-中东（Middle East），0xff-所有
	ST_byCrossLine             BYTE                     //是否压线停车（侧方停车），0-表示未知，1-不压线，2-压线
	ST_byParkingSerialNO       [SERIAL_NO_LEN]BYTE      //泊车位编号
	ST_byCrossSpaces           BYTE                     //是否跨泊车位停车（侧方停车），0-表示未知，1-未跨泊车位停车，2-跨泊车位停车
	ST_byAngledParking         BYTE                     //是否倾斜停车（侧方停车）, 0-表示未知，1-未倾斜停车，2-倾斜停车
	ST_byAlarmValidity         BYTE                     //报警置信度，可以输出驶入驶出的置信度，范围0-100；置信度越高，事件真实性越高
	ST_byDoorsStatus           BYTE                     //车门状态 0-车门关闭 1-车门开启
	ST_dwXmlLen                DWORD                    //XML报警信息长度
	ST_pXmlBuf                 *byte                    // XML报警信息指针,其XML对应到EventNotificationAlert XML Block
	ST_byVehicleHeadTailStatus BYTE                     //车头车尾状态 0-保留 1-车头 2-车尾
	ST_byRes                   [31]BYTE                 //保留
}

type NET_ITS_PLATE_RESULT struct {
	ST_dwSize         DWORD // 结构长度
	ST_dwMatchNo      DWORD // 匹配序号,由(车辆序号,数据类型,车道号)组成匹配码
	ST_byGroupNum     BYTE  // 图片组数量（一辆过车相机多次抓拍的数量，代表一组图片的总数，用于延时匹配数据）
	ST_byPicNo        BYTE  // 连拍的图片序号（接收到图片组数量后，表示接收完成;接收超时不足图片组数量时，根据需要保留或删除）
	ST_bySecondCam    BYTE  // 是否第二相机抓拍（如远近景抓拍的远景相机，或前后抓拍的后相机，特殊项目中会用到）
	ST_byFeaturePicNo BYTE  // 闯红灯电警，取第几张图作为特写图,0xff-表示不取
	ST_byDriveChan    BYTE  // 触发车道号
	ST_byVehicleType  BYTE  // 车辆类型，参考VTR_RESULT
	ST_byDetSceneID   BYTE  // 检测场景号[1,4], IPC默认是0
	//车辆属性，按位表示，0- 无附加属性(普通车)，bit1- 黄标车(类似年检的标志)，bit2- 危险品车辆，值：0- 否，1- 是
	//该节点已不再使用,使用下面的byYellowLabelCar和byDangerousVehicles判断是否黄标车和危险品车
	ST_byVehicleAttribute BYTE
	ST_wIllegalType       WORD    // 违章类型采用国标定义
	ST_byIllegalSubType   [8]BYTE // 违章子类型
	ST_byPostPicNo        BYTE    // 违章时取第几张图片作为卡口图,0xff-表示不取
	//通道号(有效，报警通道号和所在设备上传报警通道号一致，在后端和所接入的 通道号一致)
	ST_byChanIndex   BYTE
	ST_wSpeedLimit   WORD // 限速上限（超速时有效）
	ST_byChanIndexEx BYTE // byChanIndexEx*256+byChanIndex表示真实通道号。
	/*车辆方位布控标志,0~为普通车牌识别报警,
	  1~为车辆位置布控触发报警(即通过PUT /ISAPI/Traffic/channels/<ID>/vehiclePositionControl?format=json触发)。
	  2~为车辆智能布控触发报警(包含多预置点及GPS车辆布控)(即通过PUT /ISAPI/Traffic/channels/<ID>/vehicleMonitor/<taskID>/startTask触发,
	  但在unarmedVehicleDetectionEnable（非目标车辆检测使能）为true时，只上报车辆检测报警，无车辆智能布控报警，因此该字段取值为0，而不是2)。
	  3~为手动车辆布控触发报警(/ISAPI/Traffic/channels/<ID>/manualVehicleMonitor?format=json)
	  4~为日常布控(复用车辆检测配置)可通过车辆检测能力中区分是否支持日常车辆布控,
	      (即通过/ISAPI/Traffic/channels/<ID>/vehicleDetect/capabilities中isSupportDailyVehicleMonitor判断,日常车辆布控时,设备不仅会上报ANPR报警还会实时上报vehicleMonitor车辆布控报警)
	      若不返回该字段则代表为普通车辆检测
	*/
	ST_byVehiclePositionControl BYTE
	ST_struPlateInfo            NET_DVR_PLATE_INFO   // 车牌信息结构
	ST_struVehicleInfo          NET_DVR_VEHICLE_INFO // 车辆信息
	ST_byMonitoringSiteID       [48]BYTE             // 监测点编号
	ST_byDeviceID               [48]BYTE             // 设备编号
	ST_byDir                    BYTE                 // 监测方向，1-上行（反向），2-下行(正向)，3-双向，4-由东向西，5-由南向北,6-由西向东，7-由北向南，8-其它
	ST_byDetectType             BYTE                 // 检测方式,1-地感触发，2-视频触发，3-多帧识别，4-雷达触发
	//关联车道方向类型，参考ITC_RELA_LANE_DIRECTION_TYPE
	//该参数为车道方向参数，与关联车道号对应，确保车道唯一性。
	ST_byRelaLaneDirectionType BYTE
	ST_byCarDirectionType      BYTE // 车辆具体行驶的方向，0表示从上往下，1表示从下往上（根据实际车辆的行驶方向来的区分）,2表示未知
	//当wIllegalType参数为空时，使用该参数。若wIllegalType参数为有值时，以wIllegalType参数为准，该参数无效。
	ST_dwCustomIllegalType DWORD // 违章类型定义(用户自定义)
	/*为0~数字格式时，为老的违章类型，wIllegalType、dwCustomIllegalType参数生效，赋值国标违法代码。
	  为1~字符格式时，pIllegalInfoBuf参数生效。老的违章类型，wIllegalType、dwCustomIllegalType参数依然赋值国标违法代码*/
	ST_pIllegalInfoBuf     *BYTE //违法代码字符信息结构体指针；指向NET_ITS_ILLEGAL_INFO
	ST_byIllegalFromatType BYTE  // 违章信息格式类型； 0~数字格式， 1~字符格式
	ST_byPendant           BYTE  //  0-表示未知,1-车窗有悬挂物，2-车窗无悬挂物
	ST_byDataAnalysis      BYTE  // 0-数据未分析, 1-数据已分析
	ST_byYellowLabelCar    BYTE  // 0-表示未知, 1-非黄标车,2-黄标车
	ST_byDangerousVehicles BYTE  // 0-表示未知, 1-非危险品车,2-危险品车
	//以下字段包含Pilot字符均为主驾驶，含Copilot字符均为副驾驶
	ST_byPilotSafebelt   BYTE // 0-表示未知,1-系安全带,2-不系安全带
	ST_byCopilotSafebelt BYTE // 0-表示未知,1-系安全带,2-不系安全带
	ST_byPilotSunVisor   BYTE // 0-表示未知,1-不打开遮阳板,2-打开遮阳板
	ST_byCopilotSunVisor BYTE // 0-表示未知, 1-不打开遮阳板,2-打开遮阳板
	ST_byPilotCall       BYTE //  0-表示未知, 1-不打电话,2-打电话
	//0-开闸，1-未开闸 (专用于历史数据中相机根据黑白名单匹配后，是否开闸成功的标志)；当byAlarmDataType为0-实时数据时 0-未开闸 1-开闸
	byBarrierGateCtrlType   BYTE
	ST_byAlarmDataType      BYTE                    // 0-实时数据，1-历史数据
	ST_struSnapFirstPicTime NET_DVR_TIME_V30        // 端点时间(ms)（抓拍第一张图片的时间）
	ST_dwIllegalTime        DWORD                   // 违法持续时间（ms） = 抓拍最后一张图片的时间 - 抓拍第一张图片的时间
	ST_dwPicNum             DWORD                   // 图片数量（与picGroupNum不同，代表本条信息附带的图片数量，图片信息由struVehicleInfoEx定义
	ST_struPicInfo          [6]NET_ITS_PICTURE_INFO // 图片信息,单张回调，最多6张图，由序号区分
}

type NET_DVR_PLATE_RESULT struct {
	ST_dwSize         DWORD    // 结构长度
	ST_byResultType   BYTE     // 0-视频识别结果，1图像识别结果 2 大于10M时走下载路线
	ST_byChanIndex    BYTE     // 通道号
	ST_wAlarmRecordID WORD     // 报警录像ID(用于查询录像，仅当byResultType为2时有效)
	ST_dwRelativeTime DWORD    // 相对时间点
	ST_byAbsTime      [32]BYTE // 绝对时间点,yyyymmddhhmmssxxx,e.g.20090810235959999（毫秒）
	ST_dwPicLen       DWORD    // 图片长度(近景图)
	ST_dwPicPlateLen  DWORD    // 车牌小图片长度
	ST_dwVideoLen     DWORD    // 录像内容长度
	ST_byTrafficLight BYTE     // 0-非红绿灯抓拍，1-绿灯时抓拍；2-红灯时抓拍
	ST_byPicNum       BYTE     // 连拍的图片序号
	ST_byDriveChan    BYTE     // 触发车道号
	ST_byVehicleType  BYTE     // 车辆类型，参考VTR_RESULT
	ST_dwBinPicLen    DWORD    // 车牌二值图长度
	ST_dwCarPicLen    DWORD    // 车辆原图长度
	ST_dwFarCarPicLen DWORD    // 远景图长度
	ST_pBuffer3       *BYTE    // 车牌二值图
	ST_pBuffer4       *BYTE    // 车辆原图
	ST_pBuffer5       *BYTE    // 远景图
	//关联车道方向类型，参考ITC_RELA_LANE_DIRECTION_TYPE
	//该参数为车道方向参数，与关联车道号对应，确保车道唯一性。
	ST_byRelaLaneDirectionType BYTE
	ST_byCarDirectionType      BYTE // 车辆具体行驶的方向，0表示从上往下，1表示从下往上（根据实际车辆的行驶方向来的区分）,2表示未知
	ST_byRes3                  [6]BYTE
	ST_struPlateInfo           NET_DVR_PLATE_INFO   // 车牌信息结构
	ST_struVehicleInfo         NET_DVR_VEHICLE_INFO // 车辆信息
	ST_pBuffer1                *BYTE                //  当上传的是图片(近景图)，指针指向图片信息，当上传的是视频，指针指向视频信息，如果不想获得图片或视频信息，传NULL(DVS车辆近景图)
	ST_pBuffer2                *BYTE                //  当上传的是图片(车牌图)时，指向车牌图片的指针（DVS车牌彩图）
}

//后面紧跟图片数据和录像数据，只传一种，图片数据为场景图片+车牌小图片

type NET_ITS_ILLEGAL_INFO struct {
	ST_byIllegalInfo [MAX_ILLEGAL_LEN] /*64*/ BYTE // 违章类型信息（字符格式）
	ST_byRes         [256]BYTE
}

type NET_DVR_ALARMINFO_FIXED_HEADER struct {
	/*0-信号量报警,1-硬盘满,2-信号丢失，3－移动侦测，4－硬盘未格式化,5-写硬盘出错,6-遮挡报警，
	  7-制式不匹配, 8-非法访问，9-视频信号异常，10-录像异常，11-智能场景变化，12-阵列异常，13-前端/录像分辨率不匹配,
	  14-申请解码资源失败,15-智能侦测报警, 16-热备异常，17-录播主机报警，18-语音对讲请求报警,19-音频丢失，20-开启录像，
	  21-关闭录像，22-车辆检测算法异常，23-脉冲报警,24-人脸库硬盘异常,25-人脸库变更,26-人脸库图片变更,27-POC异常,28-相机视角异常，
	  30-缺少SD卡,32-云台堵转异常*/
	ST_dwAlarmType      DWORD           //报警类型
	ST_struAlarmTime    NET_DVR_TIME_EX //发生报警的时间
	ST_byUnionLen       [116]BYTE       //分出去8个字节用于扩展时区
	ST_pRes             *DWORD          //用于兼容64位下结构体字节不对齐问题
	ST_byTimeDiffFlag   BYTE            /*时差字段是否有效  0-时差无效， 1-时差有效 */
	ST_cTimeDifferenceH BYTE            /*与UTC的时差（小时），-12 ... +14， +表示东区,，byTimeDiffFlag为1时有效*/
	ST_cTimeDifferenceM BYTE            /*与UTC的时差（分钟），-30, 30, 45， +表示东区，byTimeDiffFlag为1时有效*/
	ST_byRes2           [5]BYTE         //保留
}

// type NET_DVR_ALARMINFO_V40 struct {
// 	ST_struAlarmFixedHeader NET_DVR_ALRAM_FIXED_HEADER //报警固定部分
// 	ST_pAlarmData           *DWORD                     //报警可变部分内容
// }

type NET_ITC_LANE_LOGIC_PARAM struct {
	ST_byUseageType     BYTE     //车道用途类型，详见ITC_LANE_USEAGE_TYPE
	ST_byDirectionType  BYTE     //车道方向类型，详见ITC_LANE_DIRECTION_TYPE
	ST_byCarDriveDirect BYTE     //车辆行驶方向，详见ITC_LANE_CAR_DRIVE_DIRECT
	ST_byRes            [33]BYTE //保留
}

type NET_ITC_POLYGON struct {
	ST_dwPointNum DWORD                                    //有效点 大于等于3，若是3点在一条线上认为是无效区域，线交叉认为是无效区域
	ST_struPos    [ITC_MAX_POLYGON_POINT_NUM]NET_VCA_POINT //多边形边界点,最多20个
}

type NET_VCA_LINE struct {
	ST_struStart NET_VCA_POINT //起点
	ST_struEnd   NET_VCA_POINT //终点
}

type NET_ITC_LINE struct {
	ST_struLine   NET_VCA_LINE //线参数
	ST_byLineType BYTE         //线类型，详见ITC_LINE_TYPE
	ST_byRes      [7]BYTE
}

type NET_ITC_PLATE_RECOG_PARAM struct {
	ST_byDefaultCHN [MAX_CHJC_NUM]BYTE /*设备运行省份的汉字简写*/
	ST_byEnable     BYTE               //是否启用该区域牌识，0-否，1-是
	ST_dwRecogMode  DWORD              /*识别的类型，
	  bit0-背向识别：0-正向车牌识别，1-背向识别(尾牌识别) ；
	  bit1-大车牌识别或小车牌识别：0-小车牌识别，1-大车牌识别 ；
	  bit2-车身颜色识别：0-不采用车身颜色识别，在背向识别或小车牌识别时禁止启用，1-车身颜色识别；
	  bit3-农用车识别：0-不采用农用车识别，1-农用车识别；
	  bit4-模糊识别：0-不采用模糊识别，1-模糊识别；
	  bit5-帧定位或场定位：0-帧定位，1-场定位；
	  bit6-帧识别或场识别：0-帧识别，1-场识别；
	  bit7-晚上或白天：0-白天，1-晚上
	  bit8-摩托车识别：0-不采用摩托车识别，1-摩托车识别;
	  bit9-场景模式：0-电警/多帧，1-卡口；
	  bit10-微小车牌：0-不启用，1-启用微小车牌识别(像素60～80)
	  bit11-安全带检测：0-不启用，1-启用安全带检测
	  bit12-民航车牌识别: 0-不启用，1-开启民航车牌识别
	  bit13-车牌过渡倾斜处理: 0-不启用，1-开启过渡倾斜处理（PRS）
	  bit14-超大车牌识别: 0-不启用，1-开启超大车牌识别（PRS）
	  bit15-遮阳板检测：0-不启用，1-启用遮阳板检测
	  bit16-黄标车检测：0-不启用，1-启用黄标车检测
	  bit17-危险品车辆检测：0-不启用，1-启用危险品车辆检测
	  bit18-使馆车牌识别：0-不启用，1-启用使馆车牌识别
	  bit19-车辆子品牌识别：0-不启用，1-启用车辆子品牌识别
	  bit20-打电话识别：0-不启用，1-启用
	  bit21-车窗悬挂物识别：0-不启用，1-启用
	*/
	ST_byVehicleLogoRecog BYTE //车标识别 0-不启用，1-启用
	/*
	   0-保留，1-澳，2-京，3-渝，4-闽，5-甘，6-粤，7-桂，8-贵，9-琼，10-冀，11-豫，
	   12-黑，13-鄂，14-湘，15-吉，16-苏，17-赣，18-辽，19-蒙，20-宁，21-青，22-鲁，
	   23-晋，24-陕，25-沪，26-川，27-台，28-津，29-藏，30-港，31-新，32-云，33-浙，
	   34-皖，0xff-全部
	*/
	ST_byProvince          BYTE //省份索引值
	ST_byRegion            BYTE // 区域索引值 0-保留，1-欧洲，2-俄语区域, 3-欧洲&俄罗斯(EU&CIS),4-中东(Middle East)
	ST_byCountry           BYTE //国家索引，参照枚举COUNTRY_INDEX(不支持“COUNTRY_ALL = 0xff,//ALL 全部”)
	ST_wPlatePixelWidthMin WORD //车牌像素识别宽度最小值（单位是像素）当前推荐范围[130,500]
	ST_wPlatePixelWidthMax WORD //车牌像素识别宽度最大值（单位是像素）当前推荐范围[130,500]
	ST_byRes               [24]BYTE
}

type NET_ITC_INTERVAL_PARAM struct {
	ST_byIntervalType BYTE //间隔类型（默认按时间），0-时间起效,1-距离起效
	ST_byRes1         [3]BYTE
	ST_wInterval      [MAX_INTERVAL_NUM]WORD //连拍间隔时间（单位ms）或连拍间隔距离（单位分米），当byIntervalType为0时，表示间隔时间，当byIntervalType为1时，表示距离
	ST_byRes          [8]BYTE
}

type NET_ITC_VIOLATION_DETECT_LINE struct {
	struLaneLine     NET_ITC_LINE //车道线参数
	struStopLine     NET_ITC_LINE //停止线参数
	struRedLightLine NET_ITC_LINE //闯红灯触发线参数
	struCancelLine   NET_ITC_LINE //直行触发位置取消线
	struWaitLine     NET_ITC_LINE //待行区停止线参数
	struRes          [8]NET_ITC_LINE
}

type NET_ITC_POST_MOBILE_PARAM struct {
	ST_byEnable    BYTE
	ST_bySceneMode BYTE //场景模式 0-高速公路 1-城市道路
	/*抓拍类型
	  bit0-卡口,bit1-大车占道,bit2-压硬路肩
	*/
	ST_wExpressWayCapType WORD //高速公路
	/*抓拍类型
	  bit0-卡口,bit1-机占非,bit2-占用专用车道
	*/
	ST_wUrbanRoadCapType    WORD                                                            //城市道路
	ST_byCapNum             BYTE                                                            //抓拍张数 [2,3]
	ST_byRecordEnable       BYTE                                                            //违章录像使能 0-关闭，1-开启
	ST_dwPreRecordTime      DWORD                                                           //录像预录时间(s)
	ST_dwOverRecordTime     DWORD                                                           //录像超时时间(s)
	ST_struLane             NET_ITC_LANE_LOGIC_PARAM                                        //车道属性
	ST_struPolygon          [MAX_MOBILE_POLYGON_NUM] /*3*/ NET_ITC_POLYGON                  //牌识区域参数
	ST_struLine             [MAX_MOBILE_DETECTLINE_NUM] /*3*/ NET_ITC_VIOLATION_DETECT_LINE //违规检测线
	ST_struLaneBoundaryLine NET_ITC_LINE                                                    //车道边界线（最右边车道的右车道线）
	ST_struPlateRecog       NET_ITC_PLATE_RECOG_PARAM                                       //牌识参数
	ST_struInterval         NET_ITC_INTERVAL_PARAM                                          //抓拍间隔参数（20byte）
	ST_byRes                [256]BYTE
}

type NET_ITC_VIOLATION_DETECT_PARAM struct {
	ST_dwVioDetectType       DWORD //违规检测类型, 按位表示, 详见ITC_VIOLATION_DETECT_TYPE ,0-不启用,1-启用
	ST_byDriveLineSnapTimes  BYTE  //压车道线抓拍张数,2-3
	ST_byReverseSnapTimes    BYTE  //逆行抓拍,2-3
	ST_wStayTime             WORD  //机占非停留时间（该时间后抓拍），单位s
	ST_byNonDriveSnapTimes   BYTE  //机占非抓拍张数2-3
	ST_byChangeLaneTimes     BYTE  //违法变道抓拍张数 2-3
	ST_bybanTimes            BYTE  //违法禁令抓拍张数2-3
	ST_byDriveLineSnapSen    BYTE  // 压线灵敏度(0~100)(3.7Ver)
	ST_wSnapPosFixPixel      WORD  //第2,3张抓拍位置最小偏移(违反信号灯时起效)（单位：像素） 命名需改进
	ST_bySpeedTimes          BYTE  //违法超速抓拍张数2-3(3.8Ver)
	ST_byTurnAroundEnable    BYTE  //违章掉头使能 0~关闭 1~开启
	ST_byThirdPlateRecogTime BYTE  //第三张牌识时间 0~180s
	ST_byPostSnapTimes       BYTE  //卡口抓拍张数,1-2张
	ST_byRes1                [18]BYTE
	ST_wStopLineDis          WORD //电警第2张违规图片与停止线的最短距离，[0,300]单位(像素)
	ST_byRes                 [14]BYTE
}

type NET_IPC_LANE_HVT_PARAM struct {
	ST_byLaneNO         BYTE
	ST_byCarDriveDirect BYTE
	//车辆行驶方向 ITC_LANE_CAR_DRIVE_DIRECT
	ST_byRes        [62]BYTE
	ST_struLaneLine NET_ITC_LINE
	//车道线
	ST_struPlateRecog NET_ITC_POLYGON
	//牌识区域
	ST_byRes1 [256]BYTE
}

type NET_ITC_PARKING_DETECTION struct {
	ST_byEnable           BYTE     //是否启用，0-不启用，1-启用
	ST_byRes              BYTE     //保留字节
	ST_wDuration          WORD     //检测时间间隔，单位：s
	ST_wAlarmIntervalTime WORD     //上传时间间隔，单位：s
	ST_byRes1             [58]BYTE //保留字节
}

// NET_IPC_POST_HVT_PARAM 卡口混合模式
type NET_IPC_POST_HVT_PARAM struct {
	ST_byEnable             BYTE
	ST_byLaneNum            BYTE
	ST_byEnhancedMode       BYTE //增强模式，0-不启用，1-启用
	ST_byPicRecognition     BYTE //车辆报警图片二次识别，0-不启用，1-启用
	ST_byRes                [60]BYTE
	ST_struLaneBoundaryLine NET_ITC_LINE              //车道边界线（最左边车道的左边界线）
	ST_struPlateRecog       NET_ITC_PLATE_RECOG_PARAM //牌识参数  40
	ST_struLaneParam        [MAX_ITC_LANE_NUM] /*6*/ NET_IPC_LANE_HVT_PARAM
	ST_szSceneName          [NAME_LEN] /*32*/ BYTE    //场景名称
	ST_struSnapLine         NET_VCA_LINE              //抓拍线，抓在相机的架设方式为侧装，配置这个参数生效（即4.5中NET_DVR_CAMERA_SETUPCFG结构体中byErectMethod字段为1时 ）
	ST_struParkingDetection NET_ITC_PARKING_DETECTION //停车检测参数
	ST_byRes1               [328]BYTE
}

type NET_DVR_GEOGLOCATION struct {
	iRes   [2]int /*保留*/
	dwCity DWORD  /*城市，详见PROVINCE_CITY_IDX */
}

//抓拍机4.0新增
type NET_ITC_LANE_HVT_PARAM_V50 struct {
	ST_byLaneNO              BYTE //关联的车道号1～255(用于叠加和上传)
	ST_byFlashMode           BYTE //闪光灯闪烁模式，0-同时闪，1-轮流闪
	ST_bySignSpeed           BYTE //小车标志限高速，单位km/h
	ST_bySpeedLimit          BYTE //小车限高速值，单位km/h
	ST_bySignLowSpeed        BYTE //小车标志限底速，单位km/h
	ST_byLowSpeedLimit       BYTE //小车限底速值，单位km/h
	ST_byBigCarSignSpeed     BYTE //大车标志限高速，单位km/h（新交规）
	ST_byBigCarSpeedLimit    BYTE //大车限高速值，单位km/h（新交规）
	ST_byBigCarSignLowSpeed  BYTE //大车标志限底速，单位km/h
	ST_byBigCarLowSpeedLimit BYTE //大车限底速值，单位km/h
	ST_bySnapTimes           BYTE //卡口抓拍张数，1~3
	ST_byDriveLineSnapTime   BYTE // 压线抓拍张数 1~3
	ST_byHighSpeedSnapTime   BYTE // 超高速抓拍张数1~3
	ST_byLowSpeedSnapTime    BYTE // 超低速抓拍张数1~3
	ST_byBanSnapTime         BYTE // 违反禁令抓拍张数  1~3
	ST_byReverseSnapTime     BYTE //逆行抓拍张数  1~3
	ST_byRelatedDriveWay     BYTE //关联车道号，用于匹配车检器
	ST_byLaneType            BYTE //车道类型，0-未配置、1-高速公路、2-城市快速路、0xff-其他道路
	//关联车道方向类型，参考ITC_RELA_LANE_DIRECTION_TYPE
	//该参数为车道方向参数，与关联车道号byRelatedDriveWay对应，确保车道唯一性。
	ST_byRelaLaneDirectionType BYTE
	ST_byRes1                  [27]BYTE
	ST_byChangeLaneEnable      BYTE //违章变道抓拍使能，0-关闭，1-开启
	ST_byChangeLaneCapNo       BYTE //违章变道抓拍张数2-3
	//目前仅使用第一个车道的，以后可能会扩展为多车道分别配置
	//类型, 按位表示,0-不启用,1-启用参考 ITC_VIOLATION_DETECT_TYPE
	ST_dwVioDetectType DWORD
	ST_dwRelatedIOOut  DWORD                    //关联的IO输出口(可以同时关联多个)，按位表示IO输出口，第0位表示IO输出口1，以此类推，0-不关联，1-关联
	ST_struTrigLine    NET_ITC_LINE             //触发线，目前仅使用第一个车道的，以后可能会扩展为多车道分别配置
	ST_struLineLeft    NET_ITC_LINE             //左车道线
	ST_struPlateRecog  NET_ITC_POLYGON          //牌识区域
	ST_struLane        NET_ITC_LANE_LOGIC_PARAM //车道属性，用byUseageType和byCarDriveDirect
	ST_struInterval    NET_ITC_INTERVAL_PARAM   //抓拍间隔参数（20byte）
	ST_byRes2          [280]BYTE
}

// NET_ITC_POST_HVT_PARAM_V50 卡口混合模式V50
type NET_ITC_POST_HVT_PARAM_V50 struct {
	ST_byLaneNum            BYTE //识别的车道个数，1-6
	ST_byCapType            BYTE //抓拍类型，0-机、非、人（默认），1-机动车
	ST_byCapMode            BYTE //抓拍方式，0-视频抽帧，1-打断抓拍，2-混合模式，
	ST_bySecneMode          BYTE //场景模式，0-城区道路（默认），1-小区出入口，2-高速公路
	ST_bySpeedMode          BYTE //测速模式，0-无测速，1-雷达测速，2-视频测速
	ST_byLineRuleEffect     BYTE //触发规则线有效性,每一位代表一条触发线,0-无效;1-有效。bit0-左触发线;bit1-右触发线;bit2-视频检测区域
	ST_byRes1               [78]BYTE
	ST_struLeftTrigLine     NET_ITC_LINE                                        //左触发线(一条垂直线)
	ST_struRigtTrigLine     NET_ITC_LINE                                        //右触发线(一条垂直线)
	ST_struLaneBoundaryLine NET_ITC_LINE                                        //车道边界线（最右边车道的右车道线）
	ST_struDetectArea       NET_ITC_POLYGON                                     //视频检测区域
	ST_struGeogLocation     NET_DVR_GEOGLOCATION                                //地理位置（默认浙江省）计算时区
	ST_struLaneParam        [MAX_ITC_LANE_NUM] /*6*/ NET_ITC_LANE_HVT_PARAM_V50 //单车道属性
	ST_struPlateRecog       NET_ITC_PLATE_RECOG_PARAM                           //牌识参数
	ST_byRes2               [260]BYTE
}

type NET_ITC_TRIGGER_PARAM_UNION struct {
	uLen [1070]DWORD
	// NET_ITC_POST_IOSPEED_PARAM        struIOSpeed      //（卡口）IO测速参数
	// NET_ITC_POST_SINGLEIO_PARAM       struSingleIO     //（卡口）单IO参数
	// NET_ITC_POST_RS485_PARAM          struPostRs485    //（卡口）RS485车检器参数
	// NET_ITC_POST_RS485_RADAR_PARAM    struPostRadar    //（卡口）RS485雷达参数
	// NET_ITC_POST_VTCOIL_PARAM         struVtCoil       //（卡口）虚拟线圈参数
	// NET_ITC_POST_HVT_PARAM            struHvt          //（卡口）混行卡口参数
	// NET_ITC_EPOLICE_IOTL_PARAM        struIOTL         //（电警）IO红绿灯参数
	// NET_ITC_EPOLICE_RS485_PARAM       struEpoliceRs485 //（电警）RS485车检器触发参数
	// NET_ITC_EPOLICE_RS485_PARAM       struPERs485      //（卡式电警）RS485车检器触发参数
	// NET_ITC_POST_MPR_PARAM            struPostMpr      //多帧检测触发（MPR）
	// NET_DVR_VIA_VTCOIL_PARAM          struViaVtCoil    //(VIA)视频检测参数
	// NET_ITC_POST_IMT_PARAM            struPostImt      //智慧监控触发
	// NET_ITC_POST_PRS_PARAM            struPostPrs      //视频检测触发
	// NET_IPC_POST_HVT_PARAM            struIpcHvt       //(IPC) 混行卡口参数
	// NET_ITC_POST_HVT_PARAM_V50        struHvtV50       /*（卡口）混行卡口参数V50*/
	// NET_ITC_POST_MOBILE_PARAM         struPostMobile   // 移动交通触发模式
	// NET_ITC_NOCOMITY_PEDESTRIAN_PARAM struNoComityPed  //不礼让行人参数
	// NET_ITC_REDLIGHT_PEDESTRIAN_PARAM struRedLightPed  //行人闯红灯参数
}

type NET_ITC_SINGLE_TRIGGERCFG struct {
	ST_byEnable      BYTE //是否启用该触发模式，0-否，1-是
	ST_byRes1        [3]BYTE
	ST_dwTriggerType DWORD                       //触发类型，详见ITC_TRIGGERMODE_TYPE
	ST_uTriggerParam NET_ITC_TRIGGER_PARAM_UNION //触发参数
	ST_byRes         [64]BYTE
}

type NET_ITC_TRIGGERCFG struct {
	ST_dwSize           DWORD                     //结构长度
	ST_struTriggerParam NET_ITC_SINGLE_TRIGGERCFG //单个触发参数
	ST_byRes            [32]BYTE
}

func (cfg *NET_ITC_SINGLE_TRIGGERCFG) NET_ITC_POST_MOBILE_PARAM() *NET_ITC_POST_MOBILE_PARAM {
	ptr := unsafe.Pointer(&cfg.ST_uTriggerParam.uLen[0])

	var ncfg NET_ITC_POST_MOBILE_PARAM = *(*NET_ITC_POST_MOBILE_PARAM)(ptr)
	return &ncfg
}

func (cfg *NET_ITC_SINGLE_TRIGGERCFG) NET_ITC_VIOLATION_DETECT_PARAM() *NET_ITC_VIOLATION_DETECT_PARAM {
	ptr := unsafe.Pointer(&cfg.ST_uTriggerParam.uLen[0])

	var ncfg NET_ITC_VIOLATION_DETECT_PARAM = *(*NET_ITC_VIOLATION_DETECT_PARAM)(ptr)
	return &ncfg
}

func (cfg *NET_ITC_SINGLE_TRIGGERCFG) NET_ITC_POST_HVT_PARAM_V50() *NET_ITC_POST_HVT_PARAM_V50 {
	ptr := unsafe.Pointer(&cfg.ST_uTriggerParam.uLen[0])

	var ncfg NET_ITC_POST_HVT_PARAM_V50 = *(*NET_ITC_POST_HVT_PARAM_V50)(ptr)
	return &ncfg
}
