package netsdk

/*
#cgo CFLAGS: -I ../include
#cgo LDFLAGS: -L ../lib -lhcnetsdk

#include <stdio.h>
#include <stdlib.h>
#include "HCNetSDK.h"

extern BOOL export_ExceptionCallBack(DWORD dwType, LONG lUserID, LONG lHandle, void* pUser);
extern void export_MessageCallabck(LONG lCommand, NET_DVR_ALARMER *pAlarmer, char *c, DWORD dwBufLen, void* pUser);
extern void export_RealDataCallback(LONG lPlayHandle, DWORD dwDataType, BYTE *pBuffer, DWORD dwBufSize, void* pUser);

void CALLBACK MessageCallback(LONG lCommand, NET_DVR_ALARMER *pAlarmer, char *c, DWORD dwBufLen, void* pUser) {
	export_MessageCallabck(lCommand, pAlarmer, c, dwBufLen, pUser);
}

void CALLBACK RealDataCallback(LONG lPlayHandle, DWORD dwDataType, BYTE *pBuffer, DWORD dwBufSize, void* pUser) {
	export_RealDataCallback(lPlayHandle, dwDataType, pBuffer, dwBufSize, pUser);
}
*/
import "C"

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/mattn/go-pointer"
)

type (
	ExceptionCallback func(typ int, userdata interface{}) bool
)

func Init() error {
	ret := C.NET_DVR_Init()
	if ret > 0 {
		return Err(int(C.NET_DVR_GetLastError()))
	}
	return nil
}

func Cleanup() error {
	C.NET_DVR_Cleanup()
	return Err(GetLastErrorN())
}

func GetLastErrorN() int {
	return int(C.NET_DVR_GetLastError())
}

func SDKVersion() string {
	v := C.NET_DVR_GetSDKVersion()
	major, minor := v&0xff0000>>16, v&0xffff
	bv := C.NET_DVR_GetSDKBuildVersion()
	return fmt.Sprintf("%d.%d.%d", major, minor, bv)
}

type ExceptionVisitor struct {
	UserData interface{}
	Callback ExceptionCallback
}

func SetExceptionCallback(cb ExceptionCallback, userData interface{}) bool {
	var v = ExceptionVisitor{
		UserData: userData,
		Callback: cb,
	}

	p := pointer.Save(v)
	b := C.NET_DVR_SetExceptionCallBack_V30(0, nil, C.fExceptionCallBack(C.export_ExceptionCallBack), p)
	if b > 0 {
		return true
	}
	return false
}

func SetConnectTime(waitTime time.Duration, retry int) bool {
	b := C.NET_DVR_SetConnectTime(C.DWORD(waitTime/time.Millisecond), C.DWORD(retry))
	return b > 0
}

func SetReconnect(interval time.Duration, enable bool) bool {
	var enb C.BOOL
	if enable {
		enb = 1
	}
	b := C.NET_DVR_SetReconnect(C.DWORD(interval/time.Millisecond), enb)
	return b > 0
}

func SetRecvTimeOut(timeout time.Duration) bool {
	b := C.NET_DVR_SetRecvTimeOut(C.DWORD(timeout / time.Millisecond))
	return b > 0
}

// type AutoRegisterFunc func(clientIp string, clientPort int)

// func StartListen(ip string, port int, cb AutoRegisterFunc) error {
// 	panic("nonimplement")
// }

func Login(addr string, user, pass string) (*Client, error) {
	var (
		cli       Client
		ss        = strings.SplitN(addr, ":", 2)
		ip, sport = ss[0], ss[1]
	)

	port, err := strconv.Atoi(sport)
	if err != nil {
		return nil, err
	}

	handle := int64(C.NET_DVR_Login_V30(C.CString(ip), C.WORD(port), C.CString(user), C.CString(pass),
		(C.LPNET_DVR_DEVICEINFO_V30)(unsafe.Pointer(&cli.DeviceInfo)),
	))
	if handle < 0 {
		return nil, Err(GetLastErrorN())
	}
	clientsLock.Lock()
	clientsMap[handle] = &cli
	clientsLock.Unlock()

	cli.LoginID = handle
	return &cli, nil
}

func (cli *Client) Logout() error {
	C.NET_DVR_Logout_V30(C.LONG(cli.LoginID))
	clientsLock.Lock()
	delete(clientsMap, cli.LoginID)
	clientsLock.Unlock()

	return Err(GetLastErrorN())
}

func (cli *Client) ConfigFile() ([]byte, error) {
	var (
		buf  = make([]byte, 10*1024)
		rlen int
	)

	r := C.NET_DVR_GetConfigFile_V30(C.int(cli.LoginID), (*C.char)(unsafe.Pointer(&buf[0])), (C.DWORD)(len(buf)), (*C.uint)(unsafe.Pointer(&rlen)))
	if r > 0 {
		return buf[:rlen], nil
	}

	return nil, Err(GetLastErrorN())
}

func (cli *Client) DVRConfig(cmd DVRGetSet, channel int) (interface{}, error) {
	return GetDVRConfig(int(cli.LoginID), cmd, channel)
}

type MesasgeCallBackFunc func(cmd CommAlarm, client *Client, alarm *NET_DVR_ALARMER, info interface{})

type MessageCallbackVisitor struct {
	Callback MesasgeCallBackFunc
}

type MessageServer struct {
	HandleID int
}

type RealDataCallbackFunc func(dwDataType DWORD, buf []byte)

type RealDataVisitor struct {
	UserData interface{}
	Callback RealDataCallbackFunc
}

func StartListen(ip string, port int, cb MesasgeCallBackFunc) (*MessageServer, error) {
	var (
		server  MessageServer
		visitor = MessageCallbackVisitor{
			Callback: cb,
		}
		p = pointer.Save(visitor)
	)

	handle := C.NET_DVR_StartListen_V30(C.CString(ip), C.ushort(port), C.MSGCallBack(C.MessageCallback), p)
	server.HandleID = int(handle)
	if err := Err(GetLastErrorN()); err != nil {
		return nil, err
	}
	return &server, nil
}

func (msgSrv *MessageServer) Stop() error {
	ret := C.NET_DVR_StopListen_V30(C.int(msgSrv.HandleID))
	if ret != 0 {
		return Err(GetLastErrorN())
	}

	return nil
}

func SetMessageCallback(cb MesasgeCallBackFunc) error {
	var (
		visitor = MessageCallbackVisitor{
			Callback: cb,
		}
		p = pointer.Save(visitor)
	)

	C.NET_DVR_SetDVRMessageCallBack_V31(C.MSGCallBack(C.MessageCallback), p)
	return Err(GetLastErrorN())
}

func SetSetupAlarmChan(loginId int, param *NET_DVR_SETUPALARM_PARAM) (int, error) {
	ret := C.NET_DVR_SetupAlarmChan_V41(C.int(loginId), (*C.struct_tagNET_DVR_SETUPALARM_PARAM)(unsafe.Pointer(param)))
	if ret < 0 {
		return 0, Err(GetLastErrorN())
	}
	return int(ret), nil
}

func CloseAlarmChan(alarmHandle int) error {
	b := C.NET_DVR_CloseAlarmChan_V30(C.int(alarmHandle))
	if b > 0 {
		return nil
	}
	return Err(GetLastErrorN())
}

func GetDVRConfig(loginId int, cmd DVRGetSet, channel int) (interface{}, error) {

	switch cmd {
	case NetItcGetTriggercfg:
		var (
			cfg NET_ITC_TRIGGERCFG
			l   uint
		)

		b := C.NET_DVR_GetDVRConfig(C.int(loginId), C.uint(cmd), C.int(channel), C.LPVOID(unsafe.Pointer(&cfg)), C.uint(unsafe.Sizeof(cfg)), (*C.uint)(unsafe.Pointer(&l)))
		if b == 0 {
			return nil, Err(GetLastErrorN())
		}
		return &cfg, nil
	case NetDvrGetDevicecfgV40:
		var (
			cfg NET_DVR_DEVICECFG_V40
			l   uint
		)
		b := C.NET_DVR_GetDVRConfig(C.int(loginId), C.uint(cmd), C.int(channel), C.LPVOID(unsafe.Pointer(&cfg)), C.uint(unsafe.Sizeof(cfg)), (*C.uint)(unsafe.Pointer(&l)))
		if b == 0 {
			return nil, Err(GetLastErrorN())
		}
		return &cfg, nil
	case NetDvrGetTrackCfg:
		var (
			cfg NET_DVR_TRACK_CFG
			l   uint
		)
		b := C.NET_DVR_GetDVRConfig(C.int(loginId), C.uint(cmd), C.int(channel), C.LPVOID(unsafe.Pointer(&cfg)), C.uint(unsafe.Sizeof(cfg)), (*C.uint)(unsafe.Pointer(&l)))
		if b == 0 {
			return nil, Err(GetLastErrorN())
		}
		return &cfg, nil

	case NetDvrGetTrackParamcfg:
		var (
			cfg NET_DVR_TRACK_PARAMCFG
			l   uint
		)
		b := C.NET_DVR_GetDVRConfig(C.int(loginId), C.uint(cmd), C.int(channel), C.LPVOID(unsafe.Pointer(&cfg)), C.uint(unsafe.Sizeof(cfg)), (*C.uint)(unsafe.Pointer(&l)))
		if b == 0 {
			return nil, Err(GetLastErrorN())
		}
		return &cfg, nil
	default:
		return nil, errors.New("nonimplement")
	}
}

func SetDvrConfig(loginId int, cmd DVRGetSet, channel int, incfg interface{}) error {
	switch cmd {
	// case NetItcGetTriggercfg:
	// 	var (
	// 		cfg NET_ITC_TRIGGERCFG
	// 		l   uint
	// 	)

	// 	b := C.NET_DVR_GetDVRConfig(C.int(loginId), C.uint(cmd), C.int(channel), C.LPVOID(unsafe.Pointer(&cfg)), C.uint(unsafe.Sizeof(cfg)), (*C.uint)(unsafe.Pointer(&l)))
	// 	if b == 0 {
	// 		return Err(GetLastErrorN())
	// 	}
	// 	return nil
	// case NetDvrGetDevicecfgV40:
	// 	var (
	// 		cfg NET_DVR_DEVICECFG_V40
	// 		l   uint
	// 	)
	// 	b := C.NET_DVR_GetDVRConfig(C.int(loginId), C.uint(cmd), C.int(channel), C.LPVOID(unsafe.Pointer(&cfg)), C.uint(unsafe.Sizeof(cfg)), (*C.uint)(unsafe.Pointer(&l)))
	// 	if b == 0 {
	// 		return Err(GetLastErrorN())
	// 	}
	// 	return nil
	// case NetDvrGetTrackCfg:
	// 	var (
	// 		cfg NET_DVR_TRACK_CFG
	// 		l   uint
	// 	)
	// 	b := C.NET_DVR_GetDVRConfig(C.int(loginId), C.uint(cmd), C.int(channel), C.LPVOID(unsafe.Pointer(&cfg)), C.uint(unsafe.Sizeof(cfg)), (*C.uint)(unsafe.Pointer(&l)))
	// 	if b == 0 {
	// 		return Err(GetLastErrorN())
	// 	}
	// 	return nil

	case NetDvrSetTrackParamcfg:
		cfg, ok := incfg.(*NET_DVR_TRACK_PARAMCFG)
		if !ok {
			return errors.New("invalid in cfg struct")
		}
		b := C.NET_DVR_SetDVRConfig(C.int(loginId), C.uint(cmd), C.int(channel), C.LPVOID(unsafe.Pointer(cfg)), C.uint(unsafe.Sizeof(*cfg)))
		if b == 0 {
			return Err(GetLastErrorN())
		}
		return nil
	default:
		return errors.New("nonimplement")

	}
}

func DvrRealPlayV40(usrId uint64, info *NET_DVR_PREVIEWINFO, fn RealDataCallbackFunc) (*RealPlay, error) {
	var (
		play    = &RealPlay{}
		visitor = RealDataVisitor{
			UserData: play,
			Callback: fn,
		}
		p = pointer.Save(visitor)
	)

	iret := C.NET_DVR_RealPlay_V40(C.LONG(usrId), C.LPNET_DVR_PREVIEWINFO(unsafe.Pointer(info)), C.REALDATACALLBACK(C.RealDataCallback), p)
	if iret < 0 {
		return nil, Err(GetLastErrorN())
	}
	play.realHandle = uint64(iret)
	return play, nil
}

type RealPlay struct {
	realHandle uint64
}

func (play *RealPlay) Stop() error {
	r := C.NET_DVR_StopRealPlay(C.int(play.realHandle))
	if r > 0 {
		return nil
	}

	return Err(GetLastErrorN())
}

func (play *RealPlay) PtzCruise(dwPTZCruiseCmd, byCruiseRoute, byCruisePoint, input int) error {
	r := C.NET_DVR_PTZCruise((C.LONG)(play.realHandle), C.DWORD(dwPTZCruiseCmd), C.BYTE(byCruiseRoute), C.BYTE(byCruisePoint), C.WORD(input))
	if r > 0 {
		return nil
	}
	return Err(GetLastErrorN())
}

func GetPTZCruise(userId int, channel, route int) (*NET_DVR_CRUISE_RET, error) {
	var cruiseRet = &NET_DVR_CRUISE_RET{}
	r := C.NET_DVR_GetPTZCruise(C.LONG(userId), C.LONG(channel), C.LONG(route), C.LPNET_DVR_CRUISE_RET(unsafe.Pointer(cruiseRet)))
	if r > 0 {
		return cruiseRet, nil
	}

	return nil, Err(GetLastErrorN())
}

func GetDeviceAbility(userId int, typ DeviceAbilityKind, in string) ([]byte, error) {
	var (
		buf      = make([]byte, 10*1024)
		rlen int = len(buf)
		r    C.BOOL
	)

	if len(in) == 0 {
		r = C.NET_DVR_GetDeviceAbility(C.LONG(userId), C.uint(typ), nil, C.uint(len(in)), (*C.char)(unsafe.Pointer(&buf[0])), (C.uint(rlen)))
	} else {
		r = C.NET_DVR_GetDeviceAbility(C.LONG(userId), C.uint(typ), C.CString(in), C.uint(len(in)), (*C.char)(unsafe.Pointer(&buf[0])), (C.uint(rlen)))
	}
	if r > 0 {
		return buf[:rlen], nil
	}

	return nil, Err(GetLastErrorN())
}

func GetPtzPosition(userId int, channel int, posID int) (*NET_DVR_PTZPOS, error) {
	var pos NET_DVR_PTZPOS
	r := C.NET_DVR_GetPtzPosition(C.int(userId), C.int(channel), C.int(posID), (*C.NET_DVR_PTZ_POSITION)(unsafe.Pointer(&pos)))
	if r > 0 {
		return &pos, nil
	}

	return nil, Err(GetLastErrorN())
}

func DvrCapture(userId int, snap *NET_DVR_MANUALSNAP) (*NET_DVR_PLATE_RESULT, error) {
	var reslt NET_DVR_PLATE_RESULT

	r := C.NET_DVR_ManualSnap(C.int(userId), (*C.NET_DVR_MANUALSNAP)(unsafe.Pointer(snap)), (*C.NET_DVR_PLATE_RESULT)(unsafe.Pointer(&reslt)))
	if r > 0 {
		return &reslt, nil
	}

	return nil, Err(GetLastErrorN())
}

func ContinuousShoot(userId int, snap *NET_DVR_SNAPCFG) error {
	r := C.NET_DVR_ContinuousShoot(C.int(userId), (*C.NET_DVR_SNAPCFG)(unsafe.Pointer(snap)))
	if r > 0 {
		return nil
	}

	return Err(GetLastErrorN())
}
