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
	"time"
)

type OfferPatchRequest struct {

	Title string `json:"title,omitempty"`

	Description string `json:"description,omitempty"`

	Giz int32 `json:"giz,omitempty"`

	ChatUrl string `json:"chat_url,omitempty"`

	IsPublic bool `json:"is_public,omitempty"`

	Deadline time.Time `json:"deadline,omitempty"`
}
