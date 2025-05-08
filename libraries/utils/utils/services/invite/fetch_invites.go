package invite

import "net/http"

// TODO: You can't extend a struct can you?
type FetchInvitesRequest struct {
}

type FetchInvitesResponse struct {
}

func FetchInvites(request FetchInvitesRequest) (FetchInvitesResponse, error) {
	return FetchInvitesResponse{}, nil
}

func FetchInvitesParams(req *http.Request) (FetchInvitesRequest, error) {
	return FetchInvitesRequest{}, nil
}
