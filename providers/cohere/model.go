package cohere

import (
	"errors"
	"net/http"
	"net/url"
)

func (p *CohereProvider) GetModelList() ([]string, error) {
	params := url.Values{}
	params.Add("page_size", "1000")
	// params.Add("endpoint[]", "chat")
	// params.Add("endpoint[]", "rerank")
	queryString := params.Encode()

	fullRequestURL := p.GetFullRequestURL(p.Config.ModelList) + "?" + queryString
	headers := p.GetRequestHeaders()

	req, err := p.Requester.NewRequest(http.MethodGet, fullRequestURL, p.Requester.WithHeader(headers))
	if err != nil {
		return nil, errors.New("new_request_failed")
	}

	response := &ModelListResponse{}
	_, errWithCode := p.Requester.SendRequest(req, response, false)
	if errWithCode != nil {
		return nil, errors.New(errWithCode.Message)
	}

	var modelList []string
	for _, model := range response.Models {
		for _, endpoint := range model.Endpoints {
			if endpoint == "chat" || endpoint == "rerank" {
				modelList = append(modelList, model.Name)
				break
			}
		}
	}

	return modelList, nil
}
