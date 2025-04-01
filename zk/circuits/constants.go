package circuits

const (
	MaxBase64JwtHeaderLen  = 256
	MaxBase64JwtPayloadLen = 256
	MaxBase64JwtLen        = MaxBase64JwtHeaderLen + MaxBase64JwtPayloadLen + 1

	MaxJwtHeaderKidValueLen = 64 // Google uses thumbprint which is ~40 chars

	MaxJwtPayloadIssLen = 64 // Likely too small for Microsoft.
	MaxJwtPayloadAudLen = 64
	MaxJwtPayloadSubLen = 64
)

const (
	ExpectedTypJson = `"typ":"JWT"`
	ExpectedAlgJson = `"alg":"RSA"`

	ExpectedKidPrefixJson = `"kid":`

	ExpectedIssPrefixJson = `"iss":`
	ExpectedAudPrefixJson = `"aud":`
	ExpectedSubPrefixJson = `"sub":`
)
