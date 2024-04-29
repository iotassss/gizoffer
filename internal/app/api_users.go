/*
 * Gizoffer
 *
 * Gizoffer is a platform that helps you find the best deals with Gizomo employees using GIZ.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package app

import (
	"github.com/gin-gonic/gin"
)

type UsersAPI struct {
}

// Get /users
// Get all users 
func (api *UsersAPI) UsersGet(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Delete /users/:id
// Delete a user by ID 
func (api *UsersAPI) UsersIdDelete(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /users/:id
// Get a user by ID 
func (api *UsersAPI) UsersIdGet(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Put /users/:id
// Update a user by ID 
func (api *UsersAPI) UsersIdPut(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Post /users
// Create a user 
func (api *UsersAPI) UsersPost(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

