package mingo

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"sync"
)

var mockSrv = mockServer{
	mocks: make(map[string]*Mock),
}

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex
	mocks       map[string]*Mock
}

func StartMockServer() {
	mockSrv.serverMutex.Lock()
	defer mockSrv.serverMutex.Unlock()
	mockSrv.enabled = true
}

func StopMockServer() {
	mockSrv.serverMutex.Lock()
	defer mockSrv.serverMutex.Unlock()
	mockSrv.enabled = false
}

func AddMock(mock Mock) {
	mockSrv.serverMutex.Lock()
	defer mockSrv.serverMutex.Unlock()
	key := mockSrv.getMockKey(mock.Method, mock.Url, mock.ReqBody)
	mockSrv.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
	return key
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\t", "")
	return body
}

func (m *mockServer) getMock(method, url, body string) *Mock {
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

func FlushMocks() {
	mockSrv.serverMutex.Lock()
	defer mockSrv.serverMutex.Unlock()

	mockSrv.mocks = make(map[string]*Mock)
}
