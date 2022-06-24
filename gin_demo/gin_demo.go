package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Comment struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	CreateTime string `json:"create_time"`
	Content    string `json:"content"`
}

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 dbname=comment_system user=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := gin.Default()
	r.Use(Cors())
	r.POST("/comment/submit", func(ctx *gin.Context) {
		title, _ := ctx.GetPostForm("title")
		content, _ := ctx.GetPostForm("content")
		println(title)
		println(content)
		sqlStr := "insert into review (title,content,create_time)values($1,$2,$3)"
		_, err := db.Exec(sqlStr, title, content, time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			log.Fatal(err)
		}
		// id, err := res.LastInsertId()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// println(id)
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "success",
		})
	})
	r.GET("/comment/list", func(ctx *gin.Context) {
		// 读取 comment列表
		sqlStr := "select * from review order by id desc"
		rows, err := db.Query(sqlStr)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var comments = make([]Comment, 0)
		var com = Comment{}
		for rows.Next() {
			if err = rows.Scan(&com.ID, &com.Title, &com.CreateTime, &com.Content); err != nil {
				log.Fatal(err)
			}
			// 放在 []Comment 中
			comments = append(comments, com)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"data":    comments,
			"message": "success",
		})
	})
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
