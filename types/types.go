package types

import "github.com/go-openapi/swag"

// Device holds the device definition
type Device struct {
	ID       int    `json:"id"`
	Hardware string `json:"hardware"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Type     string `json:"type"`
	Unit     string `json:"unit"`
	State    int    `json:"state"`
}

func (d *Device) MarshallBinary() ([]byte, error) {
	if d == nil {
		return nil, nil
	}
	return swag.WriteJSON(d)
}

// UnmarshalBinary interface implementation
func (d *Device) UnmarshalBinary(b []byte) error {
	var res Device
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*d = res
	return nil
}

type Devices struct {
	Items []Device
}

func (d *Devices) MarshallBinary() ([]byte, error) {
	if d == nil {
		return nil, nil
	}
	return swag.WriteJSON(d)
}

// UnmarshalBinary interface implementation
func (d *Devices) UnmarshalBinary(b []byte) error {
	var res Devices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*d = res
	return nil
}
