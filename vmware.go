package main

import (
	"github.com/calmh/ipfix"
	"encoding/binary"
	"encoding/hex"
	"strconv"
)

const (
	VENDOR_VMWARE_NSX = "vendor-vmware-nsx"
	VENDOR_VMWARE_VDS = "vendor-vmware-vds"
	LAYER2_SEGMENT_ID = "layer2SegmentId"
	NSX_SEGMENT_ID    = "nsxSegmentId"
)

func includeVmwareNsxFields(i *ipfix.Interpreter) {
	vmware_950 := ipfix.DictionaryEntry{
		Name:         "vmware_950",
		FieldID:      950,
		EnterpriseID: 6876,
		Type:         ipfix.Int32,
	}
	i.AddDictionaryEntry(vmware_950)

	vmware_951 := ipfix.DictionaryEntry{
		Name:         "vmUuid",
		FieldID:      951,
		EnterpriseID: 6876,
		Type:         ipfix.Int32,
	}
	i.AddDictionaryEntry(vmware_951)

	vmware_952 := ipfix.DictionaryEntry{
		Name:         "vnicIndex",
		FieldID:      952,
		EnterpriseID: 6876,
		Type:         ipfix.Int32,
	}
	i.AddDictionaryEntry(vmware_952)

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
	vmware_880 := ipfix.DictionaryEntry{
		Name:         "tenantProtocol",
		FieldID:      880,
		EnterpriseID: 6876,
		Type:         ipfix.Int8,
	}
	i.AddDictionaryEntry(vmware_880)

	vmware_881 := ipfix.DictionaryEntry{
		Name:         "tenantSourceIPv4",
		FieldID:      881,
		EnterpriseID: 6876,
		Type:         ipfix.Ipv4Address,
	}
	i.AddDictionaryEntry(vmware_881)

	vmware_882 := ipfix.DictionaryEntry{
		Name:         "tenantDestIPv4",
		FieldID:      882,
		EnterpriseID: 6876,
		Type:         ipfix.Ipv4Address,
	}
	i.AddDictionaryEntry(vmware_882)

	vmware_883 := ipfix.DictionaryEntry{
		Name:         "tenantSourceIPv6",
		FieldID:      883,
		EnterpriseID: 6876,
		Type:         ipfix.String,
	}
	i.AddDictionaryEntry(vmware_883)

	vmware_884 := ipfix.DictionaryEntry{
		Name:         "tenantDestIPv6",
		FieldID:      884,
		EnterpriseID: 6876,
		Type:         ipfix.String,
	}
	i.AddDictionaryEntry(vmware_884)

	vmware_885 := ipfix.DictionaryEntry{
		Name:         "vmware_885",
		FieldID:      885,
		EnterpriseID: 6876,
		Type:         ipfix.String,
	}
	i.AddDictionaryEntry(vmware_885)

	vmware_886 := ipfix.DictionaryEntry{
		Name:         "tenantSourcePort",
		FieldID:      886,
		EnterpriseID: 6876,
		Type:         ipfix.Int16,
	}
	i.AddDictionaryEntry(vmware_886)

	vmware_887 := ipfix.DictionaryEntry{
		Name:         "tenantDestPort",
		FieldID:      887,
		EnterpriseID: 6876,
		Type:         ipfix.Int16,
	}
	i.AddDictionaryEntry(vmware_887)

	vmware_888 := ipfix.DictionaryEntry{
		Name:         "egressInterfaceAttr",
		FieldID:      888,
		EnterpriseID: 6876,
		Type:         ipfix.Int16,
	}
	i.AddDictionaryEntry(vmware_888)

	vmware_889 := ipfix.DictionaryEntry{
		Name:         "vxlanExportRole",
		FieldID:      889,
		EnterpriseID: 6876,
		Type:         ipfix.Int8,
	}
	i.AddDictionaryEntry(vmware_889)

	vmware_890 := ipfix.DictionaryEntry{
		Name:         "ingressInterfaceAttr",
		FieldID:      890,
		EnterpriseID: 6876,
		Type:         ipfix.Int16,
	}
	i.AddDictionaryEntry(vmware_890)

}

// return the NSX segment ID given the value of `layer2SegmentId`
// the last three bytes of this field will give us the segment ID for the flow
func getNSXSegmentId(layer2SegmentId uint64) uint64 {

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, layer2SegmentId)

	str := hex.EncodeToString(b)
	hex_str := "0x" + str[4:6] + str[2:4] + str[0:2]

	segmentId, _ := strconv.ParseInt(hex_str, 0, 64)

	return uint64(segmentId)

}
