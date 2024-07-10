// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
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
        "/tasks/start": {
            "post": {
                "description": "Start task with name for specified user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Start task",
                "parameters": [
                    {
                        "description": "task data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.StartTask"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "task started",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Retrieve users info with filtering and pagination support",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get users",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 1,
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 10,
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "г. Москва, ул. Ленина, д. 5, кв. 1",
                        "name": "address",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 1,
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "maxLength": 255,
                        "type": "string",
                        "example": "Иван",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "maxLength": 6,
                        "minLength": 6,
                        "type": "string",
                        "example": "567890",
                        "name": "passport_number",
                        "in": "query"
                    },
                    {
                        "maxLength": 4,
                        "minLength": 4,
                        "type": "string",
                        "example": "1234",
                        "name": "passport_series",
                        "in": "query"
                    },
                    {
                        "maxLength": 255,
                        "type": "string",
                        "example": "Иванович",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "maxLength": 255,
                        "type": "string",
                        "example": "Иванов",
                        "name": "surname",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.InternalServerErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create user by passport number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "Full Passport Number",
                        "name": "passportNumber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user created",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 1,
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.InternalServerErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.InternalServerErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update all or several user fields",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 1,
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user updated",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.InternalServerErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.InternalServerErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "the server encountered a problem and could not process your request"
                }
            }
        },
        "http.Response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "message"
                }
            }
        },
        "model.CreateUser": {
            "type": "object",
            "required": [
                "passportNumber"
            ],
            "properties": {
                "passportNumber": {
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11,
                    "example": "1234 567890"
                }
            }
        },
        "model.StartTask": {
            "type": "object",
            "required": [
                "name",
                "user_id"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 1,
                    "example": "do stuff"
                },
                "user_id": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 1
                }
            }
        },
        "model.Task": {
            "type": "object",
            "properties": {
                "end_time": {
                    "type": "string",
                    "example": "2025-07-10T07:00:43.047939731+03:00"
                },
                "name": {
                    "type": "string",
                    "example": "do stuff"
                },
                "start_time": {
                    "type": "string",
                    "example": "2024-07-10T07:00:43.047939731+03:00"
                },
                "task_id": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 1
                },
                "user_id": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 1
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "г. Москва, ул. Ленина, д. 5, кв. 1"
                },
                "id": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Иван"
                },
                "passport_number": {
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 6,
                    "example": "567890"
                },
                "passport_series": {
                    "type": "string",
                    "maxLength": 4,
                    "minLength": 4,
                    "example": "1234"
                },
                "patronymic": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Иванович"
                },
                "surname": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Иванов"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Time Tracker API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
