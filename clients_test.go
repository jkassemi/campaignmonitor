package campaignmonitor

import (
  "testing"
  "net/http"
  "net/http/httptest"
)

func TestCreateClient(t *testing.T){
  server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte(`"aa164e9c8ab0471294fe6148fc9cf634"`))
  }))

  apiClient := &ApiClient{
    Endpoint: server.URL,
    ApiKey: "abc",
  }

  clientId, e := apiClient.CreateClient("Test Client", "United States", "(GMT+10:00) Canberra, Melbourne, Sydney")

  if e != nil {
    t.Fatal(e.Error())
  }

  if clientId != "aa164e9c8ab0471294fe6148fc9cf634" {
    t.Fatal("Unexpted client id")
  }
}

func TestGetClient(t *testing.T){
  server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte(`{
      "ApiKey": "639d8cc27198202f5fe6037a8b17a29a59984b86d3289bc9",
      "BasicDetails": {
          "ClientID": "4a397ccaaa55eb4e6aa1221e1e2d7122",
          "CompanyName": "Client One",
          "Country": "Australia",
          "TimeZone": "(GMT+10:00) Canberra, Melbourne, Sydney",
          "PrimaryContactName": "Sally",
          "PrimaryContactEmail": "sally@me.com"
      },
      "BillingDetails": {
          "CanPurchaseCredits": true,
          "Credits": 500,
          "MarkupOnDesignSpamTest": 0.0,
          "ClientPays": true,
          "BaseRatePerRecipient": 1.0,
          "MarkupPerRecipient": 0.0,
          "MarkupOnDelivery": 0.0,
          "BaseDeliveryRate": 5.0,
          "Currency": "USD",
          "BaseDesignSpamTestRate": 5.0
      }
    }`))
  }))

  apiClient := &ApiClient{
    Endpoint: server.URL,
    ApiKey: "abc",
  }

  client, e := apiClient.GetClient("abc")

  if e != nil {
    t.Fatal(e.Error())
  }

  if client.BasicDetails.ClientID != "4a397ccaaa55eb4e6aa1221e1e2d7122" {
    t.Fatal("Unexpected client id")
  }

  if client.BillingDetails.CanPurchaseCredits != true {
    t.Fatal("Unexpected billing detail")
  }

  // TODO: We can do a lot more testing here if we wanted to get more thorough.
}

func TestGetClientSubscriberList(t *testing.T){
  server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte(`
  [
    {
        "ListID": "a58ee1d3039b8bec838e6d1482a8a965",
        "Name": "List One"
    },
    {
        "ListID": "99bc35084a5739127a8ab81eae5bd305",
        "Name": "List Two"
    }
  ]  
    `))
  }))

  apiClient := &ApiClient{
    Endpoint: server.URL,
    ApiKey: "abc",
  }

  list, e := apiClient.GetClientSubscriberLists("abc")

  if e != nil {
    t.Fatal(e.Error())
  }

  if len(list) != 2 {
    t.Fatal("Unexpected length of return")
  }

  if list[0].ListID != "a58ee1d3039b8bec838e6d1482a8a965" {
    t.Fatal("Bad data in list")
  }

  if list[0].Name != "List One" {
    t.Fatal("Bad data in list")
  }
}

func TestGetClientSubscriberListsForEmail(t *testing.T){
  server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte(`
  [
    {
        "ListID": "a58ee1d3039b8bec838e6d1482a8a965",
        "ListName": "List One",
        "SubscriberState": "Active",
        "DateSubscriberAdded": "2010-03-19 11:15:00"
    },
    {
        "ListID": "99bc35084a5739127a8ab81eae5bd305",
        "ListName": "List Two",
        "SubscriberState": "Unsubscribed",
        "DateSubscriberAdded": "2011-04-01 01:27:00"
    }
]  
    `))
  }))

  apiClient := &ApiClient{
    Endpoint: server.URL,
    ApiKey: "abc",
  }

  list, e := apiClient.GetClientSubscriberListsForEmail("abc", "test@example.com")

  if e != nil {
    t.Fatal(e.Error())
  }

  if len(list) != 2 {
    t.Fatal("Unexpected length of return")
  }

  if list[0].ListID != "a58ee1d3039b8bec838e6d1482a8a965" {
    t.Fatal("Bad data in list")
  }

  if list[0].ListName != "List One" {
    t.Fatal("Bad data in list")
  }

  if list[0].SubscriberState != "Active" {
    t.Fatal("Bad data in list")
  }

  /* if list[0].DateSubscriberAdded != 2010-03-19 11:15:00 {
    t.Fatal("Unexpected date")
  } */
}

func TestDeleteClient(t *testing.T){
  server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte(``))
  }))

  apiClient := &ApiClient{
    Endpoint: server.URL,
    ApiKey: "abc",
  }


  if e := apiClient.DeleteClient("abc"); e != nil {
    t.Fatal(e.Error())
  }
}

