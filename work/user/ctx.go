package user

type Ctx interface {
	Get(userId string) (*User, error)
	GetUserinfo(code string) (*UserinfoResponse, error)
	SimpleList(departmentId, fetchChild int) (*SimpleListResponse, error)
	List(departmentId, fetchChild int) (*ListResponse, error)
	Delete(userId string) error
}

type ctx struct {
	tokenFunc func() (string, error)
}

func New(t func() (string, error)) Ctx {
	return &ctx{
		tokenFunc: t,
	}
}

//get 读取成员
func (u *ctx) Get(userId string) (*User, error) {
	t, err := u.tokenFunc()
	if err != nil {
		return nil, err
	}
	return get(t, userId)
}

//GetUserinfo 获取访问用户身份
func (u *ctx) GetUserinfo(code string) (*UserinfoResponse, error) {
	t, err := u.tokenFunc()
	if err != nil {
		return nil, err
	}
	return getUserinfo(t, code)
}

//SimpleList 获取访问用户身份
func (u *ctx) SimpleList(departmentId, fetchChild int) (*SimpleListResponse, error) {
	t, err := u.tokenFunc()
	if err != nil {
		return nil, err
	}
	return simpleList(t, departmentId, fetchChild)
}

//List 获取部门成员详情
func (u *ctx) List(departmentId, fetchChild int) (*ListResponse, error) {
	t, err := u.tokenFunc()
	if err != nil {
		return nil, err
	}
	return list(t, departmentId, fetchChild)
}

//delete 删除成员
func (u *ctx) Delete(userId string) error {
	t, err := u.tokenFunc()
	if err != nil {
		return err
	}
	return delete(t, userId)
}
