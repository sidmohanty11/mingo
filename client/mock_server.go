package mingo

import "sync"

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
	return method + url + body
}

func (m *mockServer) getMock(method, url, body string) *Mock {
	return m.mocks[m.getMockKey(method, url, body)]
}
