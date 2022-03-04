package controllers

import (
	"context"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"alef_education_devops_challenge/database"

	"alef_education_devops_challenge/models"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

var taskCollection *mongo.Collection = database.OpenCollection(database.Client, "task")

var validate = validator.New()

//Add Task
func AddTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var task models.Tasks

		defer cancel()
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(task)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		task.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		task.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		task.ID = primitive.NewObjectID()
		task.Task_id = task.ID.Hex()

		resultInsertionNumber, insertErr := taskCollection.InsertOne(ctx, task)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "task was not created"})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

// //GetWishList to get a single wishlist of a user
// func GetWishList() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		whishListID := c.Param("whishListId")
// 		user_id := c.Param("user_id")

// 		var wishList models.WishList
// 		var user models.User

// 		check_user := userCollection.FindOne(ctx, bson.M{"user_id": user_id}).Decode(&user)
// 		defer cancel()
// 		if check_user != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "The user was not found"})
// 			return
// 		}

// 		filter := bson.D{{"user_id", user_id}, {"wishlist_id", whishListID}}
// 		err := wishListCollection.FindOne(ctx, filter).Decode(&wishList)

// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing wishlist"})
// 		}

// 		c.JSON(http.StatusOK, wishList)
// 	}
// }

// //Find all wishlist for a helpee
// func GetAllWishListByUserID() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		user_id := c.Param("user_id")
// 		var user models.User

// 		check_user := userCollection.FindOne(ctx, bson.M{"user_id": user_id}).Decode(&user)
// 		defer cancel()
// 		if check_user != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "The user was not found"})
// 			return
// 		}

// 		result, err := wishListCollection.Find(context.TODO(), bson.M{"user_id": user_id})

// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing invoice items"})
// 		}
// 		var allWishList []bson.M
// 		if err = result.All(ctx, &allWishList); err != nil {
// 			log.Fatal(err)
// 		}
// 		c.JSON(http.StatusOK, allWishList)
// 	}
// }

// //Check minimum wish list, confirm if the wishlist is 5
// func CountWhishList() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		whishListID := c.Param("whishListId")
// 		user_id := c.Param("user_id")

// 		filter := bson.D{{"user_id", user_id}, {"wishlist_id", whishListID}}
// 		count, err := wishListCollection.CountDocuments(ctx, filter)

// 		defer cancel()
// 		if err != nil {
// 			log.Panic(err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking counting documents"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, count)
// 	}
// }

// // Delete a single wish list
// func DeleteWishListByID() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		whishListID := c.Param("whishListId")
// 		user_id := c.Param("user_id")

// 		filter := bson.D{{"user_id", user_id}, {"wishlist_id", whishListID}}
// 		result, err := wishListCollection.DeleteOne(ctx, filter)

// 		defer cancel()
// 		if err != nil {
// 			log.Panic(err)
// 			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Error occured while trying to delete the document"})
// 			return
// 		}

// 		if result.DeletedCount > 0 {
// 			c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Wish list successfully deleted"})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{"status": "success", "message": "No wish list to be deleted"})
// 		}
// 	}
// }

// //UpdateWishList is the api used to update wishlist
// func UpdateWishList() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		var wishList models.WishList
// 		var interest models.Interest
// 		var user models.User

// 		wishListId := c.Param("whishListId")
// 		user_id := c.Param("user_id")

// 		if err := c.BindJSON(&wishList); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		var updateObj primitive.D

// 		if wishList.Name != nil {
// 			updateObj = append(updateObj, bson.E{"name", wishList.Name})
// 		}

// 		wishList.Date, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 		updateObj = append(updateObj, bson.E{"date", wishList.Date})

// 		if wishList.Description != nil {
// 			updateObj = append(updateObj, bson.E{"description", wishList.Description})
// 		}

// 		wishList.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 		updateObj = append(updateObj, bson.E{"created_at", wishList.Created_at})

// 		wishList.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 		updateObj = append(updateObj, bson.E{"updated_at", wishList.Updated_at})

// 		//Check if interest_id exist
// 		if wishList.Interest_id != "" {
// 			err := interestsCollection.FindOne(ctx, bson.M{"interest_id": wishList.Interest_id}).Decode(&interest)
// 			defer cancel()
// 			if err != nil {
// 				msg := fmt.Sprintf("message: interest was not found")
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 				return
// 			}
// 			updateObj = append(updateObj, bson.E{"interest_id", wishList.Interest_id})
// 		}

// 		//Check if User_id exist
// 		if wishList.User_id != nil {
// 			err := userCollection.FindOne(ctx, bson.M{"user_id": wishList.User_id}).Decode(&user)

// 			defer cancel()
// 			if err != nil {
// 				msg := fmt.Sprintf("message: user id was not found")
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 				return
// 			}
// 			updateObj = append(updateObj, bson.E{"user_id", wishList.User_id})
// 		}

// 		if user_id != *wishList.User_id {
// 			c.JSON(http.StatusBadRequest, gin.H{"Status": "Failed", "message": "The user id and wishlist user id is not equal"})
// 			return
// 		}

// 		//Check if Wishlist_id exist
// 		if wishList.Wishlist_id != "" {
// 			err := wishListCollection.FindOne(ctx, bson.M{"wishlist_id": wishList.Wishlist_id}).Decode(&interest)
// 			defer cancel()
// 			if err != nil {
// 				msg := fmt.Sprintf("message: wishlist id was not found")
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 				return
// 			}
// 			updateObj = append(updateObj, bson.E{"wishlist_id", wishList.Wishlist_id})
// 		}

// 		upsert := true
// 		filter := bson.M{"wishlist_id": wishListId}
// 		opt := options.UpdateOptions{
// 			Upsert: &upsert,
// 		}

// 		result, err := wishListCollection.UpdateOne(
// 			ctx,
// 			filter,
// 			bson.D{
// 				{"$set", updateObj},
// 			},
// 			&opt,
// 		)

// 		if err != nil {
// 			msg := fmt.Sprintf("Wish list item update failed")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 			return
// 		}

// 		if result.ModifiedCount > 0 {
// 			c.JSON(http.StatusOK, gin.H{"status": "Successful", "message": "The wish list is successfully updated"})
// 		}
// 	}
// }
