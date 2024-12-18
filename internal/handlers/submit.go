package handlers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/garrettladley/prods/internal/algo"
	"github.com/garrettladley/prods/internal/model/product"
	"github.com/garrettladley/prods/internal/xerr"
	"github.com/garrettladley/prods/internal/xhttp"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"

	go_json "github.com/goccy/go-json"
)

type submitRequestBody struct {
	URL string `json:"url"`
}

// Submit godoc
//
//	@Summary		Submit and score a solution
//	@Description	Scores a solution by hitting the provided endpoint and checking against the expected solution.
//	@Description	Before checking your solution, we will make a GET request to the provided URL's `/health` endpoint.
//	@Description	If this request does not return a 200 status code, the scoring process will exit, and you will receive a score of -1.
//	@Description	Your endpoint should be located at URL+"/api/v1/products".
//	@Tags			solutions
//	@Accept			json
//	@Produce		json
//	@Param			token	path		string				true	"Submission token"	format(uuid)
//	@Param			body	body		submitRequestBody	true	"Submission data containing the solution URL"
//	@Success		201		{integer}	integer				"Solution successfully scored"
//	@Failure		400		{object}	xerr.APIError		"Invalid token or malformed request body"
//	@Failure		408		{object}	xerr.APIError		"Request timeout exceeded"
//	@Failure		429		{object}	xerr.APIError		"Too many requests"
//	@Failure		500		{object}	xerr.APIError		"Internal server error"
//	@Router			/api/v1/{token}/submit [post]
func (s *Service) Submit(c *fiber.Ctx) error {
	rawToken := c.Params("token")
	token, err := uuid.Parse(rawToken)
	if err != nil {
		return xerr.BadRequest(fmt.Errorf("failed to parse token. got: %s", rawToken))
	}

	var r submitRequestBody
	if err := c.BodyParser(&r); err != nil {
		slog.Error("invalid JSON request data", "error", err)
		return xerr.InvalidJSON()
	}

	baseCtx, cancel := context.WithTimeout(c.Context(), 30*time.Second)
	defer cancel()

	ok, err := health(baseCtx, r.URL)
	if err != nil || !ok {
		score := -1
		if err := s.storage.Submit(baseCtx, token, score); err != nil {
			return err
		}

		return c.Status(http.StatusCreated).JSON(score)
	}

	eg, ctx := errgroup.WithContext(baseCtx)

	var expected algo.Solution
	eg.Go(func() error {
		solution, err := s.storage.Solution(ctx, token)
		if err != nil {
			return err
		}
		expected = solution
		return nil
	})

	var actual [][]product.Product
	eg.Go(func() error {
		solution, err := test(ctx, r.URL)
		if err != nil {
			return err
		}
		actual = solution
		return nil
	})

	if err := eg.Wait(); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return xerr.Timeout("took too long to test your solution")
		} else {
			return err
		}
	}

	score := s.algo.Score(baseCtx, expected, actual)
	if err := s.storage.Submit(baseCtx, token, score); err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(score)
}

func health(ctx context.Context, url string) (ok bool, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url+"/health", nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}

func test(ctx context.Context, url string) ([][]product.Product, error) {
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(8)

	solutions := make([][]product.Product, len(algo.EncodedFilters))

	for idx, filter := range algo.EncodedFilters {
		func(idx int, filter string) {
			eg.Go(func() error {
				req, err := http.NewRequestWithContext(ctx, http.MethodGet, url+"/api/v1/products?"+filter, nil)
				if err != nil {
					return fmt.Errorf("failed to create request: %w", err)
				}
				req.Header.Set(xhttp.HeaderContentType, xhttp.HeaderApplicationJSON)

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return fmt.Errorf("failed to make request: %w", err)
				}
				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
				}

				var products []product.Product
				if err := go_json.NewDecoder(resp.Body).Decode(&products); err != nil {
					return fmt.Errorf("failed to decode response: %w", err)
				}

				solutions[idx] = products

				return nil
			})
		}(idx, filter)
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return solutions, nil
}
