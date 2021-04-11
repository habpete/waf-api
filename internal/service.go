package internal

import "waf/internal/storage"

type IServiceFX interface {}

type serviceFX struct {
	reports reports.IReports,
	storage storage.IStorage
}

func New () IServiceFX {
	return &serviceFX{}
}

