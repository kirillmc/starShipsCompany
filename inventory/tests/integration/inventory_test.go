///go:build integration

package integration

import (
	"context"

	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ = Describe("InventoryService", func() {
	var (
		ctx             context.Context
		cancel          context.CancelFunc
		inventoryClient inventoryV1.InventoryServiceClient
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(suiteCtx)

		conn, err := grpc.NewClient(
			env.App.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		Expect(err).ToNot(HaveOccurred(), "ожидали успешное подключение к gRPC приложению")

		inventoryClient = inventoryV1.NewInventoryServiceClient(conn)
	})

	AfterEach(func() {
		err := env.ClearPartsCollection(ctx)
		Expect(err).ToNot(HaveOccurred(), "ожидали успешную очистку коллекции parts")

		cancel()
	})

	Describe("ListParts", func() {
		It("должен найти созданных по умолчанию деталей", func() {
			resp, err := inventoryClient.ListParts(ctx, &inventoryV1.ListPartsRequest{Filter: &inventoryV1.PartsFilter{}})
			Expect(err).ToNot(HaveOccurred())
			Expect(resp).ToNot(BeNil())
			Expect(len(resp.Parts)).To(BeNumerically(">", 0))
		})

		It("должен  вернуть пустой набор деталей", func() {
			resp, err := inventoryClient.ListParts(ctx, env.GetWrongListPartsParams())
			Expect(err).ToNot(HaveOccurred())
			Expect(resp).ToNot(BeNil())
			Expect(resp.Parts).To(BeNil())
		})
	})

	Describe("Get", func() {
		var partUUID string

		BeforeEach(func() {
			var err error
			partUUID, err = env.InsertToPartsCollection(ctx)
			Expect(err).ToNot(HaveOccurred(),
				"ожидали успешную вставку тестовой детали в MongoDB")
		})

		It("должен успешно возвращать деталь по UUID", func() {
			resp, err := inventoryClient.GetPart(ctx, &inventoryV1.GetPartRequest{
				Uuid: partUUID,
			})

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.GetPart()).ToNot(BeNil())
			Expect(resp.GetPart().Uuid).To(Equal(partUUID))
			Expect(resp.GetPart().GetCreatedAt()).ToNot(BeNil())
			Expect(resp.GetPart().GetDimensions()).ToNot(BeNil())
		})
	})
})
