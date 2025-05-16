package controllers

import (
	"encoding/csv"
	"net/http"
	"os"
	"registrasi-tamu/config"
	"registrasi-tamu/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func CreateGuest(c *gin.Context) {
	var guest models.Tamu
	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&guest)
	c.JSON(http.StatusOK, guest)
}

func GetGuests(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	name := c.Query("name")

	var guests []models.Tamu
	query := config.DB.Model(&models.Tamu{})

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	var total int64
	query.Count(&total)

	query.Order("created_at desc").Limit(limit).Offset(offset).Find(&guests)

	c.JSON(http.StatusOK, gin.H{
		"data":       guests,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}

func GetGuestsToday(c *gin.Context) {
	var guests []models.Tamu
	today := time.Now().Format("2006-01-02")
	config.DB.Where("DATE(created_at) = ?", today).Find(&guests)
	c.JSON(http.StatusOK, guests)
}

func ExportCSV(c *gin.Context) {
	var guests []models.Tamu
	config.DB.Order("created_at desc").Find(&guests)

	file, err := os.Create("guests.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Name", "Company", "Visiting", "IDCard", "CreatedAt"})
	for _, guest := range guests {
		writer.Write([]string{
			strconv.Itoa(int(guest.ID)), guest.Name, guest.Company, guest.Visiting, guest.IDCard, guest.CreatedAt.Format(time.RFC3339),
		})
	}
	c.File("guests.csv")
}

func ExportPDF(c *gin.Context) {
	var guests []models.Tamu
	config.DB.Order("created_at desc").Find(&guests)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 15)
	pdf.Cell(40, 10, "List Tamu")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 11)
	for _, guest := range guests {
		pdf.CellFormat(0, 10, "| Nama: "+guest.Name+" | ID Card: "+guest.IDCard+" | Asal: "+guest.Company+" | Mengunjungi: "+guest.Visiting+" | Waktu: "+guest.CreatedAt.Format("2006-01-02 15:04")+" |", "0", 1, "", false, 0, "")
	}

	filename := "guests.pdf"
	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.File(filename)
}
