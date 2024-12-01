package controllers

import (
	"context"
	"net/http"
	"owlint/models"
	"owlint/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddComment(c *gin.Context) {
	targetId := c.Param("targetId")

	var comment models.Comment

	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.TargetID = targetId

	if comment.TextEn == "" && comment.TextFr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "textEn AND textFr can't be empty"})
		return
	}

	// Determine the source language and translate here

	if comment.AuthorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authorId is required"})
		return
	}

	if comment.PublishedAt == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PublishedAt is required"})
		return
	}
	//generate a new ID for the comment
	comment.ID = primitive.NewObjectID()
	commentCollection := utils.GetCollection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := commentCollection.InsertOne(ctx, comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while inserting comment"})
		return
	}
	//send the comment to the external service, it is async to not block the main thread
	go utils.SendCommentNotification(comment)

	c.JSON(http.StatusCreated, comment)
}

func GetComments(c *gin.Context) {
	targetId := c.Param("targetId")

	commentCollection := utils.GetCollection("comments")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filter comments with the targetId given
	filter := bson.M{"targetId": targetId}

	// cursor is pointing to the first document in the collection
	cursor, err := commentCollection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while fetching comments"})
		return
	}
	defer cursor.Close(ctx)

	var comments []models.Comment
	if err = cursor.All(ctx, &comments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while reading comments"})
		return
	}

	// for each comment, fetch its replies recursively
	for i, comment := range comments {
		comments[i].Replies, err = fetchReplies(ctx, comment.ID.Hex(), commentCollection)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while fetching replies"})
			return
		}
	}

	c.JSON(http.StatusOK, comments)
}

func fetchReplies(ctx context.Context, commentId string, commentCollection *mongo.Collection) ([]models.Comment, error) {

	filter := bson.M{"targetId": commentId}

	cursor, err := commentCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var replies []models.Comment
	if err = cursor.All(ctx, &replies); err != nil {
		return nil, err
	}

	// for each reply, fetch its replies recursively
	for i, reply := range replies {
		replies[i].Replies, err = fetchReplies(ctx, reply.ID.Hex(), commentCollection)
		if err != nil {
			return nil, err
		}
	}

	return replies, nil
}
