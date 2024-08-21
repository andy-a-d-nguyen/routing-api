package main_test

import (
	testHelpers "code.cloudfoundry.org/routing-api/test_helpers"
	"fmt"
	"net/http"
	"os"
	"time"

	routingAPI "code.cloudfoundry.org/routing-api"
	"code.cloudfoundry.org/routing-api/db"
	"code.cloudfoundry.org/routing-api/models"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/ghttp"
	"github.com/tedsuo/ifrit"
	ginkgomon "github.com/tedsuo/ifrit/ginkgomon_v2"
)

const (
	TokenKeyEndpoint       = "/token_key"
	OpenidConfigEndpoint   = "/.well-known/openid-configuration"
	DefaultRouterGroupName = "default-tcp"
)

var _ = Describe("Main", func() {
	var (
		session        *Session
		routingAPIArgs testHelpers.Args
		configFilePath string
	)

	BeforeEach(func() {
		oAuthServer.Reset()

		routingAPIConfig := testHelpers.GetRoutingAPIConfig(defaultConfig)
		configFilePath = testHelpers.WriteConfigToTempFile(routingAPIConfig)
		routingAPIArgs = testHelpers.Args{
			IP:         testHelpers.RoutingAPIIP,
			ConfigPath: configFilePath,
			DevMode:    true,
		}
	})
	AfterEach(func() {
		if session != nil {
			Eventually(session.Kill()).Should(Exit())
		}

		err := os.RemoveAll(configFilePath)
		Expect(err).ToNot(HaveOccurred())
	})

	It("exits 1 if no config file is provided", func() {
		session = testHelpers.RoutingAPISession(routingAPIBinPath)
		Eventually(session).Should(Exit(1))
		Eventually(session).Should(Say("No configuration file provided"))
	})

	It("exits 1 if no ip address is provided", func() {
		session = testHelpers.RoutingAPISession(routingAPIBinPath, "-config=../../example_config/example.yml")
		Eventually(session).Should(Exit(1))
		Eventually(session).Should(Say("No ip address provided"))
	})

	It("exits 1 if the uaa verification key is not a valid PEM format", func() {
		oAuthServer.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", OpenidConfigEndpoint),
				ghttp.RespondWith(http.StatusOK, `{"issuer": "https://uaa.domain.com"}`),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", TokenKeyEndpoint),
				ghttp.RespondWith(http.StatusOK, `{"alg":"rsa", "value": "Invalid PEM key" }`),
			),
		)
		args := routingAPIArgs
		args.DevMode = false
		session = testHelpers.RoutingAPISession(routingAPIBinPath, args.ArgSlice()...)
		Eventually(session).Should(Exit(1))
		Eventually(session).Should(Say("Public uaa token must be PEM encoded"))
	})

	It("exits 1 if the uaa issuer cannot be fetched on startup and non dev mode", func() {
		oAuthServer.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", OpenidConfigEndpoint),
				ghttp.RespondWith(http.StatusInternalServerError, `{}`),
			),
		)
		args := routingAPIArgs
		args.DevMode = false
		session = testHelpers.RoutingAPISession(routingAPIBinPath, args.ArgSlice()...)
		Eventually(session).Should(Exit(1))
		Eventually(session).Should(Say("Failed to get issuer configuration from UAA"))
	})

	It("logs the uaa issuer when successfully fetched on startup and non dev mode", func() {
		oAuthServer.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", OpenidConfigEndpoint),
				ghttp.RespondWith(http.StatusOK, `{"issuer": "https://uaa.domain.com"}`),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", TokenKeyEndpoint),
				ghttp.RespondWith(http.StatusInternalServerError, `{}`),
			),
		)
		args := routingAPIArgs
		args.DevMode = false
		session = testHelpers.RoutingAPISession(routingAPIBinPath, args.ArgSlice()...)
		Eventually(session).Should(Say("received-issuer"))
		Eventually(session).Should(Say("https://uaa.domain.com"))
		Eventually(session).Should(Exit(1))
	})

	It("exits 1 if the uaa_verification_key cannot be fetched on startup and non dev mode", func() {
		oAuthServer.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", OpenidConfigEndpoint),
				ghttp.RespondWith(http.StatusOK, `{"issuer": "https://uaa.domain.com"}`),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", TokenKeyEndpoint),
				ghttp.RespondWith(http.StatusInternalServerError, `{}`),
			),
		)
		args := routingAPIArgs
		args.DevMode = false
		session = testHelpers.RoutingAPISession(routingAPIBinPath, args.ArgSlice()...)
		Eventually(session).Should(Exit(1))
		Eventually(session).Should(Say("Failed to get verification key from UAA"))
	})

	It("exits 1 if the SQL db fails to initialize", func() {
		session := testHelpers.RoutingAPISession(routingAPIBinPath, "-config=../../example_config/example.yml", "-ip='1.1.1.1'")
		Eventually(session).Should(Exit(1))
		Eventually(session).Should(Say("failed-initialize-sql-connection"))
	})

	Context("when initialized correctly and db is running", func() {
		BeforeEach(func() {
			oAuthServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", OpenidConfigEndpoint),
					ghttp.RespondWith(http.StatusOK, `{"issuer": "https://uaa.domain.com"}`),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", TokenKeyEndpoint),
					ghttp.RespondWith(http.StatusOK, `{"alg":"rsa", "value": "-----BEGIN PUBLIC KEY-----MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDHFr+KICms+tuT1OXJwhCUmR2dKVy7psa8xzElSyzqx7oJyfJ1JZyOzToj9T5SfTIq396agbHJWVfYphNahvZ/7uMXqHxf+ZH9BL1gk9Y6kCnbM5R60gfwjyW1/dQPjOzn9N394zd2FJoFHwdq9Qs0wBugspULZVNRxq7veq/fzwIDAQAB-----END PUBLIC KEY-----" }`),
				),
			)
		})

		It("unregisters from the db when the process exits", func() {
			routingAPIRunner := testHelpers.New(routingAPIBinPath, routingAPIArgs)
			proc := ifrit.Invoke(routingAPIRunner)

			routingAPIConfig := testHelpers.GetRoutingAPIConfig(defaultConfig)
			connectionString, err := db.ConnectionString(&routingAPIConfig.SqlDB)
			Expect(err).NotTo(HaveOccurred())
			gormDB, err := gorm.Open(routingAPIConfig.SqlDB.Type, connectionString)
			Expect(err).NotTo(HaveOccurred())

			getRoutes := func() string {
				var routes []models.Route
				err := gormDB.Find(&routes).Error
				Expect(err).ToNot(HaveOccurred())

				var routeUrl string
				if len(routes) > 0 {
					routeUrl = routes[0].Route
				}
				return routeUrl
			}
			Eventually(getRoutes).Should(ContainSubstring("api.example.com/routing"))

			ginkgomon.Interrupt(proc)

			Eventually(getRoutes).ShouldNot(ContainSubstring("api.example.com/routing"))
			Eventually(routingAPIRunner.ExitCode()).Should(Equal(0))
		})

		It("closes open event streams when the process exits", func() {
			routingAPIRunner := testHelpers.New(routingAPIBinPath, routingAPIArgs)
			proc := ifrit.Invoke(routingAPIRunner)
			client = routingAPI.NewClient(
				fmt.Sprintf("http://%s:%d", testHelpers.RoutingAPIIP, routingAPIPort),
				false,
			)

			events, err := client.SubscribeToEvents()
			Expect(err).ToNot(HaveOccurred())

			route := models.NewRoute("some-route", 1234, "234.32.43.4", "some-guid", "", 120)
			err = client.UpsertRoutes([]models.Route{route})
			Expect(err).ToNot(HaveOccurred())

			Eventually(func() string {
				event, _ := events.Next()
				return event.Action
			}).Should(Equal("Upsert"))

			ginkgomon.Interrupt(proc)

			Eventually(func() error {
				_, err = events.Next()
				return err
			}).Should(HaveOccurred())

			Eventually(routingAPIRunner.ExitCode(), 2*time.Second).Should(Equal(0))
		})
	})

	Context("when multiple router groups with the same name are seeded", func() {
		var (
			gormDB     *gorm.DB
			configPath string
		)
		BeforeEach(func() {
			routingAPIConfig := testHelpers.GetRoutingAPIConfig(defaultConfig)
			routingAPIConfig.RouterGroups = models.RouterGroups{
				models.RouterGroup{
					Name:            "default-tcp",
					Type:            "tcp",
					ReservablePorts: "2000",
				},
				models.RouterGroup{
					Name:            "default-tcp",
					Type:            "tcp",
					ReservablePorts: "10000-65535",
				},
			}
			configPath = testHelpers.WriteConfigToTempFile(routingAPIConfig)
			routingAPIArgs = testHelpers.Args{
				IP:         testHelpers.RoutingAPIIP,
				ConfigPath: configPath,
				DevMode:    true,
			}
			connectionString, err := db.ConnectionString(&routingAPIConfig.SqlDB)
			Expect(err).NotTo(HaveOccurred())
			gormDB, err = gorm.Open(routingAPIConfig.SqlDB.Type, connectionString)
			Expect(err).NotTo(HaveOccurred())
		})
		AfterEach(func() {
			gormDB.AutoMigrate(&models.RouterGroupDB{})
			Expect(os.Remove(configPath)).To(Succeed())
		})
		It("should fail with an error", func() {
			routingAPIRunner := testHelpers.New(routingAPIBinPath, routingAPIArgs)
			proc := ifrit.Invoke(routingAPIRunner)

			db := gormDB.Raw("SHOW TABLES LIKE 'router_groups';")
			Expect(db.Error).ToNot(HaveOccurred())
			Expect(int(db.RowsAffected)).To(Equal(0))
			Eventually(routingAPIRunner).Should(Exit(1))
			ginkgomon.Interrupt(proc)
		})
	})
})
