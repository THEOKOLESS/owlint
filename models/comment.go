package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TargetID    string             `bson:"targetId" json:"targetId"`
	AuthorID    string             `bson:"authorId" json:"authorId"`
	PublishedAt string             `bson:"publishedAt" json:"publishedAt"`
	TextEn      string             `bson:"textEn" json:"textEn"`
	TextFr      string             `bson:"textFr" json:"textFr"`
	Replies     []Comment          `bson:"replies,omitempty" json:"replies,omitempty"`
}
