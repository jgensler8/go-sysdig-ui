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

type Dashboard struct {

	Id int64 `json:"id,omitempty"`

	Version int32 `json:"version,omitempty"`

	CreatedOn int64 `json:"createdOn,omitempty"`

	ModifiedOn int64 `json:"modifiedOn,omitempty"`

	TeamId int64 `json:"teamId,omitempty"`

	Schema int32 `json:"schema,omitempty"`

	IsLegendExpanded bool `json:"isLegendExpanded,omitempty"`

	FilterExtNodes bool `json:"filterExtNodes,omitempty"`

	MapdDataLimit interface{} `json:"mapdDataLimit,omitempty"`

	LinksNoiseFilder interface{} `json:"linksNoiseFilder,omitempty"`

	GridConfiguration interface{} `json:"gridConfiguration,omitempty"`

	Sorting []interface{} `json:"sorting,omitempty"`

	FilterRoot bool `json:"filterRoot,omitempty"`

	SingleTimeNavigation bool `json:"singleTimeNavigation,omitempty"`

	ShowAs interface{} `json:"showAs,omitempty"`

	CompareWith interface{} `json:"compareWith,omitempty"`

	Format Format `json:"format,omitempty"`

	SortAscending bool `json:"sortAscending,omitempty"`

	ShowAsType interface{} `json:"showAsType,omitempty"`

	NodesNoiseFilter interface{} `json:"nodesNoiseFilter,omitempty"`

	InheritTimeNavigation bool `json:"inheritTimeNavigation,omitempty"`

	Responsive string `json:"responsive,omitempty"`

	Name string `json:"name,omitempty"`

	LinkMetrics []interface{} `json:"linkMetrics,omitempty"`

	IsNewPanel bool `json:"isNewPanel,omitempty"`

	TimeMode TimeMode `json:"timeMode,omitempty"`

	Time Time `json:"time,omitempty"`

	Metrics []interface{} `json:"metrics,omitempty"`

	Items []Panel `json:"items,omitempty"`

	Username string `json:"username,omitempty"`
}
