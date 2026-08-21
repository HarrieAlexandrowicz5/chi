package chi

// ... existing imports ...

func (mx *Mux) route(w http.ResponseWriter, r *http.Request, path string) bool {
	// ... existing logic ...
	// When entering a subrouter, ensure the current pattern is appended to the context
	rctx := RouteContext(r.Context())
	rctx.RoutePatterns = append(rctx.RoutePatterns, path)
	// ... existing logic ...
}