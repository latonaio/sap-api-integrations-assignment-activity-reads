package sap_api_output_formatter

type AssignmentActivity struct {
	ConnectionKey          string `json:"connection_key"`
	Result                 bool   `json:"result"`
	RedisKey               string `json:"redis_key"`
	Filepath               string `json:"filepath"`
	APISchema              string `json:"api_schema"`
	AssignmentActivityCode string `json:"assignment_activity_code"`
	Deleted                bool   `json:"deleted"`
}

type ServiceAssignmentCollection struct {
	ObjectID                       string `json:"ObjectID"`
	ETag                           string `json:"ETag"`
	ProcessingTypeCode             string `json:"ProcessingTypeCode"`
	ProcessingTypeCodeText         string `json:"ProcessingTypeCodeText"`
	CustomerTypeCode               string `json:"CustomerTypeCode"`
	CustomerTypeCodeText           string `json:"CustomerTypeCodeText"`
	Customer                       string `json:"Customer"`
	CustomerID                     string `json:"CustomerID"`
	ServiceTechnicianTypeCode      string `json:"ServiceTechnicianTypeCode"`
	ServiceTechnicianTypeCodeText  string `json:"ServiceTechnicianTypeCodeText"`
	ServiceTechnician              string `json:"ServiceTechnician"`
	ServiceTechnicianID            string `json:"ServiceTechnicianID"`
	TypeCode                       string `json:"TypeCode"`
	TypeCodeText                   string `json:"TypeCodeText"`
	Subject                        string `json:"Subject"`
	LifeCycleStatusCode            string `json:"LifeCycleStatusCode"`
	LifeCycleStatusCodeText        string `json:"LifeCycleStatusCodeText"`
	ID                             string `json:"ID"`
	FullDayIndicator               bool   `json:"FullDayIndicator"`
	Status                         string `json:"Status"`
	StatusText                     string `json:"StatusText"`
	AssignmentUUID                 string `json:"AssignmentUUID"`
	BusinessActivityUUID           string `json:"BusinessActivityUUID"`
	StartDateTimeZoneCode          string `json:"StartDateTimeZoneCode"`
	StartDateTimeZoneCodeText      string `json:"StartDateTimeZoneCodeText"`
	Priority                       string `json:"Priority"`
	PriorityText                   string `json:"PriorityText"`
	DivisionCode                   string `json:"DivisionCode"`
	DivisionCodeText               string `json:"DivisionCodeText"`
	Released                       string `json:"Released"`
	ReleasedText                   string `json:"ReleasedText"`
	Fixed                          string `json:"Fixed"`
	FixedText                      string `json:"FixedText"`
	PrimaryContactID               string `json:"PrimaryContactID"`
	PrimaryContact                 string `json:"PrimaryContact"`
	PrimaryContactTypeCode         string `json:"PrimaryContactTypeCode"`
	PrimaryContactTypeCodeText     string `json:"PrimaryContactTypeCodeText"`
	CreationDateTime               string `json:"CreationDateTime"`
	CreatedBy                      string `json:"CreatedBy"`
	LastChangedBy                  string `json:"LastChangedBy"`
	OrganizerID                    string `json:"OrganizerID"`
	OrganizerName                  string `json:"OrganizerName"`
	OrganizerTypeCode              string `json:"OrganizerTypeCode"`
	OrganizerTypeCodeText          string `json:"OrganizerTypeCodeText"`
	EndDateTime                    string `json:"EndDateTime"`
	EndDateTimeZoneCode            string `json:"EndDateTimeZoneCode"`
	EndDateTimeZoneCodeText        string `json:"EndDateTimeZoneCodeText"`
	StartDateTime                  string `json:"StartDateTime"`
	AssignmentSchedulingSource     string `json:"AssignmentSchedulingSource"`
	AssignmentSchedulingSourceText string `json:"AssignmentSchedulingSourceText"`
	EntityLastChangedOn            string `json:"EntityLastChangedOn"`
}
