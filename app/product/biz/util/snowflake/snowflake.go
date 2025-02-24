package snoyflake

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineId uint16
)

func getMachineId() (id uint16, err error) {
	return sonyMachineId, nil
}
func Init(starttime string, machineId uint16) (err error) {
	sonyMachineId = machineId
	t, _ := time.Parse("2006-01-02", starttime) // 设置开始时间
	setting := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineId,
	}
	sonyFlake = sonyflake.NewSonyflake(setting) //用配置生成sonyflake节点
	return
}

// GetID 返回生成的id
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sonyflake not init")
		return 0, err
	}
	return sonyFlake.NextID()
}
