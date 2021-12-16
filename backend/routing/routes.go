package routing

import (
	"github.com/RodolfoMurguia/beat-invoice/controllers/invoice"
	"github.com/RodolfoMurguia/beat-invoice/controllers/ride"
	"github.com/RodolfoMurguia/beat-invoice/controllers/tax"
	"github.com/RodolfoMurguia/beat-invoice/controllers/user"
	fiber "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	//Notations Routing
	app.Get("/api/taxprofiles/", tax.GetTaxProfiles)
	app.Get("/api/taxprofiles/:name", tax.GetTaxProfilesByName)
	app.Post("/api/taxprofiles/", tax.AddTaxProfile)

	//Invoice Routing
	app.Get("/api/invoices/", invoice.GetInvoices)
	app.Get("/api/invoices/:id", invoice.GetInvoicesByUserId)
	app.Get("/api/invoices/pending/", invoice.GetPendingInvoices)
	app.Post("/api/invoices/", invoice.AddInvoice)
	app.Post("/api/invoices/export/", invoice.ExportInvoices)

	//User Routing
	app.Get("/api/users/", user.GetUsers)
	app.Get("/api/users/:id", user.GetUser)
	app.Post("/api/users/", user.AddUser)
	app.Post("/api/users/login/", user.LoginUser)
	//app.Post("/api/users/logout/", user.LogoutUser)
	app.Patch("/api/users/:id", user.UpdateUser)
	app.Delete("/api/users/:id", user.DeactivateUser)

	//Ride Routing
	app.Get("/api/rides/:id", ride.GetRidesbyUser)
	app.Get("/api/ride/:id", ride.GetRideById)
	app.Post("/api/rides/", ride.AddRide)
	//app.Post("/api/rides/toinvoice/", ride.GenerateInvoiceByRide)

}
