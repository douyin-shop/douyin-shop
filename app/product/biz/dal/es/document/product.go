package document

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/product/biz/code"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/model"
	"github.com/olivere/elastic/v7"
)

func CreateProduct(client *elastic.Client, indexName string, product model.Product) error {
	_, err := client.Index().
		Index(indexName).
		Id(strconv.FormatUint(uint64(product.ID), 10)). // 使用产品名称作为 ID，可以根据需要调整
		BodyJson(product).
		Do(context.Background())

	if err != nil {
		return fmt.Errorf("error creating document: %v", err)
	}
	klog.Debugf("Successfully created document with ID %s", product.Name)
	return nil
}

func UpdateProduct(client *elastic.Client, indexName string, OldProduct, NewProduct model.Product) error {
	updateService := client.Update().
		Index(indexName).
		Id(strconv.FormatUint(uint64(OldProduct.ID), 10)).
		Doc(map[string]interface{}{
			"product-name":        NewProduct.Name,
			"product-description": NewProduct.Description,
			"product-price":       NewProduct.Price,
			"image-name":          NewProduct.ImageName,
			"image-url":           NewProduct.ImageURL,
			"categories":          NewProduct.Category,
		})

	_, err := updateService.Do(context.Background())
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}

	klog.Debugf("Successfully updated document with ID %s", OldProduct.Name)
	return nil
}

func DeleteProduct(client *elastic.Client, indexName string, docID uint) error {
	deleteService := client.Delete().
		Index(indexName).
		Id(strconv.FormatUint(uint64(docID), 10))

	_, err := deleteService.Do(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting document: %v", err)
	}

	klog.Debugf("Successfully deleted document with ID %s", docID)
	return nil
}

//搜索商品(关键词模糊匹配与分类精确匹配)

// FuzzySearchProduct 关键词模糊匹配
func FuzzySearchProduct(client *elastic.Client, indexName string, keyword string, pageNum, pageSize int) ([]model.Product, error) {
	searchService := client.Search().Index(indexName).
		Query(elastic.NewMultiMatchQuery(keyword, "product-name", "product-description").
			Type("best_fields").
			Analyzer("ik_max_word"))

	// 如果 pageNum 和 pageSize 都不为 0，则启用分页
	if pageNum > 0 && pageSize > 0 {
		from := (pageNum - 1) * pageSize // 计算起始位置
		searchService = searchService.From(from).Size(pageSize)
	}

	searchResult, err := searchService.Do(context.Background())
	if err != nil {
		return nil, err
	}

	var products []model.Product
	for _, hit := range searchResult.Hits.Hits {
		var product model.Product
		err := json.Unmarshal(hit.Source, &product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	klog.Debugf("Successfully fuzzy searched for keyword: %s with pageNum: %d and pageSize: %d", keyword, pageNum, pageSize)
	return products, nil
}

// 分类精确匹配(支持分类和价格区间)
func ExactSearchProduct(client *elastic.Client, indexName string, categoryName string, minPrice, maxPrice float64, pageNum, pageSize int) ([]model.Product, int) {
	boolQuery := elastic.NewBoolQuery()

	if categoryName != "" {
		boolQuery = boolQuery.Must(elastic.NewNestedQuery("categories", elastic.NewTermQuery("categories.name", categoryName)))
	}

	if minPrice > 0 && maxPrice > 0 {
		boolQuery = boolQuery.Must(elastic.NewRangeQuery("product-price").Gte(minPrice).Lte(maxPrice))
	} else if minPrice > 0 {
		boolQuery = boolQuery.Must(elastic.NewRangeQuery("product-price").Gte(minPrice))
	} else if maxPrice > 0 {
		boolQuery = boolQuery.Must(elastic.NewRangeQuery("product-price").Lte(maxPrice))
	}

	searchService := client.Search().Index(indexName).Query(boolQuery)

	// 如果 pageNum 和 pageSize 都不为 0，则启用分页
	if pageNum > 0 && pageSize > 0 {
		from := (pageNum - 1) * pageSize // 计算起始位置
		searchService = searchService.From(from).Size(pageSize)
	}

	searchResult, err := searchService.Do(context.Background())
	if err != nil {
		return nil, code.Error
	}

	var products []model.Product
	for _, hit := range searchResult.Hits.Hits {
		var product model.Product
		err := json.Unmarshal(hit.Source, &product)
		if err != nil {
			return nil, code.Error
		}
		products = append(products, product)
	}

	klog.Debugf("Successfully exact searched with category: %s, minPrice: %.2f, maxPrice: %.2f, pageNum: %d, pageSize: %d", categoryName, minPrice, maxPrice, pageNum, pageSize)
	return products, code.Success
}

func CombinedSearchProduct(client *elastic.Client, indexName string, keyword string, categoryName string, minPrice, maxPrice float64, pageNum, pageSize int) ([]model.Product, error) {
	var products []model.Product

	// 模糊搜索关键词不为空时，进行模糊搜索
	if keyword != "" {
		fuzzyProducts, err := FuzzySearchProduct(client, indexName, keyword, pageNum, pageSize)
		if err != nil {
			return nil, err
		}
		products = fuzzyProducts
	}

	// 精确匹配分类名称不为空时，进行精确搜索或对模糊搜索的结果进一步过滤
	if categoryName != "" || minPrice > 0 || maxPrice > 0 {
		boolQuery := elastic.NewBoolQuery()

		if categoryName != "" {
			boolQuery = boolQuery.Must(elastic.NewNestedQuery("categories", elastic.NewTermQuery("categories.name", categoryName)))
		}

		if minPrice > 0 && maxPrice > 0 {
			boolQuery = boolQuery.Must(elastic.NewRangeQuery("product-price").Gte(minPrice).Lte(maxPrice))
		} else if minPrice > 0 {
			boolQuery = boolQuery.Must(elastic.NewRangeQuery("product-price").Gte(minPrice))
		} else if maxPrice > 0 {
			boolQuery = boolQuery.Must(elastic.NewRangeQuery("product-price").Lte(maxPrice))
		}

		searchService := client.Search().Index(indexName).Query(boolQuery)

		// 如果进行了模糊搜索，这里需要使用post filter来进行过滤
		if keyword != "" {
			searchService = searchService.PostFilter(boolQuery)
		}

		// 如果 pageNum 和 pageSize 都不为 0，则启用分页
		if pageNum > 0 && pageSize > 0 {
			from := (pageNum - 1) * pageSize // 计算起始位置
			searchService = searchService.From(from).Size(pageSize)
		}

		searchResult, err := searchService.Do(context.Background())
		if err != nil {
			return nil, err
		}

		products = make([]model.Product, 0)
		for _, hit := range searchResult.Hits.Hits {
			var product model.Product
			err := json.Unmarshal(hit.Source, &product)
			if err != nil {
				return nil, err
			}
			products = append(products, product)
		}
	}
	return products, nil
}
