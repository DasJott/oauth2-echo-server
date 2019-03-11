package echoserver

import (
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/server"
)

// SetTokenType token type
func SetTokenType(tokenType string) {
	eServer.Config.TokenType = tokenType
}

// SetAllowGetAccessRequest to allow GET requests for the token
func SetAllowGetAccessRequest(allow bool) {
	eServer.Config.AllowGetAccessRequest = allow
}

// SetAllowedResponseType allow the authorization types
func SetAllowedResponseType(types ...oauth2.ResponseType) {
	eServer.Config.AllowedResponseTypes = types
}

// SetAllowedGrantType allow the grant types
func SetAllowedGrantType(types ...oauth2.GrantType) {
	eServer.Config.AllowedGrantTypes = types
}

// SetClientInfoHandler get client info from request
func SetClientInfoHandler(handler server.ClientInfoHandler) {
	eServer.ClientInfoHandler = handler
}

// SetClientAuthorizedHandler check the client allows to use this authorization grant type
func SetClientAuthorizedHandler(handler server.ClientAuthorizedHandler) {
	eServer.ClientAuthorizedHandler = handler
}

// SetClientScopeHandler check the client allows to use scope
func SetClientScopeHandler(handler server.ClientScopeHandler) {
	eServer.ClientScopeHandler = handler
}

// SetUserAuthorizationHandler get user id from request authorization
func SetUserAuthorizationHandler(handler server.UserAuthorizationHandler) {
	eServer.UserAuthorizationHandler = handler
}

// SetPasswordAuthorizationHandler get user id from username and password
func SetPasswordAuthorizationHandler(handler server.PasswordAuthorizationHandler) {
	eServer.PasswordAuthorizationHandler = handler
}

// SetRefreshingScopeHandler check the scope of the refreshing token
func SetRefreshingScopeHandler(handler server.RefreshingScopeHandler) {
	eServer.RefreshingScopeHandler = handler
}

// SetResponseErrorHandler response error handling
func SetResponseErrorHandler(handler server.ResponseErrorHandler) {
	eServer.ResponseErrorHandler = handler
}

// SetInternalErrorHandler internal error handling
func SetInternalErrorHandler(handler server.InternalErrorHandler) {
	eServer.InternalErrorHandler = handler
}

// SetExtensionFieldsHandler in response to the access token with the extension of the field
func SetExtensionFieldsHandler(handler server.ExtensionFieldsHandler) {
	eServer.ExtensionFieldsHandler = handler
}

// SetAccessTokenExpHandler set expiration date for the access token
func SetAccessTokenExpHandler(handler server.AccessTokenExpHandler) {
	eServer.AccessTokenExpHandler = handler
}

// SetAuthorizeScopeHandler set scope for the access token
func SetAuthorizeScopeHandler(handler server.AuthorizeScopeHandler) {
	eServer.AuthorizeScopeHandler = handler
}
