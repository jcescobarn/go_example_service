package handlers

import (
	"NIST/repositories"
	"NIST/services"
	"NIST/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppRulesRequestBody struct {
	App_id  int    `json:"app_id"`
	Rule_id string `json:"rule_id"`
}

type SeverityRuleRequest struct {
	Severity string `json:"severity"`
}

func GetAllRules(c *gin.Context) {
	rules, err := services.GetAllRules()
	if err != nil {
		c.JSON(500, gin.H{"error al obtener las reglas": err.Error()})
	}
	c.JSON(200, &rules)
}

func GetAllRulesBySeverity(c *gin.Context) {
	severity := c.Param("severity")
	rules, err := services.GetAllRulesPerSeverity(severity)
	if err != nil {
		c.JSON(500, gin.H{"error al obtener las reglas": err.Error()})
	}
	c.JSON(200, &rules)
}

func AppRuleCreate(c *gin.Context) {
	var body AppRulesRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result, err := repositories.RegisterAppRule(body.Rule_id, body.App_id)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert", "message": err})
		return
	}
	c.JSON(200, gin.H{"rule_id": body.Rule_id, "app_id": body.App_id, "created": result})
	return

}

func GetAppRulesWithoutFixed(c *gin.Context) {

	app_id := c.Param("app_id")
	converted_data, err := strconv.Atoi(app_id)
	app_fixed_rules, err := repositories.GetAppRulesById(converted_data)
	if err != nil {
		c.JSON(500, gin.H{"error al obtener las reglas": err.Error()})
	}
	severity := c.Param("severity")
	rules, err := services.GetAllRulesPerSeverity(severity)
	if err != nil {
		c.JSON(500, gin.H{"error al obtener las reglas": err.Error()})
	}
	fmt.Println(rules)

	data := utils.GetDifference(app_fixed_rules, rules)
	c.JSON(200, data)
}

func GetFixedAppRules(c *gin.Context) {
	app_id := c.Param("app_id")
	converted_data, err := strconv.Atoi(app_id)
	rules, err := repositories.GetAppRulesById(converted_data)
	if err != nil {
		c.JSON(500, gin.H{"error al obtener las reglas": err.Error()})
	}
	var rules_data []utils.RuleResponse
	for _, rule := range rules {
		rule, err := services.GetRule(rule.Rule_id)
		if err != nil {
			c.JSON(500, gin.H{"error al obtener las reglas": err.Error()})
		}
		rules_data = append(rules_data, utils.RuleResponse{
			Id:            rule.Id,
			Description:   rule.Description,
			Base_severity: rule.Base_severity,
		})
	}
	c.JSON(200, &rules_data)
}

func DeleteAppRuleByApp(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("app_id"), 10, 0)
	rule_id := c.Param("rule_id")
	if err != nil {
		c.JSON(500, gin.H{"Error": "Hubo un error con el id ingresado", "message": err.Error})
		return
	}
	result, err := repositories.DeleteAppRule(int(id), rule_id)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Hubo un error con el id ingresado", "message": err.Error})
		return
	}
	c.JSON(200, gin.H{"deleted": result})
	return
}
