package mingo

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"sync"
)

// mockSrv is the global mock server.
var mockSrv = mockServer{
	mocks: make(map[string]*Mock),
}

// mockServer struct takes mocks and stores them as key value pairs for running in test cases
// it will only proceed when enabled is true (check getMock for more info).
type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex
	mocks       map[string]*Mock
}

// starts up the mock server
func StartMockServer() {
	mockSrv.serverMutex.Lock()
	defer mockSrv.serverMutex.Unlock()
	mockSrv.enabled = true
}

// stops the mock server
func StopMockServer() {
	mockSrv.serverMutex.Lock()
	defer mockSrv.serverMutex.Unlock()
	mockSrv.enabled = false
}

// takes in a mock and assigns necessary stuffs to the mockSrv(mock server)
func AddMock(mock Mock) {
	mockSrv.serverMutex.Lock()
	defer mockSrv.serverMutex.Unlock()
	key := mockSrv.getMockKey(mock.Method, mock.Url, mock.ReqBody)
	mockSrv.mocks[key] = &mock
}

// hashes the mock key with md5 and returs it,
// uses method + url + body to hash.
func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
	return key
}

// a helper function that cleans the body passed (all spaces/tabs)
func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\t", "")
	return body
}

// returns a mock from the hashed key from getMockKey function
func (m *mockServer) getMock(method, url, body string) *Mock {
	// if not enabled then return nil
	if !m.enabled {
		return nil
	}

	if mock := m.mocks[m.getMockKey(method, url, body)]; mock != nil {
		return mock
	}

	return &Mock{
		Error: errors.New("no mock matching"),
	}
}

// cleans up the mocks present earlier.
func FlushMocks() {
	mockSrv.serverMutex.Lock()
	defer mockSrv.serverMutex.Unlock()

	mockSrv.mocks = make(map[string]*Mock)
}
