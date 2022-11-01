// Copyright 2022 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model_v2

type storageState string

const (
	STORAGESTATE_UNSPECIFIED storageState = "STORAGESTATE_UNSPECIFIED"
	STORAGESTATE_AVAILABLE   storageState = "STORAGESTATE_AVAILABLE"
	STORAGESTATE_ARCHIVED    storageState = "STORAGESTATE_ARCHIVED"
)

type Experiment struct {
	UUID           string       `gorm:"column:UUID; not null; primary_key"`
	DisplayName    string       `gorm:"column:DisplayName; not null; unique_index:idx_name_namespace"`
	Description    string       `gorm:"column:Description; not null"`
	CreatedAtInSec int64        `gorm:"column:CreatedAtInSec; not null"`
	Namespace      string       `gorm:"column:Namespace; not null; unique_index:idx_name_namespace"`
	StorageState   storageState `gorm:"type:StorageState; not null;"`
}

func (e Experiment) GetValueOfPrimaryKey() string {
	return e.UUID
}

func GetExperimentTablePrimaryKeyColumn() string {
	return "UUID"
}

// PrimaryKeyColumnName returns the primary key for model Experiment.
func (e *Experiment) PrimaryKeyColumnName() string {
	return "UUID"
}

// DefaultSortField returns the default sorting field for model Experiment.
func (e *Experiment) DefaultSortField() string {
	return "CreatedAtInSec"
}

var experimentAPIToModelFieldMap = map[string]string{
	"id":            "UUID",
	"name":          "Name",
	"created_at":    "CreatedAtInSec",
	"description":   "Description",
	"namespace":     "Namespace",
	"storage_state": "StorageState",
}

// APIToModelFieldMap returns a map from API names to field names for model
// Experiment.
func (e *Experiment) APIToModelFieldMap() map[string]string {
	return experimentAPIToModelFieldMap
}

// GetModelName returns table name used as sort field prefix
func (e *Experiment) GetModelName() string {
	return "experiments"
}

func (e *Experiment) GetField(name string) (string, bool) {
	if field, ok := experimentAPIToModelFieldMap[name]; ok {
		return field, true
	}
	return "", false
}

func (e *Experiment) GetFieldValue(name string) interface{} {
	switch name {
	case "UUID":
		return e.UUID
	case "Name":
		return e.Name
	case "CreatedAtInSec":
		return e.CreatedAtInSec
	case "Description":
		return e.Description
	case "Namespace":
		return e.Namespace
	case "StorageState":
		return e.StorageState
	default:
		return nil
	}
}

func (e *Experiment) GetSortByFieldPrefix(name string) string {
	return "experiments."
}

func (e *Experiment) GetKeyFieldPrefix() string {
	return "experiments."
}
