package mingo

import "time"

const (
	// all header constants
	HEADER_CONTENT_TYPE = "Content-Type"
	HEADER_USER_AGENT   = "User-Agent"

	// all content type constants
	CONTENT_TYPE_JSON = "application/json"
	CONTENT_TYPE_XML  = "application/xml"

	// default values for timeouts
	DEFAULT_MAX_IDLE_CONNECTIONS = 10
	DEFAULT_RESPONSE_TIMEOUT     = 10 * time.Second
	DEFAULT_CONNECTION_TIMEOUT   = 10 * time.Second
)
