package api

import "github.com/go-chi/chi/v5"

func InitRoute(r chi.Router) {
	InitAibox_eventRoute(r)
	InitAibox_deviceRoute(r)
	InitAibox_event_infoRoute(r)
	InitAibox_device_infoRoute(r)
	InitMessageRoute(r)
	InitAibox_updateRoute(r)
	InitAibox_update_infoRoute(r)
	InitAiboxUpdateExtRoute(r)
}
