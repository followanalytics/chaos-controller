package grpc

import (
	"context"
	"log"
	"time"

	pb "github.com/DataDog/chaos-controller/grpc/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ExecuteSendDisruption(client pb.DisruptionListenerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	endpointAlteration := pb.EndpointAlteration{TargetEndpoint: "/chaos_dogfood.ChaosDogfood/order", ErrorToReturn: "InvalidRequest"}
	endpointAlterationList := []*pb.EndpointAlteration{&endpointAlteration}

	_, err := client.SendDisruption(ctx, &pb.DisruptionSpec{EndpointAlterations: endpointAlterationList})

	if err != nil {
		log.Printf("Received an error: %v", err)
	}
}

/*
func ExecuteDisruptionStatus(client pb.DisruptionListenerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.DisruptionStatus(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("Received an error: %v", err)
	}

	for _, alt := range res.GetEndpointAlterations() {
		log.Printf("\nalteration: for endpoint %s", alt.TargetEndpoint)
		if alt.TargetEndpoint != "" {
			log.Printf("error is %s", alt.ErrorToReturn)
		}
		if alt.OverrideToReturn != "" {
			log.Printf("override is %s", alt.OverrideToReturn)
		}
	}
}
*/
func ExecuteCleanDisruption(client pb.DisruptionListenerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := client.CleanDisruption(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("Received an error: %v", err)
	}
}

/*
func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewDisruptionListenerClient(conn)

	sendDisruption(client)
	time.Sleep(3 * time.Second)
	disruptionStatus(client)
	time.Sleep(3 * time.Second)
	cleanDisruption(client)
}
*/
