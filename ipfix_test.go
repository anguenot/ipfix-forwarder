package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMapToJson(t *testing.T) {

	myMap := map[string]interface{}{
		"packetDeltaCount":         "6",
		"flowStartMilliseconds":    "2017-08-16 15:19:57 -0400 EDT",
		"destinationTransportPort": "137",
		"ingressInterface":         "67108865",
		"protocolIdentifier":       "17",
		"ipClassOfService":         "0",
		"maximumTTL":               "128",
		"sourceIPv4Address":        "192.168.103.73",
		"destinationIPv4Address":   "192.168.103.255",
		"flowEndMilliseconds":      "2017-08-16 15:19:59 -0400 EDT",
		"sourceTransportPort":      "137",
		"egressInterface":          "12620",
		"layer2SegmentId":          "72057594037934957",
		"flowDirection":            "1",
		"vmware_888":               "2",
		"paddingOctets":            "[0]",
		"octetDeltaCount":          "468",
		"nsxSegmentId":             "7021",
		"flowEndReason":            "1",
		"tcpControlBits":           "0",
		"vmware_890":               "3",
		"vmware_889":               "1",
	}

	jsonStr := mapToJSON(myMap)

	expected := "{\"destinationIPv4Address\":\"192.168.103.255\"," +
		"\"destinationTransportPort\":\"137\",\"egressInterface\":\"12620\"," +
		"\"flowDirection\":\"1\"," +
		"\"flowEndMilliseconds\":\"2017-08-16 15:19:59 -0400 EDT\"," +
		"\"flowEndReason\":\"1\"," +
		"\"flowStartMilliseconds\":\"2017-08-16 15:19:57 -0400 EDT\"," +
		"\"ingressInterface\":\"67108865\",\"ipClassOfService\":\"0\"," +
		"\"layer2SegmentId\":\"72057594037934957\",\"maximumTTL\":\"128\"," +
		"\"nsxSegmentId\":\"7021\",\"octetDeltaCount\":\"468\"," +
		"\"packetDeltaCount\":\"6\",\"paddingOctets\":\"[0]\"," +
		"\"protocolIdentifier\":\"17\"," +
		"\"sourceIPv4Address\":\"192.168.103.73\"," +
		"\"sourceTransportPort\":\"137\",\"tcpControlBits\":\"0\"," +
		"\"vmware_888\":\"2\",\"vmware_889\":\"1\",\"vmware_890\":\"3\"}"

	if jsonStr != expected {
		t.Error(
			"Expected", expected,
			"got", jsonStr,
		)
	}

}

func TestInitIpfixContext(t *testing.T) {
	c := initIpfixContext()
	assert.NotNil(t, c)
	assert.NotNil(t, c.session)
	assert.NotNil(t, c.interpreter)
}

func TestInitIpfixVendors(t *testing.T) {
	assert.Empty(t, globalServerOptions.vendors)
	c := initIpfixContext()
	initIpfixVendors(c.interpreter)

	globalServerOptions.vendors = []string{VendorVmwareNSX}
	initIpfixVendors(c.interpreter)

	globalServerOptions.vendors = []string{VendorVmwareVDS}
	initIpfixVendors(c.interpreter)

	globalServerOptions.vendors = []string{VendorVmwareVDS, VendorVmwareNSX}
	initIpfixVendors(c.interpreter)

}
