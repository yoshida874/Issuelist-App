package dbprovider

import (
	"context"
	"fmt"
	"log"

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

func Read() map[string]interface{} {
	ctx := context.Background()
	// 初期化する
	cilent, err := Init(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	dsnap, err := cilent.Collection("Issue").Doc("UYku2uPU4QdAP3rylfyX").Get(ctx)
	if err != nil {
		log.Fatalf("Failed to iterate: %v", err)
	}
	data := dsnap.Data()

	defer cilent.Close()

	return data
}
