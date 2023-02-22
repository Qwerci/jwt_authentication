package models

import(
	"gorm.io/gorm"

)

type Performance struct {
	gorm.Model
	First_name 				string `json:"first_name"`
	Middle_name 			string `json:"middle_name"`
	Last_name 				string `json:"last_name"`
	Quality 				uint `json:"quality"`
	Speed 					uint `json:"speed"`
	Tool_proefficiency 		uint `json:"tool_proefficiency"`
	Deliberate_omit   		uint `json:"deliberate_omit"`
	Accuracy          		uint `json:"accuracy"`
	Attention_Details		uint `json:"attention_details"`
	UnAnnotated       		uint `json:"unannotated"`
	Understanding_Concept 	uint `json:"understanding_concept"`
	Count_Assets      		uint `json:"count_assets"`
	Competency 				uint `json:"competency"`
	Sense_Responsibility 	uint `json:"sense_responsibility"`
	Commitment_task    		uint `json:"commitment_task"`
	Ability_initiative  	uint `json:"ability_initiative"`
	Willingness         	uint `json:"willingness"`
	Communication     		uint `json:"communication"`
	RTR 					uint `json:"rtr"`
	Attendance 				uint `json:"attendance"`
	Project_Target_Score 	uint `json:"project_target_score"`
	Total 					uint `json:"total"`
	Month 					string `json:"month"`
	Year 					string `json:"year"`
}







