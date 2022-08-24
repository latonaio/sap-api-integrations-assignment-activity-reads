package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-assignment-activity-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToServiceAssignmentCollection(raw []byte, l *logger.Logger) ([]ServiceAssignmentCollection, error) {
	pm := &responses.ServiceAssignmentCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ServiceAssignmentCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	serviceAssignmentCollection := make([]ServiceAssignmentCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		serviceAssignmentCollection = append(serviceAssignmentCollection, ServiceAssignmentCollection{
			ObjectID:                       data.ObjectID,
			ETag:                           data.ETag,
			ProcessingTypeCode:             data.ProcessingTypeCode,
			ProcessingTypeCodeText:         data.ProcessingTypeCodeText,
			CustomerTypeCode:               data.CustomerTypeCode,
			CustomerTypeCodeText:           data.CustomerTypeCodeText,
			Customer:                       data.Customer,
			CustomerID:                     data.CustomerID,
			ServiceTechnicianTypeCode:      data.ServiceTechnicianTypeCode,
			ServiceTechnicianTypeCodeText:  data.ServiceTechnicianTypeCodeText,
			ServiceTechnician:              data.ServiceTechnician,
			ServiceTechnicianID:            data.ServiceTechnicianID,
			TypeCode:                       data.TypeCode,
			TypeCodeText:                   data.TypeCodeText,
			Subject:                        data.Subject,
			LifeCycleStatusCode:            data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:        data.LifeCycleStatusCodeText,
			ID:                             data.ID,
			FullDayIndicator:               data.FullDayIndicator,
			Status:                         data.Status,
			StatusText:                     data.StatusText,
			AssignmentUUID:                 data.AssignmentUUID,
			BusinessActivityUUID:           data.BusinessActivityUUID,
			StartDateTimeZoneCode:          data.StartDateTimeZoneCode,
			StartDateTimeZoneCodeText:      data.StartDateTimeZoneCodeText,
			Priority:                       data.Priority,
			PriorityText:                   data.PriorityText,
			DivisionCode:                   data.DivisionCode,
			DivisionCodeText:               data.DivisionCodeText,
			Released:                       data.Released,
			ReleasedText:                   data.ReleasedText,
			Fixed:                          data.Fixed,
			FixedText:                      data.FixedText,
			PrimaryContactID:               data.PrimaryContactID,
			PrimaryContact:                 data.PrimaryContact,
			PrimaryContactTypeCode:         data.PrimaryContactTypeCode,
			PrimaryContactTypeCodeText:     data.PrimaryContactTypeCodeText,
			CreationDateTime:               data.CreationDateTime,
			CreatedBy:                      data.CreatedBy,
			LastChangedBy:                  data.LastChangedBy,
			OrganizerID:                    data.OrganizerID,
			OrganizerName:                  data.OrganizerName,
			OrganizerTypeCode:              data.OrganizerTypeCode,
			OrganizerTypeCodeText:          data.OrganizerTypeCodeText,
			EndDateTime:                    data.EndDateTime,
			EndDateTimeZoneCode:            data.EndDateTimeZoneCode,
			EndDateTimeZoneCodeText:        data.EndDateTimeZoneCodeText,
			StartDateTime:                  data.StartDateTime,
			AssignmentSchedulingSource:     data.AssignmentSchedulingSource,
			AssignmentSchedulingSourceText: data.AssignmentSchedulingSourceText,
			EntityLastChangedOn:            data.EntityLastChangedOn,
		})
	}

	return serviceAssignmentCollection, nil
}
