package ginGenericHttpTemplates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(ctx *gin.Context, reason string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": reason,
	})
}

func InternalServerError(ctx *gin.Context, reason string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": reason,
	})
}

func Ok[T any](ctx *gin.Context, result T) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func OkTemplate(ctx *gin.Context, templateName string, data map[string]any) {
	ctx.HTML(http.StatusOK, templateName, data)
}

func OkOrBadRequest[T any](ctx *gin.Context, result T, err error) {
	if err != nil {
		BadRequest(ctx, err.Error())
		return
	}

	Ok(ctx, result)
}

func OkOrNotFoundTemplate(ctx *gin.Context, okTemplateName string, notFoundTemplateName string, data map[string]any, err error) {
	if err != nil {
		NotFoundTemplate(ctx, notFoundTemplateName)
		return
	}

	OkTemplate(ctx, okTemplateName, data)
}

func NoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func NotFound(ctx *gin.Context, reason string) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": reason,
	})
}

func NotFoundTemplate(ctx *gin.Context, templateName string) {
	ctx.HTML(http.StatusNotFound, templateName, gin.H{
		"message": "Page not found",
	})
}

func UnprocessableEntity(ctx *gin.Context, binder error) {
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": binder.Error(),
	})
}
