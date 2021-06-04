package controllers

import middlewares "github.com/DLzer/go-product-api/api/middleware"

// initializeRoutes declares all routes used in the application
func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	s.Router.HandleFunc("/products", middlewares.SetMiddlewareJSON(s.FindAllProducts)).Methods("GET")
	s.Router.HandleFunc("/product", middlewares.SetMiddlewareJSON(s.CreateProduct)).Methods("POST")
	s.Router.HandleFunc("/product/{id:[0-9]+}", middlewares.SetMiddlewareJSON(s.GetProduct)).Methods("GET")
	s.Router.HandleFunc("/product/{id:[0-9]+}", middlewares.SetMiddlewareJSON(s.UpdateProduct)).Methods("PUT")
	s.Router.HandleFunc("/product/{id:[0-9]+}", middlewares.SetMiddlewareJSON(s.DeleteProduct)).Methods("DELETE")
}
