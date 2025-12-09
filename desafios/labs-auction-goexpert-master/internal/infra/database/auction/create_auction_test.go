package auction_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/infra/database/auction"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Use MONGO_URI se definido, caso contrário assume mongodb://localhost:27017
func mongoURI() string {
	if uri := os.Getenv("MONGO_URI"); uri != "" {
		return uri
	}
	return "mongodb://admin:admin@localhost:27017/auctions?authSource=admin"
}

func TestCreateAuction_ClosesAutomatically(t *testing.T) {
	// contexto com timeout para evitar bloqueio indefinido
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// conecta no mongo
	clientOpts := options.Client().ApplyURI(mongoURI())
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		t.Fatalf("erro conectando ao mongo: %v\n(Verifique se o Mongo está rodando - ex: docker-compose -f docker-compose.test.yml up -d)", err)
	}
	defer func() {
		_ = client.Disconnect(context.Background())
	}()

	// Verifica conectividade com Ping para dar erro claro caso Mongo não esteja pronto
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()
	if err := client.Ping(pingCtx, nil); err != nil {
		t.Fatalf("impossível pingar o mongo em %s: %v\n(Se estiver usando docker-compose, rode: docker-compose -f docker-compose.test.yml up -d)", mongoURI(), err)
	}

	// usa um database isolado por teste
	dbName := "test_fullcycle_auctions_" + time.Now().Format("20060102150405")
	db := client.Database(dbName)
	coll := db.Collection("auctions")

	// garante limpeza ao final
	defer func() {
		_ = db.Drop(context.Background())
	}()

	// seta intervalo curto para o teste (1s)
	if err := os.Setenv("AUCTION_INTERVAL", "1s"); err != nil {
		t.Fatalf("erro ao setar AUCTION_INTERVAL: %v", err)
	}

	repo := auction.NewAuctionRepository(db)

	now := time.Now().UTC()
	auctionEntity := &auction_entity.Auction{
		Id:          "test-auction-" + now.Format("150405"),
		ProductName: "Produto Teste",
		Category:    "categoria-teste",
		Description: "descricao teste",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   now,
	}

	// chama CreateAuction que insere e dispara a goroutine que fecha o leilão
	if ierr := repo.CreateAuction(context.Background(), auctionEntity); ierr != nil {
		// imprime o erro interno se houver detalhe dentro (o seu internal_error não está expondo o err do driver;
		// aqui fornecemos instrução de debug)
		t.Fatalf("CreateAuction retornou erro: %v\n(Verifique se a coleção/DB está acessível e se as versões do driver e do Mongo são compatíveis).", ierr)
	}

	// Verificação 1: pouco antes do tempo de fechamento o leilão ainda deve estar aberto.
	time.Sleep(300 * time.Millisecond)

	var docBefore bson.M
	err = coll.FindOne(context.Background(), bson.M{"_id": auctionEntity.Id}).Decode(&docBefore)
	if err != nil {
		t.Fatalf("erro lendo documento do mongo (antes): %v", err)
	}

	statusBefore := fmt.Sprintf("%v", docBefore["status"])
	if statusBefore == fmt.Sprintf("%v", auction_entity.Completed) {
		t.Fatalf("status já é Completed antes do intervalo (esperado Open), got: %s", statusBefore)
	}

	// Espera até o fechamento automático: usa timeout para não travar
	waitTimeout := time.After(10 * time.Second)
	tick := time.Tick(250 * time.Millisecond)

	for {
		select {
		case <-waitTimeout:
			t.Fatalf("leilão não foi fechado automaticamente dentro do timeout")
		case <-tick:
			var doc bson.M
			err := coll.FindOne(context.Background(), bson.M{"_id": auctionEntity.Id}).Decode(&doc)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					// documento ainda não presente (comportamento improvável), continua
					continue
				}
				t.Fatalf("erro lendo documento do mongo: %v", err)
			}

			dbStatusStr := fmt.Sprintf("%v", doc["status"])
			expectedStatusStr := fmt.Sprintf("%v", auction_entity.Completed)

			if dbStatusStr == expectedStatusStr {
				// sucesso
				t.Logf("leilão %s fechado automaticamente (status=%s)", auctionEntity.Id, dbStatusStr)
				return
			}
			// continua polling
		}
	}
}
