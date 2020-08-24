package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/orlmonteverde/go-postgres-microblog/internal/middleware"
	"github.com/orlmonteverde/go-postgres-microblog/pkg/post"
	"github.com/orlmonteverde/go-postgres-microblog/pkg/response"
)

// PostRouter is the router of the posts.
type PostRouter struct {
	Repository post.Repository
}

// CreateHandler Create a new post.
func (pr *PostRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var p post.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = pr.Repository.Create(ctx, &p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), p.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"post": p})
}

// GetAllHandler response all the posts.
func (pr *PostRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	posts, err := pr.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"posts": posts})
}

// GetOneHandler response one post by id.
func (pr *PostRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	p, err := pr.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"post": p})
}

// UpdateHandler update a stored post by id.
func (pr *PostRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var p post.Post
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = pr.Repository.Update(ctx, uint(id), p)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

// DeleteHandler Remove a post by ID.
func (pr *PostRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = pr.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{})
}

// GetByUserHandler response posts by user id.
func (pr *PostRouter) GetByUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userId")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	posts, err := pr.Repository.GetByUser(ctx, uint(userID))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"posts": posts})
}

// Routes returns post router with each endpoint.
func (pr *PostRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Authorizator)

	r.Get("/user/{userId}", pr.GetByUserHandler)

	r.Get("/", pr.GetAllHandler)

	r.Post("/", pr.CreateHandler)

	r.Get("/{id}", pr.GetOneHandler)

	r.Put("/{id}", pr.UpdateHandler)

	r.Delete("/{id}", pr.DeleteHandler)

	return r
}
