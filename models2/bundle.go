package models2

import (
	"time"
	"encoding/json"
	"github.com/eug48/fhir/models"
)

type ShallowBundle struct {
	ResourceType string                        `json:"resourceType,omitempty"`
	Meta         *models.Meta                  `json:"meta,omitempty"`
	Type         string                        `json:"type,omitempty"`
	Id           string                        `json:"id,omitempty"`
	Total        *uint32                       `json:"total,omitempty"`
	Entry        []ShallowBundleEntryComponent `json:"entry,omitempty"`
	Link         []models.BundleLinkComponent  `json:"link,omitempty"`
}
type ShallowBundleEntryComponent struct {
	Resource *Resource                            `json:"resource,omitempty"`
	FullUrl  string                               `json:"fullUrl,omitempty"`
	Search   *models.BundleEntrySearchComponent   `json:"search,omitempty"`
	Request  *models.BundleEntryRequestComponent  `json:"request,omitempty"`
	Response *models.BundleEntryResponseComponent `json:"response,omitempty"`
}

func (r *ShallowBundle) MarshalJSON() ([]byte, error) {
	r.ResourceType = "Bundle"
	if r.Meta == nil {
		r.Meta = &models.Meta {
			LastUpdated: &models.FHIRDateTime {
				Time: time.Now(),
				Precision: models.Timestamp,
			},
		}
	}
	return json.Marshal(*r)
}

func (b *ShallowBundle) SetTransformReferencesMap(transformReferencesMap map[string]string) {
	for _, e := range b.Entry {
		if e.Resource != nil {
			e.Resource.SetTransformReferencesMap(transformReferencesMap)
		}
	}
}
