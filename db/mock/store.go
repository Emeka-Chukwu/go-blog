// Code generated by MockGen. DO NOT EDIT.
// Source: blog-api/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	db "blog-api/db/sqlc"
	context "context"
	sql "database/sql"
	reflect "reflect"

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

// CreateCategory mocks base method.
func (m *MockStore) CreateCategory(arg0 context.Context, arg1 db.CreateCategoryParams) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockStoreMockRecorder) CreateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockStore)(nil).CreateCategory), arg0, arg1)
}

// CreateComment mocks base method.
func (m *MockStore) CreateComment(arg0 context.Context, arg1 db.CreateCommentParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockStoreMockRecorder) CreateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockStore)(nil).CreateComment), arg0, arg1)
}

// CreatePost mocks base method.
func (m *MockStore) CreatePost(arg0 context.Context, arg1 db.CreatePostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockStoreMockRecorder) CreatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockStore)(nil).CreatePost), arg0, arg1)
}

// CreateTags mocks base method.
func (m *MockStore) CreateTags(arg0 context.Context, arg1 db.CreateTagsParams) (db.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTags", arg0, arg1)
	ret0, _ := ret[0].(db.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTags indicates an expected call of CreateTags.
func (mr *MockStoreMockRecorder) CreateTags(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTags", reflect.TypeOf((*MockStore)(nil).CreateTags), arg0, arg1)
}

// CreateTagsToPost mocks base method.
func (m *MockStore) CreateTagsToPost(arg0 context.Context, arg1 db.CreateTagsToPostParams) (db.PostTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTagsToPost", arg0, arg1)
	ret0, _ := ret[0].(db.PostTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTagsToPost indicates an expected call of CreateTagsToPost.
func (mr *MockStoreMockRecorder) CreateTagsToPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTagsToPost", reflect.TypeOf((*MockStore)(nil).CreateTagsToPost), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteCategory mocks base method.
func (m *MockStore) DeleteCategory(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockStoreMockRecorder) DeleteCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockStore)(nil).DeleteCategory), arg0, arg1)
}

// DeleteComment mocks base method.
func (m *MockStore) DeleteComment(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockStoreMockRecorder) DeleteComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockStore)(nil).DeleteComment), arg0, arg1)
}

// DeletePosts mocks base method.
func (m *MockStore) DeletePosts(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePosts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePosts indicates an expected call of DeletePosts.
func (mr *MockStoreMockRecorder) DeletePosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePosts", reflect.TypeOf((*MockStore)(nil).DeletePosts), arg0, arg1)
}

// DeleteTag mocks base method.
func (m *MockStore) DeleteTag(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTag indicates an expected call of DeleteTag.
func (mr *MockStoreMockRecorder) DeleteTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTag", reflect.TypeOf((*MockStore)(nil).DeleteTag), arg0, arg1)
}

// DeleteTagsOfPost mocks base method.
func (m *MockStore) DeleteTagsOfPost(arg0 context.Context, arg1 sql.NullInt32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTagsOfPost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTagsOfPost indicates an expected call of DeleteTagsOfPost.
func (mr *MockStoreMockRecorder) DeleteTagsOfPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTagsOfPost", reflect.TypeOf((*MockStore)(nil).DeleteTagsOfPost), arg0, arg1)
}

// DissociatePostZFromTag mocks base method.
func (m *MockStore) DissociatePostZFromTag(arg0 context.Context, arg1 db.DissociatePostZFromTagParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DissociatePostZFromTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DissociatePostZFromTag indicates an expected call of DissociatePostZFromTag.
func (mr *MockStoreMockRecorder) DissociatePostZFromTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DissociatePostZFromTag", reflect.TypeOf((*MockStore)(nil).DissociatePostZFromTag), arg0, arg1)
}

// GetCategories mocks base method.
func (m *MockStore) GetCategories(arg0 context.Context) ([]db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories", arg0)
	ret0, _ := ret[0].([]db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockStoreMockRecorder) GetCategories(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockStore)(nil).GetCategories), arg0)
}

// GetCategoryById mocks base method.
func (m *MockStore) GetCategoryById(arg0 context.Context, arg1 int32) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryById", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryById indicates an expected call of GetCategoryById.
func (mr *MockStoreMockRecorder) GetCategoryById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryById", reflect.TypeOf((*MockStore)(nil).GetCategoryById), arg0, arg1)
}

// GetCommentById mocks base method.
func (m *MockStore) GetCommentById(arg0 context.Context, arg1 int32) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentById", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentById indicates an expected call of GetCommentById.
func (mr *MockStoreMockRecorder) GetCommentById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentById", reflect.TypeOf((*MockStore)(nil).GetCommentById), arg0, arg1)
}

// GetComments mocks base method.
func (m *MockStore) GetComments(arg0 context.Context, arg1 int32) ([]db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComments", arg0, arg1)
	ret0, _ := ret[0].([]db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComments indicates an expected call of GetComments.
func (mr *MockStoreMockRecorder) GetComments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockStore)(nil).GetComments), arg0, arg1)
}

// GetPostById mocks base method.
func (m *MockStore) GetPostById(arg0 context.Context, arg1 int32) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostById", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostById indicates an expected call of GetPostById.
func (mr *MockStoreMockRecorder) GetPostById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostById", reflect.TypeOf((*MockStore)(nil).GetPostById), arg0, arg1)
}

// GetPosts mocks base method.
func (m *MockStore) GetPosts(arg0 context.Context, arg1 db.GetPostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosts", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPosts indicates an expected call of GetPosts.
func (mr *MockStoreMockRecorder) GetPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*MockStore)(nil).GetPosts), arg0, arg1)
}

// GetTagId mocks base method.
func (m *MockStore) GetTagId(arg0 context.Context, arg1 int32) (db.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagId", arg0, arg1)
	ret0, _ := ret[0].(db.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagId indicates an expected call of GetTagId.
func (mr *MockStoreMockRecorder) GetTagId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagId", reflect.TypeOf((*MockStore)(nil).GetTagId), arg0, arg1)
}

// GetTags mocks base method.
func (m *MockStore) GetTags(arg0 context.Context) ([]db.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTags", arg0)
	ret0, _ := ret[0].([]db.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags.
func (mr *MockStoreMockRecorder) GetTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockStore)(nil).GetTags), arg0)
}

// GetUserByEmail mocks base method.
func (m *MockStore) GetUserByEmail(arg0 context.Context, arg1 sql.NullString) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockStoreMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockStore)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockStore) GetUserById(arg0 context.Context, arg1 int32) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockStoreMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockStore)(nil).GetUserById), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockStore) GetUsers(arg0 context.Context) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockStoreMockRecorder) GetUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockStore)(nil).GetUsers), arg0)
}

// ListPostWithCommentAndTags mocks base method.
func (m *MockStore) ListPostWithCommentAndTags(arg0 context.Context) ([]db.ListPostWithCommentAndTagsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPostWithCommentAndTags", arg0)
	ret0, _ := ret[0].([]db.ListPostWithCommentAndTagsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPostWithCommentAndTags indicates an expected call of ListPostWithCommentAndTags.
func (mr *MockStoreMockRecorder) ListPostWithCommentAndTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPostWithCommentAndTags", reflect.TypeOf((*MockStore)(nil).ListPostWithCommentAndTags), arg0)
}

// ListPostbyCategories mocks base method.
func (m *MockStore) ListPostbyCategories(arg0 context.Context, arg1 int32) ([]db.ListPostbyCategoriesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPostbyCategories", arg0, arg1)
	ret0, _ := ret[0].([]db.ListPostbyCategoriesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPostbyCategories indicates an expected call of ListPostbyCategories.
func (mr *MockStoreMockRecorder) ListPostbyCategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPostbyCategories", reflect.TypeOf((*MockStore)(nil).ListPostbyCategories), arg0, arg1)
}

// ListPostbyTag mocks base method.
func (m *MockStore) ListPostbyTag(arg0 context.Context, arg1 sql.NullInt32) ([]db.ListPostbyTagRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPostbyTag", arg0, arg1)
	ret0, _ := ret[0].([]db.ListPostbyTagRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPostbyTag indicates an expected call of ListPostbyTag.
func (mr *MockStoreMockRecorder) ListPostbyTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPostbyTag", reflect.TypeOf((*MockStore)(nil).ListPostbyTag), arg0, arg1)
}

// UpdateCategory mocks base method.
func (m *MockStore) UpdateCategory(arg0 context.Context, arg1 db.UpdateCategoryParams) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockStoreMockRecorder) UpdateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockStore)(nil).UpdateCategory), arg0, arg1)
}

// UpdateComment mocks base method.
func (m *MockStore) UpdateComment(arg0 context.Context, arg1 db.UpdateCommentParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockStoreMockRecorder) UpdateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockStore)(nil).UpdateComment), arg0, arg1)
}

// UpdatePost mocks base method.
func (m *MockStore) UpdatePost(arg0 context.Context, arg1 db.UpdatePostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockStoreMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockStore)(nil).UpdatePost), arg0, arg1)
}

// UpdateTag mocks base method.
func (m *MockStore) UpdateTag(arg0 context.Context, arg1 db.UpdateTagParams) (db.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTag", arg0, arg1)
	ret0, _ := ret[0].(db.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTag indicates an expected call of UpdateTag.
func (mr *MockStoreMockRecorder) UpdateTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTag", reflect.TypeOf((*MockStore)(nil).UpdateTag), arg0, arg1)
}

// UpdateTagsPost mocks base method.
func (m *MockStore) UpdateTagsPost(arg0 context.Context, arg1 db.UpdateTagsPostParams) (db.PostTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTagsPost", arg0, arg1)
	ret0, _ := ret[0].(db.PostTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTagsPost indicates an expected call of UpdateTagsPost.
func (mr *MockStoreMockRecorder) UpdateTagsPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTagsPost", reflect.TypeOf((*MockStore)(nil).UpdateTagsPost), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
