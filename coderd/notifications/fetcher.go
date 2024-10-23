package notifications

import (
	"context"
	"database/sql"
	"errors"
	"text/template"

	"golang.org/x/xerrors"

	"cdr.dev/slog"
)

func (n *notifier) fetchHelpers(ctx context.Context) (map[string]any, error) {
	appName, err := n.fetchAppName(ctx)
	if err != nil {
		n.log.Error(ctx, "failed to fetch app name", slog.Error(err))
		return nil, xerrors.Errorf("fetch app name: %w", err)
	}
	logoURL, err := n.fetchLogoURL(ctx)
	if err != nil {
		n.log.Error(ctx, "failed to fetch logo URL", slog.Error(err))
		return nil, xerrors.Errorf("fetch logo URL: %w", err)
	}

	helpers := make(template.FuncMap)
	for k, v := range n.helpers {
		helpers[k] = v
	}

	helpers["app_name"] = func() string { return appName }
	helpers["logo_url"] = func() string { return logoURL }

	return helpers, nil
}

func (n *notifier) fetchAppName(ctx context.Context) (string, error) {
	appName, err := n.store.GetApplicationName(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return notificationsDefaultAppName, nil
		}
		return "", xerrors.Errorf("get application name: %w", err)
	}
	return appName, nil
}

func (n *notifier) fetchLogoURL(ctx context.Context) (string, error) {
	logoURL, err := n.store.GetLogoURL(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return notificationsDefaultLogoURL, nil
		}
		return "", xerrors.Errorf("get logo URL: %w", err)
	}
	return logoURL, nil
}