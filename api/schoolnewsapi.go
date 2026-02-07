package main

import (
	"encoding/json"
	"nknu-core/school_news"
	"nknu-core/utils"
)

func CountNewsApi() string {
	count, err := school_news.CountNews()
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	res, _ := json.Marshal(count)
	return utils.FormatBase64Output(string(res), nil)
}

func CountNewsByCategoryApi(category string) string {
	count, err := school_news.CountNewsByCategory(category)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	res, _ := json.Marshal(count)
	return utils.FormatBase64Output(string(res), nil)
}

func CountNewsByPublisherApi(publisher string) string {
	count, err := school_news.CountNewsByPublisher(publisher)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	res, _ := json.Marshal(count)
	return utils.FormatBase64Output(string(res), nil)
}

func GetNewsApi(start, end int) string {
	items, err := school_news.GetNews(start, end)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	res, err := json.Marshal(items)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(res), nil)
}

func GetNewsByCategoryApi(category string, start, end int) string {
	items, err := school_news.GetNewsByCategory(category, start, end)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	res, err := json.Marshal(items)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(res), nil)
}

func GetNewsByPublisherApi(publisher string, start, end int) string {
	items, err := school_news.GetNewsByPublisher(publisher, start, end)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	res, err := json.Marshal(items)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(res), nil)
}

func ForceRefreshNewsApi() string {
	err := school_news.ForceRefresh()
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output("success", nil)
}

func GetLastNewsRefreshTimeApi() string {
	timeStr := school_news.GetLastRefreshTime()
	return utils.FormatBase64Output(timeStr, nil)
}
