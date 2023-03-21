// Code generated by protoc-gen-go-flipt-sdk. DO NOT EDIT.

package http

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	auth "go.flipt.io/flipt/rpc/flipt/auth"
	sdk "go.flipt.io/flipt/sdk"
	grpc "google.golang.org/grpc"
	protojson "google.golang.org/protobuf/encoding/protojson"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http "net/http"
	url "net/url"
	strconv "strconv"
)

type authClient struct {
	client *http.Client
	addr   string
}

func (t authClient) PublicAuthenticationServiceClient() auth.PublicAuthenticationServiceClient {
	return &publicAuthenticationServiceClient{client: t.client, addr: t.addr}
}

type publicAuthenticationServiceClient struct {
	client *http.Client
	addr   string
}

func (x *publicAuthenticationServiceClient) ListAuthenticationMethods(ctx context.Context, v *emptypb.Empty, _ ...grpc.CallOption) (*auth.ListAuthenticationMethodsResponse, error) {
	var body io.Reader
	var values url.Values
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, x.addr+"/auth/v1/method", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output auth.ListAuthenticationMethodsResponse
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (t authClient) AuthenticationServiceClient() auth.AuthenticationServiceClient {
	return &authenticationServiceClient{client: t.client, addr: t.addr}
}

type authenticationServiceClient struct {
	client *http.Client
	addr   string
}

func (x *authenticationServiceClient) GetAuthenticationSelf(ctx context.Context, v *emptypb.Empty, _ ...grpc.CallOption) (*auth.Authentication, error) {
	var body io.Reader
	var values url.Values
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, x.addr+"/auth/v1/self", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output auth.Authentication
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *authenticationServiceClient) GetAuthentication(ctx context.Context, v *auth.GetAuthenticationRequest, _ ...grpc.CallOption) (*auth.Authentication, error) {
	var body io.Reader
	var values url.Values
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, x.addr+fmt.Sprintf("/auth/v1/tokens/%v", v.Id), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output auth.Authentication
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *authenticationServiceClient) ListAuthentications(ctx context.Context, v *auth.ListAuthenticationsRequest, _ ...grpc.CallOption) (*auth.ListAuthenticationsResponse, error) {
	var body io.Reader
	values := url.Values{}
	values.Set("method", fmt.Sprintf("%v", v.Method))
	values.Set("limit", fmt.Sprintf("%v", v.Limit))
	values.Set("pageToken", v.PageToken)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, x.addr+"/auth/v1/tokens", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output auth.ListAuthenticationsResponse
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *authenticationServiceClient) DeleteAuthentication(ctx context.Context, v *auth.DeleteAuthenticationRequest, _ ...grpc.CallOption) (*emptypb.Empty, error) {
	var body io.Reader
	var values url.Values
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, x.addr+fmt.Sprintf("/auth/v1/tokens/%v", v.Id), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output emptypb.Empty
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *authenticationServiceClient) ExpireAuthenticationSelf(ctx context.Context, v *auth.ExpireAuthenticationSelfRequest, _ ...grpc.CallOption) (*emptypb.Empty, error) {
	var body io.Reader
	var field []byte
	var err error
	var unquote = func(v []byte) string {
		s, err := strconv.Unquote(string(v))
		if err == nil {
			return s
		}
		return string(v)
	}
	values := url.Values{}
	field, err = protojson.Marshal(v.ExpiresAt)
	if err != nil {
		return nil, err
	}

	values.Set("expiresAt", unquote(field))
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, x.addr+"/auth/v1/self/expire", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output emptypb.Empty
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (t authClient) AuthenticationMethodTokenServiceClient() auth.AuthenticationMethodTokenServiceClient {
	return &authenticationMethodTokenServiceClient{client: t.client, addr: t.addr}
}

type authenticationMethodTokenServiceClient struct {
	client *http.Client
	addr   string
}

func (x *authenticationMethodTokenServiceClient) CreateToken(ctx context.Context, v *auth.CreateTokenRequest, _ ...grpc.CallOption) (*auth.CreateTokenResponse, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, x.addr+"/auth/v1/method/token", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output auth.CreateTokenResponse
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (t authClient) AuthenticationMethodOIDCServiceClient() auth.AuthenticationMethodOIDCServiceClient {
	return &authenticationMethodOIDCServiceClient{client: t.client, addr: t.addr}
}

type authenticationMethodOIDCServiceClient struct {
	client *http.Client
	addr   string
}

func (x *authenticationMethodOIDCServiceClient) AuthorizeURL(ctx context.Context, v *auth.AuthorizeURLRequest, _ ...grpc.CallOption) (*auth.AuthorizeURLResponse, error) {
	var body io.Reader
	values := url.Values{}
	values.Set("state", v.State)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, x.addr+fmt.Sprintf("/auth/v1/method/oidc/%v/authorize", v.Provider), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output auth.AuthorizeURLResponse
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *authenticationMethodOIDCServiceClient) Callback(ctx context.Context, v *auth.CallbackRequest, _ ...grpc.CallOption) (*auth.CallbackResponse, error) {
	var body io.Reader
	values := url.Values{}
	values.Set("code", v.Code)
	values.Set("state", v.State)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, x.addr+fmt.Sprintf("/auth/v1/method/oidc/%v/callback", v.Provider), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output auth.CallbackResponse
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (t authClient) AuthenticationMethodKubernetesServiceClient() auth.AuthenticationMethodKubernetesServiceClient {
	return &authenticationMethodKubernetesServiceClient{client: t.client, addr: t.addr}
}

type authenticationMethodKubernetesServiceClient struct {
	client *http.Client
	addr   string
}

func (x *authenticationMethodKubernetesServiceClient) VerifyServiceAccount(ctx context.Context, v *auth.VerifyServiceAccountRequest, _ ...grpc.CallOption) (*auth.VerifyServiceAccountResponse, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, x.addr+"/auth/v1/method/kubernetes/serviceaccount", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output auth.VerifyServiceAccountResponse
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (t Transport) AuthClient() sdk.AuthClient {
	return authClient{client: t.client, addr: t.addr}
}
