package netsdk

const (
	MAX_NAMELEN           = 16  //DVR本地登陆名
	MAX_RIGHT             = 32  //设备支持的权限（1-12表示本地权限，13-32表示远程权限）
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
