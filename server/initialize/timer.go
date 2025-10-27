package initialize

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/task"
	"github.com/robfig/cron/v3"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB) // 定时任务方法定在task文件包中
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "定时清理数据库【日志，黑名单】内容", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 其他定时任务定在这里 参考上方使用方法
		// 2. 新增：每 10 分钟同步 freqtrade 策略
		_, err = global.GVA_Timer.AddTaskByFunc("FreqStrategySync", "0 */10 * * * *", func() {
			new(system.FreqStrategyService).SyncListStrategies(context.Background())
		}, "每 10 分钟同步 freqtrade list-strategies 结果", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 3. 新增：每 5 分钟批量检查用户API并同步USDT永续资金
		_, err = global.GVA_Timer.AddTaskByFunc("UserApiBatchCheck", "0 */5 * * * *", func() {
			service.ServiceGroupApp.RunningServiceGroup.SysUserApiService.BatchCheckAndSyncApis(context.Background())
		}, "每 5 分钟批量检查用户API并同步USDT永续资金", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 4. 新增：每 10 分钟自动为合格用户启动 freqtrade 交易机器人
		_, err = global.GVA_Timer.AddTaskByFunc("AibotAutoRun", "0 */10 * * * *", func() {
			service.ServiceGroupApp.RunningServiceGroup.SysUserAibotService.AutoRunFreqtradeBots(context.Background())
		}, "每 10 分钟扫描可跑策略用户并按risk启动freqtrade", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		//_, err := global.GVA_Timer.AddTaskByFunc("定时任务标识", "corn表达式", func() {
		//	具体执行内容...
		//  ......
		//}, option...)
		//if err != nil {
		//	fmt.Println("add timer error:", err)
		//}
	}()
}
