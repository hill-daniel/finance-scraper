// +build unit

// Code generated by MockGen. DO NOT EDIT.
// Source: finance.go

// Package test is a generated GoMock package.
package test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	finance "github.com/hill-daniel/finance-scraper"
)

// MockScraper is a mock of Scraper interface.
type MockScraper struct {
	ctrl     *gomock.Controller
	recorder *MockScraperMockRecorder
}

// MockScraperMockRecorder is the mock recorder for MockScraper.
type MockScraperMockRecorder struct {
	mock *MockScraper
}

// NewMockScraper creates a new mock instance.
func NewMockScraper(ctrl *gomock.Controller) *MockScraper {
	mock := &MockScraper{ctrl: ctrl}
	mock.recorder = &MockScraperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScraper) EXPECT() *MockScraperMockRecorder {
	return m.recorder
}

// Scrape mocks base method.
func (m *MockScraper) Scrape(symbol string) (finance.ScrapeResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scrape", symbol)
	ret0, _ := ret[0].(finance.ScrapeResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Scrape indicates an expected call of Scrape.
func (mr *MockScraperMockRecorder) Scrape(symbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scrape", reflect.TypeOf((*MockScraper)(nil).Scrape), symbol)
}

// MockAnalyzer is a mock of Analyzer interface.
type MockAnalyzer struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyzerMockRecorder
}

// MockAnalyzerMockRecorder is the mock recorder for MockAnalyzer.
type MockAnalyzerMockRecorder struct {
	mock *MockAnalyzer
}

// NewMockAnalyzer creates a new mock instance.
func NewMockAnalyzer(ctrl *gomock.Controller) *MockAnalyzer {
	mock := &MockAnalyzer{ctrl: ctrl}
	mock.recorder = &MockAnalyzerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnalyzer) EXPECT() *MockAnalyzerMockRecorder {
	return m.recorder
}

// Analyze mocks base method.
func (m *MockAnalyzer) Analyze(quote finance.Quote) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Analyze", quote)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Analyze indicates an expected call of Analyze.
func (mr *MockAnalyzerMockRecorder) Analyze(quote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Analyze", reflect.TypeOf((*MockAnalyzer)(nil).Analyze), quote)
}
