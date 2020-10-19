package netsdk

const (
	MAX_NAMELEN           = 16 //DVR本地登陆名
	MAX_RIGHT             = 32 //设备支持的权限（1-12表示本地权限，13-32表示远程权限）
	MAX_CALIB_PT          = 6
	NAME_LEN              = 32  //用户名长度
	MIN_PASSWD_LEN        = 8   //最小密码长度
	PASSWD_LEN            = 16  //密码长度
	STREAM_PASSWD_LEN     = 12  //码流加密密钥最大长度
	MAX_PASSWD_LEN_EX     = 64  //密码长度64位
	GUID_LEN              = 16  //GUID长度
	DEV_TYPE_NAME_LEN     = 24  //设备类型名称长度
	SERIALNO_LEN          = 48  //序列号长度
	MACADDR_LEN           = 6   //mac地址长度
	MAC_ADDRESS_NUM       = 48  //Mac地址长度
	MAX_SENCE_NUM         = 16  //场景数
	RULE_REGION_MAX       = 128 //最大区域
	MAX_ETHERNET          = 2   //设备可配以太网络
	MAX_NETWORK_CARD      = 4   //设备可配最大网卡数目
	MAX_NETWORK_CARD_EX   = 12  //设备可配最大网卡数目扩展
	PATHNAME_LEN          = 128 //路径长度
	MAX_PRESET_V13        = 16  //预置点
	MAX_TEST_COMMAND_NUM  = 32  //产线测试保留字段长度
	MAX_NUMBER_LEN        = 32  //号码最大长度
	MAX_NAME_LEN          = 128 //设备名称最大长度
	MAX_INDEX_LED         = 8   //LED索引最大值 2013-11-19
	MAX_CUSTOM_DIR        = 64  //自定义目录最大长度
	URL_LEN_V40           = 256 //最大URL长度
	CLOUD_NAME_LEN        = 48  //云存储服务器用户名长度
	CLOUD_PASSWD_LEN      = 48  //云存储服务器密码长度
	MAX_SENSORNAME_LEN    = 64  //传感器名称长度
	MAX_SENSORCHAN_LEN    = 32  //传感器通道长度
	MAX_DESCRIPTION_LEN   = 32  //传感器描述长度
	MAX_DEVNAME_LEN_EX    = 64  //设备名称长度扩展
	NET_SDK_MAX_FILE_PATH = 256 //文件路径长度
	MAX_TMEVOICE_LEN      = 64  //TME语音播报内容长度
	ISO_8601_LEN          = 32  //ISO_8601时间长度
	MODULE_INFO_LEN       = 32  //模块信息长度
	VERSION_INFO_LEN      = 32  //版本信息长度

	MAX_NUM_INPUT_BOARD      = 512 //输入板最大个数
	MAX_SHIPSDETE_REGION_NUM = 8   // 船只检测区域列表最大数目

	MAX_RES_NUM_ONE_VS_INPUT_CHAN = 8  //一个虚拟屏输入通道支持的分辨率的最大数量
	MAX_VS_INPUT_CHAN_NUM         = 16 //虚拟屏输入通道最大数量

	NET_SDK_MAX_FDID_LEN          = 256 //人脸库ID最大长度
	NET_SDK_MAX_PICID_LEN         = 256 //人脸ID最大长度
	NET_SDK_FDPIC_CUSTOM_INFO_LEN = 96  //人脸库图片自定义信息长度
	NET_DVR_MAX_FACE_ANALYSIS_NUM = 32  //最大支持单张图片识别出的人脸区域个数
	NET_DVR_MAX_FACE_SEARCH_NUM   = 5   //最大支持搜索人脸区域个数
	NET_SDK_SECRETKEY_LEN         = 128 //配置文件密钥长度
	NET_SDK_CUSTOM_LEN            = 512 //自定义信息最大长度
	NET_SDK_CHECK_CODE_LEN        = 128 //校验码长度
	RELATIVE_CHANNEL_LEN          = 2   //报警关联的通道号的数量
	NET_SDK_MAX_CALLEDTARGET_NAME = 32  //呗呼叫目标的用户名
	NET_SDK_MAX_HBDID_LEN         = 256 /*256 人体库ID最大长度*/
	//小间距LED控制器
	MAX_LEN_TEXT_CONTENT      = 128 //字符内容长度
	MAX_NUM_INPUT_SOURCE_TEXT = 32  //信号源可叠加的文本数量
	MAX_NUM_OUTPUT_CHANNEL    = 512 //LED区域包含的输出口个数

	//子窗口解码OSD
	MAX_LEN_OSD_CONTENT    = 256 //OSD信息最大长度
	MAX_NUM_OSD_ONE_SUBWND = 8   //单个子窗口支持的最大OSD数量
	MAX_NUM_SPLIT_WND      = 64  //单个窗口支持的最大分屏窗口数量（即子窗口数量）
	MAX_NUM_OSD            = 8   //2013-11-19
	MAX_DEVNAME_LEN        = 32  //设备名称最大长度
	MAX_LED_INFO           = 256 //屏幕字体显示信息最大长度
	MAX_TIME_LEN           = 32  //时间最大长度
	MAX_CARD_LEN           = 24  //卡号最大长度
	MAX_OPERATORNAME_LEN   = 32  //操作人员名称最大长度

	THERMOMETRY_ALARMRULE_NUM          = 40  //热成像报警规则数
	MAX_THERMOMETRY_REGION_NUM         = 40  //热度图检测区域最大支持数
	MAX_THERMOMETRY_DIFFCOMPARISON_NUM = 40  //热成像温差报警规则数
	MAX_SHIPS_NUM                      = 20  //船只检测最大船只数
	KEY_WORD_NUM                       = 3   //关键字个数
	KEY_WORD_LEN                       = 128 //关键字长度
	//异步登录回调状态宏定义
	ASYN_LOGIN_SUCC   = 1 //异步登录成功
	ASYN_LOGIN_FAILED = 0 //异步登录失败

	NET_SDK_MAX_VERIFICATION_CODE_LEN = 32  //萤石云验证码长度
	NET_SDK_MAX_OPERATE_CODE_LEN      = 64  //萤石云操作码长度
	MAX_TIMESEGMENT_V30               = 8   //9000设备最大时间段数
	MAX_TIMESEGMENT                   = 4   //8000设备最大时间段数
	MAX_ICR_NUM                       = 8   //抓拍机红外滤光片预置点数2013-07-09
	MAX_VEHICLEFLOW_INFO              = 24  //车流量信息最大个数
	MAX_SHELTERNUM                    = 4   //8000设备最大遮挡区域数
	MAX_DAYS                          = 7   //每周天数
	PHONENUMBER_LEN                   = 32  //pppoe拨号号码最大长度
	MAX_ACCESSORY_CARD                = 256 //配件板信息最大长度
	MAX_DISKNUM_V30                   = 33  //9000设备最大硬盘数/* 最多33个硬盘(包括16个内置SATA硬盘、1个eSATA硬盘和16个NFS盘) */

	NET_SDK_DISK_LOCATION_LEN = 16 //硬盘位置长度
	NET_SDK_SUPPLIER_NAME_LEN = 32 //供应商名称长度
	NET_SDK_DISK_MODEL_LEN    = 64 //硬盘型号长度
	NET_SDK_MAX_DISK_VOLUME   = 33 //最大硬盘卷个数
	NET_SDK_DISK_VOLUME_LEN   = 36 //硬盘卷名称长度

	MAX_DISKNUM             = 16 //8000设备最大硬盘数
	MAX_DISKNUM_V10         = 8  //1.2版本之前版本
	CARD_READER_DESCRIPTION = 32 //读卡器描述
	MAX_FACE_NUM            = 2  //最大人脸数

	MAX_WINDOW_V30 = 32 //9000设备本地显示最大播放窗口数
	MAX_WINDOW_V40 = 64 //Netra 2.3.1扩展
	MAX_WINDOW     = 16 //8000设备最大硬盘数
	MAX_VGA_V30    = 4  //9000设备最大可接VGA数
	MAX_VGA        = 1  //8000设备最大可接VGA数

	MAX_USERNUM_V30        = 32  //9000设备最大用户数
	MAX_USERNUM            = 16  //8000设备最大用户数
	MAX_EXCEPTIONNUM_V30   = 32  //9000设备最大异常处理数
	MAX_EXCEPTIONNUM       = 16  //8000设备最大异常处理数
	MAX_LINK               = 6   //8000设备单通道最大视频流连接数
	MAX_ITC_EXCEPTIONOUT   = 32  //抓拍机最大报警输出
	MAX_SCREEN_DISPLAY_LEN = 512 //屏幕显示字符长度

	MAX_DECPOOLNUM     = 4  //单路解码器每个解码通道最大可循环解码数
	MAX_DECNUM         = 4  //单路解码器的最大解码通道数（实际只有一个，其他三个保留）
	MAX_TRANSPARENTNUM = 2  //单路解码器可配置最大透明通道数
	MAX_CYCLE_CHAN     = 16 //单路解码器最大轮巡通道数
	MAX_CYCLE_CHAN_V30 = 64 //最大轮巡通道数（扩展）
	MAX_DIRNAME_LENGTH = 80 //最大目录长度
	MAX_WINDOWS        = 16 //最大窗口数

	MAX_STRINGNUM_V30        = 8    //9000设备最大OSD字符行数数
	MAX_STRINGNUM            = 4    //8000设备最大OSD字符行数数
	MAX_STRINGNUM_EX         = 8    //8000定制扩展
	MAX_AUXOUT_V30           = 16   //9000设备最大辅助输出数
	MAX_AUXOUT               = 4    //8000设备最大辅助输出数
	MAX_HD_GROUP             = 16   //9000设备最大硬盘组数
	MAX_HD_GROUP_V40         = 32   //设备最大硬盘组数
	MAX_NFS_DISK             = 8    //8000设备最大NFS硬盘数
	NET_SDK_VERSION_LIST_LEN = 64   //算法库版本最大值
	IW_ESSID_MAX_SIZE        = 32   //WIFI的SSID号长度
	IW_ENCODING_TOKEN_MAX    = 32   //WIFI密锁最大字节数
	MAX_SERIAL_NUM           = 64   //最多支持的透明通道路数
	MAX_DDNS_NUMS            = 10   //9000设备最大可配ddns数
	MAX_DOMAIN_NAME          = 64   // 最大域名长度
	MAX_EMAIL_ADDR_LEN       = 48   //最大email地址长度
	MAX_EMAIL_PWD_LEN        = 32   //最大email密码长度
	MAX_SLAVECAMERA_NUM      = 8    //从摄像机个数
	MAX_CALIB_NUM            = 6    //标定点的个数
	MAX_CALIB_NUM_EX         = 20   //扩展标定点的个数
	MAX_LEDDISPLAYINFO_LEN   = 1024 //最大LED屏显示长度
	MAX_PEOPLE_DETECTION_NUM = 8    //最大人员检测区域数
	MAXPROGRESS              = 100  //回放时的最大百分率
	MAX_SERIALNUM            = 2    //8000设备支持的串口数 1-232， 2-485
	CARDNUM_LEN              = 20   //卡号长度
	PATIENTID_LEN            = 64
	CARDNUM_LEN_OUT          = 32 //外部结构体卡号长度
	MAX_VIDEOOUT_V30         = 4  //9000设备的视频输出数
	MAX_VIDEOOUT             = 2  //8000设备的视频输出数

	MAX_PRESET_V30 = 256 /* 9000设备支持的云台预置点数 */
	MAX_TRACK_V30  = 256 /* 9000设备支持的云台轨迹数 */
	MAX_CRUISE_V30 = 256 /* 9000设备支持的云台巡航数 */
	MAX_PRESET     = 128 /* 8000设备支持的云台预置点数 */
	MAX_TRACK      = 128 /* 8000设备支持的云台轨迹数 */
	MAX_CRUISE     = 128 /* 8000设备支持的云台巡航数 */

	MAX_PRESET_V40          = 300 /* 云台支持的最大预置点数 */
	MAX_CRUISE_POINT_NUM    = 128 /* 最大支持的巡航点的个数 */
	MAX_CRUISEPOINT_NUM_V50 = 256 //最大支持的巡航点的个数扩展

	CRUISE_MAX_PRESET_NUMS = 32 /* 一条巡航最多的巡航点 */
	MAX_FACE_PIC_NUM       = 30 /*人脸子图个数*/
	LOCKGATE_TIME_NUM      = 4  //锁闸时间段个数

	MAX_SERIAL_PORT  = 8     //9000设备支持232串口数
	MAX_PREVIEW_MODE = 8     /* 设备支持最大预览模式数目 1画面,4画面,9画面,16画面.... */
	MAX_MATRIXOUT    = 16    /* 最大模拟矩阵输出个数 */
	LOG_INFO_LEN     = 11840 /* 日志附加信息 */
	DESC_LEN         = 16    /* 云台描述字符串长度 */
	PTZ_PROTOCOL_NUM = 200   /* 9000最大支持的云台协议数 */
	IPC_PROTOCOL_NUM = 50    //ipc 协议最大个数

	MAX_AUDIO     = 1  //8000语音对讲通道数
	MAX_AUDIO_V30 = 2  //9000语音对讲通道数
	MAX_CHANNUM   = 16 //8000设备最大通道数
	MAX_ALARMIN   = 16 //8000设备最大报警输入数
	MAX_ALARMOUT  = 4  //8000设备最大报警输出数
	//9000 IPC接入
	MAX_ANALOG_CHANNUM  = 32 //最大32个模拟通道
	MAX_ANALOG_ALARMOUT = 32 //最大32路模拟报警输出
	MAX_ANALOG_ALARMIN  = 32 //最大32路模拟报警输入

	MAX_IP_DEVICE       = 32   //允许接入的最大IP设备数
	MAX_IP_DEVICE_V40   = 64   // 允许接入的最大IP设备数 最多可添加64个 IVMS 2000等新设备
	MAX_IP_CHANNEL      = 32   //允许加入的最多IP通道数
	MAX_IP_ALARMIN      = 128  //允许加入的最多报警输入数
	MAX_IP_ALARMOUT     = 64   //允许加入的最多报警输出数
	MAX_IP_ALARMIN_V40  = 4096 //允许加入的最多报警输入数
	MAX_IP_ALARMOUT_V40 = 4096 //允许加入的最多报警输出数

	MAX_RECORD_FILE_NUM = 20 // 每次删除或者刻录的最大文件数
	//SDK_V31 ATM
	MAX_ACTION_TYPE      = 12   //自定义协议叠加交易行为最大行为个数
	MAX_ATM_PROTOCOL_NUM = 256  //每种输入方式对应的ATM最大协议数
	ATM_CUSTOM_PROTO     = 1025 //自定义协议 值为1025
	ATM_PROTOCOL_SORT    = 4    //ATM协议段数
	ATM_DESC_LEN         = 32   //ATM描述字符串长度
	// SDK_V31 ATM

	MAX_IPV6_LEN    = 64 //IPv6地址最大长度
	MAX_EVENTID_LEN = 64 //事件ID长度

	INVALID_VALUE_UINT32   = 0xffffffff //无效值
	MAX_CHANNUM_V40        = 512
	MAX_MULTI_AREA_NUM     = 24  //SDK 录播主机
	COURSE_NAME_LEN        = 32  //课程名称
	INSTRUCTOR_NAME_LEN    = 16  //授课教师
	COURSE_DESCRIPTION_LEN = 256 //课程信息

	MAX_TIMESEGMENT_V40 = 16 //每节课信息

	MAX_MIX_CHAN_NUM      = 16 /*目前支持的最大混音通道数，背景通道 + MIC + LINE IN + 最多4个小画面*/
	MAX_LINE_IN_CHAN_NUM  = 16 //最大line in通道数
	MAX_MIC_CHAN_NUM      = 16 //最大MIC通道数
	INQUEST_CASE_NO_LEN   = 64 //审讯案件编号长度
	INQUEST_CASE_NAME_LEN = 64 //审讯案件名称长度
	CUSTOM_INFO_LEN       = 64 //自定义信息长度
	INQUEST_CASE_LEN      = 64 //审讯信息长度

	MAX_FILE_ID_LEN  = 128 //视图库项目中文件ID的最大长度
	MAX_PIC_NAME_LEN = 128 //图片名称长度

	/* 最大支持的通道数 最大模拟加上最大IP支持 */
	MAX_CHANNUM_V30                  = (MAX_ANALOG_CHANNUM + MAX_IP_CHANNEL)       //64
	MAX_ALARMOUT_V40                 = (MAX_IP_ALARMOUT_V40 + MAX_ANALOG_ALARMOUT) //4128
	MAX_ALARMOUT_V30                 = (MAX_ANALOG_ALARMOUT + MAX_IP_ALARMOUT)     //96
	MAX_ALARMIN_V30                  = (MAX_ANALOG_ALARMIN + MAX_IP_ALARMIN)       //160
	MAX_ALARMIN_V40                  = (MAX_IP_ALARMIN_V40 + MAX_ANALOG_ALARMOUT)  //4128
	MAX_ANALOG_ALARM_WITH_VOLT_LIMIT = 16                                          //受电压限定的模拟报警最大输入数

	MAX_ROIDETECT_NUM     = 8  //支持的ROI区域数
	MAX_LANERECT_NUM      = 5  //最大车牌识别区域数
	MAX_FORTIFY_NUM       = 10 //最大布防个数
	MAX_INTERVAL_NUM      = 4  //最大时间间隔个数
	MAX_CHJC_NUM          = 3  //最大车辆省份简称字符个数
	MAX_VL_NUM            = 5  //最大虚拟线圈个数
	MAX_DRIVECHAN_NUM     = 16 //最大车道数
	MAX_COIL_NUM          = 3  //最大线圈个数
	MAX_SIGNALLIGHT_NUM   = 6  //最大信号灯个数
	LEN_16                = 16
	LEN_32                = 32
	LEN_64                = 64
	LEN_31                = 31
	MAX_LINKAGE_CHAN_NUM  = 16 //报警联动的通道的最大数量
	MAX_CABINET_COUNT     = 8  //最大支持机柜数量
	MAX_ID_LEN            = 48
	MAX_PARKNO_LEN        = 16
	MAX_ALARMREASON_LEN   = 32
	MAX_UPGRADE_INFO_LEN  = 48  //获取升级文件匹配信息(模糊升级)
	MAX_CUSTOMDIR_LEN     = 32  //自定义目录长度
	MAX_LED_INFO_LEN      = 512 //LED内容长度
	MAX_VOICE_INFO_LEN    = 128 //语音播报内容长度
	MAX_LITLE_INFO_LEN    = 64  //纸票标题内容长度
	MAX_CUSTOM_INFO_LEN   = 64  //纸票自定义信息内容长度
	MAX_PHONE_NUM_LEN     = 16  //联系电话内容长度
	MAX_APP_SERIALNUM_LEN = 32  //应用序列号长度

	AUDIOTALKTYPE_G722    = 0
	AUDIOTALKTYPE_G711_MU = 1
	AUDIOTALKTYPE_G711_A  = 2
	AUDIOTALKTYPE_MP2L2   = 5
	AUDIOTALKTYPE_G726    = 6
	AUDIOTALKTYPE_AAC     = 7
	AUDIOTALKTYPE_PCM     = 8
	AUDIOTALKTYPE_G722C   = 9
	AUDIOTALKTYPE_MP3     = 15 //packet type
	FILE_HEAD             = 0  //file head
	VIDEO_I_FRAME         = 1  //video I frame
	VIDEO_B_FRAME         = 2  //video B frame
	VIDEO_P_FRAME         = 3  //video P frame
	AUDIO_PACKET          = 10 //audio packet
	PRIVT_PACKET          = 11 //private packet
	//E frame
	HIK_H264_E_FRAME           = (1 << 6) // 以前E帧不用了,深P帧也没用到
	MAX_TRANSPARENT_CHAN_NUM   = 4        //每个串口允许建立的最大透明通道数
	MAX_TRANSPARENT_ACCESS_NUM = 4        //每个监听端口允许接入的最大主机数

	//ITS
	MAX_PARKING_STATUS = 8 //车位状态 0代表无车，1代表有车，2代表压线(优先级最高), 3特殊车位
	MAX_PARKING_NUM    = 4 //一个通道最大4个车位 (从左到右车位 数组0～3)

	MAX_ITS_SCENE_NUM     = 16  //最大场景数量
	MAX_SCENE_TIMESEG_NUM = 16  //最大场景时间段数量
	MAX_IVMS_IP_CHANNEL   = 128 //最大IP通道数
	DEVICE_ID_LEN         = 48  //设备编号长度
	MONITORSITE_ID_LEN    = 48  //监测点编号长度
	MAX_AUXAREA_NUM       = 16  //辅助区域最大数目
	MAX_SLAVE_CHANNEL_NUM = 16  //最大从通道数量
	MAX_DEVDESC_LEN       = 64  //设备描述信息最大长度
	ILLEGAL_LEN           = 32  //违法代码长度
	MAX_TRUCK_AXLE_NUM    = 10  //货车轴最大数
	MAX_CATEGORY_LEN      = 8   //车牌附加信息最大字符
	SERIAL_NO_LEN         = 16  //泊车位编号

	MAX_SECRETKEY_LEN          = 512 //最大秘钥长度
	MAX_INDEX_CODE_LEN         = 64  //最大序号长度
	MAX_ILLEGAL_LEN            = 64  //违法代码最大字符长度
	CODE_LEN                   = 64  //授权码
	ALIAS_LEN                  = 32  //别名，只读
	MAX_SCH_TASKS_NUM          = 10
	MAX_SERVERID_LEN           = 64  //最大服务器ID的长度
	MAX_SERVERDOMAIN_LEN       = 128 //服务器域名最大长度
	MAX_AUTHENTICATEID_LEN     = 64  //认证ID最大长度
	MAX_AUTHENTICATEPASSWD_LEN = 32  //认证密码最大长度
	MAX_SERVERNAME_LEN         = 64  //最大服务器用户名
	MAX_COMPRESSIONID_LEN      = 64  //编码ID的最大长度
	MAX_SIPSERVER_ADDRESS_LEN  = 128 //SIP服务器地址支持域名和IP地址
	//压线报警
	MAX_PlATE_NO_LEN = 32 //车牌号码最大长度 2013-09-27
	UPNP_PORT_NUM    = 12 //upnp端口映射端口数目

	MAX_NOTICE_NUMBER_LEN = 32   //公告编号最大长度
	MAX_NOTICE_THEME_LEN  = 64   //公告主题最大长度
	MAX_NOTICE_DETAIL_LEN = 1024 //公告详情最大长度
	MAX_NOTICE_PIC_NUM    = 6    //公告信息最大图片数量
	MAX_DEV_NUMBER_LEN    = 32   //设备编号最大长度
	LOCK_NAME_LEN         = 32   //锁名称
	STREAM_ID_LEN         = 32

	HOLIDAY_GROUP_NAME_LEN          = 32  //假日组名称长度
	MAX_HOLIDAY_PLAN_NUM            = 16  //假日组最大假日计划数
	TEMPLATE_NAME_LEN               = 32  //计划模板名称长度
	MAX_HOLIDAY_GROUP_NUM           = 16  //计划模板最大假日组数
	DOOR_NAME_LEN                   = 32  //门名称
	STRESS_PASSWORD_LEN             = 8   //胁迫密码长度
	SUPER_PASSWORD_LEN              = 8   //胁迫密码长度
	GROUP_NAME_LEN                  = 32  //群组名称长度
	GROUP_COMBINATION_NUM           = 8   //群组组合数
	MULTI_CARD_GROUP_NUM            = 4   //单门最大多重卡组数
	ACS_CARD_NO_LEN                 = 32  //门禁卡号长度
	NET_SDK_EMPLOYEE_NO_LEN         = 32  //工号长度
	NET_SDK_UUID_LEN                = 36  //UUID长度
	NET_SDK_EHOME_KEY_LEN           = 32  //EHome Key长度
	CARD_PASSWORD_LEN               = 8   //卡密码长度
	MAX_DOOR_NUM                    = 32  //最大门数
	MAX_CARD_RIGHT_PLAN_NUM         = 4   //卡权限最大计划个数
	MAX_GROUP_NUM_128               = 128 //最大群组数
	MAX_CARD_READER_NUM             = 64  //最大读卡器数
	MAX_SNEAK_PATH_NODE             = 8   //最大后续读卡器数
	MAX_MULTI_DOOR_INTERLOCK_GROUP  = 8   //最大多门互锁组数
	MAX_INTER_LOCK_DOOR_NUM         = 8   //一个多门互锁组中最大互锁门数
	MAX_CASE_SENSOR_NUM             = 8   //最大case sensor触发器数
	MAX_DOOR_NUM_256                = 256 //最大门数
	MAX_READER_ROUTE_NUM            = 16  //最大刷卡循序路径
	MAX_FINGER_PRINT_NUM            = 10  //最大指纹个数
	MAX_CARD_READER_NUM_512         = 512 //最大读卡器数
	NET_SDK_MULTI_CARD_GROUP_NUM_20 = 20  //单门最大多重卡组数

	ERROR_MSG_LEN       = 32 //下发错误信息
	MAX_DOOR_CODE_LEN   = 8  //房间代码长度
	MAX_LOCK_CODE_LEN   = 8  //锁代码长度
	PER_RING_PORT_NUM   = 2  //每个环的端口数
	SENSORNAME_LEN      = 32 //传感器名称长度
	MAX_SENSORDESCR_LEN = 64 //传感器描述长度
	MAX_DNS_SERVER_NUM  = 2  //最大DNS个数
	SENSORUNIT_LEN      = 32 //最大单位长度

	WEP_KEY_MAX_SIZE = 32 //最大WEP加密密钥长度
	WEP_KEY_MAX_NUM  = 4  //最大WEP加密密钥个数
	WPA_KEY_MAX_SIZE = 64 //最大WPA共享密钥长度

	MAX_SINGLE_FTPPICNAME_LEN = 20 //最大单个FTP通道名称
	MAX_CAMNAME_LEN           = 32 //最大通道名称
	MAX_FTPNAME_NUM           = 12 //TFP名称数

	MAX_IDCODE_LEN  = 128 //  识别码最大长度
	MAX_VERSIIN_LEN = 64  //版本最大长度
	MAX_IDCODE_NUM  = 32  // 识别码个数
	SDK_LEN_2048    = 2048
	SDK_MAX_IP_LEN  = 48
	RECT_POINT_NUM  = 4 //矩形角数

	MAX_PUBLIC_KEY_LEN = 512 // 最大公钥长度
	CHIP_SERIALNO_LEN  = 32  //加密芯片序列号长度
	ENCRYPT_DEV_ID_LEN = 20  //设备ID长度

	//MCU相关的
	MAX_SEARCH_ID_LEN = 36  //搜索标识符最大长度
	TERMINAL_NAME_LEN = 64  //终端名称长度
	MAX_URL_LEN       = 512 //URL长度
	REGISTER_NAME_LEN = 64  //终端注册GK名称最大长度

	//光纤
	MAX_PORT_NUM                 = 64  //最大端口数
	MAX_SINGLE_CARD_PORT_NO      = 4   //光纤收发器单卡最大端口数
	MAX_FUNC_CARD_NUM            = 32  //光纤收发器最大功能卡数
	MAX_FC_CARD_NUM              = 33  //光纤收发器最大卡数
	MAX_REMARKS_LEN              = 128 //注释最大长度
	MAX_OUTPUT_PORT_NUM          = 32  //单路输出包含的最大输出端口数
	MAX_SINGLE_PORT_RECVCARD_NUM = 64  //单个端口连接的最大接收卡数
	MAX_GAMMA_X_VALUE            = 256 //GAMMA表X轴取值个数
	NET_DEV_NAME_LEN             = 64  //设备名称长度
	NET_DEV_TYPE_NAME_LEN        = 64  //设备类型名称长度
	ABNORMAL_INFO_NUM            = 4   //异常时间段个数

	PLAYLIST_NAME_LEN = 64 //播放表名称长度
	PLAYLIST_ITEM_NUM = 64 //播放项数目

	//后端相关
	NET_SDK_MAX_LOGIN_PASSWORD_LEN = 128 //用户登录密码最大长度
	NET_SDK_MAX_ANSWER_LEN         = 256 //安全问题答案最大长度
	NET_SDK_MAX_QUESTION_LIST_LEN  = 32  //安全问题列表最大长度

	MAX_SCREEN_AREA_NUM            = 128 //屏幕区域最大数量
	NET_SDK_MAX_THERMOMETRYALGNAME = 128 //测温算法库版本最大长度
	NET_SDK_MAX_SHIPSALGNAME       = 128 //船只算法库版本最大长度
	NET_SDK_MAX_FIRESALGNAME       = 128 //火点算法库版本最大长度

	MAX_PASSPORT_NUM_LEN     = 16   //最大护照证件号长度
	MAX_PASSPORT_INFO_LEN    = 128  //最大护照通用信息长度
	MAX_PASSPORT_NAME_LEN    = 64   //最大护照姓名长度
	MAX_PASSPORT_MONITOR_LEN = 1024 //最大护照监护信息长度
	MAX_NATIONALITY_LEN      = 16   //最大护照国籍长度
	MAX_PASSPORT_TYPE_LEN    = 4    //最大护照证件类型长度

	MAX_LICENSE_LEN = 16
)

type CommAlarm int

const (

	//报警回调命令
	CaCommAlarm CommAlarm = 0x1100 //8000报警信息主动上传
	//对应NET_VCA_RULE_ALARM
	CommAlarmRule          CommAlarm = 0x1102 //行为分析报警信息
	CommAlarmPdc           CommAlarm = 0x1103 //人数统计报警信息
	CommAlarmVideoplatform CommAlarm = 0x1104 //视频综合平台报警
	CommAlarmAlarmhost     CommAlarm = 0x1105 //网络报警主机报警
	CommAlarmFace          CommAlarm = 0x1106 //人脸检测识别报警信息
	CommRuleInfoUpload     CommAlarm = 0x1107 // 事件数据信息上传
	CommAlarmAid           CommAlarm = 0x1110 //交通事件报警信息
	CommAlarmTps           CommAlarm = 0x1111 //交通参数统计报警信息
	//智能人脸抓拍结果上传
	CommUploadFacesnapResult            CommAlarm = 0x1112 //人脸识别结果上传
	CommAlarmTfs                        CommAlarm = 0x1113 //交通取证报警信息
	CommAlarmTpsV41                     CommAlarm = 0x1114 //交通参数统计报警信息扩展
	CommAlarmAidV41                     CommAlarm = 0x1115 //交通事件报警信息扩展
	CommAlarmVqdEx                      CommAlarm = 0x1116 //视频质量诊断报警
	CommAlarmNotificationReport         CommAlarm = 0x1117 //通知事件上报
	CommSensorValueUpload               CommAlarm = 0x1120 //模拟量数据实时上传
	CommSensorAlarm                     CommAlarm = 0x1121 //模拟量报警上传
	CommSwitchAlarm                     CommAlarm = 0x1122 //开关量报警
	CommAlarmhostException              CommAlarm = 0x1123 //报警主机故障报警
	CommAlarmhostOperateeventAlarm      CommAlarm = 0x1124 //操作事件报警上传
	CommAlarmhostSafetycabinstate       CommAlarm = 0x1125 //防护舱状态
	CommAlarmhostAlarmoutstatus         CommAlarm = 0x1126 //报警输出口/警号状态
	CommAlarmhostCidAlarm               CommAlarm = 0x1127 //报告报警上传
	CommAlarmhostExternalDeviceAlarm    CommAlarm = 0x1128 //报警主机外接设备报警上传
	CommAlarmhostDataUpload             CommAlarm = 0x1129 //报警数据上传
	CommFacecaptureStatisticsResult     CommAlarm = 0x112a //人脸抓拍统计上传
	CommAlarmWirelessInfo               CommAlarm = 0x122b // 无线网络信息上传
	CommScenechangeDetectionUpload      CommAlarm = 0x1130 //场景变更报警上传(布防)2013-7-16
	CommCrosslineAlarm                  CommAlarm = 0x1131 //压线报警(监听) 2013-09-27
	CommUploadVideoIntercomEvent        CommAlarm = 0x1132 //可视对讲事件记录上传
	CommAlarmVideoIntercom              CommAlarm = 0x1133 //可视对讲报警上传
	CommUploadNoticeData                CommAlarm = 0x1134 //可视对讲公告信息上传
	CommAlarmAudioexception             CommAlarm = 0x1150 //声音报警信息
	CommAlarmDefocus                    CommAlarm = 0x1151 //虚焦报警信息
	CommAlarmButtonDownException        CommAlarm = 0x1152 //按钮按下报警信息
	CommAlarmAlarmgps                   CommAlarm = 0x1202 //GPS报警信息上传
	CommTradeinfo                       CommAlarm = 0x1500 //ATMDVR主动上传交易信息
	CommUploadPlateResult               CommAlarm = 0x2800 //上传车牌信息
	CommItcStatusDetectResult           CommAlarm = 0x2810 //实时状态检测结果上传(智能高清IPC)
	CommIpcAuxalarmResult               CommAlarm = 0x2820 //PIR报警、无线报警、呼救报警上传
	CommUploadPictureinfo               CommAlarm = 0x2900 //上传图片信息
	CommSnapMatchAlarm                  CommAlarm = 0x2902 //黑名单比对结果上传
	CommItsPlateResult                  CommAlarm = 0x3050 //终端图片上传
	CommItsTrafficCollect               CommAlarm = 0x3051 //终端统计数据上传
	CommItsGateVehicle                  CommAlarm = 0x3052 //出入口车辆抓拍数据上传
	CommItsGateFace                     CommAlarm = 0x3053 //出入口人脸抓拍数据上传
	CommItsGateCostitem                 CommAlarm = 0x3054 //出入口过车收费明细 2013-11-19
	CommItsGateHandover                 CommAlarm = 0x3055 //出入口交接班数据 2013-11-19
	CommItsParkVehicle                  CommAlarm = 0x3056 //停车场数据上传
	CommItsBlacklistAlarm               CommAlarm = 0x3057 //黑名单报警上传
	CommVehicleControlListDsalarm       CommAlarm = 0x3058 //黑白名单数据需要同步报警2013-11-04
	CommVehicleControlAlarm             CommAlarm = 0x3059 //车辆报警2013-11-04
	CommFireAlarm                       CommAlarm = 0x3060 //消防报警2013-11-04
	CommItsGateAlarminfo                CommAlarm = 0x3061 //出入口控制机数据上传
	CommVehicleRecogResult              CommAlarm = 0x3062 //车辆二次识别结果上传 2014-11-12
	CommPlateResultV50                  CommAlarm = 0x3063 //车牌上传 V50
	CommGateChargeinfoUpload            CommAlarm = 0x3064 //出入口付费信息上传
	CommTmeVehicleIndentification       CommAlarm = 0x3065 //TME车辆抓图上传
	CommGateCardinfoUpload              CommAlarm = 0x3066 //出入口卡片信息上传
	CommLoadingDockOperateinfo          CommAlarm = 0x3067 //月台作业上传
	CommAlarmSensorinfoUpload           CommAlarm = 0x3077 //传感器上传信息
	CommAlarmCaptureUpload              CommAlarm = 0x3078 //抓拍图片上传
	CommItsRadarinfo                    CommAlarm = 0x3079 //雷达报警上传
	CommSignalLampAbnormal              CommAlarm = 0x3080 //信号灯异常检测上传
	CommAlarmTpsRealTime                CommAlarm = 0x3081 //TPS实时过车数据上传
	CommAlarmTpsStatistics              CommAlarm = 0x3082 //TPS统计过车数据上传
	CommAlarmV30                        CommAlarm = 0x4000 //9000报警信息主动上传
	CommIpccfg                          CommAlarm = 0x4001 //9000设备IPC接入配置改变报警信息主动上传
	CommIpccfgV31                       CommAlarm = 0x4002 //9000设备IPC接入配置改变报警信息主动上传扩展 9000_1.1
	CommIpccfgV40                       CommAlarm = 0x4003 // IVMS 2000 编码服务器 NVR IPC接入配置改变时报警信息上传
	CommAlarmDevice                     CommAlarm = 0x4004 //设备报警内容，由于通道值大于256而扩展
	CommAlarmCvr                        CommAlarm = 0x4005 //CVR 2.0.X外部报警类型
	CommAlarmHotSpare                   CommAlarm = 0x4006 //热备异常报警（N+1模式异常报警）
	CommAlarmV40                        CommAlarm = 0x4007 //移动侦测，视频丢失，遮挡，IO信号量等报警信息主动上传，报警数据为可变长
	CommUploadHeatmapResult             CommAlarm = 0x4008 //热度图报警上传 2014-03-21
	CommAlarmDeviceV40                  CommAlarm = 0x4009 //设备报警内容扩展
	CommAlarmFaceDetection              CommAlarm = 0x4010 //人脸侦测报警
	CommAlarmTargetLeftRegion           CommAlarm = 0x4011 //检测目标离开检测区域报警(教师走向学生报警(用于联动切换录播主机控制检测学生的球机))
	CommGisinfoUpload                   CommAlarm = 0x4012 //GIS信息上传
	CommVandalproofAlarm                CommAlarm = 0x4013 //上传防破坏报警信息
	CommPeopleDetectionUpload           CommAlarm = 0x4014 //人员侦测信息上传
	CommAlarmStorageDetection           CommAlarm = 0x4015 //存储智能检测报警上传
	CommMvmRegister                     CommAlarm = 0x4016 //地磁管理器（Magnetic Vehicle Manager）注册
	CommMvmStatusInfo                   CommAlarm = 0x4017 //地磁管理器（Magnetic Vehicle Manager）状态上报
	CommUploadHeatmapResultPdc          CommAlarm = 0x4018 //热度图按人数统计数据上传事件
	CommUploadHeatmapResultDuration     CommAlarm = 0x4019 //热度图按人员停留时间统计数据上传事件
	CommUploadHeatmapResultIntersection CommAlarm = 0x4020 //路口分析热度值结果上传
	CommUploadAiopVideo                 CommAlarm = 0x4021 //设备支持AI开放平台接入，上传视频检测数据
	CommUploadAiopPicture               CommAlarm = 0x4022 //设备支持AI开放平台接入，上传图片检测数据
	CommUploadAiopPollingSnap           CommAlarm = 0x4023 //设备支持AI开放平台接入，上传轮巡抓图图片检测数据 对应的结构体(NET_AIOP_POLLING_PICTURE_HEAD)
	CommUploadAiopPollingVideo          CommAlarm = 0x4024 //设备支持AI开放平台接入，上传轮巡视频检测数据 对应的结构体(NET_AIOP_POLLING_VIDEO_HEAD)
	CommItsRoadException                CommAlarm = 0x4500 //路口设备异常报警
	CommItsExternalControlAlarm         CommAlarm = 0x4520 //外控报警
	CommAlarmShipsdetection             CommAlarm = 0x4521 // 船只检测报警信息
	CommVcaDbdAlarm                     CommAlarm = 0x4550 //驾驶行为报警信息
	CommVcaAdasAlarm                    CommAlarm = 0x4551 //高级辅助驾驶报警信息
	CommVehRealtimeInfo                 CommAlarm = 0x4552 //行车实时数据信息
	CommVcaAttendAlarm                  CommAlarm = 0x4553 //考勤事件报警信息
	CommFiredetectionAlarm              CommAlarm = 0x4991 //火点检测报警
	CommAlarmDensefogdetection          CommAlarm = 0x4992 //大雾检测报警信息
	CommVcaAlarm                        CommAlarm = 0x4993 //智能检测报警
	CommFaceThermometryAlarm            CommAlarm = 0x4994 //人脸测温报警
	CommTapeArchiveAlarm                CommAlarm = 0x4996 //磁带库归档报警
	CommScreenAlarm                     CommAlarm = 0x5000 //多屏控制器报警类型
	CommDvcsStateAlarm                  CommAlarm = 0x5001 //分布式大屏控制器报警上传
	CommAlarmAcs                        CommAlarm = 0x5002 //门禁主机报警
	CommAlarmFiberConvert               CommAlarm = 0x5003 //光纤收发器报警
	CommAlarmSwitchConvert              CommAlarm = 0x5004 //交换机报警
	CommAlarmDecVca                     CommAlarm = 0x5010 //智能解码报警
	CommAlarmLcd                        CommAlarm = 0x5011 //屏幕报警
	CommConferenceCallAlarm             CommAlarm = 0x5012 //会议呼叫告警
	CommAlarmWallConfernece             CommAlarm = 0x5015 //MCU单个已开会的会议信息报警
	CommDiagnosisUpload                 CommAlarm = 0x5100 //诊断服务器VQD报警上传
	CommHighDensityUpload               CommAlarm = 0x5101 //人员聚集密度输出报警上传
	CommIdInfoAlarm                     CommAlarm = 0x5200 //身份证信息上传
	CommPassnumInfoAlarm                CommAlarm = 0x5201 //通行人数上报
	CommPassportAlarm                   CommAlarm = 0x5202 //护照信息上传
	CommThermometryDiffAlarm            CommAlarm = 0x5211 //温差报警上传
	CommThermometryAlarm                CommAlarm = 0x5212 //温度报警上传
	CommPanoramicLinkageAlarm           CommAlarm = 0x5213 //全景联动到位上传
	CommTagInfoAlarm                    CommAlarm = 0x5215 // 标签信息上传
	CommAlarmVqd                        CommAlarm = 0x6000 //VQD主动报警上传
	CommPushUpdateRecordInfo            CommAlarm = 0x6001 //推模式录像信息上传
	CommSwitchLampAlarm                 CommAlarm = 0x6002 //开关灯检测
	CommInquestAlarm                    CommAlarm = 0x6005 // 审讯主机报警上传
	CommVideoParkingPoleAlarm           CommAlarm = 0x6006 //视频桩报警
	CommGpsStatusAlarm                  CommAlarm = 0x6010 // GPS状态上传
	CommBaseStationInfoAlarm            CommAlarm = 0x6011 //基站信息上传
	CommAlarmSubscribeEvent             CommAlarm = 0x6012 //订阅结果上报
	CommFacesnapRawdataAlarm            CommAlarm = 0x6015 //人脸比对报警（数据透传方式）
	CommClusterAlarm                    CommAlarm = 0x6020 //集群报警上传
	CommIsapiAlarm                      CommAlarm = 0x6009
)

const (
	ITC_MAX_POLYGON_POINT_NUM = 20 //检测区域最多支持20个点的多边形
	MAX_MOBILE_POLYGON_NUM    = 3  //移动布控支持最大牌识区域个数
	MAX_MOBILE_DETECTLINE_NUM = 3  //移动布控支持最大违规检测线个数

	MAX_ITC_LANE_NUM = 6
)

type DVRGetSet int

const (
	NetDvrGetDevicecfg    DVRGetSet = 100 //获取设备参数
	NetDvrSetDevicecfg    DVRGetSet = 101 //设置设备参数
	NetDvrGetNetcfg       DVRGetSet = 102 //获取网络参数
	NetDvrSetNetcfg       DVRGetSet = 103 //设置网络参数
	NetDvrGetPiccfg       DVRGetSet = 104 //获取图象参数
	NetDvrSetPiccfg       DVRGetSet = 105 //设置图象参数
	NetDvrGetCompresscfg  DVRGetSet = 106 //获取压缩参数
	NetDvrSetCompresscfg  DVRGetSet = 107 //设置压缩参数
	NetDvrGetRecordcfg    DVRGetSet = 108 //获取录像时间参数
	NetDvrSetRecordcfg    DVRGetSet = 109 //设置录像时间参数
	NetDvrGetDecodercfg   DVRGetSet = 110 //获取解码器参数
	NetDvrSetDecodercfg   DVRGetSet = 111 //设置解码器参数
	NetDvrGetRs232Cfg     DVRGetSet = 112 //获取232串口参数
	NetDvrSetRs232Cfg     DVRGetSet = 113 //设置232串口参数
	NetDvrGetAlarmincfg   DVRGetSet = 114 //获取报警输入参数
	NetDvrSetAlarmincfg   DVRGetSet = 115 //设置报警输入参数
	NetDvrGetAlarmoutcfg  DVRGetSet = 116 //获取报警输出参数
	NetDvrSetAlarmoutcfg  DVRGetSet = 117 //设置报警输出参数
	NetDvrGetTimecfg      DVRGetSet = 118 //获取DVR时间
	NetDvrSetTimecfg      DVRGetSet = 119 //设置DVR时间
	NetDvrGetPreviewcfg   DVRGetSet = 120 //获取预览参数
	NetDvrSetPreviewcfg   DVRGetSet = 121 //设置预览参数
	NetDvrGetVideooutcfg  DVRGetSet = 122 //获取视频输出参数
	NetDvrSetVideooutcfg  DVRGetSet = 123 //设置视频输出参数
	NetDvrGetUsercfg      DVRGetSet = 124 //获取用户参数
	NetDvrSetUsercfg      DVRGetSet = 125 //设置用户参数
	NetDvrGetExceptioncfg DVRGetSet = 126 //获取异常参数
	NetDvrSetExceptioncfg DVRGetSet = 127 //设置异常参数
	NetDvrGetZoneanddst   DVRGetSet = 128 //获取时区和夏时制参数
	NetDvrSetZoneanddst   DVRGetSet = 129 //设置时区和夏时制参数

	//注：该命令只支持4条OSD的类型，通常用于V30以下的设备版本。
	NetDvrGetShowstring DVRGetSet = 130 //获取叠加字符参数
	NetDvrSetShowstring DVRGetSet = 131 //设置叠加字符参数

	NetDvrGetEventcompcfg  DVRGetSet = 132 //获取事件触发录像参数
	NetDvrSetEventcompcfg  DVRGetSet = 133 //设置事件触发录像参数
	NetDvrGetFtpcfg        DVRGetSet = 134 //获取抓图的FTP参数(基线)
	NetDvrSetFtpcfg        DVRGetSet = 135 //设置抓图的FTP参数(基线)
	NetDvrGetAuxoutcfg     DVRGetSet = 140 //获取报警触发辅助输出设置(HS设备辅助输出2006-02-28)
	NetDvrSetAuxoutcfg     DVRGetSet = 141 //设置报警触发辅助输出设置(HS设备辅助输出2006-02-28)
	NetDvrGetPreviewcfgAux DVRGetSet = 142 //获取-s系列双输出预览参数(-s系列双输出2006-04-13)
	NetDvrSetPreviewcfgAux DVRGetSet = 143 //设置-s系列双输出预览参数(-s系列双输出2006-04-13)

	NetDvrGetPasswordManageCfg DVRGetSet = 144 //获取密码管理配置
	NetDvrSetPasswordManageCfg DVRGetSet = 145 //设置密码管理配置
	NetDvrUnlockUser           DVRGetSet = 146 //用户解锁
	NetDvrGetSecurityCfg       DVRGetSet = 147 //获取安全认证配置
	NetDvrSetSecurityCfg       DVRGetSet = 148 //设置安全认证配置
	NetDvrGetLockedInfoList    DVRGetSet = 149 //获取所有被锁定信息

	/*********************************智能部分接口 begin***************************************/
	//行为对应（NET_VCA_RULECFG）
	NetDvrSetRulecfg DVRGetSet = 152 //设置行为分析规则
	NetDvrGetRulecfg DVRGetSet = 153 //获取行为分析规则
	//球机标定参数（NET_DVR_TRACK_CFG ）
	NetDvrSetTrackCfg DVRGetSet = 160 //设置球机的配置参数
	NetDvrGetTrackCfg DVRGetSet = 161 //获取球机的配置参数

	//智能分析仪取流配置结构
	NetDvrSetIvmsStreamcfg DVRGetSet = 162 //设置智能分析仪取流参数
	NetDvrGetIvmsStreamcfg DVRGetSet = 163 //获取智能分析仪取流参数
	//智能控制参数结构
	NetDvrSetVcaCtrlcfg DVRGetSet = 164 //设置智能控制参数
	NetDvrGetVcaCtrlcfg DVRGetSet = 165 //获取智能控制参数
	//屏蔽区域NET_VCA_MASK_REGION_LIST
	NetDvrSetVcaMaskRegion DVRGetSet = 166 //设置屏蔽区域参数
	NetDvrGetVcaMaskRegion DVRGetSet = 167 //获取屏蔽区域参数

	//ATM进入区域 NET_VCA_ENTER_REGION
	NetDvrSetVcaEnterRegion DVRGetSet = 168 //设置进入区域参数
	NetDvrGetVcaEnterRegion DVRGetSet = 169 //获取进入区域参数

	//标定线配置NET_VCA_LINE_SEGMENT_LIST
	NetDvrSetVcaLineSegment DVRGetSet = 170 //设置标定线
	NetDvrGetVcaLineSegment DVRGetSet = 171 //获取标定线

	// ivms屏蔽区域NET_IVMS_MASK_REGION_LIST
	NetDvrSetIvmsMaskRegion DVRGetSet = 172 //设置IVMS屏蔽区域参数
	NetDvrGetIvmsMaskRegion DVRGetSet = 173 //获取IVMS屏蔽区域参数
	// ivms进入检测区域NET_IVMS_ENTER_REGION
	NetDvrSetIvmsEnterRegion DVRGetSet = 174 //设置IVMS进入区域参数
	NetDvrGetIvmsEnterRegion DVRGetSet = 175 //获取IVMS进入区域参数
	NetDvrSetIvmsBehaviorcfg DVRGetSet = 176 //设置智能分析仪行为规则参数
	NetDvrGetIvmsBehaviorcfg DVRGetSet = 177 //获取智能分析仪行为规则参数

	// IVMS 回放检索
	NetDvrIvmsSetSearchcfg   DVRGetSet = 178 //设置IVMS回放检索参数
	NetDvrIvmsGetSearchcfg   DVRGetSet = 179 //获取IVMS回放检索参数
	NetDvrSetPositionTrack   DVRGetSet = 180 // 设置场景跟踪配置信息
	NetDvrGetPositionTrack   DVRGetSet = 181 // 获取场景跟踪配置信息
	NetDvrSetCalibration     DVRGetSet = 182 // 设置标定信息
	NetDvrGetCalibration     DVRGetSet = 183 // 获取标定信息
	NetDvrSetPdcRulecfg      DVRGetSet = 184 // 设置人流量统计规则
	NetDvrGetPdcRulecfg      DVRGetSet = 185 // 获取人流量统计规则
	NetDvrSetPuStreamcfg     DVRGetSet = 186 // 设置前段取流设备信息
	NetDvrGetPuStreamcfg     DVRGetSet = 187 // 获取前段取流设备信息
	NetVcaSetIvmsBehaviorCfg DVRGetSet = 192 // 设置IVMS行为规则配置 不带时间段
	NetVcaGetIvmsBehaviorCfg DVRGetSet = 193 // 获取IVMS行为规则配置 不带时间段
	NetVcaSetSizeFilter      DVRGetSet = 194 // 设置全局尺寸过滤器
	NetVcaGetSizeFilter      DVRGetSet = 195 // 获取全局尺寸过滤器

	NetDvrSetTrackParamcfg     DVRGetSet = 196 // 设置球机本地菜单规则
	NetDvrGetTrackParamcfg     DVRGetSet = 197 // 获取球机本地菜单规则
	NetDvrSetDomeMovementParam DVRGetSet = 198 // 设置球机机芯参数
	NetDvrGetDomeMovementParam DVRGetSet = 199 // 获取球机机芯参数
	NetDvrGetPiccfgEx          DVRGetSet = 200 //获取图象参数(SDK_V14扩展命令)
	NetDvrSetPiccfgEx          DVRGetSet = 201 //设置图象参数(SDK_V14扩展命令)
	NetDvrGetUsercfgEx         DVRGetSet = 202 //获取用户参数(SDK_V15扩展命令)
	NetDvrSetUsercfgEx         DVRGetSet = 203 //设置用户参数(SDK_V15扩展命令)
	NetDvrGetCompresscfgEx     DVRGetSet = 204 //获取压缩参数(SDK_V15扩展命令2006-05-15)
	NetDvrSetCompresscfgEx     DVRGetSet = 205 //设置压缩参数(SDK_V15扩展命令2006-05-15)
	NetDvrGetNetappcfg         DVRGetSet = 222 //获取网络应用参数 NTP/DDNS/EMAIL
	NetDvrSetNetappcfg         DVRGetSet = 223 //设置网络应用参数 NTP/DDNS/EMAIL
	NetDvrGetNtpcfg            DVRGetSet = 224 //获取网络应用参数 NTP
	NetDvrSetNtpcfg            DVRGetSet = 225 //设置网络应用参数 NTP
	NetDvrGetDdnscfg           DVRGetSet = 226 //获取网络应用参数 DDNS
	NetDvrSetDdnscfg           DVRGetSet = 227 //设置网络应用参数 DDNS
	//对应NET_DVR_EMAILPARA
	NetDvrGetEmailcfg DVRGetSet = 228 //获取网络应用参数 EMAIL
	NetDvrSetEmailcfg DVRGetSet = 229 //设置网络应用参数 EMAIL
	NetDvrGetNfscfg   DVRGetSet = 230 /* NFS disk config */
	NetDvrSetNfscfg   DVRGetSet = 231 /* NFS disk config */
	/*注：该命令为定制，只支持8条OSD的类型，不会兼容V30设备版本之前的
	  NET_DVR_GET_SHOWSTRING 、NET_DVR_SET_SHOWSTRING 命令。（不建议使用）*/
	NetDvrGetShowstringEx DVRGetSet = 238 //获取叠加字符参数扩展(支持8条字符)
	NetDvrSetShowstringEx DVRGetSet = 239 //设置叠加字符参数扩展(支持8条字符)
	NetDvrGetNetcfgOther  DVRGetSet = 244 //获取网络参数
	NetDvrSetNetcfgOther  DVRGetSet = 245 //设置网络参数

	//对应NET_DVR_EMAILCFG结构
	NetDvrGetEmailparacfg DVRGetSet = 250 //Get EMAIL parameters
	NetDvrSetEmailparacfg DVRGetSet = 251 //Setup EMAIL parameters
	etDvrGetDdnscfgEx     DVRGetSet = 274 //获取扩展DDNS参数
	NetDvrSetDdnscfgEx    DVRGetSet = 275 //设置扩展DDNS参数
	NetDvrSetPtzpos       DVRGetSet = 292 //云台设置PTZ位置
	NetDvrGetPtzpos       DVRGetSet = 293 //云台获取PTZ位置
	NetDvrGetPtzscope     DVRGetSet = 294 //云台获取PTZ范围
	NetDvrGetApInfoList   DVRGetSet = 305 //获取无线网络资源参数
	NetDvrSetWifiCfg      DVRGetSet = 306 //设置IP监控设备无线参数
	NetDvrGetWifiCfg      DVRGetSet = 307 //获取IP监控设备无线参数
	NetDvrSetWifiWorkmode DVRGetSet = 308 //设置IP监控设备网口工作模式参数
	NetDvrGetWifiWorkmode DVRGetSet = 309 //获取IP监控设备网口工作模式参数
	NetDvrGetWifiStatus   DVRGetSet = 310 //获取设备当前wifi连接状态
	/*********************************智能交通事件begin***************************************/
	NetDvrGetReferenceRegion   DVRGetSet = 400 //获取参考区域
	NetDvrSetReferenceRegion   DVRGetSet = 401 //设置参考区域
	NetDvrGetTrafficMaskRegion DVRGetSet = 402 //获取交通事件屏蔽区域
	NetDvrSetTrafficMaskRegion DVRGetSet = 403 //设置交通事件屏蔽区域
	NetDvrSetAidRulecfg        DVRGetSet = 404 //设置交通事件规则参数
	NetDvrGetAidRulecfg        DVRGetSet = 405 //获取交通事件规则参数
	NetDvrSetTpsRulecfg        DVRGetSet = 406 //设置交通统计规则参数
	NetDvrGetTpsRulecfg        DVRGetSet = 407 //获取交通统计规则参数
	NetDvrSetLanecfg           DVRGetSet = 408 //设置车道规则
	NetDvrGetLanecfg           DVRGetSet = 409 //获取车道规则
	NetDvrGetVcaRuleColorCfg   DVRGetSet = 410 //获取智能规则关联的颜色参数
	NetDvrSetVcaRuleColorCfg   DVRGetSet = 411 //设置智能规则关联的颜色参数
	NetDvrGetSwitchLampCfg     DVRGetSet = 412 //获取开关灯检测规则配置参数
	NetDvrSetSwitchLampCfg     DVRGetSet = 413 //设置开关灯检测规则配置参数

	/*********************************智能交通事件end***************************************/
	NetDvrSetFacedetectRulecfg DVRGetSet = 420 // 设置人脸检测规则
	NetDvrGetFacedetectRulecfg DVRGetSet = 421 // 获取人脸检测规则
	NetDvrSetVehicleRecogTask  DVRGetSet = 422 //车辆二次识别任务提交
	NetDvrGetVehicleRecogTask  DVRGetSet = 423 //车辆二次识别任务获取
	NetDvrSetTimecorrect       DVRGetSet = 432 //校时配置（只做校时操作，不记录校时配置 eg.NET_DVR_SET_TIMECFG 会修改设备的校时配置（NTP校时，会被修改为手动校时））
	NetDvrGetConnectList       DVRGetSet = 433 //获取连接设备列表信息

	/***************************DS9000新增命令(_V30) begin *****************************/
	//网络(NET_DVR_NETCFG_V30结构)
	NetDvrGetNetcfgV30 DVRGetSet = 1000 //获取网络参数
	NetDvrSetNetcfgV30 DVRGetSet = 1001 //设置网络参数
	//图象(NET_DVR_PICCFG_V30结构)
	NetDvrGetPiccfgV30 DVRGetSet = 1002 //获取图象参数
	NetDvrSetPiccfgV30 DVRGetSet = 1003 //设置图象参数
	//录像时间(NET_DVR_RECORD_V30结构)
	NetDvrGetRecordcfgV30 DVRGetSet = 1004 //获取录像参数
	NetDvrSetRecordcfgV30 DVRGetSet = 1005 //设置录像参数

	//用户(NET_DVR_USER_V30结构)
	NetDvrGetUsercfgV30 DVRGetSet = 1006 //获取用户参数
	NetDvrSetUsercfgV30 DVRGetSet = 1007 //设置用户参数

	//录像时间(NET_DVR_RECORD_V40结构)
	NetDvrGetRecordcfgV40 DVRGetSet = 1008 //获取录像参数(扩展)
	NetDvrSetRecordcfgV40 DVRGetSet = 1009 //设置录像参数(扩展)

	//9000DDNS参数配置(NET_DVR_DDNSPARA_V30结构)
	NetDvrGetDdnscfgV30 DVRGetSet = 1010 //获取DDNS(9000扩展)
	NetDvrSetDdnscfgV30 DVRGetSet = 1011 //设置DDNS(9000扩展)

	//EMAIL功能(NET_DVR_EMAILCFG_V30结构)
	NetDvrGetEmailcfgV30 DVRGetSet = 1012 //获取EMAIL参数
	NetDvrSetEmailcfgV30 DVRGetSet = 1013 //设置EMAIL参数

	NetDvrGetNetcfgV50 DVRGetSet = 1015 //获取网络参数配置(V50)
	NetDvrSetNetcfgV50 DVRGetSet = 1016 //设置网络参数配置(V50)

	NetGetCruisepointV40 DVRGetSet = 1018 //获取巡航路径配置

	//巡航参数 (NET_DVR_CRUISE_PARA结构)
	NetDvrGetCruise         DVRGetSet = 1020
	NetDvrSetCruise         DVRGetSet = 1021 //报警输入结构参数 (NET_DVR_ALARMINCFG_V30结构)
	NetDvrGetAlarmincfgV30  DVRGetSet = 1024
	NetDvrSetAlarmincfgV30  DVRGetSet = 1025 //报警输出结构参数 (NET_DVR_ALARMOUTCFG_V30结构)
	NetDvrGetAlarmoutcfgV30 DVRGetSet = 1026
	NetDvrSetAlarmoutcfgV30 DVRGetSet = 1027 //视频输出结构参数 (NET_DVR_VIDEOOUT_V30结构)
	NetDvrGetVideooutcfgV30 DVRGetSet = 1028
	NetDvrSetVideooutcfgV30 DVRGetSet = 1029 /*该命令支持8条OSD的类型（即设备版本为V30以上时），并会通过设备版本的匹配，
	同时兼容之前的NET_DVR_GET_SHOWSTRING 、NET_DVR_SET_SHOWSTRING 命令。（建议使用）*/
	//叠加字符结构参数 (NET_DVR_SHOWSTRING_V30结构)
	NetDvrGetShowstringV30   DVRGetSet = 1030
	NetDvrSetShowstringV30   DVRGetSet = 1031 //异常结构参数 (NET_DVR_EXCEPTION_V30结构)
	NetDvrGetExceptioncfgV30 DVRGetSet = 1034
	NetDvrSetExceptioncfgV30 DVRGetSet = 1035 //串口232结构参数 (NET_DVR_RS232CFG_V30结构)
	NetDvrGetRs232CfgV30     DVRGetSet = 1036
	NetDvrSetRs232CfgV30     DVRGetSet = 1037 //网络硬盘接入结构参数 (NET_DVR_NET_DISKCFG结构)
	NetDvrGetNetDiskcfg      DVRGetSet = 1038 //网络硬盘接入获取
	NetDvrSetNetDiskcfg      DVRGetSet = 1039 //网络硬盘接入设置
	//压缩参数 (NET_DVR_COMPRESSIONCFG_V30结构)
	NetDvrGetCompresscfgV30 DVRGetSet = 1040
	NetDvrSetCompresscfgV30 DVRGetSet = 1041 //获取485解码器参数 (NET_DVR_DECODERCFG_V30结构)
	NetDvrGetDecodercfgV30  DVRGetSet = 1042 //获取解码器参数
	NetDvrSetDecodercfgV30  DVRGetSet = 1043 //设置解码器参数

	//获取预览参数 (NET_DVR_PREVIEWCFG_V30结构)
	NetDvrGetPreviewcfgV30 DVRGetSet = 1044 //获取预览参数
	NetDvrSetPreviewcfgV30 DVRGetSet = 1045 //设置预览参数

	//辅助预览参数 (NET_DVR_PREVIEWCFG_AUX_V30结构)
	NetDvrGetPreviewcfgAuxV30 DVRGetSet = 1046 //获取辅助预览参数
	NetDvrSetPreviewcfgAuxV30 DVRGetSet = 1047 //设置辅助预览参数

	//IP接入配置参数 （NET_DVR_IPPARACFG结构）
	NetDvrGetIpparacfg DVRGetSet = 1048 //获取IP接入配置信息
	NetDvrSetIpparacfg DVRGetSet = 1049 //设置IP接入配置信息

	//IP报警输入接入配置参数 （NET_DVR_IPALARMINCFG结构）
	NetDvrGetIpalarmincfg DVRGetSet = 1050 //获取IP报警输入接入配置信息
	NetDvrSetIpalarmincfg DVRGetSet = 1051 //设置IP报警输入接入配置信息
	//IP报警输出接入配置参数 （NET_DVR_IPALARMOUTCFG结构）
	NetDvrGetIpalarmoutcfg DVRGetSet = 1052 //获取IP报警输出接入配置信息
	NetDvrSetIpalarmoutcfg DVRGetSet = 1053 //设置IP报警输出接入配置信息
	//硬盘管理的参数获取 (NET_DVR_HDCFG结构)
	NetDvrGetHdcfg DVRGetSet = 1054 //获取硬盘管理配置参数
	NetDvrSetHdcfg DVRGetSet = 1055 //设置硬盘管理配置参数
	//盘组管理的参数获取 (NET_DVR_HDGROUP_CFG结构)
	NetDvrGetHdgroupCfg DVRGetSet = 1056 //获取盘组管理配置参数
	NetDvrSetHdgroupCfg DVRGetSet = 1057 //设置盘组管理配置参数

	//设备编码类型配置(NET_DVR_COMPRESSION_AUDIO结构)
	NetDvrGetCompresscfgAud DVRGetSet = 1058 //获取设备语音对讲编码参数
	NetDvrSetCompresscfgAud DVRGetSet = 1059 //设置设备语音对讲编码参数

	//IP接入配置参数 （NET_DVR_IPPARACFG_V31结构）
	NetDvrGetIpparacfgV31 DVRGetSet = 1060 //获取IP接入配置信息
	NetDvrSetIpparacfgV31 DVRGetSet = 1061 //设置IP接入配置信息

	// 通道资源配置 (NET_DVR_IPPARACFG_V40结构)
	NetDvrGetIpparacfgV40          DVRGetSet = 1062 // 获取IP接入配置
	NetDvrSetIpparacfgV40          DVRGetSet = 1063 // 设置IP接入配置
	NetDvrGetCcdparamcfg           DVRGetSet = 1067 //IPC获取CCD参数配置
	NetDvrSetCcdparamcfg           DVRGetSet = 1068 //IPC设置CCD参数配置
	NetDvrGetIoincfg               DVRGetSet = 1070 //获取抓拍机IO输入参数
	NetDvrSetIoincfg               DVRGetSet = 1071 //设置抓拍机IO输入参数
	NetDvrGetIooutcfg              DVRGetSet = 1072 //获取抓拍机IO输出参数
	NetDvrSetIooutcfg              DVRGetSet = 1073 //设置抓拍机IO输出参数
	NetDvrGetFlashcfg              DVRGetSet = 1074 //获取IO闪光灯输出参数
	NetDvrSetFlashcfg              DVRGetSet = 1075 //设置IO闪光灯输出参数
	NetDvrGetLightsnapcfg          DVRGetSet = 1076 //获取抓拍机红绿灯参数
	NetDvrSetLightsnapcfg          DVRGetSet = 1077 //设置抓拍机红绿灯参数
	NetDvrGetMeasurespeedcfg       DVRGetSet = 1078 //获取抓拍机测速参数
	NetDvrSetMeasurespeedcfg       DVRGetSet = 1079 //设置抓拍机测速参数
	NetDvrGetImageoverlaycfg       DVRGetSet = 1080 //获取抓拍机图像叠加信息参数
	NetDvrSetImageoverlaycfg       DVRGetSet = 1081 //设置抓拍机图像叠加信息参数
	NetDvrGetSnapcfg               DVRGetSet = 1082 //获取单IO触发抓拍功能配置
	NetDvrSetSnapcfg               DVRGetSet = 1083 //设置单IO触发抓拍功能配置
	NetDvrGetVtpparam              DVRGetSet = 1084 //获取虚拟线圈参数
	NetDvrSetVtpparam              DVRGetSet = 1085 //设置虚拟线圈参数
	NetDvrGetSnapenablecfg         DVRGetSet = 1086 //获取抓拍机使能参数
	NetDvrSetSnapenablecfg         DVRGetSet = 1087 //设置抓拍机使能参数
	NetDvrGetPostepolicecfg        DVRGetSet = 1088 //获取卡口电警参数
	NetDvrSetPostepolicecfg        DVRGetSet = 1089 //设置卡口电警参数
	NetDvrGetJpegcfgV30            DVRGetSet = 1090 //获取抓图的JPEG参数(基线)
	NetDvrSetJpegcfgV30            DVRGetSet = 1091 //设置抓图的JPEG参数(基线)
	NetDvrGetSprcfg                DVRGetSet = 1092 //获取车牌识别参数
	NetDvrSetSprcfg                DVRGetSet = 1093 //设置车牌识别参数
	NetDvrGetPlccfg                DVRGetSet = 1094 //获取车牌亮度补偿参数
	NetDvrSetPlccfg                DVRGetSet = 1095 //设置车牌亮度补偿参数
	NetDvrGetDevicestatecfg        DVRGetSet = 1096 //获取设备当前状态参数
	NetDvrSetCalibrateTime         DVRGetSet = 1097 //设置扩展时间校时
	NetDvrGetCalibrateTime         DVRGetSet = 1098 //获取扩展时间校时
	NetDvrGetDevicecfgV40          DVRGetSet = 1100 //获取扩展设备参数
	NetDvrSetDevicecfgV40          DVRGetSet = 1101 //设置扩展设备参数
	NetDvrGetZerochancfg           DVRGetSet = 1102 //获取零通道压缩参数
	NetDvrSetZerochancfg           DVRGetSet = 1103 //设置零通道压缩参数
	NetDvrGetZeroPreviewcfgV30     DVRGetSet = 1104 // 获取零通道预览参数配置
	NetDvrSetZeroPreviewcfgV30     DVRGetSet = 1105 // 设置零通道预览参数配置
	NetDvrSetZeroZoom              DVRGetSet = 1106 //设置零通道的缩放配置
	NetDvrGetZeroZoom              DVRGetSet = 1107 //获取零通道的缩放配置
	NetDvrNatassociatecfgGet       DVRGetSet = 1110 //获取NAT功能相关信息
	NetDvrNatassociatecfgSet       DVRGetSet = 1111 //设置NAT功能相关信息
	NetDvrGetSnmpcfg               DVRGetSet = 1112 //获取SNMP参数
	NetDvrSetSnmpcfg               DVRGetSet = 1113 //设置SNMP参数
	NetDvrGetSnmpcfgV30            DVRGetSet = 1114 //获取SNMPv30参数
	NetDvrSetSnmpcfgV30            DVRGetSet = 1115 //设置SNMPv30参数
	NetDvrVideoplatformalarmcfgGet DVRGetSet = 1130 //获取视频综合平台报警配置
	NetDvrVideoplatformalarmcfgSet DVRGetSet = 1131 //设置视频综合平台报警配置
	NetDvrGetRaidAdapterInfo       DVRGetSet = 1134 // 获取适配器信息
	NetDvrSetRaidAdapterInfo       DVRGetSet = 1135 // 设置适配器信息
	NetDvrMatrixBigscreencfgGet    DVRGetSet = 1140 //获取大屏拼接参数
	NetDvrMatrixBigscreencfgSet    DVRGetSet = 1141 //设置大屏拼接参数
	NetDvrGetMbPlatformpara        DVRGetSet = 1145 //获取平台登录参数
	NetDvrSetMbPlatformpara        DVRGetSet = 1146 //设置平台登录参数
	NetDvrGetMbDevstatus           DVRGetSet = 1147 //获取移动设备状态
	NetDvrGetDecoderJointChan      DVRGetSet = 1151 //获取解码关联通道
	NetDvrSetDecoderJointChan      DVRGetSet = 1152 //设置解码关联通道

	//多网卡配置
	NetDvrGetNetcfgMulti            DVRGetSet = 1161 //获取多网卡配置
	NetDvrSetNetcfgMulti            DVRGetSet = 1162 //设置多网卡配置
	NetDvrGetNetcfgMultiV50         DVRGetSet = 1163 //获取多网卡配置(分组)
	NetDvrSetNetcfgMultiV50         DVRGetSet = 1164 //设置多网卡配置(分组)
	NetDvrGetCmspara                DVRGetSet = 1170 //获取平台参数
	NetDvrSetCmspara                DVRGetSet = 1171 //设置平台参数
	NetDvrGetDialstatus             DVRGetSet = 1172 //获取拨号状态参数
	NetDvrGetSmsrelativepara        DVRGetSet = 1173 //获取短信相关参数
	NetDvrSetSmsrelativepara        DVRGetSet = 1174 //设置短信相关参数
	NetDvrGetPinstatus              DVRGetSet = 1175 //获取Pin状态
	NetDvrSetPincmd                 DVRGetSet = 1176 //设置PIN命令
	NetDvrSetSensorCfg              DVRGetSet = 1180 //设置模拟量参数
	NetDvrGetSensorCfg              DVRGetSet = 1181 //获取模拟量参数
	NetDvrSetAlarminParam           DVRGetSet = 1182 //设置报警输入参数
	NetDvrGetAlarminParam           DVRGetSet = 1183 //获取报警输入参数
	NetDvrSetAlarmoutParam          DVRGetSet = 1184 //设置报警输出参数
	NetDvrGetAlarmoutParam          DVRGetSet = 1185 //获取报警输出参数
	NetDvrSetSirenParam             DVRGetSet = 1186 //设置警号参数
	NetDvrGetSirenParam             DVRGetSet = 1187 //获取警号参数
	NetDvrSetAlarmRs485Cfg          DVRGetSet = 1188 //设置报警主机485参数
	NetDvrGetAlarmRs485Cfg          DVRGetSet = 1189 //获取报警主机485参数
	NetDvrGetAlarmhostMainStatus    DVRGetSet = 1190 //获取报警主机主要状态
	NetDvrGetAlarmhostOtherStatus   DVRGetSet = 1191 //获取报警主机其他状态
	NetDvrSetAlarmhostEnablecfg     DVRGetSet = 1192 //获取报警主机使能状态
	NetDvrGetAlarmhostEnablecfg     DVRGetSet = 1193 //设置报警主机使能状态
	NetDvrSetAlarmCamcfg            DVRGetSet = 1194 //设置视频综合平台报警触发CAM操作配置
	NetDvrGetAlarmCamcfg            DVRGetSet = 1195 //设置视频综合平台报警触发CAM操作配置
	NetDvrGetGatewayCfg             DVRGetSet = 1196 //获取门禁参数配置
	NetDvrSetGatewayCfg             DVRGetSet = 1197 //设置门禁参数配置
	NetDvrGetAlarmdialmodecfg       DVRGetSet = 1198 //获取报警主机拨号参数
	NetDvrSetAlarmdialmodecfg       DVRGetSet = 1199 //设置报警主机拨号参数
	NetDvrSetAlarminParamV50        DVRGetSet = 1200 // 设置防区参数V50
	NetDvrGetAlarminParamV50        DVRGetSet = 1201 // 获取防区参数V50
	NetDvrSetWincfg                 DVRGetSet = 1202 //窗口参数设置
	NetDvrGetAlarmhostdialsetupmode DVRGetSet = 1204 //获取报警主机拨号启用方式
	NetDvrSetAlarmhostdialsetupmode DVRGetSet = 1205 //设置报警主机拨号启用方式

	//视频报警主机海外版命令(视频报警主机 V1.3)
	NetDvrSetSubsystemAlarm        DVRGetSet = 1210 //设置子系统布/撤防
	NetDvrGetSubsystemAlarm        DVRGetSet = 1211 //获取子系统布/撤防
	NetDvrGetWhitelistAlarm        DVRGetSet = 1215 //获取白名单参数
	NetDvrSetWhitelistAlarm        DVRGetSet = 1216 //设置白名单参数
	NetDvrGetAlarmhostModuleList   DVRGetSet = 1222 //获取所有模块
	NetDvrSetPriorAlarm            DVRGetSet = 1223 //设置子系统布/撤防
	NetDvrGetPriorAlarm            DVRGetSet = 1224 //获取子系统布/撤防
	NetDvrSetTamperAlarminParam    DVRGetSet = 1225 // 设置防区防拆参数
	NetDvrGetTamperAlarminParam    DVRGetSet = 1226 // 获取防区防拆参数
	NetDvrGetHolidayParamCfg       DVRGetSet = 1240 // 获取节假日参数
	NetDvrSetHolidayParamCfg       DVRGetSet = 1241 // 设置节假日参数
	NetDvrGetMotionHolidayHandle   DVRGetSet = 1242 // 获取移动侦测假日报警处理方式
	NetDvrSetMotionHolidayHandle   DVRGetSet = 1243 // 获取移动侦测假日报警处理方式
	NetDvrGetVilostHolidayHandle   DVRGetSet = 1244 // 获取视频信号丢失假日报警处理方式
	NetDvrSetVilostHolidayHandle   DVRGetSet = 1245 // 获取视频信号丢失假日报警处理方式
	NetDvrGetHideHolidayHandle     DVRGetSet = 1246 // 获取遮盖假日报警处理方式
	NetDvrSetHideHolidayHandle     DVRGetSet = 1247 // 设置遮盖假日报警处理方式
	NetDvrGetAlarminHolidayHandle  DVRGetSet = 1248 // 获取报警输入假日报警处理方式
	NetDvrSetAlarminHolidayHandle  DVRGetSet = 1249 // 设置报警输入假日报警处理方式
	NetDvrGetAlarmoutHolidayHandle DVRGetSet = 1250 // 获取报警输出假日报警处理方式
	NetDvrSetAlarmoutHolidayHandle DVRGetSet = 1251 // 设置报警输出假日报警处理方式
	NetDvrGetHolidayRecord         DVRGetSet = 1252 // 获取假日录像参数
	NetDvrSetHolidayRecord         DVRGetSet = 1253 // 设置假日录像参数
	NetDvrGetNetworkBonding        DVRGetSet = 1254 // 获取BONDING网络参数
	NetDvrSetNetworkBonding        DVRGetSet = 1255 // 设置BONDING网络参数
	NetDvrGetLinkStatus            DVRGetSet = 1256 // 获取通道IP工作状态
	NetDvrGetDiskQuotaCfg          DVRGetSet = 1278 // 获取磁盘配额信息
	NetDvrSetDiskQuotaCfg          DVRGetSet = 1279 // 设置磁盘配额信息
	NetDvrGetJpegCaptureCfg        DVRGetSet = 1280 // 获取DVR抓图配置
	NetDvrSetJpegCaptureCfg        DVRGetSet = 1281 // 设置DVR抓图配置
	NetDvrGetSchedCapturecfg       DVRGetSet = 1282 // 获取抓图计划
	NetDvrSetSchedCapturecfg       DVRGetSet = 1283 // 设置抓图计划
	NetDvrGetVgaPreviewcfg         DVRGetSet = 1284 // 获取VGA预览配置
	NetDvrSetVgaPreviewcfg         DVRGetSet = 1285 // 设置VGA预览配置
	NetDvrGetVideoInputEffect      DVRGetSet = 1286 // 获取通道视频输入图像参数
	NetDvrSetVideoInputEffect      DVRGetSet = 1287 // 设置通道视频输入图像参数

	NetDvrGetStorageServerSwitch DVRGetSet = 1290 //获取存储服务器开关状态
	NetDvrSetStorageServerSwitch DVRGetSet = 1291 //设置存储服务器开关状态

	NetDvrGetDiskQuotaCfgV60 DVRGetSet = 1292 //获取磁盘配额信息V60
	NetDvrSetDiskQuotaCfgV60 DVRGetSet = 1293 //设置磁盘配额信息V60
	NetDvrGetOpticalChannel  DVRGetSet = 1300 //获取光端子系统通道关联信息
	NetDvrSetOpticalChannel  DVRGetSet = 1301 //设置光端子系统通道关联信息
	NetDvrGetFiberCascade    DVRGetSet = 1302 //获取光纤级联模式
	NetDvrSetFiberCascade    DVRGetSet = 1303 //设置光纤级联模式
	NetDvrGetSpartanStatus   DVRGetSet = 1304 //获取畅显状态
	NetDvrSetSpartanStatus   DVRGetSet = 1305 //设置畅显状态
	NetDvrGetEthernetChannel DVRGetSet = 1306 //获取端口聚合参数
	NetDvrSetEthermetChannel DVRGetSet = 1307 //设置端口聚合参数
	NetDvrOpticalReboot      DVRGetSet = 1320 //光端机重启
	NetDvrSetAudiochanCfg    DVRGetSet = 1321 //设置音频切换参数
	NetDvrGetAudiochanCfg    DVRGetSet = 1322 //获取音频切换参数
	//SDI矩阵1.0
	NetDvrSetMatrixBaseCfg         DVRGetSet = 1332 //设置矩阵基本参数
	NetDvrGetMatrixBaseCfg         DVRGetSet = 1333 //获取矩阵基本参数
	NetDvrSwitchMatrixIo           DVRGetSet = 1334 //矩阵输入输出切换
	NetDvrGetMatrixIoRelation      DVRGetSet = 1335 //获取矩阵输入输入关联关系
	NetDvrV6PsubsystemaramGet      DVRGetSet = 1501 //获取V6子系统配置
	NetDvrV6PsubsystemaramSet      DVRGetSet = 1502 //设置V6子系统配置
	NetDvrGetAllwincfg             DVRGetSet = 1503 //窗口参数获取
	NetDvrBigscreenassociatecfgGet DVRGetSet = 1511 //获取大屏关联配置
	NetDvrBigscreenassociatecfgSet DVRGetSet = 1512 //设置大屏关联配置

	//1200起
	NetDvrGetscreeninfo                 DVRGetSet = 1601 //获取大屏信息配置
	NetDvrSetscreeninfo                 DVRGetSet = 1602 //设置大屏信息配置
	NetDvrGetScreenWincfg               DVRGetSet = 1603 //单个窗口参数获取
	NetDvrLayoutlistGet                 DVRGetSet = 1605 //获取布局列表
	NetDvrSetLayoutcfg                  DVRGetSet = 1606 //布局设置
	NetDvrLayoutctrl                    DVRGetSet = 1607 //布局控制，1-open，2-close
	NetDvrInputlistGet                  DVRGetSet = 1608 //获取输入信号源列表
	NetDvrSetInputstreamcfg             DVRGetSet = 1609 //输入信号源设置
	NetDvrOutputSet                     DVRGetSet = 1610 //输出参数设置
	NetDvrOutputGet                     DVRGetSet = 1611 //输出参数获取
	NetDvrSetOsdcfg                     DVRGetSet = 1612 //OSD参数设置
	NetDvrGetOsdcfg                     DVRGetSet = 1613 //OSD参数获取
	NetDvrBigscreenGetserial            DVRGetSet = 1614 //获取大屏串口信息
	NetDvrGetPlanlist                   DVRGetSet = 1615 //获取预案列表
	NetDvrSetPlan                       DVRGetSet = 1616 //设置预案
	NetDvrCtrlPlan                      DVRGetSet = 1617 //控制预案
	NetDvrGetDeviceRunStatus            DVRGetSet = 1618 //获取设备运行状态
	NetDvrGetExternalMatrixCfg          DVRGetSet = 1619 //获取矩阵信息
	NetDvrSetExternalMatrixCfg          DVRGetSet = 1620 //设置矩阵信息
	NetDvrGetOutputScreenRelation       DVRGetSet = 1621 //获取输出和屏幕的绑定关系
	NetDvrSetOutputScreenRelation       DVRGetSet = 1622 //设置输出和屏幕的绑定关系
	NetDvrGetVcsUserCfg                 DVRGetSet = 1623 //获取用户信息配置
	NetDvrSetVcsUserCfg                 DVRGetSet = 1624 //设置用户信息配置
	NetDvrControlScreen                 DVRGetSet = 1625 //屏幕控制
	NetDvrGetExternalMatrixCfgV50       DVRGetSet = 1626 //获取矩阵信息
	NetDvrSetExternalMatrixCfgV50       DVRGetSet = 1627 //设置矩阵信息
	NetDvrGetDevBaseinfo                DVRGetSet = 1650 //获取单个设备信息
	NetDvrSetDevBaseinfo                DVRGetSet = 1651 //设置单个设备信息
	NetDvrGetDevNetinfo                 DVRGetSet = 1652 //获取设备的网络信息
	NetDvrSetDevNetinfo                 DVRGetSet = 1653 //设置设备的网络信息
	NetDvrGetSignalSourceInfo           DVRGetSet = 1654 //获取信号源信息
	NetDvrSetSignalSourceInfo           DVRGetSet = 1655 //设置信号源信息
	NetDvrAdjustPicV40                  DVRGetSet = 1656 //图像微调
	NetDvrRestoreV40                    DVRGetSet = 1657 //恢复默认参数
	NetDvrSetNetSignal                  DVRGetSet = 1658 //设置网络信号源
	NetDvrRebootV40                     DVRGetSet = 1659 //重启
	NetDvrControlPictureV41             DVRGetSet = 1660 //图片控制V41
	NetDvrGetAutoRebootCfg              DVRGetSet = 1710 //获取自动重启参数
	NetDvrSetAutoRebootCfg              DVRGetSet = 1711 //设置自动重启参数
	NetDvrGetTrunkUseState              DVRGetSet = 1713 //获取指定干线使用状态
	NetDvrSetPtzCtrlInfo                DVRGetSet = 1714 //设置PTZ控制参数
	NetDvrGetPtzCtrlInfo                DVRGetSet = 1715 //获取PTZ控制参数
	NetDvrGetPtzStatus                  DVRGetSet = 1716 //获取PTZ状态
	NetDvrGetDispRouteList              DVRGetSet = 1717 //获取显示路径列表
	NetDvrGetDecResourceList            DVRGetSet = 1720 //获取可用解码资源列表
	NetDvrSetDecResourceList            DVRGetSet = 1721 //预分配解码资源
	NetDvrGetDecYuv                     DVRGetSet = 1722 //获取解码通道关联YUV输出参数
	NetDvrSetDecYuv                     DVRGetSet = 1723 //设置解码通道关联YUV输出参数
	NetDvrGetDecResouce                 DVRGetSet = 1724 //向视频综合平台申请解码资源
	NetDvrFreeDecResource               DVRGetSet = 1725 //释放解码资源
	NetDvrSetVideowalldisplaymode       DVRGetSet = 1730 //设置电视墙拼接模式
	NetDvrGetVideowalldisplaymode       DVRGetSet = 1731 //获取电视墙拼接模式
	NetDvrGetVideowalldisplayno         DVRGetSet = 1732 //获取设备显示输出号
	NetDvrSetVideowalldisplayposition   DVRGetSet = 1733 //设置显示输出位置参数
	NetDvrGetVideowalldisplayposition   DVRGetSet = 1734 //获取显示输出位置参数
	NetDvrGetVideowallwindowposition    DVRGetSet = 1735 //获取电视墙窗口参数
	NetDvrSetVideowallwindowposition    DVRGetSet = 1736 //设置电视墙窗口参数
	NetDvrVideowallwindowCloseall       DVRGetSet = 1737 //电视墙关闭所有窗口
	NetDvrSetVirtualled                 DVRGetSet = 1738 //虚拟LED设置
	NetDvrGetVirtualled                 DVRGetSet = 1739 //虚拟LED获取
	NetDvrGetImageCutMode               DVRGetSet = 1740 //获取图像切割模式
	NetDvrSetImageCutMode               DVRGetSet = 1741 //设置图像切割模式
	NetDvrGetUsingSerialport            DVRGetSet = 1742 //获取当前使用串口
	NetDvrSetUsingSerialport            DVRGetSet = 1743 //设置当前使用串口
	NetDvrSceneControl                  DVRGetSet = 1744 //场景控制
	NetDvrGetCurrentScene               DVRGetSet = 1745 //获取当前场景号
	NetDvrGetVwSceneParam               DVRGetSet = 1746 //获取电视墙场景模式参数
	NetDvrSetVwSceneParam               DVRGetSet = 1747 //设置电视墙场景模式参数
	NetDvrDisplayChannoControl          DVRGetSet = 1748 //电视墙显示编号控制
	NetDvrGetWinDecInfo                 DVRGetSet = 1749 //获取窗口解码信息（批量）
	NetDvrResetVideowalldisplayposition DVRGetSet = 1750 //解除电视墙输出接口绑定
	NetDvrSetVwAudioCfg                 DVRGetSet = 1752 //设置音频切换参数
	NetDvrGetVwAudioCfg                 DVRGetSet = 1753 //获取音频切换参数
	NetDvrGetGbt28181DecchaninfoCfg     DVRGetSet = 1754 //获取GBT28181协议接入设备的解码通道信息
	NetDvrSetGbt28181DecchaninfoCfg     DVRGetSet = 1755 //设置GBT28181协议接入设备的解码通道信息
	NetDvrSetMainboardSerial            DVRGetSet = 1756 //设置主控板串口参数
	NetDvrGetMainboardSerial            DVRGetSet = 1757 //获取主控板串口参数
	NetDvrGetSubboardInfo               DVRGetSet = 1758 //获取子板信息
	NetDvrGetSubboardException          DVRGetSet = 1759 //获取异常子板异常信息
	NetDvrGetCamerachanSerialcfg        DVRGetSet = 1760 //获取Camera通道绑定串口配置
	NetDvrSetCamerachanSerialcfg        DVRGetSet = 1761 //设置Camera通道绑定串口配置
	NetDvrGetMatrixStatus               DVRGetSet = 1762 //获取视频综合平台状态
	NetSetMultifunctionSerialcfg        DVRGetSet = 1763 //设置多功能串口配置
	NetGetMultifunctionSerialcfg        DVRGetSet = 1764 //获取多功能串口配置
	NetDvrPtz_3DSpeed                   DVRGetSet = 1765 // 3维带速度的云台控制
	NetDvrGetSignalJoint                DVRGetSet = 1766 //获取信号源绑定配置
	NetDvrSetSignalJoint                DVRGetSet = 1767 //设置信号源绑定配置
	NetDvrSignalCut                     DVRGetSet = 1768 //信号源裁剪
	NetDvrDynamicDecodeBatch            DVRGetSet = 1769 //批量动态解码
	NetDvrDecswitchSetBatch             DVRGetSet = 1770 //批量设置解码通道开关
	NetDvrDecswitchGetBatch             DVRGetSet = 1771 //批量获取解码通道开关
	NetDvrGetAllSignalJoint             DVRGetSet = 1772 //获取所有信号源绑定配置
	NetDvrGetPlayingPlan                DVRGetSet = 1773 //获取正在执行预案
	NetDvrWallRelationGet               DVRGetSet = 1774 //获取设备墙与物理墙的关联
	NetDvrWallRelationSet               DVRGetSet = 1775 //设置设备墙与物理墙的关联
	NetDvrSetInputstreamcfgV40          DVRGetSet = 1776 //输入信号源设置
	NetDvrPtzcfgInputstreamGet          DVRGetSet = 1777 //获取输入源反向云台控制配置
	NetDvrPtzcfgInputstreamSet          DVRGetSet = 1778 //设置输入源反向云台控制配置
	NetDvrSignalCutparamGet             DVRGetSet = 1779 //获取信号源裁剪参数
	NetDvrGetSubsystemNetcfg            DVRGetSet = 1780 //获取子系统网卡参数
	NetDvrSetSubsystemNetcfg            DVRGetSet = 1781 //设置子系统网卡参数
	NetDvrDelSignalJoint                DVRGetSet = 1782 //删除拼接信号源
	NetDvrGetPictureInfo                DVRGetSet = 1783 //获取图片信息
	NetDvrSetPictureInfo                DVRGetSet = 1784 //设置图片信息
	NetDvrGetVideoInfo                  DVRGetSet = 1785 //获取视频信息
	NetDvrSetVideoInfo                  DVRGetSet = 1786 //设置视频信息
	NetDvrSetPlaylist                   DVRGetSet = 1787 //设置播放列表
	NetDvrGetPlaylist                   DVRGetSet = 1788 //获取播放列表
	NetDvrGetAllPlaylist                DVRGetSet = 1789 //获取所有播放列表
	NetDvrPlayitemControl               DVRGetSet = 1790 //播放项操作
	NetDvrSetPlayplanTemplate           DVRGetSet = 1791 //设置播放计划模板
	NetDvrGetPlayplanTemplate           DVRGetSet = 1792 //获取播放计划
	NetDvrGetAllPlayplanTemplate        DVRGetSet = 1793 //获取所有播放计划
	NetDvrSetWindowPlayplan             DVRGetSet = 1794 //设置窗口播放计划
	NetDvrGetWindowPlayplan             DVRGetSet = 1795 //获取窗口播放计划
	NetDvrToplayItem                    DVRGetSet = 1796 //指定播放
	NetDvrDevicePlayControl             DVRGetSet = 1797 //设备播放控制
	NetDvrGetPlayInfo                   DVRGetSet = 1798 //获取当前播放信息
	NetDvrGetAllPictureInfo             DVRGetSet = 1799 //获取图片信息
	NetDvrGetAllVideoInfo               DVRGetSet = 1800 //获取视频信息
	NetDvrDeleteVideoFile               DVRGetSet = 1801 //删除视频

	NetDvrGetAlarmhostsubsystemCfg         DVRGetSet = 2001 //报警主机获取子系统参数
	NetDvrSetAlarmhostsubsystemCfg         DVRGetSet = 2002 //报警主机设置子系统参数
	NetDvrGetextendalarmininfo             DVRGetSet = 2003 //获取防区编号信息
	NetDvrModifyalarminno                  DVRGetSet = 2004 //修改防区编号信息
	NetDvrGetAlarmhostWirelessNetworkCfg   DVRGetSet = 2005 //获取无线网络参数配置
	NetDvrSetAlarmhostWirelessNetworkCfg   DVRGetSet = 2006 //设置无线网络参数配置
	NetDvrGetAlarmhostNetcfg               DVRGetSet = 2007 //获取网络参数配置
	NetDvrSetAlarmhostNetcfg               DVRGetSet = 2008 //设置网络参数配置
	NetDvrGetLedScreenCfg                  DVRGetSet = 2009 // 获取LED屏幕参数
	NetDvrSetLedScreenCfg                  DVRGetSet = 2010 // 设置LED屏幕参数
	NetDvrGetLedContentCfg                 DVRGetSet = 2011 // 获取LED屏显内容
	NetDvrSetLedContentCfg                 DVRGetSet = 2012 // 设置LED屏显内容
	NetDvrTurnonLed                        DVRGetSet = 2013 // 打开LED屏
	NetDvrTurnoffLed                       DVRGetSet = 2014 // 关闭LED屏
	NetDvrGetLedTimerSwitch                DVRGetSet = 2015 // 获取LED屏定时开关参数
	NetDvrSetLedTimerSwitch                DVRGetSet = 2016 // 设置LED屏定时开关参数
	NetDvrSetLedBrightness                 DVRGetSet = 2017 // 手动设置LED屏亮度
	NetDvrGetLedTimerBrightness            DVRGetSet = 2018 //设置分时LED屏亮度
	NetDvrSetLedTimerBrightness            DVRGetSet = 2019 //获取分时LED屏亮度
	NetDvrLedChecktime                     DVRGetSet = 2020 //LED校时
	NetDvrGetAlarmhostAudioAssociateAlarm  DVRGetSet = 2021 //获取音频跟随报警事件
	NetDvrSetAlarmhostAudioAssociateAlarm  DVRGetSet = 2022 //设置音频跟随报警事件
	NetDvrGetLedStatus                     DVRGetSet = 2023 //获取LED屏状态
	NetDvrCloseSubsystemFaultAlarm         DVRGetSet = 2027 //关闭子系统故障提示音
	NetDvrSetSubsystemBypass               DVRGetSet = 2028 //子系统旁路
	NetDvrCancelSubsystemBypass            DVRGetSet = 2029 //子系统旁路恢复
	NetDvrGetAlarmhostSubsystemCfgEx       DVRGetSet = 2030 //获取子系统扩展参数
	NetDvrSetAlarmhostSubsystemCfgEx       DVRGetSet = 2031 //设置子系统扩展参数
	NetDvrGetAlarmhostPrinterCfg           DVRGetSet = 2032 //获取打印机打印使能
	NetDvrSetAlarmhostPrinterCfg           DVRGetSet = 2033 //设置打印机打印使能
	NetDvrGetAlarmhostZoneListInSubsystem  DVRGetSet = 2034 //获取指定子系统内的所有防区
	NetDvrGetAlarmhostTriggerList          DVRGetSet = 2035 //获取所有触发器
	NetDvrArmAlarmhostSubsystem            DVRGetSet = 2036 //按布防类型对子系统布防
	NetDvrGetAlarmhostEventTrigAlarmoutCfg DVRGetSet = 2037 // 获取事件触发报警输出配置
	NetDvrSetAlarmhostEventTrigAlarmoutCfg DVRGetSet = 2038 // 设置事件触发报警输出配置
	NetDvrGetAlarmhostFaultCfg             DVRGetSet = 2039 // 获取故障处理配置
	NetDvrSetAlarmhostFaultCfg             DVRGetSet = 2040 // 设置故障处理配置
	NetDvrSearchArmhostExternalModule      DVRGetSet = 2041 //自动搜索
	NetDvrRegisterAlarmhostExternalModule  DVRGetSet = 2042 //自动注册
	NetDvrCloseAlarmhostOverallFaultAlarm  DVRGetSet = 2043 //关闭全局故障提示音
	NetDvrGetSafetycabinWorkMode           DVRGetSet = 2044 //获取防护舱工作模式参数
	NetDvrSetSafetycabinWorkMode           DVRGetSet = 2045 //设置防护舱工作模式参数
	NetDvrGetSafetycabinPersonSignalCfg    DVRGetSet = 2046 //获取防护舱人信号探测参数
	NetDvrSetSafetycabinPersonSignalCfg    DVRGetSet = 2047 //设置防护舱人信号探测参数
	NetDvrGetAlarmhostModuleCfg            DVRGetSet = 2048 //获取模块信息
	//NET_DVR_SET_ALARMHOST_MODULE_CFG DVRGetSet = 2049 //设置模块信息(预留)

	NetDvrGetAlarmhostExternalDeviceState      DVRGetSet = 2050 //获取485外接设备状态
	NetDvrSetAlarmhostExternalDeviceLimitValue DVRGetSet = 2051 //设置外接设备报警限值
	NetDvrGetAlarmhostExternalDeviceLimitValue DVRGetSet = 2052 //获取外接设备报警限值
	NetDvrGetAlarmhostSensorJointCfg           DVRGetSet = 2053 // 获取模拟量关联配置
	NetDvrSetAlarmhostSensorJointCfg           DVRGetSet = 2054 // 设置模拟量关联配置
	NetDvrSetAlarmhostRs485SlotCfg             DVRGetSet = 2055 // 设置报警主机485槽位参数
	NetDvrGetAlarmhostRs485SlotCfg             DVRGetSet = 2056 // 获取报警主机485槽位参数
	NetDvrGetAllVariableInfo                   DVRGetSet = 2057 // 获取所有变量元素信息
	NetDvrGetAlarmPointCfg                     DVRGetSet = 2058 // 获取点号信息
	NetDvrSetAlarmPointCfg                     DVRGetSet = 2059 // 设置点号信息
	NetDvrGetHistoryValue                      DVRGetSet = 2060 // 获取历史数据
	NetDvrGetAlarmhostAlarmMode                DVRGetSet = 2061 // 获取数据上传方式
	NetDvrSetAlarmhostAlarmMode                DVRGetSet = 2062 // 设置数据上传方式
	NetDvrGetAlarmhostSensorValue              DVRGetSet = 2063 // 获取模拟量实时数据
	NetDvrGetAlarmhostReportCenterV40          DVRGetSet = 2064 // 获取数据上传方式
	NetDvrSetAlarmhostReportCenterV40          DVRGetSet = 2065 // 设置数据上传方式
	NetDvrGetOutputScheduleRulecfg             DVRGetSet = 2068 // 获取时控输出参数
	NetDvrSetOutputScheduleRulecfg             DVRGetSet = 2069 // 设置时控输出参数
	NetDvrGetCmsCfg                            DVRGetSet = 2070
	NetDvrSetCmsCfg                            DVRGetSet = 2071
	NetDvrGetPassthroughCap                    DVRGetSet = 2073 //获取设备透传能力集
	NetDvrGetAlarmhostMainStatusV40            DVRGetSet = 2072 // 获取主要状态V40
	NetDvrGetAlarmhostMainStatusV51            DVRGetSet = 2083 // 获取主要状态V51

	/*************************************视频报警主机1.3 begin*************************************/
	NetDvrGetAlarmCaptrueCfg          DVRGetSet = 2074 //获取报警抓图参数配置
	NetDvrSetAlarmCaptrueCfg          DVRGetSet = 2075 //设置报警抓图参数配置
	NetDvrGetOneOutputSchRulecfgV40   DVRGetSet = 2078 // 获取单个时控输出参数V40
	NetDvrSetOneOutputSchRulecfgV40   DVRGetSet = 2079 // 设置单个时控输出参数V40
	NetDvrGetOutputScheduleRulecfgV40 DVRGetSet = 2080 // 获取时控输出参数V40
	NetDvrSetOutputScheduleRulecfgV40 DVRGetSet = 2081 // 设置时控输出参数V40
	NetDvrAlarmhostCloseSubsystem     DVRGetSet = 2082 //对子系统撤防操作
	/*************************************视频报警主机1.3 end**************************************/

	NetDvrGetWeekPlanCfg            DVRGetSet = 2100 //获取门状态周计划参数
	NetDvrSetWeekPlanCfg            DVRGetSet = 2101 //设置门状态周计划参数
	NetDvrGetDoorStatusHolidayPlan  DVRGetSet = 2102 //获取门状态假日计划参数
	NetDvrSetDoorStatusHolidayPlan  DVRGetSet = 2103 //设置门状态假日计划参数
	NetDvrGetDoorStatusHolidayGroup DVRGetSet = 2104 //获取门状态假日组参数
	NetDvrSetDoorStatusHolidayGroup DVRGetSet = 2105 //设置门状态假日组参数
	NetDvrGetDoorStatusPlanTemplate DVRGetSet = 2106 //获取门状态计划模板参数
	NetDvrSetDoorStatusPlanTemplate DVRGetSet = 2107 //设置门状态计划模板参数
	NetDvrGetDoorCfg                DVRGetSet = 2108 //获取门参数
	NetDvrSetDoorCfg                DVRGetSet = 2109 //设置门参数
	NetDvrGetDoorStatusPlan         DVRGetSet = 2110 //获取门状态计划参数
	NetDvrSetDoorStatusPlan         DVRGetSet = 2111 //设置门状态计划参数
	NetDvrGetGroupCfg               DVRGetSet = 2112 //获取群组参数
	NetDvrSetGroupCfg               DVRGetSet = 2113 //设置群组参数
	NetDvrGetMultiCardCfg           DVRGetSet = 2114 //获取多重卡参数
	NetDvrSetMultiCardCfg           DVRGetSet = 2115 //设置多重卡参数
	NetDvrGetCardCfg                DVRGetSet = 2116 //获取卡参数
	NetDvrSetCardCfg                DVRGetSet = 2117 //设置卡参数
	NetDvrClearAcsParam             DVRGetSet = 2118 //清空门禁主机参数
	NetDvrGetSneakCfg               DVRGetSet = 2119 //获取反潜回参数
	NetDvrSetSneakCfg               DVRGetSet = 2120 //设置反潜回参数
	NetDvrGetMultiDoorInterlockCfg  DVRGetSet = 2121 //获取多门互锁参数
	NetDvrSetMultiDoorInterlockCfg  DVRGetSet = 2122 //设置多门互锁参数
	NetDvrGetAcsWorkStatus          DVRGetSet = 2123 //获取门禁主机工作状态
	NetDvrGetVerifyWeekPlan         DVRGetSet = 2124 //获取读卡器验证方式周计划参数
	NetDvrSetVerifyWeekPlan         DVRGetSet = 2125 //设置读卡器验证方式周计划参数
	NetDvrGetCardRightWeekPlan      DVRGetSet = 2126 //获取卡权限周计划参数
	NetDvrSetCardRightWeekPlan      DVRGetSet = 2127 //设置卡权限周计划参数
	NetDvrGetVerifyHolidayPlan      DVRGetSet = 2128 //获取读卡器验证方式假日计划参数
	NetDvrSetVerifyHolidayPlan      DVRGetSet = 2129 //设置读卡器验证方式假日计划参数
	NetDvrGetCardRightHolidayPlan   DVRGetSet = 2130 //获取卡权限假日计划参数
	NetDvrSetCardRightHolidayPlan   DVRGetSet = 2131 //设置卡权限假日计划参数
	NetDvrGetVerifyHolidayGroup     DVRGetSet = 2132 //获取读卡器验证方式假日组参数
	NetDvrSetVerifyHolidayGroup     DVRGetSet = 2133 //设置读卡器验证方式假日组参数
	NetDvrGetCardRightHolidayGroup  DVRGetSet = 2134 //获取卡权限假日组参数
	NetDvrSetCardRightHolidayGroup  DVRGetSet = 2135 //设置卡权限假日组参数
	NetDvrGetVerifyPlanTemplate     DVRGetSet = 2136 //获取读卡器验证方式计划模板参数
	NetDvrSetVerifyPlanTemplate     DVRGetSet = 2137 //设置读卡器验证方式计划模板参数
	NetDvrGetCardRightPlanTemplate  DVRGetSet = 2138 //获取卡权限计划模板参数
	NetDvrSetCardRightPlanTemplate  DVRGetSet = 2139 //设置卡权限计划模板参数
	NetDvrGetCardReaderCfg          DVRGetSet = 2140 //获取读卡器参数
	NetDvrSetCardReaderCfg          DVRGetSet = 2141 //设置读卡器参数
	NetDvrGetCardReaderPlan         DVRGetSet = 2142 //获取读卡器验证计划参数
	NetDvrSetCardReaderPlan         DVRGetSet = 2143 //设置读卡器验证计划参数
	NetDvrGetCaseSensorCfg          DVRGetSet = 2144 //获取事件触发器参数
	NetDvrSetCaseSensorCfg          DVRGetSet = 2145 //设置事件触发器参数
	NetDvrGetCardReaderAntiSneakCfg DVRGetSet = 2146 //获取读卡器反潜回参数
	NetDvrSetCardReaderAntiSneakCfg DVRGetSet = 2147 //设置读卡器反潜回参数
	NetDvrGetPhoneDoorRightCfg      DVRGetSet = 2148 //获取手机关联门权限参数
	NetDvrSetPhoneDoorRightCfg      DVRGetSet = 2149 //获取手机关联门权限参数
	NetDvrGetFingerprintCfg         DVRGetSet = 2150 //获取指纹参数
	NetDvrSetFingerprintCfg         DVRGetSet = 2151 //设置指纹参数
	NetDvrDelFingerprintCfg         DVRGetSet = 2152 //删除指纹参数
	NetDvrGetEventCardLinkageCfg    DVRGetSet = 2153 //获取事件卡号联动配置参数
	NetDvrSetEventCardLinkageCfg    DVRGetSet = 2154 //设置事件卡号联动配置参数
	NetDvrGetAntiSneakHostCfg       DVRGetSet = 2155 //获取主机组反潜回参数
	NetDvrSetAntiSneakHostCfg       DVRGetSet = 2156 //设置主机组反潜回参数
	NetDvrGetReaderAntiSneakHostCfg DVRGetSet = 2157 //获取主机组读卡器反潜回参数
	NetDvrSetReaderAntiSneakHostCfg DVRGetSet = 2158 //设置主机组读卡器反潜回参数
	NetDvrGetAcsCfg                 DVRGetSet = 2159 //获取门禁主机参数
	NetDvrSetAcsCfg                 DVRGetSet = 2160 //设置门禁主机参数
	NetDvrGetCardPasswdCfg          DVRGetSet = 2161 //获取卡密码开门使能配置
	NetDvrSetCardPasswdCfg          DVRGetSet = 2162 //设置卡密码开门使能配置
	NetDvrGetCardUserinfoCfg        DVRGetSet = 2163 //获取卡号关联用户信息参数
	NetDvrSetCardUserinfoCfg        DVRGetSet = 2164 //设置卡号关联用户信息参数

	NetDvrGetAcsExternalDevCfg      DVRGetSet = 2165 //获取门禁主机串口外设参数
	NetDvrSetAcsExternalDevCfg      DVRGetSet = 2166 //设置门禁主机串口外设参数
	NetDvrGetPersonnelChannelCfg    DVRGetSet = 2167 //获取人员通道参数
	NetDvrSetPersonnelChannelCfg    DVRGetSet = 2168 //设置人员通道参数
	NetDvrSetPlatformVerifyCfg      DVRGetSet = 2169 //下发平台认证结果
	NetDvrGetPersonStatisticsCfg    DVRGetSet = 2170 //获取人数统计参数
	NetDvrSetPersonStatisticsCfg    DVRGetSet = 2171 //设置人数统计参数
	NetDvrGetAcsScreenDisplayCfg    DVRGetSet = 2172 //获取屏幕字符串显示参数
	NetDvrSetAcsScreenDisplayCfg    DVRGetSet = 2173 //设置屏幕字符串显示参数
	NetDvrGetGateTimeCfg            DVRGetSet = 2174 //获取人员通道闸门时间参数
	NetDvrSetGateTimeCfg            DVRGetSet = 2175 //设置人员通道闸门时间参数
	NetDvrGetLocalControllerStatus  DVRGetSet = 2176 //获取就地控制器状态
	NetDvrGetOnlineLocalController  DVRGetSet = 2177 //搜索在线就地控制器
	NetDvrGetCardCfgV50             DVRGetSet = 2178 //获取新卡参数(V50)
	NetDvrSetCardCfgV50             DVRGetSet = 2179 //设置新卡参数(V50)
	NetDvrGetAcsWorkStatusV50       DVRGetSet = 2180 //获取门禁主机工作状态(V50)
	NetDvrGetEventCardLinkageCfgV50 DVRGetSet = 2181 //获取事件卡号联动配置参数(V50)
	NetDvrSetEventCardLinkageCfgV50 DVRGetSet = 2182 //设置事件卡号联动配置参数(V50)
	NetDvrGetFingerprintCfgV50      DVRGetSet = 2183 //获取指纹参数V50
	NetDvrSetFingerprintCfgV50      DVRGetSet = 2184 //设置指纹参数V50

	NetDvrGetSafetycabinState DVRGetSet = 2197 //获取防护舱状态
	NetDvrGetRs485CascadeCfg  DVRGetSet = 2198 //获取Rs485级联设备配置
	NetDvrSetRs485CascadeCfg  DVRGetSet = 2199 //设置Rs485级联设备配置

	/*************************************视频报警主机2.0 begin*************************************/
	NetDvrGetRemotecontrollerPermisionCfg   DVRGetSet = 2200 //获取遥控器权限参数
	NetDvrSetRemotecontrollerPermisionCfg   DVRGetSet = 2201 //设置遥控器权限参数
	NetDvrGetKeyboardCfg                    DVRGetSet = 2202 //获取键盘参数配置
	NetDvrSetKeyboardCfg                    DVRGetSet = 2203 //设置键盘参数配置
	NetDvrGetAlarmhostWirelessBusinnessInfo DVRGetSet = 2204 //无线业务查询
	NetDvrGetAllRemotecontrollerList        DVRGetSet = 2205 //获取所有遥控器
	NetDvrGetPreviewDelayCfg                DVRGetSet = 2206 //获取延迟预览参数配置
	NetDvrSetPreviewDelayCfg                DVRGetSet = 2207 //设置延迟预览参数配置
	NetDvrGetZoneChannelLinkageCfg          DVRGetSet = 2208 //获取防区联动视频通道配置
	NetDvrSetZoneChannelLinkageCfg          DVRGetSet = 2209 //设置防区联动视频通道配置
	NetDvrGetCenterServerCfg                DVRGetSet = 2210 //获取报警中心服务器
	NetDvrSetCenterServerCfg                DVRGetSet = 2211 //设置报警中心服务器
	/*************************************视频报警主机2.0 end**************************************/

	/********************************一键式紧急报警产品V1.0.0 begin********************************/
	NetDvrGetEmergenceAlarmProductCap  DVRGetSet = 2212 //获取一键式紧急报警产品能力
	NetDvrGetCallWaittingCfgCap        DVRGetSet = 2213 //获取呼叫等待参数配置能力
	NetDvrGetCallWaittingCfg           DVRGetSet = 2214 //获取呼叫等待参数配置
	NetDvrSetCallWaittingCfg           DVRGetSet = 2215 //设置呼叫等待参数配置
	NetDvrGetAlarmLampCfgCap           DVRGetSet = 2216 //获取警灯参数配置能力
	NetDvrGetAlarmLampCfg              DVRGetSet = 2217 //获取警灯参数配置
	NetDvrSetAlarmLampCfg              DVRGetSet = 2218 //设置警灯参数配置
	NetDvrGetVoicePromptionCfgCap      DVRGetSet = 2219 //获取语音提示配置能力
	NetDvrGetVoicePromptionCfg         DVRGetSet = 2220 //获取语音提示配置
	NetDvrSetVoicePromptionCfg         DVRGetSet = 2221 //设置语音提示配置
	NetDvrGetEmergenceAlarmResponseCap DVRGetSet = 2222 //获取紧急报警处理能力
	NetDvrEmergenceAlarmResponseCtrl   DVRGetSet = 2223 //紧急报警处理控制
	/********************************一键式紧急报警产品V1.0.0 end**********************************/

	//网络报警主机 V2.2
	NetDvrGetAlarmhostNetcfgV50 DVRGetSet = 2224 //获取报警主机网络参数配置V50
	NetDvrSetAlarmhostNetcfgV50 DVRGetSet = 2225 //设置报警主机网络参数配置V50
	NetDvrRegisterAlarmRs485    DVRGetSet = 2226 //RS485重新注册
	/**********************************动环报警主机V3.0****************************************/

	NetDvrGetAlarminParamList DVRGetSet = 2227 //获取防区参数列表
	//无线报警主机1.0.0
	NetDvrGetAlarmhostOtherStatusV50   DVRGetSet = 2228 //获取报警主机其他状态v50
	NetDvrGetAlarmhostOtherStatusV51   DVRGetSet = 2236 //获取报警主机其他状态V51
	NetDvrGetAlarminAssociatedChanList DVRGetSet = 2229 //获取防区防区联动视频通道参数列表
	NetDvrGetAlarminTrigger            DVRGetSet = 2230 //获取报警主机防区联动配置
	NetDvrSetAlarminTrigger            DVRGetSet = 2231 //设置报警主机防区联动配置
	NetDvrGetEmergencyCallHelpTrigger  DVRGetSet = 2232 //获取报警主机紧急求助联动配置
	NetDvrSetEmergencyCallHelpTrigger  DVRGetSet = 2233 //设置报警主机紧急求助联动配置
	NetDvrGetConsultTrigger            DVRGetSet = 2234 //获取报警主机业务咨询联动配置
	NetDvrSetConsultTrigger            DVRGetSet = 2235 //设置报警主机业务咨询联动配置
	NetDvrGetAlarminParamListV50       DVRGetSet = 2237 //获取防区参数列表V50

	NetDvrGetCardRightWeekPlanV50     DVRGetSet = 2304 //获取卡权限周计划参数V50
	NetDvrSetCardRightWeekPlanV50     DVRGetSet = 2305 //设置卡权限周计划参数V50
	NetDvrGetCardRightHolidayPlanV50  DVRGetSet = 2310 //获取卡权限假日计划参数V50
	NetDvrSetCardRightHolidayPlanV50  DVRGetSet = 2311 //设置卡权限假日计划参数V50
	NetDvrGetCardRightHolidayGroupV50 DVRGetSet = 2316 //获卡权限假日组参数V50
	NetDvrSetCardRightHolidayGroupV50 DVRGetSet = 2317 //设置卡权限假日组参数V50
	NetDvrGetCardRightPlanTemplateV50 DVRGetSet = 2322 //获取卡权限计划模板参数V50
	NetDvrSetCardRightPlanTemplateV50 DVRGetSet = 2323 //设置卡权限计划模板参数V50
	/**********************************经济型指纹门禁产品V1.0 设备不做****************************************/
	NetDvrGetScheduleInfo          DVRGetSet = 2500 //获取排班信息
	NetDvrGetAttendanceSummaryInfo DVRGetSet = 2501 //获取考勤汇总信息
	NetDvrGetAttendanceRecordInfo  DVRGetSet = 2502 //获取考勤记录信息
	NetDvrGetAbnormalInfo          DVRGetSet = 2503 //获取异常统计信息
	/**********************************经济型指纹门禁产品V1.0****************************************/

	/*************************************视频门禁一体机1.0 begin**************************************/
	NetDvrCaptureFingerprintInfo DVRGetSet = 2504 //采集指纹信息
	/*************************************视频门禁一体机1.0 end**************************************/

	/*************************************嵌入式智能终端1.0 begin**************************************/
	NetDvrBulkUploadBlackListPicture DVRGetSet = 2520 //批量上传黑名单图片
	NetDvrBulkUploadIdBlackList      DVRGetSet = 2521 //批量上传身份证黑名单
	NetDvrGetFailedFaceInfo          DVRGetSet = 2522 //获取设备升级建模失败的人脸记录
	NetDvrGetFaceAndTemplate         DVRGetSet = 2523 //获取人脸及模板数据
	NetDvrSetFaceAndTemplate         DVRGetSet = 2524 //设置人脸及模板数据
	/*************************************嵌入式智能终端1.0 end**************************************/

	/*************************************人脸识别门禁一体机1.0 begin**************************************/
	NetDvrGetCardReaderCfgV50 DVRGetSet = 2505 //获取读卡器参数(V50)
	NetDvrSetCardReaderCfgV50 DVRGetSet = 2506 //设置读卡器参数(V50)
	NetDvrGetFaceParamCfg     DVRGetSet = 2507 //获取人脸参数
	NetDvrSetFaceParamCfg     DVRGetSet = 2508 //设置人脸参数
	NetDvrDelFaceParamCfg     DVRGetSet = 2509 //删除人脸参数
	NetDvrCaptureFaceInfo     DVRGetSet = 2510 //采集人脸信息
	/*************************************人脸识别门禁一体机1.0 end**************************************/
	NetDvrGetRegisterInfo DVRGetSet = 2511 //登记信息获取

	NetDvrGetSmsrelativeparaV50     DVRGetSet = 2512 //获取短信相关参数
	NetDvrSetSmsrelativeparaV50     DVRGetSet = 2513 //设置短信相关参数
	NetDvrGetAcsEvent               DVRGetSet = 2514 //设备事件获取
	NetDvrGetMultiCardCfgV50        DVRGetSet = 2515 //获取多重卡参数V50
	NetDvrSetMultiCardCfgV50        DVRGetSet = 2516 //设置多重卡参数V50
	NetDvrDelFingerprintCfgV50      DVRGetSet = 2517 //删除指纹参数V50
	NetDvrGetEventCardLinkageCfgV51 DVRGetSet = 2518 //获取事件卡号联动配置参数(V51)
	NetDvrSetEventCardLinkageCfgV51 DVRGetSet = 2519 //设置事件卡号联动配置参数(V51)

	NetDvrSetExamInfo                 DVRGetSet = 2530 //考试信息下发
	NetDvrSetExamineeInfo             DVRGetSet = 2531 //考生信息下发
	NetDvrSearchExamCompareResult     DVRGetSet = 2532 //考试比对结果查询
	NetDvrBulkCheckFacePicture        DVRGetSet = 2533 //批量校验人脸图片
	NetDvrJsonConfig                  DVRGetSet = 2550 //JSON透传数据
	NetDvrFaceDataRecord              DVRGetSet = 2551 //添加人脸数据到人脸库
	NetDvrFaceDataSearch              DVRGetSet = 2552 //查询人脸库中的人脸数据
	NetDvrFaceDataModify              DVRGetSet = 2553 //修改人脸库中的人脸数据
	NetDvrCaptureDataSearch           DVRGetSet = 2554 //查询离线采集数据集中数据
	NetDvrGetCard                     DVRGetSet = 2560
	NetDvrSetCard                     DVRGetSet = 2561
	NetDvrDelCard                     DVRGetSet = 2562
	NetDvrGetFingerprint              DVRGetSet = 2563
	NetDvrSetFingerprint              DVRGetSet = 2564
	NetDvrDelFingerprint              DVRGetSet = 2565
	NetDvrGetFace                     DVRGetSet = 2566
	NetDvrSetFace                     DVRGetSet = 2567
	NetDvrGetAllAlarmRs485Cfg         DVRGetSet = 2705 // 获取485参数
	NetDvrGetAllAlarmhostRs485SlotCfg DVRGetSet = 2706 // 获取所有报警主机485槽位参数
	NetDvrGetDeviceSelfCheckState     DVRGetSet = 2707 //获取设备自检功能
	NetDvrGetAllAlarmPointCfg         DVRGetSet = 2708 // 获取所有点号参数
	NetDvrGetAllAlarmSensorCfg        DVRGetSet = 2709 // 获取所有模拟量参数
	NetDvrGetAllAlarmSensorJoint      DVRGetSet = 2710 //获取所有模拟量联动参数
	NetDvrGetAirConditionParam        DVRGetSet = 2711 //获取空调参数
	NetDvrGetOutScaleCfg              DVRGetSet = 2712 //获取主辅口输出配置
	NetDvrSetOutScaleCfg              DVRGetSet = 2713 //设置主辅口输出配置
	NetDvrGetAlarmChanAblitity        DVRGetSet = 2714 //获取报警相关通道参
	/**********************************动环报警主机V3.0****************************************/

	//动环报警主机D2000 V1.0
	NetDvrGetAlarmcenterNetcfg      DVRGetSet = 2715 //获取报警中心网络参数配置
	NetDvrSetAlarmcenterNetcfg      DVRGetSet = 2716 //获取报警中心网络参数配置
	NetItcGetTriggercfg             DVRGetSet = 3003 //获取触发参数
	NetItcSetTriggercfg             DVRGetSet = 3004 //设置触发参数
	NetItcGetIooutParamCfg          DVRGetSet = 3005 //获取IO输出参数（3.1含之后版本）
	NetItcSetIooutParamCfg          DVRGetSet = 3006 //设置IO输出参数（3.1含之后版本）
	NetDvrGetCameraSetupcfg         DVRGetSet = 3007 //获取相机架设参数
	NetDvrSetCameraSetupcfg         DVRGetSet = 3008 //设置相机架设参数
	NetItcGetTriggerDefaultcfg      DVRGetSet = 3013 //获取触发模式推荐参数
	NetDvrGetStatusDetectcfg        DVRGetSet = 3015 //获取状态检测使能参数
	NetDvrSetStatusDetectcfg        DVRGetSet = 3016 //设置状态检测使能参数
	NetItcGetVideoTriggercfg        DVRGetSet = 3017 //获取视频电警触发参数
	NetItcSetVideoTriggercfg        DVRGetSet = 3018 //设置视频电警触发参数
	NetDvrGetTpsAlarmcfg            DVRGetSet = 3019 //获取交通统计报警参数
	NetDvrSetTpsAlarmcfg            DVRGetSet = 3020 //设置交通统计报警参数
	NetDvrGetRedareacfg             DVRGetSet = 3100 //获取红绿灯区域参数
	NetDvrSetRedareacfg             DVRGetSet = 3101 //设置红绿灯区域参数
	NetDvrGetTestSpot               DVRGetSet = 3102 //获取SPOT口测试总步数和当前第几步
	NetDvrSetTestSpot               DVRGetSet = 3103 //设置SPOT口测试总步数和当前第几步
	NetDvrGetCabinetcfg             DVRGetSet = 3104 //机柜参数配置获取
	NetDvrSetCabinetcfg             DVRGetSet = 3105 //机柜参数配置设置
	NetDvrVehicleCheckStart         DVRGetSet = 3106 //黑名单稽查数据回传
	NetDvrSetCapturepicCfg          DVRGetSet = 3107 //车载抓图配置设置命令
	NetDvrGetCapturepicCfg          DVRGetSet = 3108 //车载抓图配置获取命令
	NetDvrSetMobileplateRecogCfg    DVRGetSet = 3109 //车载车牌识别配置设置命令
	NetDvrGetMobileplateRecogCfg    DVRGetSet = 3110 //车载车牌识别配置获取命令
	NetDvrSetMobileRadarCfg         DVRGetSet = 3111 //车载雷达配置设置命令
	NetDvrGetMobileRadarCfg         DVRGetSet = 3112 //车载雷达配置获取命令
	NetDvrSetMobileLocalplatechkCfg DVRGetSet = 3113 //车载黑名单本地对比配置设置命令
	NetDvrGetMobileLocalplatechkCfg DVRGetSet = 3114 //车载黑名单本地对比配置获取命令
	NetItcGetIcrcfg                 DVRGetSet = 3115 //获取ICR配置切换
	NetItcSetIcrcfg                 DVRGetSet = 3116 //设置ICR配置切换
	NetItcGetRs485Accessinfo        DVRGetSet = 3117 //获取Rs485关联接入设备的信息
	NetItcSetRs485Accessinfo        DVRGetSet = 3118 //设置Rs485关联接入设备的信息
	NetItcGetExceptioncfg           DVRGetSet = 3119 //获取异常参数
	NetItcSetExceptioncfg           DVRGetSet = 3120 //设置异常参数
	NetItcGetFtpcfg                 DVRGetSet = 3121 //获取ITC  FTP设置参数
	NetItcSetFtpcfg                 DVRGetSet = 3122 //设置ITC FTP设置参数
	NetDvrVehicleControlListStart   DVRGetSet = 3123 //设置车辆黑白名单信息
	NetDvrGetAllVehicleControlList  DVRGetSet = 3124 //获取所有车辆黑白名单信息
	NetDvrVehicleDelinfoCtrl        DVRGetSet = 3125 //删除设备内黑名单数据库信息
	NetDvrGetEntranceParamcfg       DVRGetSet = 3126 //获取出入口控制参数
	NetDvrSetEntranceParamcfg       DVRGetSet = 3127 //设置出入口控制参数
	NetDvrBarriergateCtrl           DVRGetSet = 3128 //远程控制道闸
	NetDvrGatelampCtrl              DVRGetSet = 3129 //常亮灯功能
	NetDvrGetCurtriggermode         DVRGetSet = 3130 //获取设备当前触发模式
	NetDvrGetGpsdatacfg             DVRGetSet = 3131 //获取GPS参数
	NetDvrSetGpsdatacfg             DVRGetSet = 3132 //设置GPS参数
	NetDvrVehiclelistCtrlStart      DVRGetSet = 3133 //设置车辆黑白名单信息

	NetDvrGetGuardcfg        DVRGetSet = 3134 //获取车牌识别检测计划
	NetDvrSetGuardcfg        DVRGetSet = 3135 //设置车牌识别检测计划
	NetDvrGetSnapinfoCfg     DVRGetSet = 3136 //获取抓拍图片参数
	NetDvrSetSnapinfoCfg     DVRGetSet = 3137 //设置抓拍图片参数
	NetDvrGetSnapinfoCfgV40  DVRGetSet = 3138 //获取抓拍图片参数扩展
	NetDvrSetSnapinfoCfgV40  DVRGetSet = 3139 //设置抓拍图片参数扩展
	NetDvrSetCurtriggermode  DVRGetSet = 3140 //设置设备当前触发模式(仅IPC/D支持)
	NetDvrGetTrafficData     DVRGetSet = 3141 //长连接获取交通数据
	NetDvrGetTrafficFlow     DVRGetSet = 3142 //长连接获取交通流量
	NetDvrParkingVehicleSend DVRGetSet = 3143 //停车车辆信息下发
	NetDvrParkingCardSend    DVRGetSet = 3144 //停车卡下发
	NetDvrParkingCardCtrl    DVRGetSet = 3145 //停车场停车卡控制接口

	NetDvrGetAlarmctrlCapabilities DVRGetSet = 3146 //获取报警控制能力
	NetDvrSetAlarmctrlCfg          DVRGetSet = 3147 //设置报警控制
	NetDvrGetAlarmctrlCfg          DVRGetSet = 3148 //获取报警控制

	NetDvrGetAudioInput            DVRGetSet = 3201 //获取音频输入参数
	NetDvrSetAudioInput            DVRGetSet = 3202 //设置音频输入参数
	NetDvrGetCameraDehazeCfg       DVRGetSet = 3203 //获取透雾参数配置
	NetDvrSetCameraDehazeCfg       DVRGetSet = 3204 //设置透雾参数配置
	NetDvrRemotecontrolAlarm       DVRGetSet = 3205 //远程控制遥控器布防
	NetDvrRemotecontrolDisalarm    DVRGetSet = 3206 //远程控制遥控器撤防
	NetDvrRemotecontrolStudy       DVRGetSet = 3207 //远程控制遥控器学习
	NetDvrWirelessAlarmStudy       DVRGetSet = 3208 //远程控制无线报警学习
	NetIpcGetAuxAlarmcfg           DVRGetSet = 3209 //获取辅助报警参数配置
	NetIpcSetAuxAlarmcfg           DVRGetSet = 3210 //设置辅助报警参数配置
	NetDvrGetPreviewDisplaycfg     DVRGetSet = 3211 //获取预览显示参数
	NetDvrSetPreviewDisplaycfg     DVRGetSet = 3212 //设置预览显示参数
	NetDvrRemotecontrolPtz         DVRGetSet = 3213 //远程控制PTZ
	NetDvrRemotecontrolPresetpoint DVRGetSet = 3214 //远程控制预置点
	NetDvrRemotecontrolCruise      DVRGetSet = 3215 //远程控制巡航

	NetDvrGetMultiStreamCompressioncfg DVRGetSet = 3216 //远程获取多码流压缩参数
	NetDvrSetMultiStreamCompressioncfg DVRGetSet = 3217 //远程设置多码流压缩参数

	NetDvrGetWpscfg              DVRGetSet = 3218 //获取WPS参数
	NetDvrSetWpscfg              DVRGetSet = 3219 //设置WPS参数
	NetDvrWpsConnect             DVRGetSet = 3220 //远程启用WPS连接
	NetDvrGetDevicePin           DVRGetSet = 3221 //获取设备PIN码
	NetDvrUpdatePin              DVRGetSet = 3223 //更新设备PIN码
	NetDvrGetPresetcfg           DVRGetSet = 3224 //获取预置点参数
	NetDvrGetPtzcruisecfg        DVRGetSet = 3225 //获取巡航路径参数
	NetDvrGetPresetNum           DVRGetSet = 3226 //获取预置点个数
	NetDvrGetPtzcruiseNum        DVRGetSet = 3227 //获取巡航路径个数
	NetDvrGetMotionTrackCfg      DVRGetSet = 3228 //获取跟踪参数
	NetDvrSetMotionTrackCfg      DVRGetSet = 3229 //设置跟踪参数
	NetDvrClearIpcParam          DVRGetSet = 3230 //清空前端参数
	NetDvrGetIpaddrFiltercfg     DVRGetSet = 3232 //获取IP地址过滤参数
	NetDvrSetIpaddrFiltercfg     DVRGetSet = 3233 //设置IP地址过滤参数
	NetDvrGetLogoOverlaycfg      DVRGetSet = 3234 //获取LOGO图片叠加参数
	NetDvrSetLogoOverlaycfg      DVRGetSet = 3235 //设置LOGO图片叠加参数
	NetDvrGetIpv6List            DVRGetSet = 3236 //获取网卡的全部IPV6地址信息
	NetDvrGetAudiooutVolume      DVRGetSet = 3237 //获取输出音频大小
	NetDvrSetAudiooutVolume      DVRGetSet = 3238 //设置输出音频大小
	NetDvrGetFuzzyUpgrade        DVRGetSet = 3239 //获取模糊升级匹配信息
	NetDvrGetBvCorrectParam      DVRGetSet = 3240 //获取相机校正参数
	NetDvrSetBvCorrectParam      DVRGetSet = 3241 //设置相机校正参数
	NetDvrGetOutputVideoType     DVRGetSet = 3242 //获取输出视频类型
	NetDvrSetOutputVideoType     DVRGetSet = 3243 //设置输出视频类型
	NetDvrFisheyeCfg             DVRGetSet = 3244 //鱼眼长连接配置
	NetDvrGetPtzPoint            DVRGetSet = 3245 //获取PTZ点坐标
	NetDvrSetPtzPoint            DVRGetSet = 3246 //设置PTZ点坐标
	NetDvrRemotecontrolDevParam  DVRGetSet = 3247 //设置设备登录客户端参数
	NetDvrGetFisheyeStreamStatus DVRGetSet = 3248 //获取鱼眼码流状态
	NetDvrGetGbt28181AccessCfg   DVRGetSet = 3249 //获取GBT28181协议接入配置
	NetDvrSetGbt28181AccessCfg   DVRGetSet = 3250 //设置GBT28181协议接入配置
	NetDvrGetGbt28181ChaninfoCfg DVRGetSet = 3251 //获取GBT28181协议接入设备的通道信息
	NetDvrSetGbt28181ChaninfoCfg DVRGetSet = 3252 //设置GBT28181协议接入设备的通道信息
	NetDvrGetGbt28181Alarmincfg  DVRGetSet = 3253 //获取GBT28181协议接入设备的报警信息
	NetDvrSetGbt28181Alarmincfg  DVRGetSet = 3254 //设置GBT28181协议接入设备的报警信息
	NetDvrGetIspCameraparamcfg   DVRGetSet = 3255 //获取ISP前端参数配置
	NetDvrSetIspCameraparamcfg   DVRGetSet = 3256 //设置ISP前端参数配置
	NetDvrGetDevserverCfg        DVRGetSet = 3257 //获取模块服务配置
	NetDvrSetDevserverCfg        DVRGetSet = 3258 //设置模块服务配置

	//2013-11-25
	NetDvrGetWiperinfoCfg               DVRGetSet = 3259 //雨刷配置获取
	NetDvrSetWiperinfoCfg               DVRGetSet = 3260 //雨刷配置设置
	NetDvrGetTrackDevParam              DVRGetSet = 3261 //获取跟踪设备参数
	NetDvrSetTrackDevParam              DVRGetSet = 3262 //设置跟踪设备参数
	NetDvrGetPtzTrackParam              DVRGetSet = 3263 //获取PTZ跟踪参数
	NetDvrSetPtzTrackParam              DVRGetSet = 3264 //设置PTZ跟踪参数
	NetDvrGetCenterPointCfg             DVRGetSet = 3265 //获取中心点参数
	NetDvrSetCenterPointCfg             DVRGetSet = 3266 //设置中心点参数
	NetDvrGetCenterPointCfgCapabilities DVRGetSet = 3267 //获取中心点参数能力
	NetDvrGetFisheyeCapabilities        DVRGetSet = 3268 //获取鱼眼能力

	NetDvrGetBasicparamcfg           DVRGetSet = 3270 //获取PTZ配置基本参数信息
	NetDvrSetBasicparamcfg           DVRGetSet = 3271 //设置PTZ配置基本参数信息
	NetDvrGetPtzosdcfg               DVRGetSet = 3272 //获取PTZ OSD配置参数信息
	NetDvrSetPtzosdcfg               DVRGetSet = 3273 //设置PTZ OSD配置参数信息
	NetDvrGetPoweroffmemcfg          DVRGetSet = 3274 //获取掉电记忆模式参数信息
	NetDvrSetPoweroffmemcfg          DVRGetSet = 3275 //设置掉电记忆模式参数信息
	NetDvrGetLimitcfg                DVRGetSet = 3276 //获取限位参数配置信息
	NetDvrSetLimitcfg                DVRGetSet = 3277 //设置限位参数配置信息
	NetDvrPtzlimitCtrl               DVRGetSet = 3278 //清除限位参数控制
	NetDvrPtzClearctrl               DVRGetSet = 3279 //清除配置信息控制接口
	NetDvrGetPrioritizecfg           DVRGetSet = 3281 //获取云台优先配置信息
	NetDvrSetPrioritizecfg           DVRGetSet = 3282 //设置云台优先配置信息
	NetDvrPtzInitialpositionctrl     DVRGetSet = 3283 //零方位角控制
	NetDvrGetPrivacyMaskscfg         DVRGetSet = 3285 //获取隐私遮蔽参数
	NetDvrSetPrivacyMaskscfg         DVRGetSet = 3286 //设置隐私遮蔽参数
	NetDvrGetPtzlockcfg              DVRGetSet = 3287 //获取云台锁定信息
	NetDvrSetPtzlockcfg              DVRGetSet = 3288 //设置云台锁定信息
	NetDvrPtzZoomratioctrl           DVRGetSet = 3289 //设置跟踪倍率
	NetDvrGetPtzlockinfo             DVRGetSet = 3290 //获取云台锁定剩余秒数
	NetDvrGetPrivacyMasksEnablecfg   DVRGetSet = 3291 //获取全局使能
	NetDvrSetPrivacyMasksEnablecfg   DVRGetSet = 3292 //设置全局使能
	NetDvrGetSmarttrackcfg           DVRGetSet = 3293 //获取智能运动跟踪配置信息
	NetDvrSetSmarttrackcfg           DVRGetSet = 3294 //设置智能运动跟踪配置信息
	NetDvrGetEptzCfg                 DVRGetSet = 3295 //获取EPTZ参数
	NetDvrSetEptzCfg                 DVRGetSet = 3296 //设置EPTZ参数
	NetDvrGetEptzCfgCapabilities     DVRGetSet = 3297 //获取EPTZ参数能力
	NetDvrGetLowLightcfg             DVRGetSet = 3303 //获取快球低照度设置信息
	NetDvrSetLowLightcfg             DVRGetSet = 3304 //设置快球低照度设置信息
	NetDvrGetFocusmodecfg            DVRGetSet = 3305 //获取快球聚焦模式信息
	NetDvrSetFocusmodecfg            DVRGetSet = 3306 //设置快球聚焦模式信息
	NetDvrGetInfrarecfg              DVRGetSet = 3307 //获取快球红外设置信息
	NetDvrSetInfrarecfg              DVRGetSet = 3308 //设置快球红外设置信息
	NetDvrGetAemodecfg               DVRGetSet = 3309 //获取快球其他设置信息
	NetDvrSetAemodecfg               DVRGetSet = 3310 //设置快球其他设置信息
	NetDvrControlRestoreSupport      DVRGetSet = 3311 //恢复前端默认参数(参数能力中有的前端参数配置相关的都恢复)
	NetDvrControlRestartSupport      DVRGetSet = 3312 //快球机芯重启
	NetDvrControlPtzPattern          DVRGetSet = 3313 //云台花样扫描
	NetDvrGetPtzParkactionCfg        DVRGetSet = 3314 //获取云台守望参数
	NetDvrSetPtzParkactionCfg        DVRGetSet = 3315 //设置云台守望参数
	NetDvrControlPtzManualtrace      DVRGetSet = 3316 //手动跟踪定位
	NetDvrGetRoiDetectNum            DVRGetSet = 3349 //获取ROI检测区域编号数目
	NetDvrGetRoiDetect               DVRGetSet = 3350 //获取ROI检测区域配置
	NetDvrSetRoiDetect               DVRGetSet = 3351 //设置ROI检测区域配置
	NetDvrGetFaceDetect              DVRGetSet = 3352 //获取人脸侦测配置
	NetDvrSetFaceDetect              DVRGetSet = 3353 //设置人脸侦测配置
	NetDvrGetCorridorMode            DVRGetSet = 3354 //获取走廊模式功能配置
	NetDvrSetCorridorMode            DVRGetSet = 3355 //设置走廊模式功能配置
	NetDvrGetScenechangeDetectioncfg DVRGetSet = 3356 //获取场景变更报警配置
	NetDvrSetScenechangeDetectioncfg DVRGetSet = 3357 //设置场景变更报警配置
	NetDvrGetTraversePlaneDetection  DVRGetSet = 3360
	NetDvrSetTraversePlaneDetection  DVRGetSet = 3361
	NetDvrGetFieldDetection          DVRGetSet = 3362 //获取区域侦测配置
	NetDvrSetFieldDetection          DVRGetSet = 3363 //设置区域侦测配置
	NetDvrGetDefocusparam            DVRGetSet = 3364 //获取虚焦侦测参数配置
	NetDvrSetDefocusparam            DVRGetSet = 3365 //设置虚焦侦测参数配置
	NetDvrGetAudioexceptionparam     DVRGetSet = 3366 //获取音频异常配置
	NetDvrSetAudioexceptionparam     DVRGetSet = 3367 //设置音频异常配置
	NetDvrGetCcdparamcfgEx           DVRGetSet = 3368 //获取CCD参数配置
	NetDvrSetCcdparamcfgEx           DVRGetSet = 3369 //设置CCD参数配置
	NetDvrStartGetInputvolume        DVRGetSet = 3370 //开始获取音量
	NetDvrSetSchTask                 DVRGetSet = 3380 //设置球机定时任务
	NetDvrGetSchTask                 DVRGetSet = 3381 //获取球机定时任务
	NetDvrSetPresetName              DVRGetSet = 3382 //设置预置点名称
	NetDvrGetPresetName              DVRGetSet = 3383 //获取预置点名称
	NetDvrSetAudioName               DVRGetSet = 3384 //设置语音名称
	NetDvrGetAudioName               DVRGetSet = 3385 //获取语音名称
	NetDvrResumeInitrackpos          DVRGetSet = 3386 //恢复跟踪初始位
	NetDvrNtpServerTest              DVRGetSet = 3387 //NTP服务器测试
	NetDvrNasServerTest              DVRGetSet = 3388 //NAS服务器测试
	NetDvrEmailServerTest            DVRGetSet = 3389 //Email服务器测试
	NetDvrFtpServerTest              DVRGetSet = 3390 //FTP服务器测试
	NetDvrIpTest                     DVRGetSet = 3391 //IP测试
	NetDvrGetNetDiskcfgV40           DVRGetSet = 3392 //网络硬盘接入获取v40
	NetDvrSetNetDiskcfgV40           DVRGetSet = 3393 //网络硬盘接入设置v40
	NetDvrGetIooutCfg                DVRGetSet = 3394 //获取补光灯参数
	NetDvrSetIooutCfg                DVRGetSet = 3395 //设置补光灯参数
	NetDvrGetSignalSync              DVRGetSet = 3396 //获取信号灯同步配置参数
	NetDvrSetSignalSync              DVRGetSet = 3397 //设置信号灯同步配置参数
	NetDvrGetEzvizAccessCfg          DVRGetSet = 3398 //获取EZVIZ接入参数
	NetDvrSetEzvizAccessCfg          DVRGetSet = 3399 //设置EZVIZ接入参数
	NetDvrGetScheduleAutoTrackCfg    DVRGetSet = 3400 //获取定时智能跟踪参数
	NetDvrSetScheduleAutoTrackCfg    DVRGetSet = 3401 //设置定时智能跟踪参数
	NetDvrMakeIFrame                 DVRGetSet = 3402 //强制I帧
	NetDvrGetAlarmRelate             DVRGetSet = 3403 //获取报警联动通道功能参数
	NetDvrSetAlarmRelate             DVRGetSet = 3404 //设置报警联动通道功能参数
	NetDvrGetPdcRulecfgV42           DVRGetSet = 3405 //设置人流量统计规则(扩展)
	NetDvrSetPdcRulecfgV42           DVRGetSet = 3406 //获取人流量统计规则(扩展)
	NetDvrGetHeatmapCfg              DVRGetSet = 3407 //设置热度图参数配置
	NetDvrSetHeatmapCfg              DVRGetSet = 3408 //获取热度图参数配置
	NetDvrRemotecontrolLinearscan    DVRGetSet = 3409 //设置左右边界参数 2014-03-15
	NetDvrDpcCtrl                    DVRGetSet = 3410 //坏点校正控制
	NetDvrFfcManualCtrl              DVRGetSet = 3411 //非均匀性校正(FFC)手动模式
	NetDvrFfcBackcompCtrl            DVRGetSet = 3412 //非均匀性校正(FFC)背景补偿
	NetDvrGetFocusingPositionState   DVRGetSet = 3413 //获取聚焦到位状态参数
	NetDvrGetPrivateProtocolCfg      DVRGetSet = 3414 //获取 私有关键信息上传配置接口配置
	NetDvrSetPrivateProtocolCfg      DVRGetSet = 3415 //设置 私有关键信息上传配置接口配置
	NetDvrCompleteRestoreCtrl        DVRGetSet = 3420 //设置完全恢复出厂值
	NetDvrCloudstorageServerTest     DVRGetSet = 3421 //云存储服务器测试
	NetDvrPhoneNumTest               DVRGetSet = 3422 //电话号码测试
	NetDvrGetRemotecontrolStatus     DVRGetSet = 3423 //获取无线布防状态
	NetDvrGetMonitorLocationInfo     DVRGetSet = 3424 //获取监测点信息
	NetDvrSetMonitorLocationInfo     DVRGetSet = 3425 //设置监测点信息

	NetDvrGetSmartCapabilities                     DVRGetSet = 3500 //获取Smart能力
	NetDvrGetEventTriggersCapabilities             DVRGetSet = 3501 //获取事件触发能力
	NetDvrGetRegionEntranceCapabilities            DVRGetSet = 3502 //获取进入区域侦测能力
	NetDvrGetRegionEntrDetection                   DVRGetSet = 3503 //获取进入区域配置
	NetDvrSetRegionEntrDetection                   DVRGetSet = 3504 //设置进入区域配置
	NetDvrGetRegionEntrRegion                      DVRGetSet = 3505 //获取进入区域的单个区域配置
	NetDvrSetRegionEntrRegion                      DVRGetSet = 3506 //设置进入区域的单个区域配置
	NetDvrGetRegionEntrTrigger                     DVRGetSet = 3507 //获取进入区域联动配置
	NetDvrSetRegionEntrTrigger                     DVRGetSet = 3508 //设置进入区域联动配置
	NetDvrGetRegionEntrSchedule                    DVRGetSet = 3509 //获取进入区域布防时间配置
	NetDvrSetRegionEntrSchedule                    DVRGetSet = 3510 //设置进入区域布防时间配置
	NetDvrGetRegionExitintCapabilities             DVRGetSet = 3511 //获取离开区域侦测能力
	NetDvrGetRegionExitingDetection                DVRGetSet = 3512 //获取离开区域配置
	NetDvrSetRegionExitingDetection                DVRGetSet = 3513 //设置离开区域配置
	NetDvrGetRegionExitingRegion                   DVRGetSet = 3514 //获取离开区域的单个区域配置
	NetDvrSetRegionExitingRegion                   DVRGetSet = 3515 //设置离开区域的单个区域配置
	NetDvrGetRegionExitTrigger                     DVRGetSet = 3516 //获取离开区域联动配置
	NetDvrSetRegionExitTrigger                     DVRGetSet = 3517 //设置离开区域联动配置
	NetDvrGetRegionExitSchedule                    DVRGetSet = 3518 //获取离开区域布防时间配置
	NetDvrSetRegionExitSchedule                    DVRGetSet = 3519 //设置离开区域布防时间配置
	NetDvrGetLoiteringCapabilities                 DVRGetSet = 3520 //获取徘徊侦测能力
	NetDvrGetLoiteringDetection                    DVRGetSet = 3521 //获取徘徊侦测配置
	NetDvrSetLoiteringDetection                    DVRGetSet = 3522 //设置徘徊侦测配置
	NetDvrGetLoiteringRegion                       DVRGetSet = 3523 //获取徘徊的单个区域配置
	NetDvrSetLoiteringRegion                       DVRGetSet = 3524 //设置徘徊的单个区域配置
	NetDvrGetLoiteringTrigger                      DVRGetSet = 3525 //获取徘徊联动配置
	NetDvrSetLoiteringTrigger                      DVRGetSet = 3526 //设置徘徊联动配置
	NetDvrGetLoiteringSchedule                     DVRGetSet = 3527 //获取徘徊布防时间配置
	NetDvrSetLoiteringSchedule                     DVRGetSet = 3528 //设置徘徊布防时间配置
	NetDvrGetGroupdetectionCapabilities            DVRGetSet = 3529 //获取人员聚集侦测能力
	NetDvrGetGroupDetection                        DVRGetSet = 3530 //获取人员聚集侦测配置
	NetDvrSetGroupDetection                        DVRGetSet = 3531 //设置人员聚集侦测配置
	NetDvrGetGroupdetectionRegion                  DVRGetSet = 3532 //获取人员聚集的单个区域配置
	NetDvrSetGroupdetectionRegion                  DVRGetSet = 3533 //设置人员聚集的单个区域配置
	NetDvrGetGroupdetectionTrigger                 DVRGetSet = 3534 //获取人员聚集联动配置
	NetDvrSetGroupdetectionTrigger                 DVRGetSet = 3535 //设置人员聚集联动配置
	NetDvrGetGroupdetectionSchedule                DVRGetSet = 3536 //获取人员聚集布防时间配置
	NetDvrSetGroupdetectionSchedule                DVRGetSet = 3537 //设置人员聚集布防时间配置
	NetDvrGetRapidmoveCapabilities                 DVRGetSet = 3538 //获取快速运动侦测能力
	NetDvrGetRapidmoveDetection                    DVRGetSet = 3539 //获取快速运动侦测配置
	NetDvrSetRapidmoveDetection                    DVRGetSet = 3540 //设置快速运动侦测配置
	NetDvrGetRapidmoveRegion                       DVRGetSet = 3541 //获取快速运动的单个区域配置
	NetDvrSetRapidmoveRegion                       DVRGetSet = 3542 //设置快速运动的单个区域配置
	NetDvrGetRapidmoveTrigger                      DVRGetSet = 3543 //获取快速运动联动配置
	NetDvrSetRapidmoveTrigger                      DVRGetSet = 3544 //设置快速运动联动配置
	NetDvrGetRapidmoveSchedule                     DVRGetSet = 3545 //获取快速运动的布防时间配置
	NetDvrSetRapidmoveSchedule                     DVRGetSet = 3546 //设置快速运动的布防时间配置
	NetDvrGetPatkingCapabilities                   DVRGetSet = 3547 //获取停车侦测能力
	NetDvrGetParkingDetection                      DVRGetSet = 3548 //获取停车侦测配置
	NetDvrSetParkingDetection                      DVRGetSet = 3549 //设置停车侦测配置
	NetDvrGetParkingRegion                         DVRGetSet = 3550 //获取停车侦测的单个区域配置
	NetDvrSetParkingRegion                         DVRGetSet = 3551 //设置停车侦测的单个区域配置
	NetDvrGetParkingTrigger                        DVRGetSet = 3552 //获取停车侦测联动配置
	NetDvrSetParkingTrigger                        DVRGetSet = 3553 //设置停车侦测联动配置
	NetDvrGetParkingSchedule                       DVRGetSet = 3554 //获取停车侦测的布防时间配置
	NetDvrSetParkingSchedule                       DVRGetSet = 3555 //设置停车侦测的布防时间配置
	NetDvrGetUnattendedBaggageCapabilities         DVRGetSet = 3556 //获取物品遗留侦测能力
	NetDvrGetUnattendedBaggageDetection            DVRGetSet = 3557 //获取物品遗留侦测配置
	NetDvrSetUnattendedBaggageDetection            DVRGetSet = 3558 //设置物品遗留侦测配置
	NetDvrGetUnattendedBaggageRegion               DVRGetSet = 3559 //获取物品遗留侦测的单个区域配置
	NetDvrSetUnattendedBaggageRegion               DVRGetSet = 3560 //设置物品遗留侦测的单个区域配置
	NetDvrGetUnattendedBaggageTrigger              DVRGetSet = 3561 //获取物品遗留侦测联动配置
	NetDvrSetUnattendedBaggageTrigger              DVRGetSet = 3562 //设置物品遗留侦测联动配置
	NetDvrGetUnattendedBaggageSchedule             DVRGetSet = 3563 //获取物品遗留侦测的布防时间配置
	NetDvrSetUnattendedBaggageSchedule             DVRGetSet = 3564 //设置物品遗留侦测的布防时间配置
	NetDvrGetAttendedbaggageCapabilities           DVRGetSet = 3565 //获取物品拿取侦测能力
	NetDvrGetAttendedbaggageDetection              DVRGetSet = 3566 //获取物品拿取侦测配置
	NetDvrSetAttendedbaggageDetection              DVRGetSet = 3567 //设置物品拿取侦测配置
	NetDvrGetAttendedbaggageRegion                 DVRGetSet = 3568 //获取物品拿取侦测的单个区域配置
	NetDvrSetAttendedbaggageRegion                 DVRGetSet = 3569 //设置物品拿取侦测的单个区域配置
	NetDvrGetAttendedbaggageTrigger                DVRGetSet = 3570 //获取物品拿取侦测联动配置
	NetDvrSetAttendedbaggageTrigger                DVRGetSet = 3571 //设置物品拿取侦测联动配置
	NetDvrGetAttendedbaggageSchedule               DVRGetSet = 3572 //获取物品遗留侦测的布防时间配置
	NetDvrSetAttendedbaggageSchedule               DVRGetSet = 3573 //设置物品拿取侦测的布防时间配置
	NetDvrGetRegionclipCapabilities                DVRGetSet = 3574 //获取区域裁剪能力
	NetDvrGetRegionClip                            DVRGetSet = 3575 //获取区域裁剪配置
	NetDvrSetRegionClip                            DVRGetSet = 3576 //设置区域裁剪配置
	NetDvrGetNetworkCapabilities                   DVRGetSet = 3577 //获取网络能力
	NetDvrGetWirelessDial                          DVRGetSet = 3578 //获取无线参数配置
	NetDvrSetWirelessDial                          DVRGetSet = 3579 //设置无线参数配置
	NetDvrGetWirelessdialCapabilities              DVRGetSet = 3580 //获取无线拨号参数能力
	NetDvrGetWirelessdialSchedule                  DVRGetSet = 3581 //获取拨号计划配置
	NetDvrSetWirelessdialSchedule                  DVRGetSet = 3582 //设置拨号计划配置
	NetDvrGetWirelessdialStatus                    DVRGetSet = 3583 //获取拨号状态
	NetDvrGetRegionEntranceScheduleCapabilities    DVRGetSet = 3584 //获取进入区域侦测布防时间能力
	NetDvrGetRegionExitingScheduleCapabilities     DVRGetSet = 3585 //获取离开区域侦测布防时间能力
	NetDvrGetLoiteringScheduleCapabilities         DVRGetSet = 3586 //获取徘徊侦测布防时间能力
	NetDvrGetGroupScheduleCapabilities             DVRGetSet = 3587 //获取人员聚集侦测布防时间能力
	NetDvrGetRapidmoveScheduleCapabilities         DVRGetSet = 3588 //获取快速运动侦测布防时间能力
	NetDvrGetParkingScheduleCapabilities           DVRGetSet = 3589 //获取停车侦测布防时间能力
	NetDvrGetUnattendedbaggageScheduleCapabilities DVRGetSet = 3590 //获取物品遗留侦测布防时间能力
	NetDvrGetAttendedbaggageScheduleCapabilities   DVRGetSet = 3591 //获取物品拿取侦测布防时间能力
	NetDvrGetWirelessdialScheduleCapabilities      DVRGetSet = 3592 //获取拨号计划能力
	NetDvrWirelessdialConnect                      DVRGetSet = 3593 //控制无线网络连网断网

	NetDvrGetLitestorage                     DVRGetSet = 3594 //获取轻存储配置
	NetDvrSetLitestorage                     DVRGetSet = 3595 //设置轻存储配置
	NetDvrGetLitestorageCapabilities         DVRGetSet = 3596 //获取轻存储能力
	NetDvrGetVehicleCapabilities             DVRGetSet = 3597 //获取车俩检测标定能力
	NetDvrGetVehicleCalibration              DVRGetSet = 3598 //获取车辆检测标定
	NetDvrGetSlavecameraCapabilities         DVRGetSet = 3599 //获取从摄像机IP信息配置能力
	NetDvrGetSlavecamera                     DVRGetSet = 3600 //获取从摄像机IP信息配置
	NetDvrSetSlavecamera                     DVRGetSet = 3601 //设置从摄像机IP信息配置
	NetDvrGetSlavecameraStatus               DVRGetSet = 3602 //获取从摄像机连接状态
	NetDvrGetSlavecameraCalibCapabilities    DVRGetSet = 3603 //获取从摄像机配置&&标定能力
	NetDvrGetSlavecameraCalib                DVRGetSet = 3604 //获取从摄像机标定配置
	NetDvrSetSlavecameraCalib                DVRGetSet = 3605 //设置从摄像机标定配置
	NetDvrGetPhyRatio                        DVRGetSet = 3606 //获取物理倍率坐标信息
	NetDvrSetPhyRatio                        DVRGetSet = 3607 //设置物理倍率坐标信息
	NetDvrGetMasterslavetrackingCapabilities DVRGetSet = 3608 //获取主从跟踪能力
	NetDvrSetTrackingratio                   DVRGetSet = 3610 //设置从摄像机跟踪倍率
	NetDvrGetTracking                        DVRGetSet = 3611 //获取主从跟踪功能相机跟踪配置
	NetDvrSetTracking                        DVRGetSet = 3612 //设置主从跟踪功能相机跟踪配置
	NetDvrGetTrackingCapabilities            DVRGetSet = 3613 //获取主从跟踪功能相机跟踪配置能力
	NetDvrGetSlavecameraCalibV50             DVRGetSet = 3614 //获取从摄像机标定配置V50
	NetDvrSetSlavecameraCalibV50             DVRGetSet = 3615 //设置从摄像机标定配置V50
	NetDvrSetTrackingratioManual             DVRGetSet = 3616 //设置从摄像机手动跟踪倍率
	NetDvrGetTrackingratioManual             DVRGetSet = 3617 //获取从摄像机手动跟踪倍率
	NetDvrSetTrackInitpostion                DVRGetSet = 3618 //设置从摄像机初始跟踪位置
	NetDvrGetPtzCapabilities                 DVRGetSet = 3619 //获取ptz球机控制能力

	NetDvrGetThermometryBasicparamCapabilities DVRGetSet = 3620 //获取测温配置能力
	NetDvrGetThermometryBasicparam             DVRGetSet = 3621 //获取测温配置参数
	NetDvrSetThermometryBasicparam             DVRGetSet = 3622 //设置测温配置参数
	NetDvrGetThermometrySceneCapabilities      DVRGetSet = 3623 //获取测温预置点关联配置能力
	NetDvrGetThermometryPresetinfo             DVRGetSet = 3624 //获取测温预置点关联配置参数
	NetDvrSetThermometryPresetinfo             DVRGetSet = 3625 //设置测温预置点关联配置参数
	NetDvrGetThermometryAlarmruleCapabilities  DVRGetSet = 3626 //获取测温报警方式配置能力
	NetDvrGetThermometryAlarmrule              DVRGetSet = 3627 //获取测温预置点报警规则配置参数
	NetDvrSetThermometryAlarmrule              DVRGetSet = 3628 //设置测温预置点报警规则配置参数
	NetDvrGetRealtimeThermometry               DVRGetSet = 3629 //实时温度检测
	NetDvrGetThermometryDiffcomparison         DVRGetSet = 3630 //获取测温预置点温差规则配置参数
	NetDvrSetThermometryDiffcomparison         DVRGetSet = 3631 //设置测温预置点温差规则配置参数
	NetDvrGetThermometryTrigger                DVRGetSet = 3632 //获取测温联动配置
	NetDvrSetThermometryTrigger                DVRGetSet = 3633 //设置测温联动配置
	NetDvrGetThermalCapabilities               DVRGetSet = 3634 //获取热成像（Thermal）能力
	NetDvrGetFiredetectionCapabilities         DVRGetSet = 3635 //获取火点检测配置能力
	NetDvrGetFiredetection                     DVRGetSet = 3636 //获取火点检测参数
	NetDvrSetFiredetection                     DVRGetSet = 3637 //设置火点检测参数
	NetDvrGetFiredetectionTrigger              DVRGetSet = 3638 //获取火点检测联动配置
	NetDvrSetFiredetectionTrigger              DVRGetSet = 3639 //设置火点检测联动配置
	NetDvrGetOisCapabilities                   DVRGetSet = 3640 //获取光学防抖参数配置能力
	NetDvrGetOisCfg                            DVRGetSet = 3641 //获取光学防抖配置
	NetDvrSetOisCfg                            DVRGetSet = 3642 //设置光学防抖配置
	NetDvrGetMacfilterCapabilities             DVRGetSet = 3643 //获取MAC地址过滤配置能力
	NetDvrGetMacfilterCfg                      DVRGetSet = 3644 //获取MAC地址过滤配置
	NetDvrSetMacfilterCfg                      DVRGetSet = 3645 //设置MAC地址过滤配置
	NetDvrGetEaglefocusCalcfgCapabilities      DVRGetSet = 3646 //鹰视聚焦标定配置能力
	NetDvrGetEaglefocusingCalcfg               DVRGetSet = 3647 //获取鹰视聚焦标定配置
	NetDvrSetEaglefocusingCalcfg               DVRGetSet = 3648 //设置鹰视聚焦标定配置
	NetDvrGetEaglefocusingCfgCapabilities      DVRGetSet = 3649 //获取鹰视聚焦配置能力
	NetDvrGetEaglefocusingCtrl                 DVRGetSet = 3650 //获取鹰视聚焦配置
	NetDvrSetEaglefocusingCtrl                 DVRGetSet = 3651 //设置鹰视聚焦配置
	NetDvrGetPxofflineCapabilities             DVRGetSet = 3652 //获取停车场票箱脱机下参数配置 能力
	NetDvrSetPxofflineCfg                      DVRGetSet = 3653 //设置停车场票箱脱机下参数配置信息
	NetDvrGetPxofflineCfg                      DVRGetSet = 3654 //获取停车场票箱脱机下参数配置信息
	NetDvrGetPaperchargeinfoCapabilities       DVRGetSet = 3655 //获取停车场出入口纸票信息下发 能力
	NetDvrSetPaperchargeinfo                   DVRGetSet = 3656 //设置停车场出入口纸票信息下发
	NetDvrGetParkingsapceCapabilities          DVRGetSet = 3657 //获取停车场出入口停车位信息下发 能力
	NetDvrSetParkingsapceInfo                  DVRGetSet = 3658 //设置停车场出入口停车位信息下发
	NetDvrGetPxmultictrlCapabilities           DVRGetSet = 3659 //获取停车场票箱从属设备多角度参数配置 能力
	NetDvrGetChargeaccountCapabilities         DVRGetSet = 3661 //获取停车场票箱参数配置能力
	NetDvrSetChargeAccountinfo                 DVRGetSet = 3662 //设置缴费金额信息
	NetDvrSetPxmultictrlCfg                    DVRGetSet = 3663 //设置停车场票箱从属设备多角度参数配置信息
	NetDvrGetPxmultictrlCfg                    DVRGetSet = 3664 //获取停车场票箱从属设备多角度参数配置信息
	NetDvrGetTmeChargerule                     DVRGetSet = 3665 //获取停车场出入口车卡收费规则规则
	NetDvrSetTmeChargerule                     DVRGetSet = 3666 //设置停车场出入口车卡收费规则规则
	NetDvrGetTmeChargeruleCapabilities         DVRGetSet = 3667 //获取停车场出入口 车卡收费信息配置能力
	NetDvrGetIllegalcardfilteringCapabilities  DVRGetSet = 3668 //获取停车场票箱参数配置能力
	NetDvrGetIllegalcardfilteringCfg           DVRGetSet = 3669 //获取停车场票箱参数配置
	NetDvrSetIllegalcardfilteringCfg           DVRGetSet = 3670 //设置停车场票箱参数配置
	NetDvrGetLeddisplayCapabilities            DVRGetSet = 3671 //获取LED屏幕显示参数配置参数能力
	NetDvrSetLeddisplayCfg                     DVRGetSet = 3672 //设置LED屏幕显示参数
	NetDvrGetLeddisplayCfg                     DVRGetSet = 3673 //获取LED屏幕显示参数
	NetDvrGetVoicebroadcastCapabilities        DVRGetSet = 3674 //获取语音播报控制参数配置参数能力
	NetDvrSetVoicebroadcastCfg                 DVRGetSet = 3675 //设置语音播报控制参数
	NetDvrGetPaperprintformatCapabilities      DVRGetSet = 3676 //获取纸票打印格式配置能力
	NetDvrGetPaperprintformatCfg               DVRGetSet = 3677 //获取纸票打印格式参数配置
	NetDvrSetPaperprintformatCfg               DVRGetSet = 3678 //设置纸票打印格式参数配置
	NetDvrGetLoCkGateCapabilities              DVRGetSet = 3679 //获取智能锁闸配置能力
	NetDvrGetLockgateCfg                       DVRGetSet = 3680 //获取智能锁闸参数配置
	NetDvrSetLockgateCfg                       DVRGetSet = 3681 //设置智能锁闸参数配置
	NetDvrGetParkingDatastate                  DVRGetSet = 3682 //获取数据同步状态
	NetDvrSetParkingDatastate                  DVRGetSet = 3683 //设置数据同步状态
	NetDvrGetTmeCapabilities                   DVRGetSet = 3684 //获取停车场出入口设备 能力
	NetDvrGetTmevoiceCapabilities              DVRGetSet = 3686 //获取语音配置信息能力
	NetDvrSetTmevoiceCfg                       DVRGetSet = 3687 //设置语音参数配置
	NetDvrGetTmevoiceCfg                       DVRGetSet = 3688 //获取语音参数配置
	NetDvrDelTmevoiceCfg                       DVRGetSet = 3689 //删除语音参数配置
	NetDvrGetPosition                          DVRGetSet = 3698 // 获取方位矫正配置参数
	NetDvrSetPosition                          DVRGetSet = 3699 // 设置方位矫正配置参数
	NetDvrGetCentralizedctrlCapabilities       DVRGetSet = 3700 //获取集中布控能力
	NetDvrGetCentralizedctrl                   DVRGetSet = 3701 //获取集中布控参数配置
	NetDvrSetCentralizedctrl                   DVRGetSet = 3702 //设置集中布控参数配置
	NetDvrGetCompassCapabilities               DVRGetSet = 3703 //获取电子罗盘能力
	NetDvrGetVandalproofalarm                  DVRGetSet = 3704 //获取防破坏报警参数配置
	NetDvrSetVandalproofalarm                  DVRGetSet = 3705 //设置防破坏报警参数配置
	NetDvrCompassCalibrateCtrl                 DVRGetSet = 3706 //电子罗盘矫正控制接口
	NetDvrCompassNorthCtrl                     DVRGetSet = 3707 //电子罗盘指向正北控制接口
	NetDvrGetAzimuthinfo                       DVRGetSet = 3708 //获取方位角度参数配置
	NetDvrGetSatellitetime                     DVRGetSet = 3709 //获取卫星定位参数配置
	NetDvrSetSatellitetime                     DVRGetSet = 3710 //设置卫星定位参数配置
	NetDvrGetGisinfo                           DVRGetSet = 3711 //获取当前球机的GIS信息数据
	NetDvrGetStreamingCapabilities             DVRGetSet = 3712 //获取视频流的能力
	NetDvrGetRefreshframeCapabilities          DVRGetSet = 3713 //获取刷新帧的能力
	NetDvrStreamingRefreshFrame                DVRGetSet = 3714 //取流预览的强制刷新帧
	NetDvrFacecaptureStatistics                DVRGetSet = 3715 //长连接人员统计
	NetDvrGetWirelessserverCapabilities        DVRGetSet = 3716 //获取热点功能配置协议的能力
	NetDvrGetWirelessserver                    DVRGetSet = 3717 //获取热点功能配置协议
	NetDvrSetWirelessserver                    DVRGetSet = 3718 //设置热点功能配置协议
	NetDvrGetConnectListCapabilities           DVRGetSet = 3719 //获取连接设备列表信息的能力
	NetDvrGetThscreenCapabilities              DVRGetSet = 3720 //获取温湿度配置协议的能力
	NetDvrGetThscreen                          DVRGetSet = 3721 //获取温湿度配置协议
	NetDvrGetExternaldeviceCapabilities        DVRGetSet = 3722 //获取外设配置协议的能力
	NetDvrGetExternaldevice                    DVRGetSet = 3723 //获取外设配置协议
	NetDvrSetExternaldevice                    DVRGetSet = 3724 //设置外设配置协议
	NetDvrGetLeddisplayinfoCapabilities        DVRGetSet = 3725 //获取LED显示信息的能力
	NetDvrSetLeddisplayinfo                    DVRGetSet = 3726 //设置LED显示信息
	NetDvrGetSupplementlightCapabilities       DVRGetSet = 3727 //获取内置补光灯配置协议的能力 (球机支持，软件实现，补光灯是设计在设备内部的)
	NetDvrGetSupplementlight                   DVRGetSet = 3728 //获取内置补光灯配置协议
	NetDvrSetSupplementlight                   DVRGetSet = 3729 //设置内置补光灯配置协议
	NetDvrSetThscreen                          DVRGetSet = 3730 //设置温湿度配置协议
	NetDvrGetLowpowerCapabilities              DVRGetSet = 3731 //获取低功耗配置协议的能力
	NetDvrGetLowpower                          DVRGetSet = 3732 //获取低功耗配置协议
	NetDvrSetLowpower                          DVRGetSet = 3733 //设置低功耗配置协议
	NetDvrGetZoomlinkageCapabilities           DVRGetSet = 3734 //获取变倍联动配置协议的能力
	NetDvrGetZoomlinkage                       DVRGetSet = 3735 //获取变倍联动配置协议
	NetDvrSetZoomlinkage                       DVRGetSet = 3736 //设置变倍联动配置协议
	NetDvrThscreenTiming                       DVRGetSet = 3737 //温湿度
	NetDvrGetOsdBatteryPowerCfg                DVRGetSet = 3741 //获取OSD电池电量显示参数
	NetDvrSetOsdBatteryPowerCfg                DVRGetSet = 3742 //设置OSD电池电量显示参数
	NetDvrGetOsdBatteryPowerCfgCapabilities    DVRGetSet = 3743 //OSD电池电量显示参数的能力
	NetDvrGetVandalproofalarmTrigger           DVRGetSet = 3744 //获取防破坏报警联动配置
	NetDvrSetVandalproofalarmTrigger           DVRGetSet = 3745 //设置防破坏报警联动配置
	NetDvrGetPanoramaimageCapabilities         DVRGetSet = 3746 //获取全景图像的能力
	NetDvrGetPanoramaimage                     DVRGetSet = 3747 //获取全景图像参数的协议
	NetDvrSetPanoramaimage                     DVRGetSet = 3748 //设置全景图像参数的协议
	NetDvrGetStreamencryption                  DVRGetSet = 3749 //获取码流加密配置
	NetDvrSetStreamencryption                  DVRGetSet = 3750 //设置码流加密配置
	NetDvrGetStreamencryptionCapabilities      DVRGetSet = 3751 //获取码流加密能力
	NetDvrGetReviseGpsCapabilities             DVRGetSet = 3752 //获取校准GPS经纬度能力
	NetDvrGetReviseGps                         DVRGetSet = 3753 //获取校准GPS经纬度能力
	NetDvrSetReviseGps                         DVRGetSet = 3754 //设置校准GPS经纬度能力
	NetDvrGetPdcRecommend                      DVRGetSet = 3755 //获取客流统计表示推荐值
	NetDvrRemoveFlashstorage                   DVRGetSet = 3756 //客流数据清除操作
	NetDvrGetCountingCapabilities              DVRGetSet = 3757 //获取客流量统计能力
	NetDvrSetSensorAdjustment                  DVRGetSet = 3758 //设置Sensor 调节参数的协议
	NetDvrGetSensorAdjustmentCapabilities      DVRGetSet = 3759 //获取Sensor 调节参数的协议的能力
	NetDvrGetWirelessserverFullversionCfg      DVRGetSet = 3760 //获取wifi热点参数配置(完整版)
	NetDvrSetWirelessserverFullversionCfg      DVRGetSet = 3761 //设置wifi热点参数配置(完整版)
	NetDvrGetOnlineuserInfo                    DVRGetSet = 3762 //长连接获取用户在线信息
	NetDvrGetSensorAdjustmentInfo              DVRGetSet = 3763 //获取指定sensor调节参数
	NetDvrSensorResetCtrl                      DVRGetSet = 3764 //Sensor 调节复位
	NetDvrGetPostradarCapabilities             DVRGetSet = 3765 //获取雷达测速配置能力
	NetDvrGetPostradarspeedCfg                 DVRGetSet = 3766 //获取雷达测速配置
	NetDvrSetPostradarspeedCfg                 DVRGetSet = 3767 //设置雷达测速配置
	NetDvrGetPostradarspeedRecomCfg            DVRGetSet = 3768 //获取雷达测速推荐值
	NetDvrGetPostradarparamCfg                 DVRGetSet = 3769 //获取雷达参数配置
	NetDvrSetPostradarparamCfg                 DVRGetSet = 3770 //设置雷达参数配置
	NetDvrGetPostradarparamRecomCfg            DVRGetSet = 3771 //获取雷达参数推荐值
	NetDvrGetEncryptDeviceInfo                 DVRGetSet = 3772 //获取加密设备信息
	NetDvrGetAnrArmingHost                     DVRGetSet = 3773 //获取断网续传的主机信息
	NetDvrGetFirmwareVersion                   DVRGetSet = 3776 //GET firmware version
	/********************************IPC基线FF车牌****************************/
	NetDvrGetFtpCapabilities              DVRGetSet = 3782 //获取ftp能力
	NetDvrGetFtpuploadCfg                 DVRGetSet = 3783 //获取ftp上传信息规整参数
	NetDvrSetFtpuploadCfg                 DVRGetSet = 3784 //设置ftp上传信息规整参数
	NetDvrGetVehicleInformation           DVRGetSet = 3785 //获取车辆信息
	NetDvrGetDdnsCountryAbility           DVRGetSet = 3800 //获取设备支持的DDNS国家能力列表
	NetDvrGetDevicecfgV50                 DVRGetSet = 3801 //获取设备参数
	NetDvrSetDevicecfgV50                 DVRGetSet = 3802 //设置设备参数
	NetDvrSetVehicleRecogTaskV50          DVRGetSet = 3851 //车辆二次识别任务提交V50扩展
	NetDvrGetSmartcalibrationCapabilities DVRGetSet = 3900 // Smart行为标定过滤尺寸功能能力
	NetDvrGetTemperatureTrigger           DVRGetSet = 3903 //获取测温差联动配置
	NetDvrSetTemperatureTrigger           DVRGetSet = 3904 //设置测温差联动配置

	NetDvrGetSmartcalibrationCfg              DVRGetSet = 3910 //获取Smart行为标定过滤尺寸功能
	NetDvrSetSmartcalibrationCfg              DVRGetSet = 3911 //设置Smart行为标定过滤尺寸功能
	NetDvrPostSetupCalib                      DVRGetSet = 3912 //架设标定
	NetDvrSetPosInfoOverlay                   DVRGetSet = 3913 //设置Pos信息码流叠加控制
	NetDvrGetPosInfoOverlay                   DVRGetSet = 3914 //获取Pos信息码流叠加控制
	NetDvrGetCameraWorkMode                   DVRGetSet = 3915 //设置相机工作模式参数
	NetDvrSetCameraWorkMode                   DVRGetSet = 3916 //获取相机工作模式参数
	NetDvrGetResolutionSwitchCapabilities     DVRGetSet = 3917 //获取分辨率模式切换能力
	NetDvrGetResolutionSwitch                 DVRGetSet = 3918 //获取分辨率模式切换配置
	NetDvrSetResolutionSwitch                 DVRGetSet = 3919 //设置分辨率模式切换配置
	NetDvrGetConfirmMechanismCapabilities     DVRGetSet = 3920 //报警上传确认机制控制能力
	NetDvrConfirmMechanismCtrl                DVRGetSet = 3921 //报警上传确认机制控制
	NetDvrGetVehiclleResultCapabilities       DVRGetSet = 3951 //获取获取车辆信息结果能力
	NetDvrGetCalibCapabilities                DVRGetSet = 3952 //获取架设标定能力
	NetDvrGetPosinfoOverlayCapabilities       DVRGetSet = 3953 //获取获取Pos叠加能力
	NetSdkFindmedicalfile                     DVRGetSet = 3954 //慧影科技智慧医疗查找录像文件
	NetSdkFindmedicalpicture                  DVRGetSet = 3955 //慧影科技智慧医疗查找图片文件
	NetDvrSetPosinfoOverlay                   DVRGetSet = 3960 //设置Pos叠加
	NetDvrGetPosinfoOverlay                   DVRGetSet = 3961 //获取Pos叠加
	NetDvrGetFacelibTrigger                   DVRGetSet = 3962 //获取人脸比对库的联动配置
	NetDvrSetFacelibTrigger                   DVRGetSet = 3963 //设置人脸比对库的联动配置
	NetDvrGetFacecontrastTrigger              DVRGetSet = 3965 //获取人脸比对联动配置
	NetDvrSetFacecontrastTrigger              DVRGetSet = 3966 //设置人脸比对联动配置
	NetDvrGetFacecontrastScheduleCapabilities DVRGetSet = 3967 //获取人脸比对布防时间能力
	NetDvrGetFacecontrastSchedule             DVRGetSet = 3968 //获取人脸比对布防时间配置
	NetDvrSetFacecontrastSchedule             DVRGetSet = 3969 //设置人脸比对布防时间配置
	NetDvrGetFacelibScheduleCapabilities      DVRGetSet = 3970 //获取人脸比对库的布防时间能力
	NetDvrGetVcaVersionList                   DVRGetSet = 3973 //获取算法库版本
	NetDvrGetSetupCalib                       DVRGetSet = 3974 //获取架设标定
	NetDvrGetPanoramaLinkage                  DVRGetSet = 3975 //获取联动抓图上传使能配置
	NetDvrSetPanoramaLinkage                  DVRGetSet = 3976 //设置联动抓图上传使能配置
	NetDvrGetFacelibSchedule                  DVRGetSet = 3977 //获取人脸比对库的布防时间配置
	NetDvrSetFacelibSchedule                  DVRGetSet = 3978 //设置人脸比对库的布防时间配置
	NetDvrGetSoftwareServiceCapabilities      DVRGetSet = 3980 //获取软件服务能力
	NetDvrGetSoftwareService                  DVRGetSet = 3981 //获取软件服务配置
	NetDvrSetSoftwareService                  DVRGetSet = 3982 //设置软件服务配置
	NetDvrGetPreviewModeCapabilities          DVRGetSet = 3983 //获取预览模式配置能力
	NetDvrSetEagleFocusGotoscene              DVRGetSet = 3984 //鹰式聚焦设置摄像机转向指定的场景ID
	NetDvrEagleFocusSceneDel                  DVRGetSet = 3985 //删除鹰式聚焦标定的场景
	NetDvrGetSafetyHelmetTrigger              DVRGetSet = 3986 //获取安全帽检测联动配置
	NetDvrSetSafetyHelmetTrigger              DVRGetSet = 3987 //设置安全帽检测联动配置
	NetDvrGetSafetyHelmetScheduleCapabilities DVRGetSet = 3988 //获取安全帽检测布防时间能力
	NetDvrGetSafetyHelmetSchedule             DVRGetSet = 3989 //获取安全帽检测布防时间配置
	NetDvrSetSafetyHelmetSchedule             DVRGetSet = 3990 //设置安全帽检测布防时间配置

	NetDvrGetSignAbnormalTrigger DVRGetSet = 4150 //获取体征异常联动配置
	NetDvrSetSignAbnormalTrigger DVRGetSet = 4151 //设置体征异常联动配置

	NetDvrOneKeyConfigSanV50              DVRGetSet = 4152 //一键配置SAN(V50)
	NetDvrGetHdcfgV50                     DVRGetSet = 4153 //获取硬盘信息参数V50
	NetDvrSetHdcfgV50                     DVRGetSet = 4154 //设置硬盘信息参数V50
	NetDvrGetHdvolumeCfg                  DVRGetSet = 4155 //获取硬盘卷信息
	NetDvrSetHdvolumeCfg                  DVRGetSet = 4156 //设置硬盘卷信息
	NetDvrGetPowerSupplyCabinetTrigger    DVRGetSet = 4157 //获取机柜供电检测的联动配置
	NetDvrSetPowerSupplyCabinetTrigger    DVRGetSet = 4158 //设置机柜供电检测的联动配置
	NetDvrGetSensorTrigger                DVRGetSet = 4159 //获取传感器检测的联动配置
	NetDvrSetSensorTrigger                DVRGetSet = 4160 //设置传感器检测的联动配置
	NetDvrGetFacesnapTrigger              DVRGetSet = 4161 //获取人脸抓拍联动配置
	NetDvrSetFacesnapTrigger              DVRGetSet = 4162 //设置人脸抓拍联动配置
	NetDvrGetFacesnapScheduleCapabilities DVRGetSet = 4163 //获取人脸抓拍布防时间能力
	NetDvrGetFacesnapSchedule             DVRGetSet = 4164 //获取人脸抓拍布防时间配置
	NetDvrSetFacesnapSchedule             DVRGetSet = 4165 //设置人脸抓拍布防时间配置
	NetDvrSetScreenSwitch                 DVRGetSet = 4171 //画面切换控制
	NetDvrGetBvCalibPic                   DVRGetSet = 4172 //获取设备抓取图片和附加信息
	NetDvrGetBvCalibResult                DVRGetSet = 4173 //获取双目外参标定结果
	NetDvrGetBvHcorrection                DVRGetSet = 4174 //获取双目高度矫正数据
	NetDvrDelBvCalibPic                   DVRGetSet = 4175 //删除样本图片
	NetDvrGetTvScreenCfg                  DVRGetSet = 4176 //获取导播画面停留时间配置
	NetDvrSetTvScreenCfg                  DVRGetSet = 4177 //设置导播画面停留时间配置
	NetDvrAdjustBvCalib                   DVRGetSet = 4178 //双目标定微调
	NetDvrGetHumanCalib                   DVRGetSet = 4179 //获取人体坐标标定配置
	NetDvrSetHumanCalib                   DVRGetSet = 4180 //设置人体坐标标定配置
	NetDvrGetUsercfgV51                   DVRGetSet = 4181 //获取用户参数
	NetDvrSetUsercfgV51                   DVRGetSet = 4182 //设置用户参数
	NetDvrGetSoftioTrigger                DVRGetSet = 4183 //获取SoftIO联动配置
	NetDvrSetSoftioTrigger                DVRGetSet = 4184 //设置SoftIO联动配置
	NetDvrGetSoftioScheduleCapabilities   DVRGetSet = 4185 //获取SoftIO布防时间能力
	NetDvrGetSoftioSchedule               DVRGetSet = 4186 //获取SoftIO布防时间配置
	NetDvrSetSoftioSchedule               DVRGetSet = 4187 //设置SoftIO布防时间配置
	NetDvrGetHfpdTrigger                  DVRGetSet = 4188 //获取高频人员侦测联动配置
	NetDvrSetHfpdTrigger                  DVRGetSet = 4189 //设置高频人员侦测联动配置
	NetDvrGetHfpdScheduleCapabilities     DVRGetSet = 4190 //获取高频人员侦测布防时间能力
	NetDvrGetHfpdSchedule                 DVRGetSet = 4191 //获取高频人员侦测布防时间配置
	NetDvrSetHfpdSchedule                 DVRGetSet = 4192 //设置高频人员侦测布防时间配置
	NetDvrGetAlarmInfo                    DVRGetSet = 4193 //获取报警事件信息
	NetDvrGetUsercfgV52                   DVRGetSet = 4194 //获取用户参数
	NetDvrSetUsercfgV52                   DVRGetSet = 4195 //设置用户参数

	/********************************NVR_后端产品线****************************/
	NetDvrGetMutexFunction DVRGetSet = 4353 //获取功能互斥信息

	NetDvrGetSingleChannelinfo         DVRGetSet = 4360 //获取单个通道属性数据
	NetDvrGetChannelinfo               DVRGetSet = 4361 //获取通道属性数据
	NetDvrCheckLoginPasswordcfg        DVRGetSet = 4362 //用户登录密码校验
	NetDvrGetSingleSecurityQuestionCfg DVRGetSet = 4363 //获取单个设备安全问题
	NetDvrSetSingleSecurityQuestionCfg DVRGetSet = 4364 //设置单个设备安全问题
	NetDvrGetSecurityQuestionCfg       DVRGetSet = 4365 //获取设备安全问题
	NetDvrSetSecurityQuestionCfg       DVRGetSet = 4366 //设置设备安全问题
	NetDvrGetOnlineuserlistSc          DVRGetSet = 4367 //远程获取登陆用户信息（短连接）

	NetDvrGetBlacklistFacecontrastTrigger              DVRGetSet = 4368 //获取黑名单人脸比对联动配置
	NetDvrSetBlacklistFacecontrastTrigger              DVRGetSet = 4369 //设置黑名单人脸比对联动配置
	NetDvrGetWhitelistFacecontrastTrigger              DVRGetSet = 4370 //获取白名单人脸比对联动配置
	NetDvrSetWhitelistFacecontrastTrigger              DVRGetSet = 4371 //设置白名单人脸比对联动配置
	NetDvrGetBlacklistFacecontrastScheduleCapabilities DVRGetSet = 4372 //获取黑名单人脸比对布防时间能力
	NetDvrGetBlacklistFacecontrastSchedule             DVRGetSet = 4373 //获取黑名单人脸比对布防时间配置
	NetDvrSetBlacklistFacecontrastSchedule             DVRGetSet = 4374 //设置黑名单人脸比对布防时间配置
	NetDvrGetWhitelistFacecontrastScheduleCapabilities DVRGetSet = 4375 //获取白名单人脸比对布防时间能力
	NetDvrGetWhitelistFacecontrastSchedule             DVRGetSet = 4376 //获取白名单人脸比对布防时间配置
	NetDvrSetWhitelistFacecontrastSchedule             DVRGetSet = 4377 //设置白名单人脸比对布防时间配置

	NetDvrGetHumanRecognitionScheduleCapabilities DVRGetSet = 4378 //获取人体识别布防时间能力
	NetDvrGetHumanRecognitionSchedule             DVRGetSet = 4379 //获取人体识别布防时间配置
	NetDvrSetHumanRecognitionSchedule             DVRGetSet = 4380 //设置人体识别布防时间配置
	NetDvrGetHumanRecognitionTrigger              DVRGetSet = 4381 //获取人体识别联动配置
	NetDvrSetHumanRecognitionTrigger              DVRGetSet = 4382 //设置人体识别联动配置
	NetDvrGetGbt28181AudioOutputCfg               DVRGetSet = 4383 //获取GBT28181协议接入设备的语音对讲信息
	NetDvrSetGbt28181AudioOutputCfg               DVRGetSet = 4384 //设置GBT28181协议接入设备的语音对讲信息

	NetDvrGetStudentsStoodupTrigger                   DVRGetSet = 4386 //获取学生起立检测联动配置
	NetDvrSetStudentsStoodupTrigger                   DVRGetSet = 4387 //设置学生起立检测联动配置
	NetDvrGetFramesPeopleCountingScheduleCapabilities DVRGetSet = 4388 //获取区域人数统计布防时间能力
	NetDvrGetFramesPeopleCountingSchedule             DVRGetSet = 4389 //获取区域人数统计布防时间配置
	NetDvrSetFramesPeopleCountingSchedule             DVRGetSet = 4390 //设置区域人数统计布防时间配置
	NetDvrGetFramesPeopleCountingTrigger              DVRGetSet = 4391 //获取区域人数统计联动配置
	NetDvrSetFramesPeopleCountingTrigger              DVRGetSet = 4392 //设置区域人数统计联动配置

	NetDvrGetPersondensityTrigger                DVRGetSet = 4393 //获取人员密度检测的联动配置
	NetDvrSetPersondensityTrigger                DVRGetSet = 4394 //设置人员密度检测的联动配置
	NetDvrGetPersondensityScheduleCapabilities   DVRGetSet = 4395 //获取人员密度检测的布防时间能力
	NetDvrGetPersondensitySchedule               DVRGetSet = 4396 //获取人员密度检测的布防时间配置
	NetDvrSetPersondensitySchedule               DVRGetSet = 4397 //设置人员密度检测的布防时间配置
	NetDvrGetStudentsStoodupScheduleCapabilities DVRGetSet = 4398 //获取学生起立检测布防时间能力
	NetDvrGetStudentsStoodupSchedule             DVRGetSet = 4399 //获取学生起立检测布防时间配置
	NetDvrSetStudentsStoodupSchedule             DVRGetSet = 4400 //设置学生起立检测布防时间配置
	NetDvrSetFaceThermometryTrigger              DVRGetSet = 4401 //设置人脸测温联动配置
	NetDvrGetFaceThermometryScheduleCapabilities DVRGetSet = 4402 //获取人脸测温布防时间能力
	NetDvrGetFaceThermometrySchedule             DVRGetSet = 4403 //获取人脸测温布防时间配置
	NetDvrSetFaceThermometrySchedule             DVRGetSet = 4404 //设置人脸测温布防时间配置
	NetDvrGetFaceThermometryTrigger              DVRGetSet = 4405 //获取人脸测温联动配置
	NetDvrGetPersonqueueTrigger                  DVRGetSet = 4406 //获取人员排队检测的联动配置
	NetDvrSetPersonqueueTrigger                  DVRGetSet = 4407 //设置人员排队检测的联动配置
	NetDvrGetPersonqueueScheduleCapabilities     DVRGetSet = 4408 //获取人员排队检测的布防时间能力
	NetDvrGetPersonqueueSchedule                 DVRGetSet = 4409 //获取人员排队检测的布防时间配置
	NetDvrSetPersonqueueSchedule                 DVRGetSet = 4410 //设置人员排队检测的布防时间配置

	/********************************智能人脸识别****************************/
	NetDvrGetFacesnapcfg          DVRGetSet = 5001 //获取人脸抓拍参数
	NetDvrSetFacesnapcfg          DVRGetSet = 5002 //设置人脸抓拍参数
	NetDvrGetDevaccessCfg         DVRGetSet = 5005 //获取接入设备参数
	NetDvrSetDevaccessCfg         DVRGetSet = 5006 //设置接入设备参数
	NetDvrGetSavePathCfg          DVRGetSet = 5007 //获取存储信息参数
	NetDvrSetSavePathCfg          DVRGetSet = 5008 //设置存储信息参数
	NetVcaGetRulecfgV41           DVRGetSet = 5011 //获取行为分析参数(扩展)
	NetVcaSetRulecfgV41           DVRGetSet = 5012 //设置行为分析参数(扩展)
	NetDvrGetAidRulecfgV41        DVRGetSet = 5013 //获取交通事件规则参数
	NetDvrSetAidRulecfgV41        DVRGetSet = 5014 //设置交通事件规则参数
	NetDvrGetTpsRulecfgV41        DVRGetSet = 5015 //获取交通统计规则参数(扩展)
	NetDvrSetTpsRulecfgV41        DVRGetSet = 5016 //设置交通统计规则参数(扩展)
	NetVcaGetFacedetectRulecfgV41 DVRGetSet = 5017 //获取ATM人脸检测规则(扩展)
	NetVcaSetFacedetectRulecfgV41 DVRGetSet = 5018 //设置ATM人脸检测规则(扩展)
	NetDvrGetPdcRulecfgV41        DVRGetSet = 5019 //设置人流量统计规则(扩展)
	NetDvrSetPdcRulecfgV41        DVRGetSet = 5020 //获取人流量统计规则(扩展)
	NetDvrGetTrialVersionCfg      DVRGetSet = 5021 //获取试用版信息
	NetDvrGetVcaCtrlinfoCfg       DVRGetSet = 5022 //批量获取智能控制参数
	NetDvrSetVcaCtrlinfoCfg       DVRGetSet = 5023 //批量设置智能控制参数
	NetDvrSynChannelName          DVRGetSet = 5024 //同步通道名
	NetDvrGetResetCounter         DVRGetSet = 5025 //获取统计数据清零参数（人流量、交通统计）
	NetDvrSetResetCounter         DVRGetSet = 5026 //设置统计数据清零参数（人流量、交通统计）
	NetDvrGetObjectColor          DVRGetSet = 5027 //获取物体颜色属性
	NetDvrSetObjectColor          DVRGetSet = 5028 //设置物体颜色属性
	NetDvrGetAuxArea              DVRGetSet = 5029 //获取辅助区域
	NetDvrSetAuxArea              DVRGetSet = 5030 //设置辅助区域
	NetDvrGetChanWorkmode         DVRGetSet = 5031 //获取通道工作模式
	NetDvrSetChanWorkmode         DVRGetSet = 5032 //设置通道工作模式
	NetDvrGetSlaveChannel         DVRGetSet = 5033 //获取从通道参数
	NetDvrSetSlaveChannel         DVRGetSet = 5034 //设置从通道参数
	NetDvrGetVqdEventRule         DVRGetSet = 5035 //获取视频质量诊断事件规则
	NetDvrSetVqdEventRule         DVRGetSet = 5036 //设置视频质量诊断事件规则
	NetDvrGetBaselineScene        DVRGetSet = 5037 //获取基准场景参数
	NetDvrSetBaselineScene        DVRGetSet = 5038 //设置基准场景参数
	NetDvrControlBaselineScene    DVRGetSet = 5039 //基准场景操作
	NetDvrSetVcaDetionCfg         DVRGetSet = 5040 //设置智能移动参数配置
	NetDvrGetVcaDetionCfg         DVRGetSet = 5041 //获取智能移动参数配置
	NetDvrGetStreamAttachinfoCfg  DVRGetSet = 5042 //获取码流附加信息配置
	NetDvrSetStreamAttachinfoCfg  DVRGetSet = 5043 //设置码流附加信息配置
	NetDvrGetBvCalibType          DVRGetSet = 5044 //获取双目标定类型
	NetDvrControlBvSampleCalib    DVRGetSet = 5045 //双目样本标定
	NetDvrGetBvSampleCalibCfg     DVRGetSet = 5046 //获取双目标定参数
	NetDvrGetRulecfgV42           DVRGetSet = 5049 //获取行为分析参数(支持16条规则扩展)
	NetDvrSetRulecfgV42           DVRGetSet = 5050 //设置行为分析参数(支持16条规则扩展)
	NetDvrSetVcaDetionCfgV40      DVRGetSet = 5051 //设置智能移动参数配置
	NetDvrGetVcaDetionCfgV40      DVRGetSet = 5052 //获取智能移动参数配置
	NetDvrSetFlashCfg             DVRGetSet = 5110 //写入数据到Flash 测试使用
	/********************************智能人脸识别 end****************************/

	//2014-12-03
	NetDvrGetT1TestCfg DVRGetSet = 5053 //产线测试配置接口（获取）
	NetDvrSetT1TestCfg DVRGetSet = 5054 ////产线测试配置接口（设置）

	/********************************ITS****************************/
	NetItsGetOverlapCfgV50          DVRGetSet = 5055 //获取字符叠加参数配置扩展
	NetItsSetOverlapCfgV50          DVRGetSet = 5056 //设置字符叠加参数配置扩展
	NetDvrGetParklampState          DVRGetSet = 5057 //获取停车场信号灯状态信息
	NetDvrGetCloudstorageCfg        DVRGetSet = 5058 //获取云存储配置参数
	NetDvrSetCloudstorageCfg        DVRGetSet = 5059 //设置云存储配置参数
	NetItsGetBaseInfo               DVRGetSet = 5060 //获取终端基本信息
	NetDvrGetSensorInfo             DVRGetSet = 5061 //传感器信息查询
	NetDvrSetSensorSwitch           DVRGetSet = 5062 //传感器远程控制
	NetItsGetImgmergeCfg            DVRGetSet = 5063 //获取图片合成配置参数
	NetItsSetImgmergeCfg            DVRGetSet = 5064 //设置图片合成配置参数
	NetItsGetUploadCfg              DVRGetSet = 5065 //获取数据上传配置
	NetItsSetUploadCfg              DVRGetSet = 5066 //设置数据上传配置
	NetDvrGetSensorPortCapabilities DVRGetSet = 5067 //获取传感器能力
	NetItsGetWorkstate              DVRGetSet = 5069 //获取终端工作状态
	NetItsGetIpcChanCfg             DVRGetSet = 5070 //获取通道IPC信息
	NetItsSetIpcChanCfg             DVRGetSet = 5071 //设置通道IPC信息
	NetItsGetOverlapCfg             DVRGetSet = 5072 //获取字符叠加参数配置
	NetItsSetOverlapCfg             DVRGetSet = 5073 //设置字符叠加参数配置
	NetDvrGetTriggerexCfg           DVRGetSet = 5074 //获取ITC扩展配置
	NetDvrSetTriggerexCfg           DVRGetSet = 5075 //设置ITC扩展配置
	NetItsGetRoadInfo               DVRGetSet = 5076 //获取路口信息
	NetItsRemoteDeviceControl       DVRGetSet = 5077 //设置远程设备控制
	NetItsGetGateipcChanCfg         DVRGetSet = 5078 //获取出入口参数
	NetItsSetGateipcChanCfg         DVRGetSet = 5079 //设置出入口参数
	NetItsTranschanStart            DVRGetSet = 5080 //同步数据服务器建立连接
	NetItsGetEctworkstate           DVRGetSet = 5081 //获取出入口终端工作状态
	NetItsGetEctChanInfo            DVRGetSet = 5082 //获取出入口终端通道状态
	NetDvrGetHeatmapResult          DVRGetSet = 5083 //热度图数据查找
	NetDvrSetItsExdevcfg            DVRGetSet = 5084 //设置ITS外接设备信息
	NetDvrGetItsExdevcfg            DVRGetSet = 5085 //获取ITS外接设备信息
	NetDvrGetItsExdevstatus         DVRGetSet = 5086 //获取ITS所有外接设备信息
	NetDvrSetItsEndevcmd            DVRGetSet = 5087 //设置ITS终端出入口控制命令
	NetDvrSetEnissuedDatadel        DVRGetSet = 5088 //设置ITS终端出入口控制清除
	NetDvrGetPdcResult              DVRGetSet = 5089 //客流量数据查询 2014-03-21
	NetItsGetLampCtrlcfg            DVRGetSet = 5090 //获取内外置灯参数
	NetItsSetLampCtrlcfg            DVRGetSet = 5091 //设置内外置灯参数
	NetItsGetParkspaceAttributeCfg  DVRGetSet = 5092 //获取特殊车位参数
	NetItsSetParkspaceAttributeCfg  DVRGetSet = 5093 //设置特殊车位参数
	NetItsSetLampExternalCfg        DVRGetSet = 5095 //设置外控配置参数
	NetItsSetCompelCapture          DVRGetSet = 5096 //设置车位强制抓图
	NetDvrSetTimesignCfg            DVRGetSet = 5097 //设置扩展校时自定义标记
	NetDvrGetTimesignCfg            DVRGetSet = 5098 //获取扩展校时自定义标记
	NetDvrGetSignallampStatus       DVRGetSet = 5099 //信号灯检测
	/********************************ITS end****************************/

	NetDvrGetMonitorPlanVqd    DVRGetSet = 5100 //长连接获取诊断服务器计划
	NetDvrGetMonitoridVqd      DVRGetSet = 5101 //长连接获取对应计划内的监控点信息
	NetDvrSetMonitorInfo       DVRGetSet = 5102 //批量设置计划内的监控点信息
	NetDvrDelMonitorPlanVqd    DVRGetSet = 5103 //删除计划
	NetDvrGetMonitorVqdStatus  DVRGetSet = 5104 //平台查询诊断服务器的状态
	NetDvrGetRecordInfo        DVRGetSet = 5105 //获取资源图片查询
	NetDvrGetMonitorVqdcfg     DVRGetSet = 5106 //获取服务器的监控点信息
	NetDvrSetMonitorVqdcfg     DVRGetSet = 5107 //设置服务器的监控点信息
	NetDvrSetMonitorPlanVqdcfg DVRGetSet = 5108 //设置管理计划(单独的计划)
	NetDvrSceneChangeUpdate    DVRGetSet = 5109 //场景变更数据更新
	NetDvrGetCalibratePoint    DVRGetSet = 5153 //归一化坐标转换（枪球联动设备 外部交互命令码 基线代码不实现，防止冲突，提交基线）
	/*************************智能多场景********************************/
	NetDvrGetSceneCfg                DVRGetSet = 5201 //获取场景信息
	NetDvrSetSceneCfg                DVRGetSet = 5202 //设置场景信息
	NetDvrGetSceneReferenceRegion    DVRGetSet = 5203 //获取参考区域
	NetDvrSetSceneReferenceRegion    DVRGetSet = 5204 //设置参考区域
	NetDvrGetSceneCalibration        DVRGetSet = 5205 //获取标定信息
	NetDvrSetSceneCalibration        DVRGetSet = 5206 //设置标定信息
	NetDvrGetSceneMaskRegion         DVRGetSet = 5207 //获取屏蔽区域
	NetDvrSetSceneMaskRegion         DVRGetSet = 5208 //设置屏蔽区域
	NetDvrGetSceneLanecfg            DVRGetSet = 5209 //获取车道规则
	NetDvrSetSceneLanecfg            DVRGetSet = 5210 //设置车道规则
	NetDvrGetSceneAidRulecfg         DVRGetSet = 5211 //获取交通事件规则参数
	NetDvrSetSceneAidRulecfg         DVRGetSet = 5212 //设置交通事件规则参数
	NetDvrGetSceneTpsRulecfg         DVRGetSet = 5213 //获取交通统计规则参数
	NetDvrSetSceneTpsRulecfg         DVRGetSet = 5214 //设置交通统计规则参数
	NetDvrGetSceneTimeCfg            DVRGetSet = 5215 //获取通道的场景时间段配置
	NetDvrSetSceneTimeCfg            DVRGetSet = 5216 //设置通道的场景时间段配置
	NetDvrGetForensicsMode           DVRGetSet = 5217 //获取取证方式参数
	NetDvrSetForensicsMode           DVRGetSet = 5218 //设置取证方式参数
	NetDvrForcestopForensicsCtrl     DVRGetSet = 5219 //强制停止取证
	NetDvrGetAlarmProcessCfg         DVRGetSet = 5220 //获取报警处理参数
	NetDvrSetAlarmProcessCfg         DVRGetSet = 5221 //设置报警处理参数
	NetDvrGetBlacklistAlarmInfo      DVRGetSet = 5222 //获取黑白名单报警轨迹
	NetDvrGetStorageResourceCfg      DVRGetSet = 5225 //获取存储资源参数
	NetDvrSetStorageResourceCfg      DVRGetSet = 5226 //设置存储资源参数
	NetDvrDelBlacklistAlarmRecord    DVRGetSet = 5227 //远程删除名单报警记录
	NetDvrSetBlacklistGroupInfo      DVRGetSet = 5229 //远程分组列表参数配置
	NetDvrDelBlacklistGroupInfo      DVRGetSet = 5230 //远程删除分组列表
	NetDvrGetBlacklistGroupInfo      DVRGetSet = 5231 //远程获取全部分组列表
	NetDvrSetBlacklistGroupRecordCfg DVRGetSet = 5232 //分组记录参数配置
	NetDvrGetBlacklistGroupRecordCfg DVRGetSet = 5234 //远程获取分组记录参数
	NetDvrDelBlacklistGroupRecordCfg DVRGetSet = 5235 //远程删除分组记录参数
	NetDvrGetAreaMonitorCfg          DVRGetSet = 5236 //获取区域监控点参数
	NetDvrSetAreaMonitorCfg          DVRGetSet = 5237 //设置区域监控点参数
	NetDvrDelAreaMonitorCfg          DVRGetSet = 5238 //删除区域监控点
	NetDvrRetrievalSnapRecord        DVRGetSet = 5240 //抓拍库检索
	NetDvrGetAlarmlist               DVRGetSet = 5241 //获取名单报警列表
	NetDvrDetectImage                DVRGetSet = 5242 //单张图片检测
	NetDvrGetSnapRecord              DVRGetSet = 5243 //获取抓拍记录
	NetDvrDelSnapRecord              DVRGetSet = 5244 //删除抓拍记录
	NetDvrGetFaceRecord              DVRGetSet = 5245 //远程获取人脸记录列表
	NetDvrSetFaceRecord              DVRGetSet = 5246 //添加人脸记录
	NetDvrDelFaceRecord              DVRGetSet = 5247 //删除人脸记录
	NetDvrGetFaceDatabase            DVRGetSet = 5248 //获取人脸库配置参数
	NetDvrSetFaceDatabase            DVRGetSet = 5249 //设置人脸库配置参数
	NetDvrDelFaceDatabase            DVRGetSet = 5250 //删除人脸库
	NetDvrRetrievalFaceDatabase      DVRGetSet = 5251 //人脸库检索
	NetDvrSetBlacklistRelDevCfg      DVRGetSet = 5252 //设备关联名单分组关联
	NetDvrDelBlacklistRelDev         DVRGetSet = 5253 //删除 设备关联名单分组信息
	/*************************智能多场景end*****************************/

	NetDvrGetDiskRaidInfo DVRGetSet = 6001 //获取磁盘Raid信息
	NetDvrSetDiskRaidInfo DVRGetSet = 6002 //设置磁盘Raid信息

	NetDvrGetDvrSynchronousIpc  DVRGetSet = 6005 //获取：是否为前端IPC同步设备参数
	NetDvrSetDvrSynchronousIpc  DVRGetSet = 6006 //设置：是否为前端IPC同步设备参数
	NetDvrSetDvrIpcPasswd       DVRGetSet = 6008 //设置：IPC用户名密码
	NetDvrGetDeviceNetUsingInfo DVRGetSet = 6009 //获取：当前设备网络资源使用情况
	NetDvrSetDvrIpcNet          DVRGetSet = 6012 //设置：设置前端IPC的网络地址
	NetDvrGetRecordChannelInfo  DVRGetSet = 6013 //获取：录像通道信息
	NetDvrSetRecordChannelInfo  DVRGetSet = 6014 //设置：录像通道信息
	NetDvrMountDisk             DVRGetSet = 6015 // 加载磁盘
	NetDvrUnmountDisk           DVRGetSet = 6016 // 卸载磁盘

	// CVR
	NetDvrGetStreamSrcInfo        DVRGetSet = 6017 //获取：流来源信息
	NetDvrSetStreamSrcInfo        DVRGetSet = 6018 //设置：流来源信息
	NetDvrGetStreamRecordInfo     DVRGetSet = 6019 //获取：流录像信息
	NetDvrSetStreamRecordInfo     DVRGetSet = 6020 //设置：流录像信息
	NetDvrGetStreamRecordStatus   DVRGetSet = 6021 //获取：流录像状态
	NetDvrSetStreamRecordStatus   DVRGetSet = 6022 //设置：流录像状态
	NetDvrGetStreamInfo           DVRGetSet = 6023 //获取已添加的流ID信息
	NetDvrGetStreamSrcInfoV40     DVRGetSet = 6024 //获取：流来源信息
	NetDvrSetStreamSrcInfoV40     DVRGetSet = 6025 //设置：流来源信息
	NetDvrGetRelocateInfo         DVRGetSet = 6026 //获取N+0模式下重定向信息
	NetDvrStartGopInfoPassback    DVRGetSet = 6032 //智能信息回填
	NetDvrGetChansRecordStatusCfg DVRGetSet = 6035 //获取通道录像状态信息
	NetDvrSetChansRecordStatusCfg DVRGetSet = 6036 //设置通道录像状态信息
	//NVR：96xx
	NetDvrGetIpAlarmGroupNum DVRGetSet = 6100 //获取：IP通道报警输入输出组数
	NetDvrGetIpAlarmIn       DVRGetSet = 6101 //获取：IP通道报警输入信息
	NetDvrGetIpAlarmOut      DVRGetSet = 6102 //获取：IP通道报警输出信息

	//9000 v2.2
	NetDvrGetFtpcfgSecond DVRGetSet = 6103 //获取图片上传FTP参数
	NetDvrSetFtpcfgSecond DVRGetSet = 6104 //设置图片上传FTP参数

	NetDvrGetDefaultVideoEffect        DVRGetSet = 6105 // 获取视频输入效果参数默认值
	NetDvrSetVideoEffect               DVRGetSet = 6106 // 设置通道视频输入图像参数
	NetDvrDelInvalidDisk               DVRGetSet = 6107 // 删除无效磁盘
	NetDvrGetDrawframeDiskQuotaCfg     DVRGetSet = 6109 //获取抽帧通道磁盘配额
	NetDvrSetDrawframeDiskQuotaCfg     DVRGetSet = 6110 //设置抽帧通道磁盘配额
	NetDvrGetNatCfg                    DVRGetSet = 6111 //获取NAT映射参数
	NetDvrSetNatCfg                    DVRGetSet = 6112 //设置NAT映射参数
	NetDvrGetAesKey                    DVRGetSet = 6113 //获取设备AES加密密钥
	NetDvrGetPoeCfg                    DVRGetSet = 6114 //获取POE参数
	NetDvrSetPoeCfg                    DVRGetSet = 6115 //设置POE参数
	NetDvrGetCustomProCfg              DVRGetSet = 6116 //获取自定义协议参数
	NetDvrSetCustomProCfg              DVRGetSet = 6117 //设置自定义协议参数
	NetDvrGetStreamCabac               DVRGetSet = 6118 //获取码流压缩性能选项
	NetDvrSetStreamCabac               DVRGetSet = 6119 //设置码流压缩性能选项
	NetDvrGetEsataMinisasUsageCfg      DVRGetSet = 6120 //获取eSATA和miniSAS用途
	NetDvrSetEsataMinisasUsageCfg      DVRGetSet = 6121 //设置eSATA和miniSAS用途
	NetDvrGetHdcfgV40                  DVRGetSet = 6122 //获取硬盘信息参数
	NetDvrSetHdcfgV40                  DVRGetSet = 6123 //设置硬盘信息参数
	NetDvrGetPoeChannelAddMode         DVRGetSet = 6124 //获取POE通道添加方式
	NetDvrSetPoeChannelAddMode         DVRGetSet = 6125 //设置POE通道添加方式
	NetDvrGetDigitalChannelState       DVRGetSet = 6126 //获取设备数字通道状态
	NetDvrGetBonjourCfg                DVRGetSet = 6127 // 获取Bonjour信息
	NetDvrSetBonjourCfg                DVRGetSet = 6128 // 设置Bonjour信息
	NetDvrGetSocksCfg                  DVRGetSet = 6130 //获取SOCKS信息
	NetDvrSetSocksCfg                  DVRGetSet = 6131 //设置SOCKS信息
	NetDvrGetQosCfg                    DVRGetSet = 6132 //获取QoS信息
	NetDvrSetQosCfg                    DVRGetSet = 6133 //设置QoS信息
	NetDvrGetHttpsCfg                  DVRGetSet = 6134 //获取HTTPS信息
	NetDvrSetHttpsCfg                  DVRGetSet = 6135 //设置HTTPS信息
	NetDvrGetWd1Cfg                    DVRGetSet = 6136 //远程获取WD1使能开关
	NetDvrSetWd1Cfg                    DVRGetSet = 6137 //远程设置WD1使能开关
	NetDvrCreateCert                   DVRGetSet = 6138 //创建证书
	NetDvrDeleteCert                   DVRGetSet = 6139 //删除证书
	NetDvrGetRecordLockPercentage      DVRGetSet = 6140 //获取录像段锁定比例
	NetDvrSetRecordLockPercentage      DVRGetSet = 6141 //设置录像段锁定比例
	NetDvrCmdTriggerPeriodRecord       DVRGetSet = 6144 //外部命令触发指定时间录像
	NetDvrUploadCert                   DVRGetSet = 6145 //上传证书
	NetDvrDownloadCert                 DVRGetSet = 6146 //下载证书
	NetDvrGetCert                      DVRGetSet = 6147 //获取证书
	NetDvrGetPosFilterCfg              DVRGetSet = 6148 //获取POS过滤规则
	NetDvrSetPosFilterCfg              DVRGetSet = 6149 //设置POS过滤规则
	NetDvrGetConnectPosCfg             DVRGetSet = 6150 //获取DVR与POS连接方式
	NetDvrSetConnectPosCfg             DVRGetSet = 6151 //设置DVR与POS连接方式
	NetDvrGetChanFilterCfg             DVRGetSet = 6152 //获取规则与通道关联信息
	NetDvrSetChanFilterCfg             DVRGetSet = 6153 //设置规则与通道关联信息
	NetDvrGetFtpcfgV40                 DVRGetSet = 6162 //获取FTP信息
	NetDvrSetFtpcfgV40                 DVRGetSet = 6163 //设置FTP信息
	NetDvrGetMonthlyRecordDistribution DVRGetSet = 6164 //获取月历录像分布
	NetDvrGetAccessDeviceChannelInfo   DVRGetSet = 6165 //获取待接入设备通道信息
	NetDvrGetPreviewSwitchCfg          DVRGetSet = 6166 //获取设备本地预览切换参数
	NetDvrSetPreviewSwitchCfg          DVRGetSet = 6167 //设置设备本地预览切换参数

	//Netra3.0.0
	NetDvrGetNPlusOneWorkMode        DVRGetSet = 6168 //获取N+1工作模式
	NetDvrSetNPlusOneWorkMode        DVRGetSet = 6169 //设置N+1工作模式
	NetDvrGetHdStatus                DVRGetSet = 6170 //获取硬盘状态
	NetDvrSetHdStatus                DVRGetSet = 6171 //设置硬盘状态
	NetDvrImportIpcCfgFile           DVRGetSet = 6172 //导入IPC配置文件
	NetDvrExportIpcCfgFile           DVRGetSet = 6173 //导出IPC配置文件
	NetDvrUpgradeIpc                 DVRGetSet = 6174 //升级IP通道
	NetDvrGetRaidBackgroundTaskSpeed DVRGetSet = 6175 //获取RAID后台任务速度
	NetDvrSetRaidBackgroundTaskSpeed DVRGetSet = 6176 //设置RAID后台任务速度

	//marvell 256路NVR
	NetDvrGetExceptioncfgV40      DVRGetSet = 6177 //获取异常参数配置
	NetDvrSetExceptioncfgV40      DVRGetSet = 6178 //设置异常参数配置
	NetDvrGetPiccfgV40            DVRGetSet = 6179 //获取图象参数 支持变长    NetSDK_
	NetDvrSetPiccfgV40            DVRGetSet = 6180 //设置图象参数， 支持变长
	NetDvrGetAlarmincfgV40        DVRGetSet = 6181 //获取报警输入参数，支持变长
	NetDvrSetAlarmincfgV40        DVRGetSet = 6182 //获取报警输入参数，支持变长
	NetDvrGetIpalarmincfgV40      DVRGetSet = 6183 //获取IP报警输入接入配置信息
	NetDvrGetIpalarmoutcfgV40     DVRGetSet = 6185 //获取IP报警输出接入配置信息
	NetDvrGetUsercfgV40           DVRGetSet = 6187 //获取用户参数
	NetDvrSetUsercfgV40           DVRGetSet = 6188 //设置用户参数
	NetDvrGetWorkStatus           DVRGetSet = 6189 //获取设备工作状态
	NetDvrGetJpegCaptureCfgV40    DVRGetSet = 6190 //获取DVR抓图配置
	NetDvrSetJpegCaptureCfgV40    DVRGetSet = 6191 //设置DVR抓图配置
	NetDvrGetHdgroupCfgV40        DVRGetSet = 6192 //获取盘组管理配置参数
	NetDvrSetHdgroupCfgV40        DVRGetSet = 6193 //设置盘组管理配置参数
	NetDvrGetSmdHolidayHandle     DVRGetSet = 6194 //获取简易智能假日计划
	NetDvrSetSmdHolidayHandle     DVRGetSet = 6195 //设置简易智能假日计划
	NetDvrGetPicModelCfg          DVRGetSet = 6196 //获取图片建模配置参数
	NetDvrSetPicModelCfg          DVRGetSet = 6197 //设置图片建模配置参数
	NetDvrStartLocalMouseEvent    DVRGetSet = 6198 //开启设备本地鼠标事件记录
	NetDvrStartSimulareMouseEvent DVRGetSet = 6199 //远程模拟鼠标事件
	NetDvrGetWorkStatusV50        DVRGetSet = 6200 //获取设备工作状态V50

	//91系列HD-SDI高清DVR
	NetDvrGetAccessCameraInfo DVRGetSet = 6201 // 获取前端相机信息
	NetDvrSetAccessCameraInfo DVRGetSet = 6202 // 设置前端相机信息
	NetDvrPullDisk            DVRGetSet = 6203 // 安全拔盘
	NetDvrScanRaid            DVRGetSet = 6204 // 扫描阵列
	// CVR 2.0.X
	NetDvrGetUserRightCfg     DVRGetSet = 6210 // 获取用户权限
	NetDvrSetUserRightCfg     DVRGetSet = 6211 // 设置用户权限
	NetDvrOneKeyConfig        DVRGetSet = 6212 // 一键配置CVR
	NetDvrRestartService      DVRGetSet = 6213 // 重启CVR服务
	NetDvrGetMaxMachineNumCfg DVRGetSet = 6214 // 获取备机最大个数
	NetDvrSetMaxMachineNumCfg DVRGetSet = 6215 // 设置备机最大个数
	NetDvrAddDevice           DVRGetSet = 6216 //N+1模式添加设备
	NetDvrDelDevice           DVRGetSet = 6217 //N+1模式删除设备

	NetDvrGetDataCallbackCfg DVRGetSet = 6218 // 获取数据回迁状态
	NetDvrSetDataCallbackCfg DVRGetSet = 6219 // 设置数据回迁状态

	NET_DVR_CLONE_LUN  DVRGetSet = 6220 //克隆LUN卷
	NET_DVR_EXPAND_LUN DVRGetSet = 6221 //扩展和重命名LUN卷

	NET_DVR_GET_N_PLUS_ONE_DEVICE_INFO DVRGetSet = 6222 //获取N+1设备信息
	NET_DVR_MODIFY_DVR_NET_DISK        DVRGetSet = 6223 //修改DVR网盘
	//NET_DVR_DEL_DVR_NET_DISK DVRGetSet = 6224 //删除DVR网盘

	NET_DVR_CREATE_NAS DVRGetSet = 6225 //创建NAS
	NET_DVR_DELETE_NAS DVRGetSet = 6226 //删除NAS

	NET_DVR_OPEN_ISCSI  DVRGetSet = 6227 //开启iSCSI
	NET_DVR_CLOSE_ISCSI DVRGetSet = 6228 //关闭iSCSI

	NET_DVR_GET_FC             DVRGetSet = 6229 //获取光纤信息
	NET_DVR_OPEN_FC            DVRGetSet = 6230 //开启FC
	NET_DVR_CLOSE_FC           DVRGetSet = 6231 //关闭FC
	NET_DVR_ONE_KEY_CONFIG_SAN DVRGetSet = 6232 // 一键配置SAN, 与一键配置CVR逻辑一样

	//CVR2.3.2
	NET_DVR_RECORD_CHECK                        DVRGetSet = 6233 //录像完整性检测
	NET_DVR_ADD_RECORD_PASSBACK_TASK_MANUAL     DVRGetSet = 6234 //手动添加录像回传任务
	NET_DVR_GET_ALL_RECORD_PASSBACK_TASK_MANUAL DVRGetSet = 6235 //获取所有手动添加录像回传任务
	NET_DVR_RECORD_PASSBACK_TASK_MANUAL_CTRL    DVRGetSet = 6236 //控制手动录像回传任务
	NET_DVR_DEL_RECORD_PASSBACK_TASK_MANUAL     DVRGetSet = 6237 //删除手动录像回传任务
	NET_DVR_GET_RECORD_PASSBACK_PLAN_CFG        DVRGetSet = 6238 //获取录像回传计划配置
	NET_DVR_SET_RECORD_PASSBACK_PLAN_CFG        DVRGetSet = 6239 //设置录像回传计划配置
	NET_DVR_GET_DEV_STORAGE_CFG                 DVRGetSet = 6240 //获取设备存储信息
	NET_DVR_GET_ONLINE_USER_CFG                 DVRGetSet = 6241 //获取在线用户参数
	NET_DVR_GET_RECORD_SEGMENT_CFG              DVRGetSet = 6242 //获取录像段总量

	NET_DVR_GET_REC_PASSBACK_TASK_EXECUTABLE DVRGetSet = 6243 //查询手动录像回传任务可执行性
	NET_DVR_GET_STREAM_MEDIA_CFG             DVRGetSet = 6244 //获取流媒体回传录像参数配置（流ID方式）
	NET_DVR_SET_STREAM_MEDIA_CFG             DVRGetSet = 6245 //设置流媒体回传录像参数配置（流ID方式）
	NET_DVR_GET_USERCFG_V50                  DVRGetSet = 6246 //获取用户参数V50
	NET_DVR_SET_USERCFG_V50                  DVRGetSet = 6247 //设置用户参数V50

	NET_DVR_GET_RECORD_PASSBACK_BASIC_CFG_CAP        DVRGetSet = 6248 //获取CVR回传功能基础配置能力
	NET_DVR_GET_RECORD_PASSBACK_BASIC_CFG            DVRGetSet = 6249 //获取CVR回传功能基础配置
	NET_DVR_SET_RECORD_PASSBACK_BASIC_CFG            DVRGetSet = 6250 //设置CVR回传功能基础配置
	NET_DVR_ONE_KEY_CONFIG_V50                       DVRGetSet = 6251 // 一键配置CVR(V50)
	NET_DVR_GET_RACM_CAP                             DVRGetSet = 6252 //获取存储总能力（RACM能力）
	NET_DVR_GET_THUMBNAILS                           DVRGetSet = 6253 //获取缩略图（默认是录像的缩略图）(支持流ID)
	NET_DVR_ADD_RECORD_PASSBACK_TASK_MANUAL_V50      DVRGetSet = 6254 //手动添加录像回传任务V50（返回任务ID）
	NET_DVR_GET_RECORD_PASSBACK_HISTORY_PLAN_CFG_CAP DVRGetSet = 6255 //获取CVR回传历史录像计划能力
	NET_DVR_GET_RECORD_PASSBACK_HISTORY_PLAN_CFG     DVRGetSet = 6256 //获取CVR回传历史录像计划配置
	NET_DVR_SET_RECORD_PASSBACK_HISTORY_PLAN_CFG     DVRGetSet = 6257 //设置CVR回传历史录像计划配置
	NET_DVR_ONE_KEY_CONFIG_V51                       DVRGetSet = 6258 // 一键配置CVR(V51)

	NET_DVR_GET_RECORD_PACK DVRGetSet = 6301 //获取录像打包参数
	NET_DVR_SET_RECORD_PACK DVRGetSet = 6302 //设置录像打包参数

	NET_DVR_GET_CLOUD_STORAGE_CFG DVRGetSet = 6303 //获取设备当前工作模式
	NET_DVR_SET_CLOUD_STORAGE_CFG DVRGetSet = 6304 //设置设备当前工作模式
	NET_DVR_GET_GOP_INFO          DVRGetSet = 6305 //获取GOP信息
	NET_DVR_GET_PHY_DISK_INFO     DVRGetSet = 6306 //获取物理磁盘信息
	//录播主机外部命令
	NET_DVR_GET_RECORDING_AUTO_TRACK_CFG DVRGetSet = 6307 //获取SDI自动跟踪配置信息
	NET_DVR_SET_RECORDING_AUTO_TRACK_CFG DVRGetSet = 6308 //设置SDI自动跟踪配置信息

	NET_DVR_GET_RECORDING_PUBLISH_CFG DVRGetSet = 6309 //获取一键发布信息
	NET_DVR_SET_RECORDING_PUBLISH_CFG DVRGetSet = 6310 //设置一键发布信息

	NET_DVR_RECORDING_ONEKEY_CONTROL DVRGetSet = 6311 //录播主机控制

	NET_DVR_GET_RECORDING_END_TIME DVRGetSet = 6312 //获取录播剩余时间

	NET_DVR_RECORDING_PUBLISH DVRGetSet = 6313 //一键发布录像

	NET_DVR_GET_CURRICULUM_CFG DVRGetSet = 6314 //获取课表配置信息
	NET_DVR_SET_CURRICULUM_CFG DVRGetSet = 6315 //设置课表配置信息

	NET_DVR_GET_COURSE_INDEX_CFG DVRGetSet = 6316 //获取课程信息索引
	NET_DVR_SET_COURSE_INDEX_CFG DVRGetSet = 6317 //设置课程信息索引

	NET_DVR_GET_PPT_CHANNEL    DVRGetSet = 6318 //获取PPT支持通道号
	NET_DVR_GET_PPT_DETECT_CFG DVRGetSet = 6319 //获取PPT检测参数
	NET_DVR_SET_PPT_DETECT_CFG DVRGetSet = 6320 //设置PPT检测参数

	NET_DVR_GET_RECORDINGHOST_CFG DVRGetSet = 6321 //获取录播主机配置信息
	NET_DVR_SET_RECORDINGHOST_CFG DVRGetSet = 6322 //设置录播主机配置信息
	NET_DVR_GET_BACKUP_RECORD_CFG DVRGetSet = 6323 //获取一键备份配置信息
	NET_DVR_SET_BACKUP_RECORD_CFG DVRGetSet = 6324 //设置一键备份配置信息

	//庭审主机
	NET_DVR_GET_AUDIO_ACTIVATION_CFG DVRGetSet = 6326 //获取语音激励配置参数
	NET_DVR_SET_AUDIO_ACTIVATION_CFG DVRGetSet = 6327 //设置语音激励配置参数
	NET_DVR_GET_DECODERCFG_V40       DVRGetSet = 6328 //获取解码器参数信息
	NET_DVR_SET_DECODERCFG_V40       DVRGetSet = 6329 //设置解码器参数信息

	NET_DVR_INFRARED_OUTPUT_CONTROL   DVRGetSet = 6330 //红外输出控制
	NET_DVR_GET_INFRARED_CMD_NAME_CFG DVRGetSet = 6331 //获取红外命令名称参数配置
	NET_DVR_SET_INFRARED_CMD_NAME_CFG DVRGetSet = 6332 //设置红外命令名称参数配置
	NET_DVR_START_INFRARED_LEARN      DVRGetSet = 6333 //远程红外学码

	NET_DVR_GET_TRIAL_SYSTEM_CFG        DVRGetSet = 6334 //获取庭审主机系统信息
	NET_DVR_SET_CASE_INFO               DVRGetSet = 6335 //案件信息录入
	NET_DVR_GET_TRIAL_MICROPHONE_STATUS DVRGetSet = 6336 //获取麦克风状态信息
	NET_DVR_SET_TRIAL_MICROPHONE_STATUS DVRGetSet = 6337 //获取麦克风状态信息
	NET_DVR_GET_TRIAL_HOST_STATUS       DVRGetSet = 6338 //获取庭审主机状态信息
	NET_DVR_GET_LAMP_OUT                DVRGetSet = 6339 //获取LAMP输出口信息
	NET_DVR_SET_LAMP_OUT                DVRGetSet = 6340 //设置LAMP输出口信息
	NET_DVR_LAMP_REMOTE_CONTROL         DVRGetSet = 6341 // LAMP控制
	NET_DVR_REMOTE_CONTROL_PLAY         DVRGetSet = 6342 //远程控制本地回放
	NET_DVR_GET_LOCAL_INPUT_CFG         DVRGetSet = 6343 //获取庭审主机状态信息庭审主机本地输入信息
	NET_DVR_SET_LOCAL_INPUT_CFG         DVRGetSet = 6344 //设置庭审主机本地输入信息
	NET_DVR_GET_CASE_INFO               DVRGetSet = 6345 //获取当前案件信息

	//审讯机外部命令
	NET_DVR_INQUEST_GET_CDW_STATUS            DVRGetSet = 6350 //获取审讯机刻录状态-长连接
	NET_DVR_GET_MIX_AUDIOIN_CFG               DVRGetSet = 6351 //获取混音输入口参数配置
	NET_DVR_SET_MIX_AUDIOIN_CFG               DVRGetSet = 6352 //设置混音输入口参数配置
	NET_DVR_GET_MIX_AUDIOOUT_CFG              DVRGetSet = 6353 //获取混音输出口参数配置
	NET_DVR_SET_MIX_AUDIOOUT_CFG              DVRGetSet = 6354 //设置混音输出口参数配置
	NET_DVR_GET_AUDIOIN_VOLUME_CFG            DVRGetSet = 6355 //获取音频输入口音量调节参数配置
	NET_DVR_SET_AUDIOIN_VOLUME_CFG            DVRGetSet = 6356 //设置音频输入口音量调节参数配置
	NET_DVR_GET_AREA_MASK_CFG                 DVRGetSet = 6357 //获取马赛克区域配置
	NET_DVR_SET_AREA_MASK_CFG                 DVRGetSet = 6358 //设置马赛克区域配置
	NET_DVR_GET_AUDIO_DIACRITICAL_CFG         DVRGetSet = 6359 //获取音频变音配置
	NET_DVR_SET_AUDIO_DIACRITICAL_CFG         DVRGetSet = 6360 //设置音频变音配置
	NET_DVR_GET_WIFI_DHCP_ADDR_CFG            DVRGetSet = 6361 //获WIFI DHCP 地址范围参数配置
	NET_DVR_SET_WIFI_DHCP_ADDR_CFG            DVRGetSet = 6362 //设WIFI DHCP 地址范围参数配置
	NET_DVR_GET_WIFI_CLIENT_LIST_INFO         DVRGetSet = 6363 //获取wifi热点下连接的设备信息
	NET_DVR_REMOTECONTROL_POWER_ON            DVRGetSet = 6364 //远程开机
	NET_DVR_GET_MULTISTREAM_RELATION_CHAN_CFG DVRGetSet = 6365 //获取多码流关联通道参数配置
	NET_DVR_SET_MULTISTREAM_RELATION_CHAN_CFG DVRGetSet = 6366 //设置多码流关联通道参数配置
	NET_DVR_GET_VIDEOOUT_RESOLUTION_CFG       DVRGetSet = 6367 //获取设备本地视频输出口分辨率
	NET_DVR_SET_VIDEOOUT_RESOLUTION_CFG       DVRGetSet = 6368 //设置设备本地视频输出口分辨率
	NET_DVR_GET_AUDIOOUT_VOLUME_CFG           DVRGetSet = 6369 //获取音频输出口音量调节参数配置
	NET_DVR_SET_AUDIOOUT_VOLUME_CFG           DVRGetSet = 6370 //设置音频输出口音量调节参数配置
	NET_DVR_INQUEST_PAUSE_CDW                 DVRGetSet = 6371 //暂停刻录
	NET_DVR_INQUEST_RESUME_CDW                DVRGetSet = 6372 //恢复刻录
	NET_DVR_GET_INPUT_CHAN_CFG                DVRGetSet = 6373 //获取输入通道配置
	NET_DVR_SET_INPUT_CHAN_CFG                DVRGetSet = 6374 //设置输入通道配置
	NET_DVR_GET_INQUEST_MIX_AUDIOIN_CFG       DVRGetSet = 6375 //获取审讯机音频输入混音配置
	NET_DVR_SET_INQUEST_MIX_AUDIOIN_CFG       DVRGetSet = 6376 //设置审讯机音频输入混音配置
	NET_DVR_CASE_INFO_CTRL                    DVRGetSet = 6377 //案件信息显示控制
	NET_DVR_GET_INQUEST_USER_RIGHT            DVRGetSet = 6378 //获取审讯机用户权限
	NET_DVR_SET_INQUEST_USER_RIGHT            DVRGetSet = 6379 //设置审讯机用户权限
	NET_DVR_GET_INQUEST_CASE_INFO             DVRGetSet = 6380 //获取审讯案件信息配置
	NET_DVR_SET_INQUEST_CASE_INFO             DVRGetSet = 6381 //设置审讯案件信息配置

	NET_DVR_GET_FILM_MODE_CFG     DVRGetSet = 6387 //获取电影模式
	NET_DVR_SET_FILM_MODE_CFG     DVRGetSet = 6388 //设置电影模式
	NET_DVR_GET_FILM_MODE_CFG_CAP DVRGetSet = 6389 //获取电影模式配置能力

	NET_DVR_GET_DIRECTED_STRATEGY_CFG     DVRGetSet = 6390 //获取导播策略类型
	NET_DVR_SET_DIRECTED_STRATEGY_CFG     DVRGetSet = 6391 //设置导播策略类型
	NET_DVR_GET_DIRECTED_STRATEGY_CFG_CAP DVRGetSet = 6392 //获取电影模式配置能力
	NET_DVR_GET_FRAME_CFG                 DVRGetSet = 6393 //获取画面边框
	NET_DVR_SET_FRAME_CFG                 DVRGetSet = 6394 //设置画面边框
	NET_DVR_GET_FRAME_CFG_CAP             DVRGetSet = 6395 //获取画面边框配置能力
	NET_DVR_GET_AUDIO_EFFECTIVE_CFG       DVRGetSet = 6396 //获取音频优化参数
	NET_DVR_SET_AUDIO_EFFECTIVE_CFG       DVRGetSet = 6397 //设置音频效果参数
	NET_DVR_GET_AUDIO_EFFECTIVE_CFG_CAP   DVRGetSet = 6398 //获取音频效果优化配置能力
	NET_DVR_GET_RECORD_VIDEO_CFG          DVRGetSet = 6399 //获取录制视频参数
	NET_DVR_SET_RECORD_VIDEO_CFG          DVRGetSet = 6400 //设置录制视频参数

	NET_DVR_GET_OUTPUT_CFG      DVRGetSet = 6401 //获取显示输出参数
	NET_DVR_SET_OUTPUT_CFG      DVRGetSet = 6402 //设置显示输出参数
	NET_DVR_CODER_DISPLAY_START DVRGetSet = 6403 //开始输出
	NET_DVR_CODER_DISPLAY_STOP  DVRGetSet = 6404 //停止输出
	NET_DVR_GET_WINDOW_STATUS   DVRGetSet = 6405 //获取显示窗口状态

	//VQD功能接口
	NET_DVR_GET_VQD_LOOP_DIAGNOSE_CFG DVRGetSet = 6406 //获取VQD循环诊断配置参数
	NET_DVR_SET_VQD_LOOP_DIAGNOSE_CFG DVRGetSet = 6407 //设置VQD循环诊断配置参数
	NET_DVR_GET_VQD_DIAGNOSE_INFO     DVRGetSet = 6408 //手动获取VQD诊断信息

	NET_DVR_RECORDING_PUBLISH_FILE              DVRGetSet = 6421 //文件发布
	NET_DVR_GET_RECORDING_PUBLISH_FILE_CAP      DVRGetSet = 6422 //获取文件发布能力
	NET_DVR_GET_PUBLISH_PROGRESS                DVRGetSet = 6423 //获取发布进度
	NET_DVR_GET_RECORD_VIDEO_CFG_CAP            DVRGetSet = 6424 //获取录制视频配置能力
	NET_DVR_GET_RTMP_CFG                        DVRGetSet = 6425 //获取RTMP参数
	NET_DVR_SET_RTMP_CFG                        DVRGetSet = 6426 //设置RTMP参数
	NET_DVR_GET_RTMP_CFG_CAP                    DVRGetSet = 6427 //获取RTMP配置能力
	NET_DVR_DEL_BACKGROUND_PIC                  DVRGetSet = 6428 //删除背景图片文件
	NET_DVR_GET_BACKGROUND_PIC_CFG              DVRGetSet = 6429 //查询背景图片文件
	NET_DVR_GET_BACKGROUND_PIC_INFO             DVRGetSet = 6430 //获取哪张图片作为背景图片
	NET_DVR_SET_BACKGROUND_PIC_INFO             DVRGetSet = 6431 //设置哪张图片作为背景图片
	NET_DVR_GET_BACKGROUND_PIC_INFO_CAP         DVRGetSet = 6432 //获取哪张图片作为背景图片配置能力
	NET_DVR_GET_RECORD_HOST_CAP                 DVRGetSet = 6433 //获取录播主机总能力
	NET_DVR_GET_COURSE_LIST                     DVRGetSet = 6434 //获取课程列表
	NET_DVR_GET_RECORD_STATUS                   DVRGetSet = 6435 //查询录播主机当前状态
	NET_DVR_MANUAL_CURRICULUM_CONTROL           DVRGetSet = 6436 //手动课表控制
	NET_DVR_GET_IMAGE_DIFF_DETECTION_CFG        DVRGetSet = 6437 //获取图像差分检测参数
	NET_DVR_SET_IMAGE_DIFF_DETECTION_CFG        DVRGetSet = 6438 //设置图像差分检测参数
	NET_DVR_GET_IMAGE_DIFF_DETECTION_CFG_CAP    DVRGetSet = 6439 //获取图像差分检测配置能力
	NET_DVR_GET_RECORDING_PUBLISH_FILE_INFO     DVRGetSet = 6440 //获取发布文件信息参数
	NET_DVR_SET_RECORDING_PUBLISH_FILE_INFO     DVRGetSet = 6441 //设置发布文件信息参数
	NET_DVR_GET_RECORDING_PUBLISH_FILE_INFO_CAP DVRGetSet = 6442 //获取发布文件信息配置能力
	NET_DVR_MANUAL_CURRICULUM_CONTROL_CAP       DVRGetSet = 6443 //获取手动课程录像的能力
	NET_DVR_GET_STATISTIC_DATA_LIST             DVRGetSet = 6444 //获取统计数据列表

	NET_DVR_GET_DEVICE_LAN_ENCODE             DVRGetSet = 6501 //获取设备的语言编码
	NET_DVR_GET_GBT28181_SERVICE_CFG          DVRGetSet = 6503 //获取GB28181服务器参数
	NET_DVR_SET_GBT28181_SERVICE_CFG          DVRGetSet = 6504 //设置GB28181服务器参数
	NET_DVR_GET_GBT28181_SERVICE_CAPABILITIES DVRGetSet = 6505 //获取GB28181服务器能力

	NET_DVR_GET_CLOUD_URL                       DVRGetSet = 6506 //获取云存储URL
	NET_DVR_GET_CLOUD_URL_CAP                   DVRGetSet = 6507 //获取云存储URL-能力集
	NET_DVR_GET_CLOUD_CFG                       DVRGetSet = 6508 //获取云存储配置参数
	NET_DVR_SET_CLOUD_CFG                       DVRGetSet = 6509 //设置云存储配置参数
	NET_DVR_GET_CLOUD_CFG_CAP                   DVRGetSet = 6510 //获取云存储配置-能力集
	NET_DVR_GET_CLOUD_UPLOADSTRATEGY            DVRGetSet = 6511 //获取云存储上传策略
	NET_DVR_SET_CLOUD_UPLOADSTRATEGY            DVRGetSet = 6512 //设置云存储上传策略
	NET_DVR_GET_CLOUDSTORAGE_UPLOADSTRATEGY_CAP DVRGetSet = 6513 //云存储上传策略配置-能力集

	NET_DVR_GET_VIDEO_IMAGE_DB_CFG     DVRGetSet = 6601 //获取视图库信息
	NET_DVR_SET_VIDEO_IMAGE_DB_CFG     DVRGetSet = 6602 //设置视图库信息
	NET_DVR_GET_VIDEO_IMAGE_DB_CFG_CAP DVRGetSet = 6603 //获取视图库相关能力
	NET_DVR_GET_FILE_INFO_BY_ID        DVRGetSet = 6604 //根据文件ID获取视图库中文件信息
	NET_DVR_QUERY_FILE_INFO_CAP        DVRGetSet = 6605 //根据文件名查询文件信息能力
	NET_DVR_DEL_FILE_FROM_DB           DVRGetSet = 6606 //从视图库中删除文件
	NET_DVR_GET_VIDEO_IMAGE_DB_CAP     DVRGetSet = 6607 //获取视图库总能力

	NET_DVR_GET_FIGURE DVRGetSet = 6610 //获取缩略图

	NET_DVR_SYNC_IPC_PASSWD               DVRGetSet = 6621 //同步IPC密码与NVR一致
	NET_DVR_GET_VEHICLE_BLACKLST_SCHEDULE DVRGetSet = 6622 //获取黑名单布防时间配置
	NET_DVR_SET_VEHICLE_BLACKLST_SCHEDULE DVRGetSet = 6623 //设置黑名单布防时间配置

	NET_DVR_GET_VEHICLE_WHITELST_SCHEDULE DVRGetSet = 6624 //获取白名单布防时间配置
	NET_DVR_SET_VEHICLE_WHITELST_SCHEDULE DVRGetSet = 6625 //设置白名单布防时间配置

	NET_DVR_GET_VEHICLE_BLACKLIST_EVENT_TRIGGER DVRGetSet = 6626 //获取黑名单布防联动配置
	NET_DVR_SET_VEHICLE_BLACKLIST_EVENT_TRIGGER DVRGetSet = 6627 //设置黑名单布防联动配置

	NET_DVR_GET_VEHICLE_WHITELIST_EVENT_TRIGGER DVRGetSet = 6628 //获取白名单布防联动配置
	NET_DVR_SET_VEHICLE_WHITELIST_EVENT_TRIGGER DVRGetSet = 6629 //设置白名单布防联动配置

	NET_DVR_GET_TRAFFIC_CAP                     DVRGetSet = 6630 //获取抓拍相关能力集
	NET_DVR_GET_VEHICLE_ALLLIST_EVENT_TRIGGER   DVRGetSet = 6631 //获取全部车辆检测布防联动配置
	NET_DVR_SET_VEHICLE_ALLLIST_EVENT_TRIGGER   DVRGetSet = 6632 //设置全部车辆检测布防联动配置
	NET_DVR_GET_VEHICLE_OTHERLIST_EVENT_TRIGGER DVRGetSet = 6633 //获取其他单布防联动配置
	NET_DVR_SET_VEHICLE_OTHERLIST_EVENT_TRIGGER DVRGetSet = 6634 //设置其他单布防联动配置

	NET_DVR_GET_STORAGEDETECTION_EVENT_TRIGGER         DVRGetSet = 6635 //获取存储健康检测联动配置
	NET_DVR_SET_STORAGEDETECTION_EVENT_TRIGGER         DVRGetSet = 6636 //设置存储健康检测联动配置
	NET_DVR_GET_STORAGEDETECTION_SCHEDULE_CAPABILITIES DVRGetSet = 6637 //获取存储健康检测布防时间能力
	NET_DVR_GET_STORAGEDETECTION_SCHEDULE              DVRGetSet = 6638 //获取存储健康布防时间配置
	NET_DVR_SET_STORAGEDETECTION_SCHEDULE              DVRGetSet = 6639 //设置存储健康布防时间配置
	NET_DVR_GET_STORAGEDETECTION_STATE                 DVRGetSet = 6640 //获取存储健康状态

	NET_DVR_GET_STORAGEDETECTION_RWLOCK              DVRGetSet = 6646 //获取存储侦测的读写锁配置
	NET_DVR_GET_STORAGEDETECTION_RWLOCK_CAPABILITIES DVRGetSet = 6647 //获取存储侦测的读写锁配置能力
	NET_DVR_SET_STORAGEDETECTION_RWLOCK              DVRGetSet = 6648 //设置存储侦测的读写锁配置
	NET_DVR_GET_PTZTRACKSTATUS                       DVRGetSet = 6649 //获取球机联动跟踪状态

	NET_DVR_SET_STORAGEDETECTION_UNLOCK              DVRGetSet = 6653 //设置存储侦测的解锁配置
	NET_DVR_GET_STORAGEDETECTION_UNLOCK_CAPABILITIES DVRGetSet = 6654 //获取存储侦测的解锁配置能力

	NET_DVR_SET_SHIPSDETECTION_CFG          DVRGetSet = 6655 //设置船只检测参数配置
	NET_DVR_GET_SHIPSDETECTION_CFG          DVRGetSet = 6656 //获取船只检测参数配置
	NET_DVR_GET_SHIPSDETECTION_CAPABILITIES DVRGetSet = 6657 //获取船只检测参数配置能力
	NET_DVR_GET_SHIPSDETECTION_COUNT        DVRGetSet = 6658 //获取船只计数信息
	NET_DVR_SHIPSCOUNT_DELETE_CTRL          DVRGetSet = 6659 //清空船只计数信息

	NET_DVR_GET_BAREDATAOVERLAY_CAPABILITIES         DVRGetSet = 6660 //获取裸数据叠加能力
	NET_DVR_SET_BAREDATAOVERLAY_CFG                  DVRGetSet = 6661 //设置裸数据叠加
	NET_DVR_GET_BAREDATAOVERLAY_CFG                  DVRGetSet = 6662 //获取裸数据叠加
	NET_DVR_GET_SHIPSDETECTION_SCHEDULE              DVRGetSet = 6663 //获取船只检测布防时间配置
	NET_DVR_SET_SHIPSDETECTION_SCHEDULE              DVRGetSet = 6664 //设置船只检测布防时间配置
	NET_DVR_GET_SHIPSDETECTION_EVENT_TRIGGER         DVRGetSet = 6665 //获取船只检测联动配置
	NET_DVR_SET_SHIPSDETECTION_EVENT_TRIGGER         DVRGetSet = 6666 //设置船只检测联动配置
	NET_DVR_GET_SHIPSDETECTION_SCHEDULE_CAPABILITIES DVRGetSet = 6667 //获取船只检测布防时间能力

	NET_DVR_FIRE_FOCUSZOOM_CTRL DVRGetSet = 6670 //火点可见光镜头聚焦变倍

	NET_DVR_GET_FIREDETECTION_SCHEDULE_CAPABILITIES DVRGetSet = 6671 //获取火点检测布防时间能力
	NET_DVR_GET_FIREDETECTION_SCHEDULE              DVRGetSet = 6672 //获取火点检测布防时间配置
	NET_DVR_SET_FIREDETECTION_SCHEDULE              DVRGetSet = 6673 //设置火点检测布防时间配置
	NET_DVR_GET_MANUALRANGING_CAPABILITIES          DVRGetSet = 6675 //获取手动测距配置能力
	NET_DVR_SET_MANUALRANGING                       DVRGetSet = 6677 //设置手动测距参数
	NET_DVR_GET_MANUALDEICING_CAPABILITIES          DVRGetSet = 6678 //获取手动除冰配置能力
	NET_DVR_SET_MANUALDEICING                       DVRGetSet = 6679 //设置手动除冰
	NET_DVR_GET_MANUALDEICING                       DVRGetSet = 6680 //获取手动除冰

	NET_DVR_GET_THERMALPOWER_CAPABILITIES  DVRGetSet = 6689 //获取相机电源配置能力
	NET_DVR_GET_THERMALPOWER               DVRGetSet = 6690 //获取相机电源配置参数
	NET_DVR_SET_THERMALPOWER               DVRGetSet = 6691 //设置相机电源配置参数
	NET_DVR_GET_PTZABSOLUTEEX_CAPABILITIES DVRGetSet = 6695 //获取高精度PTZ绝对位置配置扩展能力
	NET_DVR_GET_PTZABSOLUTEEX              DVRGetSet = 6696 //获取高精度PTZ绝对位置配置扩展
	NET_DVR_SET_PTZABSOLUTEEX              DVRGetSet = 6697 //设置高精度PTZ绝对位置配置扩展

	NET_DVR_GET_CRUISE_CAPABILITIES    DVRGetSet = 6698 //获取设备巡航模式配置能力
	NET_DVR_GET_CRUISE_INFO            DVRGetSet = 6699 //获取设备巡航模式
	NET_DVR_GET_TEMP_HUMI_CAPABILITIES DVRGetSet = 6700 //温湿度实时能力获取
	NET_DVR_GET_TEMP_HUMI_INFO         DVRGetSet = 6701 //温湿度实时获取

	NET_DVR_GET_MANUALTHERM_INFO         DVRGetSet = 6706 //手动测温实时获取
	NET_DVR_GET_MANUALTHERM_CAPABILITIES DVRGetSet = 6707 //获取手动测温实时数据能力
	NET_DVR_SET_MANUALTHERM              DVRGetSet = 6708 //设置手动测温数据设置

	//DVR96000
	NET_DVR_GET_ACCESSORY_CARD_INFO_CAPABILITIES DVRGetSet = 6709 //获取配件板信息能力
	NET_DVR_GET_ACCESSORY_CARD_INFO              DVRGetSet = 6710 //获取配件板信息

	NET_DVR_GET_THERMINTELL_CAPABILITIES       DVRGetSet = 6711 //获取热成像智能互斥能力
	NET_DVR_GET_THERMINTELL                    DVRGetSet = 6712 //获取热成像智能互斥配置参数
	NET_DVR_SET_THERMINTELL                    DVRGetSet = 6713 //设置热成像智能互斥配置参数
	NET_GET_CRUISEPOINT_V50                    DVRGetSet = 6714 //获取巡航路径配置扩展
	NET_DVR_GET_MANUALTHERM_BASIC_CAPABILITIES DVRGetSet = 6715 //获取手动测温基本参数配置能力
	NET_DVR_SET_MANUALTHERM_BASICPARAM         DVRGetSet = 6716 //设置手动测温基本参数
	NET_DVR_GET_MANUALTHERM_BASICPARAM         DVRGetSet = 6717 //获取手动测温基本参数

	NET_DVR_GET_FIRESHIELDMASK_CAPABILITIES DVRGetSet = 6718 //获取火点区域屏蔽能力

	NET_DVR_GET_HIDDEN_INFORMATION_CAPABILITIES DVRGetSet = 6720 //隐藏信息配置能力
	NET_DVR_GET_HIDDEN_INFORMATION              DVRGetSet = 6721 //获取隐藏信息参数
	NET_DVR_SET_HIDDEN_INFORMATION              DVRGetSet = 6722 //设置隐藏信息参数

	NET_DVR_SET_FIRESHIELDMASK_CFG DVRGetSet = 6723 //设置火点区域屏蔽参数
	NET_DVR_GET_FIRESHIELDMASK_CFG DVRGetSet = 6724 //获取火点区域屏蔽参数

	NET_DVR_GET_SMOKESHIELDMASK_CAPABILITIES DVRGetSet = 6725 //获取烟雾区域屏蔽能力
	NET_DVR_SET_SMOKESHIELDMASK_CFG          DVRGetSet = 6726 //设置烟雾区域屏蔽参数
	NET_DVR_GET_SMOKESHIELDMASK_CFG          DVRGetSet = 6727 //获取烟雾区域屏蔽参数

	NET_DVR_GET_AREASCAN_CAPABILITIES DVRGetSet = 6728 //获取区域扫描能力
	NET_DVR_GET_AREASCAN_CFG          DVRGetSet = 6730 //获取区域扫描参数

	NET_DVR_DEL_AREASCAN_CFG         DVRGetSet = 6731 //扫描区域删除
	NET_DVR_AREASCAN_INIT_CTRL       DVRGetSet = 6732 //进入区域扫描设置
	NET_DVR_AREASCAN_CONFIRM_CTRL    DVRGetSet = 6733 //保存区域扫描设置
	NET_DVR_AREASCAN_STOP_CTRL       DVRGetSet = 6734 //停止区域扫描设置
	NET_DVR_SAVE_SCANZOOM_CTRL       DVRGetSet = 6735 //设置扫描倍率；保存当前光学倍率为扫描倍率
	NET_DVR_GET_SCANZOOM_CTRL        DVRGetSet = 6736 //获取扫描倍率；将预览界面中的光学倍率返回到当前扫描倍率。
	NET_DVR_DEL_FIRESHIELDMASK_CTRL  DVRGetSet = 6737 //删除火点屏蔽区域
	NET_DVR_DEL_SMOKESHIELDMASK_CTRL DVRGetSet = 6738 //删除烟雾屏蔽区域

	NET_DVR_GET_DENSEFOG_EVENT_TRIGGER         DVRGetSet = 6740 //获取大雾检测联动配置
	NET_DVR_SET_DENSEFOG_EVENT_TRIGGER         DVRGetSet = 6741 //设置大雾检测联动配置
	NET_DVR_SET_DENSEFOGDETECTION_CFG          DVRGetSet = 6742 //设置大雾检测参数配置
	NET_DVR_GET_DENSEFOGDETECTION_CFG          DVRGetSet = 6743 //获取大雾检测参数配置
	NET_DVR_GET_DENSEFOGDETECTION_CAPABILITIES DVRGetSet = 6744 //获取大雾检测参数配置能力

	NET_DVR_GET_THERMOMETRY_SCHEDULE_CAPABILITIES DVRGetSet = 6750 //获取测温检测布防时间能力
	NET_DVR_GET_THERMOMETRY_SCHEDULE              DVRGetSet = 6751 //获取测温检测布防时间配置
	NET_DVR_SET_THERMOMETRY_SCHEDULE              DVRGetSet = 6752 //设置测温检测布防时间配置
	NET_DVR_GET_TEMPERTURE_SCHEDULE_CAPABILITIES  DVRGetSet = 6753 //获取温差布防时间能力
	NET_DVR_GET_TEMPERTURE_SCHEDULE               DVRGetSet = 6754 //获取温差布防时间配置
	NET_DVR_SET_TEMPERTURE_SCHEDULE               DVRGetSet = 6755 //设置温差布防时间配置
	NET_DVR_GET_SEARCH_LOG_CAPABILITIES           DVRGetSet = 6756 //日志类型支持能力
	NET_DVR_GET_VEHICLEFLOW                       DVRGetSet = 6758 //获取车流量数据
	NET_DVR_GET_IPADDR_FILTERCFG_V50              DVRGetSet = 6759 //获取IP地址过滤参数扩展
	NET_DVR_SET_IPADDR_FILTERCFG_V50              DVRGetSet = 6760 //设置IP地址过滤参数扩展
	NET_DVR_GET_TEMPHUMSENSOR_CAPABILITIES        DVRGetSet = 6761 //获取温湿度传感器的能力
	NET_DVR_GET_TEMPHUMSENSOR                     DVRGetSet = 6762 //获取温湿度传感器配置协议
	NET_DVR_SET_TEMPHUMSENSOR                     DVRGetSet = 6763 //设置温湿度传感器配置协议

	NET_DVR_GET_THERMOMETRY_MODE_CAPABILITIES DVRGetSet = 6764 //获取测温模式能力
	NET_DVR_GET_THERMOMETRY_MODE              DVRGetSet = 6765 //获取测温模式参数
	NET_DVR_SET_THERMOMETRY_MODE              DVRGetSet = 6766 //设置测温模式参数

	NET_DVR_GET_THERMAL_PIP_CAPABILITIES              DVRGetSet = 6767 //获取热成像画中画配置能力
	NET_DVR_GET_THERMAL_PIP                           DVRGetSet = 6768 //获取热成像画中画配置参数
	NET_DVR_SET_THERMAL_PIP                           DVRGetSet = 6769 //设置热成像画中画配置参数
	NET_DVR_GET_THERMAL_INTELRULEDISPLAY_CAPABILITIES DVRGetSet = 6770 //获取热成像智能规则显示能力
	NET_DVR_GET_THERMAL_INTELRULE_DISPLAY             DVRGetSet = 6771 //获取热成像智能规则显示参数
	NET_DVR_SET_THERMAL_INTELRULE_DISPLAY             DVRGetSet = 6772 //设置热成像智能规则显示参数
	NET_DVR_GET_THERMAL_ALGVERSION                    DVRGetSet = 6773 //获取热成像相关算法库版本
	NET_DVR_GET_CURRENT_LOCK_CAPABILITIES             DVRGetSet = 6774 //获取电流锁定配置能力
	NET_DVR_GET_CURRENT_LOCK                          DVRGetSet = 6775 //获取电流锁定配置参数
	NET_DVR_SET_CURRENT_LOCK                          DVRGetSet = 6776 //设置电流锁定配置参数

	NET_DVR_DEL_MANUALTHERM_RULE DVRGetSet = 6778 //删除手动测温规则

	NET_DVR_GET_UPGRADE_INFO DVRGetSet = 6779 //获取升级信息

	NET_DVR_SWITCH_TRANSFER        DVRGetSet = 7000
	NET_DVR_GET_MB_POWERCTRLPARA   DVRGetSet = 8000 //获取启动控制参数
	NET_DVR_SET_MB_POWERCTRLPARA   DVRGetSet = 8001 //设置启动控制参数
	NET_DVR_GET_AUTOBACKUPPARA     DVRGetSet = 8002 //获取自动备份参数
	NET_DVR_SET_AUTOBACKUPPARA     DVRGetSet = 8003 //设置自动备份参数
	NET_DVR_GET_MB_GPSPARA         DVRGetSet = 8004 //获取GPS参数
	NET_DVR_SET_MB_GPSPARA         DVRGetSet = 8005 //设置GPS参数
	NET_DVR_GET_MB_SENSORINPARA    DVRGetSet = 8006 //获取SENSOR参数
	NET_DVR_SET_MB_SENSORINPARA    DVRGetSet = 8007 //设置SENSOR参数
	NET_DVR_GET_GSENSORPARA        DVRGetSet = 8008 //获取GSENSOR参数
	NET_DVR_SET_GSENSORPARA        DVRGetSet = 8009 //设置GSENSOR参数
	NET_DVR_GET_MB_DOWNLOADSVRPARA DVRGetSet = 8010 //获取下载服务器参数
	NET_DVR_SET_MB_DOWNLOADSVRPARA DVRGetSet = 8011 //设置下载服务器参数
	NET_DVR_GET_PLATERECOG_PARA    DVRGetSet = 8012 //获取车牌识别参数
	NET_DVR_SET_PLATERECOG_PARA    DVRGetSet = 8013 //设置车牌识别参数
	NET_DVR_GET_ENFORCESYS_PARA    DVRGetSet = 8014 //获取车辆稽查参数
	NET_DVR_SET_ENFORCESYS_PARA    DVRGetSet = 8015 //设置车辆稽查参数
	NET_DVR_GET_GPS_DATA           DVRGetSet = 8016 //获取GPS数据
	NET_DVR_GET_ANALOG_ALARMINCFG  DVRGetSet = 8017 //获取模拟报警输入参数
	NET_DVR_SET_ANALOG_ALARMINCFG  DVRGetSet = 8018 //设置模拟报警输入参数

	NET_DVR_GET_SYSTEM_CAPABILITIES   DVRGetSet = 8100 //获取设备的系统能力
	NET_DVR_GET_EAGLEEYE_CAPABILITIES DVRGetSet = 8101 //获取设备鹰眼能力
	NET_DVR_GET_SLAVECAMERA_CALIB_V51 DVRGetSet = 8102 //获取从摄像机标定配置V51
	NET_DVR_SET_SLAVECAMERA_CALIB_V51 DVRGetSet = 8103 //设置从摄像机标定配置V51
	NET_DVR_SET_GOTOSCENE             DVRGetSet = 8105 //设置主摄像机转向指定的场景ID

	//I、K、E系列NVR产品升级
	NET_DVR_GET_PTZ_NOTIFICATION DVRGetSet = 8201 //获取多通道事件联动PTZ
	NET_DVR_SET_PTZ_NOTIFICATION DVRGetSet = 8202 //设置多通道事件联动PTZ
	/*****************************电视墙 start****************************/
	NET_DVR_MATRIX_WALL_SET          DVRGetSet = 9001 //设置电视墙中屏幕参数
	NET_DVR_MATRIX_WALL_GET          DVRGetSet = 9002 //获取电视墙中屏幕参数
	NET_DVR_WALLWIN_GET              DVRGetSet = 9003 //电视墙中获取窗口参数
	NET_DVR_WALLWIN_SET              DVRGetSet = 9004 //电视墙中设置窗口参数
	NET_DVR_WALLWINPARAM_SET         DVRGetSet = 9005 //设置电视墙窗口相关参数
	NET_DVR_WALLWINPARAM_GET         DVRGetSet = 9006 //获取电视墙窗口相关参数
	NET_DVR_WALLSCENEPARAM_GET       DVRGetSet = 9007 //设置场景模式参数
	NET_DVR_WALLSCENEPARAM_SET       DVRGetSet = 9008 //获取场景模式参数
	NET_DVR_MATRIX_GETWINSTATUS      DVRGetSet = 9009 //获取窗口解码状态
	NET_DVR_GET_WINASSOCIATEDDEVINFO DVRGetSet = 9010 //电视墙中获取对应资源信息
	NET_DVR_WALLOUTPUT_GET           DVRGetSet = 9011 //电视墙中获取显示输出参数
	NET_DVR_WALLOUTPUT_SET           DVRGetSet = 9012 //电视墙中设置显示输出参数
	NET_DVR_GET_UNITEDMATRIXSYSTEM   DVRGetSet = 9013 //电视墙中获取对应资源
	NET_DVR_GET_WALL_CFG             DVRGetSet = 9014 //获取电视墙全局参数
	NET_DVR_SET_WALL_CFG             DVRGetSet = 9015 //设置电视墙全局参数
	NET_DVR_CLOSE_ALL_WND            DVRGetSet = 9016 //关闭所有窗口
	NET_DVR_SWITCH_WIN_TOP           DVRGetSet = 9017 //窗口置顶
	NET_DVR_SWITCH_WIN_BOTTOM        DVRGetSet = 9018 //窗口置底

	NET_DVR_CLOSE_ALL_WND_V41        DVRGetSet = 9019 //电视墙关闭所有窗口v41（有多个电视墙）
	NET_DVR_GET_WALL_WINDOW_V41      DVRGetSet = 9020 //获取电视墙中的窗口v41
	NET_DVR_SET_WALL_WINDOW_V41      DVRGetSet = 9021 //设置电视墙中的窗口v41
	NET_DVR_GET_CURRENT_SCENE_V41    DVRGetSet = 9022 //获取当前电视墙中正在使用的场景v41
	NET_DVR_GET_WALL_SCENE_PARAM_V41 DVRGetSet = 9023 //获取当前电视墙中正在使用的场景v41
	NET_DVR_SET_WALL_SCENE_PARAM_V41 DVRGetSet = 9024 //设置当前电视墙中正在使用的场景v41
	NET_DVR_GET_MATRIX_LOGO_CFG      DVRGetSet = 9025 //获取logo参数
	NET_DVR_SET_MATRIX_LOGO_CFG      DVRGetSet = 9026 //设置logo参数
	NET_DVR_GET_WIN_LOGO_CFG         DVRGetSet = 9027 //获取窗口logo参数
	NET_DVR_SET_WIN_LOGO_CFG         DVRGetSet = 9028 //设置窗口logo参数
	NET_DVR_DELETE_LOGO              DVRGetSet = 9029 //删除logo
	NET_DVR_SET_DISPLAY_EFFECT_CFG   DVRGetSet = 9030 //设置显示输出效果参数v41
	NET_DVR_GET_DISPLAY_EFFECT_CFG   DVRGetSet = 9031 //获取显示输出效果参数v41
	NET_DVR_DEC_PLAY_REMOTE_FILE     DVRGetSet = 9032 //解码播放远程文件
	NET_DVR_DEC_PLAY_REMOTE_FILE_V50 DVRGetSet = 9314 //解码播放远程文件V50
	NET_DVR_GET_WIN_ZOOM_STATUS      DVRGetSet = 9033 //获取窗口电子放大状态
	NET_DVR_GET_ALL_MATRIX_LOGOCFG   DVRGetSet = 9034 //获取所有logo参数

	/*****************************电视墙 end******************************/

	/*******************************LCD拼接屏 begin******************************************/
	NET_DVR_SIMULATE_REMOTE_CONTROL  DVRGetSet = 9035 //模拟遥控按键 2013-09-05
	NET_DVR_SET_SCREEN_SIGNAL_CFG    DVRGetSet = 9036 //设置屏幕信号源参数
	NET_DVR_GET_SCREEN_SIGNAL_CFG    DVRGetSet = 9037 //获取屏幕信号源参数
	NET_DVR_SET_SCREEN_SPLICE_CFG    DVRGetSet = 9038 //设置屏幕拼接
	NET_DVR_GET_SCREEN_SPLICE_CFG    DVRGetSet = 9039 //获取屏幕拼接
	NET_DVR_GET_SCREEN_FAN_WORK_MODE DVRGetSet = 9040 //获取风扇工作方式
	NET_DVR_SET_SCREEN_FAN_WORK_MODE DVRGetSet = 9041 //设置风扇工作方式
	NET_DVR_SHOW_SCREEN_WORK_STATUS  DVRGetSet = 9044 //显示屏幕状态
	NET_DVR_GET_VGA_CFG              DVRGetSet = 9045 //获取VGA信号配置
	NET_DVR_SET_VGA_CFG              DVRGetSet = 9046 //设置VGA信号配置
	NET_DVR_GET_SCREEN_MENU_CFG      DVRGetSet = 9048 //获取屏幕菜单配置
	NET_DVR_SET_SCREEN_MENU_CFG      DVRGetSet = 9049 //设置屏幕菜单配置
	NET_DVR_SET_SCREEN_DISPLAY_CFG   DVRGetSet = 9050 //设置显示参数 2013-08-28
	NET_DVR_GET_SCREEN_DISPLAY_CFG   DVRGetSet = 9051 //获取显示参数 2013-08-28

	NET_DVR_SET_FUSION_CFG DVRGetSet = 9052 //设置图像融合参数
	NET_DVR_GET_FUSION_CFG DVRGetSet = 9053 //获取图像融合参数

	NET_DVR_SET_PIP_CFG             DVRGetSet = 9060 //设置画中画参数
	NET_DVR_GET_PIP_CFG             DVRGetSet = 9061 //获取画中画参数
	NET_DVR_SET_DEFOG_LCD           DVRGetSet = 9073 //设置透雾参数
	NET_DVR_GET_DEFOG_LCD           DVRGetSet = 9074 //获取透雾参数
	NET_DVR_SHOW_IP                 DVRGetSet = 9075 //显示IP
	NET_DVR_SCREEN_MAINTENANCE_WALL DVRGetSet = 9076 //屏幕维墙
	NET_DVR_SET_SCREEN_POS          DVRGetSet = 9077 //设置屏幕位置参数
	NET_DVR_GET_SCREEN_POS          DVRGetSet = 9078 //获取屏幕位置参数
	/*******************************LCD拼接屏 end******************************************/

	/*******************************LCD拼接屏V1.2 begin******************************************/
	NET_DVR_SCREEN_INDEX_SET             DVRGetSet = 9079 //屏幕索引相关参数设置
	NET_DVR_SCREEN_INDEX_GET             DVRGetSet = 9080 //屏幕索引相关参数获取
	NET_DVR_SCREEN_SPLICE_SET            DVRGetSet = 9081 //设置屏幕拼接参数
	NET_DVR_SCREEN_SPLICE_GET            DVRGetSet = 9082 //获取屏幕拼接参数
	NET_DVR_SET_SCREEN_PARAM             DVRGetSet = 9083 //设置屏幕相关参数
	NET_DVR_GET_SCREEN_PARAM             DVRGetSet = 9084 //获取屏幕相关参数
	NET_DVR_SET_SWITCH_CFG               DVRGetSet = 9085 //设置定时开关机参数
	NET_DVR_GET_SWITCH_CFG               DVRGetSet = 9086 //获取定时开关机参数
	NET_DVR_SET_POWERON_DELAY_CFG        DVRGetSet = 9087 //设置延时开机参数
	NET_DVR_GET_POWERON_DELAY_CFG        DVRGetSet = 9088 //获取延时开机参数
	NET_DVR_SET_SCREEN_POSITION          DVRGetSet = 9089 //设置屏幕位置参数
	NET_DVR_GET_SCREEN_POSITION          DVRGetSet = 9090 //获取屏幕位置参数
	NET_DVR_SCREEN_SCENE_CONTROL         DVRGetSet = 9091 //屏幕场景控制
	NET_DVR_GET_CURRENT_SCREEN_SCENE     DVRGetSet = 9092 //获取当前屏幕场景号
	NET_DVR_GET_SCREEN_SCENE_PARAM       DVRGetSet = 9093 //获取屏幕场景模式参数
	NET_DVR_SET_SCREEN_SCENE_PARAM       DVRGetSet = 9094 //设置屏幕场景模式参数
	NET_DVR_GET_EXTERNAL_MATRIX_RELATION DVRGetSet = 9095 //获取外接矩阵输入输出关联关系
	NET_DVR_GET_LCD_AUDIO_CFG            DVRGetSet = 9096 //获取LCD屏幕音频参数
	NET_DVR_SET_LCD_AUDIO_CFG            DVRGetSet = 9097 //设置LCD屏幕音频参数
	NET_DVR_GET_LCD_WORK_STATE           DVRGetSet = 9098 //获取LCD屏幕工作状态
	NET_DVR_GET_BOOT_LOGO_CFG            DVRGetSet = 9099 //获取LCD屏幕开机logo显示参数
	NET_DVR_SET_BOOT_LOGO_CFG            DVRGetSet = 9100 //设置LCD屏幕开机logo显示参数

	/*******************************LCD拼接屏V1.2 end ******************************************/
	NET_DVR_GET_STREAM_DST_COMPRESSIONINFO DVRGetSet = 9101 //获取目标压缩参数
	NET_DVR_SET_STREAM_DST_COMPRESSIONINFO DVRGetSet = 9102 //设置目标压缩参数
	NET_DVR_GET_STREAM_TRANS_STATUS        DVRGetSet = 9103 //获取流状态
	NET_DVR_GET_DEVICE_TRANS_STATUS        DVRGetSet = 9104 //获取设备转码状态
	NET_DVR_GET_ALLSTREAM_SRC_INFO         DVRGetSet = 9105 //获取所有流信息
	NET_DVR_GET_BIG_SCREEN_AUDIO           DVRGetSet = 9106 //获取大屏音频信息
	NET_DVR_SET_BIG_SCREEN_AUDIO           DVRGetSet = 9107 //设置大屏音频信息
	NET_DVR_GET_DEV_WORK_MODE              DVRGetSet = 9108 //获取转码设备工作模式
	NET_DVR_SET_DEV_WORK_MODE              DVRGetSet = 9109 //设置转码设备工作模式
	NET_DVR_APPLY_TRANS_CHAN               DVRGetSet = 9110 //按流ID申请转码通道
	NET_DVR_GET_DISPCHAN_CFG               DVRGetSet = 9111 //批量获取显示通道参数
	NET_DVR_SET_DISPCHAN_CFG               DVRGetSet = 9112 //批量设置显示通道参数

	NET_DVR_GET_DEC_CHAN_STATUS   DVRGetSet = 9113 //获取解码通道解码状态
	NET_DVR_GET_DISP_CHAN_STATUS  DVRGetSet = 9114 //获取显示通道状态
	NET_DVR_GET_ALARMIN_STATUS    DVRGetSet = 9115 //获取报警输入状态
	NET_DVR_GET_ALARMOUT_STATUS   DVRGetSet = 9116 //获取报警输出状态
	NET_DVR_GET_AUDIO_CHAN_STATUS DVRGetSet = 9117 //获取语音对讲状态

	NET_DVR_GET_VIDEO_AUDIOIN_CFG DVRGetSet = 9118 //获取视频的音频输入参数
	NET_DVR_SET_VIDEO_AUDIOIN_CFG DVRGetSet = 9119 //设置视频的音频输入参数

	NET_DVR_SET_BASEMAP_CFG         DVRGetSet = 9120 //设置底图参数
	NET_DVR_GET_BASEMAP_CFG         DVRGetSet = 9121 //获取底图参数
	NET_DVR_GET_VIRTUAL_SCREEN_CFG  DVRGetSet = 9122 //获取超高清输入子系统参数
	NET_DVR_SET_VIRTUAL_SCREEN_CFG  DVRGetSet = 9123 //设置超高清输入子系统参数
	NET_DVR_GET_BASEMAP_WIN_CFG     DVRGetSet = 9124 //获取底图窗口参数
	NET_DVR_SET_BASEMAP_WIN_CFG     DVRGetSet = 9125 //设置底图窗口参数
	NET_DVR_DELETE_PICTURE          DVRGetSet = 9126 //删除底图
	NET_DVR_GET_BASEMAP_PIC_INFO    DVRGetSet = 9127 //获取底图图片信息
	NET_DVR_SET_BASEMAP_WIN_CFG_V40 DVRGetSet = 9128 //设置底图窗口参数V40
	NET_DVR_GET_BASEMAP_WIN_CFG_V40 DVRGetSet = 9129 //获取底图窗口参数V40

	NET_DVR_GET_DEC_VCA_CFG DVRGetSet = 9130 //获取解码器智能报警参数
	NET_DVR_SET_DEC_VCA_CFG DVRGetSet = 9131 //设置解码器智能报警参数

	NET_DVR_SET_VS_INPUT_CHAN_INIT_ALL DVRGetSet = 9132 //初始化虚拟屏子板的所有输入通道
	NET_DVR_GET_VS_INPUT_CHAN_INIT_ALL DVRGetSet = 9133 //获取虚拟屏子板的所有输入通道的初始化参数
	NET_DVR_GET_VS_INPUT_CHAN_INIT     DVRGetSet = 9134 //获取虚拟屏输入通道的初始化参数
	NET_DVR_GET_VS_INPUT_CHAN_CFG      DVRGetSet = 9135 //获取虚拟屏输入通道配置参数

	NET_DVR_GET_TERMINAL_CONFERENCE_STATUS DVRGetSet = 9136 //获取终端会议状态
	NET_DVR_GET_TERMINAL_INPUT_CFG_CAP     DVRGetSet = 9137 //获取终端输入参数能力
	NET_DVR_GET_TERMINAL_INPUT_CFG         DVRGetSet = 9138 //获取终端视频会议输入参数
	NET_DVR_SET_TERMINAL_INPUT_CFG         DVRGetSet = 9139 //设置终端视频会议输入参数

	NET_DVR_GET_CONFERENCE_REGION_CAP DVRGetSet = 9140 //获取终端会议区域能力
	NET_DVR_GET_CONFERENCE_REGION     DVRGetSet = 9141 //获取终端会议区域参数
	NET_DVR_SET_CONFERENCE_REGION     DVRGetSet = 9142 //设置终端会议区域参数
	NET_DVR_GET_TERMINAL_CALL_CFG_CAP DVRGetSet = 9143 //获取终端呼叫配置能力
	NET_DVR_GET_TERMINAL_CALL_CFG     DVRGetSet = 9144 //获取终端呼叫参数
	NET_DVR_SET_TERMINAL_CALL_CFG     DVRGetSet = 9145 //设置终端呼叫参数
	NET_DVR_GET_TERMINAL_CTRL_CAP     DVRGetSet = 9146 //获取终端呼叫控制能力
	NET_DVR_TERMINAL_CTRL             DVRGetSet = 9147 //终端呼叫控制
	NET_DVR_GET_CALL_QUERY_CAP        DVRGetSet = 9148 //获取会议查找能力
	NET_DVR_GET_CALLINFO_BY_COND      DVRGetSet = 9149 //按条件查询呼叫记录

	NET_DVR_SET_FUSION_SCALE DVRGetSet = 9150 //设置图像融合规模
	NET_DVR_GET_FUSION_SCALE DVRGetSet = 9151 //获取图像融合规模

	NET_DVR_GET_VCS_CAP DVRGetSet = 9152 //获取MCU能力集

	NET_DVR_GET_TERMINAL_GK_CFG_CAP      DVRGetSet = 9153 //获取终端注册GK能力
	NET_DVR_GET_TERMINAL_GK_CFG          DVRGetSet = 9154 //获取终端注册GK参数
	NET_DVR_SET_TERMINAL_GK_CFG          DVRGetSet = 9155 //设置终端注册GK参数
	NET_DVR_GET_MCU_CONFERENCESEARCH_CAP DVRGetSet = 9156 //获取MCU设备的能力
	NET_DVR_SET_VS_INPUT_CHAN_CFG        DVRGetSet = 9157 //设置虚拟屏输入通道配置参数
	NET_DVR_GET_VS_NETSRC_CFG            DVRGetSet = 9158 //设置虚拟屏网络源参数
	NET_DVR_SET_VS_NETSRC_CFG            DVRGetSet = 9159 //设置虚拟屏网络源参数

	NET_DVR_GET_LLDP_CFG                 DVRGetSet = 9160 //获取LLDP参数
	NET_DVR_SET_LLDP_CFG                 DVRGetSet = 9161 //设置LLDP参数
	NET_DVR_GET_LLDP_CAP                 DVRGetSet = 9162 //获取LLDP能力集
	NET_DVR_GET_FIBER_CONVERT_BASIC_INFO DVRGetSet = 9163 //获取光纤收发器基本信息
	NET_DVR_GET_FIBER_CONVERT_WORK_STATE DVRGetSet = 9164 //获取光纤收发器工作状
	NET_DVR_GET_FIBER_CONVERT_TOPOLOGY   DVRGetSet = 9165 //获取光纤收发器拓扑信息
	NET_DVR_GET_FC_PORT_REMARKS          DVRGetSet = 9166 //获取光纤收发器端口注释参数
	NET_DVR_SET_FC_PORT_REMARKS          DVRGetSet = 9167 //设置光纤收发器端口注释参数
	NET_DVR_GET_PORT_REMARKS_CAP         DVRGetSet = 9168 //获取光纤收发器端口注释能力集

	NET_DVR_GET_MCU_CONFERENCECONTROL_CAP DVRGetSet = 9169 //获取会议控制能力
	NET_DVR_GET_MCU_TERMINALCONTROL_CAP   DVRGetSet = 9170 //获取终端控制能力
	NET_DVR_GET_MCU_TERIMINALGROUP_CAP    DVRGetSet = 9171 //获取终端分组能力
	NET_DVR_GET_MCU_TERMINAL_CAP          DVRGetSet = 9174 //获取终端管理能力
	NET_DVR_GET_MCU_CONFERENCE_CAP        DVRGetSet = 9175 //获取会议能力
	NET_DVR_GET_MCU_GK_CFG_CAP            DVRGetSet = 9176 //获取MCUGK配置能力
	NET_DVR_GET_MCU_GK_SERVER_CAP         DVRGetSet = 9177 //获取MCUGK服务能力

	NET_DVR_GET_EDID_CFG_FILE_INFO      DVRGetSet = 9178 //获取EDID配置文件信息
	NET_DVR_GET_EDID_CFG_FILE_INFO_LIST DVRGetSet = 9179 //获取所有EDID配置文件信息
	NET_DVR_SET_EDID_CFG_FILE_INFO      DVRGetSet = 9180 //设置EDID配置文件信息
	NET_DVR_DEL_EDID_CFG_FILE_INFO      DVRGetSet = 9181 //删除EDID配置文件信息（包括文件）
	NET_DVR_GET_EDID_CFG_FILE_INFO_CAP  DVRGetSet = 9182 //EDID配置文件信息能力集

	NET_DVR_GET_SUBWND_DECODE_OSD     DVRGetSet = 9183 //获取子窗口解码OSD信息
	NET_DVR_GET_SUBWND_DECODE_OSD_ALL DVRGetSet = 9184 //获取所有子窗口解码OSD信息
	NET_DVR_SET_SUBWND_DECODE_OSD     DVRGetSet = 9185 //设置子窗口解码OSD信息
	NET_DVR_GET_SUBWND_DECODE_OSD_CAP DVRGetSet = 9186 //获取子窗口解码OSD信息能力集
	NET_DVR_GET_DECODE_CHANNEL_OSD    DVRGetSet = 9187 //获取解码通道OSD信息
	NET_DVR_SET_DECODE_CHANNEL_OSD    DVRGetSet = 9188 //设置解码通道OSD信息

	NET_DVR_GET_OUTPUT_PIC_INFO         DVRGetSet = 9200 //获取输出口图片参数
	NET_DVR_SET_OUTPUT_PIC_INFO         DVRGetSet = 9201 //设置输出口图片参数
	NET_DVR_GET_OUTPUT_PIC_WIN_CFG      DVRGetSet = 9202 //获取输出口图片窗口参数
	NET_DVR_SET_OUTPUT_PIC_WIN_CFG      DVRGetSet = 9203 //设置输出口图片窗口参数
	NET_DVR_GET_OUTPUT_ALL_PIC_WIN_CFG  DVRGetSet = 9204 //获取输出口所有图片窗口参数
	NET_DVR_DELETE_OUPUT_PIC            DVRGetSet = 9205 //删除输出口图片
	NET_DVR_GET_OUTPUT_OSD_CFG          DVRGetSet = 9206 //获取输出口OSD参数
	NET_DVR_SET_OUTPUT_OSD_CFG          DVRGetSet = 9207 //设置输出口OSD参数
	NET_DVR_GET_OUTPUT_ALL_OSD_CFG      DVRGetSet = 9208 //获取输出口所有OSD参数
	NET_DVR_GET_CHAN_RELATION           DVRGetSet = 9209 //获取编码通道关联资源参数
	NET_DVR_SET_CHAN_RELATION           DVRGetSet = 9210 //设置编码通道关联资源参数
	NET_DVR_GET_ALL_CHAN_RELATION       DVRGetSet = 9211 //获取所有编码通道关联资源参数
	NET_DVR_GET_NS_RING_CFG             DVRGetSet = 9212 //获取光纤板环网配置
	NET_DVR_SET_NS_RING_CFG             DVRGetSet = 9213 //设置光纤板环网配置
	NET_DVR_GET_NS_RING_STATUS          DVRGetSet = 9214 //获取光纤板环网状态
	NET_DVR_GET_OPTICAL_PORT_INFO       DVRGetSet = 9220 //获取光口信息
	NET_DVR_SET_OPTICAL_PORT_INFO       DVRGetSet = 9221 //设置光口信息
	NET_DVR_GET_OPTICAL_CHAN_RELATE_CFG DVRGetSet = 9222 //获取编码通道关联光口输入源参数
	NET_DVR_SET_OPTICAL_CHAN_RELATE_CFG DVRGetSet = 9223 //设置编码通道关联光口输入源参数
	NET_DVR_GET_WIN_ROAM_SWITCH_CFG     DVRGetSet = 9224 //获取解码器窗口漫游开关参数
	NET_DVR_SET_WIN_ROAM_SWITCH_CFG     DVRGetSet = 9225 //设置解码器窗口漫游开关参数
	NET_DVR_START_SCREEN_CRTL           DVRGetSet = 9226 //开始屏幕控制
	NET_DVR_GET_SCREEN_FLIE_LIST        DVRGetSet = 9227 //获取屏幕文件列表
	NET_DVR_GET_SCREEN_FILEINFO         DVRGetSet = 9228 //获取屏幕文件信息参数
	NET_DVR_SET_SCREEN_FILEINFO         DVRGetSet = 9229 //设置屏幕文件信息参数

	/*******************************小间距LED显示屏 begin***************************************/
	NET_DVR_GET_LED_OUTPUT_CFG         DVRGetSet = 9230 //获取发送卡输出参数
	NET_DVR_SET_LED_OUTPUT_CFG         DVRGetSet = 9231 //设置发送卡输出参数
	NET_DVR_GET_LED_OUTPUT_PORT_CFG    DVRGetSet = 9232 //获取LED发送卡输出端口参数
	NET_DVR_SET_LED_OUTPUT_PORT_CFG    DVRGetSet = 9233 //设置LED发送卡输出端口参数
	NET_DVR_GET_LED_DISPLAY_AREA_CFG   DVRGetSet = 9234 //获取LED发送卡显示区域
	NET_DVR_SET_LED_DISPLAY_AREA_CFG   DVRGetSet = 9235 //设置LED发送卡显示区域
	NET_DVR_GET_LED_PORT_CFG           DVRGetSet = 9236 //获取LED发送卡端口参数
	NET_DVR_SET_LED_PORT_CFG           DVRGetSet = 9237 //设置LED发送卡端口参数
	NET_DVR_GET_LED_DISPLAY_CFG        DVRGetSet = 9238 //获取LED发送卡显示参数
	NET_DVR_SET_LED_DISPLAY_CFG        DVRGetSet = 9239 //设置LED发送卡显示参数
	NET_DVR_GET_ALL_LED_PORT_CFG       DVRGetSet = 9240 //获取LED发送卡某个输出对应
	NET_DVR_SAVE_LED_CONFIGURATION     DVRGetSet = 9241 //参数固化
	NET_DVR_GET_LED_TEST_SIGNAL_CFG    DVRGetSet = 9242 //获取LED屏测试信号参数
	NET_DVR_SET_LED_TEST_SIGNAL_CFG    DVRGetSet = 9243 //设置LED屏测试信号参数
	NET_DVR_GET_LED_NOSIGNAL_CFG       DVRGetSet = 9244 //获取LED屏无信号显示模式参数
	NET_DVR_SET_LED_NOSIGNAL_CFG       DVRGetSet = 9245 //设置LED屏无信号显示模式参数
	NET_DVR_GET_LED_INPUT_CFG          DVRGetSet = 9246 //获取LED发送卡输入参数
	NET_DVR_SET_LED_INPUT_CFG          DVRGetSet = 9247 //设置LED发送卡输入参数
	NET_DVR_GET_LED_RECV_GAMMA_CFG     DVRGetSet = 9248 //获取接收卡GAMMA表参数
	NET_DVR_SET_LED_RECV_GAMMA_CFG     DVRGetSet = 9249 //设置接收卡GAMMA表参数
	NET_DVR_GET_LED_RECV_CFG           DVRGetSet = 9250 //获取接收卡基本参数
	NET_DVR_SET_LED_RECV_CFG           DVRGetSet = 9251 //设置接收卡基本参数
	NET_DVR_GET_LED_RECV_ADVANCED_CFG  DVRGetSet = 9252 //获取接收卡高级参数
	NET_DVR_SET_LED_RECV_ADVANCED_CFG  DVRGetSet = 9253 //设置接收卡高级参数
	NET_DVR_GET_LED_SCREEN_DISPLAY_CFG DVRGetSet = 9254 //获取LED屏显示参数
	NET_DVR_SET_LED_SCREEN_DISPLAY_CFG DVRGetSet = 9255 //设置LED屏显示参数
	/*******************************小间距LED显示屏 end*****************************************/

	NET_DVR_GET_INSERTPLAY_PROGRESS DVRGetSet = 9273 //获取插播进度

	NET_DVR_GET_SCREEN_CONFIG     DVRGetSet = 9260 //获取屏幕服务器参数
	NET_DVR_SET_SCREEN_CONFIG     DVRGetSet = 9261 //设置屏幕服务器参数
	NET_DVR_GET_SCREEN_CONFIG_CAP DVRGetSet = 9262 //获取屏幕服务器参数能力集

	NET_DVR_GET_SCHEDULE_PUBLISH_PROGRESS DVRGetSet = 9271 //获取日程发布进度
	NET_DVR_GET_PUBLISH_UPGRADE_PROGRESS  DVRGetSet = 9272 //获取信息发布终端升级进度

	NET_DVR_GET_INPUT_BOARD_CFG      DVRGetSet = 9281 //获取输入板配置信息
	NET_DVR_GET_INPUT_BOARD_CFG_LIST DVRGetSet = 9282 //获取输入板配置信息列表
	NET_DVR_SET_INPUT_BOARD_CFG      DVRGetSet = 9283 //设置输入板配置信息

	NET_DVR_GET_INPUT_SOURCE_TEXT_CAP            DVRGetSet = 9284 //获取输入源字符叠加能力
	NET_DVR_GET_INPUT_SOURCE_TEXT_CFG            DVRGetSet = 9285 //获取输入源字符叠加参数
	NET_DVR_GET_INPUT_SOURCE_TEXT_CFG_LSIT       DVRGetSet = 9286 //获取输入源字符叠加参数列表
	NET_DVR_SET_INPUT_SOURCE_TEXT_CFG            DVRGetSet = 9287 //设置输入源字符叠加参数
	NET_DVR_SET_INPUT_SOURCE_TEXT_CFG_LIST       DVRGetSet = 9288 //设置输入源字符叠加参数列表
	NET_DVR_GET_INPUT_SOURCE_RESOLUTION_CAP      DVRGetSet = 9289 //获取输入源自定义分辨率能力
	NET_DVR_GET_INPUT_SOURCE_RESOLUTION_CFG      DVRGetSet = 9290 //获取输入源自定义分辨率参数
	NET_DVR_GET_INPUT_SOURCE_RESOLUTION_CFG_LIST DVRGetSet = 9291 //获取输入源自定义分辨率列表
	NET_DVR_SET_INPUT_SOURCE_RESOLUTION_CFG      DVRGetSet = 9292 //设置输入源自定义分辨率参数
	NET_DVR_SET_INPUT_SOURCE_RESOLUTION_CFG_LIST DVRGetSet = 9293 //设置输入源自定义分辨率参数
	NET_DVR_GET_LED_AREA_INFO_LIST               DVRGetSet = 9295 //获取LED区域列表

	NET_DVR_GET_DISPINPUT_CFG      DVRGetSet = 9296 //获取显示输入参数
	NET_DVR_GET_DISPINPUT_CFG_LIST DVRGetSet = 9297 //获取所有显示输入参数
	NET_DVR_SET_DISPINPUT_CFG      DVRGetSet = 9298 //设置显示输入参数
	NET_DVR_GET_DISPINPUT_CFG_CAP  DVRGetSet = 9299 //获取显示输入参数能力集
	NET_DVR_GET_CURRENT_VALID_PORT DVRGetSet = 9300 //获取当前有效的,可以连接的端口

	NET_DVR_SET_ONLINE_UPGRADE         DVRGetSet = 9301 //允许在线升级
	NET_DVR_GET_ONLINEUPGRADE_PROGRESS DVRGetSet = 9302 //获取在线升级进度
	NET_DVR_GET_FIRMWARECODE           DVRGetSet = 9303 //获取识别码
	NET_DVR_GET_ONLINEUPGRADE_SERVER   DVRGetSet = 9304 //获取升级服务器状态
	NET_DVR_GET_ONLINEUPGRADE_VERSION  DVRGetSet = 9305 //获取新版本信息
	NET_DVR_GET_RECOMMEN_VERSION       DVRGetSet = 9306 //检测是否推荐升级到此版本
	NET_DVR_GET_ONLINEUPGRADE_ABILITY  DVRGetSet = 9309 //获取在线升级能力集

	NET_DVR_GET_FIBER_CONVERT_BASIC_INFO_V50 DVRGetSet = 9310 //获取远端网管收发器基本信息V50
	NET_DVR_GET_FIBER_CONVERT_WORK_STATE_V50 DVRGetSet = 9311 //获取远端网管收发器工作状态

	NET_SDK_LED_SCREEN_CHECK        DVRGetSet = 9312 //LED屏幕校正
	NET_SDK_GENERATE_OUTPUT_CONTROL DVRGetSet = 9315 //通用扩展输出口模块控制
	NET_SDK_GET_MATRIX_STATUS_V51   DVRGetSet = 9313 /*获取视频综合平台状态*/

	//DS-19D2000-S V2.0升级 报警联动配置参数命令码
	NET_DVR_GET_ALARM_LINKAGE_CFG DVRGetSet = 9316 //获取动环报警联动配置参数
	NET_DVR_SET_ALARM_LINKAGE_CFG DVRGetSet = 9317 //设置动环报警联动配置参数

	NET_DVR_GET_RS485_WORK_MODE         DVRGetSet = 10001 //获取RS485串口工作模式
	NET_DVR_SET_RS485_WORK_MODE         DVRGetSet = 10002 //设置RS485串口工作模式
	NET_DVR_GET_SPLITTER_TRANS_CHAN_CFG DVRGetSet = 10003 //获取码分器透明通道参数
	NET_DVR_SET_SPLITTER_TRANS_CHAN_CFG DVRGetSet = 10004 //设置码分器透明通道参数

	NET_DVR_GET_RS485_PROTOCOL_VERSION  DVRGetSet = 10301 //获取RS485协议库版本信息
	NET_DVR_ALARMHOST_REGISTER_DETECTOR DVRGetSet = 10302 //自动注册探测器

	NET_DVR_GET_SIP_CFG               DVRGetSet = 11001 //IP可视化机获取SIP参数
	NET_DVR_SET_SIP_CFG               DVRGetSet = 11002 //IP可视化机设置SIP参数
	NET_DVR_GET_IP_VIEW_DEVCFG        DVRGetSet = 11003 //获取IP对讲分机配置
	NET_DVR_SET_IP_VIEW_DEVCFG        DVRGetSet = 11004 //设置IP对讲分机配置
	NET_DVR_GET_IP_VIEW_AUDIO_CFG     DVRGetSet = 11005 //获取IP对讲分机音频参数
	NET_DVR_SET_IP_VIEW_AUDIO_CFG     DVRGetSet = 11006 //设置IP对讲分机音频参数
	NET_DVR_GET_IP_VIEW_CALL_CFG      DVRGetSet = 11007 //获取IP对讲分机呼叫参数
	NET_DVR_SET_IP_VIEW_CALL_CFG      DVRGetSet = 11008 //设置IP对讲分机呼叫参数
	NET_DVR_GET_AUDIO_LIMIT_ALARM_CFG DVRGetSet = 11009 //获取声音超限配置参数
	NET_DVR_SET_AUDIO_LIMIT_ALARM_CFG DVRGetSet = 11010 //设置声音超限配置参数
	NET_DVR_GET_BUTTON_DOWN_ALARM_CFG DVRGetSet = 11011 //获取按钮按下告警配置参数
	NET_DVR_SET_BUTTON_DOWN_ALARM_CFG DVRGetSet = 11012 //设置按钮按下告警配置参数

	NET_DVR_GET_ISCSI_CFG DVRGetSet = 11070 // 获取ISCSI存储配置协议
	NET_DVR_SET_ISCSI_CFG DVRGetSet = 11071 // 获取ISCSI存储配置协议

	NET_DVR_GET_SECURITYMODE DVRGetSet = 12004 //获取当前安全模式
	//2013-11-21 获取设备当前的温度和湿度
	NET_DVR_GET_TEMP_HUMI      DVRGetSet = 12005 //2014-02-15 民用IPC自动化测试项目
	NET_DVR_SET_ALARMSOUNDMODE DVRGetSet = 12006 //设置报警声音模式
	NET_DVR_GET_ALARMSOUNDMODE DVRGetSet = 12007 //获取报警声音模式

	NET_DVR_SET_IPDEVICE_ACTIVATED           DVRGetSet = 13000 //通过NVR激活前端设备
	NET_DVR_GET_DIGITAL_CHAN_SECURITY_STATUS DVRGetSet = 13001 //获取数字通道对应设备安全状态
	NET_DVR_GET_ACTIVATE_IPC_ABILITY         DVRGetSet = 13003 //获取NVR激活IPC能力集

	/*******************************楼宇可视对讲机 start***********************************/
	NET_DVR_GET_VIDEO_INTERCOM_DEVICEID_CFG  DVRGetSet = 16001 //获取可视对讲设备编号
	NET_DVR_SET_VIDEO_INTERCOM_DEVICEID_CFG  DVRGetSet = 16002 //设置可视对讲设备编号
	NET_DVR_SET_PRIVILEGE_PASSWORD           DVRGetSet = 16003 //设置权限密码配置信息
	NET_DVR_GET_OPERATION_TIME_CFG           DVRGetSet = 16004 //获取操作时间配置
	NET_DVR_SET_OPERATION_TIME_CFG           DVRGetSet = 16005 //设置操作时间配置
	NET_DVR_GET_VIDEO_INTERCOM_RELATEDEV_CFG DVRGetSet = 16006 //获取关联网络设备参数
	NET_DVR_SET_VIDEO_INTERCOM_RELATEDEV_CFG DVRGetSet = 16007 //设置关联网络设备参数
	NET_DVR_REMOTECONTROL_NOTICE_DATA        DVRGetSet = 16008 //公告信息下发
	NET_DVR_REMOTECONTROL_GATEWAY            DVRGetSet = 16009 //远程开锁
	NET_DVR_REMOTECONTROL_OPERATION_AUTH     DVRGetSet = 16010 //操作权限验证

	NET_DVR_GET_VIDEO_INTERCOM_IOIN_CFG  DVRGetSet = 16016 //获取IO输入参数
	NET_DVR_SET_VIDEO_INTERCOM_IOIN_CFG  DVRGetSet = 16017 //设置IO输入参数
	NET_DVR_GET_VIDEO_INTERCOM_IOOUT_CFG DVRGetSet = 16018 //获取IO输出参数
	NET_DVR_SET_VIDEO_INTERCOM_IOOUT_CFG DVRGetSet = 16019 //设置IO输出参数
	NET_DVR_GET_ELEVATORCONTROL_CFG      DVRGetSet = 16020 //获取梯控器参数
	NET_DVR_SET_ELEVATORCONTROL_CFG      DVRGetSet = 16021 //设置梯控器参数
	NET_DVR_GET_VIDEOINTERCOM_STREAM     DVRGetSet = 16022 //获取可视对讲流通道参数
	NET_DVR_SET_VIDEOINTERCOM_STREAM     DVRGetSet = 16023 //设置可视对讲流通道参数
	NET_DVR_GET_WDR_CFG                  DVRGetSet = 16024 //获取宽动态参数配置
	NET_DVR_SET_WDR_CFG                  DVRGetSet = 16025 //设置宽动态参数配置
	NET_DVR_GET_VIS_DEVINFO              DVRGetSet = 16026 //获取可设备编号信息
	NET_DVR_GET_VIS_REGISTER_INFO        DVRGetSet = 16027 //获取可设备注册的设备信息
	NET_DVR_GET_ELEVATORCONTROL_CFG_V40  DVRGetSet = 16028 //获取梯控器参数-扩展
	NET_DVR_SET_ELEVATORCONTROL_CFG_V40  DVRGetSet = 16029 //设置梯控器参数-扩展
	NET_DVR_GET_CALL_ROOM_CFG            DVRGetSet = 16030 //获取按键呼叫住户配置
	NET_DVR_SET_CALL_ROOM_CFG            DVRGetSet = 16031 //设置按键呼叫住户配置
	NET_DVR_VIDEO_CALL_SIGNAL_PROCESS    DVRGetSet = 16032 //可视话对讲信令处理
	NET_DVR_GET_CALLER_INFO              DVRGetSet = 16033 //获取主叫长号信息
	NET_DVR_GET_CALL_STATUS              DVRGetSet = 16034 //获取通话状态
	NET_DVR_GET_SERVER_DEVICE_INFO       DVRGetSet = 16035 //获取设备列表
	NET_DVR_SET_CALL_SIGNAL              DVRGetSet = 16036 //可视对讲手机端发送信令
	NET_DVR_GET_VIDEO_INTERCOM_ALARM_CFG DVRGetSet = 16037 //获取可视对讲报警事件参数
	NET_DVR_SET_VIDEO_INTERCOM_ALARM_CFG DVRGetSet = 16038 //设置可视对讲报警事件参数
	NET_DVR_GET_RING_LIST                DVRGetSet = 16039 //查询铃音参数列表

	NET_DVR_GET_ROOM_CUSTOM_CFG         DVRGetSet = 16040 //房间自定义获取
	NET_DVR_SET_ROOM_CUSTOM_CFG         DVRGetSet = 16041 //房间自定义设置
	NET_DVR_GET_ELEVATORCONTROL_CFG_V50 DVRGetSet = 16042 //获取梯控器参数V50
	NET_DVR_SET_ELEVATORCONTROL_CFG_V50 DVRGetSet = 16043 //设置梯控器参数V50
	NET_DVR_GET_SIP_CFG_V50             DVRGetSet = 16044 //获取SIP参数V50
	NET_DVR_SET_SIP_CFG_V50             DVRGetSet = 16045 //设置SIP参数V50
	NET_DVR_GET_NOTICE_VIDEO_DATA       DVRGetSet = 16050 //公告视频获取

	/*******************************楼宇可视对讲机 end***********************************/

	NET_DVR_DEBUGINFO_START DVRGetSet = 18000 //网传设备调试信息启动命令
	NET_DVR_AUTO_TEST_START DVRGetSet = 18001 //自动测试长连接获取

	NET_DVR_GET_SELFCHECK_RESULT DVRGetSet = 20000 //获取设备自检结果
	NET_DVR_SET_TEST_COMMAND     DVRGetSet = 20001 //设置测试控制命令
	NET_DVR_SET_TEST_DEVMODULE   DVRGetSet = 20002 //设置测试硬件模块控制命令
	NET_DVR_GET_TEST_DEVMODULE   DVRGetSet = 20003 //获取测试硬件模块控制命令

	NET_DVR_SET_AUTOFOCUS_TEST          DVRGetSet = 20004 //保存自动对焦参数 2013-10-26
	NET_DVR_CHECK_USER_STATUS           DVRGetSet = 20005 //检测用户是否在线
	NET_DVR_GET_TEST_COMMAND            DVRGetSet = 20010 //获取测试控制命令
	NET_DVR_GET_DIAL_SWITCH_CFG         DVRGetSet = 20200 //获取拨码开关信息
	NET_DVR_SET_AGING_TRICK_SCAN        DVRGetSet = 20201 //设置老化前后工具参数
	NET_DVR_GET_ECCENTRIC_CORRECT_STATE DVRGetSet = 20202 //获取获取偏心校正状态

	NET_DVR_GET_THERMOMETRYRULE_TEMPERATURE_INFO DVRGetSet = 23001 //手动获取测温规则温度信息

	NET_DVR_T1_TEST_CMD DVRGetSet = 131073 //当测试命令来用，通过数据区域的文本内容区分具体做什么.数据长度不得大于1024
	//数据区格式为：<T1TestCmd type="0"/>//恢复设备默认参数并关机。

	// 美分定制菜单输出模式外部命令
	NET_DVR_GET_MEMU_OUTPUT_MODE DVRGetSet = 155649 // 获取菜单输出模式
	NET_DVR_SET_MEMU_OUTPUT_MODE DVRGetSet = 155650 // 设置菜单输出模式

	/***************************DS9000新增命令(_V30) end *****************************/

	NET_DVR_GET_DEV_LOGIN_RET_INFO DVRGetSet = 16777200 //设备登陆返回参数

	NET_DVR_GET_TEST_VERSION_HEAD    DVRGetSet = 268435441 //获取测试版本头
	NET_DVR_SET_TEST_VERSION_HEAD    DVRGetSet = 268435442 //设置测试版本头
	NET_DVR_GET_TEST_VERSION_HEAD_V1 DVRGetSet = 268435443 //获取测试版本头-第二版
	NET_DVR_SET_TEST_VERSION_HEAD_V1 DVRGetSet = 268435444 //设置测试版本头-第二版
	NET_DVR_GET_TEST_VERSION_HEAD_V2 DVRGetSet = 268435445 //获取测试版本头-第三版
	NET_DVR_SET_TEST_VERSION_HEAD_V2 DVRGetSet = 268435446 //设置测试版本头-第三版

	NET_DVR_GET_TEST_VERSION_HEAD_ONLY_0 DVRGetSet = 268435447 //获取测试版本头,当前仅有一个版本
	NET_DVR_SET_TEST_VERSION_HEAD_ONLY_0 DVRGetSet = 268435448 //设置测试版本头,当前仅有一个版本
)

const (
	ItcPostIospeedType            DWORD = 0x1      //IO测速（卡口）
	ItcPostSingleioType           DWORD = 0x2      //单IO触发（卡口）
	ItcPostRs485Type              DWORD = 0x4      //RS485车检器触发（卡口）
	ItcPostRs485RadarType         DWORD = 0x8      //RS485雷达触发（卡口）
	ItcPostVirtualcoilType        DWORD = 0x10     //虚拟线圈触发（卡口）
	ItcPostHvtTypeV50             DWORD = 0x20     //混行卡口视频触发V50
	ItcPostMprType                DWORD = 0x40     //多帧识别(卡口)(Ver3.7)
	ItcPostPrsType                DWORD = 0x80     //视频检测触发配置
	ItcEpoliceIoTrafficlightsType DWORD = 0x100    //IO红绿灯（电警）
	ItcEpoliceRs485Type           DWORD = 0x200    //RS485车检器电警触发（电警）
	ItcPostHvtType                DWORD = 0x400    //混行卡口视频触发（卡口）
	ItcPeRs485Type                DWORD = 0x10000  //RS485车检器卡式电警触发（卡式电警）
	ItcVideoEpoliceType           DWORD = 0x20000  //视频电警触发（卡式电警）
	ItcViaVirtualcoilType         DWORD = 0x40000  //VIA触发配置
	ItcPostImtType                DWORD = 0x80000  //智慧监控配置
	IpcPostHvtType                DWORD = 0x100000 //IPC支持的HVT
	ItcPostMobileType             DWORD = 0x200000 //移动交通触发模式
	ItcRedlightPedestrianType     DWORD = 0x400000 //行人闯红灯触发
	ItcNocomityPedestrianType     DWORD = 0x800000 //不礼让行人触发
)

const (
	EXCEPTION_EXCHANGE                 = 0x8000 //用户交互时异常
	EXCEPTION_AUDIOEXCHANGE            = 0x8001 //语音对讲异常
	EXCEPTION_ALARM                    = 0x8002 //报警异常
	EXCEPTION_PREVIEW                  = 0x8003 //网络预览异常
	EXCEPTION_SERIAL                   = 0x8004 //透明通道异常
	EXCEPTION_RECONNECT                = 0x8005 //预览时重连
	EXCEPTION_ALARMRECONNECT           = 0x8006 //报警时重连
	EXCEPTION_SERIALRECONNECT          = 0x8007 //透明通道重连
	SERIAL_RECONNECTSUCCESS            = 0x8008 //透明通道重连成功
	EXCEPTION_PLAYBACK                 = 0x8010 //回放异常
	EXCEPTION_DISKFMT                  = 0x8011 //硬盘格式化
	EXCEPTION_PASSIVEDECODE            = 0x8012 //被动解码异常
	EXCEPTION_EMAILTEST                = 0x8013 //邮件测试异常
	EXCEPTION_BACKUP                   = 0x8014 //备份异常
	PREVIEW_RECONNECTSUCCESS           = 0x8015 //预览时重连成功
	ALARM_RECONNECTSUCCESS             = 0x8016 //报警时重连成功
	RESUME_EXCHANGE                    = 0x8017 //用户交互恢复
	NETWORK_FLOWTEST_EXCEPTION         = 0x8018 //网络流量检测异常
	EXCEPTION_PICPREVIEWRECONNECT      = 0x8019 //图片预览重连
	PICPREVIEW_RECONNECTSUCCESS        = 0x8020 //图片预览重连成功
	EXCEPTION_PICPREVIEW               = 0x8021 //图片预览异常
	EXCEPTION_MAX_ALARM_INFO           = 0x8022 //报警信息缓存已达上限
	EXCEPTION_LOST_ALARM               = 0x8023 //报警丢失
	EXCEPTION_PASSIVETRANSRECONNECT    = 0x8024 //被动转码重连
	PASSIVETRANS_RECONNECTSUCCESS      = 0x8025 //被动转码重连成功
	EXCEPTION_PASSIVETRANS             = 0x8026 //被动转码异常
	SUCCESS_PUSHDEVLOGON               = 0x8030 //推模式设备注册成功
	EXCEPTION_RELOGIN                  = 0x8040 //用户重登陆
	RELOGIN_SUCCESS                    = 0x8041 //用户重登陆成功
	EXCEPTION_PASSIVEDECODE_RECONNNECT = 0x8042 //被动解码重连
	EXCEPTION_CLUSTER_CS_ARMFAILED     = 0x8043 //集群报警异常

	EXCEPTION_RELOGIN_FAILED                  = 0x8044 //重登陆失败，停止重登陆
	EXCEPTION_PREVIEW_RECONNECT_CLOSED        = 0x8045 //关闭预览重连功能
	EXCEPTION_ALARM_RECONNECT_CLOSED          = 0x8046 //关闭报警重连功能
	EXCEPTION_SERIAL_RECONNECT_CLOSED         = 0x8047 //关闭透明通道重连功能
	EXCEPTION_PIC_RECONNECT_CLOSED            = 0x8048 //关闭回显重连功能
	EXCEPTION_PASSIVE_DECODE_RECONNECT_CLOSED = 0x8049 //关闭被动解码重连功能
	EXCEPTION_PASSIVE_TRANS_RECONNECT_CLOSED  = 0x804a //关闭被动转码重连功能
	EXCEPTION_VIDEO_DOWNLOAD                  = 0x804b // [add] by yangzheng 2019/11/09 录像下载异常
)

var excpetionMap = map[int]string{
	EXCEPTION_EXCHANGE:                        "用户交互时异常",
	EXCEPTION_AUDIOEXCHANGE:                   "语音对讲异常",
	EXCEPTION_ALARM:                           "报警异常",
	EXCEPTION_PREVIEW:                         "网络预览异常",
	EXCEPTION_SERIAL:                          "透明通道异常",
	EXCEPTION_RECONNECT:                       "预览时重连",
	EXCEPTION_ALARMRECONNECT:                  "报警时重连",
	EXCEPTION_SERIALRECONNECT:                 "透明通道重连",
	SERIAL_RECONNECTSUCCESS:                   "透明通道重连成功",
	EXCEPTION_PLAYBACK:                        "回放异常",
	EXCEPTION_DISKFMT:                         "硬盘格式化",
	EXCEPTION_PASSIVEDECODE:                   "被动解码异常",
	EXCEPTION_EMAILTEST:                       "邮件测试异常",
	EXCEPTION_BACKUP:                          "备份异常",
	PREVIEW_RECONNECTSUCCESS:                  "预览时重连成功",
	ALARM_RECONNECTSUCCESS:                    "报警时重连成功",
	RESUME_EXCHANGE:                           "用户交互恢复",
	NETWORK_FLOWTEST_EXCEPTION:                "网络流量检测异常",
	EXCEPTION_PICPREVIEWRECONNECT:             "图片预览重连",
	PICPREVIEW_RECONNECTSUCCESS:               "图片预览重连成功",
	EXCEPTION_PICPREVIEW:                      "图片预览异常",
	EXCEPTION_MAX_ALARM_INFO:                  "报警信息缓存已达上限",
	EXCEPTION_LOST_ALARM:                      "报警丢失",
	EXCEPTION_PASSIVETRANSRECONNECT:           "被动转码重连",
	PASSIVETRANS_RECONNECTSUCCESS:             "被动转码重连成功",
	EXCEPTION_PASSIVETRANS:                    "被动转码异常",
	SUCCESS_PUSHDEVLOGON:                      "推模式设备注册成功",
	EXCEPTION_RELOGIN:                         "用户重登陆",
	RELOGIN_SUCCESS:                           "用户重登陆成功",
	EXCEPTION_PASSIVEDECODE_RECONNNECT:        "被动解码重连",
	EXCEPTION_CLUSTER_CS_ARMFAILED:            "集群报警异常",
	EXCEPTION_RELOGIN_FAILED:                  "重登陆失败，停止重登陆",
	EXCEPTION_PREVIEW_RECONNECT_CLOSED:        "关闭预览重连功能",
	EXCEPTION_ALARM_RECONNECT_CLOSED:          "关闭报警重连功能",
	EXCEPTION_SERIAL_RECONNECT_CLOSED:         "关闭透明通道重连功能",
	EXCEPTION_PIC_RECONNECT_CLOSED:            "关闭回显重连功能",
	EXCEPTION_PASSIVE_DECODE_RECONNECT_CLOSED: "关闭被动解码重连功能",
	EXCEPTION_PASSIVE_TRANS_RECONNECT_CLOSED:  "关闭被动转码重连功能",
	EXCEPTION_VIDEO_DOWNLOAD:                  "录像下载异常",
}

func ExceptionString(code int) string {
	return excpetionMap[code]
}
