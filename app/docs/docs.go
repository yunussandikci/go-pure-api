// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/in-memory": {
            "get": {
                "description": "This endpoints returns value of the key provided",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "values"
                ],
                "summary": "Gets a value of the key provided",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Key",
                        "name": "key",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.InMemoryRecordResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ApiError"
                        }
                    }
                }
            },
            "post": {
                "description": "This endpoints persists a new key-value in the in-memory database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "values"
                ],
                "summary": "Creates a a new key-value",
                "parameters": [
                    {
                        "description": "The key and value that will be persist.",
                        "name": "Value",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.InMemoryCreateRecordRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.InMemoryRecordResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ApiError"
                        }
                    }
                }
            }
        },
        "/mongo": {
            "post": {
                "description": "This endpoints returns records from the mongo database with the provided filter in request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "values"
                ],
                "summary": "Gets a records from database",
                "parameters": [
                    {
                        "description": "Filter for the request",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.MongoGetRecordsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.MongoRecordsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.ApiError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "request.InMemoryCreateRecordRequest": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "request.MongoGetRecordsRequest": {
            "type": "object",
            "required": [
                "endDate",
                "maxCount",
                "minCount",
                "startDate"
            ],
            "properties": {
                "endDate": {
                    "type": "string",
                    "example": "2021-01-02"
                },
                "maxCount": {
                    "type": "integer",
                    "example": 3000
                },
                "minCount": {
                    "type": "integer",
                    "example": 2800
                },
                "startDate": {
                    "type": "string",
                    "example": "2016-10-02"
                }
            }
        },
        "response.InMemoryRecordResponse": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "response.MongoRecordResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                },
                "totalCount": {
                    "type": "integer"
                }
            }
        },
        "response.MongoRecordsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.MongoRecordResponse"
                    }
                }
            }
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
	Host:        "go-pure-api.herokuapp.com",
	BasePath:    "/api/v1",
	Schemes:     []string{"https"},
	Title:       "Go Pure API",
	Description: "A REST API that allows you to get records from mongo database and read/write them to in-memory database!",
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
