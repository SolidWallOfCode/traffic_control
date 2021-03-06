/*
     Copyright 2015 Comcast Cable Communications Management, LLC

     Licensed under the Apache License, Version 2.0 (the "License");
     you may not use this file except in compliance with the License.
     You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

     Unless required by applicable law or agreed to in writing, software
     distributed under the License is distributed on an "AS IS" BASIS,
     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
     See the License for the specific language governing permissions and
     limitations under the License.
 */

package client

import "encoding/json"

type ParamResponse struct {
	Version  string      `json:"version"`
	Response []Parameter `json:"response"`
}

type Parameter struct {
	Name        string `json:"name"`
	ConfigFile  string `json:"configFile"`
	Value       string `json:"Value"`
	LastUpdated string `json:"lastUpdated"`
}

// Parameters
// Get an array of parameter structs for the profile given
func (to *Session) Parameters(profileName string) ([]Parameter, error) {
	body, err := to.getBytes("/api/1.1/parameters/profile/" + profileName + ".json")
	if err != nil {
		return nil, err
	}
	paramList, err := paramUnmarshall(body)
	return paramList.Response, err
}

func paramUnmarshall(body []byte) (ParamResponse, error) {

	var data ParamResponse
	err := json.Unmarshal(body, &data)
	return data, err
}
