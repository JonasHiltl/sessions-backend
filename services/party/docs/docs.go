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
                "description": "Create a new Party in DB",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Create a party",
                "parameters": [
                    {
                        "description": "The body to create a party",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreatePartyBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/datastruct.PublicParty"
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
        "/near": {
            "get": {
                "description": "Get a list of parties near a location",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GEO"
                ],
                "summary": "Search by location",
                "parameters": [
                    {
                        "type": "number",
                        "description": "Latitude",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Longitude",
                        "name": "long",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Geohash precision",
                        "name": "precision",
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
                                "$ref": "#/definitions/datastruct.PublicParty"
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
        "/{pId}": {
            "get": {
                "description": "Get a Party by it's id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Get party",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Party Id",
                        "name": "pId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/datastruct.PublicParty"
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
                "description": "Deletes a party from our db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Delete a party",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Party Id",
                        "name": "pId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/comtypes.MessageRes"
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
                "description": "Updates a party with the provided values",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Update a Party",
                "parameters": [
                    {
                        "description": "The body to create a Party",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UpdatePartyBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/datastruct.PublicParty"
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
        "comtypes.MessageRes": {
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
        "datastruct.PublicParty": {
            "type": "object",
            "properties": {
                "creatorId": {
                    "type": "string"
                },
                "id": {
                    "description": "PARTY#CreatorId",
                    "type": "string"
                },
                "isPublic": {
                    "type": "boolean"
                },
                "lat": {
                    "type": "number"
                },
                "long": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                },
                "ttl": {
                    "type": "string"
                }
            }
        },
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "handler.CreatePartyBody": {
            "type": "object",
            "required": [
                "lat",
                "long",
                "title"
            ],
            "properties": {
                "isPublic": {
                    "type": "boolean"
                },
                "lat": {
                    "type": "number"
                },
                "long": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "handler.UpdatePartyBody": {
            "type": "object",
            "required": [
                "lat",
                "long",
                "title"
            ],
            "properties": {
                "isPublic": {
                    "type": "boolean"
                },
                "lat": {
                    "type": "number"
                },
                "long": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8082",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Party Microservice",
	Description:      "This Microservice manages Party entities",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
