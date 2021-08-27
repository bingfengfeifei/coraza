// Copyright 2021 Juan Pablo Tosso
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

package operators

import (
	"strconv"

	engine "github.com/jptosso/coraza-waf"
)

type GeoLookup struct{}

func (o *GeoLookup) Init(data string) error {
	return nil
}

func (o *GeoLookup) Evaluate(tx *engine.Transaction, value string) bool {
	if tx.Waf.GeoDb == nil {
		return false
	}
	record, err := tx.Waf.GeoDb.Get(value)
	if err != nil {
		return false
	}
	tx.GetCollection(engine.VARIABLE_GEO).SetData(map[string][]string{
		"COUNTRY_CODE":      {record.IsoCode},
		"COUNTRY_NAME":      {record.CountryName},
		"COUNTRY_CONTINENT": {record.Continent},
		"REGION":            {record.Region},
		"CITY":              {record.City},
		"POSTAL_CODE":       {record.PostalCode},
		"LATITUDE":          {strconv.FormatFloat(record.Latitude, 'f', -1, 64)},
		"LONGITUDE":         {strconv.FormatFloat(record.Longitude, 'f', -1, 64)},
	})
	return true
}
