package baseapp

import (
	"fmt"
	"strings"

	gocontext "context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogogrpc "github.com/gogo/protobuf/grpc"
	abci "github.com/tendermint/tendermint/abci/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var protoCodec = encoding.GetCodec(proto.Name)

type QueryRouter struct {
	routes map[string]sdk.Querier
}

var _ sdk.QueryRouter = NewQueryRouter()

// NewQueryRouter returns a reference to a new QueryRouter.
func NewQueryRouter() *QueryRouter {
	return &QueryRouter{
		routes: map[string]sdk.Querier{},
	}
}

// AddRoute adds a query path to the router with a given Querier. It will panic
// if a duplicate route is given. The route must be alphanumeric.
func (qrt *QueryRouter) AddRoute(path string, q sdk.Querier) sdk.QueryRouter {
	if !isAlphaNumeric(path) {
		panic("route expressions can only contain alphanumeric characters")
	}
	if qrt.routes[path] != nil {
		panic(fmt.Sprintf("route %s has already been initialized", path))
	}

	qrt.routes[path] = q
	return qrt
}

// Route returns the Querier for a given query route path.
func (qrt *QueryRouter) Route(path string) sdk.Querier {
	return qrt.routes[path]
}

// RegisterService implements the grpc Server.RegisterService method
func (qrt *QueryRouter) RegisterService(sd *grpc.ServiceDesc, handler interface{}) {
	// adds a top-level querier based on the GRPC service name
	qrt.routes[sd.ServiceName] =
		func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
			path0 := path[0]
			for _, md := range sd.Methods {
				// checks each GRPC service method to see if it matches the path
				if md.MethodName != path0 {
					continue
				}
				res, err := md.Handler(handler, sdk.WrapSDKContext(ctx), func(i interface{}) error {
					return protoCodec.Unmarshal(req.Data, i)
				}, nil)
				if err != nil {
					return nil, err
				}
				return protoCodec.Marshal(res)
			}
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
}

// QueryServiceTestHelper provides a helper for making grpc query service
// rpc calls in unit tests. It implements both the grpc Server and ClientConn
// interfaces needed to register a query service server and create a query
// service client.
type QueryServiceTestHelper struct {
	*QueryRouter
	ctx sdk.Context
}

// NewQueryServerTestHelper creates a new QueryServiceTestHelper that wraps
// the provided sdk.Context
func NewQueryServerTestHelper(ctx sdk.Context) *QueryServiceTestHelper {
	return &QueryServiceTestHelper{QueryRouter: NewQueryRouter(), ctx: ctx}
}

// Invoke implements the grpc ClientConn.Invoke method
func (q *QueryServiceTestHelper) Invoke(_ gocontext.Context, method string, args, reply interface{}, _ ...grpc.CallOption) error {
	path := strings.Split(method, "/")
	if len(path) != 3 {
		return fmt.Errorf("unexpected method name %s", method)
	}
	querier := q.Route(path[1])
	if querier == nil {
		return fmt.Errorf("handler not found for %s", path[1])
	}
	reqBz, err := protoCodec.Marshal(args)
	if err != nil {
		return err
	}
	resBz, err := querier(q.ctx, path[2:], abci.RequestQuery{Data: reqBz})
	if err != nil {
		return err
	}
	return protoCodec.Unmarshal(resBz, reply)
}

// NewStream implements the grpc ClientConn.NewStream method
func (q *QueryServiceTestHelper) NewStream(gocontext.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("not supported")
}

var _ gogogrpc.Server = &QueryServiceTestHelper{}
var _ gogogrpc.ClientConn = &QueryServiceTestHelper{}
