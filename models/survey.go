package models

import "time"

type Survey struct {
	ID               string     `json:"id"`
	RespondentName   string     `json:"respondentName"`
	RespondentEmail  string     `json:"respondentEmail"`
	RespondentPhone  string     `json:"respondentPhone"`
	CompletedAt      *time.Time `json:"completedAt"`
	ResultPercentage int        `json:"resultPercentage"`
}
