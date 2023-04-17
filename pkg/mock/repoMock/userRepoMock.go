// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interface/userRepoInterface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	utils "github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AcceptJoinInvitation mocks base method.
func (m *MockUserRepository) AcceptJoinInvitation(user_id, organizaiton_id int, memberRole string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptJoinInvitation", user_id, organizaiton_id, memberRole)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptJoinInvitation indicates an expected call of AcceptJoinInvitation.
func (mr *MockUserRepositoryMockRecorder) AcceptJoinInvitation(user_id, organizaiton_id, memberRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptJoinInvitation", reflect.TypeOf((*MockUserRepository)(nil).AcceptJoinInvitation), user_id, organizaiton_id, memberRole)
}

// AdmitMember mocks base method.
func (m *MockUserRepository) AdmitMember(JoinStatusId int, memberRole string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdmitMember", JoinStatusId, memberRole)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdmitMember indicates an expected call of AdmitMember.
func (mr *MockUserRepositoryMockRecorder) AdmitMember(JoinStatusId, memberRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdmitMember", reflect.TypeOf((*MockUserRepository)(nil).AdmitMember), JoinStatusId, memberRole)
}

// ApplyEvent mocks base method.
func (m *MockUserRepository) ApplyEvent(applicationForm domain.ApplicationForm) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyEvent", applicationForm)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyEvent indicates an expected call of ApplyEvent.
func (mr *MockUserRepositoryMockRecorder) ApplyEvent(applicationForm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyEvent", reflect.TypeOf((*MockUserRepository)(nil).ApplyEvent), applicationForm)
}

// CreateOrganization mocks base method.
func (m *MockUserRepository) CreateOrganization(organization domain.Organizations) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganization", organization)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganization indicates an expected call of CreateOrganization.
func (mr *MockUserRepositoryMockRecorder) CreateOrganization(organization interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganization", reflect.TypeOf((*MockUserRepository)(nil).CreateOrganization), organization)
}

// DeleteMember mocks base method.
func (m *MockUserRepository) DeleteMember(user_id, organizaiton_id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMember", user_id, organizaiton_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMember indicates an expected call of DeleteMember.
func (mr *MockUserRepositoryMockRecorder) DeleteMember(user_id, organizaiton_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMember", reflect.TypeOf((*MockUserRepository)(nil).DeleteMember), user_id, organizaiton_id)
}

// FeaturizeEvent mocks base method.
func (m *MockUserRepository) FeaturizeEvent(orderid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeaturizeEvent", orderid)
	ret0, _ := ret[0].(error)
	return ret0
}

// FeaturizeEvent indicates an expected call of FeaturizeEvent.
func (mr *MockUserRepositoryMockRecorder) FeaturizeEvent(orderid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeaturizeEvent", reflect.TypeOf((*MockUserRepository)(nil).FeaturizeEvent), orderid)
}

// FindApplication mocks base method.
func (m *MockUserRepository) FindApplication(user_id, event_id int) (domain.ApplicationFormResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindApplication", user_id, event_id)
	ret0, _ := ret[0].(domain.ApplicationFormResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindApplication indicates an expected call of FindApplication.
func (mr *MockUserRepositoryMockRecorder) FindApplication(user_id, event_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindApplication", reflect.TypeOf((*MockUserRepository)(nil).FindApplication), user_id, event_id)
}

// FindJoinStatus mocks base method.
func (m *MockUserRepository) FindJoinStatus(JoinStatusId int) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindJoinStatus", JoinStatusId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindJoinStatus indicates an expected call of FindJoinStatus.
func (mr *MockUserRepositoryMockRecorder) FindJoinStatus(JoinStatusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindJoinStatus", reflect.TypeOf((*MockUserRepository)(nil).FindJoinStatus), JoinStatusId)
}

// FindOrganizationById mocks base method.
func (m *MockUserRepository) FindOrganizationById(organizaiton_id int) (domain.OrganizationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrganizationById", organizaiton_id)
	ret0, _ := ret[0].(domain.OrganizationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganizationById indicates an expected call of FindOrganizationById.
func (mr *MockUserRepositoryMockRecorder) FindOrganizationById(organizaiton_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganizationById", reflect.TypeOf((*MockUserRepository)(nil).FindOrganizationById), organizaiton_id)
}

// FindOrganizationByName mocks base method.
func (m *MockUserRepository) FindOrganizationByName(organizationName string) (domain.OrganizationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrganizationByName", organizationName)
	ret0, _ := ret[0].(domain.OrganizationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganizationByName indicates an expected call of FindOrganizationByName.
func (mr *MockUserRepositoryMockRecorder) FindOrganizationByName(organizationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganizationByName", reflect.TypeOf((*MockUserRepository)(nil).FindOrganizationByName), organizationName)
}

// FindRelation mocks base method.
func (m *MockUserRepository) FindRelation(user_id, organizaiton_id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRelation", user_id, organizaiton_id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRelation indicates an expected call of FindRelation.
func (mr *MockUserRepositoryMockRecorder) FindRelation(user_id, organizaiton_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRelation", reflect.TypeOf((*MockUserRepository)(nil).FindRelation), user_id, organizaiton_id)
}

// FindRole mocks base method.
func (m *MockUserRepository) FindRole(user_id, organizaiton_id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRole", user_id, organizaiton_id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRole indicates an expected call of FindRole.
func (mr *MockUserRepositoryMockRecorder) FindRole(user_id, organizaiton_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRole", reflect.TypeOf((*MockUserRepository)(nil).FindRole), user_id, organizaiton_id)
}

// FindUserById mocks base method.
func (m *MockUserRepository) FindUserById(user_id int) (domain.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", user_id)
	ret0, _ := ret[0].(domain.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserRepositoryMockRecorder) FindUserById(user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserRepository)(nil).FindUserById), user_id)
}

// FindUserByName mocks base method.
func (m *MockUserRepository) FindUserByName(email string) (domain.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByName", email)
	ret0, _ := ret[0].(domain.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByName indicates an expected call of FindUserByName.
func (mr *MockUserRepositoryMockRecorder) FindUserByName(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByName", reflect.TypeOf((*MockUserRepository)(nil).FindUserByName), email)
}

// GetPublicFaqas mocks base method.
func (m *MockUserRepository) GetPublicFaqas(event_id int) ([]domain.QAResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicFaqas", event_id)
	ret0, _ := ret[0].([]domain.QAResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicFaqas indicates an expected call of GetPublicFaqas.
func (mr *MockUserRepositoryMockRecorder) GetPublicFaqas(event_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicFaqas", reflect.TypeOf((*MockUserRepository)(nil).GetPublicFaqas), event_id)
}

// GetQuestions mocks base method.
func (m *MockUserRepository) GetQuestions(evnent_id int) ([]domain.FaqaResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestions", evnent_id)
	ret0, _ := ret[0].([]domain.FaqaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestions indicates an expected call of GetQuestions.
func (mr *MockUserRepositoryMockRecorder) GetQuestions(evnent_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestions", reflect.TypeOf((*MockUserRepository)(nil).GetQuestions), evnent_id)
}

// InsertUser mocks base method.
func (m *MockUserRepository) InsertUser(user domain.Users) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockUserRepositoryMockRecorder) InsertUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockUserRepository)(nil).InsertUser), user)
}

// JoinOrganization mocks base method.
func (m *MockUserRepository) JoinOrganization(organization_id, user_id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JoinOrganization", organization_id, user_id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JoinOrganization indicates an expected call of JoinOrganization.
func (mr *MockUserRepositoryMockRecorder) JoinOrganization(organization_id, user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JoinOrganization", reflect.TypeOf((*MockUserRepository)(nil).JoinOrganization), organization_id, user_id)
}

// ListJoinRequests mocks base method.
func (m *MockUserRepository) ListJoinRequests(user_id, organizaiton_id int) ([]domain.Join_StatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListJoinRequests", user_id, organizaiton_id)
	ret0, _ := ret[0].([]domain.Join_StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListJoinRequests indicates an expected call of ListJoinRequests.
func (mr *MockUserRepositoryMockRecorder) ListJoinRequests(user_id, organizaiton_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJoinRequests", reflect.TypeOf((*MockUserRepository)(nil).ListJoinRequests), user_id, organizaiton_id)
}

// ListMembers mocks base method.
func (m *MockUserRepository) ListMembers(memberRole string, organizaiton_id int) ([]domain.UserOrganizationConnectionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMembers", memberRole, organizaiton_id)
	ret0, _ := ret[0].([]domain.UserOrganizationConnectionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMembers indicates an expected call of ListMembers.
func (mr *MockUserRepositoryMockRecorder) ListMembers(memberRole, organizaiton_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMembers", reflect.TypeOf((*MockUserRepository)(nil).ListMembers), memberRole, organizaiton_id)
}

// ListOrganizations mocks base method.
func (m *MockUserRepository) ListOrganizations(pagenation utils.Filter) ([]domain.OrganizationsResponse, utils.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrganizations", pagenation)
	ret0, _ := ret[0].([]domain.OrganizationsResponse)
	ret1, _ := ret[1].(utils.Metadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListOrganizations indicates an expected call of ListOrganizations.
func (mr *MockUserRepositoryMockRecorder) ListOrganizations(pagenation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizations", reflect.TypeOf((*MockUserRepository)(nil).ListOrganizations), pagenation)
}

// PostAnswer mocks base method.
func (m *MockUserRepository) PostAnswer(answer domain.Answers, question int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostAnswer", answer, question)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostAnswer indicates an expected call of PostAnswer.
func (mr *MockUserRepositoryMockRecorder) PostAnswer(answer, question interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostAnswer", reflect.TypeOf((*MockUserRepository)(nil).PostAnswer), answer, question)
}

// PostQuestion mocks base method.
func (m *MockUserRepository) PostQuestion(question domain.Faqas) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostQuestion", question)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostQuestion indicates an expected call of PostQuestion.
func (mr *MockUserRepositoryMockRecorder) PostQuestion(question interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostQuestion", reflect.TypeOf((*MockUserRepository)(nil).PostQuestion), question)
}

// Prmotion_Faliure mocks base method.
func (m *MockUserRepository) Prmotion_Faliure(orderid, paymentid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prmotion_Faliure", orderid, paymentid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prmotion_Faliure indicates an expected call of Prmotion_Faliure.
func (mr *MockUserRepositoryMockRecorder) Prmotion_Faliure(orderid, paymentid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prmotion_Faliure", reflect.TypeOf((*MockUserRepository)(nil).Prmotion_Faliure), orderid, paymentid)
}

// Prmotion_Success mocks base method.
func (m *MockUserRepository) Prmotion_Success(orderid, paymentid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prmotion_Success", orderid, paymentid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prmotion_Success indicates an expected call of Prmotion_Success.
func (mr *MockUserRepositoryMockRecorder) Prmotion_Success(orderid, paymentid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prmotion_Success", reflect.TypeOf((*MockUserRepository)(nil).Prmotion_Success), orderid, paymentid)
}

// PromoteEvent mocks base method.
func (m *MockUserRepository) PromoteEvent(promotion domain.Promotion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PromoteEvent", promotion)
	ret0, _ := ret[0].(error)
	return ret0
}

// PromoteEvent indicates an expected call of PromoteEvent.
func (mr *MockUserRepositoryMockRecorder) PromoteEvent(promotion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PromoteEvent", reflect.TypeOf((*MockUserRepository)(nil).PromoteEvent), promotion)
}

// StoreVerificationDetails mocks base method.
func (m *MockUserRepository) StoreVerificationDetails(email, code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreVerificationDetails", email, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreVerificationDetails indicates an expected call of StoreVerificationDetails.
func (mr *MockUserRepositoryMockRecorder) StoreVerificationDetails(email, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreVerificationDetails", reflect.TypeOf((*MockUserRepository)(nil).StoreVerificationDetails), email, code)
}

// UpdatePassword mocks base method.
func (m *MockUserRepository) UpdatePassword(password, username string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", password, username)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserRepositoryMockRecorder) UpdatePassword(password, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserRepository)(nil).UpdatePassword), password, username)
}

// UpdateProfile mocks base method.
func (m *MockUserRepository) UpdateProfile(user domain.Bios, user_id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", user, user_id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockUserRepositoryMockRecorder) UpdateProfile(user, user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserRepository)(nil).UpdateProfile), user, user_id)
}

// UpdateRole mocks base method.
func (m *MockUserRepository) UpdateRole(user_id, organizaiton_id int, updatedRole string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", user_id, organizaiton_id, updatedRole)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockUserRepositoryMockRecorder) UpdateRole(user_id, organizaiton_id, updatedRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockUserRepository)(nil).UpdateRole), user_id, organizaiton_id, updatedRole)
}

// VerifyAccount mocks base method.
func (m *MockUserRepository) VerifyAccount(email, code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyAccount", email, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyAccount indicates an expected call of VerifyAccount.
func (mr *MockUserRepositoryMockRecorder) VerifyAccount(email, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAccount", reflect.TypeOf((*MockUserRepository)(nil).VerifyAccount), email, code)
}
