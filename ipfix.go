package main

import (
	"encoding/json"
	"strconv"
	"github.com/golang/glog"
	"github.com/calmh/ipfix"
	"sync"
)

var (
	ipfixSession         *ipfix.Session
	ipfixInterpreter     *ipfix.Interpreter
	ipfixSessionInitOnce sync.Once
)

// returns and ipfix session and interpreter singletons
func getIpfixSessionAndInterpreter() (*ipfix.Session, *ipfix.Interpreter) {
	ipfixSessionInitOnce.Do(func() {
		ipfixSession = ipfix.NewSession()
		ipfixInterpreter = ipfix.NewInterpreter(ipfixSession)
	})
	return ipfixSession, ipfixInterpreter
}

// golang `map[string]interface{}` to JSON string
func mapToJSON(myMap map[string]interface{}) string {
	jsonBytes, _ := json.Marshal(myMap)
	return string(jsonBytes[:])
}

func parseIpfixMessage(buf []byte, n int) (map[string]interface{}) {

	s, interpreter := getIpfixSessionAndInterpreter()

	msg, err := s.ParseBuffer(buf[0:n])
	if err != nil {
		glog.Errorln("Error recieved:", err)
	}

	for i := 0; i < len(serverOptions.vendors); i++ {
		switch serverOptions.vendors[i] {
		case VendorVmwareNSX:
			glog.V(4).Infoln("Include vendor fields",
				VendorVmwareNSX)
			includeVmwareNsxFields(interpreter)
		case VendorVmwareVDS:
			glog.V(4).Infoln("Include vendor fields",
				VendorVmwareVDS)
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
