/**
 * @author sd_sase
 * @version 1.0.0
 * @copyright www.sd-sase.com
 * @e-mail sd_sase@163.com
 * @2021/11/14
 */

package main

import (
	idoictos_mfa_service "idoictos_mfa/src/service"

	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"moul.io/banner"
)

var hosts = "0.0.0.0:22001"

func main() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	fmt.Println("starting SD-SASE IDoictOS ...")
	fmt.Println()
	fmt.Println(banner.Inline("sd-sase"))
	fmt.Println()
	fmt.Println("SD-SASE史诗级的前世今生：")
	fmt.Println()
	fmt.Println("1、SD-SASE愿景：SD-SASE是一个完全开源、免费、自主可控、开放、中立、新一代DOICT零信任云原生操作系统。 优势：内置零信任、云原生、安全、身份、计算、存储、网络、边缘计算、软件供应链安全、服务网格、低代码、应用市场，...... 价值：由开源社区、开发者、贡献者、生态合作伙伴等共同驱动的新一代开源分布式云原生操作系统。")
	fmt.Println()
	fmt.Println("2、SD-SASE官方网站：https://www.sd-sase.com")
	fmt.Println()
	fmt.Println("3、SD-SASE开源社区：https://www.github.com/sd-sase")
	fmt.Println()
	fmt.Println("4、SD-SASE微信公众号：哈希安全科技")
	fmt.Println()
	fmt.Println("5、SD-SASE开源基金会与开源社区推进委员会：北京哈希安全科技有限公司")
	fmt.Println()
	fmt.Println("6、SD-SASE创始人：宋园园")
	fmt.Println()
	fmt.Println("7、SD-SASE成立时间：2021-05-15")
	fmt.Println()
	fmt.Println("8、SD-SASE成立地址：北京")
	fmt.Println()
	fmt.Println("9、SD-SASE版权所有：北京哈希安全科技有限公司")
	fmt.Println()

	router.GET("/idoictos/idaas/mfa/keepalived", idoictos_mfa_service.Keepalived)
	router.POST("/idoictos/idaas", idoictos_mfa_service.Post)
	router.PUT("/idoictos/idaas", idoictos_mfa_service.Put)
	router.PATCH("/idoictos/idaas", idoictos_mfa_service.Patch)
	router.DELETE("/idoictos/idaas", idoictos_mfa_service.Delete)
	router.GET("/idoictos/idaas", idoictos_mfa_service.Get)

	time.Sleep(3 * time.Microsecond)
	fmt.Println()
	fmt.Println(banner.Inline("idoictos"))
	fmt.Println()
	fmt.Println("started SD-SASE IDoictOS MFA:", hosts)
	fmt.Println()

	router.Run(hosts)
}
