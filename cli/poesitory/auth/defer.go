package auth

func DeferAuthentication(webAuth bool, userToken, uploadToken string) string {
	if webAuth {
		return WebAuthentication()
	}
	if len(userToken) > 0 {
		return UserAuthentication(userToken)
	}
	if len(uploadToken) > 0 {
		return UserAuthentication(uploadToken)
	}
	return WebAuthentication()
}
