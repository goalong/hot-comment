package api

import (
	"github.com/goalong/hot-comment/router/e"
	"net/http"
	"strconv"
)
import "github.com/gin-gonic/gin"
import "github.com/goalong/hot-comment/search"

type Page struct {
	PageNum int `form:"page_num" binding:"required,max=50,min=1"` // 限制页码和每页条数的范围
	PageSize int `form:"page_size" binding:"required,max=50,min=1"`
}

func GetPage(c *gin.Context) (pageNum, pageSize, code int) {
	code = e.SUCCESS
	var page Page
	err := c.ShouldBindQuery(&page)
	if err != nil {
		code = e.INVALID_PAGE
	}
	return page.PageNum, page.PageSize, code

}

func AddPageToResp(resp map[string]interface{}, pageNum, pageSize int) map[string]interface{} {
	resp["page_num"] = pageNum
	resp["page_size"] = pageSize
	return resp
}

// @Summary 热门评论，按点赞数排行
// @Produce  json
// @Param page_num query int true "页码"
// @Param page_size query int true "每页数量"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/hot_comments [get]
func GetHotComments(c *gin.Context) {
	pageNum, pageSize, code := GetPage(c)
	data := make(map[string]interface{})
	if code == e.SUCCESS {
		data, code = search.GetCommentsByLikeCount(pageNum, pageSize)
	}
	data = AddPageToResp(data, pageNum, pageSize)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary 获取热门歌曲，按评论数排行
// @Produce  json
// @Param page_num query int true "页码"
// @Param page_size query int true "每页数量"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/hot_songs [get]
func GetHotSongs(c *gin.Context) {
	data := make(map[string]interface{})
	pageNum, pageSize, code := GetPage(c)
	if code == e.SUCCESS {
		data, code = search.GetSongsByCommentCount(pageNum, pageSize)
	}
	data = AddPageToResp(data, pageNum, pageSize)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary 按关键词搜索歌曲
// @Produce  json
// @Param page_num query int true "页码"
// @Param page_size query int true "每页数量"
// @Param keyword query string true "关键词"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/search/songs [get]
func SearchSongs(c *gin.Context) {
	data := make(map[string]interface{})
	pageNum, pageSize, code := GetPage(c)
	keyword := c.Query("keyword")

	if code == e.SUCCESS {
		data, code = search.SearchSong(keyword, pageNum, pageSize)
	}
	data = AddPageToResp(data, pageNum, pageSize)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})


}

// @Summary 按关键词搜索评论
// @Produce  json
// @Param page_num query int true "页码"
// @Param page_size query int true "每页数量"
// @Param keyword query string true "关键词"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/search/comments [get]
func SearchComments(c *gin.Context) {
	data := make(map[string]interface{})
	pageNum, pageSize, code := GetPage(c)
	keyword := c.Query("keyword")

	if code == e.SUCCESS {
		data, code = search.SearchComment(keyword, pageNum, pageSize)
	}
	data = AddPageToResp(data, pageNum, pageSize)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})


}

// @Summary 搜索歌手
// @Produce  json
// @Param page_num query int true "页码"
// @Param page_size query int true "每页数量"
// @Param keyword query string true "关键词"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/search/artists [get]
func SearchArtists(c *gin.Context) {
	data := make(map[string]interface{})
	pageNum, pageSize, code := GetPage(c)
	keyword := c.Query("keyword")

	if code == e.SUCCESS {
		data, code = search.SearchArtist(keyword, pageNum, pageSize)
	}
	data = AddPageToResp(data, pageNum, pageSize)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// @Summary 根据歌手ID获取歌曲列表
// @Produce  json
// @Param page_num query int true "页码"
// @Param page_size query int true "每页数量"
// @Param artist_id query string true "歌手ID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/songs [get]
func GetSongsByArtistId(c *gin.Context) {
	data := make(map[string]interface{})
	pageNum, pageSize, code := GetPage(c)
	artistId, err := strconv.Atoi(c.Query("artist_id"))
	if err != nil {
		code = e.INVALID_PARAMS
	}
	if code == e.SUCCESS {
		data, code = search.GetSongsByArtistId(artistId, pageNum, pageSize)
	}
	data = AddPageToResp(data, pageNum, pageSize)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}



// @Summary 根据歌曲ID获取评论列表
// @Produce  json
// @Param page_num query int true "页码"
// @Param page_size query int true "每页数量"
// @Param song_id query string true "歌曲ID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/comments [get]
func GetCommentsBySongId(c *gin.Context) {
	data := make(map[string]interface{})
	pageNum, pageSize, code := GetPage(c)
	songId, err := strconv.Atoi(c.Query("song_id"))
	if err != nil {
		code = e.INVALID_PARAMS
	}

	if code == e.SUCCESS {
		data, code = search.GetCommentsBySongId(songId, pageNum, pageSize)
	}
	data = AddPageToResp(data, pageNum, pageSize)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}