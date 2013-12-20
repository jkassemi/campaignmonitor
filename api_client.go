package campaignmonitor

import (
  "errors"
  "strings"
  "net/url"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type ApiClient struct {
  ApiKey        string
  Endpoint      string
}

type ErrorResponse struct {
  Code          int
  Message       string
}

const cmEndpoint = "https://api.createsend.com/api/v3"

func NewApiClient(apiKey string) *ApiClient {
  return &ApiClient{
    ApiKey: apiKey,
    Endpoint: cmEndpoint,
  }
}

func (c *ApiClient) request (method string, point string, params map[string]string, target interface{}) (e error) {
  if c.ApiKey == "" {
    return errors.New("ApiKey required on client")
  }

  if c.Endpoint == "" {
    return errors.New("Where'd your endpoint go?")
  }

  q := url.Values{}

  if params != nil {
    for k, v := range params {
      if v != "" {
        q.Set(k, v)
      }
    }
  }

  var req *http.Request

  if method == "GET" {
    req, e = http.NewRequest(method, c.Endpoint +  point + ".json?" + q.Encode(), nil)

    if e != nil {
      return e
    }

  } else {
    req, e = http.NewRequest(method, c.Endpoint +  point + ".json", strings.NewReader(q.Encode()))

    if e != nil {
      return e
    }

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  }

  req.SetBasicAuth(c.ApiKey, "trash")

  res, e := http.DefaultClient.Do(req)

  if e != nil {
    return e
  }

  defer res.Body.Close()

  b, e := ioutil.ReadAll(res.Body)

  if e != nil {
    return e
  }

  if res.StatusCode == 401 {
    var errorResponse ErrorResponse

    e = json.Unmarshal(b, errorResponse)

    if e != nil {
      return e
    }

    return errors.New(errorResponse.Message)

  } else if res.StatusCode >= 200 && res.StatusCode < 300 {
    if target != nil {
      e = json.Unmarshal(b, target)

      if e != nil {
        return e
      }
    }

    return nil
  }

  return errors.New("Unexpected response from server, " + res.Status)
}
