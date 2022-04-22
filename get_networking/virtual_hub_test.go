//package get_networking
//
//import (
//	"reflect"
//	"testing"
//)
//
//func TestConnect_Address(t *testing.T) {
//	type fields struct {
//		HubId   string
//		HubType string
//	}
//	type args struct {
//		sw SwitchData
//		t  TokenBuilder
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := Connect{
//				HubId:   tt.fields.HubId,
//				HubType: tt.fields.HubType,
//			}
//			got, err := c.Address(tt.args.sw, tt.args.t)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Address() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("Address() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestConnect_CheckLength(t *testing.T) {
//	type fields struct {
//		HubId   string
//		HubType string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := Connect{
//				HubId:   tt.fields.HubId,
//				HubType: tt.fields.HubType,
//			}
//			if err := c.CheckLength(); (err != nil) != tt.wantErr {
//				t.Errorf("CheckLength() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestConnect_CheckValues(t *testing.T) {
//	type fields struct {
//		HubId   string
//		HubType string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := Connect{
//				HubId:   tt.fields.HubId,
//				HubType: tt.fields.HubType,
//			}
//			if err := c.CheckValues(); (err != nil) != tt.wantErr {
//				t.Errorf("CheckValues() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestConnect_Gen(t *testing.T) {
//	type fields struct {
//		HubId   string
//		HubType string
//	}
//	type args struct {
//		sw SwitchData
//		t  TokenBuilder
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := Connect{
//				HubId:   tt.fields.HubId,
//				HubType: tt.fields.HubType,
//			}
//			got, err := c.Gen(tt.args.sw, tt.args.t)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Gen() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("Gen() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestConnect_Generate(t *testing.T) {
//	type fields struct {
//		HubId   string
//		HubType string
//	}
//	type args struct {
//		j   string
//		con Connect
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := Connect{
//				HubId:   tt.fields.HubId,
//				HubType: tt.fields.HubType,
//			}
//			got, err := c.Generate(tt.args.j, tt.args.con)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("Generate() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestGetVirtualHubConnections(t *testing.T) {
//	type args struct {
//		hubId string
//		t     TokenBuilder
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    HubConnections
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := GetVirtualHubConnections(tt.args.hubId, tt.args.t)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetVirtualHubConnections() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetVirtualHubConnections() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestGetVirtualNetworkPeerings(t *testing.T) {
//	type args struct {
//		hubId string
//		t     TokenBuilder
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    VnetConnections
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := GetVirtualNetworkPeerings(tt.args.hubId, tt.args.t)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetVirtualNetworkPeerings() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetVirtualNetworkPeerings() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestMakeConnectionSwitches(t *testing.T) {
//	type args struct {
//		j string
//		c Connect
//		t TokenBuilder
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    Confirmed
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := MakeConnectionSwitches(tt.args.j, tt.args.c, tt.args.t)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("MakeConnectionSwitches() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("MakeConnectionSwitches() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestParseHubConnections(t *testing.T) {
//	type args struct {
//		hubId string
//		t     TokenBuilder
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    []string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := ParseHubConnections(tt.args.hubId, tt.args.t)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("ParseHubConnections() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ParseHubConnections() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestParseVnetConnections(t *testing.T) {
//	type args struct {
//		hubId string
//		t     TokenBuilder
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    []string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := ParseVnetConnections(tt.args.hubId, tt.args.t)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("ParseVnetConnections() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ParseVnetConnections() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestReturnConnectable(t *testing.T) {
//	type args struct {
//		c Confirmed
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := ReturnConnectable(tt.args.c)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("ReturnConnectable() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("ReturnConnectable() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
