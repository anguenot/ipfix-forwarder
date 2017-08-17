package main

import (
	"encoding/json"
	"strconv"
	"github.com/golang/glog"
	"github.com/calmh/ipfix"
)

// golang `map[string]interface{}` to JSON string
func mapToJson(myMap map[string]interface{}) string {
	jsonBytes, _ := json.Marshal(myMap)
	return string(jsonBytes[:])
}

func parseIpfixMessage(s *ipfix.Session, buf []byte, n int) (map[string]interface{}) {

	msg, err := s.ParseBuffer(buf[0:n])
	if err != nil {
		glog.Errorln("Error recieved:", err)
	}

	interpreter := ipfix.NewInterpreter(s)
	for i := 0; i < len(serverOptions.vendors); i++ {
		switch serverOptions.vendors[i] {
		case VENDOR_VMWARE_NSX:
			glog.V(4).Infoln("Include vendor fields",
				VENDOR_VMWARE_NSX)
			includeVmwareNsxFields(interpreter)
		case VENDOR_VMWARE_VDS:
			glog.V(4).Infoln("Include vendor fields",
				VENDOR_VMWARE_VDS)
			includeVmwareVcenterFields(interpreter)
		}
	}

	if len(msg.DataRecords) > 0 {
		glog.V(4).Infoln("msg.DataRecords: ", msg.DataRecords)
	} else {
		glog.V(4).Infoln("msg.DateRecords empty. " +
			"Waiting for schema?")
	}

	var fieldList []ipfix.InterpretedField
	aliasFieldList := make(map[string]interface{})
	for a, rec := range msg.DataRecords {

		glog.V(4).Infoln("Rec: ", rec)
		glog.V(4).Infoln("a: ", a)

		fieldList = interpreter.InterpretInto(rec, fieldList[:cap(fieldList)])
		for i := 0; i < len(fieldList); i++ {

			glog.V(3).Infoln("field name=", fieldList[i].Name,
				" field value:", fieldList[i].Value)

			if fieldList[i].Name != "" {
				aliasFieldList[fieldList[i].Name] = fieldList[i].Value
			}

			if fieldList[i].Name == LAYER2_SEGMENT_ID {
				nsxSegmentId := getNSXSegmentId(fieldList[i].Value.(uint64))
				aliasFieldList[NSX_SEGMENT_ID] = strconv.Itoa(int(nsxSegmentId))
			}

		}
	}

	glog.V(3).Infoln("MSG FIELDS MAP:", aliasFieldList)
	return aliasFieldList

}
