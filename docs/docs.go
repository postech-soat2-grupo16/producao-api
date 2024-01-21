// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://fastfood.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.fastfood.io/support",
            "email": "support@fastfood.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/producao_pedidos": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProducaoPedidos"
                ],
                "summary": "Get all producao_pedidos",
                "operationId": "get-all-producao_pedidos",
                "parameters": [
                    {
                        "type": "string",
                        "description": "status search by status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/producaopedido.ProducaoPedido"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProducaoPedidos"
                ],
                "summary": "New producaopedido",
                "operationId": "create-producaopedido",
                "parameters": [
                    {
                        "description": "ProducaoPedido data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/producaopedido.ProducaoPedido"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/producaopedido.ProducaoPedido"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/producao_pedidos/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProducaoPedidos"
                ],
                "summary": "Get a producaopedido by ID",
                "operationId": "get-producaopedido-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ProducaoPedido ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/producaopedido.ProducaoPedido"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProducaoPedidos"
                ],
                "summary": "Update a producaopedido",
                "operationId": "update-producaopedido",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ProducaoPedido ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "ProducaoPedido data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/producaopedido.ProducaoPedido"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/producaopedido.ProducaoPedido"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProducaoPedidos"
                ],
                "summary": "Delete a producaopedido by ID",
                "operationId": "delete-producaopedido-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ProducaoPedido ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "producaopedido.ProducaoPedido": {
            "type": "object",
            "properties": {
                "pedido_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Producao API",
	Description:      "Here you will find everything you need to have the best possible integration with our APIs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
