// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "zayyan",
            "email": "zayyanramadhan@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/comments": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Get Comment",
                "responses": {}
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Create Comment",
                "parameters": [
                    {
                        "description": "comment info",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CommentRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/comments/{commentId}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Update Comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "commentId",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "comment info",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CommentRequestUpdate"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Delete Comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "commentId",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/photos": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Get Photo",
                "responses": {}
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Create Photo",
                "parameters": [
                    {
                        "description": "Photo info",
                        "name": "Photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PhotoRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/photos/{photoId}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Update Photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "photoId",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Photo info",
                        "name": "Photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PhotoRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Delete Photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "photoId",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/socialmedias": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialmedia"
                ],
                "summary": "Get SocialMedia",
                "responses": {}
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialmedia"
                ],
                "summary": "Create SocialMedia",
                "parameters": [
                    {
                        "description": "SocialMedia info",
                        "name": "SocialMedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SocialMediaRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/socialmedias/{socialMediaId}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialmedia"
                ],
                "summary": "Update SocialMedia",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "socialMediaId",
                        "name": "socialMediaId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "SocialMedia info",
                        "name": "SocialMedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SocialMediaRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialmedia"
                ],
                "summary": "Delete SocialMedia",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "socialMediaId",
                        "name": "socialMediaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/users": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete",
                "responses": {}
            }
        },
        "/users/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLoginRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/users/logout": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Logout",
                "responses": {}
            }
        },
        "/users/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRegisterRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/users/{userId}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserUpdate"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.CommentRequest": {
            "type": "object",
            "required": [
                "message"
            ],
            "properties": {
                "message": {
                    "type": "string",
                    "example": "zayyan"
                },
                "photo_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.CommentRequestUpdate": {
            "type": "object",
            "required": [
                "message"
            ],
            "properties": {
                "message": {
                    "type": "string",
                    "example": "zayyan"
                }
            }
        },
        "models.PhotoRequest": {
            "type": "object",
            "required": [
                "photo_url",
                "title"
            ],
            "properties": {
                "caption": {
                    "type": "string",
                    "example": "zayyan"
                },
                "photo_url": {
                    "type": "string",
                    "example": "zayyan"
                },
                "title": {
                    "type": "string",
                    "example": "zayyan"
                }
            }
        },
        "models.SocialMediaRequest": {
            "type": "object",
            "required": [
                "name",
                "social_media_url"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "zayyan"
                },
                "social_media_url": {
                    "type": "string",
                    "example": "zayyan"
                }
            }
        },
        "models.UserLoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "zayyan@mail.com"
                },
                "password": {
                    "type": "string",
                    "example": "zayyan"
                }
            }
        },
        "models.UserRegisterRequest": {
            "type": "object",
            "required": [
                "age",
                "email",
                "password",
                "username"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 9
                },
                "email": {
                    "type": "string",
                    "example": "zayyan@mail.com"
                },
                "password": {
                    "type": "string",
                    "minLength": 6,
                    "example": "zayyan"
                },
                "username": {
                    "type": "string",
                    "example": "zayyan"
                }
            }
        },
        "models.UserUpdate": {
            "type": "object",
            "required": [
                "email",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "zayyan@mail.com"
                },
                "username": {
                    "type": "string",
                    "example": "zayyan"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v2.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MyGram API",
	Description:      "Sample API Spec for MyGram",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
