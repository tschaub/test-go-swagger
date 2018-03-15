// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Testing go-swagger",
    "title": "test-go-swagger",
    "version": "1.0.0"
  },
  "paths": {
    "/test": {
      "post": {
        "summary": "Create a new animation.",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "bam": {
                  "type": "string",
                  "default": "boom"
                },
                "bar": {
                  "type": "integer",
                  "default": 42
                },
                "foo": {
                  "type": "boolean",
                  "default": true
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "It works",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    }
  }
}`))
}
