/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UpdateEventsSubscResponse201 struct {
	Events                    []AfEventSubscription        `json:"events" bson:"events"`
	NotifUri                  string                       `json:"notifUri,omitempty" bson:"notifUri"`
	UsgThres                  *UsageThreshold              `json:"usgThres,omitempty" bson:"usgThres"`
	AccessType                AccessType                   `json:"accessType,omitempty" bson:"accessType"`
	AnGwAddr                  *AnGwAddress                 `json:"anGwAddr,omitempty" bson:"anGwAddr"`
	EvSubsUri                 string                       `json:"evSubsUri" bson:"evSubsUri"`
	EvNotifs                  []AfEventNotification        `json:"evNotifs" bson:"evNotifs"`
	FailedResourcAllocReports []ResourcesAllocationInfo    `json:"failedResourcAllocReports,omitempty" bson:"failedResourcAllocReports"`
	PlmnId                    *PlmnId                      `json:"plmnId,omitempty" bson:"plmnId"`
	QncReports                []QosNotificationControlInfo `json:"qncReports,omitempty" bson:"qncReports"`
	RatType                   RatType                      `json:"ratType,omitempty" bson:"ratType"`
	UsgRep                    *AccumulatedUsage            `json:"usgRep,omitempty" bson:"usgRep"`
}