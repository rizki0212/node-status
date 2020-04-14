package controllers

import "exec/api/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/status", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

}
