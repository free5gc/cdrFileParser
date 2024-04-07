/*
Copyright © 2024 Hao Li <mr.hao.li@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cdr

type FileHeaderInfo struct {
	FileLength               int    `json:"file_length"`
	HeaderLength             int    `json:"header_length"`
	HighReleaseVersion       string `json:"high_release_version"`
	LowReleaseVersion        string `json:"low_release_version"`
	FileOpeningTimestamp     string `json:"file_opening_timestamp"`
	LastCDRAppendTimestamp   string `json:"last_cdr_append_timestamp"`
	NumberOfCDRsInFile       int    `json:"number_of_cdrs_in_file"`
	FileSequenceNumber       int    `json:"file_sequence_number"`
	FileClosureTriggerReason string `json:"file_closure_trigger_reason"`
	NodeIPAddress            string `json:"node_ip_address"`
	LostCDRIndicator         string `json:"lost_cdr_indicator"`
}

type CdrHeaderInfo struct {
	CdrLength          int    `json:"cdr_length"`
	ReleaseVersion     string `json:"release_version"`
	DataRecorderFormat string `json:"data_record_format"`
	TsNumber           string `json:"ts_number"`
}

type CdrInfo struct {
	NumberOfCDRs int             `json:"number_of_cdrs"`
	CdrHeaders   []CdrHeaderInfo `json:"cdr_headers"`
}

type FileInfo struct {
	HeaderInfo FileHeaderInfo `json:"header_info"`
	CdrInfo    CdrInfo        `json:"cdr_info"`
}
