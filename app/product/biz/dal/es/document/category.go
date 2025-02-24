package document

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/model"
	"github.com/olivere/elastic/v7"
)

func SearchCategory(client *elastic.Client, indexName string, id uint64) ([]model.Product, error) {
	searchService := client.Search().
		Index(indexName).
		Query(elastic.NewNestedQuery("categories", elastic.NewTermQuery("categories.id", id)))

	searchResult, err := searchService.Do(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error searching documents: %v", err)
	}

	var products []model.Product
	for _, hit := range searchResult.Hits.Hits {
		var product model.Product
		err := json.Unmarshal(hit.Source, &product)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling document: %v", err)
		}
		products = append(products, product)
	}

	return products, nil
}

// UpdateCategory 更新所有包含特定 category 的文档
func UpdateCategory(client *elastic.Client, indexName string, newCategory, oldCategory model.Category) {
	products, err := SearchCategory(client, indexName, oldCategory.Id)
	if err != nil {
		klog.Errorf("Error searching documents: %v", err)
		return
	}

	for _, product := range products {
		err := updateProductCategory(client, indexName, product.ID, oldCategory, newCategory)
		if err != nil {
			klog.Errorf("Error updating document with ID %s: %v", product.Name, err)
		} else {
			klog.Debugf("Successfully updated document with ID %s", product.Name)
		}
	}
}

// updateProductCategory 更新单个文档中的 category 字段
func updateProductCategory(client *elastic.Client, indexName string, docID uint, oldCategoty, newCategory model.Category) error {
	updateService := client.Update().
		Index(indexName).
		Id(strconv.FormatUint(uint64(docID), 10)).
		Script(elastic.NewScript(`for (int i = 0; i < ctx._source.categories.size(); i++) {
			if (ctx._source.categories[i].id == params.oldCategory.id) {
				ctx._source.categories[i] = params.newCategory;
			}
		}`).
			Params(map[string]interface{}{
				"oldCategory": oldCategoty,
				"newCategory": newCategory,
			}))

	_, err := updateService.Do(context.Background())
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}

	return nil
}

// DeleteCategory 移除所有包含特定 category 的文档中的指定 category
func DeleteCategory(client *elastic.Client, indexName string, category model.Category) {
	products, err := SearchCategory(client, indexName, category.Id)
	if err != nil {
		klog.Errorf("Error searching documents: %v", err)
		return
	}

	for _, product := range products {
		err := removeCategoryFromProduct(client, indexName, product.ID, category.Id)
		if err != nil {
			klog.Errorf("Error updating document with ID %s: %v", product.Name, err)
		} else {
			klog.Debugf("Successfully updated document with ID %s", product.Name)
		}
	}
}

// removeCategoryFromProduct 从单个文档中移除指定的 category
func removeCategoryFromProduct(client *elastic.Client, indexName string, docID uint, categoryID uint64) error {
	updateService := client.Update().
		Index(indexName).
		Id(strconv.FormatUint(uint64(docID), 10)).
		Script(elastic.NewScript(`ctx._source.categories.removeIf(category -> category.id == params.categoryID)`).
			Param("categoryID", categoryID))

	_, err := updateService.Do(context.Background())
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}

	return nil
}
