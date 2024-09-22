package action

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/fishingSpot"
	"storys-lab-fishing-api/app/utils"
)

type FindOneFishingSpotAction struct {
	uc  usecase.FindOneFishingSpotUseCase
	log logger.Logger
}

type FindAllFishingSpotAction struct {
	uc  usecase.FindAllFishingSpotUseCase
	log logger.Logger
}

// type FindAllFishingSpotAction struct {
// 	uc  usecase.FindAllFishingSpotUseCase
// 	log logger.Logger
// }

func NewFindOneFishingSpotAction(uc usecase.FindOneFishingSpotUseCase, log logger.Logger) FindOneFishingSpotAction {
	return FindOneFishingSpotAction{
		uc:  uc,
		log: log,
	}
}

func NewFindAllFishingSpotAction(uc usecase.FindAllFishingSpotUseCase, log logger.Logger) FindAllFishingSpotAction {
	return FindAllFishingSpotAction{
		uc:  uc,
		log: log,
	}
}

func (t FindOneFishingSpotAction) FindOne(c *gin.Context) {
	const logKey = "find_one_fish"
	fmt.Println("")
	output, err := t.uc.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the fish")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning fish")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}


func (t FindAllFishingSpotAction) FindAllByAreaId(c *gin.Context) {
	const logKey = "find_all_fishing_spot_by_area_id"
	output, err := t.uc.Execute(c.Request.Context(), utils.StrToInt(c.Param("area_id")))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the fish")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning fish")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}
