package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"golang.org/x/net/context"

	pb "github.com/Sh4d1/wat-movie-service/proto/movie"
	userService "github.com/Sh4d1/wat-user-service/proto/user"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	k8s "github.com/micro/kubernetes/go/micro"
)

var (
	srv micro.Service
)

func main() {

	if os.Getenv("DEV") == "true" {
		srv = micro.NewService(
			micro.Name("wat.movie"),
			micro.WrapHandler(AuthWrapper),
		)
	} else {
		srv = k8s.NewService(
			micro.Name("wat.movie"),
			micro.WrapHandler(AuthWrapper),
		)
	}
	srv.Init()

	pb.RegisterMovieServiceHandler(srv.Server(), &service{})

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.Forbidden("wat.movie", "No headers found")
		}
		authHeader := meta["Authorization"]
		if authHeader == "" {
			authHeader = meta["authorization"]
			if authHeader == "" {
				return errors.Forbidden("wat.movie", "Authorization header required")
			}
		}
		log.Println(authHeader)
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return errors.Forbidden("wat.movie", "Authorization requires Bearer auth")
		}

		token := authHeader[len("Bearer "):]
		//log.Println("Authenticating with token: ", token)
		authClient := userService.NewUserServiceClient("wat.user", srv.Client())
		_, err := authClient.ValidateToken(ctx, &userService.Token{
			Token: token,
		})
		if err != nil {
			var errorReturned errors.Error

			if err := json.Unmarshal([]byte(err.Error()), &errorReturned); err != nil {
				return errors.Forbidden("wat.movie", err.Error())
			}
			return errors.Forbidden("wat.movie", errorReturned.Detail)
		}
		err = fn(ctx, req, resp)
		return err
	}
}
