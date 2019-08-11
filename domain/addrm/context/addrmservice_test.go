package context

import (
	"testing"
)

func TestAddressManagementService_ApplyIpv4Address(t *testing.T) {
	if DemoApplyIpv4AddressService()!=nil{
		t.Error("Failure")
	}
}
