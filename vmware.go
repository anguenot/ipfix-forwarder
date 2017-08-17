package main

import (
	"github.com/calmh/ipfix"
	"encoding/binary"
	"encoding/hex"
	"strconv"
)

// VMware constants
const (
	VendorVmwareNSX = "vendor-vmware-nsx"
	VendorVmwareVDS = "vendor-vmware-vds"
	Layer2SegmentID = "layer2SegmentId"
	NSXSegmentID    = "nsxSegmentId"
)

func includeVmwareNsxFields(i *ipfix.Interpreter) {
	vmware950 := ipfix.DictionaryEntry{
		Name:         "vmware_950",
		FieldID:      950,
		EnterpriseID: 6876,
		Type:         ipfix.Int32,
	}
	i.AddDictionaryEntry(vmware950)

	vmware951 := ipfix.DictionaryEntry{
		Name:         "vmUuid",
		FieldID:      951,
		EnterpriseID: 6876,
		Type:         ipfix.Int32,
	}
	i.AddDictionaryEntry(vmware951)

	vmware952 := ipfix.DictionaryEntry{
		Name:         "vnicIndex",
		FieldID:      952,
		EnterpriseID: 6876,
		Type:         ipfix.Int32,
	}
	i.AddDictionaryEntry(vmware952)

}

// tenant* identifies the tenant or the inner packet attributes.
// the ingress and egress Interface attributes can take the following values
// based on the type of port:
//
// 0X01 – physical port
// 0X02 – VM port
// 0X03 – VXLAN port
//
// the vxlanExportRole defines if the exporter is an ESXi Host or any other
// network device. This should be 0x01 in most cases.
func includeVmwareVcenterFields(i *ipfix.Interpreter) {
	vmware880 := ipfix.DictionaryEntry{
		Name:         "tenantProtocol",
		FieldID:      880,
		EnterpriseID: 6876,
		Type:         ipfix.Int8,
	}
	i.AddDictionaryEntry(vmware880)

	vmware881 := ipfix.DictionaryEntry{
		Name:         "tenantSourceIPv4",
		FieldID:      881,
		EnterpriseID: 6876,
		Type:         ipfix.Ipv4Address,
	}
	i.AddDictionaryEntry(vmware881)

	vmware882 := ipfix.DictionaryEntry{
		Name:         "tenantDestIPv4",
		FieldID:      882,
		EnterpriseID: 6876,
		Type:         ipfix.Ipv4Address,
	}
	i.AddDictionaryEntry(vmware882)

	vmware883 := ipfix.DictionaryEntry{
		Name:         "tenantSourceIPv6",
		FieldID:      883,
		EnterpriseID: 6876,
		Type:         ipfix.String,
	}
	i.AddDictionaryEntry(vmware883)

	vmware884 := ipfix.DictionaryEntry{
		Name:         "tenantDestIPv6",
		FieldID:      884,
		EnterpriseID: 6876,
		Type:         ipfix.String,
	}
	i.AddDictionaryEntry(vmware884)

	vmware885 := ipfix.DictionaryEntry{
		Name:         "vmware_885",
		FieldID:      885,
		EnterpriseID: 6876,
		Type:         ipfix.String,
	}
	i.AddDictionaryEntry(vmware885)

	vmware886 := ipfix.DictionaryEntry{
		Name:         "tenantSourcePort",
		FieldID:      886,
		EnterpriseID: 6876,
		Type:         ipfix.Int16,
	}
	i.AddDictionaryEntry(vmware886)

	vmware887 := ipfix.DictionaryEntry{
		Name:         "tenantDestPort",
		FieldID:      887,
		EnterpriseID: 6876,
		Type:         ipfix.Int16,
	}
	i.AddDictionaryEntry(vmware887)

	vmware888 := ipfix.DictionaryEntry{
		Name:         "egressInterfaceAttr",
		FieldID:      888,
		EnterpriseID: 6876,
		Type:         ipfix.Int16,
	}
	i.AddDictionaryEntry(vmware888)

	vmware889 := ipfix.DictionaryEntry{
		Name:         "vxlanExportRole",
		FieldID:      889,
		EnterpriseID: 6876,
		Type:         ipfix.Int8,
	}
	i.AddDictionaryEntry(vmware889)

	vmware890 := ipfix.DictionaryEntry{
		Name:         "ingressInterfaceAttr",
		FieldID:      890,
		EnterpriseID: 6876,
		Type:         ipfix.Int16,
	}
	i.AddDictionaryEntry(vmware890)

}

// return the NSX segment ID given the value of `layer2SegmentId`
// the last three bytes of this field will give us the segment ID for the flow
func getNSXSegmentID(layer2SegmentID uint64) uint64 {

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, layer2SegmentID)

	str := hex.EncodeToString(b)
	hexStr := "0x" + str[4:6] + str[2:4] + str[0:2]

	segmentID, _ := strconv.ParseInt(hexStr, 0, 64)

	return uint64(segmentID)

}
