package service

import (
	"net/http"

	"github.com/czsilence/EvernoteTransfer/erro"
	"github.com/dghubble/oauth1"
)

// oauth for evernote
// ref: https://dev.evernote.com/doc/articles/authentication.php

var oauth_config *oauth1.Config

func oauth_init() {
	oauth_config = &oauth1.Config{
		ConsumerKey:    opt.Evernote.Key,
		ConsumerSecret: opt.Evernote.Secret,
		CallbackURL:    "http://127.0.0.1:8001/api/en/oauth/callback",
		Endpoint: oauth1.Endpoint{
			RequestTokenURL: "https://sandbox.evernote.com/oauth",
			AuthorizeURL:    "https://sandbox.evernote.com/OAuth.action",
			AccessTokenURL:  "https://sandbox.evernote.com/oauth",
		},
	}
}

func OAuth_Auth() (auth_url string, err erro.Error) {
	if tok, _, re := oauth_config.RequestToken(); re != nil {
		err = erro.E_OAUTH_FAILED.F("err: %v", re)
	} else if authorizationURL, ae := oauth_config.AuthorizationURL(tok); ae != nil {
		err = erro.E_OAUTH_FAILED.F("err: %v", ae)
	} else {
		auth_url = authorizationURL.String()
	}
	return
}

func OAuth_ParseAccessToken(req *http.Request) (tok, verifier string, err erro.Error) {
	if _tok, _verifier, pe := oauth1.ParseAuthorizationCallback(req); pe != nil {
		err = erro.E_OAUTH_FAILED.F("err: %v", pe)
	} else {
		tok, verifier = _tok, _verifier
	}
	return
}