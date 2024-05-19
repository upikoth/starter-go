package http

func (h *HTTP) startRouting() {
	h.router.GET("/api/v1/health", h.v1.CheckHealth)
}
