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

func (cli *Client) GetPTZCruise(channel, route int) (*NET_DVR_CRUISE_RET, error) {
	return GetPTZCruise(int(cli.LoginID), channel, route)
}

func (cli *Client) DeviceAbility(typ DeviceAbilityKind, in string) ([]byte, error) {
	return GetDeviceAbility(int(cli.LoginID), typ, in)
}

func (cli *Client) SetDVRConfig(cmd DVRGetSet, channel int, val interface{}) error {
	return SetDvrConfig(int(cli.LoginID), cmd, channel, val)
}

func (cli *Client) GetPtzPosition(channel int, posID int) (*NET_DVR_PTZPOS, error) {
	return GetPtzPosition(int(cli.LoginID), channel, posID)
}

func (cli *Client) Capture(snap *NET_DVR_MANUALSNAP) (*NET_DVR_PLATE_RESULT, error) {
	return DvrCapture(int(cli.LoginID), snap)
}

func (cli *Client) ContinuousShoot(snap *NET_DVR_SNAPCFG) error {
	return ContinuousShoot(int(cli.LoginID), snap)
}
