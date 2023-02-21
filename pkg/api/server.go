package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/SethukumarJ/Events_Radar_Developement/cmd/api/docs"
	handler "github.com/SethukumarJ/Events_Radar_Developement/pkg/api/handler"
	middleware "github.com/SethukumarJ/Events_Radar_Developement/pkg/api/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler handler.UserHandler,
	authHandler handler.AuthHandler,
	adminHandler handler.AdminHandler,
	eventHandler handler.EventHandler,
	middleware middleware.Middleware) *ServerHTTP {
	authHandler.InitializeOAuthGoogle()
	engine := gin.Default()

	// Use logger from Gin
	engine.Use(gin.Logger())
	// engine.Static("/templates", "./api/templates")

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// engine.LoadHTMLFiles("templates/razorpay.html")
	

	engine.GET("/", func(c *gin.Context) {
		html := `
			<html>
				<img src="https://images.all-free-download.com/images/graphiclarge/delicious_fruie_03_hd_pictures_166657.jpg" alt="alternatetext">
				<h4>Pay 12</h4>
	
				<button id="rzp-button1">Pay</button>
				<script src="https://checkout.razorpay.com/v1/checkout.js"></script>
				<script>
				var options = {
					"key": "rzp_test_kEtg65WKqGTpKd", // Enter the Key ID generated from the Dashboard
					"amount": "{{.Amount}}", // Amount is in currency subunits. Default currency is INR. Hence, 50000 refers to 50000 paise
					"currency": "INR",
					"name": "Events-Radar",
					"description": "Promote post",
					"image": "https://example.com/your_logo",
					"order_id": "{{.OrderId}}", //This is a sample Order ID. Pass the id obtained in the response of Step 1
					"handler": function (response){
						a = (response.razorpay_payment_id);
						b =(response.razorpay_order_id);
						c = (response.razorpay_signature)
						window.location.replace("http:/payment-success?paymentid="+a+"&orderid=+"+b+"&signature="+c);
					},
					"prefill": {
						"name": "{{.Name}}",
						"email": "{{.Email}}",
						"contact": "{{.Contact}}"
					},
					"notes": {
						"address": "Razorpay Corporate Office"
					},
					"theme": {
						"color": "#3399cc"
					}
				};
				var rzp1 = new Razorpay(options);
				rzp1.on('payment.failed', function (response){
						(response.error.code);
						(response.error.description);
						(response.error.source);
						(response.error.step);
						(response.error.reason);
						(response.error.metadata.order_id);
						(response.error.metadata.payment_id);
				});
				document.getElementById('rzp-button1').onclick = function(e){
					rzp1.open();
					e.preventDefault();
				}
				</script>
			</html>
		`
		c.Data(http.StatusOK, "text/html", []byte(html))
	})
	

	//User routes
	user := engine.Group("user")
	{
		user.GET("/login-gl", authHandler.GoogleAuth)
		user.GET("/callback-gl", authHandler.CallBackFromGoogle)
		user.POST("/signup", authHandler.UserSignup)
		user.POST("/login", authHandler.UserLogin)
		user.POST("/send-verification", userHandler.SendVerificationMail)
		user.GET("/verify-account", authHandler.VerifyAccount)
		user.PATCH("/update-password", userHandler.UpdatePassword)
		user.GET("/list-faqas", userHandler.GetPublicFaqas)
		user.GET("/list-organizations", userHandler.ListOrganizations)
		user.GET("/list/approved-events", eventHandler.ViewAllApprovedEvents)
		user.GET("/geteventbytitle", eventHandler.GetEventByTitle)
		user.Use(middleware.AuthorizeJwt())
		{
			user.POST("/token-refresh", authHandler.UserRefreshToken)
			user.POST("/apply-event", userHandler.ApplyEvent)
			user.POST("/event/post-question", userHandler.PostQuestion)
			user.PATCH("/update-profile", userHandler.UpdateProfile)
			user.POST("/create-organization", userHandler.CreateOrganization)
			user.POST("/create-event", eventHandler.CreateEventUser)
			user.PATCH("/join-organization", userHandler.JoinOrganization)
		}

	}

	//Admin routes
	admin := engine.Group("admin")
	{
		admin.POST("/signup", authHandler.AdminSignup)
		admin.POST("/login", authHandler.AdminLogin)

		admin.Use(middleware.AuthorizeJwt())
		{
			admin.GET("/token/refresh", authHandler.AdminRefreshToken)
			admin.PATCH("/approve-event", adminHandler.ApproveEvent)
			admin.POST("/create-event", eventHandler.CreateEventAdmin)
			admin.GET("/list-users", adminHandler.ViewAllUsers)
			admin.PATCH("/make/vip-user", adminHandler.VipUser)
			admin.GET("/list-events", adminHandler.ViewAllEvents)
			admin.PATCH("/register-organization", adminHandler.RegisterOrganization)
			admin.PATCH("/reject-organization", adminHandler.RejectOrganization)
			admin.GET("/list-organizations", adminHandler.ListOrgRequests)
		}
	}

	//Organization routes
	organization := engine.Group("organization")
	{
		organization.GET("/get-organization", userHandler.GetOrganization)
		organization.Use(middleware.AuthorizeOrg())
		{

			organization.PATCH("/update-event", eventHandler.UpdateEvent)
			organization.DELETE("/delete-event", eventHandler.DeleteEvent)
			organization.GET("/event/list-questions", userHandler.GetQuestions)
			organization.POST("/event/post-answer", userHandler.PostAnswer)
			organization.POST("/create-event", eventHandler.CreateEventOrganization)
			organization.POST("/admin/add-members", userHandler.AddMembers)
			organization.GET("/accept-invitation", userHandler.AcceptJoinInvitation)
			organization.PATCH("/admin/admit-member", userHandler.AdmitMember)
			organization.GET("/join-requests", userHandler.ListJoinRequests)

		}

	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
