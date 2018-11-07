// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/seizadi/cmdb/pkg/pb/cmdb.proto

package pb // import "github.com/seizadi/cmdb/pkg/pb"

import fmt "fmt"
import http "net/http"
import json "encoding/json"
import ioutil "io/ioutil"
import bytes "bytes"
import context "golang.org/x/net/context"
import metadata "google.golang.org/grpc/metadata"
import runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
import validate_runtime "github.com/infobloxopen/protoc-gen-atlas-validate/runtime"
import google_protobuf1 "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
import google_protobuf2 "github.com/infobloxopen/atlas-app-toolkit/query"
import google_protobuf3 "google.golang.org/genproto/protobuf/field_mask"
import proto "github.com/gogo/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "github.com/infobloxopen/atlas-app-toolkit/query"
import _ "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
import _ "github.com/infobloxopen/protoc-gen-atlas-query-validate/options"
import _ "github.com/infobloxopen/protoc-gen-atlas-validate/options"
import _ "github.com/infobloxopen/protoc-gen-gorm/options"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "google.golang.org/genproto/protobuf/field_mask"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// validate_Regions_Create_0 is an entrypoint for validating "POST" HTTP request
// that match *.pb.gw.go/pattern_Regions_Create_0.
func validate_Regions_Create_0(r json.RawMessage) (err error) {
	return validate_Object_Region(r, "", false)
}

// validate_Regions_Read_0 is an entrypoint for validating "GET" HTTP request
// that match *.pb.gw.go/pattern_Regions_Read_0.
func validate_Regions_Read_0(r json.RawMessage) (err error) {
	if len(r) != 0 {
		return fmt.Errorf("Body is not allowed")
	}
	return nil
}

// validate_Regions_Update_0 is an entrypoint for validating "PUT" HTTP request
// that match *.pb.gw.go/pattern_Regions_Update_0.
func validate_Regions_Update_0(r json.RawMessage) (err error) {
	return validate_Object_Region(r, "", false)
}

// validate_Regions_Update_1 is an entrypoint for validating "PATCH" HTTP request
// that match *.pb.gw.go/pattern_Regions_Update_1.
func validate_Regions_Update_1(r json.RawMessage) (err error) {
	return validate_Object_Region(r, "", false)
}

// validate_Regions_Delete_0 is an entrypoint for validating "DELETE" HTTP request
// that match *.pb.gw.go/pattern_Regions_Delete_0.
func validate_Regions_Delete_0(r json.RawMessage) (err error) {
	if len(r) != 0 {
		return fmt.Errorf("Body is not allowed")
	}
	return nil
}

// validate_Regions_List_0 is an entrypoint for validating "GET" HTTP request
// that match *.pb.gw.go/pattern_Regions_List_0.
func validate_Regions_List_0(r json.RawMessage) (err error) {
	if len(r) != 0 {
		return fmt.Errorf("Body is not allowed")
	}
	return nil
}

// validate_Cmdb_GetVersion_0 is an entrypoint for validating "GET" HTTP request
// that match *.pb.gw.go/pattern_Cmdb_GetVersion_0.
func validate_Cmdb_GetVersion_0(r json.RawMessage) (err error) {
	if len(r) != 0 {
		return fmt.Errorf("Body is not allowed")
	}
	return nil
}

// validate_Object_Region function validates a JSON for a given object.
func validate_Object_Region(r json.RawMessage, path string, allowUnknown bool) (err error) {
	obj := &Region{}
	if hook, ok := interface{}(obj).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("Invalid value for %q: expected object.", path)
	}
	for k, _ := range v {
		switch k {
		case "id":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			validator, ok := interface{}(&google_protobuf1.Identifier{}).(interface {
				AtlasValidateJSON(json.RawMessage, string, bool) error
			})
			if !ok {
				continue
			}
			if err = validator.AtlasValidateJSON(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		case "name":
		case "middle_name":
		default:
			if !allowUnknown {
				return fmt.Errorf("Unknown field %q", validate_runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object Region.
func (o *Region) AtlasValidateJSON(r json.RawMessage, path string, allowUnknown bool) (err error) {
	if hook, ok := interface{}(o).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	return validate_Object_Region(r, path, allowUnknown)
}

// validate_Object_CreateRegionRequest function validates a JSON for a given object.
func validate_Object_CreateRegionRequest(r json.RawMessage, path string, allowUnknown bool) (err error) {
	obj := &CreateRegionRequest{}
	if hook, ok := interface{}(obj).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("Invalid value for %q: expected object.", path)
	}
	for k, _ := range v {
		switch k {
		case "payload":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			if err = validate_Object_Region(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("Unknown field %q", validate_runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object CreateRegionRequest.
func (o *CreateRegionRequest) AtlasValidateJSON(r json.RawMessage, path string, allowUnknown bool) (err error) {
	if hook, ok := interface{}(o).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	return validate_Object_CreateRegionRequest(r, path, allowUnknown)
}

// validate_Object_CreateRegionResponse function validates a JSON for a given object.
func validate_Object_CreateRegionResponse(r json.RawMessage, path string, allowUnknown bool) (err error) {
	obj := &CreateRegionResponse{}
	if hook, ok := interface{}(obj).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("Invalid value for %q: expected object.", path)
	}
	for k, _ := range v {
		switch k {
		case "result":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			if err = validate_Object_Region(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("Unknown field %q", validate_runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object CreateRegionResponse.
func (o *CreateRegionResponse) AtlasValidateJSON(r json.RawMessage, path string, allowUnknown bool) (err error) {
	if hook, ok := interface{}(o).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	return validate_Object_CreateRegionResponse(r, path, allowUnknown)
}

// validate_Object_ReadRegionRequest function validates a JSON for a given object.
func validate_Object_ReadRegionRequest(r json.RawMessage, path string, allowUnknown bool) (err error) {
	obj := &ReadRegionRequest{}
	if hook, ok := interface{}(obj).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("Invalid value for %q: expected object.", path)
	}
	for k, _ := range v {
		switch k {
		case "id":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			validator, ok := interface{}(&google_protobuf1.Identifier{}).(interface {
				AtlasValidateJSON(json.RawMessage, string, bool) error
			})
			if !ok {
				continue
			}
			if err = validator.AtlasValidateJSON(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		case "fields":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			validator, ok := interface{}(&google_protobuf2.FieldSelection{}).(interface {
				AtlasValidateJSON(json.RawMessage, string, bool) error
			})
			if !ok {
				continue
			}
			if err = validator.AtlasValidateJSON(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("Unknown field %q", validate_runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object ReadRegionRequest.
func (o *ReadRegionRequest) AtlasValidateJSON(r json.RawMessage, path string, allowUnknown bool) (err error) {
	if hook, ok := interface{}(o).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	return validate_Object_ReadRegionRequest(r, path, allowUnknown)
}

// validate_Object_ReadRegionResponse function validates a JSON for a given object.
func validate_Object_ReadRegionResponse(r json.RawMessage, path string, allowUnknown bool) (err error) {
	obj := &ReadRegionResponse{}
	if hook, ok := interface{}(obj).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("Invalid value for %q: expected object.", path)
	}
	for k, _ := range v {
		switch k {
		case "result":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			if err = validate_Object_Region(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("Unknown field %q", validate_runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object ReadRegionResponse.
func (o *ReadRegionResponse) AtlasValidateJSON(r json.RawMessage, path string, allowUnknown bool) (err error) {
	if hook, ok := interface{}(o).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	return validate_Object_ReadRegionResponse(r, path, allowUnknown)
}

// validate_Object_UpdateRegionRequest function validates a JSON for a given object.
func validate_Object_UpdateRegionRequest(r json.RawMessage, path string, allowUnknown bool) (err error) {
	obj := &UpdateRegionRequest{}
	if hook, ok := interface{}(obj).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("Invalid value for %q: expected object.", path)
	}
	for k, _ := range v {
		switch k {
		case "payload":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			if err = validate_Object_Region(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		case "fields":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			validator, ok := interface{}(&google_protobuf3.FieldMask{}).(interface {
				AtlasValidateJSON(json.RawMessage, string, bool) error
			})
			if !ok {
				continue
			}
			if err = validator.AtlasValidateJSON(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("Unknown field %q", validate_runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object UpdateRegionRequest.
func (o *UpdateRegionRequest) AtlasValidateJSON(r json.RawMessage, path string, allowUnknown bool) (err error) {
	if hook, ok := interface{}(o).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	return validate_Object_UpdateRegionRequest(r, path, allowUnknown)
}

// validate_Object_UpdateRegionResponse function validates a JSON for a given object.
func validate_Object_UpdateRegionResponse(r json.RawMessage, path string, allowUnknown bool) (err error) {
	obj := &UpdateRegionResponse{}
	if hook, ok := interface{}(obj).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("Invalid value for %q: expected object.", path)
	}
	for k, _ := range v {
		switch k {
		case "result":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			if err = validate_Object_Region(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("Unknown field %q", validate_runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object UpdateRegionResponse.
func (o *UpdateRegionResponse) AtlasValidateJSON(r json.RawMessage, path string, allowUnknown bool) (err error) {
	if hook, ok := interface{}(o).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	return validate_Object_UpdateRegionResponse(r, path, allowUnknown)
}

// validate_Object_DeleteRegionRequest function validates a JSON for a given object.
func validate_Object_DeleteRegionRequest(r json.RawMessage, path string, allowUnknown bool) (err error) {
	obj := &DeleteRegionRequest{}
	if hook, ok := interface{}(obj).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("Invalid value for %q: expected object.", path)
	}
	for k, _ := range v {
		switch k {
		case "id":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			validator, ok := interface{}(&google_protobuf1.Identifier{}).(interface {
				AtlasValidateJSON(json.RawMessage, string, bool) error
			})
			if !ok {
				continue
			}
			if err = validator.AtlasValidateJSON(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("Unknown field %q", validate_runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object DeleteRegionRequest.
func (o *DeleteRegionRequest) AtlasValidateJSON(r json.RawMessage, path string, allowUnknown bool) (err error) {
	if hook, ok := interface{}(o).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	return validate_Object_DeleteRegionRequest(r, path, allowUnknown)
}

// validate_Object_DeleteRegionResponse function validates a JSON for a given object.
func validate_Object_DeleteRegionResponse(r json.RawMessage, path string, allowUnknown bool) (err error) {
	obj := &DeleteRegionResponse{}
	if hook, ok := interface{}(obj).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("Invalid value for %q: expected object.", path)
	}
	for k, _ := range v {
		switch k {
		default:
			if !allowUnknown {
				return fmt.Errorf("Unknown field %q", validate_runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object DeleteRegionResponse.
func (o *DeleteRegionResponse) AtlasValidateJSON(r json.RawMessage, path string, allowUnknown bool) (err error) {
	if hook, ok := interface{}(o).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	return validate_Object_DeleteRegionResponse(r, path, allowUnknown)
}

// validate_Object_ListRegionRequest function validates a JSON for a given object.
func validate_Object_ListRegionRequest(r json.RawMessage, path string, allowUnknown bool) (err error) {
	obj := &ListRegionRequest{}
	if hook, ok := interface{}(obj).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("Invalid value for %q: expected object.", path)
	}
	for k, _ := range v {
		switch k {
		case "filter":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			validator, ok := interface{}(&google_protobuf2.Filtering{}).(interface {
				AtlasValidateJSON(json.RawMessage, string, bool) error
			})
			if !ok {
				continue
			}
			if err = validator.AtlasValidateJSON(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		case "order_by":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			validator, ok := interface{}(&google_protobuf2.Sorting{}).(interface {
				AtlasValidateJSON(json.RawMessage, string, bool) error
			})
			if !ok {
				continue
			}
			if err = validator.AtlasValidateJSON(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		case "fields":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			validator, ok := interface{}(&google_protobuf2.FieldSelection{}).(interface {
				AtlasValidateJSON(json.RawMessage, string, bool) error
			})
			if !ok {
				continue
			}
			if err = validator.AtlasValidateJSON(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		case "paging":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			validator, ok := interface{}(&google_protobuf2.Pagination{}).(interface {
				AtlasValidateJSON(json.RawMessage, string, bool) error
			})
			if !ok {
				continue
			}
			if err = validator.AtlasValidateJSON(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("Unknown field %q", validate_runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object ListRegionRequest.
func (o *ListRegionRequest) AtlasValidateJSON(r json.RawMessage, path string, allowUnknown bool) (err error) {
	if hook, ok := interface{}(o).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	return validate_Object_ListRegionRequest(r, path, allowUnknown)
}

// validate_Object_ListRegionsResponse function validates a JSON for a given object.
func validate_Object_ListRegionsResponse(r json.RawMessage, path string, allowUnknown bool) (err error) {
	obj := &ListRegionsResponse{}
	if hook, ok := interface{}(obj).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("Invalid value for %q: expected object.", path)
	}
	for k, _ := range v {
		switch k {
		case "results":
			if v[k] == nil {
				continue
			}
			var vArr []json.RawMessage
			vArrPath := validate_runtime.JoinPath(path, k)
			if err = json.Unmarshal(v[k], &vArr); err != nil {
				return fmt.Errorf("Invalid value for %q: expected array.", vArrPath)
			}
			for i, vv := range vArr {
				vvPath := fmt.Sprintf("%s.[%d]", vArrPath, i)
				if err = validate_Object_Region(vv, vvPath, allowUnknown); err != nil {
					return err
				}
			}
		case "page":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := validate_runtime.JoinPath(path, k)
			validator, ok := interface{}(&google_protobuf2.PageInfo{}).(interface {
				AtlasValidateJSON(json.RawMessage, string, bool) error
			})
			if !ok {
				continue
			}
			if err = validator.AtlasValidateJSON(vv, vvPath, allowUnknown); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("Unknown field %q", validate_runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object ListRegionsResponse.
func (o *ListRegionsResponse) AtlasValidateJSON(r json.RawMessage, path string, allowUnknown bool) (err error) {
	if hook, ok := interface{}(o).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	return validate_Object_ListRegionsResponse(r, path, allowUnknown)
}

// validate_Object_VersionResponse function validates a JSON for a given object.
func validate_Object_VersionResponse(r json.RawMessage, path string, allowUnknown bool) (err error) {
	obj := &VersionResponse{}
	if hook, ok := interface{}(obj).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("Invalid value for %q: expected object.", path)
	}
	for k, _ := range v {
		switch k {
		case "version":
		default:
			if !allowUnknown {
				return fmt.Errorf("Unknown field %q", validate_runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object VersionResponse.
func (o *VersionResponse) AtlasValidateJSON(r json.RawMessage, path string, allowUnknown bool) (err error) {
	if hook, ok := interface{}(o).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	return validate_Object_VersionResponse(r, path, allowUnknown)
}

var validate_Patterns = []struct {
	pattern    runtime.Pattern
	httpMethod string
	validator  func(json.RawMessage) error
	// Included for introspection purpose.
	allowUnknown bool
}{
	{
		pattern:      pattern_Regions_Create_0,
		httpMethod:   "POST",
		validator:    validate_Regions_Create_0,
		allowUnknown: false,
	},
	{
		pattern:      pattern_Regions_Read_0,
		httpMethod:   "GET",
		validator:    validate_Regions_Read_0,
		allowUnknown: false,
	},
	{
		pattern:      pattern_Regions_Update_0,
		httpMethod:   "PUT",
		validator:    validate_Regions_Update_0,
		allowUnknown: false,
	},
	{
		pattern:      pattern_Regions_Update_1,
		httpMethod:   "PATCH",
		validator:    validate_Regions_Update_1,
		allowUnknown: false,
	},
	{
		pattern:      pattern_Regions_Delete_0,
		httpMethod:   "DELETE",
		validator:    validate_Regions_Delete_0,
		allowUnknown: false,
	},
	{
		pattern:      pattern_Regions_List_0,
		httpMethod:   "GET",
		validator:    validate_Regions_List_0,
		allowUnknown: false,
	},
	{
		pattern:      pattern_Cmdb_GetVersion_0,
		httpMethod:   "GET",
		validator:    validate_Cmdb_GetVersion_0,
		allowUnknown: false,
	},
}

// AtlasValidateAnnotator parses JSON input and validates unknown fields
// based on 'allow_unknown_fields' options specified in proto file.
func AtlasValidateAnnotator(ctx context.Context, r *http.Request) metadata.MD {
	md := make(metadata.MD)
	for _, v := range validate_Patterns {
		if r.Method == v.httpMethod && validate_runtime.PatternMatch(v.pattern, r.URL.Path) {
			var b []byte
			var err error
			if b, err = ioutil.ReadAll(r.Body); err != nil {
				md.Set("Atlas-Validation-Error", "Invalid value: unable to parse body")
				return md
			}
			r.Body = ioutil.NopCloser(bytes.NewReader(b))
			if err = v.validator(b); err != nil {
				md.Set("Atlas-Validation-Error", err.Error())
			}
			break
		}
	}
	return md
}
