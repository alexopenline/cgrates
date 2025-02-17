/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package engine

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/cgrates/cgrates/config"
	"github.com/cgrates/cgrates/utils"
)

func TestExportReqParseFieldDateTimeDaily(t *testing.T) {
	EventReq := NewExportRequest(map[string]utils.DataStorage{}, "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*daily", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "",
	}

	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("*daily", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Day(), expected.Day()) {
		t.Errorf("Expected %v but received %v", expected, result)
	}
}

func TestExportReqParseFieldDateTimeTimeZone(t *testing.T) {
	EventReq := NewExportRequest(map[string]utils.DataStorage{}, "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*daily", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}

	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("*daily", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Day(), expected.Day()) {
		t.Errorf("Expected %v but received %v", finRes.Day(), expected.Day())
	}
}

func TestExportReqParseFieldDateTimeMonthly(t *testing.T) {
	EventReq := NewExportRequest(map[string]utils.DataStorage{}, "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*monthly", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}
	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("*monthly", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Month(), expected.Month()) {
		t.Errorf("Expected %v but received %v", finRes.Month(), expected.Month())
	}
}

func TestExportReqParseFieldDateTimeMonthlyEstimated(t *testing.T) {
	EventReq := NewExportRequest(map[string]utils.DataStorage{}, "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*monthly_estimated", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}
	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("*monthly_estimated", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Month(), expected.Month()) {
		t.Errorf("Expected %v but received %v", finRes.Month(), expected.Month())
	}
}

func TestExportReqParseFieldDateTimeYearly(t *testing.T) {
	EventReq := NewExportRequest(map[string]utils.DataStorage{}, "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*yearly", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}
	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("*yearly", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Year(), expected.Year()) {
		t.Errorf("Expected %v but received %v", finRes.Year(), expected.Year())
	}
}

func TestExportReqParseFieldDateTimeMetaUnlimited(t *testing.T) {
	EventReq := NewExportRequest(map[string]utils.DataStorage{}, "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile(utils.MetaUnlimited, utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}
	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout(utils.MetaUnlimited, utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Day(), expected.Day()) {
		t.Errorf("Expected %v but received %v", finRes.Day(), expected.Day())
	}
}

func TestExportReqParseFieldDateTimeEmpty(t *testing.T) {
	EventReq := NewExportRequest(map[string]utils.DataStorage{}, "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}
	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Day(), expected.Day()) {
		t.Errorf("Expected %v but received %v", finRes.Day(), expected.Day())
	}
}

func TestExportReqParseFieldDateTimeMonthEnd(t *testing.T) {
	EventReq := NewExportRequest(map[string]utils.DataStorage{}, "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*month_endTest", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}
	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("*month_endTest", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Day(), expected.Day()) {
		t.Errorf("Expected %v but received %v", finRes.Day(), expected.Day())
	}
}

func TestExportReqParseFieldDateTimeError(t *testing.T) {
	EventReq := NewExportRequest(map[string]utils.DataStorage{}, "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*month_endTest", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "/",
	}
	_, err := EventReq.ParseField(fctTemp)
	expected := "time: invalid location name"
	if err == nil || err.Error() != expected {
		t.Errorf("Expected <%+v> but received <%+v>", expected, err)
	}
}

func TestExportReqFieldAsINterfaceOnePath(t *testing.T) {
	mS := map[string]utils.DataStorage{
		utils.MetaReq: utils.MapStorage{
			utils.AccountField: "1004",
			utils.Usage:        "20m",
			utils.AnswerTime:   time.Date(2018, time.January, 7, 16, 60, 0, 0, time.UTC),
		},
		utils.MetaOpts: utils.MapStorage{
			utils.APIKey: "attr12345",
		},
		utils.MetaVars: utils.MapStorage{
			utils.RequestType: utils.MetaRated,
			utils.Subsystems:  utils.MetaChargers,
		},
	}
	eventReq := NewExportRequest(mS, "", nil, nil)
	fldPath := []string{utils.MetaReq}
	if val, err := eventReq.FieldAsInterface(fldPath); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(val, mS[utils.MetaReq]) {
		t.Errorf("Expected %+v \n, received %+v", val, mS[utils.MetaReq])
	}

	fldPath = []string{utils.MetaOpts}
	if val, err := eventReq.FieldAsInterface(fldPath); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(val, mS[utils.MetaOpts]) {
		t.Errorf("Expected %+v \n, received %+v", val, mS[utils.MetaOpts])
	}

	fldPath = []string{utils.MetaVars}
	if val, err := eventReq.FieldAsInterface(fldPath); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(val, mS[utils.MetaVars]) {
		t.Errorf("Expected %+v \n, received %+v", val, mS[utils.MetaVars])
	}
}
func TestEventReqFieldAsInterface(t *testing.T) {
	inData := map[string]utils.DataStorage{
		utils.MetaReq: utils.MapStorage{
			"Account": "1001",
			"Usage":   "10m",
		},
	}
	eventReq := NewExportRequest(inData, "cgrates.org", nil, nil)
	fldPath := []string{utils.MetaReq, "Usage"}
	expVal := "10m"
	if rcv, err := eventReq.FieldAsInterface(fldPath); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(rcv, expVal) {
		t.Errorf("Expected %+v \n but received \n %+v", expVal, rcv)
	}

	expVal = "cgrates.org"
	fldPath = []string{utils.MetaTenant}
	if rcv, err := eventReq.FieldAsInterface(fldPath); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(rcv, expVal) {
		t.Errorf("Expected %+v \n but received \n %+v", expVal, rcv)
	}
}

func TestEventReqNewEventExporter(t *testing.T) {
	inData := map[string]utils.DataStorage{
		utils.MetaReq: utils.MapStorage{
			"Account": "1001",
			"Usage":   "10m",
		},
	}
	onm := utils.NewOrderedNavigableMap()
	fullPath := &utils.FullPath{
		PathSlice: []string{utils.MetaReq, utils.MetaTenant},
		Path:      utils.MetaTenant,
	}
	val := &utils.DataLeaf{
		Data: "value1",
	}
	onm.Append(fullPath, val)
	expData := map[string]*utils.OrderedNavigableMap{
		utils.MetaReq: onm,
	}
	expected := &ExportRequest{
		inData:  inData,
		filterS: nil,
		tnt:     "cgrates.org",
		ExpData: expData,
	}
	eventReq := NewExportRequest(inData, "cgrates.org", nil, expData)
	if !reflect.DeepEqual(expected, eventReq) {
		t.Errorf("Expected %v \n but received \n %v", expected, eventReq)
	}
}
