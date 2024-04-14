package data

import (
	"context"

	"helloworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type realWorldRepo struct {
	data *Data
	log  *log.Helper
}

// NewRealWorldRepo .
func NewRealWorldRepo(data *Data, logger log.Logger) biz.RealWorldRepo {
	return &realWorldRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *realWorldRepo) Save(ctx context.Context, g *biz.RealWorld) (*biz.RealWorld, error) {
	return g, nil
}

func (r *realWorldRepo) Update(ctx context.Context, g *biz.RealWorld) (*biz.RealWorld, error) {
	return g, nil
}

func (r *realWorldRepo) FindByID(context.Context, int64) (*biz.RealWorld, error) {
	return nil, nil
}

func (r *realWorldRepo) ListByHello(context.Context, string) ([]*biz.RealWorld, error) {
	return nil, nil
}

func (r *realWorldRepo) ListAll(context.Context) ([]*biz.RealWorld, error) {
	return nil, nil
}
