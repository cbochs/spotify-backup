package backup

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

func tokenSource(ctx context.Context, id primitive.ObjectID, s *Service,
	ts oauth2.TokenSource) *dbTokenSource {

	return &dbTokenSource{
		ctx:     ctx,
		id:      id,
		src:     ts,
		service: s,
	}
}

type dbTokenSource struct {
	ctx     context.Context
	id      primitive.ObjectID
	src     oauth2.TokenSource
	service *Service
}

func (ts *dbTokenSource) Token() (*oauth2.Token, error) {
	tok, err := ts.src.Token()
	if err != nil {
		return nil, err
	}

	user, err := ts.service.db.Users().FindByID(ts.ctx, ts.id)
	if err != nil {
		return nil, err
	}

	if tokenNeedsUpdate(user.Token, tok) {
		if err := ts.service.db.Users().SaveToken(ts.ctx, user.ID, tok); err != nil {
			return nil, err
		}
	}

	return tok, nil
}

func tokenNeedsUpdate(old *oauth2.Token, new *oauth2.Token) bool {
	return old == nil || new == nil ||
		old.AccessToken != new.AccessToken ||
		old.RefreshToken != new.RefreshToken
}
