package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mch1307/go-ws-api/pb"
)

// Devices stores the device definition
var Devices pb.Devices

//var Device pb.Device
var tmpDevices devicesLoad
var tmpDevice deviceLoad

type devicesLoad struct {
	Items []deviceLoad
}

type deviceLoad struct {
	ID       int    `json:"id"`
	Hardware string `json:"hardware"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Type     string `json:"type"`
	Unit     string `json:"unit"`
	State    int    `json:"state"`
}

func getDeviceType(s string) pb.Device_DeviceType {
	var ret pb.Device_DeviceType
	for v, k := range pb.Device_DeviceType_value {
		if v == s {
			ret = pb.Device_DeviceType(k)
			return ret
		}
	}
	return ret
}

// InitDB initialize the "db" with data.json file
func InitDB() {

	dataFile, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Println("error reading data file: ", err)
	}
	if err := json.Unmarshal(dataFile, &tmpDevices); err != nil {
		log.Println("error parsing data file ", err)
	}
	fmt.Println("db init ok: ", tmpDevices)
	for _, val := range tmpDevices.Items {
		Device := new(pb.Device)
		Device.Id = int32(val.ID)
		Device.Hardware = val.Hardware
		Device.Location = val.Location
		Device.Name = val.Name
		Device.State = int32(val.State)
		Device.Type = getDeviceType(val.Type)
		Devices.Device = append(Devices.Device, Device)
	}
}

// GetAllDevices returns all devices
func GetAllDevices() (list pb.Devices) {
	return Devices
}

// GetDeviceByID returns single device by id
func GetDeviceByID(id int32) *pb.Device {
	device := new(pb.Device)
	for _, res := range Devices.Device {
		log.Println(res)
		if res.Id == id {
			device = res
		}
	}
	return device
}

// SwitchDevice update the state of a given device
func SwitchDevice(id, val int32) (device *pb.Device, err error) {
	found := false
	for _, v := range Devices.Device {
		if v.GetId() == id {
			v.State = val
			device = v
			found = true
		}
	}
	if !found {
		err = errors.New("device not found")
	}
	return device, err

}
