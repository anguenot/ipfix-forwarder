package main

import (
	"encoding/json"
	"strconv"
	"github.com/golang/glog"
	"github.com/calmh/ipfix"
)

// golang `map[string]interface{}` to JSON string
func mapToJSON(myMap map[string]interface{}) string {
	jsonBytes, _ := json.Marshal(myMap)
	return string(jsonBytes[:])
}

func parseIpfixMessage(buf []byte, n int,
	ipfixContext *IpfixContext) (map[string]interface{}) {

	msg, err := ipfixContext.session.ParseBuffer(buf[0:n])
	if err != nil {
		glog.Errorln("Error recieved:", err)
	}

	for i := 0; i < len(serverOptions.vendors); i++ {
		switch serverOptions.vendors[i] {
		case VendorVmwareNSX:
			glog.V(4).Infoln("Include vendor fields",
				VendorVmwareNSX)
			includeVmwareNsxFields(ipfixContext.interpreter)
		case VendorVmwareVDS:
			glog.V(4).Infoln("Include vendor fields",
				VendorVmwareVDS)
			includeVmwareVcenterFields(ipfixContext.interpreter)
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

		fieldList = ipfixContext.interpreter.InterpretInto(rec,
			fieldList[:cap(fieldList)])
		for i := 0; i < len(fieldList); i++ {

			if fieldList[i].Name != "" {
				aliasFieldList[fieldList[i].Name] = fieldList[i].Value
				glog.V(3).Infoln("field name=", fieldList[i].Name,
					" field value:", fieldList[i].Value)
			}

			if fieldList[i].Name == Layer2SegmentID {
				nsxSegmentID := getNSXSegmentID(fieldList[i].Value.(uint64))
				aliasFieldList[NSXSegmentID] = strconv.Itoa(int(nsxSegmentID))
			}

		}
	}

	glog.V(3).Infoln("MSG FIELDS MAP:", aliasFieldList)
	return aliasFieldList

}
