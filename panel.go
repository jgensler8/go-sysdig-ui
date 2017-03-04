/* 
 * Sysdig REST API
 *
 * The Sysdig REST API
 *
 * OpenAPI spec version: 1.0.0
 * Contact: info@sysdig.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

type Panel struct {

	Name string `json:"name,omitempty"`

	ShowAs string `json:"showAs,omitempty"`

	SortAscending bool `json:"sortAscending,omitempty"`

	SingleTimeNavigation bool `json:"singleTimeNavigation,omitempty"`

	InheritTimeNavigation bool `json:"inheritTimeNavigation,omitempty"`

	IsLegendExpanded bool `json:"isLegendExpanded,omitempty"`

	Responsive string `json:"responsive,omitempty"`

	IsNewPanel bool `json:"isNewPanel,omitempty"`

	Time interface{} `json:"time,omitempty"`

	TimeMode interface{} `json:"timeMode,omitempty"`

	NodesNoiseFilter interface{} `json:"nodesNoiseFilter,omitempty"`

	LinksNoiseFilter interface{} `json:"linksNoiseFilter,omitempty"`

	MapDataLimit interface{} `json:"mapDataLimit,omitempty"`

	Metrics []Metric `json:"metrics,omitempty"`

	LinkMetrics []interface{} `json:"linkMetrics,omitempty"`

	CompareWith interface{} `json:"compareWith,omitempty"`

	Format Format `json:"format,omitempty"`

	Sorting []interface{} `json:"sorting,omitempty"`

	FilterProcess bool `json:"filterProcess,omitempty"`

	FilterRoot bool `json:"filterRoot,omitempty"`

	Items []interface{} `json:"items,omitempty"`

	GridConfiguration GridConfiguration `json:"gridConfiguration,omitempty"`
}
