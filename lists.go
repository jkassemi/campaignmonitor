package campaignmonitor

import (
	"errors"
)

type ListStats struct {
	TotalActiveSubscribers        int32
	NewActiveSubscribersToday     int32
	NewActiveSubscribersYesterday int32
	NewActiveSubscribersThisWeek  int32
	NewActiveSubscribersThisMonth int32
	NewActiveSubscribersThisYear  int32
	TotalUnsubscribes             int32
	UnsubscribesToday             int32
	UnsubscribesYesterday         int32
	UnsubscribesThisWeek          int32
	UnsubscribesThisMonth         int32
	UnsubscribesThisYear          int32
	TotalDeleted                  int32
	DeletedToday                  int32
	DeletedYesterday              int32
	DeletedThisWeek               int32
	DeletedThisMonth              int32
	DeletedThisYear               int32
	TotalBounces                  int32
	BouncesToday                  int32
	BouncesYesterday              int32
	BouncesThisWeek               int32
	BouncesThisMonth              int32
	BouncesThisYear               int32
}

func (a *ApiClient) CreateList(clientId, title, unsubscribePage, unsubscribeSetting string, confirmedOptIn bool, confirmationSuccessPage string) (listId string, e error) {
	// TODO: timeZone validation

	if clientId == "" {
		return "", errors.New("clientId must not be blank")
	}

	if title == "" {
		return "", errors.New("title must not be blank")
	}

	if unsubscribeSetting != "AllClientLists" && unsubscribeSetting != "OnlyThisList" {
		return "", errors.New("unsubscribeSetting must be 'AllClientLists' or 'OnlyThisList'")
	}

	var strConfirmedOptIn string

	if confirmedOptIn == true {
		strConfirmedOptIn = "true"
	} else {
		strConfirmedOptIn = "false"
	}

	e = a.request("POST", "/lists/"+clientId, map[string]string{
		"Title":                   title,
		"UnsubscribePage":         unsubscribePage,
		"UnsubscribeSetting":      unsubscribeSetting,
		"ConfirmedOptIn":          strConfirmedOptIn,
		"ConfirmationSuccessPage": confirmationSuccessPage,
	}, &listId)

	return
}

func (a *ApiClient) ListStats(listId string) (*ListStats, error) {
	if listId == "" {
		return nil, errors.New("listId must not be blank")
	}

	listStats := &ListStats{}

	e := a.request("GET", "/lists/"+listId+"/stats", nil, listStats)

	return listStats, e
}

func (a *ApiClient) DeleteList(listId string) (error) {
	if listId == "" {
		return errors.New("listId must not be blank")
	}

	return a.request("DELETE", "/lists/"+listId, nil, nil)
}

