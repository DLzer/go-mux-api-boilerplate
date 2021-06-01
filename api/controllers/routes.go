package controllers

import middlewares "github.com/DLzer/go-product-api.git/api/middleware"

// initializeRoutes declares all routes used in the application
func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	s.Router.HandleFunc("/products", middlewares.SetMiddlewareJSON(s.getProducts)).Methods("GET")
	s.Router.HandleFunc("/product", middlewares.SetMiddlewareJSON(s.createProduct)).Methods("POST")
	s.Router.HandleFunc("/product/{id:[0-9]+}", middlewares.SetMiddlewareJSON(s.getProduct)).Methods("GET")
	s.Router.HandleFunc("/product/{id:[0-9]+}", middlewares.SetMiddlewareJSON(s.updateProduct)).Methods("PUT")
	s.Router.HandleFunc("/product/{id:[0-9]+}", middlewares.SetMiddlewareJSON(s.deleteProduct)).Methods("DELETE")
}
