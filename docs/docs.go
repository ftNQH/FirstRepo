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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/addItem": {
            "post": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "ping example",
                "parameters": [
                    {
                        "description": "Thông tin item mới ",
                        "name": "NewItem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Item"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Item"
                        }
                    }
                }
            }
        },
        "/delete/:id": {
            "delete": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "ping example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id cần xóa  ",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Item"
                        }
                    }
                }
            }
        },
        "/editItem/:id": {
            "put": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "ping example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID cần sửa",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Thông tin cần sửa",
                        "name": "ItemEdit",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Item"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Item"
                        }
                    }
                }
            }
        },
        "/item": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Item"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Item": {
            "type": "object",
            "properties": {
                "arr_owner_address": {
                    "type": "object",
                    "properties": {
                        "0xce85a994e411901fc9bda228d07f7b66f8320ecd": {
                            "type": "integer"
                        }
                    }
                },
                "arr_owner_uid": {
                    "type": "object",
                    "properties": {
                        "221139": {
                            "type": "integer"
                        }
                    }
                },
                "category_type": {
                    "type": "integer"
                },
                "chain": {
                    "type": "string"
                },
                "collection_id": {
                    "type": "integer"
                },
                "contract_address": {
                    "type": "string"
                },
                "create_time": {
                    "type": "integer"
                },
                "creator_address": {
                    "type": "string"
                },
                "creator_uid": {
                    "type": "integer"
                },
                "currency_address": {
                    "type": "object",
                    "properties": {
                        "address": {
                            "type": "string"
                        },
                        "chain": {
                            "type": "string"
                        },
                        "decimal": {
                            "type": "integer"
                        },
                        "img": {
                            "type": "string"
                        },
                        "name": {
                            "type": "string"
                        },
                        "status": {
                            "type": "boolean"
                        },
                        "symbol": {
                            "type": "string"
                        }
                    }
                },
                "edition": {
                    "type": "integer"
                },
                "end_at": {
                    "type": "integer"
                },
                "event_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "extend_data": {
                    "type": "object",
                    "properties": {
                        "category": {
                            "type": "string"
                        },
                        "collection": {
                            "type": "string"
                        },
                        "create": {
                            "type": "string"
                        },
                        "like": {
                            "type": "string"
                        },
                        "owner": {
                            "type": "string"
                        },
                        "total_offer": {
                            "type": "string"
                        }
                    }
                },
                "external_link": {
                    "type": "string"
                },
                "hidden": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "is_loot_box": {
                    "type": "boolean"
                },
                "is_sensitive": {
                    "type": "boolean"
                },
                "media_type": {
                    "type": "integer"
                },
                "meta_data": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner_address": {
                    "type": "string"
                },
                "owner_uid": {
                    "type": "integer"
                },
                "price": {
                    "type": "string"
                },
                "product_no": {
                    "type": "string"
                },
                "start_at": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "token_id": {
                    "type": "string"
                },
                "total_claim": {
                    "type": "integer"
                },
                "total_edition": {
                    "type": "integer"
                },
                "total_like": {
                    "type": "integer"
                },
                "total_limit": {
                    "type": "integer"
                },
                "total_view": {
                    "type": "integer"
                },
                "type_nft": {
                    "type": "integer"
                },
                "unlockable_content": {
                    "type": "string"
                },
                "update_time": {
                    "type": "integer"
                },
                "uri": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}