package campaignmonitor

import (
  "errors"
)

type Client struct {
  ApiKey            string
  BasicDetails      struct {
    ClientID        string
    CompanyName     string
    Country         string
    TimeZone        string
    PrimaryContactName  string
    PrimaryContactEmail string
  }
  BillingDetails    struct {
    CanPurchaseCredits    bool
    Credits               int
    MarkupOnDesignSpamTest  float32
    ClientPays            bool
    BaseRatePerRecipient  float32
    MarkupOnDelivery      float32
    BaseDeliveryRate      float32
    Currency                    string
    BaseDesignSpamTestRate      float32
  }
}

type List struct {
  ListID    string
  Name      string
}

type ListState struct {
  ListID    string
  ListName  string
  SubscriberState string
  DateSubscriberAdded string
}

func (a *ApiClient) CreateClient(companyName, country, timeZone string) (clientId string, e error) {
  // TODO: timeZone validation

  if companyName == "" {
    return "", errors.New("companyName must not be blank")
  }

  if country == "" {
    return "", errors.New("country must not be blank")
  }

  if timeZone == "" {
    return "", errors.New("timeZone must not be blank")
  }

  e = a.request("POST", "/clients",  map[string]string{
    "CompanyName": companyName,
    "Country": country,
    "TimeZone": timeZone,
  }, &clientId)

  return
}

func (a *ApiClient) GetClient(clientId string) (*Client, error) {
  if clientId == "" {
    return nil, errors.New("clientId must not be blank")
  }

  client := &Client{}

  e := a.request("GET", "/clients/" + clientId, nil, client)

  return client, e
}

func (a *ApiClient) GetClientSubscriberLists(clientId string) ([]*List, error) {
  if clientId == "" {
    return nil, errors.New("clientId must not be blank")
  }

  list := make([]*List, 0)

  e := a.request("GET", "/clients/" + clientId + "/lists", nil, &list)

  return list, e
}

func (a *ApiClient) GetClientSubscriberListsForEmail(clientId, email string) ([]*ListState, error) {
  if clientId == "" {
    return nil, errors.New("clientId must not be blank")
  }

  if email == "" {
    return nil, errors.New("email must not be blank")
  }

  list := make([]*ListState, 0)

  e := a.request("GET", "/clients/" + clientId + "/listsforemail", map[string]string{
    "email": email,
  }, &list)

  return list, e
}

func (a *ApiClient) DeleteClient(clientId string) (error) {
  if clientId == "" {
    return errors.New("clientId must not be blank")
  }

  return a.request("DELETE", "/clients/" + clientId, nil, nil)
}
