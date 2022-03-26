package auth

import (
	"echo-starter/internal/wellknown"

	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	claimsprincipalServices "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/claimsprincipal"
	services_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/claimsprincipal"
)

// BuildGrpcEntrypointPermissionsClaimsMap ...
func BuildGrpcEntrypointPermissionsClaimsMap() map[string]*middleware_oidc.EntryPointConfig {
	entryPointClaimsBuilder := claimsprincipalServices.NewEntryPointClaimsBuilder()
	// HEALTH SERVICE START
	//---------------------------------------------------------------------------------------------------
	// health check is open to anyone
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.HealthPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.ReadinessPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.LoginPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.LogoutPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.OIDCCallbackPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.HomePath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.AboutPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.ErrorPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen("/css*")
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen("/assets*")
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen("/js*")

	entryPointClaimsBuilder.GetClaimsConfig(wellknown.DeepPath).
		WithGrpcEntrypointPermissionsClaimFactsMapOR(
			services_claimsprincipal.NewClaimFactType(wellknown.ClaimTypeAuthenticated),
		).GetChild().
		WithGrpcEntrypointPermissionsClaimFactsMapOR(
			services_claimsprincipal.NewClaimFactTypeAndValue(wellknown.ClaimTypeDeep, wellknown.ClaimValueRead),
			services_claimsprincipal.NewClaimFactTypeAndValue(wellknown.ClaimTypeDeep, wellknown.ClaimValueReadWrite),
			services_claimsprincipal.NewClaimFactTypeAndValue(wellknown.ClaimTypeDeep, wellknown.ClaimValueReadWriteAll),
		)
	entryPointClaimsBuilder.AddMetaData(wellknown.DeepPath, map[string]interface{}{
		"onUnauthenticated": "login",
		"onUnauthorized":    "error",
	})
	cMap := entryPointClaimsBuilder.GrpcEntrypointClaimsMap
	return cMap
}