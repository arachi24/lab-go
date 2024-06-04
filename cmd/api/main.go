package main

import (
	"example/config"
	"example/database"
	"example/internal/banner"
	"example/internal/business"
	"example/internal/director"
	"example/internal/homepage"
	invertorrelation "example/internal/invertor_relation"
	"example/internal/location"
	"example/internal/news"
	"example/internal/partner"
	"example/internal/product"
	productarea "example/internal/product_area"
	"example/internal/smart_solution"
	template "example/internal/template_data"
	"example/internal/upload"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/viper"
)

const idleTimeout = 10 * time.Second

func main() {

	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(cors.New())

	config.NewConfig()

	db := database.NewDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ARV BO API")
	})

	app.Get("/live", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "ok",
			"statusCode": 200,
		})
	})

	app.Get("/ready", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "ok",
			"statusCode": 200,
		})
	})
	v1 := app.Group("/v1")

	//repository

	homepageRepo := homepage.NewHomepageRepository(db)
	bannerRepo := banner.NewBannerRepository(db)
	directorRepo := director.NewDirectorRepository(db)
	businessRepo := business.NewBusinessRepository(db)
	investorRepo := invertorrelation.NewInvestorRelationRepository(db)
	locationRepo := location.NewLocationRepository(db)
	newseRepo := news.NewNewsRepository(db)
	partnerRepo := partner.NewPartnerRepository(db)
	productareaRepo := productarea.NewProductreaRepository(db)
	smartSolutionRepo := smart_solution.NewSmartSolutionRepository(db)
	productRepo := product.NewProductRepository(db)
	templateRepo := template.NewTemplateRepository(db)

	//routes

	homepageHandler := homepage.NewHomepageHandler(&homepageRepo, &bannerRepo)
	homepageAPI := v1.Group("/homepage")
	homepageAPI.Post("/", homepageHandler.Create)
	homepageAPI.Get("/", homepageHandler.FindAll)
	homepageAPI.Get("/:id", homepageHandler.FindOne)
	homepageAPI.Put("/:id", homepageHandler.Update)
	homepageAPI.Delete("/:id", homepageHandler.Delete)

	directorHandler := director.NewDirectorHandler(&directorRepo)
	directorAPI := v1.Group("/director")
	directorAPI.Post("/", directorHandler.Create)
	directorAPI.Get("/", directorHandler.FindAll)
	directorAPI.Get("/:id", directorHandler.FindOne)
	directorAPI.Put("/:id", directorHandler.Update)
	directorAPI.Delete("/:id", directorHandler.Delete)

	businessHandler := business.NewBusinessHandler(&businessRepo)
	businessAPI := v1.Group("/business")
	businessAPI.Post("/", businessHandler.Create)
	businessAPI.Get("/", businessHandler.FindAll)
	businessAPI.Get("/:id", businessHandler.FindOne)
	businessAPI.Put("/:id", businessHandler.Update)
	businessAPI.Delete("/:id", businessHandler.Delete)

	investorHandler := invertorrelation.NewInvestorRelationHandler(&investorRepo, &bannerRepo)
	investorAPI := v1.Group("/investor")
	investorAPI.Post("/", investorHandler.Create)
	investorAPI.Get("/", investorHandler.FindAll)
	investorAPI.Get("/:id", investorHandler.FindOne)
	investorAPI.Put("/:id", investorHandler.Update)
	investorAPI.Delete("/:id", investorHandler.Delete)

	locationHandler := location.NewLocationHandler(&locationRepo)
	locationAPI := v1.Group("/location")
	locationAPI.Post("/", locationHandler.Create)
	locationAPI.Get("/", locationHandler.FindAll)
	locationAPI.Get("/:id", locationHandler.FindOne)
	locationAPI.Put("/:id", locationHandler.Update)
	locationAPI.Delete("/:id", locationHandler.Delete)

	newsHandler := news.NewNewsHandler(&newseRepo)
	newsAPI := v1.Group("/news")
	newsAPI.Post("/", newsHandler.Create)
	newsAPI.Get("/", newsHandler.FindAll)
	newsAPI.Get("/:id", newsHandler.FindOne)
	newsAPI.Put("/:id", newsHandler.Update)
	newsAPI.Delete("/:id", newsHandler.Delete)

	partnerHandler := partner.NewPartnerHandler(&partnerRepo)
	partnerAPI := v1.Group("/partner")
	partnerAPI.Post("/", partnerHandler.Create)
	partnerAPI.Get("/", partnerHandler.FindAll)
	partnerAPI.Get("/:id", partnerHandler.FindOne)
	partnerAPI.Put("/:id", partnerHandler.Update)
	partnerAPI.Delete("/:id", partnerHandler.Delete)

	productAreaHandler := productarea.NewProductareaHandler(&productareaRepo)
	productAreaAPI := v1.Group("/product-area")
	productAreaAPI.Post("/", productAreaHandler.Create)
	productAreaAPI.Get("/", productAreaHandler.FindAll)
	productAreaAPI.Get("/:id", productAreaHandler.FindOne)
	productAreaAPI.Put("/:id", productAreaHandler.Update)
	productAreaAPI.Delete("/:id", productAreaHandler.Delete)

	smartSolutionHandler := smart_solution.NewSmartSolutionHandler(&smartSolutionRepo)
	smartApi := v1.Group("/smart-solution")
	smartApi.Post("/", smartSolutionHandler.Create)
	smartApi.Get("/", smartSolutionHandler.FindAll)
	smartApi.Get("/:id", smartSolutionHandler.FindOne)
	smartApi.Delete("/:id", smartSolutionHandler.Delete)
	smartApi.Put("/:id", smartSolutionHandler.Update)

	productHandler := product.NewProductHandler(&productRepo)
	productApi := v1.Group("/product")
	productApi.Post("/", productHandler.Create)
	productApi.Get("/", productHandler.FindAll)
	productApi.Get("/:id", productHandler.FindOne)
	productApi.Delete("/:id", productHandler.Delete)
	productApi.Put("/:id", productHandler.Update)

	templateHandler := template.NewTemplateHandler(&templateRepo)
	templateApi := v1.Group("/template")
	templateApi.Post("/", templateHandler.Create)
	templateApi.Get("/", templateHandler.FindOne)
	templateApi.Put("/:id", templateHandler.Update)

	uploadHandler := upload.NewUploadHandler()
	uploadApI := v1.Group("/upload")
	uploadApI.Post("/", uploadHandler.Create)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%v", viper.GetInt("PORT"))); err != nil {
			log.Panic(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	isShutdown := <-ch

	if isShutdown.String() == "interrupt" {
		fmt.Println("Gracefully shutting down...")
		fmt.Printf("%v\n", isShutdown.String())
		fmt.Println("Running cleanup tasks...")

		// Your cleanup tasks go here
		// database.CloseDB(db)
		fmt.Println("Fiber was disconnect database.")

		// redisConn.Close()
		fmt.Println("Fiber was successful shutdown.")
	}

}
