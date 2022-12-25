package main

import (
    "fmt"
    "log"
    "fun-in-golang/internal"
)

func main() {
    c := database.NewClient("db.json")
    err := c.EnsureDB()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("database ensured")

    user, err := c.CreateUser("test1@example.com", "password123", "Elliot Down", 24)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("user created", user)

    updatedUser, err := c.UpdateUser("test1@example.com", "password", "john doe", 18)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user updated", updatedUser)

	gotUser, err := c.GetUser("test1@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user got", gotUser)

	err = c.DeleteUser("test1@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user deleted")

	_, err = c.GetUser("test1@example.com")
	if err == nil {
		log.Fatal("shouldn't be able to get user that was deleted")
	}
	fmt.Println("user confirmed deleted")

    post, err := c.CreatePost("test@example.com", "my cat is way too fat")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("post created", post)

	secondPost, err := c.CreatePost("test@example.com", "my cat is getting skinny now")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("another post created", secondPost)

	posts, err := c.GetPosts("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got posts", posts)

	err = c.DeletePost(post.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted first post", posts)

	posts, err = c.GetPosts("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got posts", posts)

	err = c.DeletePost(secondPost.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted second post", posts)

	posts, err = c.GetPosts("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got posts", posts)

	err = c.DeleteUser("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user redeleted")
}
