# echo starter  

[auth0-golang-web-app](https://github.com/auth0-samples/auth0-golang-web-app/)  
[demo-echo-app](https://github.com/gtongy/demo-echo-app)  
[cookie auth](https://www.sohamkamani.com/golang/session-cookie-authentication/)

# TLDR
You can run the app live via docker-compose.  I haven't tested it with Auth0 where the login only sends an access_token as a reference (not jwt) and no refresh_token;
So make sure your Auth0 setup delivers a JWT access_token with a refresh_token.

## Docker-Compose

### Secrets
place the following Auth0 env variables in your OS envs.  
```env
AUTH0_CLIENT_ID=M8x**REDACTED**4Zwk
AUTH0_CLIENT_SECRET=mbTS7_63xBUx**REDACTED**lYgmRwXPbMy8ai9Pd
AUTH0_DOMAIN={{yourdomain}}.auth0.com
```
example with none working secrets
```.env
AUTH0_CLIENT_ID=M8xSKujdhflsjdfPd3yfkBTjnajz4Zwk
AUTH0_CLIENT_SECRET=mbTS7_63xBUxkjlhdsjkdfhksdjdfhnvyoWlYgmRwXPbMy8ai9Pd
AUTH0_DOMAIN=foo.auth0.com
```

```bash
docker-compose pull
docker-compose up
```
Docker-Compose using [Traefik](https://traefik.io/) to do loadbalancing and gives us an url that doesn't have a port.   
naviage to the following [echo-starter](http://echostarter.docker.localhost/)  


# Motivation  

[ECHO](https://echo.labstack.com/) is a fantastic base framework to build upon.  This project adds a lot of design patterns found in [asp.net core](https://docs.microsoft.com/en-us/aspnet/core/introduction-to-aspnet-core).  

1. Introduce [depedency injection](https://github.com/fluffy-bunny/sarulabsdi) with SINGLETON, SCOPED and TRANSIENT features  
As with asp.net core, when a request comes in, we have a context.  A scoped container is created and the handler of that request is a registered as SCOPED.  

The [home handler](internal/services/handlers/home/home.go) is an example.  

Notice the following injected SCOPED ClaimsPrincipal object;  
```go
	service struct {
		ClaimsPrincipal contracts_core_claimsprincipal.IClaimsPrincipal `inject:"claimsPrincipal"`
	}
```
Just like asp.net core, the claims principal is created on each request.  This is usually populated from an auth cookie or a jwt token.  

2. Introduce asp.net style authentication/authorization using middleware.  
Here we have a 2 phase pipeline.  First the claims principal is created but no gating action is done.  A downstream middleware only works on claims principals and gates access to paths.  This allows us to introduce any middleware that can create a claims principal from whatever auth scheme.  Cookie auth and JWT auth are 2 well known schemes that can produce a claims principal.  

3. Templates are ECHO recommendations
4. Sessions are ECHO recommendations
5. Cookies are ECHO recommendations

6. Bring in other nice asp.net standard injectable objects and funcs;  
IDistributedCache  
IMemoryCache  
func ContainerAccessor vs aps.net's IServiceProvider  
etc.   



