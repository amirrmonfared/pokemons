// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/amirrmonfared/pokemons/db/sqlc (interfaces: Store)

// Package mock_sqlc is a generated GoMock package.
package mock_sqlc

import (
	context "context"
	reflect "reflect"

	db "github.com/amirrmonfared/pokemons/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreatePokemon mocks base method.
func (m *MockStore) CreatePokemon(arg0 context.Context, arg1 db.CreatePokemonParams) (db.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePokemon", arg0, arg1)
	ret0, _ := ret[0].(db.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePokemon indicates an expected call of CreatePokemon.
func (mr *MockStoreMockRecorder) CreatePokemon(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePokemon", reflect.TypeOf((*MockStore)(nil).CreatePokemon), arg0, arg1)
}

// DeletePokemon mocks base method.
func (m *MockStore) DeletePokemon(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePokemon", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePokemon indicates an expected call of DeletePokemon.
func (mr *MockStoreMockRecorder) DeletePokemon(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePokemon", reflect.TypeOf((*MockStore)(nil).DeletePokemon), arg0, arg1)
}

// GetPokemon mocks base method.
func (m *MockStore) GetPokemon(arg0 context.Context, arg1 int64) (db.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPokemon", arg0, arg1)
	ret0, _ := ret[0].(db.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPokemon indicates an expected call of GetPokemon.
func (mr *MockStoreMockRecorder) GetPokemon(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPokemon", reflect.TypeOf((*MockStore)(nil).GetPokemon), arg0, arg1)
}

// GetPokemonByName mocks base method.
func (m *MockStore) GetPokemonByName(arg0 context.Context, arg1 string) (db.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPokemonByName", arg0, arg1)
	ret0, _ := ret[0].(db.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPokemonByName indicates an expected call of GetPokemonByName.
func (mr *MockStoreMockRecorder) GetPokemonByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPokemonByName", reflect.TypeOf((*MockStore)(nil).GetPokemonByName), arg0, arg1)
}

// ImportPokemon mocks base method.
func (m *MockStore) ImportPokemon(arg0 context.Context, arg1 db.CreatePokemonParams) (db.CreatePokemonResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportPokemon", arg0, arg1)
	ret0, _ := ret[0].(db.CreatePokemonResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportPokemon indicates an expected call of ImportPokemon.
func (mr *MockStoreMockRecorder) ImportPokemon(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportPokemon", reflect.TypeOf((*MockStore)(nil).ImportPokemon), arg0, arg1)
}

// ListPokemons mocks base method.
func (m *MockStore) ListPokemons(arg0 context.Context, arg1 db.ListPokemonsParams) ([]db.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPokemons", arg0, arg1)
	ret0, _ := ret[0].([]db.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPokemons indicates an expected call of ListPokemons.
func (mr *MockStoreMockRecorder) ListPokemons(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPokemons", reflect.TypeOf((*MockStore)(nil).ListPokemons), arg0, arg1)
}

// ListPokemonsByAbility mocks base method.
func (m *MockStore) ListPokemonsByAbility(arg0 context.Context, arg1 db.ListPokemonsByAbilityParams) ([]db.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPokemonsByAbility", arg0, arg1)
	ret0, _ := ret[0].([]db.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPokemonsByAbility indicates an expected call of ListPokemonsByAbility.
func (mr *MockStoreMockRecorder) ListPokemonsByAbility(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPokemonsByAbility", reflect.TypeOf((*MockStore)(nil).ListPokemonsByAbility), arg0, arg1)
}
