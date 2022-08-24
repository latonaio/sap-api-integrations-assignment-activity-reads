package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey               string `json:"connection_key"`
	Result                      bool   `json:"result"`
	RedisKey                    string `json:"redis_key"`
	Filepath                    string `json:"filepath"`
	ServiceAssignmentCollection struct {
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
	} `json:"ServiceAssignmentCollection"`
	APISchema              string   `json:"api_schema"`
	Accepter               []string `json:"accepter"`
	AssignmentActivityCode string   `json:"assignment_activity_code"`
	Deleted                bool     `json:"deleted"`
}
