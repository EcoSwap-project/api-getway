package handlers

import (
	item "api_getway/genproto/item_service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Itemhendler interface {
	// Items
	CreateItem(*gin.Context)
	UpdateItem(*gin.Context)
	DeleteItem(*gin.Context)
	GetItem(*gin.Context)
	GetItems(*gin.Context)
	SearchItem(*gin.Context)

	// Swaps
	CreateSwap(*gin.Context)
	UpdateSwap(*gin.Context)
	GetSwaps(*gin.Context)

	// Recycling Centers
	AddRecyclingCentr(*gin.Context)
	GetRecyclingCentrs(*gin.Context)
	AddRecyclingSubmission(*gin.Context)

	// User Ratings
	AddUserRating(*gin.Context)
	GetUserRatings(*gin.Context)

	// Item Categories
	AddItemCategory(*gin.Context)

	// Statistics //
	GetStatistics(*gin.Context)

	// User Activity
	GetUserActivity(*gin.Context)

	// Eco Challenges
	AddEcoChallenge(*gin.Context)
	ParticipateEcoChallenge(*gin.Context)
	UpdateEcoChallengeProgress(*gin.Context)

	// Eco Tips
	AddEcoTip(*gin.Context)
	GetEcoTips(*gin.Context)
}

type itemHendlerIml struct {
	ItemSer item.EcoExchangeServiceClient
}

func NewItemHendler(ItemSer item.EcoExchangeServiceClient) Itemhendler {
	return &itemHendlerIml{ItemSer: ItemSer}
}

// CreateItem handleri
// @Tags Item
// @Summary Create a new item
// @Description Create a new item with details.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.AddItemRequest true "Add item request"
// @Success 200 {object} item_service.AddItemResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /items [post]
func (h *itemHendlerIml) CreateItem(c *gin.Context) {
	var req item.AddItemRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.AddItem(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateItem handleri
// @Tags Item
// @Summary Update an existing item
// @Description Update an existing item with new details.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param itemId path string true "Item ID"
// @Param request body item_service.UpdateItemRequest true "Update item request"
// @Success 200 {object} item_service.UpdateItemResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /items [put]
func (h *itemHendlerIml) UpdateItem(c *gin.Context) {
	var req item.UpdateItemRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.UpdateItem(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteItem handleri
// @Tags Item
// @Summary Delete an existing item
// @Description Delete an existing item by ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.DeleteItemRequest true "Delete item request"
// @Success 200 {object} item_service.DeleteItemResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /items [delete]
func (h *itemHendlerIml) DeleteItem(c *gin.Context) {
	var req item.DeleteItemRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.DeleteItem(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetItem handleri
// @Tags Item
// @Summary Get an existing item
// @Description Get an existing item by ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.GetItemRequest true "Get item request"
// @Success 200 {object} item_service.GetItemResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /items [get]
func (h *itemHendlerIml) GetItem(c *gin.Context) {
	var req item.GetItemRequest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.GetItem(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetItems handleri
// @Tags Item
// @Summary Get all items
// @Description Get all items with pagination and filtering.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} item_service.GetItemsResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /items [get]
func (h *itemHendlerIml) GetItems(c *gin.Context) {
	var req item.GetItemsRequest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.GetItems(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// SearchItem handleri
// @Tags Item
// @Summary Search items
// @Description Search items by keyword.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.SearchItemsRequest true "Search items request"
// @Success 200 {object} item_service.SearchItemsResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /items/search [get]
func (h *itemHendlerIml) SearchItem(c *gin.Context) {
	var req item.SearchItemsRequest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.SearchItems(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CreateSwap handleri
// @Tags Swap
// @Summary Create a new swap
// @Description Create a new swap with details.
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param request body item_service.CreateSwapRequest true "Create swap request"
// @Success 200 {object} item_service.CreateSwapResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /swaps [post]
func (h *itemHendlerIml) CreateSwap(c *gin.Context) {
	var req item.CreateSwapRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.CreateSwap(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateSwap handleri
// @Tags Swap
// @Summary Update an existing swap
// @Description Update an existing swap with new details.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param swapId path string true "Swap ID"
// @Param request body item_service.UpdateSwapStatusRequest true "Update swap request"
// @Success 200 {object} item_service.UpdateSwapStatusResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /swaps [put]
func (h *itemHendlerIml) UpdateSwap(c *gin.Context) {
	var req item.UpdateSwapStatusRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.UpdateSwapStatus(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetSwaps handleri
// @Tags Swap
// @Summary Get all swaps
// @Description Get all swaps with pagination and filtering.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} item_service.GetSwapsResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /swaps [get]
func (h *itemHendlerIml) GetSwaps(c *gin.Context) {
	var req item.GetSwapsRequest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.GetSwaps(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// AddRecyclingCenter handleri
// @Tags recycling center
// @Summary Add a new recycling center
// @Description Add a new recycling center with details.
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param request body item_service.AddRecyclingCenterRequest true "Add recycling center request"
// @Success 200 {object} item_service.AddRecyclingCenterResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /recycling-centers [post]
func (h *itemHendlerIml) AddRecyclingCentr(c *gin.Context) {
	var req item.AddRecyclingCenterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.AddRecyclingCenter(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetRecyclingCenters handleri
// @Tags recycling center
// @Summary Get all recycling centers
// @Description Get all recycling centers with pagination and filtering.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} item_service.GetRecyclingCentersResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /recycling-centers [get]
func (h *itemHendlerIml) GetRecyclingCentrs(c *gin.Context) {
	var req item.GetRecyclingCentersRequest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.GetRecyclingCenters(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// AddRecyclingSubmission handleri
// @Tags recycling center
// @Summary Add a new recycling submission
// @Description Add a new recycling submission with details.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.AddRecyclingSubmissionRequest true "Add recycling submission request"
// @Success 200 {object} item_service.AddRecyclingSubmissionResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /recycling-submissions [post]
func (h *itemHendlerIml) AddRecyclingSubmission(c *gin.Context) {
	var req item.AddRecyclingSubmissionRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.AddRecyclingSubmission(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// AddUserRating handleri
// @Tags User Rating
// @Summary Add a new user rating
// @Description Add a new user rating with details.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.AddUserRatingRequest true "Add user rating request"
// @Success 200 {object} item_service.AddUserRatingResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /user-ratings [post]
func (h *itemHendlerIml) AddUserRating(c *gin.Context) {
	var req item.AddUserRatingRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.AddUserRating(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetUserRatings handleri
// @Tags User Rating
// @Summary Get user ratings
// @Description Get user ratings by user ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.GetUserRatingsRequest true "Get user ratings request"
// @Success 200 {object} item_service.GetUserRatingsResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /user-ratings [get]
func (h *itemHendlerIml) GetUserRatings(c *gin.Context) {
	var req item.GetUserRatingsRequest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.GetUserRatings(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// AddItemCategory handleri
// @Tags Item Category
// @Summary Add a new item category
// @Description Add a new item category with details.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.AddItemCategoryRequest true "Add item category request"
// @Success 200 {object} item_service.AddItemCategoryResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /item-categories [post]
func (h *itemHendlerIml) AddItemCategory(c *gin.Context) {
	var req item.AddItemCategoryRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.AddItemCategory(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetStatistics handleri
// @Tags Statistics
// @Summary Get statistics
// @Description Get statistics about items, swaps, and recycling centers.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} item_service.GetStatisticsResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /statistics [get]
func (h *itemHendlerIml) GetStatistics(c *gin.Context) {
	var req item.GetStatisticsRequest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.GetStatistics(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetUserActivity handleri
// @Tags User Activity
// @Summary Get user activity
// @Description Get user activity by user ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.GetUserActivityRequest true "Get user activity request"
// @Success 200 {object} item_service.GetUserActivityResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /user-activity [get]
func (h *itemHendlerIml) GetUserActivity(c *gin.Context) {
	var req item.GetUserActivityRequest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.GetUserActivity(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// AddEcoChallenge handleri
// @Tags EcoChallange
// @Summary Add a new eco challenge
// @Description Add a new eco challenge with details.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.AddEcoChallengeRequest true "Add eco challenge request"
// @Success 200 {object} item_service.AddEcoChallengeResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /eco-challenges [post]
func (h *itemHendlerIml) AddEcoChallenge(c *gin.Context) {
	var req item.AddEcoChallengeRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.AddEcoChallenge(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ParticipateEcoChallenge handleri
// @Tags EcoChallange
// @Summary Participate in an eco challenge
// @Description Participate in an eco challenge by user ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.ParticipateEcoChallengeRequest true "Participate eco challenge request"
// @Success 200 {object} item_service.ParticipateEcoChallengeResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /eco-challenges/participate [post]
func (h *itemHendlerIml) ParticipateEcoChallenge(c *gin.Context) {
	var req item.ParticipateEcoChallengeRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.ParticipateEcoChallenge(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateEcoChallengeProgress handleri
// @Tags EcoChallange
// @Summary Update eco challenge progress
// @Description Update eco challenge progress by user ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.UpdateEcoChallengeProgressRequest true "Update eco challenge progress request"
// @Success 200 {object} item_service.UpdateEcoChallengeProgressResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /eco-challenges/progress [put]
func (h *itemHendlerIml) UpdateEcoChallengeProgress(c *gin.Context) {
	var req item.UpdateEcoChallengeProgressRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.UpdateEcoChallengeProgress(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// AddEcoTip handleri
// @Tags EcoTip
// @Summary Add a new eco tip
// @Description Add a new eco tip with details.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.AddEcoTipRequest true "Add eco tip request"
// @Success 200 {object} item_service.AddEcoTipResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /eco-tips [post]
func (h *itemHendlerIml) AddEcoTip(c *gin.Context) {
	var req item.AddEcoTipRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.AddEcoTip(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetEcoTips handleri
// @Tags EcoTip
// @Summary Get eco tips
// @Description Get eco tips by user ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body item_service.GetEcoTipsRequest true "Get eco tips request"
// @Success 200 {object} item_service.GetEcoTipsResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /eco-tips [get]
func (h *itemHendlerIml) GetEcoTips(c *gin.Context) {
	var req item.GetEcoTipsRequest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ItemSer.GetEcoTips(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
