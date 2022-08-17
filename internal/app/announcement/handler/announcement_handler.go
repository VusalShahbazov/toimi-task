package handler

import (
	"github.com/VusalShahbazov/toimi-api/internal/app/announcement"
	"github.com/VusalShahbazov/toimi-api/internal/app/announcement/service"
	"github.com/VusalShahbazov/toimi-api/internal/entity"
	"github.com/VusalShahbazov/toimi-api/internal/helpers"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AnnHandler struct {
	annSrv service.AnnSrv
}

func InitAnnouncementHandler(router *gin.Engine, annSrv service.AnnSrv) {
	h := AnnHandler{annSrv: annSrv}

	router.GET("/announcements", h.GetAnnouncements)
	router.GET("/announcements/:id", h.GetAnnouncement)
	router.POST("/announcements", h.StoreAnnouncement)
}

func (h *AnnHandler) GetAnnouncements(c *gin.Context) {
	var req announcement.AnnSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid input data"})
		return
	}

	if req.SortDesc != "" && !helpers.Contains(announcement.AllowedOrderDesc, req.SortDesc) {
		c.JSON(400, gin.H{"error": "Not allowed sort desc: " + req.SortDesc})
		return
	}

	if req.SortField != "" && !helpers.Contains(announcement.AllowedOrderFields, req.SortField) {
		c.JSON(400, gin.H{"error": "Not allowed sort by: " + req.SortField})
		return
	}

	data, err := h.annSrv.Get(c, req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Something went wrong", "err": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok", "data": data})
}

func (h *AnnHandler) GetAnnouncement(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id param"})
		return
	}

	data, err := h.annSrv.GetById(c, id)
	c.JSON(200, gin.H{"message": "ok", "data": data})
}

type AnnouncementCreate struct {
	Title       string   `json:"title" binding:"required,max=200"`
	Description string   `json:"description" binding:"required,max=1000"`
	Price       float64  `json:"price" binding:"required,min=0"`
	Photos      []string `json:"photos" binding:"omitempty,max=3,min=0,dive,url"`
}

func (h *AnnHandler) StoreAnnouncement(c *gin.Context) {
	var req AnnouncementCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid data", "data": err.Error()})
		return
	}

	ann := entity.Announcement{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
	}

	var pts []entity.AnnouncementPhoto
	for _, v := range req.Photos {
		pts = append(pts, entity.AnnouncementPhoto{
			Url: v,
		})
	}

	ann.Photos = pts

	insert, err := h.annSrv.Insert(c, ann)
	if err != nil {
		c.JSON(400, gin.H{"error": "Something went wrong", "err": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Create", "inserted_id": insert.Id})
}
