package issue

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

func Read(id int)map[string]interface{} {
	ctx := context.Background()
	// 初期化する
	client, err := Init(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	res := make(map[string]interface{})
	iter := client.Collection("Issue").Where("id", "==", id).Documents(ctx)
	for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed to iterate: %v", err)
			}
			fmt.Println(doc.Ref.ID)
			fmt.Println(doc.Data())
			res = doc.Data()
			res["documentID"] = doc.Ref.ID
	}
	defer client.Close()

	return res
}


// コレクション全ての読み込み処理
func AllRead() map[string][]interface{} {
	ctx := context.Background()
	// 初期化する
	client, err := Init(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	iter := client.Collection("Issue").Documents(ctx)
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
	}
	
	defer client.Close()

	return res
}

// 解決済みのissueを表示する
func ClosedRead() map[string][]interface{} {
	ctx := context.Background()
	// 初期化する
	client, err := Init(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	
	query := client.Collection("Issue").Where("isClosed", "==", true)
	iter := query.Documents(ctx)
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
	}
	defer client.Close()

	return res
}

// 解決していないissueを表示する
func OpenRead() map[string][]interface{} {
	ctx := context.Background()
	// 初期化する
	client, err := Init(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	
	query := client.Collection("Issue").Where("isClosed", "==", false)
	iter := query.Documents(ctx)
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
	}
	defer client.Close()

	return res
}

// 新しいissueを作成する
func Create(title string, body string) error {
	ctx := context.Background()
	client, err := Init(ctx)
	if err != nil {
		defer client.Close()
		log.Fatalln(err)
		return fmt.Errorf("Init fail: %w", err)
	}

	// 連番になるIDを生成する
	collection := client.Collection("Issue")
	// 値が一番大きいIDを取得する
	query := collection.OrderBy("id", firestore.Desc).Limit(1)
	iter := query.Documents(ctx)
	doc, err := iter.Next()
	if 	err != nil {
		return err
	}
	data := doc.Data()
	// interface型なので明示的に型を書く
	id := data["id"].(int64) + 1

	time := time.Now()

	_, _, aderr := collection.Add(ctx, map[string]interface{}{
		"id": id,
        "title": title,
        "body": body,
		"createAt": time,
		"updateAt": time,
		"isClosed": false,
	})
	if aderr != nil {
		return fmt.Errorf("create fail: %w", aderr)
	}

	return nil
}

// issueのコメントupdate
func UpdComment(id string, body string) error {
	ctx := context.Background()
	client, err := Init(ctx)
	if err != nil {
		defer client.Close()
		log.Fatalln(err)
		return fmt.Errorf("Init fail: %w", err)
	}

	var p []firestore.Update
	time := time.Now()
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

	_, err = client.Collection("Issue").Doc(id).Update(ctx, p)
	defer client.Close()
	if err != nil {
		log.Fatalln(err)
		return fmt.Errorf("update fail: %w", err)
	}

	return nil
}

func UpdClose(id string, isClosed bool) error {
	ctx := context.Background()
	client, err := Init(ctx)
	if err != nil {
		defer client.Close()
		log.Fatalln(err)
		return fmt.Errorf("Init fail: %w", err)
	}
	
	return nil
}