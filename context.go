package chi

import (
	"context"
	"net/http"
	"strings"
)

// RouteContext holds the routing information for the request.
type RouteContext struct {
	RoutePatterns []string
	// ... other fields
}

// RoutePattern returns the full route pattern for the request.
func (r *RouteContext) RoutePattern() string {
	return strings.Join(r.RoutePatterns, "")
}

// ... existing implementation ...