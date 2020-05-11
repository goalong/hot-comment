package search

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goalong/hot-comment/router/e"
	"github.com/olivere/elastic/v7"
	"reflect"
)


type Artist struct {
	ID     int      `json:"id"`
	Name  string   `json:"name"`
	//Songs []int `json:"songs"`
	CreateAt string `json:"create_at"`
	Pictrue string `json:"picture"`
	//Cls string `json:"_cls"`
}

type Song struct {
	ID     int      `json:"id"`
	Name  string   `json:"name"`
	CreateAt string `json:"create_at"`
	//Cls string `json:"_cls"`
	CommentCount  int  `json:"comment_count"`
	//Comments []int `json:"comments"`
}

type Comment struct {
	ID     int      `json:"id"`
	Content  string   `json:"content"`
	CreateAt string `json:"create_at"`
	//UpdateAt string `json:"update_at"`
	//Cls string `json:"_cls"`
	LikeCount  int  `json:"like_count"`
	Timestamp int `json:"timestamp"`
	User int `json:"user"`
	Song int `json:"song"`
}

func PrintQuery(src interface{}) {
	fmt.Println("*****")
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func GetESClient() (*elastic.Client, int){
	code := e.SUCCESS
	client, err := elastic.NewClient()
	if err != nil {
		code = e.INVALID_ES_CLIENT
	}
	return client, code
}

func Search(keyword string, index_name string, field_name string, indexStruct interface{}, pageNum, pageSize int) (map[string]interface{}, int) {
	query := elastic.NewMatchQuery(field_name, keyword)
	client, code := GetESClient()
	ctx := context.Background()
	result, err := client.Search().
		Index(index_name).
		Query(query).
		From((pageNum - 1)*pageSize).Size(pageSize).
		TrackTotalHits(true).
		Do(ctx)
	data := make(map[string]interface{})

	data["total"] = result.Hits.TotalHits.Value
	if err != nil {
		return data, code
	}
	var items []interface{}
	for _, item := range result.Each(reflect.TypeOf(indexStruct)) {
		items = append(items, item)
	}
	data["items"] = items
	return data, code
}

func SearchArtist(name string, pageNum int, pageSize int) (map[string]interface{}, int) {
	data, code := Search(name, "artist", "name", Artist{}, pageNum, pageSize)
	//var ret []Artist
	//arr := data["items"]
	//for _, item := range arr {
	//	if t, ok := item.(Artist); ok {
	//		ret = append(ret, t)
	//	}
	//}
	return data, code
}

func SearchSong(keyword string, pageNum int, pageSize int) (map[string]interface{}, int) {
	data, code := Search(keyword, "song", "name", Song{}, pageNum, pageSize)
	//var ret []Song
	//for _, item := range arr {
	//	if t, ok := item.(Song); ok {
	//		ret = append(ret, t)
	//	}
	//}
	return data, code
}

func SearchComment(keyword string, pageNum int, pageSize int) (map[string]interface{}, int) {
	data, code := Search(keyword, "comment", "content", Comment{}, pageNum, pageSize)
	//var ret []Comment
	//for _, item := range arr {
	//	if t, ok := item.(Comment); ok {
	//		ret = append(ret, t)
	//	}
	//}
	return data, code
}

// 按评论数对歌曲进行排序
func GetSongsByCommentCount(pageNum int, pageSize int) (map[string]interface{}, int) {
	client, code := GetESClient()
	termQuery := elastic.NewMatchAllQuery()
	ctx := context.Background()
	searchResult, err := client.Search().
		Index("song").
		Query(termQuery).
		Sort("comment_count", false).
		From((pageNum - 1)*pageSize).Size(pageSize).
		TrackTotalHits(true).
		Pretty(true).
		Do(ctx) // 执行
	if err != nil {
		panic(err)
	}
	fmt.Println(searchResult)
	var song Song
	var ret []Song
	for _, item := range searchResult.Each(reflect.TypeOf(song)) {
		if t, ok := item.(Song); ok {
			//fmt.Printf("Song %d: %d\n", t.ID, t.CommentCount)
			ret = append(ret, t)
		}
	}
	data := make(map[string]interface{})
	data["items"] = ret
	data["total"] = searchResult.Hits.TotalHits.Value
	return data, code
}

// 按点赞数排行获取评论
func GetCommentsByLikeCount(pageNum int, pageSize int)  (map[string]interface{}, int){
	client, code := GetESClient()
	termQuery := elastic.NewMatchAllQuery()
	ctx := context.Background()
	searchResult, err := client.Search().
		Index("comment").
		Query(termQuery).
		Sort("like_count", false).
		From((pageNum - 1)*pageSize).Size(pageSize).
		TrackTotalHits(true).
		Pretty(true).
		Do(ctx) // 执行
	if err != nil {
		panic(err)
	}
	//fmt.Println(searchResult)
	var comment Comment
	var ret []Comment
	for _, item := range searchResult.Each(reflect.TypeOf(comment)) {
		if t, ok := item.(Comment); ok {
			//fmt.Printf("Song %d: %d, %s\n", t.Song, t.LikeCount, t.Content)
			ret = append(ret, t)
		}
	}
	data := make(map[string]interface{})
	data["items"] = ret
	data["total"] = searchResult.Hits.TotalHits.Value
	return data, code
}

// 根据歌手ID获取歌曲列表
func GetSongsByArtistId(artistId int, pageNum int, pageSize int)  (map[string]interface{}, int) {
	client, code := GetESClient()
	termQuery := elastic.NewTermQuery("artist", artistId)
	ctx := context.Background()
	searchResult, err := client.Search().
		Index("song").
		Query(termQuery).
		From((pageNum - 1)*pageSize).Size(pageSize).
		TrackTotalHits(true).
		Pretty(true).
		Do(ctx) // 执行
	if err != nil {
		panic(err)
	}
	//fmt.Println(searchResult)
	var song Song
	var ret []Song
	for _, item := range searchResult.Each(reflect.TypeOf(song)) {
		if t, ok := item.(Song); ok {
			//fmt.Printf("Song %d: %d, %s\n", t.Song, t.LikeCount, t.Content)
			ret = append(ret, t)
		}
	}
	data := make(map[string]interface{})
	data["items"] = ret
	data["total"] = searchResult.Hits.TotalHits.Value
	return data, code
}

// 根据歌曲ID获取评论，按评论的点赞量排序
func GetCommentsBySongId(songId int, pageNum int, pageSize int) (map[string]interface{}, int) {
	client, code := GetESClient()
	termQuery := elastic.NewTermQuery("song", songId)
	ctx := context.Background()
	searchResult, err := client.Search().
		Index("comment").
		Query(termQuery).
		Sort("like_count", false). // todo, 排序可传参
		From((pageNum - 1)*pageSize).Size(pageSize).
		TrackTotalHits(true).
		Pretty(true).
		Do(ctx) // 执行
	if err != nil {
		panic(err)
	}
	//fmt.Println(searchResult)
	var comment Comment
	var ret []Comment
	for _, item := range searchResult.Each(reflect.TypeOf(comment)) {
		if t, ok := item.(Comment); ok {
			//fmt.Printf("Song %d: %d, %s\n", t.Song, t.LikeCount, t.Content)
			ret = append(ret, t)
		}
	}
	data := make(map[string]interface{})
	data["items"] = ret
	data["total"] = searchResult.Hits.TotalHits.Value
	return data, code
}
