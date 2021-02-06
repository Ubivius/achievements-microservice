package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-achievements/data"
)

// DELETE /achievements/{id}
// Deletes a achievement with specified id from the database
func (achievementHandler *AchievementsHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	id := getAchievementId(request)
	achievementHandler.logger.Println("Handle DELETE achievement", id)

	err := data.DeleteAchievement(id)
	if err == data.ErrorAchievementNotFound {
		achievementHandler.logger.Println("[ERROR] deleting, id does not exist")
		http.Error(responseWriter, "Achievement not found", http.StatusNotFound)
		return
	}

	if err != nil {
		achievementHandler.logger.Println("[ERROR] deleting achievement", err)
		http.Error(responseWriter, "Erro deleting achievement", http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}