// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/books/delete/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiTokenAuth": []
                    }
                ],
                "description": "endpoint for delete book by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Delete book by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/books/list": {
            "get": {
                "security": [
                    {
                        "ApiTokenAuth": []
                    }
                ],
                "description": "endpoint for getting all books",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Books list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Book"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/books/update": {
            "post": {
                "security": [
                    {
                        "ApiTokenAuth": []
                    }
                ],
                "description": "endpoint for update book by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Update book by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book's id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Book's title",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Book's release year",
                        "name": "releaseYear",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Book's author",
                        "name": "authorId",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/books/{id}": {
            "get": {
                "security": [
                    {
                        "ApiTokenAuth": []
                    }
                ],
                "description": "endpoint for getting book by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Books by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/model.Book"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/users/sign-in": {
            "post": {
                "description": "endpoint for user sign in",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User sign in",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User's email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User's password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/service.Token"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/users/sign-up": {
            "post": {
                "description": "endpoint for user sign up",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User sign up",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User's email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User's password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Author": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string",
                    "example": "1948-09-20"
                },
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Book"
                    }
                },
                "first_name": {
                    "type": "string",
                    "example": "George"
                },
                "id": {
                    "type": "integer",
                    "example": 10
                },
                "is_male": {
                    "type": "boolean",
                    "example": true
                },
                "last_name": {
                    "type": "string",
                    "example": "Martin"
                }
            }
        },
        "model.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/model.Author"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "release_year": {
                    "type": "integer",
                    "example": 1997
                },
                "title": {
                    "type": "string",
                    "example": "Games of thrones"
                }
            }
        },
        "service.Token": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJcdTAwMDEifQ.O3Lp6-VLUmGrcHUO11b2LqUBXq-tQeyRTWxwAe6aNLQ"
                }
            }
        },
        "v1.ApiResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "object"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiTokenAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:8080",
	BasePath:    "/api/v1/",
	Schemes:     []string{},
	Title:       "Books & Authors API",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
