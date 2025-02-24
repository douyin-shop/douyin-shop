package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/constant"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/kafka"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/redis"
	"github.com/douyin-shop/douyin-shop/app/order/biz/utils/code"
	"github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/mysql"
	userModel "github.com/douyin-shop/douyin-shop/app/user/biz/dal/model"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	//判断订单是否合法
	if !IsPlaceOrderReqValid(req) {
		return nil, kerrors.NewBizStatusError(code.InvalidRequest, code.GetMsg(code.InvalidRequest))
	}
	//生成对应的订单Id
	orderId := generateOrderToken(req.GetUserId())
	placeOrderKey := getPlaceOrderKey(orderId)
	//创建锁实例对象
	lock := redis.NewDistributeRedisLock(placeOrderKey, 10, orderId)
	//尝试获取锁
	isLock, err := lock.TryLock()
	if err != nil {
		return nil, kerrors.NewBizStatusError(code.LockError, code.GetMsg(code.LockError))
	}
	if !isLock {
		return nil, kerrors.NewBizStatusError(code.AcquireLockFailed, code.GetMsg(code.AcquireLockFailed))
	}
	//扣减库存
	//1、预先扣减Redis中的库存数据
	items := req.GetOrderItems()
	for _, item := range items {
		//扣减Redis中的库存
		//1、扣减Redis中的库存数据
		decreaseStockKey := getDecreaseStockKey(item.Item.ProductId)
		decrease := redis.NewStockDecrease(decreaseStockKey, item.Item.Quantity)
		if tryDecrease, err := decrease.TryDecrease(); err != nil || tryDecrease == false {
			return nil, kerrors.NewBizStatusError(code.DecreaseStockError, code.GetMsg(code.DecreaseStockError))
		}
		//2、异步扣减数据库中的库存数据
		//将扣减库存的请求发送到消息队列中
		kafka.SendInventoryMessage(kafka.GetProducer(), item.Item.ProductId, item.Item.Quantity)
	}

	//todo 异步扣减用户余额 异步扣减用户积分 异步扣减用户优惠券 异步扣减用户积分 异步扣减用户积分

	//生成订单信息
	order := createOrder(req, orderId)
	db := mysql.DB
	//持久化订单信息到数据库
	if err := db.Create(&order).Error; err != nil {
		return nil, kerrors.NewBizStatusError(code.CreateOrderError, code.GetMsg(code.CreateOrderError))
	}
	//释放锁
	lock.Unlock()
	return
}

func createOrder(req *order.PlaceOrderReq, orderId string) model.Order {
	items := req.GetOrderItems()
	OrderItemIds := []uint32{}
	for _, item := range items {
		OrderItemIds = append(OrderItemIds, item.Item.ProductId)
	}
	db := mysql.DB
	var user userModel.User
	db.Where("id = ?", req.UserId).First(&user)
	reqAddress := req.Address
	address := model.Address{
		StreetAddress: reqAddress.StreetAddress,
		City:          reqAddress.City,
		State:         reqAddress.State,
		Country:       reqAddress.Country,
		ZipCode:       reqAddress.ZipCode,
	}
	order := model.Order{
		Model:           gorm.Model{},
		OrderId:         orderId,
		OrderItemIdList: OrderItemIds,
		TotalAmount:     0,
		OrderStatus:     constant.Order_Created,
		UserId:          req.UserId,
		Phone:           user.Phone,
		Email:           req.Email,
		Address:         address,
		PlaceOrderTime:  time.Now(),
	}
	return order
}

func IsPlaceOrderReqValid(req *order.PlaceOrderReq) bool {
	// 检查 UserId 是否有效
	if req.UserId == 0 {
		return false
	}

	// 检查 UserCurrency 是否有效
	if req.UserCurrency == "" {
		return false
	}

	// 检查 Address 是否有效
	if req.Address == nil || req.Address.StreetAddress == "" || req.Address.City == "" || req.Address.State == "" || req.Address.Country == "" || req.Address.ZipCode == 0 {
		return false
	}

	// 检查 Email 是否有效
	if req.Email == "" || !isValidEmail(req.Email) {
		return false
	}

	// 检查 OrderItems 是否有效
	if req.OrderItems == nil || len(req.OrderItems) == 0 {
		return false
	}

	// 检查每个 OrderItem 是否有效
	for _, item := range req.OrderItems {
		if item.Item == nil || item.Cost <= 0 {
			return false
		}
	}

	return true
}

// 辅助函数：检查电子邮件地址是否有效
func isValidEmail(email string) bool {
	// 这里可以使用正则表达式或其他方法来验证电子邮件地址的有效性
	// 为了简单起见，我们只检查是否包含 @ 符号
	return strings.Contains(email, "@")
}

func getPlaceOrderKey(orderToken string) string {
	return constant.PLACE_ORDER_LOCK + orderToken
}

func getDecreaseStockKey(ProductId uint32) string {
	return constant.PLACE_ORDER_LOCK + strconv.Itoa(int(ProductId))
}

func generateOrderToken(userId uint32) string {
	// 用户id前补零保证五位，对超出五位的保留后五位
	userIdFilledZero := fmt.Sprintf("%05d", int64(userId))
	fiveDigitsUserId := userIdFilledZero[len(userIdFilledZero)-5:]

	// 生成3位随机数
	random := rand.Intn(1000)

	// 将时间戳+3位随机数+五位id组成商户订单号，规则参考自"https://tech.meituan.com/2016/11/18/dianping-order-db-sharding.html"大众点评
	return time.Now().Format("2023120514345123") + fmt.Sprintf("%03d", random) + fiveDigitsUserId
}
