package main

import (
	"context"
	"fmt"
)

const (
	CtxKeyUserID    = "user_id"
	CtxKeyAuthToken = "auth_token"
)

func main() {
	ProcessRequest("shuryak", "secret")
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), CtxKeyUserID, userID)
	ctx = context.WithValue(ctx, CtxKeyAuthToken, authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"handling response from %s (auth: %s)\n",
		ctx.Value(CtxKeyUserID),
		ctx.Value(CtxKeyAuthToken),
	)
}
