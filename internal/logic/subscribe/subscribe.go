package subscribe

import (
	"context"
	v1 "leke/api/user/v1"
	"leke/internal/dao"
	"leke/internal/service"
	"github.com/gogf/gf/v2/os/gtime"
)

type sSubscribe struct{}

func New() *sSubscribe {
	return &sSubscribe{}
}

func init() {
	service.RegisterSubscribe(New())
}

// SubscribeList 获取关注列表
func (s *sSubscribe) SubscribeList(ctx context.Context, req *v1.SubscribeListReq) (res *v1.SubscribeListRes, err error) {
	// 创建数据库查询模型
	m := dao.Subscribe.Ctx(ctx)

	// 根据用户ID查询
	m = m.Where(dao.Subscribe.Columns().UserId, req.UserId)

	// 根据状态查询
	if req.Status != 0 {
		m = m.Where(dao.Subscribe.Columns().Status, req.Status)
	}

	// 根据被关注用户ID查询
	if req.FollowId != 0 {
		m = m.Where(dao.Subscribe.Columns().FollowId, req.FollowId)
	}

	// 查询总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页处理
	m = m.Page(req.Page, req.PageSize)

	// 查询列表数据
	var subscribeList []*v1.SubscribeListItem
	err = m.Scan(&subscribeList)
	if err != nil {
		return nil, err
	}

	res = &v1.SubscribeListRes{
		List: subscribeList,
	}

	// 设置分页信息
	req.PageResult.Total = int(total)
	req.PageResult.Page = req.Page
	req.PageResult.PageSize = req.PageSize

	return
}

// SubscribeView 获取关注详情
func (s *sSubscribe) SubscribeView(ctx context.Context, req *v1.SubscribeViewReq) (res *v1.SubscribeViewRes, err error) {
	// 创建数据库查询模型
	m := dao.Subscribe.Ctx(ctx)

	// 根据ID查询
	m = m.Where(dao.Subscribe.Columns().Id, req.Id)

	res = &v1.SubscribeViewRes{}
	err = m.Scan(&res.SubscribeListItem)
	if err != nil {
		return nil, err
	}

	return
}

// SubscribeUpdate 更新关注信息
func (s *sSubscribe) SubscribeUpdate(ctx context.Context, req *v1.SubscribeUpdateReq) (res *v1.SubscribeUpdateRes, err error) {
	// 创建数据库查询模型
	m := dao.Subscribe.Ctx(ctx)

	// 更新关注信息
	_, err = m.Data(req).Where(dao.Subscribe.Columns().Id, req.Id).Update()
	if err != nil {
		return nil, err
	}

	res = &v1.SubscribeUpdateRes{
		Id: req.Id,
	}

	return
}

// SubscribeDelete 删除关注关系
func (s *sSubscribe) SubscribeDelete(ctx context.Context, req *v1.SubscribeDeleteReq) (res *v1.SubscribeDeleteRes, err error) {
	// 根据ID删除关注关系
	_, err = dao.Subscribe.Ctx(ctx).Where(dao.Subscribe.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, err
	}

	res = &v1.SubscribeDeleteRes{}
	return
}

// SubscribeCreate 创建关注关系
func (s *sSubscribe) SubscribeCreate(ctx context.Context, req *v1.SubscribeCreateReq) (res *v1.SubscribeCreateRes, err error) {
	// 创建数据库查询模型
	m := dao.Subscribe.Ctx(ctx)

	// 准备数据
	data := map[string]interface{}{
		"user_id":    req.UserId,
		"follow_id":  req.FollowId,
		"status":     1, // 默认状态为关注中
		"remark":     req.Remark,
		"created_at": gtime.Now(),
		"updated_at": gtime.Now(),
	}

	// 插入数据
	result, err := m.Data(data).Insert()
	if err != nil {
		return nil, err
	}

	// 获取插入的ID
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	res = &v1.SubscribeCreateRes{
		Id: uint64(lastInsertId),
	}

	return
}