package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
)

func TaskOwner() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value(controllers.UserKey).(domain.User)
			if !ok {
				err := errors.New("user not found in context")
				log.Print(err)
				controllers.Unauthorized(w, err)
				return
			}

			task, ok := r.Context().Value(controllers.TaskKey).(domain.Task)
			if !ok {
				err := errors.New("task not found in context")
				log.Print(err)
				controllers.BadRequest(w, err)
				return
			}

			if task.UserId != user.Id {
				err := errors.New("access denied")
				log.Print(err)
				controllers.Forbidden(w, err)
				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(hfn)
	}
}
