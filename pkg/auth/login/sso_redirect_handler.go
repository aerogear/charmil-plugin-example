package login

import (
	"context"
	// embed static HTML file
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/aerogear/charmil-plugin-example/pkg/auth/token"
	"github.com/aerogear/charmil-plugin-example/pkg/config"
	"github.com/aerogear/charmil-plugin-example/pkg/connection"
	"github.com/aerogear/charmil/core/utils/iostreams"
	"github.com/aerogear/charmil/core/utils/localize"
	"github.com/aerogear/charmil/core/utils/logging"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

//go:embed static/sso-redirect-page.html
var ssoRedirectHTMLPage string

// handler for the SSO redirect page
type redirectPageHandler struct {
	IO            *iostreams.IOStreams
	CfgHandler    *config.CfgHandler
	Logger        logging.Logger
	ServerAddr    string
	Port          int
	AuthOptions   []oauth2.AuthCodeOption
	State         string
	Oauth2Config  *oauth2.Config
	Ctx           context.Context
	TokenVerifier *oidc.IDTokenVerifier
	AuthURL       *url.URL
	ClientID      string
	Localizer     localize.Localizer
	CancelContext context.CancelFunc
}

// nolint:funlen
func (h *redirectPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	callbackURL := fmt.Sprintf("%v%v", h.ServerAddr, r.URL.String())
	h.Logger.Infoln("Redirected to callback URL:", callbackURL)
	h.Logger.Infoln()

	if r.URL.Query().Get("state") != h.State {
		http.Error(w, "state did not match", http.StatusBadRequest)
		return
	}

	// nolint:govet
	oauth2Token, err := h.Oauth2Config.Exchange(h.Ctx, r.URL.Query().Get("code"), h.AuthOptions...)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}
	idToken, err := h.TokenVerifier.Verify(h.Ctx, rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{oauth2Token, new(json.RawMessage)}

	if err = idToken.Claims(&resp.IDTokenClaims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username, ok := token.GetUsername(oauth2Token.AccessToken)
	if !ok {
		username = "unknown"
	}

	pageTitle := h.Localizer.LocalizeByID("login.redirectPage.title")
	pageBody := h.Localizer.LocalizeByID("login.redirectPage.body", localize.NewEntry("Username", username))

	issuerURL, realm, ok := connection.SplitKeycloakRealmURL(h.AuthURL)
	if !ok {
		h.Logger.Error(h.Localizer.LocalizeByID("login.error.noRealmInURL"))
		os.Exit(1)
	}
	redirectPage := fmt.Sprintf(ssoRedirectHTMLPage, pageTitle, pageTitle, pageBody, issuerURL, realm, h.ClientID)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, redirectPage)

	// save the received tokens to the user's config
	h.CfgHandler.Cfg.AccessToken = oauth2Token.AccessToken
	h.CfgHandler.Cfg.RefreshToken = oauth2Token.RefreshToken

	h.CancelContext()
}
