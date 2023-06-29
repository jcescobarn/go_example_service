package services

import (
	"NIST/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetAllRulesPerSeverity(severity string) ([]utils.RuleResponse, error) {
	init_rules, err := GetRules(severity, "0")
	if err != nil {
		return []utils.RuleResponse{}, err
	}
	result_data := init_rules.Vulnerabilities
	total_pages := init_rules.Total_pages
	for i := 1; i <= total_pages; i++ {
		rules, err := GetRules("HIGH", fmt.Sprintf("%d", i))
		if err != nil {
			return []utils.RuleResponse{}, err
		}
		result_data = append(result_data, rules.Vulnerabilities...)
	}
	return result_data, nil
}

func GetAllRules() ([]utils.RuleResponse, error) {
	init_rules, err := GetRules("", "0")
	if err != nil {
		return []utils.RuleResponse{}, err
	}
	result_data := init_rules.Vulnerabilities
	total_pages := init_rules.Total_pages
	for i := 1; i <= total_pages; i++ {
		rules, err := GetRules("", fmt.Sprintf("%d", i))
		if err != nil {
			return []utils.RuleResponse{}, err
		}
		result_data = append(result_data, rules.Vulnerabilities...)
	}
	return result_data, nil
}

func GetRules(severity string, index string) (utils.RulesResponse, error) {
	// Configurar los parámetros de la solicitud
	fmt.Println(index)
	params := url.Values{
		"startIndex": {index},
	}
	if severity != "" {
		params.Set("cvssV2Severity", severity)
	}

	// Construir la URL completa con los parámetros
	apiURL := "https://services.nvd.nist.gov/rest/json/cves/2.0"
	fullURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	// Realizar la solicitud HTTP GET al endpoint del API
	resp, err := http.Get(fullURL)
	if err != nil {
		return utils.RulesResponse{}, err
	}
	defer resp.Body.Close()

	// Leer la respuesta en formato JSON
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return utils.RulesResponse{}, err
	}

	var object utils.Response
	err = json.Unmarshal(body, &object)

	var arr_data []utils.RuleResponse
	for _, vulnerability := range object.Vulnerabilities {
		cve := vulnerability.Cve

		baseSeverity := ""
		if len(cve.Metrics.CvssMetricV2) > 0 {
			baseSeverity = cve.Metrics.CvssMetricV2[0].BaseSeverity
		}

		rule := utils.RuleResponse{
			Id:            cve.ID,
			Description:   cve.Descriptions[0].Value,
			Base_severity: baseSeverity,
		}
		arr_data = append(arr_data, rule)
	}

	response := utils.RulesResponse{
		Total_in_page:   object.ResultsPerPage,
		Total_pages:     utils.CalculateTotalPages(object.TotalResults, object.ResultsPerPage),
		Total_results:   object.TotalResults,
		Index:           object.StartIndex,
		Vulnerabilities: arr_data,
	}
	return response, nil
}

func GetRule(rule_id string) (utils.RuleResponse, error) {
	// Configurar los parámetros de la solicitud
	params := url.Values{
		"cveid": {rule_id},
	}
	apiURL := "https://services.nvd.nist.gov/rest/json/cves/2.0"
	fullURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	// Realizar la solicitud HTTP GET al endpoint del API
	resp, err := http.Get(fullURL)
	if err != nil {
		return utils.RuleResponse{}, err
	}
	defer resp.Body.Close()

	// Leer la respuesta en formato JSON
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return utils.RuleResponse{}, err
	}

	var object utils.Response
	err = json.Unmarshal(body, &object)

	response := utils.RuleResponse{
		Id:            rule_id,
		Description:   object.Vulnerabilities[0].Cve.Descriptions[0].Value,
		Base_severity: object.Vulnerabilities[0].Cve.Metrics.CvssMetricV2[0].BaseSeverity,
	}

	return response, nil
}
