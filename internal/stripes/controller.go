package stripes

import (
	"main/internal/utility"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/webhook"
)

type controller struct {
	service StripeService
}

func RegisterRoutes(router *gin.Engine, service StripeService) {
	c := controller{service}
	s := router.Group("/stripes")
	s.POST("/webhook", c.HandleStripeWebhook)
}

func (con controller) HandleStripeWebhook(c *gin.Context) {
	var err error
	defer utility.ErrorHandleHTTP(c, err)
	body, _ := c.GetRawData()
	event, err := webhook.ConstructEvent(body, c.Request.Header.Get("Stripe-Signature"), os.Getenv("STRIPE_WEBHOOK_SECRET"))
	if err != nil {
		return
	}
	con.service.HandleStripeWebhook(event)
}
