package client

import (
	"github.com/infobloxopen/atlas-app-toolkit/query"
	"google.golang.org/grpc"
	context "golang.org/x/net/context"
	"github.com/seizadi/cmdb/pkg/pb"
	"google.golang.org/grpc/metadata"
	atlas_rpc "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"

)

type CmdbClient struct {
	// The CMDB Server
	Host string
	// The CMDB Server Connection
	Conn *grpc.ClientConn
	// The CMDB Server API Key
	ApiKey string
}

func NewCmdbClient(host string, apiKey string) (*CmdbClient, error) {
	c := CmdbClient{}
	
	err := c.GetConn(host)
	if err != nil {
		return nil, err
	}
	
	c.ApiKey = apiKey
	return &c,nil
}

func (m *CmdbClient) GetConn(host string) error{
	// First we create the connection:
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return err
	}
	m.Conn = conn
	return nil
}

func (m *CmdbClient) GetCloudProviders() (*pb.ListCloudProvidersResponse, error) {
	
	// We can now create stubs that wrap conn:
	stub := pb.NewCloudProvidersClient(m.Conn)
	
	// Now we can use the stub to make RPCs
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer " + m.ApiKey))
	reqProviders := &pb.ListCloudProviderRequest{}
	respProviders, err := stub.List(ctx, reqProviders)
	if err != nil {
		return nil, err
	}
	
	return respProviders, nil
}

func (m *CmdbClient) GetRegions() (*pb.ListRegionsResponse, error) {
	
	// We can now create stubs that wrap conn:
	stub := pb.NewRegionsClient(m.Conn)
	
	// Now we can use the stub to make RPCs
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer " + m.ApiKey))
	reqList := &pb.ListRegionRequest{}
	respList, err := stub.List(ctx, reqList)
	if err != nil {
		return nil, err
	}
	
	return respList, nil
}


func (m *CmdbClient) GetRegion(id string) (*pb.ReadRegionResponse, error) {
	
	stub := pb.NewRegionsClient(m.Conn)
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer " + m.ApiKey))
	
	resourceId := atlas_rpc.Identifier{ResourceId: id}
	fields := query.FieldSelection{}
	
	// We'll make another request and also print the response metadata
	req := &pb.ReadRegionRequest{Id: &resourceId, Fields: &fields}
	var respHdrs, respTrlrs metadata.MD
	resp, err := stub.Read(ctx, req,
		grpc.Header(&respHdrs), grpc.Trailer(&respTrlrs))
	if err != nil {
		return nil, err
	}
	
	return resp, nil
}


func (m *CmdbClient) CreateCloudProvider(req *pb.CreateCloudProviderRequest) (*pb.CreateCloudProviderResponse, error) {
	
	// We can now create stubs that wrap conn:
	stub := pb.NewCloudProvidersClient(m.Conn)
	
	// Now we can use the stub to make RPCs
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer " + m.ApiKey))
	res, err := stub.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	
	return res, nil
}


func (m *CmdbClient) CreateRegion(req *pb.CreateRegionRequest) (*pb.CreateRegionResponse, error) {
	
	// We can now create stubs that wrap conn:
	stub := pb.NewRegionsClient(m.Conn)
	
	// Now we can use the stub to make RPCs
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer " + m.ApiKey))
	res, err := stub.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	
	return res, nil
}

func (m *CmdbClient) CreateStage(req *pb.CreateStageRequest) (*pb.CreateStageResponse, error) {
	
	// We can now create stubs that wrap conn:
	stub := pb.NewStagesClient(m.Conn)
	
	// Now we can use the stub to make RPCs
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer " + m.ApiKey))
	res, err := stub.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	
	return res, nil
}

func (m *CmdbClient) CreateEnvironment(req *pb.CreateEnvironmentRequest) (*pb.CreateEnvironmentResponse, error) {
	
	// We can now create stubs that wrap conn:
	stub := pb.NewEnvironmentsClient(m.Conn)
	
	// Now we can use the stub to make RPCs
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer " + m.ApiKey))
	res, err := stub.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	
	return res, nil
}

func (m *CmdbClient) CreateApplication(req *pb.CreateApplicationRequest) (*pb.CreateApplicationResponse, error) {
	
	// We can now create stubs that wrap conn:
	stub := pb.NewApplicationsClient(m.Conn)
	
	// Now we can use the stub to make RPCs
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer " + m.ApiKey))
	res, err := stub.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	
	return res, nil
}

func (m *CmdbClient) GetApplications() (*pb.ListApplicationsResponse, error) {
	
	// We can now create stubs that wrap conn:
	stub := pb.NewApplicationsClient(m.Conn)
	
	// Now we can use the stub to make RPCs
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer " + m.ApiKey))
	reqList := &pb.ListApplicationRequest{}
	respList, err := stub.List(ctx, reqList)
	if err != nil {
		return nil, err
	}
	
	return respList, nil
}

func (m *CmdbClient) CreateApplicationInstance(req *pb.CreateApplicationInstanceRequest) (*pb.CreateApplicationInstanceResponse, error) {
	
	// We can now create stubs that wrap conn:
	stub := pb.NewApplicationInstancesClient(m.Conn)
	
	// Now we can use the stub to make RPCs
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer " + m.ApiKey))
	res, err := stub.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	
	return res, nil
}

func (m *CmdbClient) CreateChartVersion(req *pb.CreateChartVersionRequest) (*pb.CreateChartVersionResponse, error) {
	
	// We can now create stubs that wrap conn:
	stub := pb.NewChartVersionsClient(m.Conn)
	
	// Now we can use the stub to make RPCs
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer " + m.ApiKey))
	res, err := stub.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	
	return res, nil
}

func (m *CmdbClient) CreateValue(req *pb.CreateValueRequest) (*pb.CreateValueResponse, error) {
	
	// We can now create stubs that wrap conn:
	stub := pb.NewValuesClient(m.Conn)
	
	// Now we can use the stub to make RPCs
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer " + m.ApiKey))
	res, err := stub.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	
	return res, nil
}

