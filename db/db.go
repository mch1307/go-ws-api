package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	//"sgithub.sgbt.lu/champam1/go-api/types"
	"sgithub.sgbt.lu/champam1/go-api/pb"
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
		}
	}
	return ret
}

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

func GetAllDevices() (list pb.Devices) {
	return Devices
}

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

func SwitchDevice(id, val int32) (device *pb.Device, err error) {
	dev := GetDeviceByID(id)
	dev.State = val
	return dev, err

}
