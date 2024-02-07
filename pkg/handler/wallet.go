package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morpheus228/infotecs_golang_test/models"
)

func (h *Handler) createWallet(c *gin.Context) {
	wallet, err := h.services.Wallet.CreateWallet()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wallet)
}

func (h *Handler) makeTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if transaction.Amount <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "amount must be > 0")
		return
	}

	walletId := c.Param("walletId")

	if walletId == transaction.ToWalletId {
		newErrorResponse(c, http.StatusBadRequest, "incoming and outgoing wallet must be different")
	}

	wallet, err := h.services.Wallet.GetWallet(walletId)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			newErrorResponse(c, http.StatusNotFound, "not found outgoing wallet")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if wallet.Balance < transaction.Amount {
		newErrorResponse(c, http.StatusBadRequest, "the required amount is not available on the outgoing wallet")
		return
	}

	err = h.services.Wallet.MakeTransaction(walletId, transaction)

	if err != nil {
		if err.Error() == `pq: insert or update on table \"transactions\" violates foreign key constraint \"transactions_to_wallet_id_fkey\"` {
			newErrorResponse(c, http.StatusBadRequest, "not found incoming wallet")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) getWallet(c *gin.Context) {
	walletId := c.Param("walletId")

	wallet, err := h.services.Wallet.GetWallet(walletId)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			newErrorResponse(c, http.StatusNotFound, "not found this wallet id")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, wallet)
}

func (h *Handler) getWalletHistory(c *gin.Context) {
	walletId := c.Param("walletId")

	_, err := h.services.Wallet.GetWallet(walletId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			newErrorResponse(c, http.StatusNotFound, "not found this wallet id")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	transactions, err := h.services.Wallet.GetWalletHistory(walletId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, transactions)
}
