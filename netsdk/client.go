package netsdk

import (
	"log"
	"sync"
	"unsafe"
)

type Command int

type Alarm struct {
}

type Client struct {
	LoginID     int64
	DeviceInfo  NET_DVR_DEVICEINFO_V30
	alarmHandle int
	msgCb       MesasgeCallBackFunc
}

var (
	clientsMap  = make(map[int64]*Client)
	clientsLock sync.Mutex
)

type MessageFunc func(cmd Command, alarm Alarm, msg string)

func (cli *Client) StartListen(cb MessageFunc) error {
	panic("nonimplement")
}

func (cli *Client) StopListen() error {
	panic("nonimplement")
}

type DeviceAbility int

func (cli *Client) DeviceAbility() ([]DeviceAbility, error) {
	// C.NET_DVR_GetDeviceAbility(LONG lUserID, DWORD dwAbilityType, char* pInBuf, DWORD dwInLength, char* pOutBuf, DWORD dwOutLength)
	panic("nonimplement")
}

type AlarmKind int

func (cli *Client) Subscribe(kind AlarmKind, fn MesasgeCallBackFunc) error {
	var param = NET_DVR_SETUPALARM_PARAM{
		ST_dwSize:            uint32(unsafe.Sizeof(NET_DVR_SETUPALARM_PARAM{})),
		ST_byAlarmInfoType:   1,
		ST_byRetAlarmTypeV40: 1,
	}

	handle, err := SetSetupAlarmChan(int(cli.LoginID), &param)
	if err != nil {
		return err
	}
	log.Printf("kind %v", kind)
	cli.msgCb = fn
	cli.alarmHandle = handle
	return nil
}

func (cli *Client) StopSubscribe() error {
	return CloseAlarmChan(cli.alarmHandle)
}

func (cli *Client) GetDVRConfig(cmd DVRGetSet, channel int) (interface{}, error) {
	return GetDVRConfig(int(cli.LoginID), cmd, channel)
}
