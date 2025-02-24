package index

import (
	"context"

	"github.com/olivere/elastic/v7"
)

func CreateIndex(client *elastic.Client) (err error) {
	mapping := `
	{
		"settings": {
		  "analysis": {
			"analyzer": {
			  "ik_analyzer": {
				"type": "custom",
				"tokenizer": "ik_max_word"
			  }
			}
		  }
		},
		"mappings": {
		  "properties": {
			"product-name": {
			  "type": "text",
			  "analyzer": "ik_analyzer"
			},
			"product-description": {
			  "type": "text",
			  "analyzer": "ik_analyzer"
			},
			"product-price": {
			  "type": "float"
			},
			"image-name": {
			  "type": "keyword"
			},
			"image-url": {
			  "type": "keyword"
			},
			"categories": {  // 注意这里改为复数形式 categories
			  "type": "nested",
			  "properties": {
				"id": {
				  "type": "long"
				},
				"name": {
				  "type": "text",
				  "analyzer": "ik_analyzer"
				},
				"parent_id": {
				  "type": "long"
				},
				"level": {
				  "type": "long"
				}
			  }
			}
		  }
		}
	  }
	`
	_, err = client.CreateIndex("product").BodyString(mapping).Do(context.Background())
	return err
}
