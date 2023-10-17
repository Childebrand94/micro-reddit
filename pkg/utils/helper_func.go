package utils

import "github.com/Childebrand94/micro-reddit/pkg/models"

func CombinedPostComments(allPosts []models.Post, allComments []models.Comment) []models.PostWithComments {
	var result []models.PostWithComments

	for _, post := range allPosts {
		pwc := models.PostWithComments{
			Post: post,
		}
		for _, comment := range allComments {
			if comment.Post_ID == post.ID {
				pwc.Comments = append(pwc.Comments, comment)
			}
		}
		result = append(result, pwc)
	}

	return result
}
