package dbprovider

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Firestoreの初期化
func Init(ctx context.Context) (*firestore.Client, error) {

	// サービスアカウント読み込み
	sa := option.WithCredentialsFile("path/to/GcpAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return client, nil
}

func Read() map[string]interface{} {
	ctx := context.Background()
	// 初期化する
	cilent, err := Init(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	//TODO: ID固定値
	dsnap, err := cilent.Collection("Issue").Doc("UYku2uPU4QdAP3rylfyX").Get(ctx)
	if err != nil {
		log.Fatalf("Failed to iterate: %v", err)
	}
	data := dsnap.Data()

	defer cilent.Close()

	return data
}


// コレクション全ての読み込み処理
func AllRead() map[string][]interface{} {
	ctx := context.Background()
	// 初期化する
	cilent, err := Init(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	iter := cilent.Collection("Issue").Documents(ctx)
	// var allData []interface {}
	res := make(map[string][]interface{})
	for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed to iterate: %v", err)
			}
			fmt.Println(doc.Data())
			res["value"] = append(res["value"],  doc.Data())
			// allData = doc.Data()
		}
	defer cilent.Close()

	return res
}


func Update(id string, body string, isClosed bool) error {
	ctx := context.Background()
	cilent, err := Init(ctx)
	if err != nil {
		defer cilent.Close()
		log.Fatalln(err)
		return fmt.Errorf("Init fail: %w", err)
	}

	var p []firestore.Update
	time := time.Now()
	if isClosed {
		p = []firestore.Update{
			{
				Path: "body",
				Value: body,
			},
			{
				Path: "updateAt",
				Value: time,
			},
			{
				Path: "isClosed",
				Value: isClosed,
			},
			{
				Path: "closedAt",
				Value: time,
			},
		}
	} else {
		p = []firestore.Update{
			{
				Path: "body",
				Value: body,
			},
			{
				Path: "updateAt",
				Value: time,
			},
		}
	}

	_, err = cilent.Collection("Issue").Doc("UYku2uPU4QdAP3rylfyX").Update(ctx, p)
	defer cilent.Close()
	if err != nil {
		log.Fatalln(err)
		return fmt.Errorf("update fail: %w", err)
	}

	return nil
}