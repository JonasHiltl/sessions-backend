// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
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
        "/": {
            "post": {
                "description": "Saves a user in our DB",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "The body to create a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/datastruct.RequestUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/datastruct.PublicUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login with credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "The required credentials",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/datastruct.LoginBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/datastruct.AuthRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Saves a new User in the Database and returns auth token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new User",
                "parameters": [
                    {
                        "description": "The body to create a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/datastruct.RequestUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/datastruct.AuthRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/friend/search/{id}": {
            "get": {
                "description": "Search friends by username, first name or last name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friend"
                ],
                "summary": "Search friends of user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of user which friends to search",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "If only accepted friends to search",
                        "name": "accepted",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Query string",
                        "name": "query",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/datastruct.PublicUser"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/friend/{id}": {
            "put": {
                "description": "Sends friend request to user with id from params",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friend"
                ],
                "summary": "Send friend request",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Friend Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/datastruct.MessageRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/me": {
            "get": {
                "description": "Gets the user information of currently logged in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get current user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Base64 encoded JWT Payload",
                        "name": "jwt_payload",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/datastruct.PublicUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/username-exists/{username}": {
            "get": {
                "description": "Check if username is already taken",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Check Username availability",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/{id}": {
            "get": {
                "description": "Gets the user information by id user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Get user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/datastruct.PublicUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a user from our DB",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/datastruct.PublicUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates user with provided values",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "description": "The body to create a thing",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/datastruct.RequestUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/datastruct.PublicUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "datastruct.AuthRes": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "datastruct.LoginBody": {
            "type": "object",
            "required": [
                "password",
                "usernameOrEmail"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "usernameOrEmail": {
                    "type": "string"
                }
            }
        },
        "datastruct.MessageRes": {
            "type": "object",
            "required": [
                "message"
            ],
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "datastruct.PublicUser": {
            "type": "object",
            "required": [
                "firstame",
                "username"
            ],
            "properties": {
                "firstame": {
                    "type": "string"
                },
                "friendCount": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "datastruct.RequestUser": {
            "type": "object",
            "required": [
                "email",
                "firstame",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstame": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "picture": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "User Microservice",
	Description:      "This Microservice manages user entities",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
