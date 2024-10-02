package controllers

import "UrlShorteningService/models"

func ShortenResponse(stats *models.Stats) models.ShortenResponse {
	return models.ShortenResponse{
		Id:        stats.Id,
		Url:       stats.Url_info.Url,
		CreatedAt: stats.Url_info.CreatedAt,
		ShortCode: stats.Url_info.ShortCode,
		UpdatedAt: stats.Url_info.UpdatedAt,
	}
}
