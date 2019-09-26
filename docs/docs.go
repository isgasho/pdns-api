// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-09-26 11:57:44.449235 +0900 JST m=+0.183197595

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is PDNS RESTful API Server.",
        "title": "PDNS-API",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "127.0.0.1:8080",
    "basePath": "/v1",
    "paths": {
        "/domains": {
            "get": {
                "security": [
                    {
                        "ID": []
                    },
                    {
                        "Secret": []
                    }
                ],
                "description": "get domains",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get domains",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Domain ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Domain"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ID": []
                    },
                    {
                        "Secret": []
                    }
                ],
                "description": "create domain",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create domain",
                "parameters": [
                    {
                        "description": "Domain Object",
                        "name": "domain",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Domain"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Domain"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    }
                }
            }
        },
        "/domains/{name}": {
            "put": {
                "security": [
                    {
                        "ID": []
                    },
                    {
                        "Secret": []
                    }
                ],
                "description": "update domain",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "update domain",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dorain Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Domain Object",
                        "name": "domain",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Domain"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Domain"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ID": []
                    },
                    {
                        "Secret": []
                    }
                ],
                "description": "delete domain",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete domain",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Domain Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Domain"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    }
                }
            }
        },
        "/records": {
            "get": {
                "security": [
                    {
                        "ID": []
                    },
                    {
                        "Secret": []
                    }
                ],
                "description": "get records",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get records",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Record ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Domain ID",
                        "name": "domain_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Record"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ID": []
                    },
                    {
                        "Secret": []
                    }
                ],
                "description": "create record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create record",
                "parameters": [
                    {
                        "description": "Record Object",
                        "name": "record",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Record"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Record"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    }
                }
            }
        },
        "/records/disable/{id}": {
            "put": {
                "security": [
                    {
                        "ID": []
                    },
                    {
                        "Secret": []
                    }
                ],
                "description": "disable record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "disable record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Record ID ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Record Object",
                        "name": "record",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Record"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Record"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    }
                }
            }
        },
        "/records/enable/{id}": {
            "put": {
                "security": [
                    {
                        "ID": []
                    },
                    {
                        "Secret": []
                    }
                ],
                "description": "enable record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "enable record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Record ID ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Record Object",
                        "name": "record",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Record"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Record"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    }
                }
            }
        },
        "/records/{id}": {
            "put": {
                "security": [
                    {
                        "ID": []
                    },
                    {
                        "Secret": []
                    }
                ],
                "description": "update record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "update record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Record ID ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Record Object",
                        "name": "record",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Record"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Record"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ID": []
                    },
                    {
                        "Secret": []
                    }
                ],
                "description": "delete record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Record ID ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Record"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/pdns_api.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Domain": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "db": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_check": {
                    "type": "integer"
                },
                "master": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "notified_serial": {
                    "type": "integer"
                },
                "records": {
                    "type": "object",
                    "$ref": "#/definitions/model.Records"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.Record": {
            "type": "object",
            "properties": {
                "auth": {
                    "type": "boolean"
                },
                "content": {
                    "type": "string"
                },
                "db": {
                    "type": "string"
                },
                "disabled": {
                    "type": "boolean"
                },
                "domain_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "ordername": {
                    "type": "string"
                },
                "prio": {
                    "type": "integer"
                },
                "ttl": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.Records": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "auth": {
                        "type": "boolean"
                    },
                    "content": {
                        "type": "string"
                    },
                    "db": {
                        "type": "string"
                    },
                    "disabled": {
                        "type": "boolean"
                    },
                    "domain_id": {
                        "type": "integer"
                    },
                    "id": {
                        "type": "integer"
                    },
                    "name": {
                        "type": "string"
                    },
                    "ordername": {
                        "type": "string"
                    },
                    "prio": {
                        "type": "integer"
                    },
                    "ttl": {
                        "type": "integer"
                    },
                    "type": {
                        "type": "string"
                    }
                }
            }
        },
        "pdns_api.HTTPError": {
            "type": "object"
        }
    },
    "securityDefinitions": {
        "ID": {
            "type": "apiKey",
            "name": "PIR5-ID",
            "in": "header"
        },
        "Secret": {
            "type": "apiKey",
            "name": "PIR5-SECRET",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
