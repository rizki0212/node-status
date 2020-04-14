package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"exec/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {

	out, err := exec.Command("cockroach", "node", "status", "--insecure", "--format", "records").Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	status := string(out)
	s := strings.Split(status, "\n")

	responses.JSON(w, http.StatusOK, s)
}
