package api

import "encoding/json"

type Entity interface{}

type APIResponse struct {
	Status  string            `json:"status"`
	Code    int               `json:"code"`
	Data    map[string]Entity `json:"data"`
	Message string            `json:"message"`
}

func NewAPIResponse(key string, e Entity) *APIResponse {
	r := new(APIResponse)
	r.Data = make(map[string]Entity)
	r.Data[key] = e
	return r
}

func (r *APIResponse) AddEntity(key string, e Entity) {
	r.Data[key] = e
}

func (r *APIResponse) UnmarshalJSON(b []byte) error {
	data := struct {
		Status  string
		Code    int
		Data    map[string]json.RawMessage
		Message string
	}{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	r.Status = data.Status
	r.Code = data.Code
	r.Message = data.Message
	for key, ent := range r.Data {
		if raw, ok := data.Data[key]; ok {
			if err = json.Unmarshal(raw, ent); err != nil {
				return err
			}
			r.Data[key] = ent
		}
	}
	return nil
}
