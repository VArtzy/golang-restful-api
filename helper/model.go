package helper

import (
    "belajar-go-rest-api/model/domain"
    "belajar-go-rest-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
    return web.CategoryResponse{
        Id: category.Id,
        Name: category.Name,
    }
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
    var categoryResponses []web.CategoryResponse
    for _, category := range categories {
        categoryResponses = append(categoryResponses, ToCategoryResponse(category))
    }
    return categoryResponses
}
