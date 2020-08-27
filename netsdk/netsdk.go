package netsdk

/*
#cgo CFLAGS: -I ../include
#cgo LDFLAGS: -L ../lib -lhcnetsdk

#include <stdio.h>
#include <stdlib.h>
#include "HCNetSDK.h"

extern BOOL export_ExceptionCallBack(DWORD dwType, LONG lUserID, LONG lHandle, void* pUser);
extern void export_MessageCallabck(LONG lCommand, NET_DVR_ALARMER *pAlarmer, char *c, DWORD dwBufLen, void* pUser);

void CALLBACK MessageCallback(LONG lCommand, NET_DVR_ALARMER *pAlarmer, char *c, DWORD dwBufLen, void* pUser) {
	export_MessageCallabck(lCommand, pAlarmer, c, dwBufLen, pUser);
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
	return Err(getLastErrorN())
}

func getLastErrorN() int {
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

// func NetDvrGetDeviceAbility(handleID int, abilityKind DeviceAbilityKind) ([]string, error) {
// 	var (
// 		matrix       NET_DVR_MATRIX_ABILITY
// 		matrixLength = unsafe.Sizeof(matrix)
// 	)
// 	l := C.NET_DVR_GetDeviceAbility(
// 		C.LONG(handleID),
// 		C.DWORD(abilityKind),
// 		&matrix, C.DWORD(matrixLength),
// 		&matrix, C.DWORD(matrixLength),
// 	)

// 	return nil, err
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
		return nil, Err(getLastErrorN())
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

	return Err(getLastErrorN())
}

type MesasgeCallBackFunc func(cmd CommAlarm, client *Client, alarm *NET_DVR_ALARMER, info interface{})

type MessageCallbackVisitor struct {
	Callback MesasgeCallBackFunc
}

type MessageServer struct {
	HandleID int
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
	if err := Err(getLastErrorN()); err != nil {
		return nil, err
	}
	return &server, nil
}

func (msgSrv *MessageServer) Stop() error {
	ret := C.NET_DVR_StopListen_V30(C.int(msgSrv.HandleID))
	if ret != 0 {
		return Err(getLastErrorN())
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
	return Err(getLastErrorN())
}

func SetSetupAlarmChan(loginId int, param *NET_DVR_SETUPALARM_PARAM) (int, error) {
	ret := C.NET_DVR_SetupAlarmChan_V41(C.int(loginId), (*C.struct_tagNET_DVR_SETUPALARM_PARAM)(unsafe.Pointer(param)))
	if ret < 0 {
		return 0, Err(getLastErrorN())
	}
	return int(ret), nil
}

func CloseAlarmChan(alarmHandle int) error {
	b := C.NET_DVR_CloseAlarmChan_V30(C.int(alarmHandle))
	if b > 0 {
		return nil
	}
	return Err(getLastErrorN())
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
			return nil, Err(getLastErrorN())
		}
		return &cfg, nil
	default:
		return nil, errors.New("nonimplement")
	}
}
