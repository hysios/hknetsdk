package netsdk

import (
	"unsafe"

	"github.com/yudai/pp"
)

type Command int

type Alarm struct {
}
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

func (cli *Client) Subscribe(kind AlarmKind) error {
	var param = NET_DVR_SETUPALARM_PARAM{
		ST_dwSize: uint32(unsafe.Sizeof(NET_DVR_SETUPALARM_PARAM{})),
	}
	handle, err := SetSetupAlarmChan(int(cli.LoginID), &param)
	if err != nil {
		return err
	}
	pp.Print(param)
	cli.alarmHandle = handle
	return nil
}

func (cli *Client) StopSubscribe() error {
	return CloseAlarmChan(cli.alarmHandle)
}
