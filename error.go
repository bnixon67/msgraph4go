/*
Copyright 2020 Bill Nixon

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published
by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package msgraph4go

// InnerError are additional error objects that may be more specific than the top level error.
type InnerError struct {
	RequestId string `json:"request-id,omitempty"`
	Date      string `json:"date,omitempty"`
}

// ODataError contains information about the Error
type ODataError struct {
	// An error code string for the error that occured
	Code string `json:"code,omitempty"`

	// A developer ready message about the error that occured.
	// This should not be displayed to the user directly.
	Message string `json:"message,omitempty"`

	// Optional. Additional error objects that may be more specific than the top level error.
	InnerError *InnerError `json:"innerError,omitempty"`
}

// GraphErrorResponse contains a single property named error.
type GraphErrorResponse struct {
	ODataError *ODataError `json:"error,omitempty"`
}

// Error return a string representation of the error
func (e *GraphErrorResponse) Error() string {
	return VarToJsonString(e.ODataError)
}

// codeIsError return true if the code is a error Status Code per
// https://docs.microsoft.com/en-us/graph/errors?context=graph%2Fapi%2F1.0&view=graph-rest-1.0
func codeIsError(code int) bool {
	// Microsoft Graph error responses and resource types
	// https://docs.microsoft.com/en-us/graph/errors
	var errorCodes = []int{400, 401, 403, 404, 405, 406, 409, 410,
		411, 412, 413, 415, 416, 422, 423, 429, 500, 501, 503,
		504, 507, 509}

	for _, n := range errorCodes {
		if code == n {
			return true
		}
	}
	return false
}
