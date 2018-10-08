package api

//
type userHandler struct {
	s users.UserService
}

func (h *userHandler) (w http.ResponseWriter, r *http.Request) {
	//TODO will demonstrate testing handler later...
}
