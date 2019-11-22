package material

type Ctx interface {
	Upload(mold, path string, params map[string]string) (*UploadResponse, error)
}

type ctx struct {
	tokenFunc func() (string, error)
}

func New(t func() (string, error)) Ctx {
	return &ctx{tokenFunc: t}
}

//Upload 上传永久素材
func (c *ctx) Upload(mold, path string, params map[string]string) (*UploadResponse, error) {
	t, err := c.tokenFunc()
	if err != nil {
		return nil, err
	}
	return upload(t, mold, path, params)
}
