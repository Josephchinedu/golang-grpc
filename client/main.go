package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/josephchinedu/golang-grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := pb.NewAddServiceClient(conn)

	g := gin.Default()

	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseInt(ctx.Param("a"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseInt(ctx.Param("b"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid Parameter B"})
			return
		}

		req := &pb.Request{A: int64(a), B: int64(b)}

		if response, err := client.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	})

	g.GET("/multiply/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseInt(ctx.Param("a"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseInt(ctx.Param("b"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid Parameter B"})
			return
		}

		req := &pb.Request{A: int64(a), B: int64(b)}

		if response, err := client.Multiply(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	})

	if err := g.Run(":8080"); err != nil {
		panic(err)
	}
}
