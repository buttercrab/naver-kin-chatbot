package httpHandler

import (
	"encoding/json"
	"net/http"
)

type chatBotJson struct {
	UserRequest userRequest `json:"userRequest"`
	Bot         bot         `json:"bot"`
	Action      action      `json:"action"`
}

type userRequest struct {
	Timezone  string `json:"timezone"`
	Utterance string `json:"utterance"`
	Lang      string `json:"lang"`
	User      user   `json:"user"`
}

type user struct {
	Id         string         `json:"id"`
	Type       string         `json:"type"`
	Properties userProperties `json:"properties"`
}

type userProperties struct {
	PlusFriendUserKey string `json:"plusfriendUserKey"`
	AppUserId         string `json:"appUserId"`
}

type action struct {
	Id           string                 `json:"id"`
	Name         string                 `json:"name"`
	Params       map[string]string      `json:"params"`
	DetailParams map[string]detailParam `json:"detailParams"`
}

type detailParam struct {
	Origin    string `json:"origin"`
	Value     string `json:"value"`
	GroupName string `json:"groupName"`
}

type bot struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type chatBotResponse struct {
	Version  string        `json:"version"`
	Template skillTemplate `json:"template"`
}

type skillTemplate struct {
	Outputs []component `json:"outputs"`
}

type component struct {
	SimpleText *simpleText `json:"simpleText,omitempty"`
	BasicCard  *basicCard  `json:"basicCard,omitempty"`
}

type simpleText struct {
	Text string `json:"text"`
}

type basicCard struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Buttons     []button `json:"buttons"`
}

type button struct {
	Label      string `json:"label"`
	Action     string `json:"action"`
	WebLinkUrl string `json:"webLinkUrl"`
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid method", 400)
		return
	}

	var c chatBotJson
	if r.Body == nil {
		http.Error(w, "Please send a body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "JSON decoding error: "+err.Error(), 400)
		return
	}

	res := "Response Code"

	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(res)
}
